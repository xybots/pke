// Copyright © 2019 Banzai Cloud
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

package controlplane

// storageClassVsphereTemplate is a generated function returning the template as a string.
func storageClassVsphereTemplate() string {
	var tmpl = "kind: StorageClass\n" +
		"apiVersion: storage.k8s.io/v1\n" +
		"metadata:\n" +
		"  name: vsphere\n" +
		"  annotations:\n" +
		"    storageclass.kubernetes.io/is-default-class: \"true\"\n" +
		"provisioner: kubernetes.io/vsphere-volume\n" +
		"parameters:\n" +
		"  diskformat: thin\n" +
		""
	return tmpl
}
