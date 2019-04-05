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

package ecdscluster

import (
	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/kubedgeoperators/v1alpha1"
	bmgr "github.com/kubedge/kubedge-operator-base/pkg/kubedgemanager"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type ecdsclusterrenderer struct {
	base bmgr.KubedgeBaseRenderer

	spec av1.ECDSClusterSpec
}

// Update the Unstructured read in the file using the content of the Spec.
func (o ecdsclusterrenderer) updateStatefulSet(u *unstructured.Unstructured, k *av1.KubedgeSetSpec) {
	if (k != nil) && (k.Replicas != nil) {

		out := v1.StatefulSet{}
		err1 := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &out)
		if err1 != nil {
			log.Error(err1, "error converting from Unstructured")
		}

		if k.Replicas != nil {
			out.Spec.Replicas = k.Replicas
		}

		unst, err2 := runtime.DefaultUnstructuredConverter.ToUnstructured(&out)
		if err2 != nil {
			log.Error(err2, "error converting to Unstructured")
		}

		u.SetUnstructuredContent(unst)
	}
}

// Adds the ownerrefs to all the documents in a YAML file
func (o ecdsclusterrenderer) RenderFile(name string, namespace string, fileName string) (*av1.SubResourceList, error) {

	// Let render the default resourceList using the ecds-template.yaml
	rendered, err := o.base.RenderFile(name, namespace, fileName)
	if err != nil {
		return rendered, err
	}

	updated := av1.NewSubResourceList(rendered.GetNamespace(), rendered.GetName())

	// Let's adapt the template using the values passed in the spec
	for _, renderedResource := range rendered.Items {
		if renderedResource.GetKind() == "StatefulSet" {
			switch renderedResource.GetName() {
			case ECBusinessLogic.String():
				o.updateStatefulSet(&renderedResource, o.spec.BusinessLogics)
			case ECEnrichment.String():
				o.updateStatefulSet(&renderedResource, o.spec.Enrichments)
			case ECFrontend.String():
				o.updateStatefulSet(&renderedResource, o.spec.FrontEnds)
			case ECLoadbalancer.String():
				o.updateStatefulSet(&renderedResource, o.spec.LoadBalancers)
			case ECPlatform.String():
				o.updateStatefulSet(&renderedResource, o.spec.Platforms)
			}
			updated.Items = append(updated.Items, renderedResource)
		} else {
			updated.Items = append(updated.Items, renderedResource)
		}
	}

	return updated, nil

}

// NewECDSClusterRenderer creates a new OwnerRef engine with a set of metav1.OwnerReferences to be added to assets
func NewECDSClusterRenderer(refs []metav1.OwnerReference, suffix string,
	renderFiles []string, renderValues map[string]interface{}, spec av1.ECDSClusterSpec) bmgr.KubedgeResourceRenderer {
	return ecdsclusterrenderer{
		base: bmgr.KubedgeBaseRenderer{
			Refs:         refs,
			Suffix:       suffix,
			RenderFiles:  renderFiles,
			RenderValues: renderValues,
		},
		spec: spec,
	}
}
