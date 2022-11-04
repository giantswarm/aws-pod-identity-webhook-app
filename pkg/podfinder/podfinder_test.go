package podfinder

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/giantswarm/micrologger/microloggertest"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/aws-pod-identity-webhook/pkg/unittest"
)

func TestPodFinder_getServiceAccountsWithIRSAEnabled(t *testing.T) {
	saWithIrsa := v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod1",
			Namespace: "default",
			Annotations: map[string]string{
				"eks.amazonaws.com/role-arn": "changeme",
			},
		},
	}

	saWithoutIrsa := v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "pod1",
			Namespace:   "default",
			Annotations: nil,
		},
	}

	tests := []struct {
		name     string
		existing []v1.ServiceAccount
		want     []v1.ServiceAccount
		wantErr  bool
	}{
		{
			name:     "No service accounts at all",
			existing: []v1.ServiceAccount{},
			want:     []v1.ServiceAccount{},
			wantErr:  false,
		},
		{
			name: "One service account with no annotations",
			existing: []v1.ServiceAccount{
				saWithoutIrsa,
			},
			want:    []v1.ServiceAccount{},
			wantErr: false,
		},
		{
			name: "One service account with IRSA annotations",
			existing: []v1.ServiceAccount{
				saWithIrsa,
			},
			want: []v1.ServiceAccount{
				saWithIrsa,
			},
			wantErr: false,
		},
		{
			name: "Two service accounts, one with and one without IRSA annotations",
			existing: []v1.ServiceAccount{
				saWithIrsa,
				saWithoutIrsa,
			},
			want: []v1.ServiceAccount{
				saWithIrsa,
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := unittest.FakeK8sClient()

			for _, sa := range tt.existing {
				_ = client.CtrlClient().Create(ctx, &sa)
			}

			p := &PodFinder{
				ctrlClient: client.CtrlClient(),
				logger:     microloggertest.New(),
			}
			got, err := p.getServiceAccountsWithIRSAEnabled(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("getServiceAccountsWithIRSAEnabled(%s) error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			gotNames := []string{}
			wantNames := []string{}
			for _, sa := range got {
				gotNames = append(gotNames, fmt.Sprintf("%s/%s", sa.Namespace, sa.Name))
			}
			for _, sa := range tt.want {
				wantNames = append(wantNames, fmt.Sprintf("%s/%s", sa.Namespace, sa.Name))
			}
			if !reflect.DeepEqual(gotNames, wantNames) {
				t.Errorf("getServiceAccountsWithIRSAEnabled(%s) got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestPodFinder_getPodsUsingServiceAccount(t *testing.T) {
	namespace := "default"

	tests := []struct {
		name     string
		existing []v1.Pod
		saName   string
		want     []v1.Pod
		wantErr  bool
	}{
		{
			name:     "No pods at all",
			existing: []v1.Pod{},
			saName:   "",
			want:     []v1.Pod{},
			wantErr:  false,
		},
		{
			name: "One pod using a different SA",
			existing: []v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod1",
						Namespace: namespace,
					},
					Spec: v1.PodSpec{
						ServiceAccountName: "my-service-account-name",
					},
				},
			},
			saName:  "target-service-account-name",
			want:    []v1.Pod{},
			wantErr: false,
		},
		{
			name: "One pod using desired SA",
			existing: []v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod1",
						Namespace: namespace,
					},
					Spec: v1.PodSpec{
						ServiceAccountName: "my-service-account-name",
					},
				},
			},
			saName: "my-service-account-name",
			want: []v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod1",
						Namespace: namespace,
					},
				},
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := unittest.FakeK8sClient()

			for _, pod := range tt.existing {
				_ = client.CtrlClient().Create(ctx, &pod)
			}

			p := &PodFinder{
				ctrlClient: client.CtrlClient(),
				logger:     microloggertest.New(),
			}
			got, err := p.getPodsUsingServiceAccount(ctx, v1.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      tt.saName,
					Namespace: namespace,
				},
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("getPodsUsingServiceAccount(%s) error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			gotNames := []string{}
			wantNames := []string{}
			for _, pod := range got {
				// This is unfortunately needed as it seems `MatchingFields` does not work with fake ctrlclient.
				if pod.Spec.ServiceAccountName == tt.saName {
					gotNames = append(gotNames, fmt.Sprintf("%s/%s", pod.Namespace, pod.Name))
				}
			}
			for _, pod := range tt.want {
				wantNames = append(wantNames, fmt.Sprintf("%s/%s", pod.Namespace, pod.Name))
			}
			if !reflect.DeepEqual(gotNames, wantNames) {
				t.Errorf("getPodsUsingServiceAccount(%s) got = %v, want %v", tt.name, gotNames, wantNames)
			}
		})
	}
}

func TestPodFinder_needsToBeRecreated(t *testing.T) {
	tests := []struct {
		name    string
		pod     v1.Pod
		want    bool
		wantErr bool
	}{
		{
			name: "Pod with existing volume",
			pod: v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: v1.PodSpec{
					Volumes: []v1.Volume{
						{
							Name: "aws-iam-token",
						},
					},
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Pod with existing volume not in first position",
			pod: v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: v1.PodSpec{
					Volumes: []v1.Volume{
						{
							Name: "somethingelse",
						},
						{
							Name: "aws-iam-token",
						},
					},
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Pod with no existing volume",
			pod: v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod1",
					Namespace: "default",
				},
				Spec: v1.PodSpec{
					Volumes: []v1.Volume{
						{
							Name: "somethingelse",
						},
					},
				},
			},
			want:    true,
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PodFinder{
				ctrlClient: nil,
				logger:     microloggertest.New(),
			}
			got, err := p.needsToBeRecreated(ctx, tt.pod)
			if (err != nil) != tt.wantErr {
				t.Errorf("needsToBeRecreated(%s) error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("needsToBeRecreated(%s) got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
