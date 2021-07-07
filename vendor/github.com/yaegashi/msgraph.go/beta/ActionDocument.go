// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// Comments returns request builder for DocumentComment collection
func (b *DocumentRequestBuilder) Comments() *DocumentCommentsCollectionRequestBuilder {
	bb := &DocumentCommentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/comments"
	return bb
}

// DocumentCommentsCollectionRequestBuilder is request builder for DocumentComment collection
type DocumentCommentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DocumentComment collection
func (b *DocumentCommentsCollectionRequestBuilder) Request() *DocumentCommentsCollectionRequest {
	return &DocumentCommentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DocumentComment item
func (b *DocumentCommentsCollectionRequestBuilder) ID(id string) *DocumentCommentRequestBuilder {
	bb := &DocumentCommentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DocumentCommentsCollectionRequest is request for DocumentComment collection
type DocumentCommentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DocumentComment collection
func (r *DocumentCommentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DocumentComment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DocumentComment
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DocumentComment
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for DocumentComment collection, max N pages
func (r *DocumentCommentsCollectionRequest) GetN(ctx context.Context, n int) ([]DocumentComment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DocumentComment collection
func (r *DocumentCommentsCollectionRequest) Get(ctx context.Context) ([]DocumentComment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DocumentComment collection
func (r *DocumentCommentsCollectionRequest) Add(ctx context.Context, reqObj *DocumentComment) (resObj *DocumentComment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Replies returns request builder for DocumentCommentReply collection
func (b *DocumentCommentRequestBuilder) Replies() *DocumentCommentRepliesCollectionRequestBuilder {
	bb := &DocumentCommentRepliesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/replies"
	return bb
}

// DocumentCommentRepliesCollectionRequestBuilder is request builder for DocumentCommentReply collection
type DocumentCommentRepliesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DocumentCommentReply collection
func (b *DocumentCommentRepliesCollectionRequestBuilder) Request() *DocumentCommentRepliesCollectionRequest {
	return &DocumentCommentRepliesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DocumentCommentReply item
func (b *DocumentCommentRepliesCollectionRequestBuilder) ID(id string) *DocumentCommentReplyRequestBuilder {
	bb := &DocumentCommentReplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DocumentCommentRepliesCollectionRequest is request for DocumentCommentReply collection
type DocumentCommentRepliesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DocumentCommentReply collection
func (r *DocumentCommentRepliesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DocumentCommentReply, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DocumentCommentReply
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DocumentCommentReply
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for DocumentCommentReply collection, max N pages
func (r *DocumentCommentRepliesCollectionRequest) GetN(ctx context.Context, n int) ([]DocumentCommentReply, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DocumentCommentReply collection
func (r *DocumentCommentRepliesCollectionRequest) Get(ctx context.Context) ([]DocumentCommentReply, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DocumentCommentReply collection
func (r *DocumentCommentRepliesCollectionRequest) Add(ctx context.Context, reqObj *DocumentCommentReply) (resObj *DocumentCommentReply, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
