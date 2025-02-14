/*
Copyright 2021 The OpenEBS Authors

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/openebs/device-localpv/pkg/apis/openebs.io/device/v1alpha1"
	scheme "github.com/openebs/device-localpv/pkg/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeviceVolumesGetter has a method to return a DeviceVolumeInterface.
// A group's client should implement this interface.
type DeviceVolumesGetter interface {
	DeviceVolumes(namespace string) DeviceVolumeInterface
}

// DeviceVolumeInterface has methods to work with DeviceVolume resources.
type DeviceVolumeInterface interface {
	Create(*v1alpha1.DeviceVolume) (*v1alpha1.DeviceVolume, error)
	Update(*v1alpha1.DeviceVolume) (*v1alpha1.DeviceVolume, error)
	UpdateStatus(*v1alpha1.DeviceVolume) (*v1alpha1.DeviceVolume, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DeviceVolume, error)
	List(opts v1.ListOptions) (*v1alpha1.DeviceVolumeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeviceVolume, err error)
	DeviceVolumeExpansion
}

// deviceVolumes implements DeviceVolumeInterface
type deviceVolumes struct {
	client rest.Interface
	ns     string
}

// newDeviceVolumes returns a DeviceVolumes
func newDeviceVolumes(c *LocalV1alpha1Client, namespace string) *deviceVolumes {
	return &deviceVolumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deviceVolume, and returns the corresponding deviceVolume object, and an error if there is any.
func (c *deviceVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.DeviceVolume, err error) {
	result = &v1alpha1.DeviceVolume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicevolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeviceVolumes that match those selectors.
func (c *deviceVolumes) List(opts v1.ListOptions) (result *v1alpha1.DeviceVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DeviceVolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicevolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deviceVolumes.
func (c *deviceVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("devicevolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a deviceVolume and creates it.  Returns the server's representation of the deviceVolume, and an error, if there is any.
func (c *deviceVolumes) Create(deviceVolume *v1alpha1.DeviceVolume) (result *v1alpha1.DeviceVolume, err error) {
	result = &v1alpha1.DeviceVolume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("devicevolumes").
		Body(deviceVolume).
		Do().
		Into(result)
	return
}

// Update takes the representation of a deviceVolume and updates it. Returns the server's representation of the deviceVolume, and an error, if there is any.
func (c *deviceVolumes) Update(deviceVolume *v1alpha1.DeviceVolume) (result *v1alpha1.DeviceVolume, err error) {
	result = &v1alpha1.DeviceVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devicevolumes").
		Name(deviceVolume.Name).
		Body(deviceVolume).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *deviceVolumes) UpdateStatus(deviceVolume *v1alpha1.DeviceVolume) (result *v1alpha1.DeviceVolume, err error) {
	result = &v1alpha1.DeviceVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devicevolumes").
		Name(deviceVolume.Name).
		SubResource("status").
		Body(deviceVolume).
		Do().
		Into(result)
	return
}

// Delete takes name of the deviceVolume and deletes it. Returns an error if one occurs.
func (c *deviceVolumes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicevolumes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deviceVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicevolumes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched deviceVolume.
func (c *deviceVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeviceVolume, err error) {
	result = &v1alpha1.DeviceVolume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("devicevolumes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
