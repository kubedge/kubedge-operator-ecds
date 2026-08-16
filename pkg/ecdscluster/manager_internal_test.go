// Copyright 2019 The Kubedge Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Internal tests (package ecdscluster) for the unexported factory/render helpers —
// the shared contract consumers rely on: owner-ref stamping and render-value mapping.
package ecdscluster

import (
	"testing"

	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/kubedgeoperators/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestExecutionContextKindString(t *testing.T) {
	cases := map[ExecutionContextKind]string{
		ECBusinessLogic: "businesslogic",
		ECEnrichment:    "enrichment",
		ECFrontend:      "frontend",
		ECLoadbalancer:  "loadbalancer",
		ECPlatform:      "platform",
	}
	for k, want := range cases {
		if got := k.String(); got != want {
			t.Errorf("ExecutionContextKind(%q).String() = %q, want %q", string(k), got, want)
		}
	}
}

func TestInitRenderValuesCarriesStage(t *testing.T) {
	v := initRenderValues(av1.PhaseRollback)
	oslc, ok := v["oslc"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected v[\"oslc\"] to be map[string]interface{}, got %T", v["oslc"])
	}
	if got := oslc["stage"]; got != av1.PhaseRollback.String() {
		t.Errorf("oslc.stage = %v, want %v", got, av1.PhaseRollback.String())
	}
}

func TestInitRenderFilesEmpty(t *testing.T) {
	if f := initRenderFiles(av1.PhaseRollback); len(f) != 0 {
		t.Errorf("initRenderFiles = %v, want empty", f)
	}
}

// NewECDSClusterManager must stamp a controller owner-ref derived from the CR and
// carry the CR's name/namespace/source down to the base manager — this is what makes
// the operator's rendered child resources garbage-collect with their ECDSCluster.
func TestNewECDSClusterManagerStampsOwnerRef(t *testing.T) {
	r := &av1.ECDSCluster{}
	r.SetName("my-ecds")
	r.SetNamespace("kubedge")
	r.SetUID("uid-123")
	// NewControllerRef reads the owner's GVK for the ref's APIVersion/Kind.
	r.TypeMeta = metav1.TypeMeta{
		Kind:       "ECDSCluster",
		APIVersion: "kubedgeoperators.kubedge.cloud/v1alpha1",
	}

	f := managerFactory{kubeClient: nil}
	mgr, ok := f.NewECDSClusterManager(r).(*ecdsclustermgr)
	if !ok {
		t.Fatalf("NewECDSClusterManager returned %T, want *ecdsclustermgr", f.NewECDSClusterManager(r))
	}

	if mgr.PhaseName != "my-ecds" {
		t.Errorf("PhaseName = %q, want %q", mgr.PhaseName, "my-ecds")
	}
	if mgr.PhaseNamespace != "kubedge" {
		t.Errorf("PhaseNamespace = %q, want %q", mgr.PhaseNamespace, "kubedge")
	}

	if len(mgr.OwnerRefs) != 1 {
		t.Fatalf("OwnerRefs len = %d, want 1", len(mgr.OwnerRefs))
	}
	ref := mgr.OwnerRefs[0]
	if ref.Name != "my-ecds" || ref.Kind != "ECDSCluster" || string(ref.UID) != "uid-123" {
		t.Errorf("owner ref = {Name:%q Kind:%q UID:%q}, want {my-ecds ECDSCluster uid-123}", ref.Name, ref.Kind, ref.UID)
	}
	if ref.Controller == nil || !*ref.Controller {
		t.Errorf("owner ref Controller = %v, want true", ref.Controller)
	}
}

// The unexported factory constructors for the sibling CRs are intentional no-ops in
// this consumer (ecds only reconciles ECDSCluster); pin that so a future edit is deliberate.
func TestSiblingManagersAreNil(t *testing.T) {
	f := managerFactory{kubeClient: nil}
	if m := f.NewArpscanManager(&av1.Arpscan{}); m != nil {
		t.Errorf("NewArpscanManager = %v, want nil", m)
	}
	if m := f.NewMMESimManager(&av1.MMESim{}); m != nil {
		t.Errorf("NewMMESimManager = %v, want nil", m)
	}
	if m := f.NewEMBBSliceManager(&av1.EMBBSlice{}); m != nil {
		t.Errorf("NewEMBBSliceManager = %v, want nil", m)
	}
}
