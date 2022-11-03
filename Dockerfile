# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
ADD ./aws-pod-identity-restarter /aws-pod-identity-restarter
USER 65532:65532

ENTRYPOINT ["/aws-pod-identity-restarter"]
