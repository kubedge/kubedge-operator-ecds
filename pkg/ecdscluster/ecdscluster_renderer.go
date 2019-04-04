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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ecdsclusterrenderer struct {
	bmgr.KubedgeBaseRenderer

	spec av1.ECDSClusterSpec
}

// Adds the ownerrefs to all the documents in a YAML file
func (o ecdsclusterrenderer) RenderFile(name string, namespace string, fileName string) (*av1.SubResourceList, error) {
	return nil, nil
}

// NewECDSClusterRenderer creates a new OwnerRef engine with a set of metav1.OwnerReferences to be added to assets
func NewECDSClusterRenderer(refs []metav1.OwnerReference, suffix string,
	renderFiles []string, renderValues map[string]interface{}, spec av1.ECDSClusterSpec) bmgr.KubedgeResourceRenderer {
	return ecdsclusterrenderer{
		KubedgeBaseRenderer: bmgr.KubedgeBaseRenderer{
			refs:         refs,
			suffix:       suffix,
			renderFiles:  renderFiles,
			renderValues: renderValues,
		},
		spec: spec,
	}
}
