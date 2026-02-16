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
	"fmt"
	"testing"
)

// TestStandard validates the "standard" param set rendering.
func TestStandard(t *testing.T) {
	resources := renderChart(t, "standard")
	controller := mustFind(t, resources, "Deployment", "ebs-csi-controller")
	cPS := podSpec(t, controller)
	ebsPlugin := findContainer(t, cPS, "ebs-plugin")

	nodeDSObj := mustFind(t, resources, "DaemonSet", "ebs-csi-node")
	nPS := podSpec(t, nodeDSObj)
	nodePlugin := findContainer(t, nPS, "ebs-plugin")

	t.Run("controllerMetrics", func(t *testing.T) {
		svc := mustFind(t, resources, "Service", "ebs-csi-controller")
		assertServiceHasPort(t, svc, "metrics", 3301)
	})

	t.Run("nodeMetrics", func(t *testing.T) {
		svc := mustFind(t, resources, "Service", "ebs-csi-node")
		assertServiceHasPort(t, svc, "metrics", 3302)
	})

	t.Run("batching", func(t *testing.T) {
		if !hasArg(ebsPlugin, "--batching=true") {
			t.Error("controller should have --batching=true")
		}
	})

	t.Run("controllerLoggingFormat", func(t *testing.T) {
		if !hasArg(ebsPlugin, "--logging-format=json") {
			t.Error("controller should have --logging-format=json")
		}
	})

	t.Run("nodeLoggingFormat", func(t *testing.T) {
		if !hasArg(nodePlugin, "--logging-format=json") {
			t.Error("node should have --logging-format=json")
		}
	})

	t.Run("controllerUserAgentExtra", func(t *testing.T) {
		if !hasArg(ebsPlugin, "--user-agent-extra=e2e-test") {
			t.Error("controller should have --user-agent-extra=e2e-test")
		}
	})

	t.Run("reservedVolumeAttachments", func(t *testing.T) {
		if !hasArg(nodePlugin, "--reserved-volume-attachments=2") {
			t.Error("node should have --reserved-volume-attachments=2")
		}
	})

	t.Run("hostNetwork", func(t *testing.T) {
		hn, ok := nPS["hostNetwork"].(bool)
		if !ok || !hn {
			t.Error("node pod should use hostNetwork")
		}
	})

	t.Run("nodeTerminationGracePeriod", func(t *testing.T) {
		tgp, ok := nestedFloat(nPS, "terminationGracePeriodSeconds")
		if !ok || tgp != 60 {
			t.Errorf("node terminationGracePeriodSeconds: got %v, want 60", tgp)
		}
	})

	t.Run("nodeTolerateAllTaints", func(t *testing.T) {
		tolerations := nestedSlice(t, nPS, "tolerations")
		var found bool
		for _, tol := range tolerations {
			tm := tol.(obj)
			if tm["operator"] == "Exists" && (tm["key"] == nil || tm["key"] == "") {
				found = true
			}
		}
		if !found {
			t.Error("node should have tolerate-all toleration")
		}
	})

	t.Run("nodeKubeletPath", func(t *testing.T) {
		mounts, ok := nodePlugin["volumeMounts"].([]interface{})
		if !ok {
			t.Fatal("no volumeMounts on node ebs-plugin")
		}
		var found bool
		for _, m := range mounts {
			mm := m.(obj)
			if mm["mountPath"] == "/var/lib/kubelet" {
				found = true
			}
		}
		if !found {
			t.Error("node should mount kubelet path at /var/lib/kubelet")
		}
	})

	t.Run("controllerPodDisruptionBudget", func(t *testing.T) {
		mustFind(t, resources, "PodDisruptionBudget", "ebs-csi-controller")
	})

	t.Run("snapshotterForceEnable", func(t *testing.T) {
		if !hasContainer(cPS, "csi-snapshotter") {
			t.Error("controller should have csi-snapshotter when forceEnable=true")
		}
	})

	t.Run("nodeDisableMutation", func(t *testing.T) {
		cr := mustFind(t, resources, "ClusterRole", "ebs-csi-node-role")
		rules, ok := cr["rules"].([]interface{})
		if !ok {
			t.Fatal("no rules in ClusterRole")
		}
		for _, rule := range rules {
			rm := rule.(obj)
			res, _ := rm["resources"].([]interface{})
			for _, r := range res {
				if r == "nodes" {
					verbs, _ := rm["verbs"].([]interface{})
					for _, v := range verbs {
						if v == "patch" || v == "update" {
							t.Errorf("node role should not have %s on nodes when disableMutation=true", v)
						}
					}
				}
			}
		}
	})

	t.Run("storageClasses", func(t *testing.T) {
		sc := mustFind(t, resources, "StorageClass", "test-sc")
		params := nested(t, sc, "parameters")
		if params["type"] != "gp3" {
			t.Errorf("StorageClass type: got %v, want gp3", params["type"])
		}
	})

	t.Run("volumeSnapshotClasses", func(t *testing.T) {
		vsc := mustFind(t, resources, "VolumeSnapshotClass", "test-vsc")
		dp, ok := nestedString(vsc, "deletionPolicy")
		if !ok || dp != "Delete" {
			t.Errorf("VolumeSnapshotClass deletionPolicy: got %v, want Delete", dp)
		}
	})

	t.Run("defaultStorageClass", func(t *testing.T) {
		sc := mustFind(t, resources, "StorageClass", "ebs-csi-default-sc")
		meta := nested(t, sc, "metadata")
		ann := meta["annotations"].(obj)
		if ann["storageclass.kubernetes.io/is-default-class"] != "true" {
			t.Error("default StorageClass should have is-default-class=true annotation")
		}
	})

	t.Run("nodeAllocatableUpdatePeriodSeconds", func(t *testing.T) {
		csiDriver := mustFind(t, resources, "CSIDriver", "ebs.csi.aws.com")
		val, ok := nestedFloat(csiDriver, "spec", "nodeAllocatableUpdatePeriodSeconds")
		if !ok || val != 30 {
			t.Errorf("nodeAllocatableUpdatePeriodSeconds: got %v, want 30", val)
		}
	})

	// Log level tests
	logLevelTests := []struct {
		name      string
		container string
		podType   string // "controller" or "node"
	}{
		{"controllerLogLevel", "ebs-plugin", "controller"},
		{"provisionerLogLevel", "csi-provisioner", "controller"},
		{"attacherLogLevel", "csi-attacher", "controller"},
		{"snapshotterLogLevel", "csi-snapshotter", "controller"},
		{"resizerLogLevel", "csi-resizer", "controller"},
		{"nodeDriverRegistrarLogLevel", "node-driver-registrar", "node"},
		{"nodeLogLevel", "ebs-plugin", "node"},
	}
	for _, tc := range logLevelTests {
		t.Run(tc.name, func(t *testing.T) {
			var ps obj
			if tc.podType == "controller" {
				ps = cPS
			} else {
				ps = nPS
			}
			c := findContainer(t, ps, tc.container)
			if !hasArgAny(c, "-v=5", "--v=5") {
				t.Errorf("%s should have -v=5", tc.container)
			}
		})
	}

	// Leader election tests
	leaderTests := []struct {
		name      string
		container string
		enabled   bool
	}{
		{"provisionerLeaderElection", "csi-provisioner", true},
		{"attacherLeaderElection", "csi-attacher", true},
		{"resizerLeaderElection", "csi-resizer", true},
	}
	for _, tc := range leaderTests {
		t.Run(tc.name, func(t *testing.T) {
			c := findContainer(t, cPS, tc.container)
			expected := fmt.Sprintf("--leader-election=%v", tc.enabled)
			if !hasArg(c, expected) {
				t.Errorf("%s should have %s", tc.container, expected)
			}
		})
	}

	// Prometheus annotations
	for _, svcName := range []string{"ebs-csi-controller", "ebs-csi-node"} {
		t.Run("prometheusAnnotations/"+svcName, func(t *testing.T) {
			svc := mustFind(t, resources, "Service", svcName)
			meta := nested(t, svc, "metadata")
			ann, ok := meta["annotations"].(obj)
			if !ok {
				t.Fatal("no annotations on service")
			}
			if ann["prometheus.io/scrape"] != "true" {
				t.Error("service should have prometheus.io/scrape=true")
			}
		})
	}
}

// assertServiceHasPort checks that a Service has a port with the given name and number.
func assertServiceHasPort(t *testing.T, svc obj, name string, port int) {
	t.Helper()
	ports := nestedSlice(t, svc, "spec", "ports")
	for _, p := range ports {
		pm := p.(obj)
		pNum, _ := pm["port"].(float64)
		if pm["name"] == name && int(pNum) == port {
			return
		}
	}
	t.Errorf("service should have port %s/%d", name, port)
}
