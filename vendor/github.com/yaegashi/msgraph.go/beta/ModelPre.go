// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PreAuthorizedApplication undocumented
type PreAuthorizedApplication struct {
	// Object is the base model of PreAuthorizedApplication
	Object
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// PermissionIDs undocumented
	PermissionIDs []string `json:"permissionIds,omitempty"`
}