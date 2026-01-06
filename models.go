package twentygo

import (
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

////////////////////////////
// Params                 //
////////////////////////////

// RestoreManyParams defines parameters for RestoreMany functions.
type RestoreManyParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *string `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *int `form:"depth,omitempty" json:"depth,string,omitempty"`
}

// RestoreOneParams defines parameters for RestoreOne functions.
type RestoreOneParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *int `form:"depth,omitempty" json:"depth,string,omitempty"`
}

// FindDuplicatesBody defines parameters for FindDuplicates functions.
type FindDuplicatesBody struct {
	Data *[]CalendarEvent      `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// MergeManyBody defines parameters for MergeMany functions.
type MergeManyBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// FindDuplicatesParams defines parameters for FindDuplicates functions.
type FindDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *int `form:"depth,omitempty" json:"depth,string,omitempty"`
}

// MergeManyParams defines parameters for MergeMany functions.
type MergeManyParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *int `form:"depth,omitempty" json:"depth,string,omitempty"`
}

// CreateManyParams defines parameters for CreateMany functions.
type CreateManyParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *int `form:"depth,omitempty" json:"depth,string,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *bool `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// DeleteManyParams defines parameters for DeleteMany functions.
type DeleteManyParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *string `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *bool `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyParams defines parameters for FindMany functions.
type FindManyParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *string `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *string `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *int `form:"limit,omitempty" json:"limit,string,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *int `form:"depth,omitempty" json:"depth,string,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *string `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *string `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// UpdateManyParams defines parameters for UpdateMany functions.
type UpdateManyParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *int `form:"depth,omitempty" json:"depth,string,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *string `form:"filter,omitempty" json:"filter,omitempty"`
}

// CreateOneParams defines parameters for CreateOne functions.
type CreateOneParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *int `form:"depth,omitempty" json:"depth,string,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *bool `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// DeleteOneParams defines parameters for DeleteOne functions.
type DeleteOneParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *bool `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneParams defines parameters for FindOne functions.
type FindOneParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *int `form:"depth,omitempty" json:"depth,string,omitempty"`
}

// UpdateOneParams defines parameters for UpdateOne functions.
type UpdateOneParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *int `form:"depth,omitempty" json:"depth,string,omitempty"`
}

////////////////////////////
// Errors.                //
////////////////////////////

type ErrorResponse struct {
	StatusCode int      `json:"statusCode"`
	Error      string   `json:"error"`
	Messages   []string `json:"messages"`
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	StartCursor string `json:"startCursor"`
	EndCursor   string `json:"endCursor"`
}
type Response struct {
	Data       Data      `json:"data"`
	PageInfo   *PageInfo `json:"pageInfo,omitempty"`
	TotalCount *int      `json:"totalCount,omitempty"`
}
type Data struct {
	Attachment        *Attachment   `json:"attachment,omitempty"`
	Attachments       []*Attachment `json:"attachments,omitempty"`
	UpdateAttachment  *Attachment   `json:"updateAttachment,omitempty"`
	UpdateAttachments []*Attachment `json:"updateAttachments,omitempty"`
	CreateAttachment  *Attachment   `json:"createAttachment,omitempty"`
	CreateAttachments []*Attachment `json:"createAttachments,omitempty"`

	Person        *Person   `json:"person,omitempty"`
	Persons       []*Person `json:"persons,omitempty"`
	UpdatePerson  *Person   `json:"updatePerson,omitempty"`
	UpdatePersons []*Person `json:"updatePersons,omitempty"`
	CreatePerson  *Person   `json:"createPerson,omitempty"`
	CreatePersons []*Person `json:"createPersons,omitempty"`

	Opportunity         *Opportunity   `json:"opportunity,omitempty"`
	Opportunities       []*Opportunity `json:"opportunities,omitempty"`
	UpdateOpportunity   *Opportunity   `json:"updateOpportunity,omitempty"`
	UpdateOpportunities []*Opportunity `json:"updateOpportunities,omitempty"`
	CreateOpportunity   *Opportunity   `json:"createOpportunity,omitempty"`
	CreateOpportunities []*Opportunity `json:"createOpportunities,omitempty"`

	Company         *Company   `json:"company,omitempty"`
	Companies       []*Company `json:"companies,omitempty"`
	UpdateCompany   *Company   `json:"updateCompany,omitempty"`
	UpdateCompanies []*Company `json:"updateCompanies,omitempty"`
	CreateCompany   *Company   `json:"createCompany,omitempty"`
	CreateCompanies []*Company `json:"createCompanies,omitempty"`

	Task        *Task   `json:"task,omitempty"`
	Tasks       []*Task `json:"tasks,omitempty"`
	UpdateTask  *Task   `json:"updateTask,omitempty"`
	UpdateTasks []*Task `json:"updateTasks,omitempty"`
	CreateTask  *Task   `json:"createTask,omitempty"`
	CreateTasks []*Task `json:"createTasks,omitempty"`
}

// APIErrType is a field containing more specific API error types
// that may be checked by the receiver.
type APIErrType string

const (
	// APIErrTypeUnknown is for API errors that are not strongly
	// typed.
	APIErrTypeUnknown APIErrType = "unknown"

	// APIErrTypeInvalidGrant corresponds with Keycloak's
	// OAuthErrorException due to "invalid_grant".
	APIErrTypeInvalidGrant = "oauth: invalid grant"
)

// ParseAPIErrType is a convenience method for returning strongly
// typed API errors.
func ParseAPIErrType(err error) APIErrType {
	if err == nil {
		return APIErrTypeUnknown
	}
	switch {
	case strings.Contains(err.Error(), "invalid_grant"):
		return APIErrTypeInvalidGrant
	default:
		return APIErrTypeUnknown
	}
}

// APIError holds message and statusCode for api errors
type APIError struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Type    APIErrType `json:"type"`
}

// Error stringifies the APIError
func (apiError APIError) Error() string {
	return apiError.Message
}

////////////////////////////
// Client                 //
////////////////////////////

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests
	RestyClient *resty.Client
}

////////////////////////////
// Person                //
////////////////////////////

type Emails struct {
	AdditionalEmails *[]openapi_types.Email `json:"additionalEmails,omitempty"`
	PrimaryEmail     *string                `json:"primaryEmail,omitempty"`
}

type Name struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

type Phones struct {
	AdditionalPhones        *[]string `json:"additionalPhones,omitempty"`
	PrimaryPhoneCallingCode *string   `json:"primaryPhoneCallingCode,omitempty"`
	PrimaryPhoneCountryCode *string   `json:"primaryPhoneCountryCode,omitempty"`
	PrimaryPhoneNumber      *string   `json:"primaryPhoneNumber,omitempty"`
}

// NewPerson A Person
type NewPerson struct {
	// AvatarUrl Contact’s avatar
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// City Contact’s city
	City      *string             `json:"city,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *NewCreatedBy `json:"createdBy,omitempty"`

	// Emails Contact’s Emails
	Emails *Emails `json:"emails,omitempty"`

	// JobTitle Contact’s job title
	JobTitle   *string `json:"jobTitle,omitempty"`
	Keycloakid *string `json:"keycloakid,omitempty"`

	// LinkedinLink Contact’s Linkedin account
	LinkedinLink *Link `json:"linkedinLink,omitempty"`

	// Name Contact’s name
	Name       *Name `json:"name,omitempty"`
	Newsletter *bool `json:"newsletter,omitempty"`

	// Phones Contact’s phone numbers
	Phones *Phones `json:"phones,omitempty"`

	// Position Person record Position
	Position *float32 `json:"position,omitempty"`

	// XLink Contact’s X/Twitter account
	XLink *Link `json:"xLink,omitempty"`
}

// Person A person
type Person struct {
	// Attachments Attachments linked to the contact.
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// AvatarUrl Contact’s avatar
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// CalendarEventParticipants Calendar Event Participants
	CalendarEventParticipants *[]CalendarEventParticipant `json:"calendarEventParticipants,omitempty"`

	// City Contact’s city
	City *string `json:"city,omitempty"`

	// Company Contact’s company
	Company   *Company            `json:"company,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Emails Contact’s Emails
	Emails *Emails `json:"emails,omitempty"`

	// Favorites Favorites linked to the contact
	Favorites *[]Favorite `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// JobTitle Contact’s job title
	JobTitle   *string `json:"jobTitle,omitempty"`
	Keycloakid *string `json:"keycloakid,omitempty"`

	// LinkedinLink Contact’s Linkedin account
	LinkedinLink *Link `json:"linkedinLink,omitempty"`

	// MessageParticipants Message Participants
	MessageParticipants *[]MessageParticipant `json:"messageParticipants,omitempty"`

	// Name Contact’s name
	Name       *Name `json:"name,omitempty"`
	Newsletter *bool `json:"newsletter,omitempty"`

	// NoteTargets Notes tied to the contact
	NoteTargets *[]NoteTarget `json:"noteTargets,omitempty"`

	// Phones Contact’s phone numbers
	Phones *Phones `json:"phones,omitempty"`

	// PointOfContactForOpportunities List of opportunities for which that person is the point of contact
	PointOfContactForOpportunities *[]Opportunity `json:"pointOfContactForOpportunities,omitempty"`

	// Position Person record Position
	Position *float32 `json:"position,omitempty"`

	// TaskTargets Tasks tied to the contact
	TaskTargets *[]TaskTarget `json:"taskTargets,omitempty"`

	// TimelineActivities Events linked to the person
	TimelineActivities *[]TimelineActivity `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// XLink Contact’s X/Twitter account
	XLink *Link `json:"xLink,omitempty"`
}

////////////////////////////
// Attachments            //
////////////////////////////

// AttachmentFileCategory Attachment file category
type AttachmentFileCategory string

// NewAttachment An attachment
type NewAttachment struct {
	AuthorId  *openapi_types.UUID `json:"authorId,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy   *NewCreatedBy       `json:"createdBy,omitempty"`
	DashboardId *openapi_types.UUID `json:"dashboardId,omitempty"`

	// FileCategory Attachment file category
	FileCategory *AttachmentFileCategory `json:"fileCategory,omitempty"`

	// FullPath Attachment full path
	FullPath *string `json:"fullPath,omitempty"`

	// Name Attachment name
	Name          *string             `json:"name,omitempty"`
	NoteId        *openapi_types.UUID `json:"noteId,omitempty"`
	OpportunityId *openapi_types.UUID `json:"opportunityId,omitempty"`
	PersonId      *openapi_types.UUID `json:"personId,omitempty"`
	TaskId        *openapi_types.UUID `json:"taskId,omitempty"`

	// Type Attachment type (deprecated - use fileCategory)
	Type       *string             `json:"type,omitempty"`
	WorkflowId *openapi_types.UUID `json:"workflowId,omitempty"`
}

// Attachment An attachment
type Attachment struct {
	// Author Attachment author (deprecated - use createdBy)
	Author   *WorkspaceMember    `json:"author,omitempty"`
	AuthorId *openapi_types.UUID `json:"authorId,omitempty"`

	// Company Attachment company
	Company   *Company            `json:"company,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`

	// Dashboard Attachment dashboard
	Dashboard   *Dashboard          `json:"dashboard,omitempty"`
	DashboardId *openapi_types.UUID `json:"dashboardId,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// FileCategory Attachment file category
	FileCategory *AttachmentFileCategory `json:"fileCategory,omitempty"`

	// FullPath Attachment full path
	FullPath *string `json:"fullPath,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Name Attachment name
	Name *string `json:"name,omitempty"`

	// Note Attachment note
	Note   *Note               `json:"note,omitempty"`
	NoteId *openapi_types.UUID `json:"noteId,omitempty"`

	// Opportunity Attachment opportunity
	Opportunity   *Opportunity        `json:"opportunity,omitempty"`
	OpportunityId *openapi_types.UUID `json:"opportunityId,omitempty"`

	// Person Attachment person
	Person   *Person             `json:"person,omitempty"`
	PersonId *openapi_types.UUID `json:"personId,omitempty"`

	// Task Attachment task
	Task   *Task               `json:"task,omitempty"`
	TaskId *openapi_types.UUID `json:"taskId,omitempty"`

	// Type Attachment type (deprecated - use fileCategory)
	Type *string `json:"type,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Workflow Attachment workflow
	Workflow   *Workflow           `json:"workflow,omitempty"`
	WorkflowId *openapi_types.UUID `json:"workflowId,omitempty"`
}

////////////////////////////
// CalendarEventParticipant //
////////////////////////////

// CalendarEventParticipant Calendar event participants
type NewCalendarEventParticipant struct {
	CalendarEventId *openapi_types.UUID `json:"calendarEventId,omitempty"`

	// DisplayName Display Name
	DisplayName *string `json:"displayName,omitempty"`

	// Handle Handle
	Handle *string `json:"handle,omitempty"`

	// IsOrganizer Is Organizer
	IsOrganizer *bool               `json:"isOrganizer,omitempty"`
	PersonId    *openapi_types.UUID `json:"personId,omitempty"`

	// ResponseStatus Response Status
	ResponseStatus    *CalendarEventParticipantResponseStatus `json:"responseStatus,omitempty"`
	WorkspaceMemberId *openapi_types.UUID                     `json:"workspaceMemberId,omitempty"`
}

// CalendarEventParticipantResponseStatus Response Status
type CalendarEventParticipantResponseStatus string

// CalendarEventParticipantForResponse Calendar event participants
type CalendarEventParticipant struct {
	// CalendarEvent Event ID
	CalendarEvent   *CalendarEvent      `json:"calendarEvent,omitempty"`
	CalendarEventId *openapi_types.UUID `json:"calendarEventId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// DisplayName Display Name
	DisplayName *string `json:"displayName,omitempty"`

	// Handle Handle
	Handle *string `json:"handle,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// IsOrganizer Is Organizer
	IsOrganizer *bool `json:"isOrganizer,omitempty"`

	// Person Person
	Person   *Person             `json:"person,omitempty"`
	PersonId *openapi_types.UUID `json:"personId,omitempty"`

	// ResponseStatus Response Status
	ResponseStatus *CalendarEventParticipantResponseStatus `json:"responseStatus,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// WorkspaceMember Workspace Member
	WorkspaceMember   *WorkspaceMember    `json:"workspaceMember,omitempty"`
	WorkspaceMemberId *openapi_types.UUID `json:"workspaceMemberId,omitempty"`
}

////////////////////////////
// WorkspaceMember        //
////////////////////////////

// NewWorkspaceMember A workspace member
type NewWorkspaceMember struct {
	// AvatarUrl Workspace member avatar
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// CalendarStartDay User's preferred start day of the week
	CalendarStartDay *int `json:"calendarStartDay,omitempty"`

	// ColorScheme Preferred color scheme
	ColorScheme *string `json:"colorScheme,omitempty"`

	// DateFormat User's preferred date format
	DateFormat *WorkspaceMemberDateFormat `json:"dateFormat,omitempty"`

	// Locale Preferred language
	Locale *string `json:"locale,omitempty"`

	// Name Workspace member name
	Name Name `json:"name"`

	// NumberFormat User's preferred number format
	NumberFormat *WorkspaceMemberNumberFormat `json:"numberFormat,omitempty"`

	// Position Workspace member position
	Position *float32 `json:"position,omitempty"`

	// TimeFormat User's preferred time format
	TimeFormat *WorkspaceMemberTimeFormat `json:"timeFormat,omitempty"`

	// TimeZone User time zone
	TimeZone *string `json:"timeZone,omitempty"`

	// UserEmail Related user email address
	UserEmail *string `json:"userEmail,omitempty"`

	// UserId Associated User Id
	UserId openapi_types.UUID `json:"userId"`
}

// WorkspaceMemberDateFormat User's preferred date format
type WorkspaceMemberDateFormat string

// WorkspaceMemberNumberFormat User's preferred number format
type WorkspaceMemberNumberFormat string

// WorkspaceMemberTimeFormat User's preferred time format
type WorkspaceMemberTimeFormat string

// WorkspaceMember A workspace member
type WorkspaceMember struct {
	// AccountOwnerForCompanies Account owner for companies
	AccountOwnerForCompanies *[]Company `json:"accountOwnerForCompanies,omitempty"`

	// AssignedTasks Tasks assigned to the workspace member
	AssignedTasks *[]Task `json:"assignedTasks,omitempty"`

	// AuthoredAttachments Attachments created by the workspace member
	AuthoredAttachments *[]Attachment `json:"authoredAttachments,omitempty"`

	// AvatarUrl Workspace member avatar
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// Blocklist Blocklisted handles
	Blocklist *[]Blocklist `json:"blocklist,omitempty"`

	// CalendarEventParticipants Calendar Event Participants
	CalendarEventParticipants *[]CalendarEventParticipant `json:"calendarEventParticipants,omitempty"`

	// CalendarStartDay User's preferred start day of the week
	CalendarStartDay *int `json:"calendarStartDay,omitempty"`

	// ColorScheme Preferred color scheme
	ColorScheme *string `json:"colorScheme,omitempty"`

	// ConnectedAccounts Connected accounts
	ConnectedAccounts *[]ConnectedAccount `json:"connectedAccounts,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DateFormat User's preferred date format
	DateFormat *WorkspaceMemberDateFormat `json:"dateFormat,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the workspace member
	Favorites *[]Favorite `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Locale Preferred language
	Locale *string `json:"locale,omitempty"`

	// MessageParticipants Message Participants
	MessageParticipants *[]MessageParticipant `json:"messageParticipants,omitempty"`

	// Name Workspace member name
	Name *Name `json:"name,omitempty"`

	// NumberFormat User's preferred number format
	NumberFormat *WorkspaceMemberNumberFormat `json:"numberFormat,omitempty"`

	// Position Workspace member position
	Position *float32 `json:"position,omitempty"`

	// TimeFormat User's preferred time format
	TimeFormat *WorkspaceMemberTimeFormat `json:"timeFormat,omitempty"`

	// TimeZone User time zone
	TimeZone *string `json:"timeZone,omitempty"`

	// TimelineActivities Events linked to the workspace member
	TimelineActivities *[]TimelineActivity `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// UserEmail Related user email address
	UserEmail *string `json:"userEmail,omitempty"`

	// UserId Associated User Id
	UserId *openapi_types.UUID `json:"userId,omitempty"`
}

////////////////////////////
// Company                   //
////////////////////////////

type Address struct {
	AddressCity     *string  `json:"addressCity,omitempty"`
	AddressCountry  *string  `json:"addressCountry,omitempty"`
	AddressLat      *float32 `json:"addressLat,omitempty"`
	AddressLng      *float32 `json:"addressLng,omitempty"`
	AddressPostcode *string  `json:"addressPostcode,omitempty"`
	AddressState    *string  `json:"addressState,omitempty"`
	AddressStreet1  *string  `json:"addressStreet1,omitempty"`
	AddressStreet2  *string  `json:"addressStreet2,omitempty"`
}

type AnnualRecurringRevenue struct {
	AmountMicros *float32 `json:"amountMicros,omitempty"`
	CurrencyCode *string  `json:"currencyCode,omitempty"`
}

// NewCompany A company
type NewCompany struct {
	AccountOwnerId *openapi_types.UUID `json:"accountOwnerId,omitempty"`

	// Address Address of the company
	Address *Address `json:"address,omitempty"`

	// AnnualRecurringRevenue Annual Recurring Revenue: The actual or estimated annual revenue of the company
	AnnualRecurringRevenue *AnnualRecurringRevenue `json:"annualRecurringRevenue,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *NewCreatedBy `json:"createdBy,omitempty"`

	// DomainName The company website URL. We use this url to fetch the company icon
	DomainName *Link `json:"domainName,omitempty"`

	// Employees Number of employees in the company
	Employees *int `json:"employees,omitempty"`

	// IdealCustomerProfile Ideal Customer Profile:  Indicates whether the company is the most suitable and valuable customer for you
	IdealCustomerProfile *bool `json:"idealCustomerProfile,omitempty"`

	// LinkedinLink The company Linkedin account
	LinkedinLink *Link `json:"linkedinLink,omitempty"`

	// Name The company name
	Name *string `json:"name,omitempty"`

	// Position Company record position
	Position *float32 `json:"position,omitempty"`

	// XLink The company Twitter/X account
	XLink *Link `json:"xLink,omitempty"`
}

// CompanyForResponse A company
type Company struct {
	// AccountOwner Your team member responsible for managing the company account
	AccountOwner   *WorkspaceMember    `json:"accountOwner,omitempty"`
	AccountOwnerId *openapi_types.UUID `json:"accountOwnerId,omitempty"`

	// Address Address of the company
	Address *Address `json:"address,omitempty"`

	// AnnualRecurringRevenue Annual Recurring Revenue: The actual or estimated annual revenue of the company
	AnnualRecurringRevenue *AnnualRecurringRevenue `json:"annualRecurringRevenue,omitempty"`

	// Attachments Attachments linked to the company
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// DomainName The company website URL. We use this url to fetch the company icon
	DomainName *Link `json:"domainName,omitempty"`

	// Employees Number of employees in the company
	Employees *int `json:"employees,omitempty"`

	// Favorites Favorites linked to the company
	Favorites *[]Favorite `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// IdealCustomerProfile Ideal Customer Profile:  Indicates whether the company is the most suitable and valuable customer for you
	IdealCustomerProfile *bool `json:"idealCustomerProfile,omitempty"`

	// LinkedinLink The company Linkedin account
	LinkedinLink *Link `json:"linkedinLink,omitempty"`

	// Name The company name
	Name *string `json:"name,omitempty"`

	// NoteTargets Notes tied to the company
	NoteTargets *[]NoteTarget `json:"noteTargets,omitempty"`

	// Opportunities Opportunities linked to the company.
	Opportunities *[]Opportunity `json:"opportunities,omitempty"`

	// People People linked to the company.
	People *[]Person `json:"people,omitempty"`

	// Position Company record position
	Position *float32 `json:"position,omitempty"`

	// TaskTargets Tasks tied to the company
	TaskTargets *[]TaskTarget `json:"taskTargets,omitempty"`

	// TimelineActivities Timeline Activities linked to the company
	TimelineActivities *[]TimelineActivity `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// XLink The company Twitter/X account
	XLink *Link `json:"xLink,omitempty"`
}

////////////////////////////
// Favorite                   //
////////////////////////////

// NewFavorite A favorite that can be accessed from the left menu
type NewFavorite struct {
	CompanyId            *openapi_types.UUID `json:"companyId,omitempty"`
	DashboardId          *openapi_types.UUID `json:"dashboardId,omitempty"`
	FavoriteFolderId     *openapi_types.UUID `json:"favoriteFolderId,omitempty"`
	ForWorkspaceMemberId *openapi_types.UUID `json:"forWorkspaceMemberId,omitempty"`
	NoteId               *openapi_types.UUID `json:"noteId,omitempty"`
	OpportunityId        *openapi_types.UUID `json:"opportunityId,omitempty"`
	PersonId             *openapi_types.UUID `json:"personId,omitempty"`

	// Position Favorite position
	Position *int                `json:"position,omitempty"`
	TaskId   *openapi_types.UUID `json:"taskId,omitempty"`

	// ViewId ViewId
	ViewId            *openapi_types.UUID `json:"viewId,omitempty"`
	WorkflowId        *openapi_types.UUID `json:"workflowId,omitempty"`
	WorkflowRunId     *openapi_types.UUID `json:"workflowRunId,omitempty"`
	WorkflowVersionId *openapi_types.UUID `json:"workflowVersionId,omitempty"`
}

// FavoriteFolder A Folder of favorites
type NewFavoriteFolder struct {
	// Name Name of the favorite folder
	Name *string `json:"name,omitempty"`

	// Position Favorite folder position
	Position *int `json:"position,omitempty"`
}

// FavoriteFolderForResponse A Folder of favorites
type FavoriteFolder struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites in this folder
	Favorites *[]Favorite `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Name Name of the favorite folder
	Name *string `json:"name,omitempty"`

	// Position Favorite folder position
	Position *int `json:"position,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// FavoriteForResponse A favorite that can be accessed from the left menu
type Favorite struct {
	// Company Favorite company
	Company   *Company            `json:"company,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Dashboard Favorite dashboard
	Dashboard   *Dashboard          `json:"dashboard,omitempty"`
	DashboardId *openapi_types.UUID `json:"dashboardId,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// FavoriteFolder The folder this favorite belongs to
	FavoriteFolder   *FavoriteFolder     `json:"favoriteFolder,omitempty"`
	FavoriteFolderId *openapi_types.UUID `json:"favoriteFolderId,omitempty"`

	// ForWorkspaceMember Favorite workspace member
	ForWorkspaceMember   *WorkspaceMember    `json:"forWorkspaceMember,omitempty"`
	ForWorkspaceMemberId *openapi_types.UUID `json:"forWorkspaceMemberId,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Note Favorite note
	Note   *Note               `json:"note,omitempty"`
	NoteId *openapi_types.UUID `json:"noteId,omitempty"`

	// Opportunity Favorite opportunity
	Opportunity   *Opportunity        `json:"opportunity,omitempty"`
	OpportunityId *openapi_types.UUID `json:"opportunityId,omitempty"`

	// Person Favorite person
	Person   *Person             `json:"person,omitempty"`
	PersonId *openapi_types.UUID `json:"personId,omitempty"`

	// Position Favorite position
	Position *int `json:"position,omitempty"`

	// Task Favorite task
	Task   *Task               `json:"task,omitempty"`
	TaskId *openapi_types.UUID `json:"taskId,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// ViewId ViewId
	ViewId *openapi_types.UUID `json:"viewId,omitempty"`

	// Workflow Favorite workflow
	Workflow   *Workflow           `json:"workflow,omitempty"`
	WorkflowId *openapi_types.UUID `json:"workflowId,omitempty"`

	// WorkflowRun Favorite workflow run
	WorkflowRun   *WorkflowRun        `json:"workflowRun,omitempty"`
	WorkflowRunId *openapi_types.UUID `json:"workflowRunId,omitempty"`

	// WorkflowVersion Favorite workflow version
	WorkflowVersion   *WorkflowVersion    `json:"workflowVersion,omitempty"`
	WorkflowVersionId *openapi_types.UUID `json:"workflowVersionId,omitempty"`
}

////////////////////////////
// MessageParticipant     //
////////////////////////////

// MessageParticipant Message Participants
type NewMessageParticipant struct {
	// DisplayName Display Name
	DisplayName *string `json:"displayName,omitempty"`

	// Handle Handle
	Handle    *string             `json:"handle,omitempty"`
	MessageId *openapi_types.UUID `json:"messageId,omitempty"`
	PersonId  *openapi_types.UUID `json:"personId,omitempty"`

	// Role Role
	Role              *MessageParticipantRole `json:"role,omitempty"`
	WorkspaceMemberId *openapi_types.UUID     `json:"workspaceMemberId,omitempty"`
}

// MessageParticipantRole Role
type MessageParticipantRole string

// MessageParticipantForResponse Message Participants
type MessageParticipant struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// DisplayName Display Name
	DisplayName *string `json:"displayName,omitempty"`

	// Handle Handle
	Handle *string `json:"handle,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Message Message
	Message   *Message            `json:"message,omitempty"`
	MessageId *openapi_types.UUID `json:"messageId,omitempty"`

	// Person Person
	Person   *Person             `json:"person,omitempty"`
	PersonId *openapi_types.UUID `json:"personId,omitempty"`

	// Role Role
	Role *MessageParticipantRole `json:"role,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// WorkspaceMember Workspace member
	WorkspaceMember   *WorkspaceMember    `json:"workspaceMember,omitempty"`
	WorkspaceMemberId *openapi_types.UUID `json:"workspaceMemberId,omitempty"`
}

////////////////////////////
// NoteTarget             //
////////////////////////////

// NoteTarget A note target
type NewNoteTarget struct {
	CompanyId     *openapi_types.UUID `json:"companyId,omitempty"`
	NoteId        *openapi_types.UUID `json:"noteId,omitempty"`
	OpportunityId *openapi_types.UUID `json:"opportunityId,omitempty"`
	PersonId      *openapi_types.UUID `json:"personId,omitempty"`
}

// NoteTargetForResponse A note target
type NoteTarget struct {
	// Company NoteTarget company
	Company   *Company            `json:"company,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Note NoteTarget note
	Note   *Note               `json:"note,omitempty"`
	NoteId *openapi_types.UUID `json:"noteId,omitempty"`

	// Opportunity NoteTarget opportunity
	Opportunity   *Opportunity        `json:"opportunity,omitempty"`
	OpportunityId *openapi_types.UUID `json:"opportunityId,omitempty"`

	// Person NoteTarget person
	Person   *Person             `json:"person,omitempty"`
	PersonId *openapi_types.UUID `json:"personId,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

////////////////////////////
// Opportunity            //
////////////////////////////

type Amount struct {
	AmountMicros *float32 `json:"amountMicros,omitempty"`
	CurrencyCode *string  `json:"currencyCode,omitempty"`
}

// Opportunity An opportunity
type NewOpportunity struct {
	// Amount Opportunity amount
	Amount *Amount `json:"amount,omitempty"`

	// CloseDate Opportunity close date
	CloseDate *time.Time          `json:"closeDate,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy  *NewCreatedBy          `json:"createdBy,omitempty"`
	LostReason *OpportunityLostReason `json:"lostReason,omitempty"`

	// Name The opportunity name
	Name             *string             `json:"name,omitempty"`
	Organizationid   *string             `json:"organizationid,omitempty"`
	PointOfContactId *openapi_types.UUID `json:"pointOfContactId,omitempty"`

	// Position Opportunity record position
	Position *float32 `json:"position,omitempty"`

	// Stage Opportunity stage
	Stage              *OpportunityStage `json:"stage,omitempty"`
	SubscriptionActive *bool             `json:"subscriptionActive,omitempty"`
	SubscriptionAmount *int              `json:"subscriptionAmount,omitempty"`
	SubscriptionEnd    *string           `json:"subscriptionEnd,omitempty"`
	SubscriptionStart  *string           `json:"subscriptionStart,omitempty"`
}

// OpportunityLostReason defines model for Opportunity.LostReason.
type OpportunityLostReason string

// OpportunityStage Opportunity stage
type OpportunityStage string

// OpportunityForResponse An opportunity
type Opportunity struct {
	// Amount Opportunity amount
	Amount *struct {
		AmountMicros *float32 `json:"amountMicros,omitempty"`
		CurrencyCode *string  `json:"currencyCode,omitempty"`
	} `json:"amount,omitempty"`

	// Attachments Attachments linked to the opportunity
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// CloseDate Opportunity close date
	CloseDate *time.Time `json:"closeDate,omitempty"`

	// Company Opportunity company
	Company   *Company            `json:"company,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the opportunity
	Favorites *[]Favorite `json:"favorites,omitempty"`

	// Id Id
	Id         *openapi_types.UUID    `json:"id,omitempty"`
	LostReason *OpportunityLostReason `json:"lostReason,omitempty"`

	// Name The opportunity name
	Name *string `json:"name,omitempty"`

	// NoteTargets Notes tied to the opportunity
	NoteTargets    *[]NoteTarget `json:"noteTargets,omitempty"`
	Organizationid *string       `json:"organizationid,omitempty"`

	// PointOfContact Opportunity point of contact
	PointOfContact   *Person             `json:"pointOfContact,omitempty"`
	PointOfContactId *openapi_types.UUID `json:"pointOfContactId,omitempty"`

	// Position Opportunity record position
	Position *float32 `json:"position,omitempty"`

	// Stage Opportunity stage
	Stage              *OpportunityStage `json:"stage,omitempty"`
	SubscriptionActive *bool             `json:"subscriptionActive,omitempty"`
	SubscriptionAmount *int              `json:"subscriptionAmount,omitempty"`
	SubscriptionEnd    *string           `json:"subscriptionEnd,omitempty"`
	SubscriptionStart  *string           `json:"subscriptionStart,omitempty"`

	// TaskTargets Tasks tied to the opportunity
	TaskTargets *[]TaskTarget `json:"taskTargets,omitempty"`

	// TimelineActivities Timeline Activities linked to the opportunity.
	TimelineActivities *[]TimelineActivity `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

////////////////////////////
// TaskTarget                   //
////////////////////////////

// NewTaskTarget A task target
type NewTaskTarget struct {
	CompanyId     *openapi_types.UUID `json:"companyId,omitempty"`
	OpportunityId *openapi_types.UUID `json:"opportunityId,omitempty"`
	PersonId      *openapi_types.UUID `json:"personId,omitempty"`
	TaskId        *openapi_types.UUID `json:"taskId,omitempty"`
}

// TaskTarget A task target
type TaskTarget struct {
	// Company TaskTarget company
	Company   *Company            `json:"company,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Opportunity TaskTarget opportunity
	Opportunity   *Opportunity        `json:"opportunity,omitempty"`
	OpportunityId *openapi_types.UUID `json:"opportunityId,omitempty"`

	// Person TaskTarget person
	Person   *Person             `json:"person,omitempty"`
	PersonId *openapi_types.UUID `json:"personId,omitempty"`

	// Task TaskTarget task
	Task   *Task               `json:"task,omitempty"`
	TaskId *openapi_types.UUID `json:"taskId,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

////////////////////////////
// TimelineActivity       //
////////////////////////////

// NewTimelineActivity Aggregated / filtered event to be displayed on the timeline
type NewTimelineActivity struct {
	// HappensAt Creation date
	HappensAt *time.Time `json:"happensAt,omitempty"`

	// LinkedObjectMetadataId Linked Object Metadata Id
	LinkedObjectMetadataId *openapi_types.UUID `json:"linkedObjectMetadataId,omitempty"`

	// LinkedRecordCachedName Cached record name
	LinkedRecordCachedName *string `json:"linkedRecordCachedName,omitempty"`

	// LinkedRecordId Linked Record id
	LinkedRecordId *openapi_types.UUID `json:"linkedRecordId,omitempty"`

	// Name Event name
	Name *string `json:"name,omitempty"`

	// Properties Json value for event details
	Properties *map[string]interface{} `json:"properties,omitempty"`

	// TargetCompany Event company
	TargetCompany *string `json:"targetCompany,omitempty"`

	// TargetDashboard Event dashboard
	TargetDashboard *string `json:"targetDashboard,omitempty"`

	// TargetNote Event note
	TargetNote *string `json:"targetNote,omitempty"`

	// TargetOpportunity Event opportunity
	TargetOpportunity *string `json:"targetOpportunity,omitempty"`

	// TargetPerson Event person
	TargetPerson *string `json:"targetPerson,omitempty"`

	// TargetTask Event task
	TargetTask *string `json:"targetTask,omitempty"`

	// TargetWorkflow Event workflow
	TargetWorkflow *string `json:"targetWorkflow,omitempty"`

	// TargetWorkflowRun Event workflow run
	TargetWorkflowRun *string `json:"targetWorkflowRun,omitempty"`

	// TargetWorkflowVersion Event workflow version
	TargetWorkflowVersion *string             `json:"targetWorkflowVersion,omitempty"`
	WorkspaceMemberId     *openapi_types.UUID `json:"workspaceMemberId,omitempty"`
}

// TimelineActivityForResponse Aggregated / filtered event to be displayed on the timeline
type TimelineActivity struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// HappensAt Creation date
	HappensAt *time.Time `json:"happensAt,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// LinkedObjectMetadataId Linked Object Metadata Id
	LinkedObjectMetadataId *openapi_types.UUID `json:"linkedObjectMetadataId,omitempty"`

	// LinkedRecordCachedName Cached record name
	LinkedRecordCachedName *string `json:"linkedRecordCachedName,omitempty"`

	// LinkedRecordId Linked Record id
	LinkedRecordId *openapi_types.UUID `json:"linkedRecordId,omitempty"`

	// Name Event name
	Name *string `json:"name,omitempty"`

	// Properties Json value for event details
	Properties *map[string]interface{} `json:"properties,omitempty"`

	// TargetCompany Event company
	TargetCompany *string `json:"targetCompany,omitempty"`

	// TargetDashboard Event dashboard
	TargetDashboard *string `json:"targetDashboard,omitempty"`

	// TargetNote Event note
	TargetNote *string `json:"targetNote,omitempty"`

	// TargetOpportunity Event opportunity
	TargetOpportunity *string `json:"targetOpportunity,omitempty"`

	// TargetPerson Event person
	TargetPerson *string `json:"targetPerson,omitempty"`

	// TargetTask Event task
	TargetTask *string `json:"targetTask,omitempty"`

	// TargetWorkflow Event workflow
	TargetWorkflow *string `json:"targetWorkflow,omitempty"`

	// TargetWorkflowRun Event workflow run
	TargetWorkflowRun *string `json:"targetWorkflowRun,omitempty"`

	// TargetWorkflowVersion Event workflow version
	TargetWorkflowVersion *string `json:"targetWorkflowVersion,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// WorkspaceMember Event workspace member
	WorkspaceMember   *WorkspaceMember    `json:"workspaceMember,omitempty"`
	WorkspaceMemberId *openapi_types.UUID `json:"workspaceMemberId,omitempty"`
}

////////////////////////////
// Dashboard              //
////////////////////////////

// NewDashboard A dashboard
type NewDashboard struct {
	// CreatedBy The creator of the record
	CreatedBy *NewCreatedBy `json:"createdBy,omitempty"`

	// PageLayoutId Dashboard page layout
	PageLayoutId *openapi_types.UUID `json:"pageLayoutId,omitempty"`

	// Position Dashboard record Position
	Position *float32 `json:"position,omitempty"`

	// Title Dashboard title
	Title *string `json:"title,omitempty"`
}

// Dashboard A dashboard
type Dashboard struct {
	// Attachments Attachments linked to the dashboard
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the dashboard
	Favorites *[]Favorite `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// PageLayoutId Dashboard page layout
	PageLayoutId *openapi_types.UUID `json:"pageLayoutId,omitempty"`

	// Position Dashboard record Position
	Position *float32 `json:"position,omitempty"`

	// TimelineActivities Timeline activities linked to the dashboard
	TimelineActivities *[]TimelineActivity `json:"timelineActivities,omitempty"`

	// Title Dashboard title
	Title *string `json:"title,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

////////////////////////////
// Note                   //
////////////////////////////

type BodyV2 struct {
	Blocknote *string `json:"blocknote,omitempty"`
	Markdown  *string `json:"markdown,omitempty"`
}

// NewNote A note
type NewNote struct {
	// BodyV2 Note body
	BodyV2 *BodyV2 `json:"bodyV2,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *NewCreatedBy `json:"createdBy,omitempty"`

	// Position Note record position
	Position *float32 `json:"position,omitempty"`

	// Title Note title
	Title *string `json:"title,omitempty"`
}

// Note A note
type Note struct {
	// Attachments Note attachments
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// BodyV2 Note body
	BodyV2 *BodyV2 `json:"bodyV2,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the note
	Favorites *[]Favorite `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// NoteTargets Note targets
	NoteTargets *[]NoteTarget `json:"noteTargets,omitempty"`

	// Position Note record position
	Position *float32 `json:"position,omitempty"`

	// TimelineActivities Timeline Activities linked to the note.
	TimelineActivities *[]TimelineActivity `json:"timelineActivities,omitempty"`

	// Title Note title
	Title *string `json:"title,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

////////////////////////////
// Task                   //
////////////////////////////

// NewTask A task
type NewTask struct {
	AssigneeId *openapi_types.UUID `json:"assigneeId,omitempty"`

	// BodyV2 Task body
	BodyV2 *BodyV2 `json:"bodyV2,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *NewCreatedBy `json:"createdBy,omitempty"`

	// DueAt Task due date
	DueAt *time.Time `json:"dueAt,omitempty"`

	// Position Task record position
	Position *float32 `json:"position,omitempty"`

	// Status Task status
	Status *TaskStatus `json:"status,omitempty"`

	// Title Task title
	Title *string `json:"title,omitempty"`
}

// TaskStatus Task status
type TaskStatus string

// Task A task
type Task struct {
	// Assignee Task assignee
	Assignee   *WorkspaceMember    `json:"assignee,omitempty"`
	AssigneeId *openapi_types.UUID `json:"assigneeId,omitempty"`

	// Attachments Task attachments
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// BodyV2 Task body
	BodyV2 *BodyV2 `json:"bodyV2,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// DueAt Task due date
	DueAt *time.Time `json:"dueAt,omitempty"`

	// Favorites Favorites linked to the task
	Favorites *[]Favorite `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Position Task record position
	Position *float32 `json:"position,omitempty"`

	// Status Task status
	Status *TaskStatus `json:"status,omitempty"`

	// TaskTargets Task targets
	TaskTargets *[]TaskTarget `json:"taskTargets,omitempty"`

	// TimelineActivities Timeline Activities linked to the task.
	TimelineActivities *[]TimelineActivity `json:"timelineActivities,omitempty"`

	// Title Task title
	Title *string `json:"title,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

////////////////////////////
// Workflow               //
////////////////////////////

// NewWorkflow A workflow
type NewWorkflow struct {
	// CreatedBy The creator of the record
	CreatedBy *NewCreatedBy `json:"createdBy,omitempty"`

	// LastPublishedVersionId The workflow last published version id
	LastPublishedVersionId *string `json:"lastPublishedVersionId,omitempty"`

	// Name The workflow name
	Name *string `json:"name,omitempty"`

	// Position Workflow record position
	Position *float32 `json:"position,omitempty"`

	// Statuses The current statuses of the workflow versions
	Statuses *[]WorkflowStatuses `json:"statuses,omitempty"`
}

// WorkflowStatuses defines model for Workflow.Statuses.
type WorkflowStatuses string

// WorkflowAutomatedTrigger A workflow automated trigger
type NewWorkflowAutomatedTrigger struct {
	// Settings The workflow automated trigger settings
	Settings map[string]interface{} `json:"settings"`

	// Type The workflow automated trigger type
	Type       WorkflowAutomatedTriggerType `json:"type"`
	WorkflowId *openapi_types.UUID          `json:"workflowId,omitempty"`
}

// WorkflowAutomatedTriggerType The workflow automated trigger type
type WorkflowAutomatedTriggerType string

// WorkflowAutomatedTriggerForResponse A workflow automated trigger
type WorkflowAutomatedTrigger struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Settings The workflow automated trigger settings
	Settings *map[string]interface{} `json:"settings,omitempty"`

	// Type The workflow automated trigger type
	Type *WorkflowAutomatedTriggerType `json:"type,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Workflow WorkflowAutomatedTrigger workflow
	Workflow   *Workflow           `json:"workflow,omitempty"`
	WorkflowId *openapi_types.UUID `json:"workflowId,omitempty"`
}

// Workflow A workflow
type Workflow struct {
	// Attachments Attachments linked to the workflow
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// AutomatedTriggers Workflow automated triggers linked to the workflow.
	AutomatedTriggers *[]WorkflowAutomatedTrigger `json:"automatedTriggers,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the workflow
	Favorites *[]Favorite `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// LastPublishedVersionId The workflow last published version id
	LastPublishedVersionId *string `json:"lastPublishedVersionId,omitempty"`

	// Name The workflow name
	Name *string `json:"name,omitempty"`

	// Position Workflow record position
	Position *float32 `json:"position,omitempty"`

	// Runs Workflow runs linked to the workflow.
	Runs *[]WorkflowRun `json:"runs,omitempty"`

	// Statuses The current statuses of the workflow versions
	Statuses *[]WorkflowStatuses `json:"statuses,omitempty"`

	// TimelineActivities Timeline activities linked to the workflow
	TimelineActivities *[]TimelineActivity `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Versions Workflow versions linked to the workflow.
	Versions *[]WorkflowVersion `json:"versions,omitempty"`
}

// NewWorkflowVersion A workflow version
type NewWorkflowVersion struct {
	// Name The workflow version name
	Name *string `json:"name,omitempty"`

	// Position Workflow version position
	Position *float32 `json:"position,omitempty"`

	// Status The workflow version status
	Status *WorkflowVersionStatus `json:"status,omitempty"`

	// Steps Json object to provide steps
	Steps *map[string]interface{} `json:"steps,omitempty"`

	// Trigger Json object to provide trigger
	Trigger    *map[string]interface{} `json:"trigger,omitempty"`
	WorkflowId *openapi_types.UUID     `json:"workflowId,omitempty"`
}

// WorkflowVersionStatus The workflow version status
type WorkflowVersionStatus string

// WorkflowVersion A workflow version
type WorkflowVersion struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the workflow version
	Favorites *[]Favorite `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Name The workflow version name
	Name *string `json:"name,omitempty"`

	// Position Workflow version position
	Position *float32 `json:"position,omitempty"`

	// Runs Workflow runs linked to the version.
	Runs *[]WorkflowRun `json:"runs,omitempty"`

	// Status The workflow version status
	Status *WorkflowVersionStatus `json:"status,omitempty"`

	// Steps Json object to provide steps
	Steps *map[string]interface{} `json:"steps,omitempty"`

	// TimelineActivities Timeline activities linked to the version
	TimelineActivities *[]TimelineActivity `json:"timelineActivities,omitempty"`

	// Trigger Json object to provide trigger
	Trigger *map[string]interface{} `json:"trigger,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Workflow WorkflowVersion workflow
	Workflow   *Workflow           `json:"workflow,omitempty"`
	WorkflowId *openapi_types.UUID `json:"workflowId,omitempty"`
}

// NewWorkflowRun A workflow run
type NewWorkflowRun struct {
	// CreatedBy The executor of the workflow
	CreatedBy *NewCreatedBy `json:"createdBy,omitempty"`

	// EndedAt Workflow run ended at
	EndedAt *time.Time `json:"endedAt,omitempty"`

	// EnqueuedAt Workflow run enqueued at
	EnqueuedAt *time.Time `json:"enqueuedAt,omitempty"`

	// Name Name of the workflow run
	Name *string `json:"name,omitempty"`

	// Position Workflow run position
	Position *float32 `json:"position,omitempty"`

	// StartedAt Workflow run started at
	StartedAt *time.Time `json:"startedAt,omitempty"`

	// State State of the workflow run
	State map[string]interface{} `json:"state"`

	// Status Workflow run status
	Status            *WorkflowRunStatus  `json:"status,omitempty"`
	WorkflowId        *openapi_types.UUID `json:"workflowId,omitempty"`
	WorkflowVersionId *openapi_types.UUID `json:"workflowVersionId,omitempty"`
}

// WorkflowRunStatus Workflow run status
type WorkflowRunStatus string

// WorkflowRunForResponse A workflow run
type WorkflowRun struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The executor of the workflow
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// EndedAt Workflow run ended at
	EndedAt *time.Time `json:"endedAt,omitempty"`

	// EnqueuedAt Workflow run enqueued at
	EnqueuedAt *time.Time `json:"enqueuedAt,omitempty"`

	// Favorites Favorites linked to the workflow run
	Favorites *[]Favorite `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Name Name of the workflow run
	Name *string `json:"name,omitempty"`

	// Position Workflow run position
	Position *float32 `json:"position,omitempty"`

	// StartedAt Workflow run started at
	StartedAt *time.Time `json:"startedAt,omitempty"`

	// State State of the workflow run
	State *map[string]interface{} `json:"state,omitempty"`

	// Status Workflow run status
	Status *WorkflowRunStatus `json:"status,omitempty"`

	// TimelineActivities Timeline activities linked to the run
	TimelineActivities *[]TimelineActivity `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Workflow Workflow linked to the run.
	Workflow   *Workflow           `json:"workflow,omitempty"`
	WorkflowId *openapi_types.UUID `json:"workflowId,omitempty"`

	// WorkflowVersion Workflow version linked to the run.
	WorkflowVersion   *WorkflowVersion    `json:"workflowVersion,omitempty"`
	WorkflowVersionId *openapi_types.UUID `json:"workflowVersionId,omitempty"`
}

////////////////////////////
// CalendarEvent        //
////////////////////////////

// NewCalendarEvent Calendar events
type NewCalendarEvent struct {
	// ConferenceLink Meet Link
	ConferenceLink *Link `json:"conferenceLink,omitempty"`

	// ConferenceSolution Conference Solution
	ConferenceSolution *string `json:"conferenceSolution,omitempty"`

	// Description Description
	Description *string `json:"description,omitempty"`

	// EndsAt End Date
	EndsAt *time.Time `json:"endsAt,omitempty"`

	// ExternalCreatedAt Creation DateTime
	ExternalCreatedAt *time.Time `json:"externalCreatedAt,omitempty"`

	// ExternalUpdatedAt Update DateTime
	ExternalUpdatedAt *time.Time `json:"externalUpdatedAt,omitempty"`

	// ICalUid iCal UID
	ICalUid *string `json:"iCalUid,omitempty"`

	// IsCanceled Is canceled
	IsCanceled *bool `json:"isCanceled,omitempty"`

	// IsFullDay Is Full Day
	IsFullDay *bool `json:"isFullDay,omitempty"`

	// Location Location
	Location *string `json:"location,omitempty"`

	// StartsAt Start Date
	StartsAt *time.Time `json:"startsAt,omitempty"`

	// Title Title
	Title *string `json:"title,omitempty"`
}

// CalendarEventForResponse Calendar events
type CalendarEvent struct {
	// CalendarChannelEventAssociations Calendar Channel Event Associations
	CalendarChannelEventAssociations *[]CalendarChannelEventAssociation `json:"calendarChannelEventAssociations,omitempty"`

	// CalendarEventParticipants Event Participants
	CalendarEventParticipants *[]CalendarEventParticipant `json:"calendarEventParticipants,omitempty"`

	// ConferenceLink Meet Link
	ConferenceLink *Link `json:"conferenceLink,omitempty"`

	// ConferenceSolution Conference Solution
	ConferenceSolution *string `json:"conferenceSolution,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Description Description
	Description *string `json:"description,omitempty"`

	// EndsAt End Date
	EndsAt *time.Time `json:"endsAt,omitempty"`

	// ExternalCreatedAt Creation DateTime
	ExternalCreatedAt *time.Time `json:"externalCreatedAt,omitempty"`

	// ExternalUpdatedAt Update DateTime
	ExternalUpdatedAt *time.Time `json:"externalUpdatedAt,omitempty"`

	// ICalUid iCal UID
	ICalUid *string `json:"iCalUid,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// IsCanceled Is canceled
	IsCanceled *bool `json:"isCanceled,omitempty"`

	// IsFullDay Is Full Day
	IsFullDay *bool `json:"isFullDay,omitempty"`

	// Location Location
	Location *string `json:"location,omitempty"`

	// StartsAt Start Date
	StartsAt *time.Time `json:"startsAt,omitempty"`

	// Title Title
	Title *string `json:"title,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewCalendarChannelEventAssociation Calendar Channel Event Associations
type NewCalendarChannelEventAssociation struct {
	CalendarChannelId *openapi_types.UUID `json:"calendarChannelId,omitempty"`
	CalendarEventId   *openapi_types.UUID `json:"calendarEventId,omitempty"`

	// EventExternalId Event external ID
	EventExternalId *string `json:"eventExternalId,omitempty"`

	// RecurringEventExternalId Recurring Event ID
	RecurringEventExternalId *string `json:"recurringEventExternalId,omitempty"`
}

// CalendarChannelEventAssociation Calendar Channel Event Associations
type CalendarChannelEventAssociation struct {
	// CalendarChannel Channel ID
	CalendarChannel   *CalendarChannel    `json:"calendarChannel,omitempty"`
	CalendarChannelId *openapi_types.UUID `json:"calendarChannelId,omitempty"`

	// CalendarEvent Event ID
	CalendarEvent   *CalendarEvent      `json:"calendarEvent,omitempty"`
	CalendarEventId *openapi_types.UUID `json:"calendarEventId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// EventExternalId Event external ID
	EventExternalId *string `json:"eventExternalId,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// RecurringEventExternalId Recurring Event ID
	RecurringEventExternalId *string `json:"recurringEventExternalId,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewCalendarChannel Calendar Channels
type NewCalendarChannel struct {
	ConnectedAccountId *openapi_types.UUID `json:"connectedAccountId,omitempty"`

	// ContactAutoCreationPolicy Automatically create records for people you participated with in an event.
	ContactAutoCreationPolicy *CalendarChannelContactAutoCreationPolicy `json:"contactAutoCreationPolicy,omitempty"`

	// Handle Handle
	Handle *string `json:"handle,omitempty"`

	// IsContactAutoCreationEnabled Is Contact Auto Creation Enabled
	IsContactAutoCreationEnabled *bool `json:"isContactAutoCreationEnabled,omitempty"`

	// IsSyncEnabled Is Sync Enabled
	IsSyncEnabled *bool `json:"isSyncEnabled,omitempty"`

	// SyncCursor Sync Cursor. Used for syncing events from the calendar provider
	SyncCursor *string `json:"syncCursor,omitempty"`

	// SyncStage Sync stage
	SyncStage *CalendarChannelSyncStage `json:"syncStage,omitempty"`

	// SyncStageStartedAt Sync stage started at
	SyncStageStartedAt *time.Time `json:"syncStageStartedAt,omitempty"`

	// SyncStatus Sync status
	SyncStatus *CalendarChannelSyncStatus `json:"syncStatus,omitempty"`

	// SyncedAt Last sync date
	SyncedAt *time.Time `json:"syncedAt,omitempty"`

	// ThrottleFailureCount Throttle Failure Count
	ThrottleFailureCount *int `json:"throttleFailureCount,omitempty"`

	// Visibility Visibility
	Visibility *CalendarChannelVisibility `json:"visibility,omitempty"`
}

// CalendarChannelContactAutoCreationPolicy Automatically create records for people you participated with in an event.
type CalendarChannelContactAutoCreationPolicy string

// CalendarChannelSyncStage Sync stage
type CalendarChannelSyncStage string

// CalendarChannelSyncStatus Sync status
type CalendarChannelSyncStatus string

// CalendarChannelVisibility Visibility
type CalendarChannelVisibility string

// CalendarChannel Calendar Channels
type CalendarChannel struct {
	// CalendarChannelEventAssociations Calendar Channel Event Associations
	CalendarChannelEventAssociations *[]CalendarChannelEventAssociation `json:"calendarChannelEventAssociations,omitempty"`

	// ConnectedAccount Connected Account
	ConnectedAccount   *ConnectedAccount   `json:"connectedAccount,omitempty"`
	ConnectedAccountId *openapi_types.UUID `json:"connectedAccountId,omitempty"`

	// ContactAutoCreationPolicy Automatically create records for people you participated with in an event.
	ContactAutoCreationPolicy *CalendarChannelContactAutoCreationPolicy `json:"contactAutoCreationPolicy,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Handle Handle
	Handle *string `json:"handle,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// IsContactAutoCreationEnabled Is Contact Auto Creation Enabled
	IsContactAutoCreationEnabled *bool `json:"isContactAutoCreationEnabled,omitempty"`

	// IsSyncEnabled Is Sync Enabled
	IsSyncEnabled *bool `json:"isSyncEnabled,omitempty"`

	// SyncCursor Sync Cursor. Used for syncing events from the calendar provider
	SyncCursor *string `json:"syncCursor,omitempty"`

	// SyncStage Sync stage
	SyncStage *CalendarChannelSyncStage `json:"syncStage,omitempty"`

	// SyncStageStartedAt Sync stage started at
	SyncStageStartedAt *time.Time `json:"syncStageStartedAt,omitempty"`

	// SyncStatus Sync status
	SyncStatus *CalendarChannelSyncStatus `json:"syncStatus,omitempty"`

	// SyncedAt Last sync date
	SyncedAt *time.Time `json:"syncedAt,omitempty"`

	// ThrottleFailureCount Throttle Failure Count
	ThrottleFailureCount *int `json:"throttleFailureCount,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Visibility Visibility
	Visibility *CalendarChannelVisibility `json:"visibility,omitempty"`
}

////////////////////////////
// ConnectedAccount                   //
////////////////////////////

// NewConnectedAccount A connected account
type NewConnectedAccount struct {
	// AccessToken Messaging provider access token
	AccessToken    *string             `json:"accessToken,omitempty"`
	AccountOwnerId *openapi_types.UUID `json:"accountOwnerId,omitempty"`

	// AuthFailedAt Auth failed at
	AuthFailedAt *time.Time `json:"authFailedAt,omitempty"`

	// ConnectionParameters JSON object containing custom connection parameters
	ConnectionParameters *map[string]interface{} `json:"connectionParameters,omitempty"`

	// Handle The account handle (email, username, phone number, etc.)
	Handle *string `json:"handle,omitempty"`

	// HandleAliases Handle Aliases
	HandleAliases *string `json:"handleAliases,omitempty"`

	// LastCredentialsRefreshedAt Last credentials refreshed at
	LastCredentialsRefreshedAt *time.Time `json:"lastCredentialsRefreshedAt,omitempty"`

	// LastSyncHistoryId Last sync history ID
	LastSyncHistoryId *string `json:"lastSyncHistoryId,omitempty"`

	// Provider The account provider
	Provider *string `json:"provider,omitempty"`

	// RefreshToken Messaging provider refresh token
	RefreshToken *string `json:"refreshToken,omitempty"`

	// Scopes Scopes
	Scopes *[]string `json:"scopes,omitempty"`
}

// ConnectedAccount A connected account
type ConnectedAccount struct {
	// AccessToken Messaging provider access token
	AccessToken *string `json:"accessToken,omitempty"`

	// AccountOwner Account Owner
	AccountOwner   *WorkspaceMember    `json:"accountOwner,omitempty"`
	AccountOwnerId *openapi_types.UUID `json:"accountOwnerId,omitempty"`

	// AuthFailedAt Auth failed at
	AuthFailedAt *time.Time `json:"authFailedAt,omitempty"`

	// CalendarChannels Calendar Channels
	CalendarChannels *[]CalendarChannel `json:"calendarChannels,omitempty"`

	// ConnectionParameters JSON object containing custom connection parameters
	ConnectionParameters *map[string]interface{} `json:"connectionParameters,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Handle The account handle (email, username, phone number, etc.)
	Handle *string `json:"handle,omitempty"`

	// HandleAliases Handle Aliases
	HandleAliases *string `json:"handleAliases,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// LastCredentialsRefreshedAt Last credentials refreshed at
	LastCredentialsRefreshedAt *time.Time `json:"lastCredentialsRefreshedAt,omitempty"`

	// LastSyncHistoryId Last sync history ID
	LastSyncHistoryId *string `json:"lastSyncHistoryId,omitempty"`

	// MessageChannels Message Channels
	MessageChannels *[]MessageChannel `json:"messageChannels,omitempty"`

	// Provider The account provider
	Provider *string `json:"provider,omitempty"`

	// RefreshToken Messaging provider refresh token
	RefreshToken *string `json:"refreshToken,omitempty"`

	// Scopes Scopes
	Scopes *[]string `json:"scopes,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

////////////////////////////
// Blocklist                   //
////////////////////////////

// NewBlocklist Blocklist
type NewBlocklist struct {
	// Handle Handle
	Handle            *string             `json:"handle,omitempty"`
	WorkspaceMemberId *openapi_types.UUID `json:"workspaceMemberId,omitempty"`
}

// Blocklist Blocklist
type Blocklist struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Handle Handle
	Handle *string `json:"handle,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// WorkspaceMember WorkspaceMember
	WorkspaceMember   *WorkspaceMember    `json:"workspaceMember,omitempty"`
	WorkspaceMemberId *openapi_types.UUID `json:"workspaceMemberId,omitempty"`
}

////////////////////////////
// Message                   //
////////////////////////////

// NewMessage A message sent or received through a messaging channel (email, chat, etc.)
type NewMessage struct {
	// HeaderMessageId Message id from the message header
	HeaderMessageId *string             `json:"headerMessageId,omitempty"`
	MessageThreadId *openapi_types.UUID `json:"messageThreadId,omitempty"`

	// ReceivedAt The date the message was received
	ReceivedAt *time.Time `json:"receivedAt,omitempty"`

	// Subject Subject
	Subject *string `json:"subject,omitempty"`

	// Text Text
	Text *string `json:"text,omitempty"`
}

// Message A message sent or received through a messaging channel (email, chat, etc.)
type Message struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// HeaderMessageId Message id from the message header
	HeaderMessageId *string `json:"headerMessageId,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// MessageChannelMessageAssociations Messages from the channel.
	MessageChannelMessageAssociations *[]MessageChannelMessageAssociation `json:"messageChannelMessageAssociations,omitempty"`

	// MessageParticipants Message Participants
	MessageParticipants *[]MessageParticipant `json:"messageParticipants,omitempty"`

	// MessageThread Message Thread Id
	MessageThread   *MessageThread      `json:"messageThread,omitempty"`
	MessageThreadId *openapi_types.UUID `json:"messageThreadId,omitempty"`

	// ReceivedAt The date the message was received
	ReceivedAt *time.Time `json:"receivedAt,omitempty"`

	// Subject Subject
	Subject *string `json:"subject,omitempty"`

	// Text Text
	Text *string `json:"text,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

////////////////////////////
// MessageChannel         //
////////////////////////////

// NewMessageChannel Message Channels
type NewMessageChannel struct {
	ConnectedAccountId *openapi_types.UUID `json:"connectedAccountId,omitempty"`

	// ContactAutoCreationPolicy Automatically create People records when receiving or sending emails
	ContactAutoCreationPolicy *MessageChannelContactAutoCreationPolicy `json:"contactAutoCreationPolicy,omitempty"`

	// ExcludeGroupEmails Exclude group emails
	ExcludeGroupEmails *bool `json:"excludeGroupEmails,omitempty"`

	// ExcludeNonProfessionalEmails Exclude non professional emails
	ExcludeNonProfessionalEmails *bool `json:"excludeNonProfessionalEmails,omitempty"`

	// Handle Handle
	Handle *string `json:"handle,omitempty"`

	// IsContactAutoCreationEnabled Is Contact Auto Creation Enabled
	IsContactAutoCreationEnabled *bool `json:"isContactAutoCreationEnabled,omitempty"`

	// IsSyncEnabled Is Sync Enabled
	IsSyncEnabled *bool `json:"isSyncEnabled,omitempty"`

	// MessageFolderImportPolicy Message folder import policy
	MessageFolderImportPolicy *MessageChannelMessageFolderImportPolicy `json:"messageFolderImportPolicy,omitempty"`

	// PendingGroupEmailsAction Pending action for group emails
	PendingGroupEmailsAction *MessageChannelPendingGroupEmailsAction `json:"pendingGroupEmailsAction,omitempty"`

	// SyncCursor Last sync cursor
	SyncCursor *string `json:"syncCursor,omitempty"`

	// SyncStage Sync stage
	SyncStage *MessageChannelSyncStage `json:"syncStage,omitempty"`

	// SyncStageStartedAt Sync stage started at
	SyncStageStartedAt *time.Time `json:"syncStageStartedAt,omitempty"`

	// SyncStatus Sync status
	SyncStatus *MessageChannelSyncStatus `json:"syncStatus,omitempty"`

	// SyncedAt Last sync date
	SyncedAt *time.Time `json:"syncedAt,omitempty"`

	// ThrottleFailureCount Throttle Failure Count
	ThrottleFailureCount *int `json:"throttleFailureCount,omitempty"`

	// Type Channel Type
	Type *MessageChannelType `json:"type,omitempty"`

	// Visibility Visibility
	Visibility *MessageChannelVisibility `json:"visibility,omitempty"`
}

// MessageChannelContactAutoCreationPolicy Automatically create People records when receiving or sending emails
type MessageChannelContactAutoCreationPolicy string

// MessageChannelMessageFolderImportPolicy Message folder import policy
type MessageChannelMessageFolderImportPolicy string

// MessageChannelPendingGroupEmailsAction Pending action for group emails
type MessageChannelPendingGroupEmailsAction string

// MessageChannelSyncStage Sync stage
type MessageChannelSyncStage string

// MessageChannelSyncStatus Sync status
type MessageChannelSyncStatus string

// MessageChannelType Channel Type
type MessageChannelType string

// MessageChannelVisibility Visibility
type MessageChannelVisibility string

// MessageChannel Message Channels
type MessageChannel struct {
	// ConnectedAccount Connected Account
	ConnectedAccount   *ConnectedAccount   `json:"connectedAccount,omitempty"`
	ConnectedAccountId *openapi_types.UUID `json:"connectedAccountId,omitempty"`

	// ContactAutoCreationPolicy Automatically create People records when receiving or sending emails
	ContactAutoCreationPolicy *MessageChannelContactAutoCreationPolicy `json:"contactAutoCreationPolicy,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// ExcludeGroupEmails Exclude group emails
	ExcludeGroupEmails *bool `json:"excludeGroupEmails,omitempty"`

	// ExcludeNonProfessionalEmails Exclude non professional emails
	ExcludeNonProfessionalEmails *bool `json:"excludeNonProfessionalEmails,omitempty"`

	// Handle Handle
	Handle *string `json:"handle,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// IsContactAutoCreationEnabled Is Contact Auto Creation Enabled
	IsContactAutoCreationEnabled *bool `json:"isContactAutoCreationEnabled,omitempty"`

	// IsSyncEnabled Is Sync Enabled
	IsSyncEnabled *bool `json:"isSyncEnabled,omitempty"`

	// MessageChannelMessageAssociations Messages from the channel.
	MessageChannelMessageAssociations *[]MessageChannelMessageAssociation `json:"messageChannelMessageAssociations,omitempty"`

	// MessageFolderImportPolicy Message folder import policy
	MessageFolderImportPolicy *MessageChannelMessageFolderImportPolicy `json:"messageFolderImportPolicy,omitempty"`

	// MessageFolders Message Folders
	MessageFolders *[]MessageFolder `json:"messageFolders,omitempty"`

	// PendingGroupEmailsAction Pending action for group emails
	PendingGroupEmailsAction *MessageChannelPendingGroupEmailsAction `json:"pendingGroupEmailsAction,omitempty"`

	// SyncCursor Last sync cursor
	SyncCursor *string `json:"syncCursor,omitempty"`

	// SyncStage Sync stage
	SyncStage *MessageChannelSyncStage `json:"syncStage,omitempty"`

	// SyncStageStartedAt Sync stage started at
	SyncStageStartedAt *time.Time `json:"syncStageStartedAt,omitempty"`

	// SyncStatus Sync status
	SyncStatus *MessageChannelSyncStatus `json:"syncStatus,omitempty"`

	// SyncedAt Last sync date
	SyncedAt *time.Time `json:"syncedAt,omitempty"`

	// ThrottleFailureCount Throttle Failure Count
	ThrottleFailureCount *int `json:"throttleFailureCount,omitempty"`

	// Type Channel Type
	Type *MessageChannelType `json:"type,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Visibility Visibility
	Visibility *MessageChannelVisibility `json:"visibility,omitempty"`
}

// NewMessageChannelMessageAssociation Message Synced with a Message Channel
type NewMessageChannelMessageAssociation struct {
	// Direction Message Direction
	Direction        *MessageChannelMessageAssociationDirection `json:"direction,omitempty"`
	MessageChannelId *openapi_types.UUID                        `json:"messageChannelId,omitempty"`

	// MessageExternalId Message id from the messaging provider
	MessageExternalId *string             `json:"messageExternalId,omitempty"`
	MessageId         *openapi_types.UUID `json:"messageId,omitempty"`

	// MessageThreadExternalId Thread id from the messaging provider
	MessageThreadExternalId *string `json:"messageThreadExternalId,omitempty"`
}

// MessageChannelMessageAssociationDirection Message Direction
type MessageChannelMessageAssociationDirection string

// MessageChannelMessageAssociation Message Synced with a Message Channel
type MessageChannelMessageAssociation struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Direction Message Direction
	Direction *MessageChannelMessageAssociationDirection `json:"direction,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Message Message Id
	Message *Message `json:"message,omitempty"`

	// MessageChannel Message Channel Id
	MessageChannel   *MessageChannel     `json:"messageChannel,omitempty"`
	MessageChannelId *openapi_types.UUID `json:"messageChannelId,omitempty"`

	// MessageExternalId Message id from the messaging provider
	MessageExternalId *string             `json:"messageExternalId,omitempty"`
	MessageId         *openapi_types.UUID `json:"messageId,omitempty"`

	// MessageThreadExternalId Thread id from the messaging provider
	MessageThreadExternalId *string `json:"messageThreadExternalId,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewMessageFolder Folder for Message Channel
type NewMessageFolder struct {
	// ExternalId External ID
	ExternalId *string `json:"externalId,omitempty"`

	// IsSentFolder Is Sent Folder
	IsSentFolder *bool `json:"isSentFolder,omitempty"`

	// IsSynced Is Synced
	IsSynced         *bool               `json:"isSynced,omitempty"`
	MessageChannelId *openapi_types.UUID `json:"messageChannelId,omitempty"`

	// Name Folder name
	Name *string `json:"name,omitempty"`

	// ParentFolderId Parent Folder ID
	ParentFolderId *string `json:"parentFolderId,omitempty"`

	// PendingSyncAction Pending action for folder sync
	PendingSyncAction *MessageFolderPendingSyncAction `json:"pendingSyncAction,omitempty"`

	// SyncCursor Sync Cursor
	SyncCursor *string `json:"syncCursor,omitempty"`
}

// MessageFolderPendingSyncAction Pending action for folder sync
type MessageFolderPendingSyncAction string

// MessageFolder Folder for Message Channel
type MessageFolder struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// ExternalId External ID
	ExternalId *string `json:"externalId,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// IsSentFolder Is Sent Folder
	IsSentFolder *bool `json:"isSentFolder,omitempty"`

	// IsSynced Is Synced
	IsSynced *bool `json:"isSynced,omitempty"`

	// MessageChannel Message Channel
	MessageChannel   *MessageChannel     `json:"messageChannel,omitempty"`
	MessageChannelId *openapi_types.UUID `json:"messageChannelId,omitempty"`

	// Name Folder name
	Name *string `json:"name,omitempty"`

	// ParentFolderId Parent Folder ID
	ParentFolderId *string `json:"parentFolderId,omitempty"`

	// PendingSyncAction Pending action for folder sync
	PendingSyncAction *MessageFolderPendingSyncAction `json:"pendingSyncAction,omitempty"`

	// SyncCursor Sync Cursor
	SyncCursor *string `json:"syncCursor,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewMessageThread A group of related messages (e.g. email thread, chat thread)
type NewMessageThread = map[string]interface{}

// MessageThread A group of related messages (e.g. email thread, chat thread)
type MessageThread struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Messages Messages from the thread.
	Messages *[]Message `json:"messages,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

////////////////////////////
// Vars                   //
////////////////////////////

// Defines values for PersonCreatedBySource.
const (
	CreatedBySourceAGENT    CreatedBySource = "AGENT"
	CreatedBySourceAPI      CreatedBySource = "API"
	CreatedBySourceCALENDAR CreatedBySource = "CALENDAR"
	CreatedBySourceEMAIL    CreatedBySource = "EMAIL"
	CreatedBySourceIMPORT   CreatedBySource = "IMPORT"
	CreatedBySourceMANUAL   CreatedBySource = "MANUAL"
	CreatedBySourceSYSTEM   CreatedBySource = "SYSTEM"
	CreatedBySourceWEBHOOK  CreatedBySource = "WEBHOOK"
	CreatedBySourceWORKFLOW CreatedBySource = "WORKFLOW"
)

// Defines values for AttachmentFileCategory.
const (
	AttachmentFileCategoryARCHIVE      AttachmentFileCategory = "ARCHIVE"
	AttachmentFileCategoryAUDIO        AttachmentFileCategory = "AUDIO"
	AttachmentFileCategoryIMAGE        AttachmentFileCategory = "IMAGE"
	AttachmentFileCategoryOTHER        AttachmentFileCategory = "OTHER"
	AttachmentFileCategoryPRESENTATION AttachmentFileCategory = "PRESENTATION"
	AttachmentFileCategorySPREADSHEET  AttachmentFileCategory = "SPREADSHEET"
	AttachmentFileCategoryTEXTDOCUMENT AttachmentFileCategory = "TEXT_DOCUMENT"
	AttachmentFileCategoryVIDEO        AttachmentFileCategory = "VIDEO"
)

// Defines values for CalendarEventParticipantResponseStatus.
const (
	CalendarEventParticipantResponseStatusACCEPTED    CalendarEventParticipantResponseStatus = "ACCEPTED"
	CalendarEventParticipantResponseStatusDECLINED    CalendarEventParticipantResponseStatus = "DECLINED"
	CalendarEventParticipantResponseStatusNEEDSACTION CalendarEventParticipantResponseStatus = "NEEDS_ACTION"
	CalendarEventParticipantResponseStatusTENTATIVE   CalendarEventParticipantResponseStatus = "TENTATIVE"
)

// Defines values for WorkspaceMemberDateFormat.
const (
	WorkspaceMemberDateFormatDAYFIRST   WorkspaceMemberDateFormat = "DAY_FIRST"
	WorkspaceMemberDateFormatMONTHFIRST WorkspaceMemberDateFormat = "MONTH_FIRST"
	WorkspaceMemberDateFormatSYSTEM     WorkspaceMemberDateFormat = "SYSTEM"
	WorkspaceMemberDateFormatYEARFIRST  WorkspaceMemberDateFormat = "YEAR_FIRST"
)

// Defines values for WorkspaceMemberNumberFormat.
const (
	WorkspaceMemberNumberFormatAPOSTROPHEANDDOT WorkspaceMemberNumberFormat = "APOSTROPHE_AND_DOT"
	WorkspaceMemberNumberFormatCOMMASANDDOT     WorkspaceMemberNumberFormat = "COMMAS_AND_DOT"
	WorkspaceMemberNumberFormatDOTSANDCOMMA     WorkspaceMemberNumberFormat = "DOTS_AND_COMMA"
	WorkspaceMemberNumberFormatSPACESANDCOMMA   WorkspaceMemberNumberFormat = "SPACES_AND_COMMA"
	WorkspaceMemberNumberFormatSYSTEM           WorkspaceMemberNumberFormat = "SYSTEM"
)

// Defines values for WorkspaceMemberTimeFormat.
const (
	WorkspaceMemberTimeFormatHOUR12 WorkspaceMemberTimeFormat = "HOUR_12"
	WorkspaceMemberTimeFormatHOUR24 WorkspaceMemberTimeFormat = "HOUR_24"
	WorkspaceMemberTimeFormatSYSTEM WorkspaceMemberTimeFormat = "SYSTEM"
)

// Defines values for MessageParticipantRole.
const (
	MessageParticipantRoleBCC  MessageParticipantRole = "BCC"
	MessageParticipantRoleCC   MessageParticipantRole = "CC"
	MessageParticipantRoleFROM MessageParticipantRole = "FROM"
	MessageParticipantRoleTO   MessageParticipantRole = "TO"
)

// Defines values for OpportunityLostReason.
const (
	OpportunityLostReasonCOMPETITION     OpportunityLostReason = "COMPETITION"
	OpportunityLostReasonMISSINGFEATURES OpportunityLostReason = "MISSING_FEATURES"
	OpportunityLostReasonNOTUSED         OpportunityLostReason = "NOT_USED"
	OpportunityLostReasonOTHER           OpportunityLostReason = "OTHER"
	OpportunityLostReasonPRICE           OpportunityLostReason = "PRICE"
)

// Defines values for OpportunityStage.
const (
	OpportunityStageACTIVE                           OpportunityStage = "ACTIVE"
	OpportunityStageCANCELED                         OpportunityStage = "CANCELED"
	OpportunityStageCLOSERCALLVEREINBART             OpportunityStage = "CLOSER_CALL_VEREINBART"
	OpportunityStageDEALLOST                         OpportunityStage = "DEAL_LOST"
	OpportunityStageDETAILSCALLVEREINBART            OpportunityStage = "DETAILS_CALL_VEREINBART"
	OpportunityStageENTSCHEIDUNGSTRAGERHATZUGESTIMMT OpportunityStage = "ENTSCHEIDUNGSTRAGER_HAT_ZUGESTIMMT"
	OpportunityStageSETTERCALLVEREINBART             OpportunityStage = "SETTER_CALL_VEREINBART"
	OpportunityStageTRIALPERIOD                      OpportunityStage = "TRIAL_PERIOD"
	OpportunityStageVERTRAGEVERSANT                  OpportunityStage = "VERTRAGE_VERSANT"
)

// Defines values for TaskStatus.
const (
	TaskStatusDONE       TaskStatus = "DONE"
	TaskStatusINPROGRESS TaskStatus = "IN_PROGRESS"
	TaskStatusTODO       TaskStatus = "TODO"
)

// Defines values for WorkflowStatuses.
const (
	WorkflowStatusesACTIVE      WorkflowStatuses = "ACTIVE"
	WorkflowStatusesDEACTIVATED WorkflowStatuses = "DEACTIVATED"
	WorkflowStatusesDRAFT       WorkflowStatuses = "DRAFT"
)

// Defines values for WorkflowAutomatedTriggerType.
const (
	WorkflowAutomatedTriggerTypeCRON          WorkflowAutomatedTriggerType = "CRON"
	WorkflowAutomatedTriggerTypeDATABASEEVENT WorkflowAutomatedTriggerType = "DATABASE_EVENT"
)

// Defines values for WorkflowVersionStatus.
const (
	WorkflowVersionStatusACTIVE      WorkflowVersionStatus = "ACTIVE"
	WorkflowVersionStatusARCHIVED    WorkflowVersionStatus = "ARCHIVED"
	WorkflowVersionStatusDEACTIVATED WorkflowVersionStatus = "DEACTIVATED"
	WorkflowVersionStatusDRAFT       WorkflowVersionStatus = "DRAFT"
)

// Defines values for WorkflowRunStatus.
const (
	WorkflowRunStatusCOMPLETED  WorkflowRunStatus = "COMPLETED"
	WorkflowRunStatusENQUEUED   WorkflowRunStatus = "ENQUEUED"
	WorkflowRunStatusFAILED     WorkflowRunStatus = "FAILED"
	WorkflowRunStatusNOTSTARTED WorkflowRunStatus = "NOT_STARTED"
	WorkflowRunStatusRUNNING    WorkflowRunStatus = "RUNNING"
	WorkflowRunStatusSTOPPED    WorkflowRunStatus = "STOPPED"
	WorkflowRunStatusSTOPPING   WorkflowRunStatus = "STOPPING"
)

// Defines values for CalendarChannelContactAutoCreationPolicy.
const (
	CalendarChannelContactAutoCreationPolicyASORGANIZER               CalendarChannelContactAutoCreationPolicy = "AS_ORGANIZER"
	CalendarChannelContactAutoCreationPolicyASPARTICIPANT             CalendarChannelContactAutoCreationPolicy = "AS_PARTICIPANT"
	CalendarChannelContactAutoCreationPolicyASPARTICIPANTANDORGANIZER CalendarChannelContactAutoCreationPolicy = "AS_PARTICIPANT_AND_ORGANIZER"
	CalendarChannelContactAutoCreationPolicyNONE                      CalendarChannelContactAutoCreationPolicy = "NONE"
)

// Defines values for CalendarChannelSyncStage.
const (
	CalendarChannelSyncStageCALENDAREVENTLISTFETCHONGOING   CalendarChannelSyncStage = "CALENDAR_EVENT_LIST_FETCH_ONGOING"
	CalendarChannelSyncStageCALENDAREVENTLISTFETCHPENDING   CalendarChannelSyncStage = "CALENDAR_EVENT_LIST_FETCH_PENDING"
	CalendarChannelSyncStageCALENDAREVENTLISTFETCHSCHEDULED CalendarChannelSyncStage = "CALENDAR_EVENT_LIST_FETCH_SCHEDULED"
	CalendarChannelSyncStageCALENDAREVENTSIMPORTONGOING     CalendarChannelSyncStage = "CALENDAR_EVENTS_IMPORT_ONGOING"
	CalendarChannelSyncStageCALENDAREVENTSIMPORTPENDING     CalendarChannelSyncStage = "CALENDAR_EVENTS_IMPORT_PENDING"
	CalendarChannelSyncStageCALENDAREVENTSIMPORTSCHEDULED   CalendarChannelSyncStage = "CALENDAR_EVENTS_IMPORT_SCHEDULED"
	CalendarChannelSyncStageFAILED                          CalendarChannelSyncStage = "FAILED"
	CalendarChannelSyncStagePENDINGCONFIGURATION            CalendarChannelSyncStage = "PENDING_CONFIGURATION"
)

// Defines values for CalendarChannelSyncStatus.
const (
	CalendarChannelSyncStatusACTIVE                        CalendarChannelSyncStatus = "ACTIVE"
	CalendarChannelSyncStatusFAILEDINSUFFICIENTPERMISSIONS CalendarChannelSyncStatus = "FAILED_INSUFFICIENT_PERMISSIONS"
	CalendarChannelSyncStatusFAILEDUNKNOWN                 CalendarChannelSyncStatus = "FAILED_UNKNOWN"
	CalendarChannelSyncStatusNOTSYNCED                     CalendarChannelSyncStatus = "NOT_SYNCED"
	CalendarChannelSyncStatusONGOING                       CalendarChannelSyncStatus = "ONGOING"
)

// Defines values for CalendarChannelVisibility.
const (
	CalendarChannelVisibilityMETADATA        CalendarChannelVisibility = "METADATA"
	CalendarChannelVisibilitySHAREEVERYTHING CalendarChannelVisibility = "SHARE_EVERYTHING"
)

// Defines values for MessageChannelContactAutoCreationPolicy.
const (
	MessageChannelContactAutoCreationPolicyNONE            MessageChannelContactAutoCreationPolicy = "NONE"
	MessageChannelContactAutoCreationPolicySENT            MessageChannelContactAutoCreationPolicy = "SENT"
	MessageChannelContactAutoCreationPolicySENTANDRECEIVED MessageChannelContactAutoCreationPolicy = "SENT_AND_RECEIVED"
)

// Defines values for MessageChannelMessageFolderImportPolicy.
const (
	MessageChannelMessageFolderImportPolicyALLFOLDERS      MessageChannelMessageFolderImportPolicy = "ALL_FOLDERS"
	MessageChannelMessageFolderImportPolicySELECTEDFOLDERS MessageChannelMessageFolderImportPolicy = "SELECTED_FOLDERS"
)

// Defines values for MessageChannelPendingGroupEmailsAction.
const (
	MessageChannelPendingGroupEmailsActionGROUPEMAILSDELETION MessageChannelPendingGroupEmailsAction = "GROUP_EMAILS_DELETION"
	MessageChannelPendingGroupEmailsActionGROUPEMAILSIMPORT   MessageChannelPendingGroupEmailsAction = "GROUP_EMAILS_IMPORT"
	MessageChannelPendingGroupEmailsActionNONE                MessageChannelPendingGroupEmailsAction = "NONE"
)

// Defines values for MessageChannelSyncStage.
const (
	MessageChannelSyncStageFAILED                    MessageChannelSyncStage = "FAILED"
	MessageChannelSyncStageMESSAGELISTFETCHONGOING   MessageChannelSyncStage = "MESSAGE_LIST_FETCH_ONGOING"
	MessageChannelSyncStageMESSAGELISTFETCHPENDING   MessageChannelSyncStage = "MESSAGE_LIST_FETCH_PENDING"
	MessageChannelSyncStageMESSAGELISTFETCHSCHEDULED MessageChannelSyncStage = "MESSAGE_LIST_FETCH_SCHEDULED"
	MessageChannelSyncStageMESSAGESIMPORTONGOING     MessageChannelSyncStage = "MESSAGES_IMPORT_ONGOING"
	MessageChannelSyncStageMESSAGESIMPORTPENDING     MessageChannelSyncStage = "MESSAGES_IMPORT_PENDING"
	MessageChannelSyncStageMESSAGESIMPORTSCHEDULED   MessageChannelSyncStage = "MESSAGES_IMPORT_SCHEDULED"
	MessageChannelSyncStagePENDINGCONFIGURATION      MessageChannelSyncStage = "PENDING_CONFIGURATION"
)

// Defines values for MessageChannelSyncStatus.
const (
	MessageChannelSyncStatusACTIVE                        MessageChannelSyncStatus = "ACTIVE"
	MessageChannelSyncStatusFAILEDINSUFFICIENTPERMISSIONS MessageChannelSyncStatus = "FAILED_INSUFFICIENT_PERMISSIONS"
	MessageChannelSyncStatusFAILEDUNKNOWN                 MessageChannelSyncStatus = "FAILED_UNKNOWN"
	MessageChannelSyncStatusNOTSYNCED                     MessageChannelSyncStatus = "NOT_SYNCED"
	MessageChannelSyncStatusONGOING                       MessageChannelSyncStatus = "ONGOING"
)

// Defines values for MessageChannelType.
const (
	MessageChannelTypeEMAIL MessageChannelType = "EMAIL"
	MessageChannelTypeSMS   MessageChannelType = "SMS"
)

// Defines values for MessageChannelVisibility.
const (
	MessageChannelVisibilityMETADATA        MessageChannelVisibility = "METADATA"
	MessageChannelVisibilitySHAREEVERYTHING MessageChannelVisibility = "SHARE_EVERYTHING"
	MessageChannelVisibilitySUBJECT         MessageChannelVisibility = "SUBJECT"
)

// Defines values for MessageChannelMessageAssociationDirection.
const (
	MessageChannelMessageAssociationDirectionINCOMING MessageChannelMessageAssociationDirection = "INCOMING"
	MessageChannelMessageAssociationDirectionOUTGOING MessageChannelMessageAssociationDirection = "OUTGOING"
)

// Defines values for MessageFolderPendingSyncAction.
const (
	MessageFolderPendingSyncActionFOLDERDELETION MessageFolderPendingSyncAction = "FOLDER_DELETION"
	MessageFolderPendingSyncActionNONE           MessageFolderPendingSyncAction = "NONE"
)

////////////////////////////
// General                //
////////////////////////////

type CreatedBySource string

type NewCreatedBy struct {
	Source *CreatedBySource `json:"source,omitempty"`
}

type CreatedBy struct {
	Name              *string             `json:"name,omitempty"`
	Source            *CreatedBySource    `json:"source,omitempty"`
	WorkspaceMemberId *openapi_types.UUID `json:"workspaceMemberId,omitempty"`
}

type Link struct {
	PrimaryLinkLabel *string          `json:"primaryLinkLabel,omitempty"`
	PrimaryLinkUrl   *string          `json:"primaryLinkUrl,omitempty"`
	SecondaryLinks   *[]SecondaryLink `json:"secondaryLinks,omitempty"`
}

type SecondaryLink struct {
	Label *string `json:"label,omitempty"`
	Url   *string `json:"url,omitempty"`
}
