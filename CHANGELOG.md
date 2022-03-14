# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.3.2...HEAD
[0.3.2]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/giantswarm/aws-pod-identity-webhook/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/giantswarm/aws-pod-identity-webhook/releases/tag/v0.3.0
