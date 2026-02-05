# Copyright 2025 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

###
### This Makefile is documented in docs/makefile.md
###

## Variables/Functions

VERSION?=v1.55.0

PKG=github.com/kubernetes-sigs/aws-ebs-csi-driver
GIT_COMMIT?=$(shell git rev-parse HEAD)
BUILD_DATE?=$(shell date -u -Iseconds)
LDFLAGS?="-X ${PKG}/pkg/driver.driverVersion=${VERSION} -X ${PKG}/pkg/cloud.driverVersion=${VERSION} -X ${PKG}/pkg/driver.gitCommit=${GIT_COMMIT} -X ${PKG}/pkg/driver.buildDate=${BUILD_DATE} -s -w"

OS?=$(shell go env GOHOSTOS)
ARCH?=$(shell go env GOHOSTARCH)
ifeq ($(OS),windows)
	BINARY=aws-ebs-csi-driver.exe
	OSVERSION?=ltsc2022
else
	BINARY=aws-ebs-csi-driver
	OSVERSION?=al2023
endif
FIPS?=false
ifeq ($(FIPS),true)
	FIPS_DOCKER_ARGS=--build-arg=GOEXPERIMENT=boringcrypto
endif

GO_SOURCES=go.mod go.sum $(shell find pkg cmd -type f -name "*.go")

ALL_OS?=linux windows
ALL_ARCH_linux?=amd64 arm64
ALL_OSVERSION_linux?=al2023
ALL_OS_ARCH_OSVERSION_linux=$(foreach arch, $(ALL_ARCH_linux), $(foreach osversion, ${ALL_OSVERSION_linux}, linux-$(arch)-${osversion}))

ALL_ARCH_windows?=amd64
ALL_OSVERSION_windows?=ltsc2019 ltsc2022
ALL_OS_ARCH_OSVERSION_windows=$(foreach arch, $(ALL_ARCH_windows), $(foreach osversion, ${ALL_OSVERSION_windows}, windows-$(arch)-${osversion}))
ALL_OS_ARCH_OSVERSION=$(foreach os, $(ALL_OS), ${ALL_OS_ARCH_OSVERSION_${os}})

CLUSTER_NAME?=ebs-csi-e2e.k8s.local
CLUSTER_TYPE?=kops

GINKGO_WINDOWS_SKIP?="\[Disruptive\]|\[Serial\]|\[Flaky\]|\[LinuxOnly\]|\[Feature:VolumeSnapshotDataSource\]|\(xfs\)|\(ext4\)|\(block volmode\)|should resize volume when PVC is edited and the pod is re-created on the same node after controller resize is finished"
GINKGO_BOTTLEROCKET_SKIP?="\[Disruptive\]|\[Serial\]|\[Flaky\]|should not mount / map unused volumes in a pod \[LinuxOnly\]"

# split words on hyphen, access by 1-index
word-hyphen = $(word $2,$(subst -, ,$1))

.EXPORT_ALL_VARIABLES:

## Default target
# When no target is supplied, make runs the first target that does not begin with a .
# Alias that to building the binary
.PHONY: default
default: bin/$(BINARY)

## Top level targets

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: test
test:
	go test -v -race ./cmd/... ./pkg/... ./tests/sanity/...

.PHONY: test/coverage
test/coverage:
	go test -coverprofile=cover.out ./cmd/... ./pkg/...
	grep -v "mock" cover.out > filtered_cover.out
	go tool cover -html=filtered_cover.out -o coverage.html
	rm cover.out filtered_cover.out

.PHONY: tools
tools: bin/aws bin/ct bin/eksctl bin/ginkgo bin/golangci-lint bin/gomplate bin/helm bin/kops bin/kubetest2 bin/mockgen bin/shfmt

.PHONY: update
update: update/gofix update/gofmt update/golangci-fix update/kustomize update/mockgen update/gomod update/shfmt update/generate-license-header
	@echo "All updates succeeded!"

.PHONY: verify
verify: verify/govet verify/golangci-lint verify/update
	@echo "All verifications passed!"

.PHONY: cluster/create
cluster/create: bin/kops bin/eksctl bin/aws bin/gomplate
	./hack/e2e/create-cluster.sh

.PHONY: cluster/kubeconfig
cluster/kubeconfig:
	@./hack/e2e/kubeconfig.sh

.PHONY: cluster/image
cluster/image: bin/aws
	./hack/e2e/build-image.sh

.PHONY: cluster/delete
cluster/delete: bin/kops bin/eksctl
	./hack/e2e/delete-cluster.sh

.PHONY: cluster/install
cluster/install: bin/helm bin/aws
	./hack/e2e/install.sh

.PHONY: cluster/helm
cluster/helm: bin/helm bin/aws
	HELM_USE_DEFAULT_IMAGE="true" \
	./hack/e2e/install.sh

.PHONY: cluster/uninstall
cluster/uninstall: bin/helm bin/aws
	./hack/e2e/uninstall.sh

## E2E targets
# Targets to run e2e tests

## Parameter-specific e2e tests
# Run tests for specific Helm parameter sets
# Usage: make e2e/parameters PARAM_SET=standard

.PHONY: e2e/parameters
e2e/parameters: bin/helm bin/ginkgo
ifndef PARAM_SET
	$(error PARAM_SET is required. Options: standard, volume-modification, node-config, storage-classes, debug, metadata-labeler, legacy-compat, selinux, fips)
endif
ifeq ($(PARAM_SET),standard)
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:(extraCreateMetadata|k8sTagClusterId|extraVolumeTags|controllerMetrics|nodeMetrics|batching|defaultFsType|controllerLoggingFormat|nodeLoggingFormat|controllerLogLevel|nodeLogLevel|provisionerLogLevel|attacherLogLevel|snapshotterLogLevel|resizerLogLevel|nodeDriverRegistrarLogLevel|storageClasses|volumeSnapshotClasses|defaultStorageClass|snapshotterForceEnable|controllerUserAgentExtra|controllerEnablePrometheusAnnotations|nodeEnablePrometheusAnnotations|nodeKubeletPath|nodeTolerateAllTaints|controllerPodDisruptionBudget|provisionerLeaderElection|attacherLeaderElection|resizerLeaderElection)\]" \
	GINKGO_PARALLEL=5 \
	HELM_EXTRA_FLAGS="--set=controller.extraCreateMetadata=true,controller.k8sTagClusterId=e2e-param-test,controller.extraVolumeTags.TestKey=TestValue,controller.enableMetrics=true,node.enableMetrics=true,controller.batching=true,controller.defaultFsType=xfs,controller.loggingFormat=json,node.loggingFormat=json,controller.logLevel=4,node.logLevel=4,sidecars.provisioner.logLevel=4,sidecars.attacher.logLevel=4,sidecars.snapshotter.logLevel=4,sidecars.resizer.logLevel=4,sidecars.nodeDriverRegistrar.logLevel=4,defaultStorageClass.enabled=true,storageClasses[0].name=test-sc,storageClasses[0].parameters.type=gp3,volumeSnapshotClasses[0].name=test-vsc,volumeSnapshotClasses[0].deletionPolicy=Delete,sidecars.snapshotter.forceEnable=true,controller.userAgentExtra=e2e-test,controller.enablePrometheusAnnotations=true,node.enablePrometheusAnnotations=true,node.kubeletPath=/var/lib/kubelet,node.tolerateAllTaints=true,controller.podDisruptionBudget.enabled=true,sidecars.provisioner.leaderElection.enabled=true,sidecars.attacher.leaderElection.enabled=true,sidecars.resizer.leaderElection.enabled=true" \
	./hack/e2e/run.sh
else ifeq ($(PARAM_SET),volume-modification)
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:(volumeModification|volumemodifierLogLevel|volumemodifierLeaderElection)\]" \
	GINKGO_PARALLEL=5 \
	HELM_EXTRA_FLAGS="--set=controller.volumeModificationFeature.enabled=true,sidecars.provisioner.additionalArgs[0]='--feature-gates=VolumeAttributesClass=true',sidecars.resizer.additionalArgs[0]='--feature-gates=VolumeAttributesClass=true',sidecars.volumemodifier.logLevel=4,sidecars.volumemodifier.leaderElection.enabled=false" \
	./hack/e2e/run.sh
else ifeq ($(PARAM_SET),node-config)
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:(reservedVolumeAttachments|hostNetwork|nodeDisableMutation|nodeTerminationGracePeriod)\]" \
	GINKGO_PARALLEL=5 \
	HELM_EXTRA_FLAGS="--set=node.reservedVolumeAttachments=2,node.hostNetwork=true,node.serviceAccount.disableMutation=true,node.terminationGracePeriodSeconds=60" \
	./hack/e2e/run.sh
# volume-attach-limit must be separate from node-config because volumeAttachLimit and reservedVolumeAttachments are mutually exclusive
else ifeq ($(PARAM_SET),volume-attach-limit)
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:volumeAttachLimit\]" \
	GINKGO_PARALLEL=5 \
	HELM_EXTRA_FLAGS="--set=node.volumeAttachLimit=25" \
	./hack/e2e/run.sh
else ifeq ($(PARAM_SET),storage-classes)
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:(storageClasses|volumeSnapshotClasses|defaultStorageClass)\]" \
	GINKGO_PARALLEL=5 \
	HELM_EXTRA_FLAGS="--set=defaultStorageClass.enabled=true,storageClasses[0].name=test-sc,storageClasses[0].parameters.type=gp3,volumeSnapshotClasses[0].name=test-vsc,volumeSnapshotClasses[0].deletionPolicy=Delete" \
	./hack/e2e/run.sh
else ifeq ($(PARAM_SET),debug) #debugLogs=true overrides the individual logLevel settings. We need to keep debug separate or not set debugLogs when testing individual log levels. 
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:(debugLogs|sdkDebugLog)\]" \
	GINKGO_PARALLEL=5 \
	HELM_EXTRA_FLAGS="--set=debugLogs=true,controller.sdkDebugLog=true" \
	./hack/e2e/run.sh
else ifeq ($(PARAM_SET),metadata-labeler)
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:(metadataLabeler|metadataLabelerLogLevel)\]" \
	GINKGO_PARALLEL=1 \
	EBS_INSTALL_SNAPSHOT=false \
	HELM_EXTRA_FLAGS="--set=sidecars.metadataLabeler.enabled=true,node.metadataSources='metadata-labeler',sidecars.metadataLabeler.logLevel=4" \
	./hack/e2e/run.sh
else ifeq ($(PARAM_SET),legacy-compat)
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:(useOldCSIDriver|legacyXFS)\]" \
	GINKGO_PARALLEL=5 \
	HELM_EXTRA_FLAGS="--set=useOldCSIDriver=true,node.legacyXFS=true" \
	./hack/e2e/run.sh
else ifeq ($(PARAM_SET),selinux)
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:selinux\]" \
	GINKGO_PARALLEL=5 \
	HELM_EXTRA_FLAGS="--set=node.selinux=true" \
	./hack/e2e/run.sh
else ifeq ($(PARAM_SET),fips)
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:fips\]" \
	GINKGO_PARALLEL=5 \
	FIPS_TEST=true \
	HELM_EXTRA_FLAGS="--set=fips=true" \
	./hack/e2e/run.sh
else ifeq ($(PARAM_SET),infra)
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:(controllerReplicaCount|controllerPriorityClassName|controllerResources|controllerPodAnnotations|controllerPodLabels|controllerDeploymentAnnotations|controllerRevisionHistoryLimit|nodePriorityClassName|nodeResources|nodePodAnnotations|nodeDaemonSetAnnotations|nodeRevisionHistoryLimit|provisionerResources|attacherResources|snapshotterResources|resizerResources|nodeDriverRegistrarResources|livenessProbeResources|customLabels|controllerEnv|nodeEnv|controllerTopologySpreadConstraints|controllerSecurityContext|nodeSecurityContext|controllerContainerSecurityContext|controllerVolumes|controllerVolumeMounts|nodeVolumes|nodeVolumeMounts|controllerDnsConfig|nodeDnsConfig|controllerInitContainers|nodeInitContainers|imagePullPolicy)\]" \
	GINKGO_PARALLEL=5 \
	HELM_EXTRA_FLAGS="--set=controller.replicaCount=3,controller.priorityClassName=system-cluster-critical,controller.resources.requests.cpu=100m,controller.resources.limits.memory=256Mi,controller.podAnnotations.test-annotation=test-value,controller.podLabels.test-label=test-value,controller.deploymentAnnotations.deploy-annotation=deploy-value,controller.revisionHistoryLimit=5,node.priorityClassName=system-node-critical,node.resources.requests.cpu=50m,node.resources.limits.memory=128Mi,node.podAnnotations.node-annotation=node-value,node.daemonSetAnnotations.ds-annotation=ds-value,node.revisionHistoryLimit=3,sidecars.provisioner.resources.requests.cpu=20m,sidecars.attacher.resources.requests.cpu=15m,sidecars.snapshotter.resources.requests.cpu=15m,sidecars.resizer.resources.requests.cpu=15m,sidecars.nodeDriverRegistrar.resources.requests.cpu=10m,sidecars.livenessProbe.resources.requests.cpu=5m,customLabels.custom-label=custom-value,controller.env[0].name=TEST_ENV,controller.env[0].value=test-value,node.env[0].name=NODE_ENV,node.env[0].value=node-value,controller.topologySpreadConstraints[0].maxSkew=1,controller.topologySpreadConstraints[0].topologyKey=topology.kubernetes.io/zone,controller.topologySpreadConstraints[0].whenUnsatisfiable=ScheduleAnyway,controller.securityContext.runAsNonRoot=true,controller.containerSecurityContext.readOnlyRootFilesystem=true,controller.volumes[0].name=extra-volume,controller.volumes[0].configMap.name=kube-root-ca.crt,controller.volumeMounts[0].name=extra-volume,controller.volumeMounts[0].mountPath=/extra,node.volumes[0].name=node-extra-volume,node.volumes[0].configMap.name=kube-root-ca.crt,node.volumeMounts[0].name=node-extra-volume,node.volumeMounts[0].mountPath=/node-extra,controller.dnsConfig.nameservers[0]=8.8.8.8,node.dnsConfig.nameservers[0]=8.8.4.4,controller.initContainers[0].name=init-container,controller.initContainers[0].image=busybox,controller.initContainers[0].command[0]=echo,controller.initContainers[0].command[1]=init,node.initContainers[0].name=node-init-container,node.initContainers[0].image=busybox,node.initContainers[0].command[0]=echo,node.initContainers[0].command[1]=node-init,image.pullPolicy=Always" \
	./hack/e2e/run.sh
# update-strategy must be separate because Recreate conflicts with default rollingUpdate settings
else ifeq ($(PARAM_SET),update-strategy)
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\] \[param:(controllerUpdateStrategy|nodeUpdateStrategy)\]" \
	GINKGO_PARALLEL=5 \
	HELM_EXTRA_FLAGS="--set=controller.updateStrategy.type=Recreate,controller.updateStrategy.rollingUpdate=null,node.updateStrategy.type=OnDelete" \
	./hack/e2e/run.sh
else
	$(error Unknown PARAM_SET: $(PARAM_SET). Options: standard, volume-modification, node-config, volume-attach-limit, storage-classes, debug, infra, update-strategy, metadata-labeler, legacy-compat, selinux, fips)
endif

.PHONY: e2e/parameters-all
e2e/parameters-all: bin/helm bin/ginkgo
	@echo "Running all parameter sets sequentially..."
	$(MAKE) e2e/parameters PARAM_SET=standard
	$(MAKE) e2e/parameters PARAM_SET=volume-modification
	$(MAKE) e2e/parameters PARAM_SET=node-config
	$(MAKE) e2e/parameters PARAM_SET=volume-attach-limit
	$(MAKE) e2e/parameters PARAM_SET=storage-classes
	$(MAKE) e2e/parameters PARAM_SET=debug
	$(MAKE) e2e/parameters PARAM_SET=infra
	$(MAKE) e2e/parameters PARAM_SET=update-strategy
	@echo "All parameter sets completed successfully!"

.PHONY: e2e/single-az
e2e/single-az: bin/helm bin/ginkgo
	AWS_AVAILABILITY_ZONES=us-west-2a \
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[single-az\]" \
	GINKGO_SKIP="\[param:" \
	GINKGO_PARALLEL=5 \
	HELM_EXTRA_FLAGS="--set=controller.volumeModificationFeature.enabled=true,sidecars.provisioner.additionalArgs[0]='--feature-gates=VolumeAttributesClass=true',sidecars.resizer.additionalArgs[0]='--feature-gates=VolumeAttributesClass=true',node.enableMetrics=true" \
	./hack/e2e/run.sh

.PHONY: e2e/multi-az
e2e/multi-az: bin/helm bin/ginkgo
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[multi-az\]" \
	GINKGO_PARALLEL=5 \
	./hack/e2e/run.sh

.PHONY: e2e/disruptive
e2e/disruptive: bin/helm bin/ginkgo
	TEST_PATH=./tests/e2e/... \
	GINKGO_FOCUS="\[ebs-csi-e2e\] \[Disruptive\]" \
	GINKGO_SKIP="\[Flaky\]" \
	GINKGO_PARALLEL=1 \
	EBS_INSTALL_SNAPSHOT=false \
	HELM_EXTRA_FLAGS="--set=sidecars.metadataLabeler.enabled=true,node.metadataSources='metadata-labeler'" \
	./hack/e2e/run.sh

.PHONY: e2e/external
e2e/external: bin/helm bin/kubetest2
	COLLECT_METRICS="true" \
	./hack/e2e/run.sh

.PHONY: e2e/external-eks-bottlerocket
e2e/external-eks-bottlerocket: bin/helm bin/kubetest2
	GINKGO_SKIP=$(GINKGO_BOTTLEROCKET_SKIP) \
	./hack/e2e/run.sh

.PHONY: e2e/external-fips
e2e/external-fips: bin/helm bin/kubetest2
	HELM_EXTRA_FLAGS="--set=fips=true" \
	./hack/e2e/run.sh

.PHONY: e2e/external-windows
e2e/external-windows: bin/helm bin/kubetest2
	WINDOWS=true \
	GINKGO_SKIP=$(GINKGO_WINDOWS_SKIP) \
	GINKGO_PARALLEL=15 \
	EBS_INSTALL_SNAPSHOT="false" \
	./hack/e2e/run.sh

.PHONY: e2e/external-windows-fips
e2e/external-windows-fips: bin/helm bin/kubetest2
	WINDOWS=true \
	GINKGO_SKIP=$(GINKGO_WINDOWS_SKIP) \
	GINKGO_PARALLEL=15 \
	EBS_INSTALL_SNAPSHOT="false" \
	HELM_EXTRA_FLAGS="--set=fips=true" \
	./hack/e2e/run.sh

.PHONY: e2e/external-windows-hostprocess
e2e/external-windows-hostprocess: bin/helm bin/kubetest2
	WINDOWS_HOSTPROCESS=true \
	WINDOWS=true \
	GINKGO_SKIP=$(GINKGO_WINDOWS_SKIP) \
	GINKGO_PARALLEL=15 \
	EBS_INSTALL_SNAPSHOT="false" \
	./hack/e2e/run.sh

.PHONY: e2e/external-kustomize
e2e/external-kustomize: bin/kubetest2
	DEPLOY_METHOD="kustomize" \
	./hack/e2e/run.sh

.PHONY: e2e/helm-ct
e2e/helm-ct: bin/helm bin/ct
	HELM_CT_TEST="true" \
	./hack/e2e/run.sh

## Release scripts
# Targets run as part of performing a release

.PHONY: update-truth-sidecars
update-truth-sidecars: hack/release-scripts/get-latest-sidecar-images
	./hack/release-scripts/get-latest-sidecar-images

.PHONY: generate-sidecar-tags
generate-sidecar-tags: update-truth-sidecars charts/aws-ebs-csi-driver/values.yaml deploy/kubernetes/overlays/stable/gcr/kustomization.yaml hack/release-scripts/generate-sidecar-tags
	./hack/release-scripts/generate-sidecar-tags

.PHONY: update-sidecar-dependencies
update-sidecar-dependencies: update-truth-sidecars generate-sidecar-tags update/kustomize

.PHONY: update-image-dependencies
update-image-dependencies: update-sidecar-dependencies
	./hack/release-scripts/update-e2e-images

.PHONY: security
security: bin/govulncheck
	./hack/tools/check-security.sh

## CI aliases
# Targets intended to be executed mostly or only by CI jobs

.PHONY: sub-push
sub-push: all-image-registry push-manifest

.PHONY: sub-push-fips
sub-push-fips:
	$(MAKE) FIPS=true TAG=$(TAG)-fips sub-push

.PHONY: all-push
all-push: sub-push sub-push-fips

test-e2e-%:
	./hack/prow-e2e.sh test-e2e-$*

test-helm-chart:
	./hack/prow-e2e.sh test-helm-chart

.PHONY: test-images 
test-images: bin/aws 
	./hack/e2e/test-images.sh 

## Builds

bin:
	@mkdir -p $@

bin/$(BINARY): $(GO_SOURCES) | bin
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -mod=readonly -ldflags ${LDFLAGS} -o $@ ./cmd/

.PHONY: all-image-registry
all-image-registry: $(addprefix sub-image-,$(ALL_OS_ARCH_OSVERSION))

sub-image-%:
	$(MAKE) OS=$(call word-hyphen,$*,1) ARCH=$(call word-hyphen,$*,2) OSVERSION=$(call word-hyphen,$*,3) image

.PHONY: image
image:
	BUILDX_NO_DEFAULT_ATTESTATIONS=1 docker buildx build \
		--platform=$(OS)/$(ARCH) \
		--progress=plain \
		--target=$(OS)-$(OSVERSION) \
		--output=type=registry \
		-t=$(IMAGE):$(TAG)-$(OS)-$(ARCH)-$(OSVERSION) \
		--build-arg=GOPROXY=$(GOPROXY) \
		--build-arg=VERSION=$(VERSION) \
		$(FIPS_DOCKER_ARGS) \
		$(DOCKER_EXTRA_ARGS) \
		.

.PHONY: create-manifest
create-manifest: all-image-registry
# sed expression:
# LHS: match 0 or more not space characters
# RHS: replace with $(IMAGE):$(TAG)-& where & is what was matched on LHS
	docker manifest create --amend $(IMAGE):$(TAG) $(shell echo $(ALL_OS_ARCH_OSVERSION) | sed -e "s~[^ ]*~$(IMAGE):$(TAG)\-&~g")

.PHONY: push-manifest
push-manifest: create-manifest
	docker manifest push --purge $(IMAGE):$(TAG)

## Tools
# Tools necessary to perform other targets

bin/%: hack/tools/install.sh hack/tools/python-runner.sh
	@TOOLS_PATH="$(shell pwd)/bin" ./hack/tools/install.sh $*

## Updaters
# Automatic generators/formatters for code

.PHONY: update/gofix
update/gofix:
	go fix ./...

.PHONY: update/gofmt
update/gofmt:
	gofmt -s -w .

.PHONY: update/golangci-fix
update/golangci-fix: bin/golangci-lint
ifndef SKIP_GOLANGCI_FIX
	./bin/golangci-lint run --fix ./... || true
endif

.PHONY: update/kustomize
update/kustomize: bin/helm
	./hack/update-kustomize.sh

.PHONY: update/mockgen
update/mockgen: bin/mockgen
	./hack/update-mockgen.sh

.PHONY: update/gomod
update/gomod:
	go mod tidy
	go mod tidy -C tests/e2e/

.PHONY: update/shfmt
update/shfmt: bin/shfmt
	./bin/shfmt -w -i 2 -d ./hack/

.PHONY: update/generate-license-header
update/generate-license-header:
	./hack/generate-license-header.sh

.PHONY: generate-volume-limits-table
generate-volume-limits-table:
	go run ./hack/generate-volume-limits-table > pkg/cloud/limits/volume_limits_table.go
	gofmt -s -w pkg/cloud/limits/volume_limits_table.go
	go run ./hack/detect-potentially-invalid-limits

## Verifiers
# Linters and similar

.PHONY: verify/golangci-lint
verify/golangci-lint: bin/golangci-lint
	./bin/golangci-lint run --timeout=10m --verbose

.PHONY: verify/govet
verify/govet:
	go vet $$(go list ./...)

.PHONY: verify/update
verify/update: bin/helm bin/mockgen
	./hack/verify-update.sh
