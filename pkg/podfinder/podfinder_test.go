package podfinder

import (
	"context"
	"testing"

	"github.com/giantswarm/micrologger/microloggertest"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
