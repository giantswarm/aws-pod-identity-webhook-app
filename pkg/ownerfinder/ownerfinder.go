package ownerfinder

import (
	"context"
	"fmt"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	v1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Config struct {
	CtrlClient client.Client
	Logger     micrologger.Logger
}

type OwnerFinder struct {
	ctrlClient client.Client
	logger     micrologger.Logger
}

func New(config Config) (*OwnerFinder, error) {
	if config.CtrlClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.CtrlClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	return &OwnerFinder{
		ctrlClient: config.CtrlClient,
		logger:     config.Logger,
	}, nil
}

func (of *OwnerFinder) FindOwner(ctx context.Context, pod corev1.Pod) (*metav1.OwnerReference, error) {
	of.logger.Debugf(ctx, "Looking up owner of pod %s/%s", pod.Namespace, pod.Name)

	// Iterate up the ownerReference chain until we find an object whose ownerReference is nil.
	var lastOwner *metav1.OwnerReference
	owners := pod.GetOwnerReferences()
	for len(owners) > 0 {
		var obj client.Object
		switch owners[0].Kind {
		case "ReplicaSet":
			obj = &v1.ReplicaSet{}
		case "Deployment":
			obj = &v1.Deployment{}
		case "DaemonSet":
			obj = &v1.DaemonSet{}
		case "StatefulSet":
			obj = &v1.StatefulSet{}
		case "Job":
			obj = &batchv1.Job{}
		default:
			return nil, microerror.Maskf(unsupportedOwnerKindError, fmt.Sprintf("Unsupported Kind %s", owners[0].Kind))
		}
		err := of.ctrlClient.Get(ctx, client.ObjectKey{Namespace: pod.Namespace, Name: owners[0].Name}, obj)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		lastOwner = &owners[0]
		owners = obj.GetOwnerReferences()
	}

	if lastOwner == nil {
		of.logger.Debugf(ctx, "Pod %s/%s does not seem to have any owner", pod.Namespace, pod.Name)
		return nil, nil
	}

	of.logger.Debugf(ctx, "Pod %s/%s is ultimately owned by %s %s", pod.Namespace, pod.Name, lastOwner.Kind, lastOwner.Name)

	return lastOwner, nil
}
