/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"

	logicalcluster "github.com/kcp-dev/logicalcluster/v2"
	v1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// TokenReviewsGetter has a method to return a TokenReviewInterface.
// A group's client should implement this interface.
type TokenReviewsGetter interface {
	TokenReviews() TokenReviewInterface
}

// TokenReviewInterface has methods to work with TokenReview resources.
type TokenReviewInterface interface {
	Create(ctx context.Context, tokenReview *v1.TokenReview, opts metav1.CreateOptions) (*v1.TokenReview, error)
	TokenReviewExpansion
}

// tokenReviews implements TokenReviewInterface
type tokenReviews struct {
	client  rest.Interface
	cluster logicalcluster.Name
}

// newTokenReviews returns a TokenReviews
func newTokenReviews(c *AuthenticationV1Client) *tokenReviews {
	return &tokenReviews{
		client:  c.RESTClient(),
		cluster: c.cluster,
	}
}

// Create takes the representation of a tokenReview and creates it.  Returns the server's representation of the tokenReview, and an error, if there is any.
func (c *tokenReviews) Create(ctx context.Context, tokenReview *v1.TokenReview, opts metav1.CreateOptions) (result *v1.TokenReview, err error) {
	result = &v1.TokenReview{}
	err = c.client.Post().
		Cluster(c.cluster).
		Resource("tokenreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tokenReview).
		Do(ctx).
		Into(result)
	return
}