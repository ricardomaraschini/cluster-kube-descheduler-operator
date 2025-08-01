FROM brew.registry.redhat.io/rh-osbs/openshift-golang-builder:rhel_9_1.24 as builder
WORKDIR /go/src/github.com/openshift/cluster-kube-descheduler-operator
COPY . .

RUN mkdir licenses
COPY ./LICENSE licenses/.

ARG OPERATOR_IMAGE=registry.redhat.io/kube-descheduler-operator/kube-descheduler-rhel9-operator@sha256:f7675a6efa5941b3c0e2411d223b391b34d892cbada9c6d40c55b6603337b1ba
ARG OPERAND_IMAGE=registry.redhat.io/kube-descheduler-operator/descheduler-rhel9@sha256:4257096daf40a3d90c637ace85c847b903dff57b69037bb39f6327f80c2f6c32
ARG SOFTTAINER_IMAGE=registry.redhat.io/kube-descheduler-operator/kube-descheduler-rhel9-operator@sha256:f7675a6efa5941b3c0e2411d223b391b34d892cbada9c6d40c55b6603337b1ba
ARG REPLACED_OPERATOR_IMG=registry-proxy.engineering.redhat.com/rh-osbs/kube-descheduler-operator-rhel-9:latest
ARG REPLACED_OPERAND_IMG=registry-proxy.engineering.redhat.com/rh-osbs/descheduler-rhel-9:latest
ARG REPLACED_SOFTTAINER_IMG=registry-proxy.engineering.redhat.com/rh-osbs/kube-descheduler-operator-rhel-9:latest

RUN hack/replace-image.sh manifests ${REPLACED_OPERATOR_IMG} ${OPERATOR_IMAGE}
RUN hack/replace-image.sh manifests ${REPLACED_OPERAND_IMG} ${OPERAND_IMAGE}
RUN hack/replace-image.sh manifests ${REPLACED_SOFTTAINER_IMG} ${SOFTTAINER_IMAGE}

FROM registry.redhat.io/rhel9-4-els/rhel-minimal:9.4

COPY --from=builder /go/src/github.com/openshift/cluster-kube-descheduler-operator/manifests /manifests
COPY --from=builder /go/src/github.com/openshift/cluster-kube-descheduler-operator/metadata /metadata
COPY --from=builder /go/src/github.com/openshift/cluster-kube-descheduler-operator/licenses /licenses

LABEL operators.operatorframework.io.bundle.mediatype.v1="registry+v1"
LABEL operators.operatorframework.io.bundle.manifests.v1=manifests/
LABEL operators.operatorframework.io.bundle.metadata.v1=metadata/
LABEL operators.operatorframework.io.bundle.package.v1="cluster-kube-descheduler-operator"
LABEL operators.operatorframework.io.bundle.channels.v1=stable
LABEL operators.operatorframework.io.bundle.channel.default.v1=stable
LABEL operators.operatorframework.io.metrics.builder=operator-sdk-v1.34.2
LABEL operators.operatorframework.io.metrics.mediatype.v1=metrics+v1
LABEL operators.operatorframework.io.metrics.project_layout=go.kubebuilder.io/v4

LABEL com.redhat.component="kube-descheduler-operator-bundle-container"
LABEL description="Descheduler support for OpenShift based on RHEL 9"
LABEL distribution-scope="public"
LABEL name="cluster-kube-descheduler-operator-metadata-rhel-9"
LABEL release="5.2.0"
LABEL version="5.2.0"
LABEL url="https://github.com/openshift/cluster-kube-descheduler-operator"
LABEL vendor="Red Hat, Inc."
LABEL summary="Descheduler support for OpenShift"
LABEL io.openshift.expose-services=""
LABEL io.k8s.display-name="kube-descheduler-operator based on RHEL 9"
LABEL io.k8s.description="Descheduler support for OpenShift based on RHEL 9"
LABEL io.openshift.tags="openshift,kube-descheduler-operator"
LABEL com.redhat.delivery.operator.bundle=true
LABEL com.redhat.openshift.versions="v4.19"
LABEL com.redhat.delivery.appregistry=true
LABEL maintainer="AOS workloads team, <aos-workloads-staff@redhat.com>"

USER 1001
