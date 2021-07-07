// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsFactRequestBuilder struct{ BaseRequestBuilder }

// Fact action undocumented
func (b *WorkbookFunctionsRequestBuilder) Fact(reqObj *WorkbookFunctionsFactRequestParameter) *WorkbookFunctionsFactRequestBuilder {
	bb := &WorkbookFunctionsFactRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/fact"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsFactRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsFactRequestBuilder) Request() *WorkbookFunctionsFactRequest {
	return &WorkbookFunctionsFactRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsFactRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsFactDoubleRequestBuilder struct{ BaseRequestBuilder }

// FactDouble action undocumented
func (b *WorkbookFunctionsRequestBuilder) FactDouble(reqObj *WorkbookFunctionsFactDoubleRequestParameter) *WorkbookFunctionsFactDoubleRequestBuilder {
	bb := &WorkbookFunctionsFactDoubleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/factDouble"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsFactDoubleRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsFactDoubleRequestBuilder) Request() *WorkbookFunctionsFactDoubleRequest {
	return &WorkbookFunctionsFactDoubleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsFactDoubleRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
