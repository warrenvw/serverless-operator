// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	scheme "github.com/openshift-knative/serverless-operator/openshift-knative-operator/pkg/client/clientset/versioned/scheme"
	v1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AuthenticationsGetter has a method to return a AuthenticationInterface.
// A group's client should implement this interface.
type AuthenticationsGetter interface {
	Authentications() AuthenticationInterface
}

// AuthenticationInterface has methods to work with Authentication resources.
type AuthenticationInterface interface {
	Create(ctx context.Context, authentication *v1.Authentication, opts metav1.CreateOptions) (*v1.Authentication, error)
	Update(ctx context.Context, authentication *v1.Authentication, opts metav1.UpdateOptions) (*v1.Authentication, error)
	UpdateStatus(ctx context.Context, authentication *v1.Authentication, opts metav1.UpdateOptions) (*v1.Authentication, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Authentication, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AuthenticationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Authentication, err error)
	AuthenticationExpansion
}

// authentications implements AuthenticationInterface
type authentications struct {
	client rest.Interface
}

// newAuthentications returns a Authentications
func newAuthentications(c *ConfigV1Client) *authentications {
	return &authentications{
		client: c.RESTClient(),
	}
}

// Get takes name of the authentication, and returns the corresponding authentication object, and an error if there is any.
func (c *authentications) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Authentication, err error) {
	result = &v1.Authentication{}
	err = c.client.Get().
		Resource("authentications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Authentications that match those selectors.
func (c *authentications) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AuthenticationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AuthenticationList{}
	err = c.client.Get().
		Resource("authentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested authentications.
func (c *authentications) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("authentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a authentication and creates it.  Returns the server's representation of the authentication, and an error, if there is any.
func (c *authentications) Create(ctx context.Context, authentication *v1.Authentication, opts metav1.CreateOptions) (result *v1.Authentication, err error) {
	result = &v1.Authentication{}
	err = c.client.Post().
		Resource("authentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authentication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a authentication and updates it. Returns the server's representation of the authentication, and an error, if there is any.
func (c *authentications) Update(ctx context.Context, authentication *v1.Authentication, opts metav1.UpdateOptions) (result *v1.Authentication, err error) {
	result = &v1.Authentication{}
	err = c.client.Put().
		Resource("authentications").
		Name(authentication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authentication).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *authentications) UpdateStatus(ctx context.Context, authentication *v1.Authentication, opts metav1.UpdateOptions) (result *v1.Authentication, err error) {
	result = &v1.Authentication{}
	err = c.client.Put().
		Resource("authentications").
		Name(authentication.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authentication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the authentication and deletes it. Returns an error if one occurs.
func (c *authentications) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("authentications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *authentications) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("authentications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched authentication.
func (c *authentications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Authentication, err error) {
	result = &v1.Authentication{}
	err = c.client.Patch(pt).
		Resource("authentications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
