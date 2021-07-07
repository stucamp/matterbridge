// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AuditEventRequestBuilder is request builder for AuditEvent
type AuditEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns AuditEventRequest
func (b *AuditEventRequestBuilder) Request() *AuditEventRequest {
	return &AuditEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AuditEventRequest is request for AuditEvent
type AuditEventRequest struct{ BaseRequest }

// Get performs GET request for AuditEvent
func (r *AuditEventRequest) Get(ctx context.Context) (resObj *AuditEvent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AuditEvent
func (r *AuditEventRequest) Update(ctx context.Context, reqObj *AuditEvent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AuditEvent
func (r *AuditEventRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AuditLogRootRequestBuilder is request builder for AuditLogRoot
type AuditLogRootRequestBuilder struct{ BaseRequestBuilder }

// Request returns AuditLogRootRequest
func (b *AuditLogRootRequestBuilder) Request() *AuditLogRootRequest {
	return &AuditLogRootRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AuditLogRootRequest is request for AuditLogRoot
type AuditLogRootRequest struct{ BaseRequest }

// Get performs GET request for AuditLogRoot
func (r *AuditLogRootRequest) Get(ctx context.Context) (resObj *AuditLogRoot, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AuditLogRoot
func (r *AuditLogRootRequest) Update(ctx context.Context, reqObj *AuditLogRoot) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AuditLogRoot
func (r *AuditLogRootRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
