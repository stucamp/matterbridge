// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RegistryKeyState undocumented
type RegistryKeyState struct {
	// Object is the base model of RegistryKeyState
	Object
	// Hive undocumented
	Hive *RegistryHive `json:"hive,omitempty"`
	// Key undocumented
	Key *string `json:"key,omitempty"`
	// OldKey undocumented
	OldKey *string `json:"oldKey,omitempty"`
	// OldValueData undocumented
	OldValueData *string `json:"oldValueData,omitempty"`
	// OldValueName undocumented
	OldValueName *string `json:"oldValueName,omitempty"`
	// Operation undocumented
	Operation *RegistryOperation `json:"operation,omitempty"`
	// ProcessID undocumented
	ProcessID *int `json:"processId,omitempty"`
	// ValueData undocumented
	ValueData *string `json:"valueData,omitempty"`
	// ValueName undocumented
	ValueName *string `json:"valueName,omitempty"`
	// ValueType undocumented
	ValueType *RegistryValueType `json:"valueType,omitempty"`
}
