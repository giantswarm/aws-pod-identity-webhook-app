[![CircleCI](https://circleci.com/gh/giantswarm/aws-pod-identity-webhook.svg?style=shield)](https://circleci.com/gh/giantswarm/aws-pod-identity-webhook)

# aws-pod-identity-webhook chart

Helm Chart for AWS Pod Identity Webhook in Workload Clusters.

* Installs the the [amazon-eks-pod-identity-webhook](https://github.com/aws/amazon-eks-pod-identity-webhook).

# Deployment

Managed by the Giant Swarm [App Platform](https://docs.giantswarm.io/app-platform/).

# Configuration Options

- All configuration options are documented in the [values.yaml](/helm/aws-pod-identity-webhook/values.yaml) file.

# For developers

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/aws-pod-identity-webhook.git
$ cd aws-pod-identity-webhook
$ helm install helm/aws-pod-identity-webhook
```

Provide a custom `values.yaml`:

```bash
$ helm install aws-pod-identity-webhook -f values.yaml
```
