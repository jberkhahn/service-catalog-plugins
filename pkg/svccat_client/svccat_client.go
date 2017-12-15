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

package svccat_client

import (
	"errors"
	"fmt"

	"k8s.io/kubectl/pkg/framework"

	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
)

type SvcCatClient struct {
	ScClient     clientset.Interface
	PluginClient *framework.PluginClient
}

// NewClient uses the KUBECONFIG environment variable to create a new client
// based on an existing configuration
func NewClient() (*SvcCatClient, error) {
	pluginClient, err := framework.NewClient()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error obtaining a plugin client from existing configuration: %v", err))
	}
	c, err := clientset.NewForConfig(pluginClient.Config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error obtaining a service-catalog client from existing configuration: %v", err))
	}

	svcCatClient := SvcCatClient{c, pluginClient}
	return &svcCatClient, nil
}
