/*


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
// Code generated by client-gen-v0.30. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	scheme "github.com/VictoriaMetrics/operator/api/client/versioned/scheme"
	v1beta1 "github.com/VictoriaMetrics/operator/api/operator/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VLogsGetter has a method to return a VLogsInterface.
// A group's client should implement this interface.
type VLogsGetter interface {
	VLogs(namespace string) VLogsInterface
}

// VLogsInterface has methods to work with VLogs resources.
type VLogsInterface interface {
	Create(ctx context.Context, vLogs *v1beta1.VLogs, opts v1.CreateOptions) (*v1beta1.VLogs, error)
	Update(ctx context.Context, vLogs *v1beta1.VLogs, opts v1.UpdateOptions) (*v1beta1.VLogs, error)
	UpdateStatus(ctx context.Context, vLogs *v1beta1.VLogs, opts v1.UpdateOptions) (*v1beta1.VLogs, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VLogs, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VLogsList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VLogs, err error)
	VLogsExpansion
}

// vLogs implements VLogsInterface
type vLogs struct {
	client rest.Interface
	ns     string
}

// newVLogs returns a VLogs
func newVLogs(c *OperatorV1beta1Client, namespace string) *vLogs {
	return &vLogs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vLogs, and returns the corresponding vLogs object, and an error if there is any.
func (c *vLogs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VLogs, err error) {
	result = &v1beta1.VLogs{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vlogs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VLogs that match those selectors.
func (c *vLogs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VLogsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VLogsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vlogs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vLogs.
func (c *vLogs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vlogs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vLogs and creates it.  Returns the server's representation of the vLogs, and an error, if there is any.
func (c *vLogs) Create(ctx context.Context, vLogs *v1beta1.VLogs, opts v1.CreateOptions) (result *v1beta1.VLogs, err error) {
	result = &v1beta1.VLogs{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vlogs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vLogs).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vLogs and updates it. Returns the server's representation of the vLogs, and an error, if there is any.
func (c *vLogs) Update(ctx context.Context, vLogs *v1beta1.VLogs, opts v1.UpdateOptions) (result *v1beta1.VLogs, err error) {
	result = &v1beta1.VLogs{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vlogs").
		Name(vLogs.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vLogs).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vLogs) UpdateStatus(ctx context.Context, vLogs *v1beta1.VLogs, opts v1.UpdateOptions) (result *v1beta1.VLogs, err error) {
	result = &v1beta1.VLogs{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vlogs").
		Name(vLogs.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vLogs).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vLogs and deletes it. Returns an error if one occurs.
func (c *vLogs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vlogs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vLogs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vlogs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vLogs.
func (c *vLogs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VLogs, err error) {
	result = &v1beta1.VLogs{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vlogs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}