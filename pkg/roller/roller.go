package roller

import (
	"context"
	"fmt"
	"time"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	v1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/aws-pod-identity-webhook/pkg/types"
)

type Config struct {
	CtrlClient client.Client
	Logger     micrologger.Logger
}

type Roller struct {
	ctrlClient client.Client
	logger     micrologger.Logger
}

func New(config Config) (*Roller, error) {
	if config.CtrlClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.CtrlClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	return &Roller{
		ctrlClient: config.CtrlClient,
		logger:     config.Logger,
	}, nil
}

func (r *Roller) Roll(ctx context.Context, rollable types.Rollable) error {
	r.logger.Debugf(ctx, "Rolling %s %s/%s", rollable.Kind, rollable.Namespace, rollable.Name)
	switch rollable.Kind {
	case "Deployment":
		return r.rollDeployment(ctx, rollable)
	case "StatefulSet":
		return r.rollStatefulSet(ctx, rollable)
	case "DaemonSet":
		return r.rollDaemonSet(ctx, rollable)
	default:
		return microerror.Maskf(unsupportedKindError, fmt.Sprintf("Unsupported Kind %s", rollable.Kind))
	}
}

func (r *Roller) rollDeployment(ctx context.Context, rollable types.Rollable) error {
	deployment := v1.Deployment{}
	err := r.ctrlClient.Get(ctx, client.ObjectKey{Name: rollable.Name, Namespace: rollable.Namespace}, &deployment)
	if err != nil {
		return microerror.Mask(err)
	}

	if deployment.Spec.Template.Annotations == nil {
		deployment.Spec.Template.Annotations = map[string]string{}
	}
	deployment.Spec.Template.Annotations["restarted-by-aws-pod-identity-webhook"] = fmt.Sprint(time.Now().UnixNano())

	err = r.ctrlClient.Update(ctx, &deployment)
	if err != nil {
		return microerror.Mask(err)
	}

	r.logger.Debugf(ctx, "Rolled Deployment %s/%s", rollable.Namespace, rollable.Name)

	return nil
}

func (r *Roller) rollStatefulSet(ctx context.Context, rollable types.Rollable) error {
	statefulset := v1.StatefulSet{}
	err := r.ctrlClient.Get(ctx, client.ObjectKey{Name: rollable.Name, Namespace: rollable.Namespace}, &statefulset)
	if err != nil {
		return microerror.Mask(err)
	}

	if statefulset.Spec.Template.Annotations == nil {
		statefulset.Spec.Template.Annotations = map[string]string{}
	}
	statefulset.Spec.Template.Annotations["restarted-by-aws-pod-identity-webhook"] = fmt.Sprint(time.Now().UnixNano())

	err = r.ctrlClient.Update(ctx, &statefulset)
	if err != nil {
		return microerror.Mask(err)
	}

	r.logger.Debugf(ctx, "Rolled Statefulset %s/%s", rollable.Namespace, rollable.Name)

	return nil
}

func (r *Roller) rollDaemonSet(ctx context.Context, rollable types.Rollable) error {
	daemonset := v1.DaemonSet{}
	err := r.ctrlClient.Get(ctx, client.ObjectKey{Name: rollable.Name, Namespace: rollable.Namespace}, &daemonset)
	if err != nil {
		return microerror.Mask(err)
	}

	if daemonset.Spec.Template.Annotations == nil {
		daemonset.Spec.Template.Annotations = map[string]string{}
	}
	daemonset.Spec.Template.Annotations["restarted-by-aws-pod-identity-webhook"] = fmt.Sprint(time.Now().UnixNano())

	err = r.ctrlClient.Update(ctx, &daemonset)
	if err != nil {
		return microerror.Mask(err)
	}

	r.logger.Debugf(ctx, "Rolled DaemonSet %s/%s", rollable.Namespace, rollable.Name)

	return nil
}
