// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// OutlookCategory undocumented
type OutlookCategory struct {
	// Entity is the base model of OutlookCategory
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Color undocumented
	Color *CategoryColor `json:"color,omitempty"`
}

// OutlookGeoCoordinates undocumented
type OutlookGeoCoordinates struct {
	// Object is the base model of OutlookGeoCoordinates
	Object
	// Altitude undocumented
	Altitude *float64 `json:"altitude,omitempty"`
	// Latitude undocumented
	Latitude *float64 `json:"latitude,omitempty"`
	// Longitude undocumented
	Longitude *float64 `json:"longitude,omitempty"`
	// Accuracy undocumented
	Accuracy *float64 `json:"accuracy,omitempty"`
	// AltitudeAccuracy undocumented
	AltitudeAccuracy *float64 `json:"altitudeAccuracy,omitempty"`
}

// OutlookItem undocumented
type OutlookItem struct {
	// Entity is the base model of OutlookItem
	Entity
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// ChangeKey undocumented
	ChangeKey *string `json:"changeKey,omitempty"`
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
}

// OutlookTask undocumented
type OutlookTask struct {
	// OutlookItem is the base model of OutlookTask
	OutlookItem
	// AssignedTo undocumented
	AssignedTo *string `json:"assignedTo,omitempty"`
	// Body undocumented
	Body *ItemBody `json:"body,omitempty"`
	// CompletedDateTime undocumented
	CompletedDateTime *DateTimeTimeZone `json:"completedDateTime,omitempty"`
	// DueDateTime undocumented
	DueDateTime *DateTimeTimeZone `json:"dueDateTime,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// Importance undocumented
	Importance *Importance `json:"importance,omitempty"`
	// IsReminderOn undocumented
	IsReminderOn *bool `json:"isReminderOn,omitempty"`
	// Owner undocumented
	Owner *string `json:"owner,omitempty"`
	// ParentFolderID undocumented
	ParentFolderID *string `json:"parentFolderId,omitempty"`
	// Recurrence undocumented
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// ReminderDateTime undocumented
	ReminderDateTime *DateTimeTimeZone `json:"reminderDateTime,omitempty"`
	// Sensitivity undocumented
	Sensitivity *Sensitivity `json:"sensitivity,omitempty"`
	// StartDateTime undocumented
	StartDateTime *DateTimeTimeZone `json:"startDateTime,omitempty"`
	// Status undocumented
	Status *TaskStatus `json:"status,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// Attachments undocumented
	Attachments []Attachment `json:"attachments,omitempty"`
}

// OutlookTaskFolder undocumented
type OutlookTaskFolder struct {
	// Entity is the base model of OutlookTaskFolder
	Entity
	// ChangeKey undocumented
	ChangeKey *string `json:"changeKey,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// IsDefaultFolder undocumented
	IsDefaultFolder *bool `json:"isDefaultFolder,omitempty"`
	// ParentGroupKey undocumented
	ParentGroupKey *UUID `json:"parentGroupKey,omitempty"`
	// Tasks undocumented
	Tasks []OutlookTask `json:"tasks,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
}

// OutlookTaskGroup undocumented
type OutlookTaskGroup struct {
	// Entity is the base model of OutlookTaskGroup
	Entity
	// ChangeKey undocumented
	ChangeKey *string `json:"changeKey,omitempty"`
	// IsDefaultGroup undocumented
	IsDefaultGroup *bool `json:"isDefaultGroup,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// GroupKey undocumented
	GroupKey *UUID `json:"groupKey,omitempty"`
	// TaskFolders undocumented
	TaskFolders []OutlookTaskFolder `json:"taskFolders,omitempty"`
}

// OutlookUser undocumented
type OutlookUser struct {
	// Entity is the base model of OutlookUser
	Entity
	// MasterCategories undocumented
	MasterCategories []OutlookCategory `json:"masterCategories,omitempty"`
	// TaskGroups undocumented
	TaskGroups []OutlookTaskGroup `json:"taskGroups,omitempty"`
	// TaskFolders undocumented
	TaskFolders []OutlookTaskFolder `json:"taskFolders,omitempty"`
	// Tasks undocumented
	Tasks []OutlookTask `json:"tasks,omitempty"`
}
