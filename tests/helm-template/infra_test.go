/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package helmtemplate

import (
	"testing"
)

func TestInfra(t *testing.T) {
	resources := renderChart(t, "infra")
	deploy := mustFind(t, resources, "Deployment", "ebs-csi-controller")
	dSpec := nested(t, deploy, "spec")
	cPS := podSpec(t, deploy)
	ebsPlugin := findContainer(t, cPS, "ebs-plugin")

	ds := mustFind(t, resources, "DaemonSet", "ebs-csi-node")
	dsSpec := nested(t, ds, "spec")
	nPS := podSpec(t, ds)
	nodePlugin := findContainer(t, nPS, "ebs-plugin")

	t.Run("controllerReplicaCount", func(t *testing.T) {
		replicas, ok := nestedFloat(dSpec, "replicas")
		if !ok || replicas != 3 {
			t.Errorf("replicas: got %v, want 3", replicas)
		}
	})

	t.Run("controllerPriorityClassName", func(t *testing.T) {
		if cPS["priorityClassName"] != "system-node-critical" {
			t.Error("controller priorityClassName should be system-node-critical")
		}
	})

	t.Run("controllerResources", func(t *testing.T) {
		res := nested(t, ebsPlugin, "resources")
		cpu, _ := nestedString(res, "requests", "cpu")
		mem, _ := nestedString(res, "limits", "memory")
		if cpu != "100m" {
			t.Errorf("controller cpu request: got %s, want 100m", cpu)
		}
		if mem != "256Mi" {
			t.Errorf("controller memory limit: got %s, want 256Mi", mem)
		}
	})

	t.Run("controllerPodAnnotations", func(t *testing.T) {
		tmpl := nested(t, deploy, "spec", "template", "metadata")
		ann := tmpl["annotations"].(obj)
		if ann["test-annotation"] != "test-value" {
			t.Error("controller pod should have test-annotation")
		}
	})

	t.Run("controllerPodLabels", func(t *testing.T) {
		tmpl := nested(t, deploy, "spec", "template", "metadata")
		labels := tmpl["labels"].(obj)
		if labels["test-label"] != "test-value" {
			t.Error("controller pod should have test-label")
		}
	})

	t.Run("controllerDeploymentAnnotations", func(t *testing.T) {
		meta := nested(t, deploy, "metadata")
		ann := meta["annotations"].(obj)
		if ann["deploy-annotation"] != "deploy-value" {
			t.Error("deployment should have deploy-annotation")
		}
	})

	t.Run("controllerRevisionHistoryLimit", func(t *testing.T) {
		rhl, ok := nestedFloat(dSpec, "revisionHistoryLimit")
		if !ok || rhl != 5 {
			t.Errorf("revisionHistoryLimit: got %v, want 5", rhl)
		}
	})

	t.Run("controllerTopologySpreadConstraints", func(t *testing.T) {
		tscs := nestedSlice(t, cPS, "topologySpreadConstraints")
		if len(tscs) == 0 {
			t.Fatal("no topologySpreadConstraints")
		}
		tsc := tscs[0].(obj)
		if tsc["topologyKey"] != "topology.kubernetes.io/zone" {
			t.Error("topologyKey should be topology.kubernetes.io/zone")
		}
	})

	t.Run("controllerSecurityContext", func(t *testing.T) {
		sc := nested(t, cPS, "securityContext")
		if sc["runAsNonRoot"] != true {
			t.Error("controller securityContext.runAsNonRoot should be true")
		}
	})

	t.Run("controllerContainerSecurityContext", func(t *testing.T) {
		sc := nested(t, ebsPlugin, "securityContext")
		if sc["readOnlyRootFilesystem"] != true {
			t.Error("ebs-plugin securityContext.readOnlyRootFilesystem should be true")
		}
	})

	t.Run("controllerUpdateStrategy", func(t *testing.T) {
		strategy := nested(t, dSpec, "strategy")
		if strategy["type"] != "Recreate" {
			t.Errorf("strategy type: got %v, want Recreate", strategy["type"])
		}
	})

	t.Run("controllerEnv", func(t *testing.T) {
		envs := nestedSlice(t, ebsPlugin, "env")
		var found bool
		for _, e := range envs {
			em := e.(obj)
			if em["name"] == "TEST_ENV" && em["value"] == "test-value" {
				found = true
			}
		}
		if !found {
			t.Error("controller should have TEST_ENV=test-value")
		}
	})

	t.Run("controllerVolumes", func(t *testing.T) {
		vols := nestedSlice(t, cPS, "volumes")
		var found bool
		for _, v := range vols {
			if v.(obj)["name"] == "extra-volume" {
				found = true
			}
		}
		if !found {
			t.Error("controller should have extra-volume")
		}
	})

	t.Run("controllerVolumeMounts", func(t *testing.T) {
		mounts := nestedSlice(t, ebsPlugin, "volumeMounts")
		var found bool
		for _, m := range mounts {
			mm := m.(obj)
			if mm["name"] == "extra-volume" && mm["mountPath"] == "/extra" {
				found = true
			}
		}
		if !found {
			t.Error("controller should have extra-volume mount at /extra")
		}
	})

	t.Run("controllerDnsConfig", func(t *testing.T) {
		dns := nested(t, cPS, "dnsConfig")
		ns := dns["nameservers"].([]interface{})
		var found bool
		for _, n := range ns {
			if n == "8.8.8.8" {
				found = true
			}
		}
		if !found {
			t.Error("controller dnsConfig should have 8.8.8.8")
		}
	})

	t.Run("controllerInitContainers", func(t *testing.T) {
		inits := nestedSlice(t, cPS, "initContainers")
		var found bool
		for _, c := range inits {
			if c.(obj)["name"] == "init-container" {
				found = true
			}
		}
		if !found {
			t.Error("controller should have init-container")
		}
	})

	t.Run("controllerTolerations", func(t *testing.T) {
		tols := nestedSlice(t, cPS, "tolerations")
		var found bool
		for _, tol := range tols {
			tm := tol.(obj)
			if tm["key"] == "test-key" && tm["value"] == "test-value" && tm["effect"] == "NoSchedule" {
				found = true
			}
		}
		if !found {
			t.Error("controller should have test toleration")
		}
	})

	t.Run("controllerAdditionalArgs", func(t *testing.T) {
		if !hasArg(ebsPlugin, "--warn-on-invalid-tag") {
			t.Error("controller should have --warn-on-invalid-tag")
		}
	})

	t.Run("nameOverride", func(t *testing.T) {
		meta := nested(t, deploy, "metadata")
		labels := meta["labels"].(obj)
		if labels["app.kubernetes.io/name"] != "custom-ebs-name" {
			t.Error("app.kubernetes.io/name should be custom-ebs-name")
		}
	})

	t.Run("imagePullPolicy", func(t *testing.T) {
		if ebsPlugin["imagePullPolicy"] != "Always" {
			t.Error("ebs-plugin imagePullPolicy should be Always")
		}
	})

	t.Run("customLabels/controller", func(t *testing.T) {
		meta := nested(t, deploy, "metadata")
		labels := meta["labels"].(obj)
		if labels["custom-label"] != "custom-value" {
			t.Error("controller should have custom-label")
		}
	})

	// --- Node DaemonSet assertions ---

	t.Run("nodePriorityClassName", func(t *testing.T) {
		if nPS["priorityClassName"] != "system-node-critical" {
			t.Error("node priorityClassName should be system-node-critical")
		}
	})

	t.Run("nodeResources", func(t *testing.T) {
		res := nested(t, nodePlugin, "resources")
		cpu, _ := nestedString(res, "requests", "cpu")
		mem, _ := nestedString(res, "limits", "memory")
		if cpu != "50m" {
			t.Errorf("node cpu request: got %s, want 50m", cpu)
		}
		if mem != "128Mi" {
			t.Errorf("node memory limit: got %s, want 128Mi", mem)
		}
	})

	t.Run("nodePodAnnotations", func(t *testing.T) {
		tmpl := nested(t, ds, "spec", "template", "metadata")
		ann := tmpl["annotations"].(obj)
		if ann["node-annotation"] != "node-value" {
			t.Error("node pod should have node-annotation")
		}
	})

	t.Run("nodeDaemonSetAnnotations", func(t *testing.T) {
		meta := nested(t, ds, "metadata")
		ann := meta["annotations"].(obj)
		if ann["ds-annotation"] != "ds-value" {
			t.Error("daemonset should have ds-annotation")
		}
	})

	t.Run("nodeRevisionHistoryLimit", func(t *testing.T) {
		rhl, ok := nestedFloat(dsSpec, "revisionHistoryLimit")
		if !ok || rhl != 3 {
			t.Errorf("node revisionHistoryLimit: got %v, want 3", rhl)
		}
	})

	t.Run("nodeSecurityContext", func(t *testing.T) {
		if nPS["securityContext"] == nil {
			t.Error("node should have securityContext")
		}
	})

	t.Run("nodeUpdateStrategy", func(t *testing.T) {
		strategy := nested(t, dsSpec, "updateStrategy")
		if strategy["type"] != "OnDelete" {
			t.Errorf("node updateStrategy: got %v, want OnDelete", strategy["type"])
		}
	})

	t.Run("nodeEnv", func(t *testing.T) {
		envs := nestedSlice(t, nodePlugin, "env")
		var found bool
		for _, e := range envs {
			em := e.(obj)
			if em["name"] == "NODE_ENV" && em["value"] == "node-value" {
				found = true
			}
		}
		if !found {
			t.Error("node should have NODE_ENV=node-value")
		}
	})

	t.Run("nodeVolumes", func(t *testing.T) {
		vols := nestedSlice(t, nPS, "volumes")
		var found bool
		for _, v := range vols {
			if v.(obj)["name"] == "node-extra-volume" {
				found = true
			}
		}
		if !found {
			t.Error("node should have node-extra-volume")
		}
	})

	t.Run("nodeVolumeMounts", func(t *testing.T) {
		mounts := nestedSlice(t, nodePlugin, "volumeMounts")
		var found bool
		for _, m := range mounts {
			mm := m.(obj)
			if mm["name"] == "node-extra-volume" && mm["mountPath"] == "/node-extra" {
				found = true
			}
		}
		if !found {
			t.Error("node should have node-extra-volume mount at /node-extra")
		}
	})

	t.Run("nodeAdditionalArgs", func(t *testing.T) {
		if !hasArg(nodePlugin, "--logtostderr") {
			t.Error("node should have --logtostderr")
		}
	})

	t.Run("nodeDnsConfig", func(t *testing.T) {
		dns := nested(t, nPS, "dnsConfig")
		ns := dns["nameservers"].([]interface{})
		var found bool
		for _, n := range ns {
			if n == "8.8.4.4" {
				found = true
			}
		}
		if !found {
			t.Error("node dnsConfig should have 8.8.4.4")
		}
	})

	t.Run("nodeInitContainers", func(t *testing.T) {
		inits := nestedSlice(t, nPS, "initContainers")
		var found bool
		for _, c := range inits {
			if c.(obj)["name"] == "node-init-container" {
				found = true
			}
		}
		if !found {
			t.Error("node should have node-init-container")
		}
	})

	t.Run("nodeTolerations", func(t *testing.T) {
		tols := nestedSlice(t, nPS, "tolerations")
		var found bool
		for _, tol := range tols {
			tm := tol.(obj)
			if tm["key"] == "node-key" && tm["value"] == "node-value" {
				found = true
			}
		}
		if !found {
			t.Error("node should have node-key toleration")
		}
	})

	t.Run("customLabels/node", func(t *testing.T) {
		meta := nested(t, ds, "metadata")
		labels := meta["labels"].(obj)
		if labels["custom-label"] != "custom-value" {
			t.Error("node should have custom-label")
		}
	})

	// Sidecar resource tests
	sidecarTests := []struct {
		name        string
		resKind     string
		resName     string
		container   string
		expectedCPU string
	}{
		{"provisionerResources", "Deployment", "ebs-csi-controller", "csi-provisioner", "20m"},
		{"attacherResources", "Deployment", "ebs-csi-controller", "csi-attacher", "15m"},
		{"snapshotterResources", "Deployment", "ebs-csi-controller", "csi-snapshotter", "15m"},
		{"resizerResources", "Deployment", "ebs-csi-controller", "csi-resizer", "15m"},
		{"nodeDriverRegistrarResources", "DaemonSet", "ebs-csi-node", "node-driver-registrar", "10m"},
		{"livenessProbeResources", "DaemonSet", "ebs-csi-node", "liveness-probe", "5m"},
	}
	for _, tc := range sidecarTests {
		t.Run(tc.name, func(t *testing.T) {
			r := mustFind(t, resources, tc.resKind, tc.resName)
			ps := podSpec(t, r)
			c := findContainer(t, ps, tc.container)
			res := nested(t, c, "resources")
			cpu, _ := nestedString(res, "requests", "cpu")
			if cpu != tc.expectedCPU {
				t.Errorf("%s cpu request: got %s, want %s", tc.container, cpu, tc.expectedCPU)
			}
		})
	}

	t.Run("provisionerAdditionalArgs", func(t *testing.T) {
		provisioner := findContainer(t, cPS, "csi-provisioner")
		if !hasArg(provisioner, "--retry-interval-start=10s") {
			t.Error("provisioner should have --retry-interval-start=10s")
		}
	})
}

func TestDebug(t *testing.T) {
	resources := renderChart(t, "debug")
	controller := mustFind(t, resources, "Deployment", "ebs-csi-controller")
	cPS := podSpec(t, controller)
	ebsPlugin := findContainer(t, cPS, "ebs-plugin")

	t.Run("debugLogs", func(t *testing.T) {
		if !hasArgAny(ebsPlugin, "-v=7", "--v=7") {
			t.Error("controller ebs-plugin should have -v=7 when debugLogs=true")
		}
	})

	t.Run("sdkDebugLog", func(t *testing.T) {
		if !hasArg(ebsPlugin, "--aws-sdk-debug-log=true") {
			t.Error("controller should have --aws-sdk-debug-log=true")
		}
	})
}

func TestOther(t *testing.T) {
	resources := renderChart(t, "other")
	controller := mustFind(t, resources, "Deployment", "ebs-csi-controller")
	cPS := podSpec(t, controller)

	ds := mustFind(t, resources, "DaemonSet", "ebs-csi-node")
	nPS := podSpec(t, ds)
	nodePlugin := findContainer(t, nPS, "ebs-plugin")

	t.Run("volumeModification", func(t *testing.T) {
		if !hasContainer(cPS, "volumemodifier") {
			t.Error("controller should have volumemodifier sidecar")
		}
	})

	t.Run("volumemodifierLogLevel", func(t *testing.T) {
		c := findContainer(t, cPS, "volumemodifier")
		if !hasArgAny(c, "-v=5", "--v=5") {
			t.Error("volumemodifier should have -v=5")
		}
	})

	t.Run("volumemodifierLeaderElection", func(t *testing.T) {
		c := findContainer(t, cPS, "volumemodifier")
		if !hasArg(c, "--leader-election=false") {
			t.Error("volumemodifier should have --leader-election=false")
		}
	})

	t.Run("volumeAttachLimit", func(t *testing.T) {
		if !hasArg(nodePlugin, "--volume-attach-limit=25") {
			t.Error("node should have --volume-attach-limit=25")
		}
	})

	t.Run("metadataLabeler", func(t *testing.T) {
		if !hasContainer(cPS, "metadata-labeler") {
			t.Error("controller should have metadata-labeler sidecar")
		}
	})

	t.Run("metadataLabelerLogLevel", func(t *testing.T) {
		c := findContainer(t, cPS, "metadata-labeler")
		if !hasArgAny(c, "-v=5", "--v=5") {
			t.Error("metadata-labeler should have -v=5")
		}
	})

	t.Run("additionalDaemonSets", func(t *testing.T) {
		extraDS := mustFind(t, resources, "DaemonSet", "ebs-csi-node-extra")
		extraPS := podSpec(t, extraDS)
		c := findContainer(t, extraPS, "ebs-plugin")
		if !hasArg(c, "--volume-attach-limit=15") {
			t.Error("additional DaemonSet should have --volume-attach-limit=15")
		}
	})
}

func TestNodeComponentOnly(t *testing.T) {
	resources := renderChart(t, "node-component-only")

	t.Run("noController", func(t *testing.T) {
		_, found := find(resources, "Deployment", "ebs-csi-controller")
		if found {
			t.Error("controller Deployment should not exist when nodeComponentOnly=true")
		}
	})

	t.Run("nodeExists", func(t *testing.T) {
		mustFind(t, resources, "DaemonSet", "ebs-csi-node")
	})
}

// Placeholder: selinux and legacy-compat would need their own values files and tests.
// fips test checks the actual running image, so it stays in e2e.
