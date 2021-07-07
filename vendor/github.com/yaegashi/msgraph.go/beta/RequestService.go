// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ServicePrincipalRequestBuilder is request builder for ServicePrincipal
type ServicePrincipalRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServicePrincipalRequest
func (b *ServicePrincipalRequestBuilder) Request() *ServicePrincipalRequest {
	return &ServicePrincipalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServicePrincipalRequest is request for ServicePrincipal
type ServicePrincipalRequest struct{ BaseRequest }

// Get performs GET request for ServicePrincipal
func (r *ServicePrincipalRequest) Get(ctx context.Context) (resObj *ServicePrincipal, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServicePrincipal
func (r *ServicePrincipalRequest) Update(ctx context.Context, reqObj *ServicePrincipal) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServicePrincipal
func (r *ServicePrincipalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type ServicePrincipalCreatePasswordSingleSignOnCredentialsRequestBuilder struct{ BaseRequestBuilder }

// CreatePasswordSingleSignOnCredentials action undocumented
func (b *ServicePrincipalRequestBuilder) CreatePasswordSingleSignOnCredentials(reqObj *ServicePrincipalCreatePasswordSingleSignOnCredentialsRequestParameter) *ServicePrincipalCreatePasswordSingleSignOnCredentialsRequestBuilder {
	bb := &ServicePrincipalCreatePasswordSingleSignOnCredentialsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createPasswordSingleSignOnCredentials"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ServicePrincipalCreatePasswordSingleSignOnCredentialsRequest struct{ BaseRequest }

//
func (b *ServicePrincipalCreatePasswordSingleSignOnCredentialsRequestBuilder) Request() *ServicePrincipalCreatePasswordSingleSignOnCredentialsRequest {
	return &ServicePrincipalCreatePasswordSingleSignOnCredentialsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ServicePrincipalCreatePasswordSingleSignOnCredentialsRequest) Post(ctx context.Context) (resObj *PasswordSingleSignOnCredentialSet, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ServicePrincipalGetPasswordSingleSignOnCredentialsRequestBuilder struct{ BaseRequestBuilder }

// GetPasswordSingleSignOnCredentials action undocumented
func (b *ServicePrincipalRequestBuilder) GetPasswordSingleSignOnCredentials(reqObj *ServicePrincipalGetPasswordSingleSignOnCredentialsRequestParameter) *ServicePrincipalGetPasswordSingleSignOnCredentialsRequestBuilder {
	bb := &ServicePrincipalGetPasswordSingleSignOnCredentialsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getPasswordSingleSignOnCredentials"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ServicePrincipalGetPasswordSingleSignOnCredentialsRequest struct{ BaseRequest }

//
func (b *ServicePrincipalGetPasswordSingleSignOnCredentialsRequestBuilder) Request() *ServicePrincipalGetPasswordSingleSignOnCredentialsRequest {
	return &ServicePrincipalGetPasswordSingleSignOnCredentialsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ServicePrincipalGetPasswordSingleSignOnCredentialsRequest) Post(ctx context.Context) (resObj *PasswordSingleSignOnCredentialSet, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ServicePrincipalDeletePasswordSingleSignOnCredentialsRequestBuilder struct{ BaseRequestBuilder }

// DeletePasswordSingleSignOnCredentials action undocumented
func (b *ServicePrincipalRequestBuilder) DeletePasswordSingleSignOnCredentials(reqObj *ServicePrincipalDeletePasswordSingleSignOnCredentialsRequestParameter) *ServicePrincipalDeletePasswordSingleSignOnCredentialsRequestBuilder {
	bb := &ServicePrincipalDeletePasswordSingleSignOnCredentialsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/deletePasswordSingleSignOnCredentials"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ServicePrincipalDeletePasswordSingleSignOnCredentialsRequest struct{ BaseRequest }

//
func (b *ServicePrincipalDeletePasswordSingleSignOnCredentialsRequestBuilder) Request() *ServicePrincipalDeletePasswordSingleSignOnCredentialsRequest {
	return &ServicePrincipalDeletePasswordSingleSignOnCredentialsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ServicePrincipalDeletePasswordSingleSignOnCredentialsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ServicePrincipalUpdatePasswordSingleSignOnCredentialsRequestBuilder struct{ BaseRequestBuilder }

// UpdatePasswordSingleSignOnCredentials action undocumented
func (b *ServicePrincipalRequestBuilder) UpdatePasswordSingleSignOnCredentials(reqObj *ServicePrincipalUpdatePasswordSingleSignOnCredentialsRequestParameter) *ServicePrincipalUpdatePasswordSingleSignOnCredentialsRequestBuilder {
	bb := &ServicePrincipalUpdatePasswordSingleSignOnCredentialsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updatePasswordSingleSignOnCredentials"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ServicePrincipalUpdatePasswordSingleSignOnCredentialsRequest struct{ BaseRequest }

//
func (b *ServicePrincipalUpdatePasswordSingleSignOnCredentialsRequestBuilder) Request() *ServicePrincipalUpdatePasswordSingleSignOnCredentialsRequest {
	return &ServicePrincipalUpdatePasswordSingleSignOnCredentialsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ServicePrincipalUpdatePasswordSingleSignOnCredentialsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
