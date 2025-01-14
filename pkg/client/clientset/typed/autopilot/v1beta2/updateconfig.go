/*
Copyright k0s authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	context "context"

	autopilotv1beta2 "github.com/k0sproject/k0s/pkg/apis/autopilot/v1beta2"
	scheme "github.com/k0sproject/k0s/pkg/client/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// UpdateConfigsGetter has a method to return a UpdateConfigInterface.
// A group's client should implement this interface.
type UpdateConfigsGetter interface {
	UpdateConfigs() UpdateConfigInterface
}

// UpdateConfigInterface has methods to work with UpdateConfig resources.
type UpdateConfigInterface interface {
	Create(ctx context.Context, updateConfig *autopilotv1beta2.UpdateConfig, opts v1.CreateOptions) (*autopilotv1beta2.UpdateConfig, error)
	Update(ctx context.Context, updateConfig *autopilotv1beta2.UpdateConfig, opts v1.UpdateOptions) (*autopilotv1beta2.UpdateConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*autopilotv1beta2.UpdateConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*autopilotv1beta2.UpdateConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	UpdateConfigExpansion
}

// updateConfigs implements UpdateConfigInterface
type updateConfigs struct {
	*gentype.ClientWithList[*autopilotv1beta2.UpdateConfig, *autopilotv1beta2.UpdateConfigList]
}

// newUpdateConfigs returns a UpdateConfigs
func newUpdateConfigs(c *AutopilotV1beta2Client) *updateConfigs {
	return &updateConfigs{
		gentype.NewClientWithList[*autopilotv1beta2.UpdateConfig, *autopilotv1beta2.UpdateConfigList](
			"updateconfigs",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *autopilotv1beta2.UpdateConfig { return &autopilotv1beta2.UpdateConfig{} },
			func() *autopilotv1beta2.UpdateConfigList { return &autopilotv1beta2.UpdateConfigList{} },
		),
	}
}
