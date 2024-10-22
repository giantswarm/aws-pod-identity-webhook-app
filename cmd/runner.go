package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/giantswarm/k8sclient/v8/pkg/k8sclient"
	"github.com/giantswarm/k8sclient/v8/pkg/k8srestconfig"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"

	"github.com/giantswarm/aws-pod-identity-webhook/pkg/ownerfinder"
	"github.com/giantswarm/aws-pod-identity-webhook/pkg/podfinder"
	"github.com/giantswarm/aws-pod-identity-webhook/pkg/roller"
	"github.com/giantswarm/aws-pod-identity-webhook/pkg/types"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate(cmd)
	if err != nil {
		return err
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error
	var restConfig *rest.Config
	{
		c := k8srestconfig.Config{
			Logger:    r.logger,
			InCluster: true,
		}

		restConfig, err = k8srestconfig.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	c := k8sclient.ClientsConfig{
		SchemeBuilder: k8sclient.SchemeBuilder{
			corev1.AddToScheme,
		},
		Logger:     r.logger,
		RestConfig: restConfig,
	}
	k8sClient, err := k8sclient.NewClients(c)
	if err != nil {
		return microerror.Mask(err)
	}

	podFinder, err := podfinder.New(podfinder.Config{
		CtrlClient: k8sClient.CtrlClient(),
		Logger:     r.logger,
	})
	if err != nil {
		return microerror.Mask(err)
	}

	ownerFinder, err := ownerfinder.New(ownerfinder.Config{
		CtrlClient: k8sClient.CtrlClient(),
		Logger:     r.logger,
	})
	if err != nil {
		return microerror.Mask(err)
	}

	rollerObj, err := roller.New(roller.Config{
		CtrlClient: k8sClient.CtrlClient(),
		Logger:     r.logger,
	})
	if err != nil {
		return microerror.Mask(err)
	}

	// Get list of pods that need rolling.
	pods, err := podFinder.FindPodsToBeTerminated(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	// Get unique list of deployments/replicaset/daemonsets that need rolling.
	rollables := map[string]types.Rollable{}
	for _, pod := range pods {
		owner, err := ownerFinder.FindOwner(ctx, pod)
		if err != nil {
			r.logger.Debugf(ctx, "error finding owner for pod %s/%s: %s", pod.Namespace, pod.Name, microerror.Cause(err))
			continue
		}

		if owner != nil {
			key := fmt.Sprintf("%s-%s-%s", owner.Kind, pod.Namespace, owner.Name)

			rollables[key] = types.Rollable{
				Kind:      owner.Kind,
				Name:      owner.Name,
				Namespace: pod.Namespace,
			}
		}
	}

	// Roll all deployments/replicaset/daemonsets.
	for _, rollable := range rollables {
		err = rollerObj.Roll(ctx, rollable)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
