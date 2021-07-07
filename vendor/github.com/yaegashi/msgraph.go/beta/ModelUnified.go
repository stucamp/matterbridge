// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// UnifiedRoleAssignment undocumented
type UnifiedRoleAssignment struct {
	// Entity is the base model of UnifiedRoleAssignment
	Entity
	// PrincipalID undocumented
	PrincipalID *string `json:"principalId,omitempty"`
	// ResourceScope undocumented
	ResourceScope *string `json:"resourceScope,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// Principal undocumented
	Principal *DirectoryObject `json:"principal,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
}

// UnifiedRoleDefinition undocumented
type UnifiedRoleDefinition struct {
	// Entity is the base model of UnifiedRoleDefinition
	Entity
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsBuiltIn undocumented
	IsBuiltIn *bool `json:"isBuiltIn,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// ResourceScopes undocumented
	ResourceScopes []string `json:"resourceScopes,omitempty"`
	// RolePermissions undocumented
	RolePermissions []UnifiedRolePermission `json:"rolePermissions,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// Version undocumented
	Version *string `json:"version,omitempty"`
}

// UnifiedRolePermission undocumented
type UnifiedRolePermission struct {
	// Object is the base model of UnifiedRolePermission
	Object
	// AllowedResourceActions undocumented
	AllowedResourceActions []string `json:"allowedResourceActions,omitempty"`
	// Condition undocumented
	Condition *string `json:"condition,omitempty"`
}
