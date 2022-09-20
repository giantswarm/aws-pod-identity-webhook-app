# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.10.0...HEAD
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
