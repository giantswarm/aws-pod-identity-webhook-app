[![CircleCI](https://circleci.com/gh/giantswarm/aws-pod-identity-webhook.svg?style=shield)](https://circleci.com/gh/giantswarm/aws-pod-identity-webhook)

# aws-pod-identity-webhook chart and custom extension

This repository provides two components:

1. Helm Chart for AWS Pod Identity Webhook in Workload Clusters.

* Installs the the [amazon-eks-pod-identity-webhook](https://github.com/aws/amazon-eks-pod-identity-webhook).

This webhook is for mutating pods that will require AWS IAM access.

2. Custom golang binary that runs as a cronjob

This is used to restart pods that weren't catched by the webhook (because they started before the webhook).

# Deployment

Managed by the Giant Swarm [App Platform](https://docs.giantswarm.io/app-platform/).

# Configuration Options

- All configuration options are documented in the [values.yaml](/helm/aws-pod-identity-webhook/values.yaml) file.

# Usage

See the [amazon-eks-pod-identity-webhook walkthrough](https://github.com/aws/amazon-eks-pod-identity-webhook#eks-walkthrough)

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
