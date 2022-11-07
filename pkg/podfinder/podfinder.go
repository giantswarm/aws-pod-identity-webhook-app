package podfinder

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	coreV1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	irsaVolumeName               = "aws-iam-token"
	irsaServiceAccountAnnotation = "eks.amazonaws.com/role-arn"
)

type Config struct {
	CtrlClient client.Client
	Logger     micrologger.Logger
}

type PodFinder struct {
	ctrlClient client.Client
	logger     micrologger.Logger
}

func New(config Config) (*PodFinder, error) {
	if config.CtrlClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.CtrlClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	return &PodFinder{
		ctrlClient: config.CtrlClient,
		logger:     config.Logger,
	}, nil
}

func (p *PodFinder) FindPodsToBeTerminated(ctx context.Context) ([]coreV1.Pod, error) {
	ret := make([]coreV1.Pod, 0)
	sas, err := p.getServiceAccountsWithIRSAEnabled(ctx)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	// Get pods using this service account.
	for _, sa := range sas {
		pods, err := p.getPodsUsingServiceAccount(ctx, sa)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		for _, pod := range pods {
			needs, err := p.needsToBeRecreated(ctx, pod)
			if err != nil {
				return nil, microerror.Mask(err)
			}
			if needs {
				ret = append(ret, pod)
			}
		}
	}

	return ret, nil
}

func (p *PodFinder) getServiceAccountsWithIRSAEnabled(ctx context.Context) ([]coreV1.ServiceAccount, error) {
	p.logger.Debugf(ctx, "Finding all ServiceAccounts with IRSA annotation %q", irsaServiceAccountAnnotation)
	serviceaccounts := coreV1.ServiceAccountList{}
	err := p.ctrlClient.List(ctx, &serviceaccounts)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	ret := make([]coreV1.ServiceAccount, 0)

	for _, sa := range serviceaccounts.Items {
		if _, ok := sa.Annotations[irsaServiceAccountAnnotation]; ok {
			ret = append(ret, sa)
		}
	}

	p.logger.Debugf(ctx, "Found %d ServiceAccounts", len(ret))

	return ret, nil
}

func (p *PodFinder) getPodsUsingServiceAccount(ctx context.Context, sa coreV1.ServiceAccount) ([]coreV1.Pod, error) {
	p.logger.Debugf(ctx, "Looking up pods using service account %s in namespace %s", sa.Name, sa.Namespace)
	podList := coreV1.PodList{}
	err := p.ctrlClient.List(ctx, &podList, client.InNamespace(sa.Namespace), client.MatchingFields{
		"spec.serviceAccountName": sa.Name,
	})
	if err != nil {
		return nil, microerror.Mask(err)
	}

	p.logger.Debugf(ctx, "Found %d pods using service account %s in namespace %s", len(podList.Items), sa.Name, sa.Namespace)

	return podList.Items, nil
}

func (p *PodFinder) needsToBeRecreated(ctx context.Context, pod coreV1.Pod) (bool, error) {
	p.logger.Debugf(ctx, "Checking if pod %s/%s needs to be recreated", pod.Namespace, pod.Name)
	// Check if pod has needed volume.
	for _, v := range pod.Spec.Volumes {
		if v.Name == irsaVolumeName {
			p.logger.Debugf(ctx, "Checking if pod %s/%s does NOT need to be recreated", pod.Namespace, pod.Name)
			// volume present, pod does not need to be recreated.
			return false, nil
		}
	}

	p.logger.Debugf(ctx, "Checking if pod %s/%s needs to be recreated", pod.Namespace, pod.Name)

	return true, nil
}
