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

package plugin_client

import (
	"github.com/jberkhahn/service-catalog-plugins/pkg/utils"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListInstances(namespace string) (*v1beta1.ServiceInstanceList, error) {
	scClient, _ := utils.NewClient()
	instances, err := scClient.ServicecatalogV1beta1().ServiceInstances(namespace).List(v1.ListOptions{})
	return instances, err
}

func GetInstance(namespace, instanceName string) (*v1beta1.ServiceInstance, error) {
	scClient, _ := utils.NewClient()
	instance, err := scClient.ServicecatalogV1beta1().ServiceInstances(namespace).Get(instanceName, v1.GetOptions{})
	return instance, err
}
