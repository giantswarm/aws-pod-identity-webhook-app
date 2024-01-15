# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Configure `gsoci.azurecr.io` as the default container image registry.
- Use `policy/v1` for `PodDisruptionBudget` (enabled since k8s 1.21).

## [1.13.2] - 2023-10-18

### Fixed

- Disable PSP for restarter component when PSS is enforced.

## [1.13.1] - 2023-10-18

### Fixed

- Don't add `use` permission on `PSP` when PSPs are disabled.

## [1.13.0] - 2023-10-18

### Added

- Add `global.podSecurityStandards.enforced` value for PSS migration.

## [1.12.0] - 2023-10-03

### Changed

- Make PSP rendering conditional for 1.25+ compatibility

### Fixed

- In case it's not possible to determine the owner of a pod, log and continue rather than panic.
- Fix `service`'s selector to not include the restarer job during prometheus scraping.

## [1.11.1] - 2023-07-31

### Fixed

- Use `topologySpreadConstraints` instead of `podAntiAffinity` to spread deployment replicas across nodes.

## [1.11.0] - 2023-07-26

### Changed

- Use patch instead of update when restarting pods.

### Added

- Add support for restarting `Jobs`.

## [1.10.0] - 2023-07-13

### Fixed

- Add required values for pss policies.

## [1.9.1] - 2023-06-28

### Added

- Add `cluster autoscaler safe to evict` annotation to restarter pods

## [1.9.0] - 2023-05-09

### Changed

- Increase replicas to `3` and include update strategy.

## [1.8.2] - 2023-05-02

### Fixed

- Fix indentation for `seccompProfile` in `Deployment` Helm manifest.

## [1.8.1] - 2023-05-01

### Fixed

- Remove duplicated labels in Helm manifests.

## [1.8.0] - 2023-04-19

### Changed

- Change default registry in Helm chart from quay.io to docker.io.

## [1.7.1] - 2023-04-05

### Changed

- Specify `group` in the `Certificate`'s `issuerRef` to avoid renewal issues in certain clusters.

## [1.7.0] - 2023-03-15

### Changed

- Use `image.registry` value as registry domain for the `restarter` image.

## [1.6.3] - 2023-02-20

### Added

- Add service name ports

### Fixed

- Remove duplicate labels definition in service
- Group all monitoring annotations

## [1.6.2] - 2023-02-17

### Fixed

- Added `projected` and `secret` volumes to PodSecurityPolicy.

## [1.6.1] - 2023-02-17

### Fixed

- Added runAsUser: 1000 to securityContext.

## [1.6.0] - 2023-02-17

### Changed

- Added the use of the runtime/default seccomp profile.

## [1.5.0] - 2023-02-02

### Changed

- Update image to the latest version of our fork `v0.2.0`.

## [1.4.0] - 2023-01-12

- Added selecting image image based on provider
- Added `giantswarm/amazon-eks-pod-identity-webhook-gs:v0.1.0` our custom fork to which adds the aws account id to the service account `role-arn` annotation

## [1.3.0] - 2023-01-05

- Undo image update to `giantswarm/amazon-eks-pod-identity-webhook-gs:v0.1.0` our custom fork to which adds the aws account id to the service account `role-arn` annotation

## [1.2.0] - 2023-01-05

### Added

- Updated image to `giantswarm/amazon-eks-pod-identity-webhook-gs:v0.1.0` our custom fork to which adds the aws account id to the service account `role-arn` annotation

## [1.1.1] - 2022-12-08

### Fixed

- PodDisruptionBudget incorrectly matching the completed `restarter` Job pod.

## [1.1.0] - 2022-11-10

### Changed

- Use regional sts endpoint by default.

## [1.0.0] - 2022-11-07

### Added

- Added Pod Restarter `CronJob` that automatically restarts pods using IRSA that don't have the needed configurations set.

## [0.10.0] - 2022-09-20

### Changed

- Set `--in-cluster` to false because it's deprecated.

## [0.9.0] - 2022-07-15

### Changed

- Change webhook port to `9443`
- Update image to the latest version `v0.4.0`.

## [0.8.1] - 2022-05-19

### Fixed

- Allow pod-identity-webhook to inject volumes in `kube-system` namespace.

## [0.8.0] - 2022-04-21

### Changed

- Adjust helm values.

## [0.7.0] - 2022-04-21

### Added

- Add annotation config.giantswarm.io.

## [0.6.0] - 2022-04-21

### Changed

- Set token audience as value.

## [0.5.0] - 2022-03-22

### Added

- Add VerticalPodAutoscaler CR.

## [0.4.0] - 2022-03-14

### Added

- Support for forcing the webhook to run on pods in kube-system with specific label

## [0.3.2] - 2022-03-10

### Fixed

- Avoid running mutating webhook against it's own pods

## [0.3.1] - 2022-03-03

### Fixed

- Certificate name and DNS name variant

## [0.3.0] - 2022-03-01

### Added

- Initial chart release

[Unreleased]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.13.2...HEAD
[1.13.2]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.13.1...v1.13.2
[1.13.1]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.13.0...v1.13.1
[1.13.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.12.0...v1.13.0
[1.12.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.11.1...v1.12.0
[1.11.1]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.11.0...v1.11.1
[1.11.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.10.0...v1.11.0
[1.10.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.9.1...v1.10.0
[1.9.1]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.9.0...v1.9.1
[1.9.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.8.2...v1.9.0
[1.8.2]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.8.1...v1.8.2
[1.8.1]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.8.0...v1.8.1
[1.8.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.7.1...v1.8.0
[1.7.1]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.7.0...v1.7.1
[1.7.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.6.3...v1.7.0
[1.6.3]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.6.2...v1.6.3
[1.6.2]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.6.1...v1.6.2
[1.6.1]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.6.0...v1.6.1
[1.6.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.5.0...v1.6.0
[1.5.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.4.0...v1.5.0
[1.4.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.1.1...v1.2.0
[1.1.1]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.10.0...v1.0.0
[0.10.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.8.1...v0.9.0
[0.8.1]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.8.0...v0.8.1
[0.8.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.3.2...v0.4.0
[0.3.2]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/giantswarm/aws-pod-identity-webhook/releases/tag/v0.3.0
