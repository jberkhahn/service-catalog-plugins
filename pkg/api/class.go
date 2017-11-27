/*
Copyright 2016 The Kubernetes Authors.

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

package api

import (
	"github.com/jberkhahn/service-catalog-plugins/pkg/utils"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListClasses() (*v1beta1.ClusterServiceClassList, error) {
	scClient, _ := utils.NewClient()
	classes, err := scClient.ServicecatalogV1beta1().ClusterServiceClasses().List(v1.ListOptions{})
	return classes, err
}

func GetClass(className string) (*v1beta1.ClusterServiceClass, error) {
	scClient, _ := utils.NewClient()
	class, err := scClient.ServicecatalogV1beta1().ClusterServiceClasses().Get(className, v1.GetOptions{})
	return class, err
}
