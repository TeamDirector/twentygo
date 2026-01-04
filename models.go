package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for AttachmentCreatedBySource.
const (
	AttachmentCreatedBySourceAGENT    AttachmentCreatedBySource = "AGENT"
	AttachmentCreatedBySourceAPI      AttachmentCreatedBySource = "API"
	AttachmentCreatedBySourceCALENDAR AttachmentCreatedBySource = "CALENDAR"
	AttachmentCreatedBySourceEMAIL    AttachmentCreatedBySource = "EMAIL"
	AttachmentCreatedBySourceIMPORT   AttachmentCreatedBySource = "IMPORT"
	AttachmentCreatedBySourceMANUAL   AttachmentCreatedBySource = "MANUAL"
	AttachmentCreatedBySourceSYSTEM   AttachmentCreatedBySource = "SYSTEM"
	AttachmentCreatedBySourceWEBHOOK  AttachmentCreatedBySource = "WEBHOOK"
	AttachmentCreatedBySourceWORKFLOW AttachmentCreatedBySource = "WORKFLOW"
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

// Defines values for AttachmentForResponseCreatedBySource.
const (
	AttachmentForResponseCreatedBySourceAGENT    AttachmentForResponseCreatedBySource = "AGENT"
	AttachmentForResponseCreatedBySourceAPI      AttachmentForResponseCreatedBySource = "API"
	AttachmentForResponseCreatedBySourceCALENDAR AttachmentForResponseCreatedBySource = "CALENDAR"
	AttachmentForResponseCreatedBySourceEMAIL    AttachmentForResponseCreatedBySource = "EMAIL"
	AttachmentForResponseCreatedBySourceIMPORT   AttachmentForResponseCreatedBySource = "IMPORT"
	AttachmentForResponseCreatedBySourceMANUAL   AttachmentForResponseCreatedBySource = "MANUAL"
	AttachmentForResponseCreatedBySourceSYSTEM   AttachmentForResponseCreatedBySource = "SYSTEM"
	AttachmentForResponseCreatedBySourceWEBHOOK  AttachmentForResponseCreatedBySource = "WEBHOOK"
	AttachmentForResponseCreatedBySourceWORKFLOW AttachmentForResponseCreatedBySource = "WORKFLOW"
)

// Defines values for AttachmentForResponseFileCategory.
const (
	AttachmentForResponseFileCategoryARCHIVE      AttachmentForResponseFileCategory = "ARCHIVE"
	AttachmentForResponseFileCategoryAUDIO        AttachmentForResponseFileCategory = "AUDIO"
	AttachmentForResponseFileCategoryIMAGE        AttachmentForResponseFileCategory = "IMAGE"
	AttachmentForResponseFileCategoryOTHER        AttachmentForResponseFileCategory = "OTHER"
	AttachmentForResponseFileCategoryPRESENTATION AttachmentForResponseFileCategory = "PRESENTATION"
	AttachmentForResponseFileCategorySPREADSHEET  AttachmentForResponseFileCategory = "SPREADSHEET"
	AttachmentForResponseFileCategoryTEXTDOCUMENT AttachmentForResponseFileCategory = "TEXT_DOCUMENT"
	AttachmentForResponseFileCategoryVIDEO        AttachmentForResponseFileCategory = "VIDEO"
)

// Defines values for AttachmentForUpdateCreatedBySource.
const (
	AttachmentForUpdateCreatedBySourceAGENT    AttachmentForUpdateCreatedBySource = "AGENT"
	AttachmentForUpdateCreatedBySourceAPI      AttachmentForUpdateCreatedBySource = "API"
	AttachmentForUpdateCreatedBySourceCALENDAR AttachmentForUpdateCreatedBySource = "CALENDAR"
	AttachmentForUpdateCreatedBySourceEMAIL    AttachmentForUpdateCreatedBySource = "EMAIL"
	AttachmentForUpdateCreatedBySourceIMPORT   AttachmentForUpdateCreatedBySource = "IMPORT"
	AttachmentForUpdateCreatedBySourceMANUAL   AttachmentForUpdateCreatedBySource = "MANUAL"
	AttachmentForUpdateCreatedBySourceSYSTEM   AttachmentForUpdateCreatedBySource = "SYSTEM"
	AttachmentForUpdateCreatedBySourceWEBHOOK  AttachmentForUpdateCreatedBySource = "WEBHOOK"
	AttachmentForUpdateCreatedBySourceWORKFLOW AttachmentForUpdateCreatedBySource = "WORKFLOW"
)

// Defines values for AttachmentForUpdateFileCategory.
const (
	AttachmentForUpdateFileCategoryARCHIVE      AttachmentForUpdateFileCategory = "ARCHIVE"
	AttachmentForUpdateFileCategoryAUDIO        AttachmentForUpdateFileCategory = "AUDIO"
	AttachmentForUpdateFileCategoryIMAGE        AttachmentForUpdateFileCategory = "IMAGE"
	AttachmentForUpdateFileCategoryOTHER        AttachmentForUpdateFileCategory = "OTHER"
	AttachmentForUpdateFileCategoryPRESENTATION AttachmentForUpdateFileCategory = "PRESENTATION"
	AttachmentForUpdateFileCategorySPREADSHEET  AttachmentForUpdateFileCategory = "SPREADSHEET"
	AttachmentForUpdateFileCategoryTEXTDOCUMENT AttachmentForUpdateFileCategory = "TEXT_DOCUMENT"
	AttachmentForUpdateFileCategoryVIDEO        AttachmentForUpdateFileCategory = "VIDEO"
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

// Defines values for CalendarChannelForResponseContactAutoCreationPolicy.
const (
	CalendarChannelForResponseContactAutoCreationPolicyASORGANIZER               CalendarChannelForResponseContactAutoCreationPolicy = "AS_ORGANIZER"
	CalendarChannelForResponseContactAutoCreationPolicyASPARTICIPANT             CalendarChannelForResponseContactAutoCreationPolicy = "AS_PARTICIPANT"
	CalendarChannelForResponseContactAutoCreationPolicyASPARTICIPANTANDORGANIZER CalendarChannelForResponseContactAutoCreationPolicy = "AS_PARTICIPANT_AND_ORGANIZER"
	CalendarChannelForResponseContactAutoCreationPolicyNONE                      CalendarChannelForResponseContactAutoCreationPolicy = "NONE"
)

// Defines values for CalendarChannelForResponseSyncStage.
const (
	CalendarChannelForResponseSyncStageCALENDAREVENTLISTFETCHONGOING   CalendarChannelForResponseSyncStage = "CALENDAR_EVENT_LIST_FETCH_ONGOING"
	CalendarChannelForResponseSyncStageCALENDAREVENTLISTFETCHPENDING   CalendarChannelForResponseSyncStage = "CALENDAR_EVENT_LIST_FETCH_PENDING"
	CalendarChannelForResponseSyncStageCALENDAREVENTLISTFETCHSCHEDULED CalendarChannelForResponseSyncStage = "CALENDAR_EVENT_LIST_FETCH_SCHEDULED"
	CalendarChannelForResponseSyncStageCALENDAREVENTSIMPORTONGOING     CalendarChannelForResponseSyncStage = "CALENDAR_EVENTS_IMPORT_ONGOING"
	CalendarChannelForResponseSyncStageCALENDAREVENTSIMPORTPENDING     CalendarChannelForResponseSyncStage = "CALENDAR_EVENTS_IMPORT_PENDING"
	CalendarChannelForResponseSyncStageCALENDAREVENTSIMPORTSCHEDULED   CalendarChannelForResponseSyncStage = "CALENDAR_EVENTS_IMPORT_SCHEDULED"
	CalendarChannelForResponseSyncStageFAILED                          CalendarChannelForResponseSyncStage = "FAILED"
	CalendarChannelForResponseSyncStagePENDINGCONFIGURATION            CalendarChannelForResponseSyncStage = "PENDING_CONFIGURATION"
)

// Defines values for CalendarChannelForResponseSyncStatus.
const (
	CalendarChannelForResponseSyncStatusACTIVE                        CalendarChannelForResponseSyncStatus = "ACTIVE"
	CalendarChannelForResponseSyncStatusFAILEDINSUFFICIENTPERMISSIONS CalendarChannelForResponseSyncStatus = "FAILED_INSUFFICIENT_PERMISSIONS"
	CalendarChannelForResponseSyncStatusFAILEDUNKNOWN                 CalendarChannelForResponseSyncStatus = "FAILED_UNKNOWN"
	CalendarChannelForResponseSyncStatusNOTSYNCED                     CalendarChannelForResponseSyncStatus = "NOT_SYNCED"
	CalendarChannelForResponseSyncStatusONGOING                       CalendarChannelForResponseSyncStatus = "ONGOING"
)

// Defines values for CalendarChannelForResponseVisibility.
const (
	CalendarChannelForResponseVisibilityMETADATA        CalendarChannelForResponseVisibility = "METADATA"
	CalendarChannelForResponseVisibilitySHAREEVERYTHING CalendarChannelForResponseVisibility = "SHARE_EVERYTHING"
)

// Defines values for CalendarChannelForUpdateContactAutoCreationPolicy.
const (
	CalendarChannelForUpdateContactAutoCreationPolicyASORGANIZER               CalendarChannelForUpdateContactAutoCreationPolicy = "AS_ORGANIZER"
	CalendarChannelForUpdateContactAutoCreationPolicyASPARTICIPANT             CalendarChannelForUpdateContactAutoCreationPolicy = "AS_PARTICIPANT"
	CalendarChannelForUpdateContactAutoCreationPolicyASPARTICIPANTANDORGANIZER CalendarChannelForUpdateContactAutoCreationPolicy = "AS_PARTICIPANT_AND_ORGANIZER"
	CalendarChannelForUpdateContactAutoCreationPolicyNONE                      CalendarChannelForUpdateContactAutoCreationPolicy = "NONE"
)

// Defines values for CalendarChannelForUpdateSyncStage.
const (
	CalendarChannelForUpdateSyncStageCALENDAREVENTLISTFETCHONGOING   CalendarChannelForUpdateSyncStage = "CALENDAR_EVENT_LIST_FETCH_ONGOING"
	CalendarChannelForUpdateSyncStageCALENDAREVENTLISTFETCHPENDING   CalendarChannelForUpdateSyncStage = "CALENDAR_EVENT_LIST_FETCH_PENDING"
	CalendarChannelForUpdateSyncStageCALENDAREVENTLISTFETCHSCHEDULED CalendarChannelForUpdateSyncStage = "CALENDAR_EVENT_LIST_FETCH_SCHEDULED"
	CalendarChannelForUpdateSyncStageCALENDAREVENTSIMPORTONGOING     CalendarChannelForUpdateSyncStage = "CALENDAR_EVENTS_IMPORT_ONGOING"
	CalendarChannelForUpdateSyncStageCALENDAREVENTSIMPORTPENDING     CalendarChannelForUpdateSyncStage = "CALENDAR_EVENTS_IMPORT_PENDING"
	CalendarChannelForUpdateSyncStageCALENDAREVENTSIMPORTSCHEDULED   CalendarChannelForUpdateSyncStage = "CALENDAR_EVENTS_IMPORT_SCHEDULED"
	CalendarChannelForUpdateSyncStageFAILED                          CalendarChannelForUpdateSyncStage = "FAILED"
	CalendarChannelForUpdateSyncStagePENDINGCONFIGURATION            CalendarChannelForUpdateSyncStage = "PENDING_CONFIGURATION"
)

// Defines values for CalendarChannelForUpdateSyncStatus.
const (
	CalendarChannelForUpdateSyncStatusACTIVE                        CalendarChannelForUpdateSyncStatus = "ACTIVE"
	CalendarChannelForUpdateSyncStatusFAILEDINSUFFICIENTPERMISSIONS CalendarChannelForUpdateSyncStatus = "FAILED_INSUFFICIENT_PERMISSIONS"
	CalendarChannelForUpdateSyncStatusFAILEDUNKNOWN                 CalendarChannelForUpdateSyncStatus = "FAILED_UNKNOWN"
	CalendarChannelForUpdateSyncStatusNOTSYNCED                     CalendarChannelForUpdateSyncStatus = "NOT_SYNCED"
	CalendarChannelForUpdateSyncStatusONGOING                       CalendarChannelForUpdateSyncStatus = "ONGOING"
)

// Defines values for CalendarChannelForUpdateVisibility.
const (
	CalendarChannelForUpdateVisibilityMETADATA        CalendarChannelForUpdateVisibility = "METADATA"
	CalendarChannelForUpdateVisibilitySHAREEVERYTHING CalendarChannelForUpdateVisibility = "SHARE_EVERYTHING"
)

// Defines values for CalendarEventParticipantResponseStatus.
const (
	CalendarEventParticipantResponseStatusACCEPTED    CalendarEventParticipantResponseStatus = "ACCEPTED"
	CalendarEventParticipantResponseStatusDECLINED    CalendarEventParticipantResponseStatus = "DECLINED"
	CalendarEventParticipantResponseStatusNEEDSACTION CalendarEventParticipantResponseStatus = "NEEDS_ACTION"
	CalendarEventParticipantResponseStatusTENTATIVE   CalendarEventParticipantResponseStatus = "TENTATIVE"
)

// Defines values for CalendarEventParticipantForResponseResponseStatus.
const (
	CalendarEventParticipantForResponseResponseStatusACCEPTED    CalendarEventParticipantForResponseResponseStatus = "ACCEPTED"
	CalendarEventParticipantForResponseResponseStatusDECLINED    CalendarEventParticipantForResponseResponseStatus = "DECLINED"
	CalendarEventParticipantForResponseResponseStatusNEEDSACTION CalendarEventParticipantForResponseResponseStatus = "NEEDS_ACTION"
	CalendarEventParticipantForResponseResponseStatusTENTATIVE   CalendarEventParticipantForResponseResponseStatus = "TENTATIVE"
)

// Defines values for CalendarEventParticipantForUpdateResponseStatus.
const (
	ACCEPTED    CalendarEventParticipantForUpdateResponseStatus = "ACCEPTED"
	DECLINED    CalendarEventParticipantForUpdateResponseStatus = "DECLINED"
	NEEDSACTION CalendarEventParticipantForUpdateResponseStatus = "NEEDS_ACTION"
	TENTATIVE   CalendarEventParticipantForUpdateResponseStatus = "TENTATIVE"
)

// Defines values for CompanyCreatedBySource.
const (
	CompanyCreatedBySourceAGENT    CompanyCreatedBySource = "AGENT"
	CompanyCreatedBySourceAPI      CompanyCreatedBySource = "API"
	CompanyCreatedBySourceCALENDAR CompanyCreatedBySource = "CALENDAR"
	CompanyCreatedBySourceEMAIL    CompanyCreatedBySource = "EMAIL"
	CompanyCreatedBySourceIMPORT   CompanyCreatedBySource = "IMPORT"
	CompanyCreatedBySourceMANUAL   CompanyCreatedBySource = "MANUAL"
	CompanyCreatedBySourceSYSTEM   CompanyCreatedBySource = "SYSTEM"
	CompanyCreatedBySourceWEBHOOK  CompanyCreatedBySource = "WEBHOOK"
	CompanyCreatedBySourceWORKFLOW CompanyCreatedBySource = "WORKFLOW"
)

// Defines values for CompanyForResponseCreatedBySource.
const (
	CompanyForResponseCreatedBySourceAGENT    CompanyForResponseCreatedBySource = "AGENT"
	CompanyForResponseCreatedBySourceAPI      CompanyForResponseCreatedBySource = "API"
	CompanyForResponseCreatedBySourceCALENDAR CompanyForResponseCreatedBySource = "CALENDAR"
	CompanyForResponseCreatedBySourceEMAIL    CompanyForResponseCreatedBySource = "EMAIL"
	CompanyForResponseCreatedBySourceIMPORT   CompanyForResponseCreatedBySource = "IMPORT"
	CompanyForResponseCreatedBySourceMANUAL   CompanyForResponseCreatedBySource = "MANUAL"
	CompanyForResponseCreatedBySourceSYSTEM   CompanyForResponseCreatedBySource = "SYSTEM"
	CompanyForResponseCreatedBySourceWEBHOOK  CompanyForResponseCreatedBySource = "WEBHOOK"
	CompanyForResponseCreatedBySourceWORKFLOW CompanyForResponseCreatedBySource = "WORKFLOW"
)

// Defines values for CompanyForUpdateCreatedBySource.
const (
	CompanyForUpdateCreatedBySourceAGENT    CompanyForUpdateCreatedBySource = "AGENT"
	CompanyForUpdateCreatedBySourceAPI      CompanyForUpdateCreatedBySource = "API"
	CompanyForUpdateCreatedBySourceCALENDAR CompanyForUpdateCreatedBySource = "CALENDAR"
	CompanyForUpdateCreatedBySourceEMAIL    CompanyForUpdateCreatedBySource = "EMAIL"
	CompanyForUpdateCreatedBySourceIMPORT   CompanyForUpdateCreatedBySource = "IMPORT"
	CompanyForUpdateCreatedBySourceMANUAL   CompanyForUpdateCreatedBySource = "MANUAL"
	CompanyForUpdateCreatedBySourceSYSTEM   CompanyForUpdateCreatedBySource = "SYSTEM"
	CompanyForUpdateCreatedBySourceWEBHOOK  CompanyForUpdateCreatedBySource = "WEBHOOK"
	CompanyForUpdateCreatedBySourceWORKFLOW CompanyForUpdateCreatedBySource = "WORKFLOW"
)

// Defines values for DashboardCreatedBySource.
const (
	DashboardCreatedBySourceAGENT    DashboardCreatedBySource = "AGENT"
	DashboardCreatedBySourceAPI      DashboardCreatedBySource = "API"
	DashboardCreatedBySourceCALENDAR DashboardCreatedBySource = "CALENDAR"
	DashboardCreatedBySourceEMAIL    DashboardCreatedBySource = "EMAIL"
	DashboardCreatedBySourceIMPORT   DashboardCreatedBySource = "IMPORT"
	DashboardCreatedBySourceMANUAL   DashboardCreatedBySource = "MANUAL"
	DashboardCreatedBySourceSYSTEM   DashboardCreatedBySource = "SYSTEM"
	DashboardCreatedBySourceWEBHOOK  DashboardCreatedBySource = "WEBHOOK"
	DashboardCreatedBySourceWORKFLOW DashboardCreatedBySource = "WORKFLOW"
)

// Defines values for DashboardForResponseCreatedBySource.
const (
	DashboardForResponseCreatedBySourceAGENT    DashboardForResponseCreatedBySource = "AGENT"
	DashboardForResponseCreatedBySourceAPI      DashboardForResponseCreatedBySource = "API"
	DashboardForResponseCreatedBySourceCALENDAR DashboardForResponseCreatedBySource = "CALENDAR"
	DashboardForResponseCreatedBySourceEMAIL    DashboardForResponseCreatedBySource = "EMAIL"
	DashboardForResponseCreatedBySourceIMPORT   DashboardForResponseCreatedBySource = "IMPORT"
	DashboardForResponseCreatedBySourceMANUAL   DashboardForResponseCreatedBySource = "MANUAL"
	DashboardForResponseCreatedBySourceSYSTEM   DashboardForResponseCreatedBySource = "SYSTEM"
	DashboardForResponseCreatedBySourceWEBHOOK  DashboardForResponseCreatedBySource = "WEBHOOK"
	DashboardForResponseCreatedBySourceWORKFLOW DashboardForResponseCreatedBySource = "WORKFLOW"
)

// Defines values for DashboardForUpdateCreatedBySource.
const (
	DashboardForUpdateCreatedBySourceAGENT    DashboardForUpdateCreatedBySource = "AGENT"
	DashboardForUpdateCreatedBySourceAPI      DashboardForUpdateCreatedBySource = "API"
	DashboardForUpdateCreatedBySourceCALENDAR DashboardForUpdateCreatedBySource = "CALENDAR"
	DashboardForUpdateCreatedBySourceEMAIL    DashboardForUpdateCreatedBySource = "EMAIL"
	DashboardForUpdateCreatedBySourceIMPORT   DashboardForUpdateCreatedBySource = "IMPORT"
	DashboardForUpdateCreatedBySourceMANUAL   DashboardForUpdateCreatedBySource = "MANUAL"
	DashboardForUpdateCreatedBySourceSYSTEM   DashboardForUpdateCreatedBySource = "SYSTEM"
	DashboardForUpdateCreatedBySourceWEBHOOK  DashboardForUpdateCreatedBySource = "WEBHOOK"
	DashboardForUpdateCreatedBySourceWORKFLOW DashboardForUpdateCreatedBySource = "WORKFLOW"
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

// Defines values for MessageChannelForResponseContactAutoCreationPolicy.
const (
	MessageChannelForResponseContactAutoCreationPolicyNONE            MessageChannelForResponseContactAutoCreationPolicy = "NONE"
	MessageChannelForResponseContactAutoCreationPolicySENT            MessageChannelForResponseContactAutoCreationPolicy = "SENT"
	MessageChannelForResponseContactAutoCreationPolicySENTANDRECEIVED MessageChannelForResponseContactAutoCreationPolicy = "SENT_AND_RECEIVED"
)

// Defines values for MessageChannelForResponseMessageFolderImportPolicy.
const (
	MessageChannelForResponseMessageFolderImportPolicyALLFOLDERS      MessageChannelForResponseMessageFolderImportPolicy = "ALL_FOLDERS"
	MessageChannelForResponseMessageFolderImportPolicySELECTEDFOLDERS MessageChannelForResponseMessageFolderImportPolicy = "SELECTED_FOLDERS"
)

// Defines values for MessageChannelForResponsePendingGroupEmailsAction.
const (
	MessageChannelForResponsePendingGroupEmailsActionGROUPEMAILSDELETION MessageChannelForResponsePendingGroupEmailsAction = "GROUP_EMAILS_DELETION"
	MessageChannelForResponsePendingGroupEmailsActionGROUPEMAILSIMPORT   MessageChannelForResponsePendingGroupEmailsAction = "GROUP_EMAILS_IMPORT"
	MessageChannelForResponsePendingGroupEmailsActionNONE                MessageChannelForResponsePendingGroupEmailsAction = "NONE"
)

// Defines values for MessageChannelForResponseSyncStage.
const (
	MessageChannelForResponseSyncStageFAILED                    MessageChannelForResponseSyncStage = "FAILED"
	MessageChannelForResponseSyncStageMESSAGELISTFETCHONGOING   MessageChannelForResponseSyncStage = "MESSAGE_LIST_FETCH_ONGOING"
	MessageChannelForResponseSyncStageMESSAGELISTFETCHPENDING   MessageChannelForResponseSyncStage = "MESSAGE_LIST_FETCH_PENDING"
	MessageChannelForResponseSyncStageMESSAGELISTFETCHSCHEDULED MessageChannelForResponseSyncStage = "MESSAGE_LIST_FETCH_SCHEDULED"
	MessageChannelForResponseSyncStageMESSAGESIMPORTONGOING     MessageChannelForResponseSyncStage = "MESSAGES_IMPORT_ONGOING"
	MessageChannelForResponseSyncStageMESSAGESIMPORTPENDING     MessageChannelForResponseSyncStage = "MESSAGES_IMPORT_PENDING"
	MessageChannelForResponseSyncStageMESSAGESIMPORTSCHEDULED   MessageChannelForResponseSyncStage = "MESSAGES_IMPORT_SCHEDULED"
	MessageChannelForResponseSyncStagePENDINGCONFIGURATION      MessageChannelForResponseSyncStage = "PENDING_CONFIGURATION"
)

// Defines values for MessageChannelForResponseSyncStatus.
const (
	MessageChannelForResponseSyncStatusACTIVE                        MessageChannelForResponseSyncStatus = "ACTIVE"
	MessageChannelForResponseSyncStatusFAILEDINSUFFICIENTPERMISSIONS MessageChannelForResponseSyncStatus = "FAILED_INSUFFICIENT_PERMISSIONS"
	MessageChannelForResponseSyncStatusFAILEDUNKNOWN                 MessageChannelForResponseSyncStatus = "FAILED_UNKNOWN"
	MessageChannelForResponseSyncStatusNOTSYNCED                     MessageChannelForResponseSyncStatus = "NOT_SYNCED"
	MessageChannelForResponseSyncStatusONGOING                       MessageChannelForResponseSyncStatus = "ONGOING"
)

// Defines values for MessageChannelForResponseType.
const (
	MessageChannelForResponseTypeEMAIL MessageChannelForResponseType = "EMAIL"
	MessageChannelForResponseTypeSMS   MessageChannelForResponseType = "SMS"
)

// Defines values for MessageChannelForResponseVisibility.
const (
	MessageChannelForResponseVisibilityMETADATA        MessageChannelForResponseVisibility = "METADATA"
	MessageChannelForResponseVisibilitySHAREEVERYTHING MessageChannelForResponseVisibility = "SHARE_EVERYTHING"
	MessageChannelForResponseVisibilitySUBJECT         MessageChannelForResponseVisibility = "SUBJECT"
)

// Defines values for MessageChannelForUpdateContactAutoCreationPolicy.
const (
	MessageChannelForUpdateContactAutoCreationPolicyNONE            MessageChannelForUpdateContactAutoCreationPolicy = "NONE"
	MessageChannelForUpdateContactAutoCreationPolicySENT            MessageChannelForUpdateContactAutoCreationPolicy = "SENT"
	MessageChannelForUpdateContactAutoCreationPolicySENTANDRECEIVED MessageChannelForUpdateContactAutoCreationPolicy = "SENT_AND_RECEIVED"
)

// Defines values for MessageChannelForUpdateMessageFolderImportPolicy.
const (
	ALLFOLDERS      MessageChannelForUpdateMessageFolderImportPolicy = "ALL_FOLDERS"
	SELECTEDFOLDERS MessageChannelForUpdateMessageFolderImportPolicy = "SELECTED_FOLDERS"
)

// Defines values for MessageChannelForUpdatePendingGroupEmailsAction.
const (
	MessageChannelForUpdatePendingGroupEmailsActionGROUPEMAILSDELETION MessageChannelForUpdatePendingGroupEmailsAction = "GROUP_EMAILS_DELETION"
	MessageChannelForUpdatePendingGroupEmailsActionGROUPEMAILSIMPORT   MessageChannelForUpdatePendingGroupEmailsAction = "GROUP_EMAILS_IMPORT"
	MessageChannelForUpdatePendingGroupEmailsActionNONE                MessageChannelForUpdatePendingGroupEmailsAction = "NONE"
)

// Defines values for MessageChannelForUpdateSyncStage.
const (
	MessageChannelForUpdateSyncStageFAILED                    MessageChannelForUpdateSyncStage = "FAILED"
	MessageChannelForUpdateSyncStageMESSAGELISTFETCHONGOING   MessageChannelForUpdateSyncStage = "MESSAGE_LIST_FETCH_ONGOING"
	MessageChannelForUpdateSyncStageMESSAGELISTFETCHPENDING   MessageChannelForUpdateSyncStage = "MESSAGE_LIST_FETCH_PENDING"
	MessageChannelForUpdateSyncStageMESSAGELISTFETCHSCHEDULED MessageChannelForUpdateSyncStage = "MESSAGE_LIST_FETCH_SCHEDULED"
	MessageChannelForUpdateSyncStageMESSAGESIMPORTONGOING     MessageChannelForUpdateSyncStage = "MESSAGES_IMPORT_ONGOING"
	MessageChannelForUpdateSyncStageMESSAGESIMPORTPENDING     MessageChannelForUpdateSyncStage = "MESSAGES_IMPORT_PENDING"
	MessageChannelForUpdateSyncStageMESSAGESIMPORTSCHEDULED   MessageChannelForUpdateSyncStage = "MESSAGES_IMPORT_SCHEDULED"
	MessageChannelForUpdateSyncStagePENDINGCONFIGURATION      MessageChannelForUpdateSyncStage = "PENDING_CONFIGURATION"
)

// Defines values for MessageChannelForUpdateSyncStatus.
const (
	MessageChannelForUpdateSyncStatusACTIVE                        MessageChannelForUpdateSyncStatus = "ACTIVE"
	MessageChannelForUpdateSyncStatusFAILEDINSUFFICIENTPERMISSIONS MessageChannelForUpdateSyncStatus = "FAILED_INSUFFICIENT_PERMISSIONS"
	MessageChannelForUpdateSyncStatusFAILEDUNKNOWN                 MessageChannelForUpdateSyncStatus = "FAILED_UNKNOWN"
	MessageChannelForUpdateSyncStatusNOTSYNCED                     MessageChannelForUpdateSyncStatus = "NOT_SYNCED"
	MessageChannelForUpdateSyncStatusONGOING                       MessageChannelForUpdateSyncStatus = "ONGOING"
)

// Defines values for MessageChannelForUpdateType.
const (
	MessageChannelForUpdateTypeEMAIL MessageChannelForUpdateType = "EMAIL"
	MessageChannelForUpdateTypeSMS   MessageChannelForUpdateType = "SMS"
)

// Defines values for MessageChannelForUpdateVisibility.
const (
	MessageChannelForUpdateVisibilityMETADATA        MessageChannelForUpdateVisibility = "METADATA"
	MessageChannelForUpdateVisibilitySHAREEVERYTHING MessageChannelForUpdateVisibility = "SHARE_EVERYTHING"
	MessageChannelForUpdateVisibilitySUBJECT         MessageChannelForUpdateVisibility = "SUBJECT"
)

// Defines values for MessageChannelMessageAssociationDirection.
const (
	MessageChannelMessageAssociationDirectionINCOMING MessageChannelMessageAssociationDirection = "INCOMING"
	MessageChannelMessageAssociationDirectionOUTGOING MessageChannelMessageAssociationDirection = "OUTGOING"
)

// Defines values for MessageChannelMessageAssociationForResponseDirection.
const (
	MessageChannelMessageAssociationForResponseDirectionINCOMING MessageChannelMessageAssociationForResponseDirection = "INCOMING"
	MessageChannelMessageAssociationForResponseDirectionOUTGOING MessageChannelMessageAssociationForResponseDirection = "OUTGOING"
)

// Defines values for MessageChannelMessageAssociationForUpdateDirection.
const (
	INCOMING MessageChannelMessageAssociationForUpdateDirection = "INCOMING"
	OUTGOING MessageChannelMessageAssociationForUpdateDirection = "OUTGOING"
)

// Defines values for MessageFolderPendingSyncAction.
const (
	MessageFolderPendingSyncActionFOLDERDELETION MessageFolderPendingSyncAction = "FOLDER_DELETION"
	MessageFolderPendingSyncActionNONE           MessageFolderPendingSyncAction = "NONE"
)

// Defines values for MessageFolderForResponsePendingSyncAction.
const (
	MessageFolderForResponsePendingSyncActionFOLDERDELETION MessageFolderForResponsePendingSyncAction = "FOLDER_DELETION"
	MessageFolderForResponsePendingSyncActionNONE           MessageFolderForResponsePendingSyncAction = "NONE"
)

// Defines values for MessageFolderForUpdatePendingSyncAction.
const (
	MessageFolderForUpdatePendingSyncActionFOLDERDELETION MessageFolderForUpdatePendingSyncAction = "FOLDER_DELETION"
	MessageFolderForUpdatePendingSyncActionNONE           MessageFolderForUpdatePendingSyncAction = "NONE"
)

// Defines values for MessageParticipantRole.
const (
	MessageParticipantRoleBCC  MessageParticipantRole = "BCC"
	MessageParticipantRoleCC   MessageParticipantRole = "CC"
	MessageParticipantRoleFROM MessageParticipantRole = "FROM"
	MessageParticipantRoleTO   MessageParticipantRole = "TO"
)

// Defines values for MessageParticipantForResponseRole.
const (
	MessageParticipantForResponseRoleBCC  MessageParticipantForResponseRole = "BCC"
	MessageParticipantForResponseRoleCC   MessageParticipantForResponseRole = "CC"
	MessageParticipantForResponseRoleFROM MessageParticipantForResponseRole = "FROM"
	MessageParticipantForResponseRoleTO   MessageParticipantForResponseRole = "TO"
)

// Defines values for MessageParticipantForUpdateRole.
const (
	BCC  MessageParticipantForUpdateRole = "BCC"
	CC   MessageParticipantForUpdateRole = "CC"
	FROM MessageParticipantForUpdateRole = "FROM"
	TO   MessageParticipantForUpdateRole = "TO"
)

// Defines values for NoteCreatedBySource.
const (
	NoteCreatedBySourceAGENT    NoteCreatedBySource = "AGENT"
	NoteCreatedBySourceAPI      NoteCreatedBySource = "API"
	NoteCreatedBySourceCALENDAR NoteCreatedBySource = "CALENDAR"
	NoteCreatedBySourceEMAIL    NoteCreatedBySource = "EMAIL"
	NoteCreatedBySourceIMPORT   NoteCreatedBySource = "IMPORT"
	NoteCreatedBySourceMANUAL   NoteCreatedBySource = "MANUAL"
	NoteCreatedBySourceSYSTEM   NoteCreatedBySource = "SYSTEM"
	NoteCreatedBySourceWEBHOOK  NoteCreatedBySource = "WEBHOOK"
	NoteCreatedBySourceWORKFLOW NoteCreatedBySource = "WORKFLOW"
)

// Defines values for NoteForResponseCreatedBySource.
const (
	NoteForResponseCreatedBySourceAGENT    NoteForResponseCreatedBySource = "AGENT"
	NoteForResponseCreatedBySourceAPI      NoteForResponseCreatedBySource = "API"
	NoteForResponseCreatedBySourceCALENDAR NoteForResponseCreatedBySource = "CALENDAR"
	NoteForResponseCreatedBySourceEMAIL    NoteForResponseCreatedBySource = "EMAIL"
	NoteForResponseCreatedBySourceIMPORT   NoteForResponseCreatedBySource = "IMPORT"
	NoteForResponseCreatedBySourceMANUAL   NoteForResponseCreatedBySource = "MANUAL"
	NoteForResponseCreatedBySourceSYSTEM   NoteForResponseCreatedBySource = "SYSTEM"
	NoteForResponseCreatedBySourceWEBHOOK  NoteForResponseCreatedBySource = "WEBHOOK"
	NoteForResponseCreatedBySourceWORKFLOW NoteForResponseCreatedBySource = "WORKFLOW"
)

// Defines values for NoteForUpdateCreatedBySource.
const (
	NoteForUpdateCreatedBySourceAGENT    NoteForUpdateCreatedBySource = "AGENT"
	NoteForUpdateCreatedBySourceAPI      NoteForUpdateCreatedBySource = "API"
	NoteForUpdateCreatedBySourceCALENDAR NoteForUpdateCreatedBySource = "CALENDAR"
	NoteForUpdateCreatedBySourceEMAIL    NoteForUpdateCreatedBySource = "EMAIL"
	NoteForUpdateCreatedBySourceIMPORT   NoteForUpdateCreatedBySource = "IMPORT"
	NoteForUpdateCreatedBySourceMANUAL   NoteForUpdateCreatedBySource = "MANUAL"
	NoteForUpdateCreatedBySourceSYSTEM   NoteForUpdateCreatedBySource = "SYSTEM"
	NoteForUpdateCreatedBySourceWEBHOOK  NoteForUpdateCreatedBySource = "WEBHOOK"
	NoteForUpdateCreatedBySourceWORKFLOW NoteForUpdateCreatedBySource = "WORKFLOW"
)

// Defines values for OpportunityCreatedBySource.
const (
	OpportunityCreatedBySourceAGENT    OpportunityCreatedBySource = "AGENT"
	OpportunityCreatedBySourceAPI      OpportunityCreatedBySource = "API"
	OpportunityCreatedBySourceCALENDAR OpportunityCreatedBySource = "CALENDAR"
	OpportunityCreatedBySourceEMAIL    OpportunityCreatedBySource = "EMAIL"
	OpportunityCreatedBySourceIMPORT   OpportunityCreatedBySource = "IMPORT"
	OpportunityCreatedBySourceMANUAL   OpportunityCreatedBySource = "MANUAL"
	OpportunityCreatedBySourceSYSTEM   OpportunityCreatedBySource = "SYSTEM"
	OpportunityCreatedBySourceWEBHOOK  OpportunityCreatedBySource = "WEBHOOK"
	OpportunityCreatedBySourceWORKFLOW OpportunityCreatedBySource = "WORKFLOW"
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

// Defines values for OpportunityForResponseCreatedBySource.
const (
	OpportunityForResponseCreatedBySourceAGENT    OpportunityForResponseCreatedBySource = "AGENT"
	OpportunityForResponseCreatedBySourceAPI      OpportunityForResponseCreatedBySource = "API"
	OpportunityForResponseCreatedBySourceCALENDAR OpportunityForResponseCreatedBySource = "CALENDAR"
	OpportunityForResponseCreatedBySourceEMAIL    OpportunityForResponseCreatedBySource = "EMAIL"
	OpportunityForResponseCreatedBySourceIMPORT   OpportunityForResponseCreatedBySource = "IMPORT"
	OpportunityForResponseCreatedBySourceMANUAL   OpportunityForResponseCreatedBySource = "MANUAL"
	OpportunityForResponseCreatedBySourceSYSTEM   OpportunityForResponseCreatedBySource = "SYSTEM"
	OpportunityForResponseCreatedBySourceWEBHOOK  OpportunityForResponseCreatedBySource = "WEBHOOK"
	OpportunityForResponseCreatedBySourceWORKFLOW OpportunityForResponseCreatedBySource = "WORKFLOW"
)

// Defines values for OpportunityForResponseLostReason.
const (
	OpportunityForResponseLostReasonCOMPETITION     OpportunityForResponseLostReason = "COMPETITION"
	OpportunityForResponseLostReasonMISSINGFEATURES OpportunityForResponseLostReason = "MISSING_FEATURES"
	OpportunityForResponseLostReasonNOTUSED         OpportunityForResponseLostReason = "NOT_USED"
	OpportunityForResponseLostReasonOTHER           OpportunityForResponseLostReason = "OTHER"
	OpportunityForResponseLostReasonPRICE           OpportunityForResponseLostReason = "PRICE"
)

// Defines values for OpportunityForResponseStage.
const (
	OpportunityForResponseStageACTIVE                           OpportunityForResponseStage = "ACTIVE"
	OpportunityForResponseStageCANCELED                         OpportunityForResponseStage = "CANCELED"
	OpportunityForResponseStageCLOSERCALLVEREINBART             OpportunityForResponseStage = "CLOSER_CALL_VEREINBART"
	OpportunityForResponseStageDEALLOST                         OpportunityForResponseStage = "DEAL_LOST"
	OpportunityForResponseStageDETAILSCALLVEREINBART            OpportunityForResponseStage = "DETAILS_CALL_VEREINBART"
	OpportunityForResponseStageENTSCHEIDUNGSTRAGERHATZUGESTIMMT OpportunityForResponseStage = "ENTSCHEIDUNGSTRAGER_HAT_ZUGESTIMMT"
	OpportunityForResponseStageSETTERCALLVEREINBART             OpportunityForResponseStage = "SETTER_CALL_VEREINBART"
	OpportunityForResponseStageTRIALPERIOD                      OpportunityForResponseStage = "TRIAL_PERIOD"
	OpportunityForResponseStageVERTRAGEVERSANT                  OpportunityForResponseStage = "VERTRAGE_VERSANT"
)

// Defines values for OpportunityForUpdateCreatedBySource.
const (
	OpportunityForUpdateCreatedBySourceAGENT    OpportunityForUpdateCreatedBySource = "AGENT"
	OpportunityForUpdateCreatedBySourceAPI      OpportunityForUpdateCreatedBySource = "API"
	OpportunityForUpdateCreatedBySourceCALENDAR OpportunityForUpdateCreatedBySource = "CALENDAR"
	OpportunityForUpdateCreatedBySourceEMAIL    OpportunityForUpdateCreatedBySource = "EMAIL"
	OpportunityForUpdateCreatedBySourceIMPORT   OpportunityForUpdateCreatedBySource = "IMPORT"
	OpportunityForUpdateCreatedBySourceMANUAL   OpportunityForUpdateCreatedBySource = "MANUAL"
	OpportunityForUpdateCreatedBySourceSYSTEM   OpportunityForUpdateCreatedBySource = "SYSTEM"
	OpportunityForUpdateCreatedBySourceWEBHOOK  OpportunityForUpdateCreatedBySource = "WEBHOOK"
	OpportunityForUpdateCreatedBySourceWORKFLOW OpportunityForUpdateCreatedBySource = "WORKFLOW"
)

// Defines values for OpportunityForUpdateLostReason.
const (
	OpportunityForUpdateLostReasonCOMPETITION     OpportunityForUpdateLostReason = "COMPETITION"
	OpportunityForUpdateLostReasonMISSINGFEATURES OpportunityForUpdateLostReason = "MISSING_FEATURES"
	OpportunityForUpdateLostReasonNOTUSED         OpportunityForUpdateLostReason = "NOT_USED"
	OpportunityForUpdateLostReasonOTHER           OpportunityForUpdateLostReason = "OTHER"
	OpportunityForUpdateLostReasonPRICE           OpportunityForUpdateLostReason = "PRICE"
)

// Defines values for OpportunityForUpdateStage.
const (
	OpportunityForUpdateStageACTIVE                           OpportunityForUpdateStage = "ACTIVE"
	OpportunityForUpdateStageCANCELED                         OpportunityForUpdateStage = "CANCELED"
	OpportunityForUpdateStageCLOSERCALLVEREINBART             OpportunityForUpdateStage = "CLOSER_CALL_VEREINBART"
	OpportunityForUpdateStageDEALLOST                         OpportunityForUpdateStage = "DEAL_LOST"
	OpportunityForUpdateStageDETAILSCALLVEREINBART            OpportunityForUpdateStage = "DETAILS_CALL_VEREINBART"
	OpportunityForUpdateStageENTSCHEIDUNGSTRAGERHATZUGESTIMMT OpportunityForUpdateStage = "ENTSCHEIDUNGSTRAGER_HAT_ZUGESTIMMT"
	OpportunityForUpdateStageSETTERCALLVEREINBART             OpportunityForUpdateStage = "SETTER_CALL_VEREINBART"
	OpportunityForUpdateStageTRIALPERIOD                      OpportunityForUpdateStage = "TRIAL_PERIOD"
	OpportunityForUpdateStageVERTRAGEVERSANT                  OpportunityForUpdateStage = "VERTRAGE_VERSANT"
)

// Defines values for PersonCreatedBySource.
const (
	PersonCreatedBySourceAGENT    PersonCreatedBySource = "AGENT"
	PersonCreatedBySourceAPI      PersonCreatedBySource = "API"
	PersonCreatedBySourceCALENDAR PersonCreatedBySource = "CALENDAR"
	PersonCreatedBySourceEMAIL    PersonCreatedBySource = "EMAIL"
	PersonCreatedBySourceIMPORT   PersonCreatedBySource = "IMPORT"
	PersonCreatedBySourceMANUAL   PersonCreatedBySource = "MANUAL"
	PersonCreatedBySourceSYSTEM   PersonCreatedBySource = "SYSTEM"
	PersonCreatedBySourceWEBHOOK  PersonCreatedBySource = "WEBHOOK"
	PersonCreatedBySourceWORKFLOW PersonCreatedBySource = "WORKFLOW"
)

// Defines values for PersonForResponseCreatedBySource.
const (
	PersonForResponseCreatedBySourceAGENT    PersonForResponseCreatedBySource = "AGENT"
	PersonForResponseCreatedBySourceAPI      PersonForResponseCreatedBySource = "API"
	PersonForResponseCreatedBySourceCALENDAR PersonForResponseCreatedBySource = "CALENDAR"
	PersonForResponseCreatedBySourceEMAIL    PersonForResponseCreatedBySource = "EMAIL"
	PersonForResponseCreatedBySourceIMPORT   PersonForResponseCreatedBySource = "IMPORT"
	PersonForResponseCreatedBySourceMANUAL   PersonForResponseCreatedBySource = "MANUAL"
	PersonForResponseCreatedBySourceSYSTEM   PersonForResponseCreatedBySource = "SYSTEM"
	PersonForResponseCreatedBySourceWEBHOOK  PersonForResponseCreatedBySource = "WEBHOOK"
	PersonForResponseCreatedBySourceWORKFLOW PersonForResponseCreatedBySource = "WORKFLOW"
)

// Defines values for PersonForUpdateCreatedBySource.
const (
	PersonForUpdateCreatedBySourceAGENT    PersonForUpdateCreatedBySource = "AGENT"
	PersonForUpdateCreatedBySourceAPI      PersonForUpdateCreatedBySource = "API"
	PersonForUpdateCreatedBySourceCALENDAR PersonForUpdateCreatedBySource = "CALENDAR"
	PersonForUpdateCreatedBySourceEMAIL    PersonForUpdateCreatedBySource = "EMAIL"
	PersonForUpdateCreatedBySourceIMPORT   PersonForUpdateCreatedBySource = "IMPORT"
	PersonForUpdateCreatedBySourceMANUAL   PersonForUpdateCreatedBySource = "MANUAL"
	PersonForUpdateCreatedBySourceSYSTEM   PersonForUpdateCreatedBySource = "SYSTEM"
	PersonForUpdateCreatedBySourceWEBHOOK  PersonForUpdateCreatedBySource = "WEBHOOK"
	PersonForUpdateCreatedBySourceWORKFLOW PersonForUpdateCreatedBySource = "WORKFLOW"
)

// Defines values for TaskCreatedBySource.
const (
	TaskCreatedBySourceAGENT    TaskCreatedBySource = "AGENT"
	TaskCreatedBySourceAPI      TaskCreatedBySource = "API"
	TaskCreatedBySourceCALENDAR TaskCreatedBySource = "CALENDAR"
	TaskCreatedBySourceEMAIL    TaskCreatedBySource = "EMAIL"
	TaskCreatedBySourceIMPORT   TaskCreatedBySource = "IMPORT"
	TaskCreatedBySourceMANUAL   TaskCreatedBySource = "MANUAL"
	TaskCreatedBySourceSYSTEM   TaskCreatedBySource = "SYSTEM"
	TaskCreatedBySourceWEBHOOK  TaskCreatedBySource = "WEBHOOK"
	TaskCreatedBySourceWORKFLOW TaskCreatedBySource = "WORKFLOW"
)

// Defines values for TaskStatus.
const (
	TaskStatusDONE       TaskStatus = "DONE"
	TaskStatusINPROGRESS TaskStatus = "IN_PROGRESS"
	TaskStatusTODO       TaskStatus = "TODO"
)

// Defines values for TaskForResponseCreatedBySource.
const (
	TaskForResponseCreatedBySourceAGENT    TaskForResponseCreatedBySource = "AGENT"
	TaskForResponseCreatedBySourceAPI      TaskForResponseCreatedBySource = "API"
	TaskForResponseCreatedBySourceCALENDAR TaskForResponseCreatedBySource = "CALENDAR"
	TaskForResponseCreatedBySourceEMAIL    TaskForResponseCreatedBySource = "EMAIL"
	TaskForResponseCreatedBySourceIMPORT   TaskForResponseCreatedBySource = "IMPORT"
	TaskForResponseCreatedBySourceMANUAL   TaskForResponseCreatedBySource = "MANUAL"
	TaskForResponseCreatedBySourceSYSTEM   TaskForResponseCreatedBySource = "SYSTEM"
	TaskForResponseCreatedBySourceWEBHOOK  TaskForResponseCreatedBySource = "WEBHOOK"
	TaskForResponseCreatedBySourceWORKFLOW TaskForResponseCreatedBySource = "WORKFLOW"
)

// Defines values for TaskForResponseStatus.
const (
	TaskForResponseStatusDONE       TaskForResponseStatus = "DONE"
	TaskForResponseStatusINPROGRESS TaskForResponseStatus = "IN_PROGRESS"
	TaskForResponseStatusTODO       TaskForResponseStatus = "TODO"
)

// Defines values for TaskForUpdateCreatedBySource.
const (
	TaskForUpdateCreatedBySourceAGENT    TaskForUpdateCreatedBySource = "AGENT"
	TaskForUpdateCreatedBySourceAPI      TaskForUpdateCreatedBySource = "API"
	TaskForUpdateCreatedBySourceCALENDAR TaskForUpdateCreatedBySource = "CALENDAR"
	TaskForUpdateCreatedBySourceEMAIL    TaskForUpdateCreatedBySource = "EMAIL"
	TaskForUpdateCreatedBySourceIMPORT   TaskForUpdateCreatedBySource = "IMPORT"
	TaskForUpdateCreatedBySourceMANUAL   TaskForUpdateCreatedBySource = "MANUAL"
	TaskForUpdateCreatedBySourceSYSTEM   TaskForUpdateCreatedBySource = "SYSTEM"
	TaskForUpdateCreatedBySourceWEBHOOK  TaskForUpdateCreatedBySource = "WEBHOOK"
	TaskForUpdateCreatedBySourceWORKFLOW TaskForUpdateCreatedBySource = "WORKFLOW"
)

// Defines values for TaskForUpdateStatus.
const (
	DONE       TaskForUpdateStatus = "DONE"
	INPROGRESS TaskForUpdateStatus = "IN_PROGRESS"
	TODO       TaskForUpdateStatus = "TODO"
)

// Defines values for WorkflowCreatedBySource.
const (
	WorkflowCreatedBySourceAGENT    WorkflowCreatedBySource = "AGENT"
	WorkflowCreatedBySourceAPI      WorkflowCreatedBySource = "API"
	WorkflowCreatedBySourceCALENDAR WorkflowCreatedBySource = "CALENDAR"
	WorkflowCreatedBySourceEMAIL    WorkflowCreatedBySource = "EMAIL"
	WorkflowCreatedBySourceIMPORT   WorkflowCreatedBySource = "IMPORT"
	WorkflowCreatedBySourceMANUAL   WorkflowCreatedBySource = "MANUAL"
	WorkflowCreatedBySourceSYSTEM   WorkflowCreatedBySource = "SYSTEM"
	WorkflowCreatedBySourceWEBHOOK  WorkflowCreatedBySource = "WEBHOOK"
	WorkflowCreatedBySourceWORKFLOW WorkflowCreatedBySource = "WORKFLOW"
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

// Defines values for WorkflowAutomatedTriggerForResponseType.
const (
	WorkflowAutomatedTriggerForResponseTypeCRON          WorkflowAutomatedTriggerForResponseType = "CRON"
	WorkflowAutomatedTriggerForResponseTypeDATABASEEVENT WorkflowAutomatedTriggerForResponseType = "DATABASE_EVENT"
)

// Defines values for WorkflowAutomatedTriggerForUpdateType.
const (
	CRON          WorkflowAutomatedTriggerForUpdateType = "CRON"
	DATABASEEVENT WorkflowAutomatedTriggerForUpdateType = "DATABASE_EVENT"
)

// Defines values for WorkflowForResponseCreatedBySource.
const (
	WorkflowForResponseCreatedBySourceAGENT    WorkflowForResponseCreatedBySource = "AGENT"
	WorkflowForResponseCreatedBySourceAPI      WorkflowForResponseCreatedBySource = "API"
	WorkflowForResponseCreatedBySourceCALENDAR WorkflowForResponseCreatedBySource = "CALENDAR"
	WorkflowForResponseCreatedBySourceEMAIL    WorkflowForResponseCreatedBySource = "EMAIL"
	WorkflowForResponseCreatedBySourceIMPORT   WorkflowForResponseCreatedBySource = "IMPORT"
	WorkflowForResponseCreatedBySourceMANUAL   WorkflowForResponseCreatedBySource = "MANUAL"
	WorkflowForResponseCreatedBySourceSYSTEM   WorkflowForResponseCreatedBySource = "SYSTEM"
	WorkflowForResponseCreatedBySourceWEBHOOK  WorkflowForResponseCreatedBySource = "WEBHOOK"
	WorkflowForResponseCreatedBySourceWORKFLOW WorkflowForResponseCreatedBySource = "WORKFLOW"
)

// Defines values for WorkflowForResponseStatuses.
const (
	WorkflowForResponseStatusesACTIVE      WorkflowForResponseStatuses = "ACTIVE"
	WorkflowForResponseStatusesDEACTIVATED WorkflowForResponseStatuses = "DEACTIVATED"
	WorkflowForResponseStatusesDRAFT       WorkflowForResponseStatuses = "DRAFT"
)

// Defines values for WorkflowForUpdateCreatedBySource.
const (
	WorkflowForUpdateCreatedBySourceAGENT    WorkflowForUpdateCreatedBySource = "AGENT"
	WorkflowForUpdateCreatedBySourceAPI      WorkflowForUpdateCreatedBySource = "API"
	WorkflowForUpdateCreatedBySourceCALENDAR WorkflowForUpdateCreatedBySource = "CALENDAR"
	WorkflowForUpdateCreatedBySourceEMAIL    WorkflowForUpdateCreatedBySource = "EMAIL"
	WorkflowForUpdateCreatedBySourceIMPORT   WorkflowForUpdateCreatedBySource = "IMPORT"
	WorkflowForUpdateCreatedBySourceMANUAL   WorkflowForUpdateCreatedBySource = "MANUAL"
	WorkflowForUpdateCreatedBySourceSYSTEM   WorkflowForUpdateCreatedBySource = "SYSTEM"
	WorkflowForUpdateCreatedBySourceWEBHOOK  WorkflowForUpdateCreatedBySource = "WEBHOOK"
	WorkflowForUpdateCreatedBySourceWORKFLOW WorkflowForUpdateCreatedBySource = "WORKFLOW"
)

// Defines values for WorkflowForUpdateStatuses.
const (
	WorkflowForUpdateStatusesACTIVE      WorkflowForUpdateStatuses = "ACTIVE"
	WorkflowForUpdateStatusesDEACTIVATED WorkflowForUpdateStatuses = "DEACTIVATED"
	WorkflowForUpdateStatusesDRAFT       WorkflowForUpdateStatuses = "DRAFT"
)

// Defines values for WorkflowRunCreatedBySource.
const (
	WorkflowRunCreatedBySourceAGENT    WorkflowRunCreatedBySource = "AGENT"
	WorkflowRunCreatedBySourceAPI      WorkflowRunCreatedBySource = "API"
	WorkflowRunCreatedBySourceCALENDAR WorkflowRunCreatedBySource = "CALENDAR"
	WorkflowRunCreatedBySourceEMAIL    WorkflowRunCreatedBySource = "EMAIL"
	WorkflowRunCreatedBySourceIMPORT   WorkflowRunCreatedBySource = "IMPORT"
	WorkflowRunCreatedBySourceMANUAL   WorkflowRunCreatedBySource = "MANUAL"
	WorkflowRunCreatedBySourceSYSTEM   WorkflowRunCreatedBySource = "SYSTEM"
	WorkflowRunCreatedBySourceWEBHOOK  WorkflowRunCreatedBySource = "WEBHOOK"
	WorkflowRunCreatedBySourceWORKFLOW WorkflowRunCreatedBySource = "WORKFLOW"
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

// Defines values for WorkflowRunForResponseCreatedBySource.
const (
	WorkflowRunForResponseCreatedBySourceAGENT    WorkflowRunForResponseCreatedBySource = "AGENT"
	WorkflowRunForResponseCreatedBySourceAPI      WorkflowRunForResponseCreatedBySource = "API"
	WorkflowRunForResponseCreatedBySourceCALENDAR WorkflowRunForResponseCreatedBySource = "CALENDAR"
	WorkflowRunForResponseCreatedBySourceEMAIL    WorkflowRunForResponseCreatedBySource = "EMAIL"
	WorkflowRunForResponseCreatedBySourceIMPORT   WorkflowRunForResponseCreatedBySource = "IMPORT"
	WorkflowRunForResponseCreatedBySourceMANUAL   WorkflowRunForResponseCreatedBySource = "MANUAL"
	WorkflowRunForResponseCreatedBySourceSYSTEM   WorkflowRunForResponseCreatedBySource = "SYSTEM"
	WorkflowRunForResponseCreatedBySourceWEBHOOK  WorkflowRunForResponseCreatedBySource = "WEBHOOK"
	WorkflowRunForResponseCreatedBySourceWORKFLOW WorkflowRunForResponseCreatedBySource = "WORKFLOW"
)

// Defines values for WorkflowRunForResponseStatus.
const (
	WorkflowRunForResponseStatusCOMPLETED  WorkflowRunForResponseStatus = "COMPLETED"
	WorkflowRunForResponseStatusENQUEUED   WorkflowRunForResponseStatus = "ENQUEUED"
	WorkflowRunForResponseStatusFAILED     WorkflowRunForResponseStatus = "FAILED"
	WorkflowRunForResponseStatusNOTSTARTED WorkflowRunForResponseStatus = "NOT_STARTED"
	WorkflowRunForResponseStatusRUNNING    WorkflowRunForResponseStatus = "RUNNING"
	WorkflowRunForResponseStatusSTOPPED    WorkflowRunForResponseStatus = "STOPPED"
	WorkflowRunForResponseStatusSTOPPING   WorkflowRunForResponseStatus = "STOPPING"
)

// Defines values for WorkflowRunForUpdateCreatedBySource.
const (
	WorkflowRunForUpdateCreatedBySourceAGENT    WorkflowRunForUpdateCreatedBySource = "AGENT"
	WorkflowRunForUpdateCreatedBySourceAPI      WorkflowRunForUpdateCreatedBySource = "API"
	WorkflowRunForUpdateCreatedBySourceCALENDAR WorkflowRunForUpdateCreatedBySource = "CALENDAR"
	WorkflowRunForUpdateCreatedBySourceEMAIL    WorkflowRunForUpdateCreatedBySource = "EMAIL"
	WorkflowRunForUpdateCreatedBySourceIMPORT   WorkflowRunForUpdateCreatedBySource = "IMPORT"
	WorkflowRunForUpdateCreatedBySourceMANUAL   WorkflowRunForUpdateCreatedBySource = "MANUAL"
	WorkflowRunForUpdateCreatedBySourceSYSTEM   WorkflowRunForUpdateCreatedBySource = "SYSTEM"
	WorkflowRunForUpdateCreatedBySourceWEBHOOK  WorkflowRunForUpdateCreatedBySource = "WEBHOOK"
	WorkflowRunForUpdateCreatedBySourceWORKFLOW WorkflowRunForUpdateCreatedBySource = "WORKFLOW"
)

// Defines values for WorkflowRunForUpdateStatus.
const (
	COMPLETED  WorkflowRunForUpdateStatus = "COMPLETED"
	ENQUEUED   WorkflowRunForUpdateStatus = "ENQUEUED"
	FAILED     WorkflowRunForUpdateStatus = "FAILED"
	NOTSTARTED WorkflowRunForUpdateStatus = "NOT_STARTED"
	RUNNING    WorkflowRunForUpdateStatus = "RUNNING"
	STOPPED    WorkflowRunForUpdateStatus = "STOPPED"
	STOPPING   WorkflowRunForUpdateStatus = "STOPPING"
)

// Defines values for WorkflowVersionStatus.
const (
	WorkflowVersionStatusACTIVE      WorkflowVersionStatus = "ACTIVE"
	WorkflowVersionStatusARCHIVED    WorkflowVersionStatus = "ARCHIVED"
	WorkflowVersionStatusDEACTIVATED WorkflowVersionStatus = "DEACTIVATED"
	WorkflowVersionStatusDRAFT       WorkflowVersionStatus = "DRAFT"
)

// Defines values for WorkflowVersionForResponseStatus.
const (
	WorkflowVersionForResponseStatusACTIVE      WorkflowVersionForResponseStatus = "ACTIVE"
	WorkflowVersionForResponseStatusARCHIVED    WorkflowVersionForResponseStatus = "ARCHIVED"
	WorkflowVersionForResponseStatusDEACTIVATED WorkflowVersionForResponseStatus = "DEACTIVATED"
	WorkflowVersionForResponseStatusDRAFT       WorkflowVersionForResponseStatus = "DRAFT"
)

// Defines values for WorkflowVersionForUpdateStatus.
const (
	ACTIVE      WorkflowVersionForUpdateStatus = "ACTIVE"
	ARCHIVED    WorkflowVersionForUpdateStatus = "ARCHIVED"
	DEACTIVATED WorkflowVersionForUpdateStatus = "DEACTIVATED"
	DRAFT       WorkflowVersionForUpdateStatus = "DRAFT"
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

// Defines values for WorkspaceMemberForResponseDateFormat.
const (
	WorkspaceMemberForResponseDateFormatDAYFIRST   WorkspaceMemberForResponseDateFormat = "DAY_FIRST"
	WorkspaceMemberForResponseDateFormatMONTHFIRST WorkspaceMemberForResponseDateFormat = "MONTH_FIRST"
	WorkspaceMemberForResponseDateFormatSYSTEM     WorkspaceMemberForResponseDateFormat = "SYSTEM"
	WorkspaceMemberForResponseDateFormatYEARFIRST  WorkspaceMemberForResponseDateFormat = "YEAR_FIRST"
)

// Defines values for WorkspaceMemberForResponseNumberFormat.
const (
	WorkspaceMemberForResponseNumberFormatAPOSTROPHEANDDOT WorkspaceMemberForResponseNumberFormat = "APOSTROPHE_AND_DOT"
	WorkspaceMemberForResponseNumberFormatCOMMASANDDOT     WorkspaceMemberForResponseNumberFormat = "COMMAS_AND_DOT"
	WorkspaceMemberForResponseNumberFormatDOTSANDCOMMA     WorkspaceMemberForResponseNumberFormat = "DOTS_AND_COMMA"
	WorkspaceMemberForResponseNumberFormatSPACESANDCOMMA   WorkspaceMemberForResponseNumberFormat = "SPACES_AND_COMMA"
	WorkspaceMemberForResponseNumberFormatSYSTEM           WorkspaceMemberForResponseNumberFormat = "SYSTEM"
)

// Defines values for WorkspaceMemberForResponseTimeFormat.
const (
	WorkspaceMemberForResponseTimeFormatHOUR12 WorkspaceMemberForResponseTimeFormat = "HOUR_12"
	WorkspaceMemberForResponseTimeFormatHOUR24 WorkspaceMemberForResponseTimeFormat = "HOUR_24"
	WorkspaceMemberForResponseTimeFormatSYSTEM WorkspaceMemberForResponseTimeFormat = "SYSTEM"
)

// Defines values for WorkspaceMemberForUpdateDateFormat.
const (
	WorkspaceMemberForUpdateDateFormatDAYFIRST   WorkspaceMemberForUpdateDateFormat = "DAY_FIRST"
	WorkspaceMemberForUpdateDateFormatMONTHFIRST WorkspaceMemberForUpdateDateFormat = "MONTH_FIRST"
	WorkspaceMemberForUpdateDateFormatSYSTEM     WorkspaceMemberForUpdateDateFormat = "SYSTEM"
	WorkspaceMemberForUpdateDateFormatYEARFIRST  WorkspaceMemberForUpdateDateFormat = "YEAR_FIRST"
)

// Defines values for WorkspaceMemberForUpdateNumberFormat.
const (
	WorkspaceMemberForUpdateNumberFormatAPOSTROPHEANDDOT WorkspaceMemberForUpdateNumberFormat = "APOSTROPHE_AND_DOT"
	WorkspaceMemberForUpdateNumberFormatCOMMASANDDOT     WorkspaceMemberForUpdateNumberFormat = "COMMAS_AND_DOT"
	WorkspaceMemberForUpdateNumberFormatDOTSANDCOMMA     WorkspaceMemberForUpdateNumberFormat = "DOTS_AND_COMMA"
	WorkspaceMemberForUpdateNumberFormatSPACESANDCOMMA   WorkspaceMemberForUpdateNumberFormat = "SPACES_AND_COMMA"
	WorkspaceMemberForUpdateNumberFormatSYSTEM           WorkspaceMemberForUpdateNumberFormat = "SYSTEM"
)

// Defines values for WorkspaceMemberForUpdateTimeFormat.
const (
	HOUR12 WorkspaceMemberForUpdateTimeFormat = "HOUR_12"
	HOUR24 WorkspaceMemberForUpdateTimeFormat = "HOUR_24"
	SYSTEM WorkspaceMemberForUpdateTimeFormat = "SYSTEM"
)

// Defines values for Depth.
const (
	DepthN0 Depth = 0
	DepthN1 Depth = 1
)

// Defines values for FindManyAttachmentsParamsDepth.
const (
	FindManyAttachmentsParamsDepthN0 FindManyAttachmentsParamsDepth = 0
	FindManyAttachmentsParamsDepthN1 FindManyAttachmentsParamsDepth = 1
)

// Defines values for UpdateManyAttachmentsParamsDepth.
const (
	UpdateManyAttachmentsParamsDepthN0 UpdateManyAttachmentsParamsDepth = 0
	UpdateManyAttachmentsParamsDepthN1 UpdateManyAttachmentsParamsDepth = 1
)

// Defines values for CreateOneAttachmentParamsDepth.
const (
	CreateOneAttachmentParamsDepthN0 CreateOneAttachmentParamsDepth = 0
	CreateOneAttachmentParamsDepthN1 CreateOneAttachmentParamsDepth = 1
)

// Defines values for FindAttachmentDuplicatesParamsDepth.
const (
	FindAttachmentDuplicatesParamsDepthN0 FindAttachmentDuplicatesParamsDepth = 0
	FindAttachmentDuplicatesParamsDepthN1 FindAttachmentDuplicatesParamsDepth = 1
)

// Defines values for MergeManyAttachmentsParamsDepth.
const (
	MergeManyAttachmentsParamsDepthN0 MergeManyAttachmentsParamsDepth = 0
	MergeManyAttachmentsParamsDepthN1 MergeManyAttachmentsParamsDepth = 1
)

// Defines values for FindOneAttachmentParamsDepth.
const (
	FindOneAttachmentParamsDepthN0 FindOneAttachmentParamsDepth = 0
	FindOneAttachmentParamsDepthN1 FindOneAttachmentParamsDepth = 1
)

// Defines values for UpdateOneAttachmentParamsDepth.
const (
	UpdateOneAttachmentParamsDepthN0 UpdateOneAttachmentParamsDepth = 0
	UpdateOneAttachmentParamsDepthN1 UpdateOneAttachmentParamsDepth = 1
)

// Defines values for CreateManyAttachmentsParamsDepth.
const (
	CreateManyAttachmentsParamsDepthN0 CreateManyAttachmentsParamsDepth = 0
	CreateManyAttachmentsParamsDepthN1 CreateManyAttachmentsParamsDepth = 1
)

// Defines values for CreateManyBlocklistsParamsDepth.
const (
	CreateManyBlocklistsParamsDepthN0 CreateManyBlocklistsParamsDepth = 0
	CreateManyBlocklistsParamsDepthN1 CreateManyBlocklistsParamsDepth = 1
)

// Defines values for CreateManyCalendarChannelEventAssociationsParamsDepth.
const (
	CreateManyCalendarChannelEventAssociationsParamsDepthN0 CreateManyCalendarChannelEventAssociationsParamsDepth = 0
	CreateManyCalendarChannelEventAssociationsParamsDepthN1 CreateManyCalendarChannelEventAssociationsParamsDepth = 1
)

// Defines values for CreateManyCalendarChannelsParamsDepth.
const (
	CreateManyCalendarChannelsParamsDepthN0 CreateManyCalendarChannelsParamsDepth = 0
	CreateManyCalendarChannelsParamsDepthN1 CreateManyCalendarChannelsParamsDepth = 1
)

// Defines values for CreateManyCalendarEventParticipantsParamsDepth.
const (
	CreateManyCalendarEventParticipantsParamsDepthN0 CreateManyCalendarEventParticipantsParamsDepth = 0
	CreateManyCalendarEventParticipantsParamsDepthN1 CreateManyCalendarEventParticipantsParamsDepth = 1
)

// Defines values for CreateManyCalendarEventsParamsDepth.
const (
	CreateManyCalendarEventsParamsDepthN0 CreateManyCalendarEventsParamsDepth = 0
	CreateManyCalendarEventsParamsDepthN1 CreateManyCalendarEventsParamsDepth = 1
)

// Defines values for CreateManyCompaniesParamsDepth.
const (
	CreateManyCompaniesParamsDepthN0 CreateManyCompaniesParamsDepth = 0
	CreateManyCompaniesParamsDepthN1 CreateManyCompaniesParamsDepth = 1
)

// Defines values for CreateManyConnectedAccountsParamsDepth.
const (
	CreateManyConnectedAccountsParamsDepthN0 CreateManyConnectedAccountsParamsDepth = 0
	CreateManyConnectedAccountsParamsDepthN1 CreateManyConnectedAccountsParamsDepth = 1
)

// Defines values for CreateManyDashboardsParamsDepth.
const (
	CreateManyDashboardsParamsDepthN0 CreateManyDashboardsParamsDepth = 0
	CreateManyDashboardsParamsDepthN1 CreateManyDashboardsParamsDepth = 1
)

// Defines values for CreateManyFavoriteFoldersParamsDepth.
const (
	CreateManyFavoriteFoldersParamsDepthN0 CreateManyFavoriteFoldersParamsDepth = 0
	CreateManyFavoriteFoldersParamsDepthN1 CreateManyFavoriteFoldersParamsDepth = 1
)

// Defines values for CreateManyFavoritesParamsDepth.
const (
	CreateManyFavoritesParamsDepthN0 CreateManyFavoritesParamsDepth = 0
	CreateManyFavoritesParamsDepthN1 CreateManyFavoritesParamsDepth = 1
)

// Defines values for CreateManyMessageChannelMessageAssociationsParamsDepth.
const (
	CreateManyMessageChannelMessageAssociationsParamsDepthN0 CreateManyMessageChannelMessageAssociationsParamsDepth = 0
	CreateManyMessageChannelMessageAssociationsParamsDepthN1 CreateManyMessageChannelMessageAssociationsParamsDepth = 1
)

// Defines values for CreateManyMessageChannelsParamsDepth.
const (
	CreateManyMessageChannelsParamsDepthN0 CreateManyMessageChannelsParamsDepth = 0
	CreateManyMessageChannelsParamsDepthN1 CreateManyMessageChannelsParamsDepth = 1
)

// Defines values for CreateManyMessageFoldersParamsDepth.
const (
	CreateManyMessageFoldersParamsDepthN0 CreateManyMessageFoldersParamsDepth = 0
	CreateManyMessageFoldersParamsDepthN1 CreateManyMessageFoldersParamsDepth = 1
)

// Defines values for CreateManyMessageParticipantsParamsDepth.
const (
	CreateManyMessageParticipantsParamsDepthN0 CreateManyMessageParticipantsParamsDepth = 0
	CreateManyMessageParticipantsParamsDepthN1 CreateManyMessageParticipantsParamsDepth = 1
)

// Defines values for CreateManyMessageThreadsParamsDepth.
const (
	CreateManyMessageThreadsParamsDepthN0 CreateManyMessageThreadsParamsDepth = 0
	CreateManyMessageThreadsParamsDepthN1 CreateManyMessageThreadsParamsDepth = 1
)

// Defines values for CreateManyMessagesParamsDepth.
const (
	CreateManyMessagesParamsDepthN0 CreateManyMessagesParamsDepth = 0
	CreateManyMessagesParamsDepthN1 CreateManyMessagesParamsDepth = 1
)

// Defines values for CreateManyNoteTargetsParamsDepth.
const (
	CreateManyNoteTargetsParamsDepthN0 CreateManyNoteTargetsParamsDepth = 0
	CreateManyNoteTargetsParamsDepthN1 CreateManyNoteTargetsParamsDepth = 1
)

// Defines values for CreateManyNotesParamsDepth.
const (
	CreateManyNotesParamsDepthN0 CreateManyNotesParamsDepth = 0
	CreateManyNotesParamsDepthN1 CreateManyNotesParamsDepth = 1
)

// Defines values for CreateManyOpportunitiesParamsDepth.
const (
	CreateManyOpportunitiesParamsDepthN0 CreateManyOpportunitiesParamsDepth = 0
	CreateManyOpportunitiesParamsDepthN1 CreateManyOpportunitiesParamsDepth = 1
)

// Defines values for CreateManyPeopleParamsDepth.
const (
	CreateManyPeopleParamsDepthN0 CreateManyPeopleParamsDepth = 0
	CreateManyPeopleParamsDepthN1 CreateManyPeopleParamsDepth = 1
)

// Defines values for CreateManyTaskTargetsParamsDepth.
const (
	CreateManyTaskTargetsParamsDepthN0 CreateManyTaskTargetsParamsDepth = 0
	CreateManyTaskTargetsParamsDepthN1 CreateManyTaskTargetsParamsDepth = 1
)

// Defines values for CreateManyTasksParamsDepth.
const (
	CreateManyTasksParamsDepthN0 CreateManyTasksParamsDepth = 0
	CreateManyTasksParamsDepthN1 CreateManyTasksParamsDepth = 1
)

// Defines values for CreateManyTimelineActivitiesParamsDepth.
const (
	CreateManyTimelineActivitiesParamsDepthN0 CreateManyTimelineActivitiesParamsDepth = 0
	CreateManyTimelineActivitiesParamsDepthN1 CreateManyTimelineActivitiesParamsDepth = 1
)

// Defines values for CreateManyWorkflowAutomatedTriggersParamsDepth.
const (
	CreateManyWorkflowAutomatedTriggersParamsDepthN0 CreateManyWorkflowAutomatedTriggersParamsDepth = 0
	CreateManyWorkflowAutomatedTriggersParamsDepthN1 CreateManyWorkflowAutomatedTriggersParamsDepth = 1
)

// Defines values for CreateManyWorkflowRunsParamsDepth.
const (
	CreateManyWorkflowRunsParamsDepthN0 CreateManyWorkflowRunsParamsDepth = 0
	CreateManyWorkflowRunsParamsDepthN1 CreateManyWorkflowRunsParamsDepth = 1
)

// Defines values for CreateManyWorkflowVersionsParamsDepth.
const (
	CreateManyWorkflowVersionsParamsDepthN0 CreateManyWorkflowVersionsParamsDepth = 0
	CreateManyWorkflowVersionsParamsDepthN1 CreateManyWorkflowVersionsParamsDepth = 1
)

// Defines values for CreateManyWorkflowsParamsDepth.
const (
	CreateManyWorkflowsParamsDepthN0 CreateManyWorkflowsParamsDepth = 0
	CreateManyWorkflowsParamsDepthN1 CreateManyWorkflowsParamsDepth = 1
)

// Defines values for CreateManyWorkspaceMembersParamsDepth.
const (
	CreateManyWorkspaceMembersParamsDepthN0 CreateManyWorkspaceMembersParamsDepth = 0
	CreateManyWorkspaceMembersParamsDepthN1 CreateManyWorkspaceMembersParamsDepth = 1
)

// Defines values for FindManyBlocklistsParamsDepth.
const (
	FindManyBlocklistsParamsDepthN0 FindManyBlocklistsParamsDepth = 0
	FindManyBlocklistsParamsDepthN1 FindManyBlocklistsParamsDepth = 1
)

// Defines values for UpdateManyBlocklistsParamsDepth.
const (
	UpdateManyBlocklistsParamsDepthN0 UpdateManyBlocklistsParamsDepth = 0
	UpdateManyBlocklistsParamsDepthN1 UpdateManyBlocklistsParamsDepth = 1
)

// Defines values for CreateOneBlocklistParamsDepth.
const (
	CreateOneBlocklistParamsDepthN0 CreateOneBlocklistParamsDepth = 0
	CreateOneBlocklistParamsDepthN1 CreateOneBlocklistParamsDepth = 1
)

// Defines values for FindBlocklistDuplicatesParamsDepth.
const (
	FindBlocklistDuplicatesParamsDepthN0 FindBlocklistDuplicatesParamsDepth = 0
	FindBlocklistDuplicatesParamsDepthN1 FindBlocklistDuplicatesParamsDepth = 1
)

// Defines values for MergeManyBlocklistsParamsDepth.
const (
	MergeManyBlocklistsParamsDepthN0 MergeManyBlocklistsParamsDepth = 0
	MergeManyBlocklistsParamsDepthN1 MergeManyBlocklistsParamsDepth = 1
)

// Defines values for FindOneBlocklistParamsDepth.
const (
	FindOneBlocklistParamsDepthN0 FindOneBlocklistParamsDepth = 0
	FindOneBlocklistParamsDepthN1 FindOneBlocklistParamsDepth = 1
)

// Defines values for UpdateOneBlocklistParamsDepth.
const (
	UpdateOneBlocklistParamsDepthN0 UpdateOneBlocklistParamsDepth = 0
	UpdateOneBlocklistParamsDepthN1 UpdateOneBlocklistParamsDepth = 1
)

// Defines values for FindManyCalendarChannelEventAssociationsParamsDepth.
const (
	FindManyCalendarChannelEventAssociationsParamsDepthN0 FindManyCalendarChannelEventAssociationsParamsDepth = 0
	FindManyCalendarChannelEventAssociationsParamsDepthN1 FindManyCalendarChannelEventAssociationsParamsDepth = 1
)

// Defines values for UpdateManyCalendarChannelEventAssociationsParamsDepth.
const (
	UpdateManyCalendarChannelEventAssociationsParamsDepthN0 UpdateManyCalendarChannelEventAssociationsParamsDepth = 0
	UpdateManyCalendarChannelEventAssociationsParamsDepthN1 UpdateManyCalendarChannelEventAssociationsParamsDepth = 1
)

// Defines values for CreateOneCalendarChannelEventAssociationParamsDepth.
const (
	CreateOneCalendarChannelEventAssociationParamsDepthN0 CreateOneCalendarChannelEventAssociationParamsDepth = 0
	CreateOneCalendarChannelEventAssociationParamsDepthN1 CreateOneCalendarChannelEventAssociationParamsDepth = 1
)

// Defines values for FindCalendarChannelEventAssociationDuplicatesParamsDepth.
const (
	FindCalendarChannelEventAssociationDuplicatesParamsDepthN0 FindCalendarChannelEventAssociationDuplicatesParamsDepth = 0
	FindCalendarChannelEventAssociationDuplicatesParamsDepthN1 FindCalendarChannelEventAssociationDuplicatesParamsDepth = 1
)

// Defines values for MergeManyCalendarChannelEventAssociationsParamsDepth.
const (
	MergeManyCalendarChannelEventAssociationsParamsDepthN0 MergeManyCalendarChannelEventAssociationsParamsDepth = 0
	MergeManyCalendarChannelEventAssociationsParamsDepthN1 MergeManyCalendarChannelEventAssociationsParamsDepth = 1
)

// Defines values for FindOneCalendarChannelEventAssociationParamsDepth.
const (
	FindOneCalendarChannelEventAssociationParamsDepthN0 FindOneCalendarChannelEventAssociationParamsDepth = 0
	FindOneCalendarChannelEventAssociationParamsDepthN1 FindOneCalendarChannelEventAssociationParamsDepth = 1
)

// Defines values for UpdateOneCalendarChannelEventAssociationParamsDepth.
const (
	UpdateOneCalendarChannelEventAssociationParamsDepthN0 UpdateOneCalendarChannelEventAssociationParamsDepth = 0
	UpdateOneCalendarChannelEventAssociationParamsDepthN1 UpdateOneCalendarChannelEventAssociationParamsDepth = 1
)

// Defines values for FindManyCalendarChannelsParamsDepth.
const (
	FindManyCalendarChannelsParamsDepthN0 FindManyCalendarChannelsParamsDepth = 0
	FindManyCalendarChannelsParamsDepthN1 FindManyCalendarChannelsParamsDepth = 1
)

// Defines values for UpdateManyCalendarChannelsParamsDepth.
const (
	UpdateManyCalendarChannelsParamsDepthN0 UpdateManyCalendarChannelsParamsDepth = 0
	UpdateManyCalendarChannelsParamsDepthN1 UpdateManyCalendarChannelsParamsDepth = 1
)

// Defines values for CreateOneCalendarChannelParamsDepth.
const (
	CreateOneCalendarChannelParamsDepthN0 CreateOneCalendarChannelParamsDepth = 0
	CreateOneCalendarChannelParamsDepthN1 CreateOneCalendarChannelParamsDepth = 1
)

// Defines values for FindCalendarChannelDuplicatesParamsDepth.
const (
	FindCalendarChannelDuplicatesParamsDepthN0 FindCalendarChannelDuplicatesParamsDepth = 0
	FindCalendarChannelDuplicatesParamsDepthN1 FindCalendarChannelDuplicatesParamsDepth = 1
)

// Defines values for MergeManyCalendarChannelsParamsDepth.
const (
	MergeManyCalendarChannelsParamsDepthN0 MergeManyCalendarChannelsParamsDepth = 0
	MergeManyCalendarChannelsParamsDepthN1 MergeManyCalendarChannelsParamsDepth = 1
)

// Defines values for FindOneCalendarChannelParamsDepth.
const (
	FindOneCalendarChannelParamsDepthN0 FindOneCalendarChannelParamsDepth = 0
	FindOneCalendarChannelParamsDepthN1 FindOneCalendarChannelParamsDepth = 1
)

// Defines values for UpdateOneCalendarChannelParamsDepth.
const (
	UpdateOneCalendarChannelParamsDepthN0 UpdateOneCalendarChannelParamsDepth = 0
	UpdateOneCalendarChannelParamsDepthN1 UpdateOneCalendarChannelParamsDepth = 1
)

// Defines values for FindManyCalendarEventParticipantsParamsDepth.
const (
	FindManyCalendarEventParticipantsParamsDepthN0 FindManyCalendarEventParticipantsParamsDepth = 0
	FindManyCalendarEventParticipantsParamsDepthN1 FindManyCalendarEventParticipantsParamsDepth = 1
)

// Defines values for UpdateManyCalendarEventParticipantsParamsDepth.
const (
	UpdateManyCalendarEventParticipantsParamsDepthN0 UpdateManyCalendarEventParticipantsParamsDepth = 0
	UpdateManyCalendarEventParticipantsParamsDepthN1 UpdateManyCalendarEventParticipantsParamsDepth = 1
)

// Defines values for CreateOneCalendarEventParticipantParamsDepth.
const (
	CreateOneCalendarEventParticipantParamsDepthN0 CreateOneCalendarEventParticipantParamsDepth = 0
	CreateOneCalendarEventParticipantParamsDepthN1 CreateOneCalendarEventParticipantParamsDepth = 1
)

// Defines values for FindCalendarEventParticipantDuplicatesParamsDepth.
const (
	FindCalendarEventParticipantDuplicatesParamsDepthN0 FindCalendarEventParticipantDuplicatesParamsDepth = 0
	FindCalendarEventParticipantDuplicatesParamsDepthN1 FindCalendarEventParticipantDuplicatesParamsDepth = 1
)

// Defines values for MergeManyCalendarEventParticipantsParamsDepth.
const (
	MergeManyCalendarEventParticipantsParamsDepthN0 MergeManyCalendarEventParticipantsParamsDepth = 0
	MergeManyCalendarEventParticipantsParamsDepthN1 MergeManyCalendarEventParticipantsParamsDepth = 1
)

// Defines values for FindOneCalendarEventParticipantParamsDepth.
const (
	FindOneCalendarEventParticipantParamsDepthN0 FindOneCalendarEventParticipantParamsDepth = 0
	FindOneCalendarEventParticipantParamsDepthN1 FindOneCalendarEventParticipantParamsDepth = 1
)

// Defines values for UpdateOneCalendarEventParticipantParamsDepth.
const (
	UpdateOneCalendarEventParticipantParamsDepthN0 UpdateOneCalendarEventParticipantParamsDepth = 0
	UpdateOneCalendarEventParticipantParamsDepthN1 UpdateOneCalendarEventParticipantParamsDepth = 1
)

// Defines values for FindManyCalendarEventsParamsDepth.
const (
	FindManyCalendarEventsParamsDepthN0 FindManyCalendarEventsParamsDepth = 0
	FindManyCalendarEventsParamsDepthN1 FindManyCalendarEventsParamsDepth = 1
)

// Defines values for UpdateManyCalendarEventsParamsDepth.
const (
	UpdateManyCalendarEventsParamsDepthN0 UpdateManyCalendarEventsParamsDepth = 0
	UpdateManyCalendarEventsParamsDepthN1 UpdateManyCalendarEventsParamsDepth = 1
)

// Defines values for CreateOneCalendarEventParamsDepth.
const (
	CreateOneCalendarEventParamsDepthN0 CreateOneCalendarEventParamsDepth = 0
	CreateOneCalendarEventParamsDepthN1 CreateOneCalendarEventParamsDepth = 1
)

// Defines values for FindCalendarEventDuplicatesParamsDepth.
const (
	FindCalendarEventDuplicatesParamsDepthN0 FindCalendarEventDuplicatesParamsDepth = 0
	FindCalendarEventDuplicatesParamsDepthN1 FindCalendarEventDuplicatesParamsDepth = 1
)

// Defines values for MergeManyCalendarEventsParamsDepth.
const (
	MergeManyCalendarEventsParamsDepthN0 MergeManyCalendarEventsParamsDepth = 0
	MergeManyCalendarEventsParamsDepthN1 MergeManyCalendarEventsParamsDepth = 1
)

// Defines values for FindOneCalendarEventParamsDepth.
const (
	FindOneCalendarEventParamsDepthN0 FindOneCalendarEventParamsDepth = 0
	FindOneCalendarEventParamsDepthN1 FindOneCalendarEventParamsDepth = 1
)

// Defines values for UpdateOneCalendarEventParamsDepth.
const (
	UpdateOneCalendarEventParamsDepthN0 UpdateOneCalendarEventParamsDepth = 0
	UpdateOneCalendarEventParamsDepthN1 UpdateOneCalendarEventParamsDepth = 1
)

// Defines values for FindManyCompaniesParamsDepth.
const (
	FindManyCompaniesParamsDepthN0 FindManyCompaniesParamsDepth = 0
	FindManyCompaniesParamsDepthN1 FindManyCompaniesParamsDepth = 1
)

// Defines values for UpdateManyCompaniesParamsDepth.
const (
	UpdateManyCompaniesParamsDepthN0 UpdateManyCompaniesParamsDepth = 0
	UpdateManyCompaniesParamsDepthN1 UpdateManyCompaniesParamsDepth = 1
)

// Defines values for CreateOneCompanyParamsDepth.
const (
	CreateOneCompanyParamsDepthN0 CreateOneCompanyParamsDepth = 0
	CreateOneCompanyParamsDepthN1 CreateOneCompanyParamsDepth = 1
)

// Defines values for FindCompanyDuplicatesParamsDepth.
const (
	FindCompanyDuplicatesParamsDepthN0 FindCompanyDuplicatesParamsDepth = 0
	FindCompanyDuplicatesParamsDepthN1 FindCompanyDuplicatesParamsDepth = 1
)

// Defines values for MergeManyCompaniesParamsDepth.
const (
	MergeManyCompaniesParamsDepthN0 MergeManyCompaniesParamsDepth = 0
	MergeManyCompaniesParamsDepthN1 MergeManyCompaniesParamsDepth = 1
)

// Defines values for FindOneCompanyParamsDepth.
const (
	FindOneCompanyParamsDepthN0 FindOneCompanyParamsDepth = 0
	FindOneCompanyParamsDepthN1 FindOneCompanyParamsDepth = 1
)

// Defines values for UpdateOneCompanyParamsDepth.
const (
	UpdateOneCompanyParamsDepthN0 UpdateOneCompanyParamsDepth = 0
	UpdateOneCompanyParamsDepthN1 UpdateOneCompanyParamsDepth = 1
)

// Defines values for FindManyConnectedAccountsParamsDepth.
const (
	FindManyConnectedAccountsParamsDepthN0 FindManyConnectedAccountsParamsDepth = 0
	FindManyConnectedAccountsParamsDepthN1 FindManyConnectedAccountsParamsDepth = 1
)

// Defines values for UpdateManyConnectedAccountsParamsDepth.
const (
	UpdateManyConnectedAccountsParamsDepthN0 UpdateManyConnectedAccountsParamsDepth = 0
	UpdateManyConnectedAccountsParamsDepthN1 UpdateManyConnectedAccountsParamsDepth = 1
)

// Defines values for CreateOneConnectedAccountParamsDepth.
const (
	CreateOneConnectedAccountParamsDepthN0 CreateOneConnectedAccountParamsDepth = 0
	CreateOneConnectedAccountParamsDepthN1 CreateOneConnectedAccountParamsDepth = 1
)

// Defines values for FindConnectedAccountDuplicatesParamsDepth.
const (
	FindConnectedAccountDuplicatesParamsDepthN0 FindConnectedAccountDuplicatesParamsDepth = 0
	FindConnectedAccountDuplicatesParamsDepthN1 FindConnectedAccountDuplicatesParamsDepth = 1
)

// Defines values for MergeManyConnectedAccountsParamsDepth.
const (
	MergeManyConnectedAccountsParamsDepthN0 MergeManyConnectedAccountsParamsDepth = 0
	MergeManyConnectedAccountsParamsDepthN1 MergeManyConnectedAccountsParamsDepth = 1
)

// Defines values for FindOneConnectedAccountParamsDepth.
const (
	FindOneConnectedAccountParamsDepthN0 FindOneConnectedAccountParamsDepth = 0
	FindOneConnectedAccountParamsDepthN1 FindOneConnectedAccountParamsDepth = 1
)

// Defines values for UpdateOneConnectedAccountParamsDepth.
const (
	UpdateOneConnectedAccountParamsDepthN0 UpdateOneConnectedAccountParamsDepth = 0
	UpdateOneConnectedAccountParamsDepthN1 UpdateOneConnectedAccountParamsDepth = 1
)

// Defines values for FindManyDashboardsParamsDepth.
const (
	FindManyDashboardsParamsDepthN0 FindManyDashboardsParamsDepth = 0
	FindManyDashboardsParamsDepthN1 FindManyDashboardsParamsDepth = 1
)

// Defines values for UpdateManyDashboardsParamsDepth.
const (
	UpdateManyDashboardsParamsDepthN0 UpdateManyDashboardsParamsDepth = 0
	UpdateManyDashboardsParamsDepthN1 UpdateManyDashboardsParamsDepth = 1
)

// Defines values for CreateOneDashboardParamsDepth.
const (
	CreateOneDashboardParamsDepthN0 CreateOneDashboardParamsDepth = 0
	CreateOneDashboardParamsDepthN1 CreateOneDashboardParamsDepth = 1
)

// Defines values for FindDashboardDuplicatesParamsDepth.
const (
	FindDashboardDuplicatesParamsDepthN0 FindDashboardDuplicatesParamsDepth = 0
	FindDashboardDuplicatesParamsDepthN1 FindDashboardDuplicatesParamsDepth = 1
)

// Defines values for MergeManyDashboardsParamsDepth.
const (
	MergeManyDashboardsParamsDepthN0 MergeManyDashboardsParamsDepth = 0
	MergeManyDashboardsParamsDepthN1 MergeManyDashboardsParamsDepth = 1
)

// Defines values for FindOneDashboardParamsDepth.
const (
	FindOneDashboardParamsDepthN0 FindOneDashboardParamsDepth = 0
	FindOneDashboardParamsDepthN1 FindOneDashboardParamsDepth = 1
)

// Defines values for UpdateOneDashboardParamsDepth.
const (
	UpdateOneDashboardParamsDepthN0 UpdateOneDashboardParamsDepth = 0
	UpdateOneDashboardParamsDepthN1 UpdateOneDashboardParamsDepth = 1
)

// Defines values for FindManyFavoriteFoldersParamsDepth.
const (
	FindManyFavoriteFoldersParamsDepthN0 FindManyFavoriteFoldersParamsDepth = 0
	FindManyFavoriteFoldersParamsDepthN1 FindManyFavoriteFoldersParamsDepth = 1
)

// Defines values for UpdateManyFavoriteFoldersParamsDepth.
const (
	UpdateManyFavoriteFoldersParamsDepthN0 UpdateManyFavoriteFoldersParamsDepth = 0
	UpdateManyFavoriteFoldersParamsDepthN1 UpdateManyFavoriteFoldersParamsDepth = 1
)

// Defines values for CreateOneFavoriteFolderParamsDepth.
const (
	CreateOneFavoriteFolderParamsDepthN0 CreateOneFavoriteFolderParamsDepth = 0
	CreateOneFavoriteFolderParamsDepthN1 CreateOneFavoriteFolderParamsDepth = 1
)

// Defines values for FindFavoriteFolderDuplicatesParamsDepth.
const (
	FindFavoriteFolderDuplicatesParamsDepthN0 FindFavoriteFolderDuplicatesParamsDepth = 0
	FindFavoriteFolderDuplicatesParamsDepthN1 FindFavoriteFolderDuplicatesParamsDepth = 1
)

// Defines values for MergeManyFavoriteFoldersParamsDepth.
const (
	MergeManyFavoriteFoldersParamsDepthN0 MergeManyFavoriteFoldersParamsDepth = 0
	MergeManyFavoriteFoldersParamsDepthN1 MergeManyFavoriteFoldersParamsDepth = 1
)

// Defines values for FindOneFavoriteFolderParamsDepth.
const (
	FindOneFavoriteFolderParamsDepthN0 FindOneFavoriteFolderParamsDepth = 0
	FindOneFavoriteFolderParamsDepthN1 FindOneFavoriteFolderParamsDepth = 1
)

// Defines values for UpdateOneFavoriteFolderParamsDepth.
const (
	UpdateOneFavoriteFolderParamsDepthN0 UpdateOneFavoriteFolderParamsDepth = 0
	UpdateOneFavoriteFolderParamsDepthN1 UpdateOneFavoriteFolderParamsDepth = 1
)

// Defines values for FindManyFavoritesParamsDepth.
const (
	FindManyFavoritesParamsDepthN0 FindManyFavoritesParamsDepth = 0
	FindManyFavoritesParamsDepthN1 FindManyFavoritesParamsDepth = 1
)

// Defines values for UpdateManyFavoritesParamsDepth.
const (
	UpdateManyFavoritesParamsDepthN0 UpdateManyFavoritesParamsDepth = 0
	UpdateManyFavoritesParamsDepthN1 UpdateManyFavoritesParamsDepth = 1
)

// Defines values for CreateOneFavoriteParamsDepth.
const (
	CreateOneFavoriteParamsDepthN0 CreateOneFavoriteParamsDepth = 0
	CreateOneFavoriteParamsDepthN1 CreateOneFavoriteParamsDepth = 1
)

// Defines values for FindFavoriteDuplicatesParamsDepth.
const (
	FindFavoriteDuplicatesParamsDepthN0 FindFavoriteDuplicatesParamsDepth = 0
	FindFavoriteDuplicatesParamsDepthN1 FindFavoriteDuplicatesParamsDepth = 1
)

// Defines values for MergeManyFavoritesParamsDepth.
const (
	MergeManyFavoritesParamsDepthN0 MergeManyFavoritesParamsDepth = 0
	MergeManyFavoritesParamsDepthN1 MergeManyFavoritesParamsDepth = 1
)

// Defines values for FindOneFavoriteParamsDepth.
const (
	FindOneFavoriteParamsDepthN0 FindOneFavoriteParamsDepth = 0
	FindOneFavoriteParamsDepthN1 FindOneFavoriteParamsDepth = 1
)

// Defines values for UpdateOneFavoriteParamsDepth.
const (
	UpdateOneFavoriteParamsDepthN0 UpdateOneFavoriteParamsDepth = 0
	UpdateOneFavoriteParamsDepthN1 UpdateOneFavoriteParamsDepth = 1
)

// Defines values for FindManyMessageChannelMessageAssociationsParamsDepth.
const (
	FindManyMessageChannelMessageAssociationsParamsDepthN0 FindManyMessageChannelMessageAssociationsParamsDepth = 0
	FindManyMessageChannelMessageAssociationsParamsDepthN1 FindManyMessageChannelMessageAssociationsParamsDepth = 1
)

// Defines values for UpdateManyMessageChannelMessageAssociationsParamsDepth.
const (
	UpdateManyMessageChannelMessageAssociationsParamsDepthN0 UpdateManyMessageChannelMessageAssociationsParamsDepth = 0
	UpdateManyMessageChannelMessageAssociationsParamsDepthN1 UpdateManyMessageChannelMessageAssociationsParamsDepth = 1
)

// Defines values for CreateOneMessageChannelMessageAssociationParamsDepth.
const (
	CreateOneMessageChannelMessageAssociationParamsDepthN0 CreateOneMessageChannelMessageAssociationParamsDepth = 0
	CreateOneMessageChannelMessageAssociationParamsDepthN1 CreateOneMessageChannelMessageAssociationParamsDepth = 1
)

// Defines values for FindMessageChannelMessageAssociationDuplicatesParamsDepth.
const (
	FindMessageChannelMessageAssociationDuplicatesParamsDepthN0 FindMessageChannelMessageAssociationDuplicatesParamsDepth = 0
	FindMessageChannelMessageAssociationDuplicatesParamsDepthN1 FindMessageChannelMessageAssociationDuplicatesParamsDepth = 1
)

// Defines values for MergeManyMessageChannelMessageAssociationsParamsDepth.
const (
	MergeManyMessageChannelMessageAssociationsParamsDepthN0 MergeManyMessageChannelMessageAssociationsParamsDepth = 0
	MergeManyMessageChannelMessageAssociationsParamsDepthN1 MergeManyMessageChannelMessageAssociationsParamsDepth = 1
)

// Defines values for FindOneMessageChannelMessageAssociationParamsDepth.
const (
	FindOneMessageChannelMessageAssociationParamsDepthN0 FindOneMessageChannelMessageAssociationParamsDepth = 0
	FindOneMessageChannelMessageAssociationParamsDepthN1 FindOneMessageChannelMessageAssociationParamsDepth = 1
)

// Defines values for UpdateOneMessageChannelMessageAssociationParamsDepth.
const (
	UpdateOneMessageChannelMessageAssociationParamsDepthN0 UpdateOneMessageChannelMessageAssociationParamsDepth = 0
	UpdateOneMessageChannelMessageAssociationParamsDepthN1 UpdateOneMessageChannelMessageAssociationParamsDepth = 1
)

// Defines values for FindManyMessageChannelsParamsDepth.
const (
	FindManyMessageChannelsParamsDepthN0 FindManyMessageChannelsParamsDepth = 0
	FindManyMessageChannelsParamsDepthN1 FindManyMessageChannelsParamsDepth = 1
)

// Defines values for UpdateManyMessageChannelsParamsDepth.
const (
	UpdateManyMessageChannelsParamsDepthN0 UpdateManyMessageChannelsParamsDepth = 0
	UpdateManyMessageChannelsParamsDepthN1 UpdateManyMessageChannelsParamsDepth = 1
)

// Defines values for CreateOneMessageChannelParamsDepth.
const (
	CreateOneMessageChannelParamsDepthN0 CreateOneMessageChannelParamsDepth = 0
	CreateOneMessageChannelParamsDepthN1 CreateOneMessageChannelParamsDepth = 1
)

// Defines values for FindMessageChannelDuplicatesParamsDepth.
const (
	FindMessageChannelDuplicatesParamsDepthN0 FindMessageChannelDuplicatesParamsDepth = 0
	FindMessageChannelDuplicatesParamsDepthN1 FindMessageChannelDuplicatesParamsDepth = 1
)

// Defines values for MergeManyMessageChannelsParamsDepth.
const (
	MergeManyMessageChannelsParamsDepthN0 MergeManyMessageChannelsParamsDepth = 0
	MergeManyMessageChannelsParamsDepthN1 MergeManyMessageChannelsParamsDepth = 1
)

// Defines values for FindOneMessageChannelParamsDepth.
const (
	FindOneMessageChannelParamsDepthN0 FindOneMessageChannelParamsDepth = 0
	FindOneMessageChannelParamsDepthN1 FindOneMessageChannelParamsDepth = 1
)

// Defines values for UpdateOneMessageChannelParamsDepth.
const (
	UpdateOneMessageChannelParamsDepthN0 UpdateOneMessageChannelParamsDepth = 0
	UpdateOneMessageChannelParamsDepthN1 UpdateOneMessageChannelParamsDepth = 1
)

// Defines values for FindManyMessageFoldersParamsDepth.
const (
	FindManyMessageFoldersParamsDepthN0 FindManyMessageFoldersParamsDepth = 0
	FindManyMessageFoldersParamsDepthN1 FindManyMessageFoldersParamsDepth = 1
)

// Defines values for UpdateManyMessageFoldersParamsDepth.
const (
	UpdateManyMessageFoldersParamsDepthN0 UpdateManyMessageFoldersParamsDepth = 0
	UpdateManyMessageFoldersParamsDepthN1 UpdateManyMessageFoldersParamsDepth = 1
)

// Defines values for CreateOneMessageFolderParamsDepth.
const (
	CreateOneMessageFolderParamsDepthN0 CreateOneMessageFolderParamsDepth = 0
	CreateOneMessageFolderParamsDepthN1 CreateOneMessageFolderParamsDepth = 1
)

// Defines values for FindMessageFolderDuplicatesParamsDepth.
const (
	FindMessageFolderDuplicatesParamsDepthN0 FindMessageFolderDuplicatesParamsDepth = 0
	FindMessageFolderDuplicatesParamsDepthN1 FindMessageFolderDuplicatesParamsDepth = 1
)

// Defines values for MergeManyMessageFoldersParamsDepth.
const (
	MergeManyMessageFoldersParamsDepthN0 MergeManyMessageFoldersParamsDepth = 0
	MergeManyMessageFoldersParamsDepthN1 MergeManyMessageFoldersParamsDepth = 1
)

// Defines values for FindOneMessageFolderParamsDepth.
const (
	FindOneMessageFolderParamsDepthN0 FindOneMessageFolderParamsDepth = 0
	FindOneMessageFolderParamsDepthN1 FindOneMessageFolderParamsDepth = 1
)

// Defines values for UpdateOneMessageFolderParamsDepth.
const (
	UpdateOneMessageFolderParamsDepthN0 UpdateOneMessageFolderParamsDepth = 0
	UpdateOneMessageFolderParamsDepthN1 UpdateOneMessageFolderParamsDepth = 1
)

// Defines values for FindManyMessageParticipantsParamsDepth.
const (
	FindManyMessageParticipantsParamsDepthN0 FindManyMessageParticipantsParamsDepth = 0
	FindManyMessageParticipantsParamsDepthN1 FindManyMessageParticipantsParamsDepth = 1
)

// Defines values for UpdateManyMessageParticipantsParamsDepth.
const (
	UpdateManyMessageParticipantsParamsDepthN0 UpdateManyMessageParticipantsParamsDepth = 0
	UpdateManyMessageParticipantsParamsDepthN1 UpdateManyMessageParticipantsParamsDepth = 1
)

// Defines values for CreateOneMessageParticipantParamsDepth.
const (
	CreateOneMessageParticipantParamsDepthN0 CreateOneMessageParticipantParamsDepth = 0
	CreateOneMessageParticipantParamsDepthN1 CreateOneMessageParticipantParamsDepth = 1
)

// Defines values for FindMessageParticipantDuplicatesParamsDepth.
const (
	FindMessageParticipantDuplicatesParamsDepthN0 FindMessageParticipantDuplicatesParamsDepth = 0
	FindMessageParticipantDuplicatesParamsDepthN1 FindMessageParticipantDuplicatesParamsDepth = 1
)

// Defines values for MergeManyMessageParticipantsParamsDepth.
const (
	MergeManyMessageParticipantsParamsDepthN0 MergeManyMessageParticipantsParamsDepth = 0
	MergeManyMessageParticipantsParamsDepthN1 MergeManyMessageParticipantsParamsDepth = 1
)

// Defines values for FindOneMessageParticipantParamsDepth.
const (
	FindOneMessageParticipantParamsDepthN0 FindOneMessageParticipantParamsDepth = 0
	FindOneMessageParticipantParamsDepthN1 FindOneMessageParticipantParamsDepth = 1
)

// Defines values for UpdateOneMessageParticipantParamsDepth.
const (
	UpdateOneMessageParticipantParamsDepthN0 UpdateOneMessageParticipantParamsDepth = 0
	UpdateOneMessageParticipantParamsDepthN1 UpdateOneMessageParticipantParamsDepth = 1
)

// Defines values for FindManyMessageThreadsParamsDepth.
const (
	FindManyMessageThreadsParamsDepthN0 FindManyMessageThreadsParamsDepth = 0
	FindManyMessageThreadsParamsDepthN1 FindManyMessageThreadsParamsDepth = 1
)

// Defines values for UpdateManyMessageThreadsParamsDepth.
const (
	UpdateManyMessageThreadsParamsDepthN0 UpdateManyMessageThreadsParamsDepth = 0
	UpdateManyMessageThreadsParamsDepthN1 UpdateManyMessageThreadsParamsDepth = 1
)

// Defines values for CreateOneMessageThreadParamsDepth.
const (
	CreateOneMessageThreadParamsDepthN0 CreateOneMessageThreadParamsDepth = 0
	CreateOneMessageThreadParamsDepthN1 CreateOneMessageThreadParamsDepth = 1
)

// Defines values for FindMessageThreadDuplicatesParamsDepth.
const (
	FindMessageThreadDuplicatesParamsDepthN0 FindMessageThreadDuplicatesParamsDepth = 0
	FindMessageThreadDuplicatesParamsDepthN1 FindMessageThreadDuplicatesParamsDepth = 1
)

// Defines values for MergeManyMessageThreadsParamsDepth.
const (
	MergeManyMessageThreadsParamsDepthN0 MergeManyMessageThreadsParamsDepth = 0
	MergeManyMessageThreadsParamsDepthN1 MergeManyMessageThreadsParamsDepth = 1
)

// Defines values for FindOneMessageThreadParamsDepth.
const (
	FindOneMessageThreadParamsDepthN0 FindOneMessageThreadParamsDepth = 0
	FindOneMessageThreadParamsDepthN1 FindOneMessageThreadParamsDepth = 1
)

// Defines values for UpdateOneMessageThreadParamsDepth.
const (
	UpdateOneMessageThreadParamsDepthN0 UpdateOneMessageThreadParamsDepth = 0
	UpdateOneMessageThreadParamsDepthN1 UpdateOneMessageThreadParamsDepth = 1
)

// Defines values for FindManyMessagesParamsDepth.
const (
	FindManyMessagesParamsDepthN0 FindManyMessagesParamsDepth = 0
	FindManyMessagesParamsDepthN1 FindManyMessagesParamsDepth = 1
)

// Defines values for UpdateManyMessagesParamsDepth.
const (
	UpdateManyMessagesParamsDepthN0 UpdateManyMessagesParamsDepth = 0
	UpdateManyMessagesParamsDepthN1 UpdateManyMessagesParamsDepth = 1
)

// Defines values for CreateOneMessageParamsDepth.
const (
	CreateOneMessageParamsDepthN0 CreateOneMessageParamsDepth = 0
	CreateOneMessageParamsDepthN1 CreateOneMessageParamsDepth = 1
)

// Defines values for FindMessageDuplicatesParamsDepth.
const (
	FindMessageDuplicatesParamsDepthN0 FindMessageDuplicatesParamsDepth = 0
	FindMessageDuplicatesParamsDepthN1 FindMessageDuplicatesParamsDepth = 1
)

// Defines values for MergeManyMessagesParamsDepth.
const (
	MergeManyMessagesParamsDepthN0 MergeManyMessagesParamsDepth = 0
	MergeManyMessagesParamsDepthN1 MergeManyMessagesParamsDepth = 1
)

// Defines values for FindOneMessageParamsDepth.
const (
	FindOneMessageParamsDepthN0 FindOneMessageParamsDepth = 0
	FindOneMessageParamsDepthN1 FindOneMessageParamsDepth = 1
)

// Defines values for UpdateOneMessageParamsDepth.
const (
	UpdateOneMessageParamsDepthN0 UpdateOneMessageParamsDepth = 0
	UpdateOneMessageParamsDepthN1 UpdateOneMessageParamsDepth = 1
)

// Defines values for FindManyNoteTargetsParamsDepth.
const (
	FindManyNoteTargetsParamsDepthN0 FindManyNoteTargetsParamsDepth = 0
	FindManyNoteTargetsParamsDepthN1 FindManyNoteTargetsParamsDepth = 1
)

// Defines values for UpdateManyNoteTargetsParamsDepth.
const (
	UpdateManyNoteTargetsParamsDepthN0 UpdateManyNoteTargetsParamsDepth = 0
	UpdateManyNoteTargetsParamsDepthN1 UpdateManyNoteTargetsParamsDepth = 1
)

// Defines values for CreateOneNoteTargetParamsDepth.
const (
	CreateOneNoteTargetParamsDepthN0 CreateOneNoteTargetParamsDepth = 0
	CreateOneNoteTargetParamsDepthN1 CreateOneNoteTargetParamsDepth = 1
)

// Defines values for FindNoteTargetDuplicatesParamsDepth.
const (
	FindNoteTargetDuplicatesParamsDepthN0 FindNoteTargetDuplicatesParamsDepth = 0
	FindNoteTargetDuplicatesParamsDepthN1 FindNoteTargetDuplicatesParamsDepth = 1
)

// Defines values for MergeManyNoteTargetsParamsDepth.
const (
	MergeManyNoteTargetsParamsDepthN0 MergeManyNoteTargetsParamsDepth = 0
	MergeManyNoteTargetsParamsDepthN1 MergeManyNoteTargetsParamsDepth = 1
)

// Defines values for FindOneNoteTargetParamsDepth.
const (
	FindOneNoteTargetParamsDepthN0 FindOneNoteTargetParamsDepth = 0
	FindOneNoteTargetParamsDepthN1 FindOneNoteTargetParamsDepth = 1
)

// Defines values for UpdateOneNoteTargetParamsDepth.
const (
	UpdateOneNoteTargetParamsDepthN0 UpdateOneNoteTargetParamsDepth = 0
	UpdateOneNoteTargetParamsDepthN1 UpdateOneNoteTargetParamsDepth = 1
)

// Defines values for FindManyNotesParamsDepth.
const (
	FindManyNotesParamsDepthN0 FindManyNotesParamsDepth = 0
	FindManyNotesParamsDepthN1 FindManyNotesParamsDepth = 1
)

// Defines values for UpdateManyNotesParamsDepth.
const (
	UpdateManyNotesParamsDepthN0 UpdateManyNotesParamsDepth = 0
	UpdateManyNotesParamsDepthN1 UpdateManyNotesParamsDepth = 1
)

// Defines values for CreateOneNoteParamsDepth.
const (
	CreateOneNoteParamsDepthN0 CreateOneNoteParamsDepth = 0
	CreateOneNoteParamsDepthN1 CreateOneNoteParamsDepth = 1
)

// Defines values for FindNoteDuplicatesParamsDepth.
const (
	FindNoteDuplicatesParamsDepthN0 FindNoteDuplicatesParamsDepth = 0
	FindNoteDuplicatesParamsDepthN1 FindNoteDuplicatesParamsDepth = 1
)

// Defines values for MergeManyNotesParamsDepth.
const (
	MergeManyNotesParamsDepthN0 MergeManyNotesParamsDepth = 0
	MergeManyNotesParamsDepthN1 MergeManyNotesParamsDepth = 1
)

// Defines values for FindOneNoteParamsDepth.
const (
	FindOneNoteParamsDepthN0 FindOneNoteParamsDepth = 0
	FindOneNoteParamsDepthN1 FindOneNoteParamsDepth = 1
)

// Defines values for UpdateOneNoteParamsDepth.
const (
	UpdateOneNoteParamsDepthN0 UpdateOneNoteParamsDepth = 0
	UpdateOneNoteParamsDepthN1 UpdateOneNoteParamsDepth = 1
)

// Defines values for FindManyOpportunitiesParamsDepth.
const (
	FindManyOpportunitiesParamsDepthN0 FindManyOpportunitiesParamsDepth = 0
	FindManyOpportunitiesParamsDepthN1 FindManyOpportunitiesParamsDepth = 1
)

// Defines values for UpdateManyOpportunitiesParamsDepth.
const (
	UpdateManyOpportunitiesParamsDepthN0 UpdateManyOpportunitiesParamsDepth = 0
	UpdateManyOpportunitiesParamsDepthN1 UpdateManyOpportunitiesParamsDepth = 1
)

// Defines values for CreateOneOpportunityParamsDepth.
const (
	CreateOneOpportunityParamsDepthN0 CreateOneOpportunityParamsDepth = 0
	CreateOneOpportunityParamsDepthN1 CreateOneOpportunityParamsDepth = 1
)

// Defines values for FindOpportunityDuplicatesParamsDepth.
const (
	FindOpportunityDuplicatesParamsDepthN0 FindOpportunityDuplicatesParamsDepth = 0
	FindOpportunityDuplicatesParamsDepthN1 FindOpportunityDuplicatesParamsDepth = 1
)

// Defines values for MergeManyOpportunitiesParamsDepth.
const (
	MergeManyOpportunitiesParamsDepthN0 MergeManyOpportunitiesParamsDepth = 0
	MergeManyOpportunitiesParamsDepthN1 MergeManyOpportunitiesParamsDepth = 1
)

// Defines values for FindOneOpportunityParamsDepth.
const (
	FindOneOpportunityParamsDepthN0 FindOneOpportunityParamsDepth = 0
	FindOneOpportunityParamsDepthN1 FindOneOpportunityParamsDepth = 1
)

// Defines values for UpdateOneOpportunityParamsDepth.
const (
	UpdateOneOpportunityParamsDepthN0 UpdateOneOpportunityParamsDepth = 0
	UpdateOneOpportunityParamsDepthN1 UpdateOneOpportunityParamsDepth = 1
)

// Defines values for FindManyPeopleParamsDepth.
const (
	FindManyPeopleParamsDepthN0 FindManyPeopleParamsDepth = 0
	FindManyPeopleParamsDepthN1 FindManyPeopleParamsDepth = 1
)

// Defines values for UpdateManyPeopleParamsDepth.
const (
	UpdateManyPeopleParamsDepthN0 UpdateManyPeopleParamsDepth = 0
	UpdateManyPeopleParamsDepthN1 UpdateManyPeopleParamsDepth = 1
)

// Defines values for CreateOnePersonParamsDepth.
const (
	CreateOnePersonParamsDepthN0 CreateOnePersonParamsDepth = 0
	CreateOnePersonParamsDepthN1 CreateOnePersonParamsDepth = 1
)

// Defines values for FindPersonDuplicatesParamsDepth.
const (
	FindPersonDuplicatesParamsDepthN0 FindPersonDuplicatesParamsDepth = 0
	FindPersonDuplicatesParamsDepthN1 FindPersonDuplicatesParamsDepth = 1
)

// Defines values for MergeManyPeopleParamsDepth.
const (
	MergeManyPeopleParamsDepthN0 MergeManyPeopleParamsDepth = 0
	MergeManyPeopleParamsDepthN1 MergeManyPeopleParamsDepth = 1
)

// Defines values for FindOnePersonParamsDepth.
const (
	FindOnePersonParamsDepthN0 FindOnePersonParamsDepth = 0
	FindOnePersonParamsDepthN1 FindOnePersonParamsDepth = 1
)

// Defines values for UpdateOnePersonParamsDepth.
const (
	UpdateOnePersonParamsDepthN0 UpdateOnePersonParamsDepth = 0
	UpdateOnePersonParamsDepthN1 UpdateOnePersonParamsDepth = 1
)

// Defines values for RestoreManyAttachmentsParamsDepth.
const (
	RestoreManyAttachmentsParamsDepthN0 RestoreManyAttachmentsParamsDepth = 0
	RestoreManyAttachmentsParamsDepthN1 RestoreManyAttachmentsParamsDepth = 1
)

// Defines values for RestoreOneAttachmentParamsDepth.
const (
	RestoreOneAttachmentParamsDepthN0 RestoreOneAttachmentParamsDepth = 0
	RestoreOneAttachmentParamsDepthN1 RestoreOneAttachmentParamsDepth = 1
)

// Defines values for RestoreManyBlocklistsParamsDepth.
const (
	RestoreManyBlocklistsParamsDepthN0 RestoreManyBlocklistsParamsDepth = 0
	RestoreManyBlocklistsParamsDepthN1 RestoreManyBlocklistsParamsDepth = 1
)

// Defines values for RestoreOneBlocklistParamsDepth.
const (
	RestoreOneBlocklistParamsDepthN0 RestoreOneBlocklistParamsDepth = 0
	RestoreOneBlocklistParamsDepthN1 RestoreOneBlocklistParamsDepth = 1
)

// Defines values for RestoreManyCalendarChannelEventAssociationsParamsDepth.
const (
	RestoreManyCalendarChannelEventAssociationsParamsDepthN0 RestoreManyCalendarChannelEventAssociationsParamsDepth = 0
	RestoreManyCalendarChannelEventAssociationsParamsDepthN1 RestoreManyCalendarChannelEventAssociationsParamsDepth = 1
)

// Defines values for RestoreOneCalendarChannelEventAssociationParamsDepth.
const (
	RestoreOneCalendarChannelEventAssociationParamsDepthN0 RestoreOneCalendarChannelEventAssociationParamsDepth = 0
	RestoreOneCalendarChannelEventAssociationParamsDepthN1 RestoreOneCalendarChannelEventAssociationParamsDepth = 1
)

// Defines values for RestoreManyCalendarChannelsParamsDepth.
const (
	RestoreManyCalendarChannelsParamsDepthN0 RestoreManyCalendarChannelsParamsDepth = 0
	RestoreManyCalendarChannelsParamsDepthN1 RestoreManyCalendarChannelsParamsDepth = 1
)

// Defines values for RestoreOneCalendarChannelParamsDepth.
const (
	RestoreOneCalendarChannelParamsDepthN0 RestoreOneCalendarChannelParamsDepth = 0
	RestoreOneCalendarChannelParamsDepthN1 RestoreOneCalendarChannelParamsDepth = 1
)

// Defines values for RestoreManyCalendarEventParticipantsParamsDepth.
const (
	RestoreManyCalendarEventParticipantsParamsDepthN0 RestoreManyCalendarEventParticipantsParamsDepth = 0
	RestoreManyCalendarEventParticipantsParamsDepthN1 RestoreManyCalendarEventParticipantsParamsDepth = 1
)

// Defines values for RestoreOneCalendarEventParticipantParamsDepth.
const (
	RestoreOneCalendarEventParticipantParamsDepthN0 RestoreOneCalendarEventParticipantParamsDepth = 0
	RestoreOneCalendarEventParticipantParamsDepthN1 RestoreOneCalendarEventParticipantParamsDepth = 1
)

// Defines values for RestoreManyCalendarEventsParamsDepth.
const (
	RestoreManyCalendarEventsParamsDepthN0 RestoreManyCalendarEventsParamsDepth = 0
	RestoreManyCalendarEventsParamsDepthN1 RestoreManyCalendarEventsParamsDepth = 1
)

// Defines values for RestoreOneCalendarEventParamsDepth.
const (
	RestoreOneCalendarEventParamsDepthN0 RestoreOneCalendarEventParamsDepth = 0
	RestoreOneCalendarEventParamsDepthN1 RestoreOneCalendarEventParamsDepth = 1
)

// Defines values for RestoreManyCompaniesParamsDepth.
const (
	RestoreManyCompaniesParamsDepthN0 RestoreManyCompaniesParamsDepth = 0
	RestoreManyCompaniesParamsDepthN1 RestoreManyCompaniesParamsDepth = 1
)

// Defines values for RestoreOneCompanyParamsDepth.
const (
	RestoreOneCompanyParamsDepthN0 RestoreOneCompanyParamsDepth = 0
	RestoreOneCompanyParamsDepthN1 RestoreOneCompanyParamsDepth = 1
)

// Defines values for RestoreManyConnectedAccountsParamsDepth.
const (
	RestoreManyConnectedAccountsParamsDepthN0 RestoreManyConnectedAccountsParamsDepth = 0
	RestoreManyConnectedAccountsParamsDepthN1 RestoreManyConnectedAccountsParamsDepth = 1
)

// Defines values for RestoreOneConnectedAccountParamsDepth.
const (
	RestoreOneConnectedAccountParamsDepthN0 RestoreOneConnectedAccountParamsDepth = 0
	RestoreOneConnectedAccountParamsDepthN1 RestoreOneConnectedAccountParamsDepth = 1
)

// Defines values for RestoreManyDashboardsParamsDepth.
const (
	RestoreManyDashboardsParamsDepthN0 RestoreManyDashboardsParamsDepth = 0
	RestoreManyDashboardsParamsDepthN1 RestoreManyDashboardsParamsDepth = 1
)

// Defines values for RestoreOneDashboardParamsDepth.
const (
	RestoreOneDashboardParamsDepthN0 RestoreOneDashboardParamsDepth = 0
	RestoreOneDashboardParamsDepthN1 RestoreOneDashboardParamsDepth = 1
)

// Defines values for RestoreManyFavoriteFoldersParamsDepth.
const (
	RestoreManyFavoriteFoldersParamsDepthN0 RestoreManyFavoriteFoldersParamsDepth = 0
	RestoreManyFavoriteFoldersParamsDepthN1 RestoreManyFavoriteFoldersParamsDepth = 1
)

// Defines values for RestoreOneFavoriteFolderParamsDepth.
const (
	RestoreOneFavoriteFolderParamsDepthN0 RestoreOneFavoriteFolderParamsDepth = 0
	RestoreOneFavoriteFolderParamsDepthN1 RestoreOneFavoriteFolderParamsDepth = 1
)

// Defines values for RestoreManyFavoritesParamsDepth.
const (
	RestoreManyFavoritesParamsDepthN0 RestoreManyFavoritesParamsDepth = 0
	RestoreManyFavoritesParamsDepthN1 RestoreManyFavoritesParamsDepth = 1
)

// Defines values for RestoreOneFavoriteParamsDepth.
const (
	RestoreOneFavoriteParamsDepthN0 RestoreOneFavoriteParamsDepth = 0
	RestoreOneFavoriteParamsDepthN1 RestoreOneFavoriteParamsDepth = 1
)

// Defines values for RestoreManyMessageChannelMessageAssociationsParamsDepth.
const (
	RestoreManyMessageChannelMessageAssociationsParamsDepthN0 RestoreManyMessageChannelMessageAssociationsParamsDepth = 0
	RestoreManyMessageChannelMessageAssociationsParamsDepthN1 RestoreManyMessageChannelMessageAssociationsParamsDepth = 1
)

// Defines values for RestoreOneMessageChannelMessageAssociationParamsDepth.
const (
	RestoreOneMessageChannelMessageAssociationParamsDepthN0 RestoreOneMessageChannelMessageAssociationParamsDepth = 0
	RestoreOneMessageChannelMessageAssociationParamsDepthN1 RestoreOneMessageChannelMessageAssociationParamsDepth = 1
)

// Defines values for RestoreManyMessageChannelsParamsDepth.
const (
	RestoreManyMessageChannelsParamsDepthN0 RestoreManyMessageChannelsParamsDepth = 0
	RestoreManyMessageChannelsParamsDepthN1 RestoreManyMessageChannelsParamsDepth = 1
)

// Defines values for RestoreOneMessageChannelParamsDepth.
const (
	RestoreOneMessageChannelParamsDepthN0 RestoreOneMessageChannelParamsDepth = 0
	RestoreOneMessageChannelParamsDepthN1 RestoreOneMessageChannelParamsDepth = 1
)

// Defines values for RestoreManyMessageFoldersParamsDepth.
const (
	RestoreManyMessageFoldersParamsDepthN0 RestoreManyMessageFoldersParamsDepth = 0
	RestoreManyMessageFoldersParamsDepthN1 RestoreManyMessageFoldersParamsDepth = 1
)

// Defines values for RestoreOneMessageFolderParamsDepth.
const (
	RestoreOneMessageFolderParamsDepthN0 RestoreOneMessageFolderParamsDepth = 0
	RestoreOneMessageFolderParamsDepthN1 RestoreOneMessageFolderParamsDepth = 1
)

// Defines values for RestoreManyMessageParticipantsParamsDepth.
const (
	RestoreManyMessageParticipantsParamsDepthN0 RestoreManyMessageParticipantsParamsDepth = 0
	RestoreManyMessageParticipantsParamsDepthN1 RestoreManyMessageParticipantsParamsDepth = 1
)

// Defines values for RestoreOneMessageParticipantParamsDepth.
const (
	RestoreOneMessageParticipantParamsDepthN0 RestoreOneMessageParticipantParamsDepth = 0
	RestoreOneMessageParticipantParamsDepthN1 RestoreOneMessageParticipantParamsDepth = 1
)

// Defines values for RestoreManyMessageThreadsParamsDepth.
const (
	RestoreManyMessageThreadsParamsDepthN0 RestoreManyMessageThreadsParamsDepth = 0
	RestoreManyMessageThreadsParamsDepthN1 RestoreManyMessageThreadsParamsDepth = 1
)

// Defines values for RestoreOneMessageThreadParamsDepth.
const (
	RestoreOneMessageThreadParamsDepthN0 RestoreOneMessageThreadParamsDepth = 0
	RestoreOneMessageThreadParamsDepthN1 RestoreOneMessageThreadParamsDepth = 1
)

// Defines values for RestoreManyMessagesParamsDepth.
const (
	RestoreManyMessagesParamsDepthN0 RestoreManyMessagesParamsDepth = 0
	RestoreManyMessagesParamsDepthN1 RestoreManyMessagesParamsDepth = 1
)

// Defines values for RestoreOneMessageParamsDepth.
const (
	RestoreOneMessageParamsDepthN0 RestoreOneMessageParamsDepth = 0
	RestoreOneMessageParamsDepthN1 RestoreOneMessageParamsDepth = 1
)

// Defines values for RestoreManyNoteTargetsParamsDepth.
const (
	RestoreManyNoteTargetsParamsDepthN0 RestoreManyNoteTargetsParamsDepth = 0
	RestoreManyNoteTargetsParamsDepthN1 RestoreManyNoteTargetsParamsDepth = 1
)

// Defines values for RestoreOneNoteTargetParamsDepth.
const (
	RestoreOneNoteTargetParamsDepthN0 RestoreOneNoteTargetParamsDepth = 0
	RestoreOneNoteTargetParamsDepthN1 RestoreOneNoteTargetParamsDepth = 1
)

// Defines values for RestoreManyNotesParamsDepth.
const (
	RestoreManyNotesParamsDepthN0 RestoreManyNotesParamsDepth = 0
	RestoreManyNotesParamsDepthN1 RestoreManyNotesParamsDepth = 1
)

// Defines values for RestoreOneNoteParamsDepth.
const (
	RestoreOneNoteParamsDepthN0 RestoreOneNoteParamsDepth = 0
	RestoreOneNoteParamsDepthN1 RestoreOneNoteParamsDepth = 1
)

// Defines values for RestoreManyOpportunitiesParamsDepth.
const (
	RestoreManyOpportunitiesParamsDepthN0 RestoreManyOpportunitiesParamsDepth = 0
	RestoreManyOpportunitiesParamsDepthN1 RestoreManyOpportunitiesParamsDepth = 1
)

// Defines values for RestoreOneOpportunityParamsDepth.
const (
	RestoreOneOpportunityParamsDepthN0 RestoreOneOpportunityParamsDepth = 0
	RestoreOneOpportunityParamsDepthN1 RestoreOneOpportunityParamsDepth = 1
)

// Defines values for RestoreManyPeopleParamsDepth.
const (
	RestoreManyPeopleParamsDepthN0 RestoreManyPeopleParamsDepth = 0
	RestoreManyPeopleParamsDepthN1 RestoreManyPeopleParamsDepth = 1
)

// Defines values for RestoreOnePersonParamsDepth.
const (
	RestoreOnePersonParamsDepthN0 RestoreOnePersonParamsDepth = 0
	RestoreOnePersonParamsDepthN1 RestoreOnePersonParamsDepth = 1
)

// Defines values for RestoreManyTaskTargetsParamsDepth.
const (
	RestoreManyTaskTargetsParamsDepthN0 RestoreManyTaskTargetsParamsDepth = 0
	RestoreManyTaskTargetsParamsDepthN1 RestoreManyTaskTargetsParamsDepth = 1
)

// Defines values for RestoreOneTaskTargetParamsDepth.
const (
	RestoreOneTaskTargetParamsDepthN0 RestoreOneTaskTargetParamsDepth = 0
	RestoreOneTaskTargetParamsDepthN1 RestoreOneTaskTargetParamsDepth = 1
)

// Defines values for RestoreManyTasksParamsDepth.
const (
	RestoreManyTasksParamsDepthN0 RestoreManyTasksParamsDepth = 0
	RestoreManyTasksParamsDepthN1 RestoreManyTasksParamsDepth = 1
)

// Defines values for RestoreOneTaskParamsDepth.
const (
	RestoreOneTaskParamsDepthN0 RestoreOneTaskParamsDepth = 0
	RestoreOneTaskParamsDepthN1 RestoreOneTaskParamsDepth = 1
)

// Defines values for RestoreManyTimelineActivitiesParamsDepth.
const (
	RestoreManyTimelineActivitiesParamsDepthN0 RestoreManyTimelineActivitiesParamsDepth = 0
	RestoreManyTimelineActivitiesParamsDepthN1 RestoreManyTimelineActivitiesParamsDepth = 1
)

// Defines values for RestoreOneTimelineActivityParamsDepth.
const (
	RestoreOneTimelineActivityParamsDepthN0 RestoreOneTimelineActivityParamsDepth = 0
	RestoreOneTimelineActivityParamsDepthN1 RestoreOneTimelineActivityParamsDepth = 1
)

// Defines values for RestoreManyWorkflowAutomatedTriggersParamsDepth.
const (
	RestoreManyWorkflowAutomatedTriggersParamsDepthN0 RestoreManyWorkflowAutomatedTriggersParamsDepth = 0
	RestoreManyWorkflowAutomatedTriggersParamsDepthN1 RestoreManyWorkflowAutomatedTriggersParamsDepth = 1
)

// Defines values for RestoreOneWorkflowAutomatedTriggerParamsDepth.
const (
	RestoreOneWorkflowAutomatedTriggerParamsDepthN0 RestoreOneWorkflowAutomatedTriggerParamsDepth = 0
	RestoreOneWorkflowAutomatedTriggerParamsDepthN1 RestoreOneWorkflowAutomatedTriggerParamsDepth = 1
)

// Defines values for RestoreManyWorkflowRunsParamsDepth.
const (
	RestoreManyWorkflowRunsParamsDepthN0 RestoreManyWorkflowRunsParamsDepth = 0
	RestoreManyWorkflowRunsParamsDepthN1 RestoreManyWorkflowRunsParamsDepth = 1
)

// Defines values for RestoreOneWorkflowRunParamsDepth.
const (
	RestoreOneWorkflowRunParamsDepthN0 RestoreOneWorkflowRunParamsDepth = 0
	RestoreOneWorkflowRunParamsDepthN1 RestoreOneWorkflowRunParamsDepth = 1
)

// Defines values for RestoreManyWorkflowVersionsParamsDepth.
const (
	RestoreManyWorkflowVersionsParamsDepthN0 RestoreManyWorkflowVersionsParamsDepth = 0
	RestoreManyWorkflowVersionsParamsDepthN1 RestoreManyWorkflowVersionsParamsDepth = 1
)

// Defines values for RestoreOneWorkflowVersionParamsDepth.
const (
	RestoreOneWorkflowVersionParamsDepthN0 RestoreOneWorkflowVersionParamsDepth = 0
	RestoreOneWorkflowVersionParamsDepthN1 RestoreOneWorkflowVersionParamsDepth = 1
)

// Defines values for RestoreManyWorkflowsParamsDepth.
const (
	RestoreManyWorkflowsParamsDepthN0 RestoreManyWorkflowsParamsDepth = 0
	RestoreManyWorkflowsParamsDepthN1 RestoreManyWorkflowsParamsDepth = 1
)

// Defines values for RestoreOneWorkflowParamsDepth.
const (
	RestoreOneWorkflowParamsDepthN0 RestoreOneWorkflowParamsDepth = 0
	RestoreOneWorkflowParamsDepthN1 RestoreOneWorkflowParamsDepth = 1
)

// Defines values for RestoreManyWorkspaceMembersParamsDepth.
const (
	RestoreManyWorkspaceMembersParamsDepthN0 RestoreManyWorkspaceMembersParamsDepth = 0
	RestoreManyWorkspaceMembersParamsDepthN1 RestoreManyWorkspaceMembersParamsDepth = 1
)

// Defines values for RestoreOneWorkspaceMemberParamsDepth.
const (
	RestoreOneWorkspaceMemberParamsDepthN0 RestoreOneWorkspaceMemberParamsDepth = 0
	RestoreOneWorkspaceMemberParamsDepthN1 RestoreOneWorkspaceMemberParamsDepth = 1
)

// Defines values for FindManyTaskTargetsParamsDepth.
const (
	FindManyTaskTargetsParamsDepthN0 FindManyTaskTargetsParamsDepth = 0
	FindManyTaskTargetsParamsDepthN1 FindManyTaskTargetsParamsDepth = 1
)

// Defines values for UpdateManyTaskTargetsParamsDepth.
const (
	UpdateManyTaskTargetsParamsDepthN0 UpdateManyTaskTargetsParamsDepth = 0
	UpdateManyTaskTargetsParamsDepthN1 UpdateManyTaskTargetsParamsDepth = 1
)

// Defines values for CreateOneTaskTargetParamsDepth.
const (
	CreateOneTaskTargetParamsDepthN0 CreateOneTaskTargetParamsDepth = 0
	CreateOneTaskTargetParamsDepthN1 CreateOneTaskTargetParamsDepth = 1
)

// Defines values for FindTaskTargetDuplicatesParamsDepth.
const (
	FindTaskTargetDuplicatesParamsDepthN0 FindTaskTargetDuplicatesParamsDepth = 0
	FindTaskTargetDuplicatesParamsDepthN1 FindTaskTargetDuplicatesParamsDepth = 1
)

// Defines values for MergeManyTaskTargetsParamsDepth.
const (
	MergeManyTaskTargetsParamsDepthN0 MergeManyTaskTargetsParamsDepth = 0
	MergeManyTaskTargetsParamsDepthN1 MergeManyTaskTargetsParamsDepth = 1
)

// Defines values for FindOneTaskTargetParamsDepth.
const (
	FindOneTaskTargetParamsDepthN0 FindOneTaskTargetParamsDepth = 0
	FindOneTaskTargetParamsDepthN1 FindOneTaskTargetParamsDepth = 1
)

// Defines values for UpdateOneTaskTargetParamsDepth.
const (
	UpdateOneTaskTargetParamsDepthN0 UpdateOneTaskTargetParamsDepth = 0
	UpdateOneTaskTargetParamsDepthN1 UpdateOneTaskTargetParamsDepth = 1
)

// Defines values for FindManyTasksParamsDepth.
const (
	FindManyTasksParamsDepthN0 FindManyTasksParamsDepth = 0
	FindManyTasksParamsDepthN1 FindManyTasksParamsDepth = 1
)

// Defines values for UpdateManyTasksParamsDepth.
const (
	UpdateManyTasksParamsDepthN0 UpdateManyTasksParamsDepth = 0
	UpdateManyTasksParamsDepthN1 UpdateManyTasksParamsDepth = 1
)

// Defines values for CreateOneTaskParamsDepth.
const (
	CreateOneTaskParamsDepthN0 CreateOneTaskParamsDepth = 0
	CreateOneTaskParamsDepthN1 CreateOneTaskParamsDepth = 1
)

// Defines values for FindTaskDuplicatesParamsDepth.
const (
	FindTaskDuplicatesParamsDepthN0 FindTaskDuplicatesParamsDepth = 0
	FindTaskDuplicatesParamsDepthN1 FindTaskDuplicatesParamsDepth = 1
)

// Defines values for MergeManyTasksParamsDepth.
const (
	MergeManyTasksParamsDepthN0 MergeManyTasksParamsDepth = 0
	MergeManyTasksParamsDepthN1 MergeManyTasksParamsDepth = 1
)

// Defines values for FindOneTaskParamsDepth.
const (
	FindOneTaskParamsDepthN0 FindOneTaskParamsDepth = 0
	FindOneTaskParamsDepthN1 FindOneTaskParamsDepth = 1
)

// Defines values for UpdateOneTaskParamsDepth.
const (
	UpdateOneTaskParamsDepthN0 UpdateOneTaskParamsDepth = 0
	UpdateOneTaskParamsDepthN1 UpdateOneTaskParamsDepth = 1
)

// Defines values for FindManyTimelineActivitiesParamsDepth.
const (
	FindManyTimelineActivitiesParamsDepthN0 FindManyTimelineActivitiesParamsDepth = 0
	FindManyTimelineActivitiesParamsDepthN1 FindManyTimelineActivitiesParamsDepth = 1
)

// Defines values for UpdateManyTimelineActivitiesParamsDepth.
const (
	UpdateManyTimelineActivitiesParamsDepthN0 UpdateManyTimelineActivitiesParamsDepth = 0
	UpdateManyTimelineActivitiesParamsDepthN1 UpdateManyTimelineActivitiesParamsDepth = 1
)

// Defines values for CreateOneTimelineActivityParamsDepth.
const (
	CreateOneTimelineActivityParamsDepthN0 CreateOneTimelineActivityParamsDepth = 0
	CreateOneTimelineActivityParamsDepthN1 CreateOneTimelineActivityParamsDepth = 1
)

// Defines values for FindTimelineActivityDuplicatesParamsDepth.
const (
	FindTimelineActivityDuplicatesParamsDepthN0 FindTimelineActivityDuplicatesParamsDepth = 0
	FindTimelineActivityDuplicatesParamsDepthN1 FindTimelineActivityDuplicatesParamsDepth = 1
)

// Defines values for MergeManyTimelineActivitiesParamsDepth.
const (
	MergeManyTimelineActivitiesParamsDepthN0 MergeManyTimelineActivitiesParamsDepth = 0
	MergeManyTimelineActivitiesParamsDepthN1 MergeManyTimelineActivitiesParamsDepth = 1
)

// Defines values for FindOneTimelineActivityParamsDepth.
const (
	FindOneTimelineActivityParamsDepthN0 FindOneTimelineActivityParamsDepth = 0
	FindOneTimelineActivityParamsDepthN1 FindOneTimelineActivityParamsDepth = 1
)

// Defines values for UpdateOneTimelineActivityParamsDepth.
const (
	UpdateOneTimelineActivityParamsDepthN0 UpdateOneTimelineActivityParamsDepth = 0
	UpdateOneTimelineActivityParamsDepthN1 UpdateOneTimelineActivityParamsDepth = 1
)

// Defines values for FindManyWorkflowAutomatedTriggersParamsDepth.
const (
	FindManyWorkflowAutomatedTriggersParamsDepthN0 FindManyWorkflowAutomatedTriggersParamsDepth = 0
	FindManyWorkflowAutomatedTriggersParamsDepthN1 FindManyWorkflowAutomatedTriggersParamsDepth = 1
)

// Defines values for UpdateManyWorkflowAutomatedTriggersParamsDepth.
const (
	UpdateManyWorkflowAutomatedTriggersParamsDepthN0 UpdateManyWorkflowAutomatedTriggersParamsDepth = 0
	UpdateManyWorkflowAutomatedTriggersParamsDepthN1 UpdateManyWorkflowAutomatedTriggersParamsDepth = 1
)

// Defines values for CreateOneWorkflowAutomatedTriggerParamsDepth.
const (
	CreateOneWorkflowAutomatedTriggerParamsDepthN0 CreateOneWorkflowAutomatedTriggerParamsDepth = 0
	CreateOneWorkflowAutomatedTriggerParamsDepthN1 CreateOneWorkflowAutomatedTriggerParamsDepth = 1
)

// Defines values for FindWorkflowAutomatedTriggerDuplicatesParamsDepth.
const (
	FindWorkflowAutomatedTriggerDuplicatesParamsDepthN0 FindWorkflowAutomatedTriggerDuplicatesParamsDepth = 0
	FindWorkflowAutomatedTriggerDuplicatesParamsDepthN1 FindWorkflowAutomatedTriggerDuplicatesParamsDepth = 1
)

// Defines values for MergeManyWorkflowAutomatedTriggersParamsDepth.
const (
	MergeManyWorkflowAutomatedTriggersParamsDepthN0 MergeManyWorkflowAutomatedTriggersParamsDepth = 0
	MergeManyWorkflowAutomatedTriggersParamsDepthN1 MergeManyWorkflowAutomatedTriggersParamsDepth = 1
)

// Defines values for FindOneWorkflowAutomatedTriggerParamsDepth.
const (
	FindOneWorkflowAutomatedTriggerParamsDepthN0 FindOneWorkflowAutomatedTriggerParamsDepth = 0
	FindOneWorkflowAutomatedTriggerParamsDepthN1 FindOneWorkflowAutomatedTriggerParamsDepth = 1
)

// Defines values for UpdateOneWorkflowAutomatedTriggerParamsDepth.
const (
	UpdateOneWorkflowAutomatedTriggerParamsDepthN0 UpdateOneWorkflowAutomatedTriggerParamsDepth = 0
	UpdateOneWorkflowAutomatedTriggerParamsDepthN1 UpdateOneWorkflowAutomatedTriggerParamsDepth = 1
)

// Defines values for FindManyWorkflowRunsParamsDepth.
const (
	FindManyWorkflowRunsParamsDepthN0 FindManyWorkflowRunsParamsDepth = 0
	FindManyWorkflowRunsParamsDepthN1 FindManyWorkflowRunsParamsDepth = 1
)

// Defines values for UpdateManyWorkflowRunsParamsDepth.
const (
	UpdateManyWorkflowRunsParamsDepthN0 UpdateManyWorkflowRunsParamsDepth = 0
	UpdateManyWorkflowRunsParamsDepthN1 UpdateManyWorkflowRunsParamsDepth = 1
)

// Defines values for CreateOneWorkflowRunParamsDepth.
const (
	CreateOneWorkflowRunParamsDepthN0 CreateOneWorkflowRunParamsDepth = 0
	CreateOneWorkflowRunParamsDepthN1 CreateOneWorkflowRunParamsDepth = 1
)

// Defines values for FindWorkflowRunDuplicatesParamsDepth.
const (
	FindWorkflowRunDuplicatesParamsDepthN0 FindWorkflowRunDuplicatesParamsDepth = 0
	FindWorkflowRunDuplicatesParamsDepthN1 FindWorkflowRunDuplicatesParamsDepth = 1
)

// Defines values for MergeManyWorkflowRunsParamsDepth.
const (
	MergeManyWorkflowRunsParamsDepthN0 MergeManyWorkflowRunsParamsDepth = 0
	MergeManyWorkflowRunsParamsDepthN1 MergeManyWorkflowRunsParamsDepth = 1
)

// Defines values for FindOneWorkflowRunParamsDepth.
const (
	FindOneWorkflowRunParamsDepthN0 FindOneWorkflowRunParamsDepth = 0
	FindOneWorkflowRunParamsDepthN1 FindOneWorkflowRunParamsDepth = 1
)

// Defines values for UpdateOneWorkflowRunParamsDepth.
const (
	UpdateOneWorkflowRunParamsDepthN0 UpdateOneWorkflowRunParamsDepth = 0
	UpdateOneWorkflowRunParamsDepthN1 UpdateOneWorkflowRunParamsDepth = 1
)

// Defines values for FindManyWorkflowVersionsParamsDepth.
const (
	FindManyWorkflowVersionsParamsDepthN0 FindManyWorkflowVersionsParamsDepth = 0
	FindManyWorkflowVersionsParamsDepthN1 FindManyWorkflowVersionsParamsDepth = 1
)

// Defines values for UpdateManyWorkflowVersionsParamsDepth.
const (
	UpdateManyWorkflowVersionsParamsDepthN0 UpdateManyWorkflowVersionsParamsDepth = 0
	UpdateManyWorkflowVersionsParamsDepthN1 UpdateManyWorkflowVersionsParamsDepth = 1
)

// Defines values for CreateOneWorkflowVersionParamsDepth.
const (
	CreateOneWorkflowVersionParamsDepthN0 CreateOneWorkflowVersionParamsDepth = 0
	CreateOneWorkflowVersionParamsDepthN1 CreateOneWorkflowVersionParamsDepth = 1
)

// Defines values for FindWorkflowVersionDuplicatesParamsDepth.
const (
	FindWorkflowVersionDuplicatesParamsDepthN0 FindWorkflowVersionDuplicatesParamsDepth = 0
	FindWorkflowVersionDuplicatesParamsDepthN1 FindWorkflowVersionDuplicatesParamsDepth = 1
)

// Defines values for MergeManyWorkflowVersionsParamsDepth.
const (
	MergeManyWorkflowVersionsParamsDepthN0 MergeManyWorkflowVersionsParamsDepth = 0
	MergeManyWorkflowVersionsParamsDepthN1 MergeManyWorkflowVersionsParamsDepth = 1
)

// Defines values for FindOneWorkflowVersionParamsDepth.
const (
	FindOneWorkflowVersionParamsDepthN0 FindOneWorkflowVersionParamsDepth = 0
	FindOneWorkflowVersionParamsDepthN1 FindOneWorkflowVersionParamsDepth = 1
)

// Defines values for UpdateOneWorkflowVersionParamsDepth.
const (
	UpdateOneWorkflowVersionParamsDepthN0 UpdateOneWorkflowVersionParamsDepth = 0
	UpdateOneWorkflowVersionParamsDepthN1 UpdateOneWorkflowVersionParamsDepth = 1
)

// Defines values for FindManyWorkflowsParamsDepth.
const (
	FindManyWorkflowsParamsDepthN0 FindManyWorkflowsParamsDepth = 0
	FindManyWorkflowsParamsDepthN1 FindManyWorkflowsParamsDepth = 1
)

// Defines values for UpdateManyWorkflowsParamsDepth.
const (
	UpdateManyWorkflowsParamsDepthN0 UpdateManyWorkflowsParamsDepth = 0
	UpdateManyWorkflowsParamsDepthN1 UpdateManyWorkflowsParamsDepth = 1
)

// Defines values for CreateOneWorkflowParamsDepth.
const (
	CreateOneWorkflowParamsDepthN0 CreateOneWorkflowParamsDepth = 0
	CreateOneWorkflowParamsDepthN1 CreateOneWorkflowParamsDepth = 1
)

// Defines values for FindWorkflowDuplicatesParamsDepth.
const (
	FindWorkflowDuplicatesParamsDepthN0 FindWorkflowDuplicatesParamsDepth = 0
	FindWorkflowDuplicatesParamsDepthN1 FindWorkflowDuplicatesParamsDepth = 1
)

// Defines values for MergeManyWorkflowsParamsDepth.
const (
	MergeManyWorkflowsParamsDepthN0 MergeManyWorkflowsParamsDepth = 0
	MergeManyWorkflowsParamsDepthN1 MergeManyWorkflowsParamsDepth = 1
)

// Defines values for FindOneWorkflowParamsDepth.
const (
	FindOneWorkflowParamsDepthN0 FindOneWorkflowParamsDepth = 0
	FindOneWorkflowParamsDepthN1 FindOneWorkflowParamsDepth = 1
)

// Defines values for UpdateOneWorkflowParamsDepth.
const (
	UpdateOneWorkflowParamsDepthN0 UpdateOneWorkflowParamsDepth = 0
	UpdateOneWorkflowParamsDepthN1 UpdateOneWorkflowParamsDepth = 1
)

// Defines values for FindManyWorkspaceMembersParamsDepth.
const (
	FindManyWorkspaceMembersParamsDepthN0 FindManyWorkspaceMembersParamsDepth = 0
	FindManyWorkspaceMembersParamsDepthN1 FindManyWorkspaceMembersParamsDepth = 1
)

// Defines values for UpdateManyWorkspaceMembersParamsDepth.
const (
	UpdateManyWorkspaceMembersParamsDepthN0 UpdateManyWorkspaceMembersParamsDepth = 0
	UpdateManyWorkspaceMembersParamsDepthN1 UpdateManyWorkspaceMembersParamsDepth = 1
)

// Defines values for CreateOneWorkspaceMemberParamsDepth.
const (
	CreateOneWorkspaceMemberParamsDepthN0 CreateOneWorkspaceMemberParamsDepth = 0
	CreateOneWorkspaceMemberParamsDepthN1 CreateOneWorkspaceMemberParamsDepth = 1
)

// Defines values for FindWorkspaceMemberDuplicatesParamsDepth.
const (
	FindWorkspaceMemberDuplicatesParamsDepthN0 FindWorkspaceMemberDuplicatesParamsDepth = 0
	FindWorkspaceMemberDuplicatesParamsDepthN1 FindWorkspaceMemberDuplicatesParamsDepth = 1
)

// Defines values for MergeManyWorkspaceMembersParamsDepth.
const (
	MergeManyWorkspaceMembersParamsDepthN0 MergeManyWorkspaceMembersParamsDepth = 0
	MergeManyWorkspaceMembersParamsDepthN1 MergeManyWorkspaceMembersParamsDepth = 1
)

// Defines values for FindOneWorkspaceMemberParamsDepth.
const (
	FindOneWorkspaceMemberParamsDepthN0 FindOneWorkspaceMemberParamsDepth = 0
	FindOneWorkspaceMemberParamsDepthN1 FindOneWorkspaceMemberParamsDepth = 1
)

// Defines values for UpdateOneWorkspaceMemberParamsDepth.
const (
	N0 UpdateOneWorkspaceMemberParamsDepth = 0
	N1 UpdateOneWorkspaceMemberParamsDepth = 1
)

// Attachment An attachment
type Attachment struct {
	AuthorId  *openapi_types.UUID `json:"authorId,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *AttachmentCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`
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

// AttachmentCreatedBySource defines model for Attachment.CreatedBy.Source.
type AttachmentCreatedBySource string

// AttachmentFileCategory Attachment file category
type AttachmentFileCategory string

// AttachmentForResponse An attachment
type AttachmentForResponse struct {
	// Author Attachment author (deprecated - use createdBy)
	Author   *AttachmentForResponse_Author `json:"author,omitempty"`
	AuthorId *openapi_types.UUID           `json:"authorId,omitempty"`

	// Company Attachment company
	Company   *AttachmentForResponse_Company `json:"company,omitempty"`
	CompanyId *openapi_types.UUID            `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Name              *string                               `json:"name,omitempty"`
		Source            *AttachmentForResponseCreatedBySource `json:"source,omitempty"`
		WorkspaceMemberId *openapi_types.UUID                   `json:"workspaceMemberId,omitempty"`
	} `json:"createdBy,omitempty"`

	// Dashboard Attachment dashboard
	Dashboard   *AttachmentForResponse_Dashboard `json:"dashboard,omitempty"`
	DashboardId *openapi_types.UUID              `json:"dashboardId,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// FileCategory Attachment file category
	FileCategory *AttachmentForResponseFileCategory `json:"fileCategory,omitempty"`

	// FullPath Attachment full path
	FullPath *string `json:"fullPath,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Name Attachment name
	Name *string `json:"name,omitempty"`

	// Note Attachment note
	Note   *AttachmentForResponse_Note `json:"note,omitempty"`
	NoteId *openapi_types.UUID         `json:"noteId,omitempty"`

	// Opportunity Attachment opportunity
	Opportunity   *AttachmentForResponse_Opportunity `json:"opportunity,omitempty"`
	OpportunityId *openapi_types.UUID                `json:"opportunityId,omitempty"`

	// Person Attachment person
	Person   *AttachmentForResponse_Person `json:"person,omitempty"`
	PersonId *openapi_types.UUID           `json:"personId,omitempty"`

	// Task Attachment task
	Task   *AttachmentForResponse_Task `json:"task,omitempty"`
	TaskId *openapi_types.UUID         `json:"taskId,omitempty"`

	// Type Attachment type (deprecated - use fileCategory)
	Type *string `json:"type,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Workflow Attachment workflow
	Workflow   *AttachmentForResponse_Workflow `json:"workflow,omitempty"`
	WorkflowId *openapi_types.UUID             `json:"workflowId,omitempty"`
}

// AttachmentForResponse_Author Attachment author (deprecated - use createdBy)
type AttachmentForResponse_Author struct {
	union json.RawMessage
}

// AttachmentForResponse_Company Attachment company
type AttachmentForResponse_Company struct {
	union json.RawMessage
}

// AttachmentForResponseCreatedBySource defines model for AttachmentForResponse.CreatedBy.Source.
type AttachmentForResponseCreatedBySource string

// AttachmentForResponse_Dashboard Attachment dashboard
type AttachmentForResponse_Dashboard struct {
	union json.RawMessage
}

// AttachmentForResponseFileCategory Attachment file category
type AttachmentForResponseFileCategory string

// AttachmentForResponse_Note Attachment note
type AttachmentForResponse_Note struct {
	union json.RawMessage
}

// AttachmentForResponse_Opportunity Attachment opportunity
type AttachmentForResponse_Opportunity struct {
	union json.RawMessage
}

// AttachmentForResponse_Person Attachment person
type AttachmentForResponse_Person struct {
	union json.RawMessage
}

// AttachmentForResponse_Task Attachment task
type AttachmentForResponse_Task struct {
	union json.RawMessage
}

// AttachmentForResponse_Workflow Attachment workflow
type AttachmentForResponse_Workflow struct {
	union json.RawMessage
}

// AttachmentForUpdate An attachment
type AttachmentForUpdate struct {
	AuthorId  *openapi_types.UUID `json:"authorId,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *AttachmentForUpdateCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`
	DashboardId *openapi_types.UUID `json:"dashboardId,omitempty"`

	// FileCategory Attachment file category
	FileCategory *AttachmentForUpdateFileCategory `json:"fileCategory,omitempty"`

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

// AttachmentForUpdateCreatedBySource defines model for AttachmentForUpdate.CreatedBy.Source.
type AttachmentForUpdateCreatedBySource string

// AttachmentForUpdateFileCategory Attachment file category
type AttachmentForUpdateFileCategory string

// Blocklist Blocklist
type Blocklist struct {
	// Handle Handle
	Handle            *string             `json:"handle,omitempty"`
	WorkspaceMemberId *openapi_types.UUID `json:"workspaceMemberId,omitempty"`
}

// BlocklistForResponse Blocklist
type BlocklistForResponse struct {
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
	WorkspaceMember   *BlocklistForResponse_WorkspaceMember `json:"workspaceMember,omitempty"`
	WorkspaceMemberId *openapi_types.UUID                   `json:"workspaceMemberId,omitempty"`
}

// BlocklistForResponse_WorkspaceMember WorkspaceMember
type BlocklistForResponse_WorkspaceMember struct {
	union json.RawMessage
}

// BlocklistForUpdate Blocklist
type BlocklistForUpdate struct {
	// Handle Handle
	Handle            *string             `json:"handle,omitempty"`
	WorkspaceMemberId *openapi_types.UUID `json:"workspaceMemberId,omitempty"`
}

// CalendarChannel Calendar Channels
type CalendarChannel struct {
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

// CalendarChannelEventAssociation Calendar Channel Event Associations
type CalendarChannelEventAssociation struct {
	CalendarChannelId *openapi_types.UUID `json:"calendarChannelId,omitempty"`
	CalendarEventId   *openapi_types.UUID `json:"calendarEventId,omitempty"`

	// EventExternalId Event external ID
	EventExternalId *string `json:"eventExternalId,omitempty"`

	// RecurringEventExternalId Recurring Event ID
	RecurringEventExternalId *string `json:"recurringEventExternalId,omitempty"`
}

// CalendarChannelEventAssociationForResponse Calendar Channel Event Associations
type CalendarChannelEventAssociationForResponse struct {
	// CalendarChannel Channel ID
	CalendarChannel   *CalendarChannelEventAssociationForResponse_CalendarChannel `json:"calendarChannel,omitempty"`
	CalendarChannelId *openapi_types.UUID                                         `json:"calendarChannelId,omitempty"`

	// CalendarEvent Event ID
	CalendarEvent   *CalendarChannelEventAssociationForResponse_CalendarEvent `json:"calendarEvent,omitempty"`
	CalendarEventId *openapi_types.UUID                                       `json:"calendarEventId,omitempty"`

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

// CalendarChannelEventAssociationForResponse_CalendarChannel Channel ID
type CalendarChannelEventAssociationForResponse_CalendarChannel struct {
	union json.RawMessage
}

// CalendarChannelEventAssociationForResponse_CalendarEvent Event ID
type CalendarChannelEventAssociationForResponse_CalendarEvent struct {
	union json.RawMessage
}

// CalendarChannelEventAssociationForUpdate Calendar Channel Event Associations
type CalendarChannelEventAssociationForUpdate struct {
	CalendarChannelId *openapi_types.UUID `json:"calendarChannelId,omitempty"`
	CalendarEventId   *openapi_types.UUID `json:"calendarEventId,omitempty"`

	// EventExternalId Event external ID
	EventExternalId *string `json:"eventExternalId,omitempty"`

	// RecurringEventExternalId Recurring Event ID
	RecurringEventExternalId *string `json:"recurringEventExternalId,omitempty"`
}

// CalendarChannelForResponse Calendar Channels
type CalendarChannelForResponse struct {
	// CalendarChannelEventAssociations Calendar Channel Event Associations
	CalendarChannelEventAssociations *[]CalendarChannelEventAssociationForResponse `json:"calendarChannelEventAssociations,omitempty"`

	// ConnectedAccount Connected Account
	ConnectedAccount   *CalendarChannelForResponse_ConnectedAccount `json:"connectedAccount,omitempty"`
	ConnectedAccountId *openapi_types.UUID                          `json:"connectedAccountId,omitempty"`

	// ContactAutoCreationPolicy Automatically create records for people you participated with in an event.
	ContactAutoCreationPolicy *CalendarChannelForResponseContactAutoCreationPolicy `json:"contactAutoCreationPolicy,omitempty"`

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
	SyncStage *CalendarChannelForResponseSyncStage `json:"syncStage,omitempty"`

	// SyncStageStartedAt Sync stage started at
	SyncStageStartedAt *time.Time `json:"syncStageStartedAt,omitempty"`

	// SyncStatus Sync status
	SyncStatus *CalendarChannelForResponseSyncStatus `json:"syncStatus,omitempty"`

	// SyncedAt Last sync date
	SyncedAt *time.Time `json:"syncedAt,omitempty"`

	// ThrottleFailureCount Throttle Failure Count
	ThrottleFailureCount *int `json:"throttleFailureCount,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Visibility Visibility
	Visibility *CalendarChannelForResponseVisibility `json:"visibility,omitempty"`
}

// CalendarChannelForResponse_ConnectedAccount Connected Account
type CalendarChannelForResponse_ConnectedAccount struct {
	union json.RawMessage
}

// CalendarChannelForResponseContactAutoCreationPolicy Automatically create records for people you participated with in an event.
type CalendarChannelForResponseContactAutoCreationPolicy string

// CalendarChannelForResponseSyncStage Sync stage
type CalendarChannelForResponseSyncStage string

// CalendarChannelForResponseSyncStatus Sync status
type CalendarChannelForResponseSyncStatus string

// CalendarChannelForResponseVisibility Visibility
type CalendarChannelForResponseVisibility string

// CalendarChannelForUpdate Calendar Channels
type CalendarChannelForUpdate struct {
	ConnectedAccountId *openapi_types.UUID `json:"connectedAccountId,omitempty"`

	// ContactAutoCreationPolicy Automatically create records for people you participated with in an event.
	ContactAutoCreationPolicy *CalendarChannelForUpdateContactAutoCreationPolicy `json:"contactAutoCreationPolicy,omitempty"`

	// Handle Handle
	Handle *string `json:"handle,omitempty"`

	// IsContactAutoCreationEnabled Is Contact Auto Creation Enabled
	IsContactAutoCreationEnabled *bool `json:"isContactAutoCreationEnabled,omitempty"`

	// IsSyncEnabled Is Sync Enabled
	IsSyncEnabled *bool `json:"isSyncEnabled,omitempty"`

	// SyncCursor Sync Cursor. Used for syncing events from the calendar provider
	SyncCursor *string `json:"syncCursor,omitempty"`

	// SyncStage Sync stage
	SyncStage *CalendarChannelForUpdateSyncStage `json:"syncStage,omitempty"`

	// SyncStageStartedAt Sync stage started at
	SyncStageStartedAt *time.Time `json:"syncStageStartedAt,omitempty"`

	// SyncStatus Sync status
	SyncStatus *CalendarChannelForUpdateSyncStatus `json:"syncStatus,omitempty"`

	// SyncedAt Last sync date
	SyncedAt *time.Time `json:"syncedAt,omitempty"`

	// ThrottleFailureCount Throttle Failure Count
	ThrottleFailureCount *int `json:"throttleFailureCount,omitempty"`

	// Visibility Visibility
	Visibility *CalendarChannelForUpdateVisibility `json:"visibility,omitempty"`
}

// CalendarChannelForUpdateContactAutoCreationPolicy Automatically create records for people you participated with in an event.
type CalendarChannelForUpdateContactAutoCreationPolicy string

// CalendarChannelForUpdateSyncStage Sync stage
type CalendarChannelForUpdateSyncStage string

// CalendarChannelForUpdateSyncStatus Sync status
type CalendarChannelForUpdateSyncStatus string

// CalendarChannelForUpdateVisibility Visibility
type CalendarChannelForUpdateVisibility string

// CalendarEvent Calendar events
type CalendarEvent struct {
	// ConferenceLink Meet Link
	ConferenceLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"conferenceLink,omitempty"`

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
type CalendarEventForResponse struct {
	// CalendarChannelEventAssociations Calendar Channel Event Associations
	CalendarChannelEventAssociations *[]CalendarChannelEventAssociationForResponse `json:"calendarChannelEventAssociations,omitempty"`

	// CalendarEventParticipants Event Participants
	CalendarEventParticipants *[]CalendarEventParticipantForResponse `json:"calendarEventParticipants,omitempty"`

	// ConferenceLink Meet Link
	ConferenceLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"conferenceLink,omitempty"`

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

// CalendarEventForUpdate Calendar events
type CalendarEventForUpdate struct {
	// ConferenceLink Meet Link
	ConferenceLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"conferenceLink,omitempty"`

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

// CalendarEventParticipant Calendar event participants
type CalendarEventParticipant struct {
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
type CalendarEventParticipantForResponse struct {
	// CalendarEvent Event ID
	CalendarEvent   *CalendarEventParticipantForResponse_CalendarEvent `json:"calendarEvent,omitempty"`
	CalendarEventId *openapi_types.UUID                                `json:"calendarEventId,omitempty"`

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
	Person   *CalendarEventParticipantForResponse_Person `json:"person,omitempty"`
	PersonId *openapi_types.UUID                         `json:"personId,omitempty"`

	// ResponseStatus Response Status
	ResponseStatus *CalendarEventParticipantForResponseResponseStatus `json:"responseStatus,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// WorkspaceMember Workspace Member
	WorkspaceMember   *CalendarEventParticipantForResponse_WorkspaceMember `json:"workspaceMember,omitempty"`
	WorkspaceMemberId *openapi_types.UUID                                  `json:"workspaceMemberId,omitempty"`
}

// CalendarEventParticipantForResponse_CalendarEvent Event ID
type CalendarEventParticipantForResponse_CalendarEvent struct {
	union json.RawMessage
}

// CalendarEventParticipantForResponse_Person Person
type CalendarEventParticipantForResponse_Person struct {
	union json.RawMessage
}

// CalendarEventParticipantForResponseResponseStatus Response Status
type CalendarEventParticipantForResponseResponseStatus string

// CalendarEventParticipantForResponse_WorkspaceMember Workspace Member
type CalendarEventParticipantForResponse_WorkspaceMember struct {
	union json.RawMessage
}

// CalendarEventParticipantForUpdate Calendar event participants
type CalendarEventParticipantForUpdate struct {
	CalendarEventId *openapi_types.UUID `json:"calendarEventId,omitempty"`

	// DisplayName Display Name
	DisplayName *string `json:"displayName,omitempty"`

	// Handle Handle
	Handle *string `json:"handle,omitempty"`

	// IsOrganizer Is Organizer
	IsOrganizer *bool               `json:"isOrganizer,omitempty"`
	PersonId    *openapi_types.UUID `json:"personId,omitempty"`

	// ResponseStatus Response Status
	ResponseStatus    *CalendarEventParticipantForUpdateResponseStatus `json:"responseStatus,omitempty"`
	WorkspaceMemberId *openapi_types.UUID                              `json:"workspaceMemberId,omitempty"`
}

// CalendarEventParticipantForUpdateResponseStatus Response Status
type CalendarEventParticipantForUpdateResponseStatus string

// Company A company
type Company struct {
	AccountOwnerId *openapi_types.UUID `json:"accountOwnerId,omitempty"`

	// Address Address of the company
	Address *struct {
		AddressCity     *string  `json:"addressCity,omitempty"`
		AddressCountry  *string  `json:"addressCountry,omitempty"`
		AddressLat      *float32 `json:"addressLat,omitempty"`
		AddressLng      *float32 `json:"addressLng,omitempty"`
		AddressPostcode *string  `json:"addressPostcode,omitempty"`
		AddressState    *string  `json:"addressState,omitempty"`
		AddressStreet1  *string  `json:"addressStreet1,omitempty"`
		AddressStreet2  *string  `json:"addressStreet2,omitempty"`
	} `json:"address,omitempty"`

	// AnnualRecurringRevenue Annual Recurring Revenue: The actual or estimated annual revenue of the company
	AnnualRecurringRevenue *struct {
		AmountMicros *float32 `json:"amountMicros,omitempty"`
		CurrencyCode *string  `json:"currencyCode,omitempty"`
	} `json:"annualRecurringRevenue,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *CompanyCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// DomainName The company website URL. We use this url to fetch the company icon
	DomainName *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"domainName,omitempty"`

	// Employees Number of employees in the company
	Employees *int `json:"employees,omitempty"`

	// IdealCustomerProfile Ideal Customer Profile:  Indicates whether the company is the most suitable and valuable customer for you
	IdealCustomerProfile *bool `json:"idealCustomerProfile,omitempty"`

	// LinkedinLink The company Linkedin account
	LinkedinLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"linkedinLink,omitempty"`

	// Name The company name
	Name *string `json:"name,omitempty"`

	// Position Company record position
	Position *float32 `json:"position,omitempty"`

	// XLink The company Twitter/X account
	XLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"xLink,omitempty"`
}

// CompanyCreatedBySource defines model for Company.CreatedBy.Source.
type CompanyCreatedBySource string

// CompanyForResponse A company
type CompanyForResponse struct {
	// AccountOwner Your team member responsible for managing the company account
	AccountOwner   *CompanyForResponse_AccountOwner `json:"accountOwner,omitempty"`
	AccountOwnerId *openapi_types.UUID              `json:"accountOwnerId,omitempty"`

	// Address Address of the company
	Address *struct {
		AddressCity     *string  `json:"addressCity,omitempty"`
		AddressCountry  *string  `json:"addressCountry,omitempty"`
		AddressLat      *float32 `json:"addressLat,omitempty"`
		AddressLng      *float32 `json:"addressLng,omitempty"`
		AddressPostcode *string  `json:"addressPostcode,omitempty"`
		AddressState    *string  `json:"addressState,omitempty"`
		AddressStreet1  *string  `json:"addressStreet1,omitempty"`
		AddressStreet2  *string  `json:"addressStreet2,omitempty"`
	} `json:"address,omitempty"`

	// AnnualRecurringRevenue Annual Recurring Revenue: The actual or estimated annual revenue of the company
	AnnualRecurringRevenue *struct {
		AmountMicros *float32 `json:"amountMicros,omitempty"`
		CurrencyCode *string  `json:"currencyCode,omitempty"`
	} `json:"annualRecurringRevenue,omitempty"`

	// Attachments Attachments linked to the company
	Attachments *[]AttachmentForResponse `json:"attachments,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Name              *string                            `json:"name,omitempty"`
		Source            *CompanyForResponseCreatedBySource `json:"source,omitempty"`
		WorkspaceMemberId *openapi_types.UUID                `json:"workspaceMemberId,omitempty"`
	} `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// DomainName The company website URL. We use this url to fetch the company icon
	DomainName *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"domainName,omitempty"`

	// Employees Number of employees in the company
	Employees *int `json:"employees,omitempty"`

	// Favorites Favorites linked to the company
	Favorites *[]FavoriteForResponse `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// IdealCustomerProfile Ideal Customer Profile:  Indicates whether the company is the most suitable and valuable customer for you
	IdealCustomerProfile *bool `json:"idealCustomerProfile,omitempty"`

	// LinkedinLink The company Linkedin account
	LinkedinLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"linkedinLink,omitempty"`

	// Name The company name
	Name *string `json:"name,omitempty"`

	// NoteTargets Notes tied to the company
	NoteTargets *[]NoteTargetForResponse `json:"noteTargets,omitempty"`

	// Opportunities Opportunities linked to the company.
	Opportunities *[]OpportunityForResponse `json:"opportunities,omitempty"`

	// People People linked to the company.
	People *[]PersonForResponse `json:"people,omitempty"`

	// Position Company record position
	Position *float32 `json:"position,omitempty"`

	// TaskTargets Tasks tied to the company
	TaskTargets *[]TaskTargetForResponse `json:"taskTargets,omitempty"`

	// TimelineActivities Timeline Activities linked to the company
	TimelineActivities *[]TimelineActivityForResponse `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// XLink The company Twitter/X account
	XLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"xLink,omitempty"`
}

// CompanyForResponse_AccountOwner Your team member responsible for managing the company account
type CompanyForResponse_AccountOwner struct {
	union json.RawMessage
}

// CompanyForResponseCreatedBySource defines model for CompanyForResponse.CreatedBy.Source.
type CompanyForResponseCreatedBySource string

// CompanyForUpdate A company
type CompanyForUpdate struct {
	AccountOwnerId *openapi_types.UUID `json:"accountOwnerId,omitempty"`

	// Address Address of the company
	Address *struct {
		AddressCity     *string  `json:"addressCity,omitempty"`
		AddressCountry  *string  `json:"addressCountry,omitempty"`
		AddressLat      *float32 `json:"addressLat,omitempty"`
		AddressLng      *float32 `json:"addressLng,omitempty"`
		AddressPostcode *string  `json:"addressPostcode,omitempty"`
		AddressState    *string  `json:"addressState,omitempty"`
		AddressStreet1  *string  `json:"addressStreet1,omitempty"`
		AddressStreet2  *string  `json:"addressStreet2,omitempty"`
	} `json:"address,omitempty"`

	// AnnualRecurringRevenue Annual Recurring Revenue: The actual or estimated annual revenue of the company
	AnnualRecurringRevenue *struct {
		AmountMicros *float32 `json:"amountMicros,omitempty"`
		CurrencyCode *string  `json:"currencyCode,omitempty"`
	} `json:"annualRecurringRevenue,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *CompanyForUpdateCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// DomainName The company website URL. We use this url to fetch the company icon
	DomainName *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"domainName,omitempty"`

	// Employees Number of employees in the company
	Employees *int `json:"employees,omitempty"`

	// IdealCustomerProfile Ideal Customer Profile:  Indicates whether the company is the most suitable and valuable customer for you
	IdealCustomerProfile *bool `json:"idealCustomerProfile,omitempty"`

	// LinkedinLink The company Linkedin account
	LinkedinLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"linkedinLink,omitempty"`

	// Name The company name
	Name *string `json:"name,omitempty"`

	// Position Company record position
	Position *float32 `json:"position,omitempty"`

	// XLink The company Twitter/X account
	XLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"xLink,omitempty"`
}

// CompanyForUpdateCreatedBySource defines model for CompanyForUpdate.CreatedBy.Source.
type CompanyForUpdateCreatedBySource string

// ConnectedAccount A connected account
type ConnectedAccount struct {
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

// ConnectedAccountForResponse A connected account
type ConnectedAccountForResponse struct {
	// AccessToken Messaging provider access token
	AccessToken *string `json:"accessToken,omitempty"`

	// AccountOwner Account Owner
	AccountOwner   *ConnectedAccountForResponse_AccountOwner `json:"accountOwner,omitempty"`
	AccountOwnerId *openapi_types.UUID                       `json:"accountOwnerId,omitempty"`

	// AuthFailedAt Auth failed at
	AuthFailedAt *time.Time `json:"authFailedAt,omitempty"`

	// CalendarChannels Calendar Channels
	CalendarChannels *[]CalendarChannelForResponse `json:"calendarChannels,omitempty"`

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
	MessageChannels *[]MessageChannelForResponse `json:"messageChannels,omitempty"`

	// Provider The account provider
	Provider *string `json:"provider,omitempty"`

	// RefreshToken Messaging provider refresh token
	RefreshToken *string `json:"refreshToken,omitempty"`

	// Scopes Scopes
	Scopes *[]string `json:"scopes,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// ConnectedAccountForResponse_AccountOwner Account Owner
type ConnectedAccountForResponse_AccountOwner struct {
	union json.RawMessage
}

// ConnectedAccountForUpdate A connected account
type ConnectedAccountForUpdate struct {
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

// Dashboard A dashboard
type Dashboard struct {
	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *DashboardCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// PageLayoutId Dashboard page layout
	PageLayoutId *openapi_types.UUID `json:"pageLayoutId,omitempty"`

	// Position Dashboard record Position
	Position *float32 `json:"position,omitempty"`

	// Title Dashboard title
	Title *string `json:"title,omitempty"`
}

// DashboardCreatedBySource defines model for Dashboard.CreatedBy.Source.
type DashboardCreatedBySource string

// DashboardForResponse A dashboard
type DashboardForResponse struct {
	// Attachments Attachments linked to the dashboard
	Attachments *[]AttachmentForResponse `json:"attachments,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Name              *string                              `json:"name,omitempty"`
		Source            *DashboardForResponseCreatedBySource `json:"source,omitempty"`
		WorkspaceMemberId *openapi_types.UUID                  `json:"workspaceMemberId,omitempty"`
	} `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the dashboard
	Favorites *[]FavoriteForResponse `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// PageLayoutId Dashboard page layout
	PageLayoutId *openapi_types.UUID `json:"pageLayoutId,omitempty"`

	// Position Dashboard record Position
	Position *float32 `json:"position,omitempty"`

	// TimelineActivities Timeline activities linked to the dashboard
	TimelineActivities *[]TimelineActivityForResponse `json:"timelineActivities,omitempty"`

	// Title Dashboard title
	Title *string `json:"title,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// DashboardForResponseCreatedBySource defines model for DashboardForResponse.CreatedBy.Source.
type DashboardForResponseCreatedBySource string

// DashboardForUpdate A dashboard
type DashboardForUpdate struct {
	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *DashboardForUpdateCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// PageLayoutId Dashboard page layout
	PageLayoutId *openapi_types.UUID `json:"pageLayoutId,omitempty"`

	// Position Dashboard record Position
	Position *float32 `json:"position,omitempty"`

	// Title Dashboard title
	Title *string `json:"title,omitempty"`
}

// DashboardForUpdateCreatedBySource defines model for DashboardForUpdate.CreatedBy.Source.
type DashboardForUpdateCreatedBySource string

// Favorite A favorite that can be accessed from the left menu
type Favorite struct {
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
type FavoriteFolder struct {
	// Name Name of the favorite folder
	Name *string `json:"name,omitempty"`

	// Position Favorite folder position
	Position *int `json:"position,omitempty"`
}

// FavoriteFolderForResponse A Folder of favorites
type FavoriteFolderForResponse struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites in this folder
	Favorites *[]FavoriteForResponse `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Name Name of the favorite folder
	Name *string `json:"name,omitempty"`

	// Position Favorite folder position
	Position *int `json:"position,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// FavoriteFolderForUpdate A Folder of favorites
type FavoriteFolderForUpdate struct {
	// Name Name of the favorite folder
	Name *string `json:"name,omitempty"`

	// Position Favorite folder position
	Position *int `json:"position,omitempty"`
}

// FavoriteForResponse A favorite that can be accessed from the left menu
type FavoriteForResponse struct {
	// Company Favorite company
	Company   *FavoriteForResponse_Company `json:"company,omitempty"`
	CompanyId *openapi_types.UUID          `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Dashboard Favorite dashboard
	Dashboard   *FavoriteForResponse_Dashboard `json:"dashboard,omitempty"`
	DashboardId *openapi_types.UUID            `json:"dashboardId,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// FavoriteFolder The folder this favorite belongs to
	FavoriteFolder   *FavoriteForResponse_FavoriteFolder `json:"favoriteFolder,omitempty"`
	FavoriteFolderId *openapi_types.UUID                 `json:"favoriteFolderId,omitempty"`

	// ForWorkspaceMember Favorite workspace member
	ForWorkspaceMember   *FavoriteForResponse_ForWorkspaceMember `json:"forWorkspaceMember,omitempty"`
	ForWorkspaceMemberId *openapi_types.UUID                     `json:"forWorkspaceMemberId,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Note Favorite note
	Note   *FavoriteForResponse_Note `json:"note,omitempty"`
	NoteId *openapi_types.UUID       `json:"noteId,omitempty"`

	// Opportunity Favorite opportunity
	Opportunity   *FavoriteForResponse_Opportunity `json:"opportunity,omitempty"`
	OpportunityId *openapi_types.UUID              `json:"opportunityId,omitempty"`

	// Person Favorite person
	Person   *FavoriteForResponse_Person `json:"person,omitempty"`
	PersonId *openapi_types.UUID         `json:"personId,omitempty"`

	// Position Favorite position
	Position *int `json:"position,omitempty"`

	// Task Favorite task
	Task   *FavoriteForResponse_Task `json:"task,omitempty"`
	TaskId *openapi_types.UUID       `json:"taskId,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// ViewId ViewId
	ViewId *openapi_types.UUID `json:"viewId,omitempty"`

	// Workflow Favorite workflow
	Workflow   *FavoriteForResponse_Workflow `json:"workflow,omitempty"`
	WorkflowId *openapi_types.UUID           `json:"workflowId,omitempty"`

	// WorkflowRun Favorite workflow run
	WorkflowRun   *FavoriteForResponse_WorkflowRun `json:"workflowRun,omitempty"`
	WorkflowRunId *openapi_types.UUID              `json:"workflowRunId,omitempty"`

	// WorkflowVersion Favorite workflow version
	WorkflowVersion   *FavoriteForResponse_WorkflowVersion `json:"workflowVersion,omitempty"`
	WorkflowVersionId *openapi_types.UUID                  `json:"workflowVersionId,omitempty"`
}

// FavoriteForResponse_Company Favorite company
type FavoriteForResponse_Company struct {
	union json.RawMessage
}

// FavoriteForResponse_Dashboard Favorite dashboard
type FavoriteForResponse_Dashboard struct {
	union json.RawMessage
}

// FavoriteForResponse_FavoriteFolder The folder this favorite belongs to
type FavoriteForResponse_FavoriteFolder struct {
	union json.RawMessage
}

// FavoriteForResponse_ForWorkspaceMember Favorite workspace member
type FavoriteForResponse_ForWorkspaceMember struct {
	union json.RawMessage
}

// FavoriteForResponse_Note Favorite note
type FavoriteForResponse_Note struct {
	union json.RawMessage
}

// FavoriteForResponse_Opportunity Favorite opportunity
type FavoriteForResponse_Opportunity struct {
	union json.RawMessage
}

// FavoriteForResponse_Person Favorite person
type FavoriteForResponse_Person struct {
	union json.RawMessage
}

// FavoriteForResponse_Task Favorite task
type FavoriteForResponse_Task struct {
	union json.RawMessage
}

// FavoriteForResponse_Workflow Favorite workflow
type FavoriteForResponse_Workflow struct {
	union json.RawMessage
}

// FavoriteForResponse_WorkflowRun Favorite workflow run
type FavoriteForResponse_WorkflowRun struct {
	union json.RawMessage
}

// FavoriteForResponse_WorkflowVersion Favorite workflow version
type FavoriteForResponse_WorkflowVersion struct {
	union json.RawMessage
}

// FavoriteForUpdate A favorite that can be accessed from the left menu
type FavoriteForUpdate struct {
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

// Message A message sent or received through a messaging channel (email, chat, etc.)
type Message struct {
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

// MessageChannel Message Channels
type MessageChannel struct {
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

// MessageChannelForResponse Message Channels
type MessageChannelForResponse struct {
	// ConnectedAccount Connected Account
	ConnectedAccount   *MessageChannelForResponse_ConnectedAccount `json:"connectedAccount,omitempty"`
	ConnectedAccountId *openapi_types.UUID                         `json:"connectedAccountId,omitempty"`

	// ContactAutoCreationPolicy Automatically create People records when receiving or sending emails
	ContactAutoCreationPolicy *MessageChannelForResponseContactAutoCreationPolicy `json:"contactAutoCreationPolicy,omitempty"`

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
	MessageChannelMessageAssociations *[]MessageChannelMessageAssociationForResponse `json:"messageChannelMessageAssociations,omitempty"`

	// MessageFolderImportPolicy Message folder import policy
	MessageFolderImportPolicy *MessageChannelForResponseMessageFolderImportPolicy `json:"messageFolderImportPolicy,omitempty"`

	// MessageFolders Message Folders
	MessageFolders *[]MessageFolderForResponse `json:"messageFolders,omitempty"`

	// PendingGroupEmailsAction Pending action for group emails
	PendingGroupEmailsAction *MessageChannelForResponsePendingGroupEmailsAction `json:"pendingGroupEmailsAction,omitempty"`

	// SyncCursor Last sync cursor
	SyncCursor *string `json:"syncCursor,omitempty"`

	// SyncStage Sync stage
	SyncStage *MessageChannelForResponseSyncStage `json:"syncStage,omitempty"`

	// SyncStageStartedAt Sync stage started at
	SyncStageStartedAt *time.Time `json:"syncStageStartedAt,omitempty"`

	// SyncStatus Sync status
	SyncStatus *MessageChannelForResponseSyncStatus `json:"syncStatus,omitempty"`

	// SyncedAt Last sync date
	SyncedAt *time.Time `json:"syncedAt,omitempty"`

	// ThrottleFailureCount Throttle Failure Count
	ThrottleFailureCount *int `json:"throttleFailureCount,omitempty"`

	// Type Channel Type
	Type *MessageChannelForResponseType `json:"type,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Visibility Visibility
	Visibility *MessageChannelForResponseVisibility `json:"visibility,omitempty"`
}

// MessageChannelForResponse_ConnectedAccount Connected Account
type MessageChannelForResponse_ConnectedAccount struct {
	union json.RawMessage
}

// MessageChannelForResponseContactAutoCreationPolicy Automatically create People records when receiving or sending emails
type MessageChannelForResponseContactAutoCreationPolicy string

// MessageChannelForResponseMessageFolderImportPolicy Message folder import policy
type MessageChannelForResponseMessageFolderImportPolicy string

// MessageChannelForResponsePendingGroupEmailsAction Pending action for group emails
type MessageChannelForResponsePendingGroupEmailsAction string

// MessageChannelForResponseSyncStage Sync stage
type MessageChannelForResponseSyncStage string

// MessageChannelForResponseSyncStatus Sync status
type MessageChannelForResponseSyncStatus string

// MessageChannelForResponseType Channel Type
type MessageChannelForResponseType string

// MessageChannelForResponseVisibility Visibility
type MessageChannelForResponseVisibility string

// MessageChannelForUpdate Message Channels
type MessageChannelForUpdate struct {
	ConnectedAccountId *openapi_types.UUID `json:"connectedAccountId,omitempty"`

	// ContactAutoCreationPolicy Automatically create People records when receiving or sending emails
	ContactAutoCreationPolicy *MessageChannelForUpdateContactAutoCreationPolicy `json:"contactAutoCreationPolicy,omitempty"`

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
	MessageFolderImportPolicy *MessageChannelForUpdateMessageFolderImportPolicy `json:"messageFolderImportPolicy,omitempty"`

	// PendingGroupEmailsAction Pending action for group emails
	PendingGroupEmailsAction *MessageChannelForUpdatePendingGroupEmailsAction `json:"pendingGroupEmailsAction,omitempty"`

	// SyncCursor Last sync cursor
	SyncCursor *string `json:"syncCursor,omitempty"`

	// SyncStage Sync stage
	SyncStage *MessageChannelForUpdateSyncStage `json:"syncStage,omitempty"`

	// SyncStageStartedAt Sync stage started at
	SyncStageStartedAt *time.Time `json:"syncStageStartedAt,omitempty"`

	// SyncStatus Sync status
	SyncStatus *MessageChannelForUpdateSyncStatus `json:"syncStatus,omitempty"`

	// SyncedAt Last sync date
	SyncedAt *time.Time `json:"syncedAt,omitempty"`

	// ThrottleFailureCount Throttle Failure Count
	ThrottleFailureCount *int `json:"throttleFailureCount,omitempty"`

	// Type Channel Type
	Type *MessageChannelForUpdateType `json:"type,omitempty"`

	// Visibility Visibility
	Visibility *MessageChannelForUpdateVisibility `json:"visibility,omitempty"`
}

// MessageChannelForUpdateContactAutoCreationPolicy Automatically create People records when receiving or sending emails
type MessageChannelForUpdateContactAutoCreationPolicy string

// MessageChannelForUpdateMessageFolderImportPolicy Message folder import policy
type MessageChannelForUpdateMessageFolderImportPolicy string

// MessageChannelForUpdatePendingGroupEmailsAction Pending action for group emails
type MessageChannelForUpdatePendingGroupEmailsAction string

// MessageChannelForUpdateSyncStage Sync stage
type MessageChannelForUpdateSyncStage string

// MessageChannelForUpdateSyncStatus Sync status
type MessageChannelForUpdateSyncStatus string

// MessageChannelForUpdateType Channel Type
type MessageChannelForUpdateType string

// MessageChannelForUpdateVisibility Visibility
type MessageChannelForUpdateVisibility string

// MessageChannelMessageAssociation Message Synced with a Message Channel
type MessageChannelMessageAssociation struct {
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

// MessageChannelMessageAssociationForResponse Message Synced with a Message Channel
type MessageChannelMessageAssociationForResponse struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Direction Message Direction
	Direction *MessageChannelMessageAssociationForResponseDirection `json:"direction,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Message Message Id
	Message *MessageChannelMessageAssociationForResponse_Message `json:"message,omitempty"`

	// MessageChannel Message Channel Id
	MessageChannel   *MessageChannelMessageAssociationForResponse_MessageChannel `json:"messageChannel,omitempty"`
	MessageChannelId *openapi_types.UUID                                         `json:"messageChannelId,omitempty"`

	// MessageExternalId Message id from the messaging provider
	MessageExternalId *string             `json:"messageExternalId,omitempty"`
	MessageId         *openapi_types.UUID `json:"messageId,omitempty"`

	// MessageThreadExternalId Thread id from the messaging provider
	MessageThreadExternalId *string `json:"messageThreadExternalId,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// MessageChannelMessageAssociationForResponseDirection Message Direction
type MessageChannelMessageAssociationForResponseDirection string

// MessageChannelMessageAssociationForResponse_Message Message Id
type MessageChannelMessageAssociationForResponse_Message struct {
	union json.RawMessage
}

// MessageChannelMessageAssociationForResponse_MessageChannel Message Channel Id
type MessageChannelMessageAssociationForResponse_MessageChannel struct {
	union json.RawMessage
}

// MessageChannelMessageAssociationForUpdate Message Synced with a Message Channel
type MessageChannelMessageAssociationForUpdate struct {
	// Direction Message Direction
	Direction        *MessageChannelMessageAssociationForUpdateDirection `json:"direction,omitempty"`
	MessageChannelId *openapi_types.UUID                                 `json:"messageChannelId,omitempty"`

	// MessageExternalId Message id from the messaging provider
	MessageExternalId *string             `json:"messageExternalId,omitempty"`
	MessageId         *openapi_types.UUID `json:"messageId,omitempty"`

	// MessageThreadExternalId Thread id from the messaging provider
	MessageThreadExternalId *string `json:"messageThreadExternalId,omitempty"`
}

// MessageChannelMessageAssociationForUpdateDirection Message Direction
type MessageChannelMessageAssociationForUpdateDirection string

// MessageFolder Folder for Message Channel
type MessageFolder struct {
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

// MessageFolderForResponse Folder for Message Channel
type MessageFolderForResponse struct {
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
	MessageChannel   *MessageFolderForResponse_MessageChannel `json:"messageChannel,omitempty"`
	MessageChannelId *openapi_types.UUID                      `json:"messageChannelId,omitempty"`

	// Name Folder name
	Name *string `json:"name,omitempty"`

	// ParentFolderId Parent Folder ID
	ParentFolderId *string `json:"parentFolderId,omitempty"`

	// PendingSyncAction Pending action for folder sync
	PendingSyncAction *MessageFolderForResponsePendingSyncAction `json:"pendingSyncAction,omitempty"`

	// SyncCursor Sync Cursor
	SyncCursor *string `json:"syncCursor,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// MessageFolderForResponse_MessageChannel Message Channel
type MessageFolderForResponse_MessageChannel struct {
	union json.RawMessage
}

// MessageFolderForResponsePendingSyncAction Pending action for folder sync
type MessageFolderForResponsePendingSyncAction string

// MessageFolderForUpdate Folder for Message Channel
type MessageFolderForUpdate struct {
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
	PendingSyncAction *MessageFolderForUpdatePendingSyncAction `json:"pendingSyncAction,omitempty"`

	// SyncCursor Sync Cursor
	SyncCursor *string `json:"syncCursor,omitempty"`
}

// MessageFolderForUpdatePendingSyncAction Pending action for folder sync
type MessageFolderForUpdatePendingSyncAction string

// MessageForResponse A message sent or received through a messaging channel (email, chat, etc.)
type MessageForResponse struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// HeaderMessageId Message id from the message header
	HeaderMessageId *string `json:"headerMessageId,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// MessageChannelMessageAssociations Messages from the channel.
	MessageChannelMessageAssociations *[]MessageChannelMessageAssociationForResponse `json:"messageChannelMessageAssociations,omitempty"`

	// MessageParticipants Message Participants
	MessageParticipants *[]MessageParticipantForResponse `json:"messageParticipants,omitempty"`

	// MessageThread Message Thread Id
	MessageThread   *MessageForResponse_MessageThread `json:"messageThread,omitempty"`
	MessageThreadId *openapi_types.UUID               `json:"messageThreadId,omitempty"`

	// ReceivedAt The date the message was received
	ReceivedAt *time.Time `json:"receivedAt,omitempty"`

	// Subject Subject
	Subject *string `json:"subject,omitempty"`

	// Text Text
	Text *string `json:"text,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// MessageForResponse_MessageThread Message Thread Id
type MessageForResponse_MessageThread struct {
	union json.RawMessage
}

// MessageForUpdate A message sent or received through a messaging channel (email, chat, etc.)
type MessageForUpdate struct {
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

// MessageParticipant Message Participants
type MessageParticipant struct {
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
type MessageParticipantForResponse struct {
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
	Message   *MessageParticipantForResponse_Message `json:"message,omitempty"`
	MessageId *openapi_types.UUID                    `json:"messageId,omitempty"`

	// Person Person
	Person   *MessageParticipantForResponse_Person `json:"person,omitempty"`
	PersonId *openapi_types.UUID                   `json:"personId,omitempty"`

	// Role Role
	Role *MessageParticipantForResponseRole `json:"role,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// WorkspaceMember Workspace member
	WorkspaceMember   *MessageParticipantForResponse_WorkspaceMember `json:"workspaceMember,omitempty"`
	WorkspaceMemberId *openapi_types.UUID                            `json:"workspaceMemberId,omitempty"`
}

// MessageParticipantForResponse_Message Message
type MessageParticipantForResponse_Message struct {
	union json.RawMessage
}

// MessageParticipantForResponse_Person Person
type MessageParticipantForResponse_Person struct {
	union json.RawMessage
}

// MessageParticipantForResponseRole Role
type MessageParticipantForResponseRole string

// MessageParticipantForResponse_WorkspaceMember Workspace member
type MessageParticipantForResponse_WorkspaceMember struct {
	union json.RawMessage
}

// MessageParticipantForUpdate Message Participants
type MessageParticipantForUpdate struct {
	// DisplayName Display Name
	DisplayName *string `json:"displayName,omitempty"`

	// Handle Handle
	Handle    *string             `json:"handle,omitempty"`
	MessageId *openapi_types.UUID `json:"messageId,omitempty"`
	PersonId  *openapi_types.UUID `json:"personId,omitempty"`

	// Role Role
	Role              *MessageParticipantForUpdateRole `json:"role,omitempty"`
	WorkspaceMemberId *openapi_types.UUID              `json:"workspaceMemberId,omitempty"`
}

// MessageParticipantForUpdateRole Role
type MessageParticipantForUpdateRole string

// MessageThread A group of related messages (e.g. email thread, chat thread)
type MessageThread = map[string]interface{}

// MessageThreadForResponse A group of related messages (e.g. email thread, chat thread)
type MessageThreadForResponse struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Messages Messages from the thread.
	Messages *[]MessageForResponse `json:"messages,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// MessageThreadForUpdate A group of related messages (e.g. email thread, chat thread)
type MessageThreadForUpdate = map[string]interface{}

// Note A note
type Note struct {
	// BodyV2 Note body
	BodyV2 *struct {
		Blocknote *string `json:"blocknote,omitempty"`
		Markdown  *string `json:"markdown,omitempty"`
	} `json:"bodyV2,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *NoteCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// Position Note record position
	Position *float32 `json:"position,omitempty"`

	// Title Note title
	Title *string `json:"title,omitempty"`
}

// NoteCreatedBySource defines model for Note.CreatedBy.Source.
type NoteCreatedBySource string

// NoteForResponse A note
type NoteForResponse struct {
	// Attachments Note attachments
	Attachments *[]AttachmentForResponse `json:"attachments,omitempty"`

	// BodyV2 Note body
	BodyV2 *struct {
		Blocknote *string `json:"blocknote,omitempty"`
		Markdown  *string `json:"markdown,omitempty"`
	} `json:"bodyV2,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Name              *string                         `json:"name,omitempty"`
		Source            *NoteForResponseCreatedBySource `json:"source,omitempty"`
		WorkspaceMemberId *openapi_types.UUID             `json:"workspaceMemberId,omitempty"`
	} `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the note
	Favorites *[]FavoriteForResponse `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// NoteTargets Note targets
	NoteTargets *[]NoteTargetForResponse `json:"noteTargets,omitempty"`

	// Position Note record position
	Position *float32 `json:"position,omitempty"`

	// TimelineActivities Timeline Activities linked to the note.
	TimelineActivities *[]TimelineActivityForResponse `json:"timelineActivities,omitempty"`

	// Title Note title
	Title *string `json:"title,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NoteForResponseCreatedBySource defines model for NoteForResponse.CreatedBy.Source.
type NoteForResponseCreatedBySource string

// NoteForUpdate A note
type NoteForUpdate struct {
	// BodyV2 Note body
	BodyV2 *struct {
		Blocknote *string `json:"blocknote,omitempty"`
		Markdown  *string `json:"markdown,omitempty"`
	} `json:"bodyV2,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *NoteForUpdateCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// Position Note record position
	Position *float32 `json:"position,omitempty"`

	// Title Note title
	Title *string `json:"title,omitempty"`
}

// NoteForUpdateCreatedBySource defines model for NoteForUpdate.CreatedBy.Source.
type NoteForUpdateCreatedBySource string

// NoteTarget A note target
type NoteTarget struct {
	CompanyId     *openapi_types.UUID `json:"companyId,omitempty"`
	NoteId        *openapi_types.UUID `json:"noteId,omitempty"`
	OpportunityId *openapi_types.UUID `json:"opportunityId,omitempty"`
	PersonId      *openapi_types.UUID `json:"personId,omitempty"`
}

// NoteTargetForResponse A note target
type NoteTargetForResponse struct {
	// Company NoteTarget company
	Company   *NoteTargetForResponse_Company `json:"company,omitempty"`
	CompanyId *openapi_types.UUID            `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Note NoteTarget note
	Note   *NoteTargetForResponse_Note `json:"note,omitempty"`
	NoteId *openapi_types.UUID         `json:"noteId,omitempty"`

	// Opportunity NoteTarget opportunity
	Opportunity   *NoteTargetForResponse_Opportunity `json:"opportunity,omitempty"`
	OpportunityId *openapi_types.UUID                `json:"opportunityId,omitempty"`

	// Person NoteTarget person
	Person   *NoteTargetForResponse_Person `json:"person,omitempty"`
	PersonId *openapi_types.UUID           `json:"personId,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NoteTargetForResponse_Company NoteTarget company
type NoteTargetForResponse_Company struct {
	union json.RawMessage
}

// NoteTargetForResponse_Note NoteTarget note
type NoteTargetForResponse_Note struct {
	union json.RawMessage
}

// NoteTargetForResponse_Opportunity NoteTarget opportunity
type NoteTargetForResponse_Opportunity struct {
	union json.RawMessage
}

// NoteTargetForResponse_Person NoteTarget person
type NoteTargetForResponse_Person struct {
	union json.RawMessage
}

// NoteTargetForUpdate A note target
type NoteTargetForUpdate struct {
	CompanyId     *openapi_types.UUID `json:"companyId,omitempty"`
	NoteId        *openapi_types.UUID `json:"noteId,omitempty"`
	OpportunityId *openapi_types.UUID `json:"opportunityId,omitempty"`
	PersonId      *openapi_types.UUID `json:"personId,omitempty"`
}

// Opportunity An opportunity
type Opportunity struct {
	// Amount Opportunity amount
	Amount *struct {
		AmountMicros *float32 `json:"amountMicros,omitempty"`
		CurrencyCode *string  `json:"currencyCode,omitempty"`
	} `json:"amount,omitempty"`

	// CloseDate Opportunity close date
	CloseDate *time.Time          `json:"closeDate,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *OpportunityCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`
	LostReason *OpportunityLostReason `json:"lostReason,omitempty"`

	// Name The opportunity name
	Name             *string             `json:"name,omitempty"`
	Organizationid   *string             `json:"organizationid,omitempty"`
	PointOfContactId *openapi_types.UUID `json:"pointOfContactId,omitempty"`

	// Position Opportunity record position
	Position *float32 `json:"position,omitempty"`

	// Stage Opportunity stage
	Stage              *OpportunityStage   `json:"stage,omitempty"`
	SubscriptionActive *bool               `json:"subscriptionActive,omitempty"`
	SubscriptionAmount *int                `json:"subscriptionAmount,omitempty"`
	SubscriptionEnd    *openapi_types.Date `json:"subscriptionEnd,omitempty"`
	SubscriptionStart  *openapi_types.Date `json:"subscriptionStart,omitempty"`
}

// OpportunityCreatedBySource defines model for Opportunity.CreatedBy.Source.
type OpportunityCreatedBySource string

// OpportunityLostReason defines model for Opportunity.LostReason.
type OpportunityLostReason string

// OpportunityStage Opportunity stage
type OpportunityStage string

// OpportunityForResponse An opportunity
type OpportunityForResponse struct {
	// Amount Opportunity amount
	Amount *struct {
		AmountMicros *float32 `json:"amountMicros,omitempty"`
		CurrencyCode *string  `json:"currencyCode,omitempty"`
	} `json:"amount,omitempty"`

	// Attachments Attachments linked to the opportunity
	Attachments *[]AttachmentForResponse `json:"attachments,omitempty"`

	// CloseDate Opportunity close date
	CloseDate *time.Time `json:"closeDate,omitempty"`

	// Company Opportunity company
	Company   *OpportunityForResponse_Company `json:"company,omitempty"`
	CompanyId *openapi_types.UUID             `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Name              *string                                `json:"name,omitempty"`
		Source            *OpportunityForResponseCreatedBySource `json:"source,omitempty"`
		WorkspaceMemberId *openapi_types.UUID                    `json:"workspaceMemberId,omitempty"`
	} `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the opportunity
	Favorites *[]FavoriteForResponse `json:"favorites,omitempty"`

	// Id Id
	Id         *openapi_types.UUID               `json:"id,omitempty"`
	LostReason *OpportunityForResponseLostReason `json:"lostReason,omitempty"`

	// Name The opportunity name
	Name *string `json:"name,omitempty"`

	// NoteTargets Notes tied to the opportunity
	NoteTargets    *[]NoteTargetForResponse `json:"noteTargets,omitempty"`
	Organizationid *string                  `json:"organizationid,omitempty"`

	// PointOfContact Opportunity point of contact
	PointOfContact   *OpportunityForResponse_PointOfContact `json:"pointOfContact,omitempty"`
	PointOfContactId *openapi_types.UUID                    `json:"pointOfContactId,omitempty"`

	// Position Opportunity record position
	Position *float32 `json:"position,omitempty"`

	// Stage Opportunity stage
	Stage              *OpportunityForResponseStage `json:"stage,omitempty"`
	SubscriptionActive *bool                        `json:"subscriptionActive,omitempty"`
	SubscriptionAmount *int                         `json:"subscriptionAmount,omitempty"`
	SubscriptionEnd    *openapi_types.Date          `json:"subscriptionEnd,omitempty"`
	SubscriptionStart  *openapi_types.Date          `json:"subscriptionStart,omitempty"`

	// TaskTargets Tasks tied to the opportunity
	TaskTargets *[]TaskTargetForResponse `json:"taskTargets,omitempty"`

	// TimelineActivities Timeline Activities linked to the opportunity.
	TimelineActivities *[]TimelineActivityForResponse `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// OpportunityForResponse_Company Opportunity company
type OpportunityForResponse_Company struct {
	union json.RawMessage
}

// OpportunityForResponseCreatedBySource defines model for OpportunityForResponse.CreatedBy.Source.
type OpportunityForResponseCreatedBySource string

// OpportunityForResponseLostReason defines model for OpportunityForResponse.LostReason.
type OpportunityForResponseLostReason string

// OpportunityForResponse_PointOfContact Opportunity point of contact
type OpportunityForResponse_PointOfContact struct {
	union json.RawMessage
}

// OpportunityForResponseStage Opportunity stage
type OpportunityForResponseStage string

// OpportunityForUpdate An opportunity
type OpportunityForUpdate struct {
	// Amount Opportunity amount
	Amount *struct {
		AmountMicros *float32 `json:"amountMicros,omitempty"`
		CurrencyCode *string  `json:"currencyCode,omitempty"`
	} `json:"amount,omitempty"`

	// CloseDate Opportunity close date
	CloseDate *time.Time          `json:"closeDate,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *OpportunityForUpdateCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`
	LostReason *OpportunityForUpdateLostReason `json:"lostReason,omitempty"`

	// Name The opportunity name
	Name             *string             `json:"name,omitempty"`
	Organizationid   *string             `json:"organizationid,omitempty"`
	PointOfContactId *openapi_types.UUID `json:"pointOfContactId,omitempty"`

	// Position Opportunity record position
	Position *float32 `json:"position,omitempty"`

	// Stage Opportunity stage
	Stage              *OpportunityForUpdateStage `json:"stage,omitempty"`
	SubscriptionActive *bool                      `json:"subscriptionActive,omitempty"`
	SubscriptionAmount *int                       `json:"subscriptionAmount,omitempty"`
	SubscriptionEnd    *openapi_types.Date        `json:"subscriptionEnd,omitempty"`
	SubscriptionStart  *openapi_types.Date        `json:"subscriptionStart,omitempty"`
}

// OpportunityForUpdateCreatedBySource defines model for OpportunityForUpdate.CreatedBy.Source.
type OpportunityForUpdateCreatedBySource string

// OpportunityForUpdateLostReason defines model for OpportunityForUpdate.LostReason.
type OpportunityForUpdateLostReason string

// OpportunityForUpdateStage Opportunity stage
type OpportunityForUpdateStage string

// Person A person
type Person struct {
	// AvatarUrl Contact’s avatar
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// City Contact’s city
	City      *string             `json:"city,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *PersonCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// Emails Contact’s Emails
	Emails *struct {
		AdditionalEmails *[]openapi_types.Email `json:"additionalEmails,omitempty"`
		PrimaryEmail     *string                `json:"primaryEmail,omitempty"`
	} `json:"emails,omitempty"`

	// JobTitle Contact’s job title
	JobTitle   *string `json:"jobTitle,omitempty"`
	Keycloakid *string `json:"keycloakid,omitempty"`

	// LinkedinLink Contact’s Linkedin account
	LinkedinLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"linkedinLink,omitempty"`

	// Name Contact’s name
	Name *struct {
		FirstName *string `json:"firstName,omitempty"`
		LastName  *string `json:"lastName,omitempty"`
	} `json:"name,omitempty"`
	Newsletter *bool `json:"newsletter,omitempty"`

	// Phones Contact’s phone numbers
	Phones *struct {
		AdditionalPhones        *[]string `json:"additionalPhones,omitempty"`
		PrimaryPhoneCallingCode *string   `json:"primaryPhoneCallingCode,omitempty"`
		PrimaryPhoneCountryCode *string   `json:"primaryPhoneCountryCode,omitempty"`
		PrimaryPhoneNumber      *string   `json:"primaryPhoneNumber,omitempty"`
	} `json:"phones,omitempty"`

	// Position Person record Position
	Position *float32 `json:"position,omitempty"`

	// XLink Contact’s X/Twitter account
	XLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"xLink,omitempty"`
}

// PersonCreatedBySource defines model for Person.CreatedBy.Source.
type PersonCreatedBySource string

// PersonForResponse A person
type PersonForResponse struct {
	// Attachments Attachments linked to the contact.
	Attachments *[]AttachmentForResponse `json:"attachments,omitempty"`

	// AvatarUrl Contact’s avatar
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// CalendarEventParticipants Calendar Event Participants
	CalendarEventParticipants *[]CalendarEventParticipantForResponse `json:"calendarEventParticipants,omitempty"`

	// City Contact’s city
	City *string `json:"city,omitempty"`

	// Company Contact’s company
	Company   *PersonForResponse_Company `json:"company,omitempty"`
	CompanyId *openapi_types.UUID        `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Name              *string                           `json:"name,omitempty"`
		Source            *PersonForResponseCreatedBySource `json:"source,omitempty"`
		WorkspaceMemberId *openapi_types.UUID               `json:"workspaceMemberId,omitempty"`
	} `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Emails Contact’s Emails
	Emails *struct {
		AdditionalEmails *[]openapi_types.Email `json:"additionalEmails,omitempty"`
		PrimaryEmail     *string                `json:"primaryEmail,omitempty"`
	} `json:"emails,omitempty"`

	// Favorites Favorites linked to the contact
	Favorites *[]FavoriteForResponse `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// JobTitle Contact’s job title
	JobTitle   *string `json:"jobTitle,omitempty"`
	Keycloakid *string `json:"keycloakid,omitempty"`

	// LinkedinLink Contact’s Linkedin account
	LinkedinLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"linkedinLink,omitempty"`

	// MessageParticipants Message Participants
	MessageParticipants *[]MessageParticipantForResponse `json:"messageParticipants,omitempty"`

	// Name Contact’s name
	Name *struct {
		FirstName *string `json:"firstName,omitempty"`
		LastName  *string `json:"lastName,omitempty"`
	} `json:"name,omitempty"`
	Newsletter *bool `json:"newsletter,omitempty"`

	// NoteTargets Notes tied to the contact
	NoteTargets *[]NoteTargetForResponse `json:"noteTargets,omitempty"`

	// Phones Contact’s phone numbers
	Phones *struct {
		AdditionalPhones        *[]string `json:"additionalPhones,omitempty"`
		PrimaryPhoneCallingCode *string   `json:"primaryPhoneCallingCode,omitempty"`
		PrimaryPhoneCountryCode *string   `json:"primaryPhoneCountryCode,omitempty"`
		PrimaryPhoneNumber      *string   `json:"primaryPhoneNumber,omitempty"`
	} `json:"phones,omitempty"`

	// PointOfContactForOpportunities List of opportunities for which that person is the point of contact
	PointOfContactForOpportunities *[]OpportunityForResponse `json:"pointOfContactForOpportunities,omitempty"`

	// Position Person record Position
	Position *float32 `json:"position,omitempty"`

	// TaskTargets Tasks tied to the contact
	TaskTargets *[]TaskTargetForResponse `json:"taskTargets,omitempty"`

	// TimelineActivities Events linked to the person
	TimelineActivities *[]TimelineActivityForResponse `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// XLink Contact’s X/Twitter account
	XLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"xLink,omitempty"`
}

// PersonForResponse_Company Contact’s company
type PersonForResponse_Company struct {
	union json.RawMessage
}

// PersonForResponseCreatedBySource defines model for PersonForResponse.CreatedBy.Source.
type PersonForResponseCreatedBySource string

// PersonForUpdate A person
type PersonForUpdate struct {
	// AvatarUrl Contact’s avatar
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// City Contact’s city
	City      *string             `json:"city,omitempty"`
	CompanyId *openapi_types.UUID `json:"companyId,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *PersonForUpdateCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// Emails Contact’s Emails
	Emails *struct {
		AdditionalEmails *[]openapi_types.Email `json:"additionalEmails,omitempty"`
		PrimaryEmail     *string                `json:"primaryEmail,omitempty"`
	} `json:"emails,omitempty"`

	// JobTitle Contact’s job title
	JobTitle   *string `json:"jobTitle,omitempty"`
	Keycloakid *string `json:"keycloakid,omitempty"`

	// LinkedinLink Contact’s Linkedin account
	LinkedinLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"linkedinLink,omitempty"`

	// Name Contact’s name
	Name *struct {
		FirstName *string `json:"firstName,omitempty"`
		LastName  *string `json:"lastName,omitempty"`
	} `json:"name,omitempty"`
	Newsletter *bool `json:"newsletter,omitempty"`

	// Phones Contact’s phone numbers
	Phones *struct {
		AdditionalPhones        *[]string `json:"additionalPhones,omitempty"`
		PrimaryPhoneCallingCode *string   `json:"primaryPhoneCallingCode,omitempty"`
		PrimaryPhoneCountryCode *string   `json:"primaryPhoneCountryCode,omitempty"`
		PrimaryPhoneNumber      *string   `json:"primaryPhoneNumber,omitempty"`
	} `json:"phones,omitempty"`

	// Position Person record Position
	Position *float32 `json:"position,omitempty"`

	// XLink Contact’s X/Twitter account
	XLink *struct {
		PrimaryLinkLabel *string `json:"primaryLinkLabel,omitempty"`
		PrimaryLinkUrl   *string `json:"primaryLinkUrl,omitempty"`
		SecondaryLinks   *[]struct {
			Label *string `json:"label,omitempty"`
			Url   *string `json:"url,omitempty"`
		} `json:"secondaryLinks,omitempty"`
	} `json:"xLink,omitempty"`
}

// PersonForUpdateCreatedBySource defines model for PersonForUpdate.CreatedBy.Source.
type PersonForUpdateCreatedBySource string

// Task A task
type Task struct {
	AssigneeId *openapi_types.UUID `json:"assigneeId,omitempty"`

	// BodyV2 Task body
	BodyV2 *struct {
		Blocknote *string `json:"blocknote,omitempty"`
		Markdown  *string `json:"markdown,omitempty"`
	} `json:"bodyV2,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *TaskCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// DueAt Task due date
	DueAt *time.Time `json:"dueAt,omitempty"`

	// Position Task record position
	Position *float32 `json:"position,omitempty"`

	// Status Task status
	Status *TaskStatus `json:"status,omitempty"`

	// Title Task title
	Title *string `json:"title,omitempty"`
}

// TaskCreatedBySource defines model for Task.CreatedBy.Source.
type TaskCreatedBySource string

// TaskStatus Task status
type TaskStatus string

// TaskForResponse A task
type TaskForResponse struct {
	// Assignee Task assignee
	Assignee   *TaskForResponse_Assignee `json:"assignee,omitempty"`
	AssigneeId *openapi_types.UUID       `json:"assigneeId,omitempty"`

	// Attachments Task attachments
	Attachments *[]AttachmentForResponse `json:"attachments,omitempty"`

	// BodyV2 Task body
	BodyV2 *struct {
		Blocknote *string `json:"blocknote,omitempty"`
		Markdown  *string `json:"markdown,omitempty"`
	} `json:"bodyV2,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Name              *string                         `json:"name,omitempty"`
		Source            *TaskForResponseCreatedBySource `json:"source,omitempty"`
		WorkspaceMemberId *openapi_types.UUID             `json:"workspaceMemberId,omitempty"`
	} `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// DueAt Task due date
	DueAt *time.Time `json:"dueAt,omitempty"`

	// Favorites Favorites linked to the task
	Favorites *[]FavoriteForResponse `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Position Task record position
	Position *float32 `json:"position,omitempty"`

	// Status Task status
	Status *TaskForResponseStatus `json:"status,omitempty"`

	// TaskTargets Task targets
	TaskTargets *[]TaskTargetForResponse `json:"taskTargets,omitempty"`

	// TimelineActivities Timeline Activities linked to the task.
	TimelineActivities *[]TimelineActivityForResponse `json:"timelineActivities,omitempty"`

	// Title Task title
	Title *string `json:"title,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// TaskForResponse_Assignee Task assignee
type TaskForResponse_Assignee struct {
	union json.RawMessage
}

// TaskForResponseCreatedBySource defines model for TaskForResponse.CreatedBy.Source.
type TaskForResponseCreatedBySource string

// TaskForResponseStatus Task status
type TaskForResponseStatus string

// TaskForUpdate A task
type TaskForUpdate struct {
	AssigneeId *openapi_types.UUID `json:"assigneeId,omitempty"`

	// BodyV2 Task body
	BodyV2 *struct {
		Blocknote *string `json:"blocknote,omitempty"`
		Markdown  *string `json:"markdown,omitempty"`
	} `json:"bodyV2,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *TaskForUpdateCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// DueAt Task due date
	DueAt *time.Time `json:"dueAt,omitempty"`

	// Position Task record position
	Position *float32 `json:"position,omitempty"`

	// Status Task status
	Status *TaskForUpdateStatus `json:"status,omitempty"`

	// Title Task title
	Title *string `json:"title,omitempty"`
}

// TaskForUpdateCreatedBySource defines model for TaskForUpdate.CreatedBy.Source.
type TaskForUpdateCreatedBySource string

// TaskForUpdateStatus Task status
type TaskForUpdateStatus string

// TaskTarget A task target
type TaskTarget struct {
	CompanyId     *openapi_types.UUID `json:"companyId,omitempty"`
	OpportunityId *openapi_types.UUID `json:"opportunityId,omitempty"`
	PersonId      *openapi_types.UUID `json:"personId,omitempty"`
	TaskId        *openapi_types.UUID `json:"taskId,omitempty"`
}

// TaskTargetForResponse A task target
type TaskTargetForResponse struct {
	// Company TaskTarget company
	Company   *TaskTargetForResponse_Company `json:"company,omitempty"`
	CompanyId *openapi_types.UUID            `json:"companyId,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Opportunity TaskTarget opportunity
	Opportunity   *TaskTargetForResponse_Opportunity `json:"opportunity,omitempty"`
	OpportunityId *openapi_types.UUID                `json:"opportunityId,omitempty"`

	// Person TaskTarget person
	Person   *TaskTargetForResponse_Person `json:"person,omitempty"`
	PersonId *openapi_types.UUID           `json:"personId,omitempty"`

	// Task TaskTarget task
	Task   *TaskTargetForResponse_Task `json:"task,omitempty"`
	TaskId *openapi_types.UUID         `json:"taskId,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// TaskTargetForResponse_Company TaskTarget company
type TaskTargetForResponse_Company struct {
	union json.RawMessage
}

// TaskTargetForResponse_Opportunity TaskTarget opportunity
type TaskTargetForResponse_Opportunity struct {
	union json.RawMessage
}

// TaskTargetForResponse_Person TaskTarget person
type TaskTargetForResponse_Person struct {
	union json.RawMessage
}

// TaskTargetForResponse_Task TaskTarget task
type TaskTargetForResponse_Task struct {
	union json.RawMessage
}

// TaskTargetForUpdate A task target
type TaskTargetForUpdate struct {
	CompanyId     *openapi_types.UUID `json:"companyId,omitempty"`
	OpportunityId *openapi_types.UUID `json:"opportunityId,omitempty"`
	PersonId      *openapi_types.UUID `json:"personId,omitempty"`
	TaskId        *openapi_types.UUID `json:"taskId,omitempty"`
}

// TimelineActivity Aggregated / filtered event to be displayed on the timeline
type TimelineActivity struct {
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
type TimelineActivityForResponse struct {
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
	WorkspaceMember   *TimelineActivityForResponse_WorkspaceMember `json:"workspaceMember,omitempty"`
	WorkspaceMemberId *openapi_types.UUID                          `json:"workspaceMemberId,omitempty"`
}

// TimelineActivityForResponse_WorkspaceMember Event workspace member
type TimelineActivityForResponse_WorkspaceMember struct {
	union json.RawMessage
}

// TimelineActivityForUpdate Aggregated / filtered event to be displayed on the timeline
type TimelineActivityForUpdate struct {
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

// Workflow A workflow
type Workflow struct {
	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *WorkflowCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// LastPublishedVersionId The workflow last published version id
	LastPublishedVersionId *string `json:"lastPublishedVersionId,omitempty"`

	// Name The workflow name
	Name *string `json:"name,omitempty"`

	// Position Workflow record position
	Position *float32 `json:"position,omitempty"`

	// Statuses The current statuses of the workflow versions
	Statuses *[]WorkflowStatuses `json:"statuses,omitempty"`
}

// WorkflowCreatedBySource defines model for Workflow.CreatedBy.Source.
type WorkflowCreatedBySource string

// WorkflowStatuses defines model for Workflow.Statuses.
type WorkflowStatuses string

// WorkflowAutomatedTrigger A workflow automated trigger
type WorkflowAutomatedTrigger struct {
	// Settings The workflow automated trigger settings
	Settings map[string]interface{} `json:"settings"`

	// Type The workflow automated trigger type
	Type       WorkflowAutomatedTriggerType `json:"type"`
	WorkflowId *openapi_types.UUID          `json:"workflowId,omitempty"`
}

// WorkflowAutomatedTriggerType The workflow automated trigger type
type WorkflowAutomatedTriggerType string

// WorkflowAutomatedTriggerForResponse A workflow automated trigger
type WorkflowAutomatedTriggerForResponse struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Settings The workflow automated trigger settings
	Settings *map[string]interface{} `json:"settings,omitempty"`

	// Type The workflow automated trigger type
	Type *WorkflowAutomatedTriggerForResponseType `json:"type,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Workflow WorkflowAutomatedTrigger workflow
	Workflow   *WorkflowAutomatedTriggerForResponse_Workflow `json:"workflow,omitempty"`
	WorkflowId *openapi_types.UUID                           `json:"workflowId,omitempty"`
}

// WorkflowAutomatedTriggerForResponseType The workflow automated trigger type
type WorkflowAutomatedTriggerForResponseType string

// WorkflowAutomatedTriggerForResponse_Workflow WorkflowAutomatedTrigger workflow
type WorkflowAutomatedTriggerForResponse_Workflow struct {
	union json.RawMessage
}

// WorkflowAutomatedTriggerForUpdate A workflow automated trigger
type WorkflowAutomatedTriggerForUpdate struct {
	// Settings The workflow automated trigger settings
	Settings *map[string]interface{} `json:"settings,omitempty"`

	// Type The workflow automated trigger type
	Type       *WorkflowAutomatedTriggerForUpdateType `json:"type,omitempty"`
	WorkflowId *openapi_types.UUID                    `json:"workflowId,omitempty"`
}

// WorkflowAutomatedTriggerForUpdateType The workflow automated trigger type
type WorkflowAutomatedTriggerForUpdateType string

// WorkflowForResponse A workflow
type WorkflowForResponse struct {
	// Attachments Attachments linked to the workflow
	Attachments *[]AttachmentForResponse `json:"attachments,omitempty"`

	// AutomatedTriggers Workflow automated triggers linked to the workflow.
	AutomatedTriggers *[]WorkflowAutomatedTriggerForResponse `json:"automatedTriggers,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The creator of the record
	CreatedBy *struct {
		Name              *string                             `json:"name,omitempty"`
		Source            *WorkflowForResponseCreatedBySource `json:"source,omitempty"`
		WorkspaceMemberId *openapi_types.UUID                 `json:"workspaceMemberId,omitempty"`
	} `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the workflow
	Favorites *[]FavoriteForResponse `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// LastPublishedVersionId The workflow last published version id
	LastPublishedVersionId *string `json:"lastPublishedVersionId,omitempty"`

	// Name The workflow name
	Name *string `json:"name,omitempty"`

	// Position Workflow record position
	Position *float32 `json:"position,omitempty"`

	// Runs Workflow runs linked to the workflow.
	Runs *[]WorkflowRunForResponse `json:"runs,omitempty"`

	// Statuses The current statuses of the workflow versions
	Statuses *[]WorkflowForResponseStatuses `json:"statuses,omitempty"`

	// TimelineActivities Timeline activities linked to the workflow
	TimelineActivities *[]TimelineActivityForResponse `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Versions Workflow versions linked to the workflow.
	Versions *[]WorkflowVersionForResponse `json:"versions,omitempty"`
}

// WorkflowForResponseCreatedBySource defines model for WorkflowForResponse.CreatedBy.Source.
type WorkflowForResponseCreatedBySource string

// WorkflowForResponseStatuses defines model for WorkflowForResponse.Statuses.
type WorkflowForResponseStatuses string

// WorkflowForUpdate A workflow
type WorkflowForUpdate struct {
	// CreatedBy The creator of the record
	CreatedBy *struct {
		Source *WorkflowForUpdateCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

	// LastPublishedVersionId The workflow last published version id
	LastPublishedVersionId *string `json:"lastPublishedVersionId,omitempty"`

	// Name The workflow name
	Name *string `json:"name,omitempty"`

	// Position Workflow record position
	Position *float32 `json:"position,omitempty"`

	// Statuses The current statuses of the workflow versions
	Statuses *[]WorkflowForUpdateStatuses `json:"statuses,omitempty"`
}

// WorkflowForUpdateCreatedBySource defines model for WorkflowForUpdate.CreatedBy.Source.
type WorkflowForUpdateCreatedBySource string

// WorkflowForUpdateStatuses defines model for WorkflowForUpdate.Statuses.
type WorkflowForUpdateStatuses string

// WorkflowRun A workflow run
type WorkflowRun struct {
	// CreatedBy The executor of the workflow
	CreatedBy *struct {
		Source *WorkflowRunCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

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

// WorkflowRunCreatedBySource defines model for WorkflowRun.CreatedBy.Source.
type WorkflowRunCreatedBySource string

// WorkflowRunStatus Workflow run status
type WorkflowRunStatus string

// WorkflowRunForResponse A workflow run
type WorkflowRunForResponse struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy The executor of the workflow
	CreatedBy *struct {
		Name              *string                                `json:"name,omitempty"`
		Source            *WorkflowRunForResponseCreatedBySource `json:"source,omitempty"`
		WorkspaceMemberId *openapi_types.UUID                    `json:"workspaceMemberId,omitempty"`
	} `json:"createdBy,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// EndedAt Workflow run ended at
	EndedAt *time.Time `json:"endedAt,omitempty"`

	// EnqueuedAt Workflow run enqueued at
	EnqueuedAt *time.Time `json:"enqueuedAt,omitempty"`

	// Favorites Favorites linked to the workflow run
	Favorites *[]FavoriteForResponse `json:"favorites,omitempty"`

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
	Status *WorkflowRunForResponseStatus `json:"status,omitempty"`

	// TimelineActivities Timeline activities linked to the run
	TimelineActivities *[]TimelineActivityForResponse `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Workflow Workflow linked to the run.
	Workflow   *WorkflowRunForResponse_Workflow `json:"workflow,omitempty"`
	WorkflowId *openapi_types.UUID              `json:"workflowId,omitempty"`

	// WorkflowVersion Workflow version linked to the run.
	WorkflowVersion   *WorkflowRunForResponse_WorkflowVersion `json:"workflowVersion,omitempty"`
	WorkflowVersionId *openapi_types.UUID                     `json:"workflowVersionId,omitempty"`
}

// WorkflowRunForResponseCreatedBySource defines model for WorkflowRunForResponse.CreatedBy.Source.
type WorkflowRunForResponseCreatedBySource string

// WorkflowRunForResponseStatus Workflow run status
type WorkflowRunForResponseStatus string

// WorkflowRunForResponse_Workflow Workflow linked to the run.
type WorkflowRunForResponse_Workflow struct {
	union json.RawMessage
}

// WorkflowRunForResponse_WorkflowVersion Workflow version linked to the run.
type WorkflowRunForResponse_WorkflowVersion struct {
	union json.RawMessage
}

// WorkflowRunForUpdate A workflow run
type WorkflowRunForUpdate struct {
	// CreatedBy The executor of the workflow
	CreatedBy *struct {
		Source *WorkflowRunForUpdateCreatedBySource `json:"source,omitempty"`
	} `json:"createdBy,omitempty"`

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
	State *map[string]interface{} `json:"state,omitempty"`

	// Status Workflow run status
	Status            *WorkflowRunForUpdateStatus `json:"status,omitempty"`
	WorkflowId        *openapi_types.UUID         `json:"workflowId,omitempty"`
	WorkflowVersionId *openapi_types.UUID         `json:"workflowVersionId,omitempty"`
}

// WorkflowRunForUpdateCreatedBySource defines model for WorkflowRunForUpdate.CreatedBy.Source.
type WorkflowRunForUpdateCreatedBySource string

// WorkflowRunForUpdateStatus Workflow run status
type WorkflowRunForUpdateStatus string

// WorkflowVersion A workflow version
type WorkflowVersion struct {
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

// WorkflowVersionForResponse A workflow version
type WorkflowVersionForResponse struct {
	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the workflow version
	Favorites *[]FavoriteForResponse `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Name The workflow version name
	Name *string `json:"name,omitempty"`

	// Position Workflow version position
	Position *float32 `json:"position,omitempty"`

	// Runs Workflow runs linked to the version.
	Runs *[]WorkflowRunForResponse `json:"runs,omitempty"`

	// Status The workflow version status
	Status *WorkflowVersionForResponseStatus `json:"status,omitempty"`

	// Steps Json object to provide steps
	Steps *map[string]interface{} `json:"steps,omitempty"`

	// TimelineActivities Timeline activities linked to the version
	TimelineActivities *[]TimelineActivityForResponse `json:"timelineActivities,omitempty"`

	// Trigger Json object to provide trigger
	Trigger *map[string]interface{} `json:"trigger,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Workflow WorkflowVersion workflow
	Workflow   *WorkflowVersionForResponse_Workflow `json:"workflow,omitempty"`
	WorkflowId *openapi_types.UUID                  `json:"workflowId,omitempty"`
}

// WorkflowVersionForResponseStatus The workflow version status
type WorkflowVersionForResponseStatus string

// WorkflowVersionForResponse_Workflow WorkflowVersion workflow
type WorkflowVersionForResponse_Workflow struct {
	union json.RawMessage
}

// WorkflowVersionForUpdate A workflow version
type WorkflowVersionForUpdate struct {
	// Name The workflow version name
	Name *string `json:"name,omitempty"`

	// Position Workflow version position
	Position *float32 `json:"position,omitempty"`

	// Status The workflow version status
	Status *WorkflowVersionForUpdateStatus `json:"status,omitempty"`

	// Steps Json object to provide steps
	Steps *map[string]interface{} `json:"steps,omitempty"`

	// Trigger Json object to provide trigger
	Trigger    *map[string]interface{} `json:"trigger,omitempty"`
	WorkflowId *openapi_types.UUID     `json:"workflowId,omitempty"`
}

// WorkflowVersionForUpdateStatus The workflow version status
type WorkflowVersionForUpdateStatus string

// WorkspaceMember A workspace member
type WorkspaceMember struct {
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
	Name struct {
		FirstName *string `json:"firstName,omitempty"`
		LastName  *string `json:"lastName,omitempty"`
	} `json:"name"`

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

// WorkspaceMemberForResponse A workspace member
type WorkspaceMemberForResponse struct {
	// AccountOwnerForCompanies Account owner for companies
	AccountOwnerForCompanies *[]CompanyForResponse `json:"accountOwnerForCompanies,omitempty"`

	// AssignedTasks Tasks assigned to the workspace member
	AssignedTasks *[]TaskForResponse `json:"assignedTasks,omitempty"`

	// AuthoredAttachments Attachments created by the workspace member
	AuthoredAttachments *[]AttachmentForResponse `json:"authoredAttachments,omitempty"`

	// AvatarUrl Workspace member avatar
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// Blocklist Blocklisted handles
	Blocklist *[]BlocklistForResponse `json:"blocklist,omitempty"`

	// CalendarEventParticipants Calendar Event Participants
	CalendarEventParticipants *[]CalendarEventParticipantForResponse `json:"calendarEventParticipants,omitempty"`

	// CalendarStartDay User's preferred start day of the week
	CalendarStartDay *int `json:"calendarStartDay,omitempty"`

	// ColorScheme Preferred color scheme
	ColorScheme *string `json:"colorScheme,omitempty"`

	// ConnectedAccounts Connected accounts
	ConnectedAccounts *[]ConnectedAccountForResponse `json:"connectedAccounts,omitempty"`

	// CreatedAt Creation date
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// DateFormat User's preferred date format
	DateFormat *WorkspaceMemberForResponseDateFormat `json:"dateFormat,omitempty"`

	// DeletedAt Date when the record was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	// Favorites Favorites linked to the workspace member
	Favorites *[]FavoriteForResponse `json:"favorites,omitempty"`

	// Id Id
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Locale Preferred language
	Locale *string `json:"locale,omitempty"`

	// MessageParticipants Message Participants
	MessageParticipants *[]MessageParticipantForResponse `json:"messageParticipants,omitempty"`

	// Name Workspace member name
	Name *struct {
		FirstName *string `json:"firstName,omitempty"`
		LastName  *string `json:"lastName,omitempty"`
	} `json:"name,omitempty"`

	// NumberFormat User's preferred number format
	NumberFormat *WorkspaceMemberForResponseNumberFormat `json:"numberFormat,omitempty"`

	// Position Workspace member position
	Position *float32 `json:"position,omitempty"`

	// TimeFormat User's preferred time format
	TimeFormat *WorkspaceMemberForResponseTimeFormat `json:"timeFormat,omitempty"`

	// TimeZone User time zone
	TimeZone *string `json:"timeZone,omitempty"`

	// TimelineActivities Events linked to the workspace member
	TimelineActivities *[]TimelineActivityForResponse `json:"timelineActivities,omitempty"`

	// UpdatedAt Last time the record was changed
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// UserEmail Related user email address
	UserEmail *string `json:"userEmail,omitempty"`

	// UserId Associated User Id
	UserId *openapi_types.UUID `json:"userId,omitempty"`
}

// WorkspaceMemberForResponseDateFormat User's preferred date format
type WorkspaceMemberForResponseDateFormat string

// WorkspaceMemberForResponseNumberFormat User's preferred number format
type WorkspaceMemberForResponseNumberFormat string

// WorkspaceMemberForResponseTimeFormat User's preferred time format
type WorkspaceMemberForResponseTimeFormat string

// WorkspaceMemberForUpdate A workspace member
type WorkspaceMemberForUpdate struct {
	// AvatarUrl Workspace member avatar
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// CalendarStartDay User's preferred start day of the week
	CalendarStartDay *int `json:"calendarStartDay,omitempty"`

	// ColorScheme Preferred color scheme
	ColorScheme *string `json:"colorScheme,omitempty"`

	// DateFormat User's preferred date format
	DateFormat *WorkspaceMemberForUpdateDateFormat `json:"dateFormat,omitempty"`

	// Locale Preferred language
	Locale *string `json:"locale,omitempty"`

	// Name Workspace member name
	Name *struct {
		FirstName *string `json:"firstName,omitempty"`
		LastName  *string `json:"lastName,omitempty"`
	} `json:"name,omitempty"`

	// NumberFormat User's preferred number format
	NumberFormat *WorkspaceMemberForUpdateNumberFormat `json:"numberFormat,omitempty"`

	// Position Workspace member position
	Position *float32 `json:"position,omitempty"`

	// TimeFormat User's preferred time format
	TimeFormat *WorkspaceMemberForUpdateTimeFormat `json:"timeFormat,omitempty"`

	// TimeZone User time zone
	TimeZone *string `json:"timeZone,omitempty"`

	// UserEmail Related user email address
	UserEmail *string `json:"userEmail,omitempty"`

	// UserId Associated User Id
	UserId *openapi_types.UUID `json:"userId,omitempty"`
}

// WorkspaceMemberForUpdateDateFormat User's preferred date format
type WorkspaceMemberForUpdateDateFormat string

// WorkspaceMemberForUpdateNumberFormat User's preferred number format
type WorkspaceMemberForUpdateNumberFormat string

// WorkspaceMemberForUpdateTimeFormat User's preferred time format
type WorkspaceMemberForUpdateTimeFormat string

// Depth defines model for depth.
type Depth int

// EndingBefore defines model for endingBefore.
type EndingBefore = string

// Filter defines model for filter.
type Filter = string

// IdPath defines model for idPath.
type IdPath = openapi_types.UUID

// Limit defines model for limit.
type Limit = int

// OrderBy defines model for orderBy.
type OrderBy = string

// SoftDelete defines model for softDelete.
type SoftDelete = bool

// StartingAfter defines model for startingAfter.
type StartingAfter = string

// Upsert defines model for upsert.
type Upsert = bool

// N400 defines model for 400.
type N400 struct {
	Error      *string   `json:"error,omitempty"`
	Messages   *[]string `json:"messages,omitempty"`
	StatusCode *float32  `json:"statusCode,omitempty"`
}

// N401 defines model for 401.
type N401 struct {
	Error      *string  `json:"error,omitempty"`
	Message    *string  `json:"message,omitempty"`
	StatusCode *float32 `json:"statusCode,omitempty"`
}

// DeleteManyAttachmentsParams defines parameters for DeleteManyAttachments.
type DeleteManyAttachmentsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyAttachmentsParams defines parameters for FindManyAttachments.
type FindManyAttachmentsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyAttachmentsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyAttachmentsParamsDepth defines parameters for FindManyAttachments.
type FindManyAttachmentsParamsDepth int

// UpdateManyAttachmentsParams defines parameters for UpdateManyAttachments.
type UpdateManyAttachmentsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyAttachmentsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyAttachmentsParamsDepth defines parameters for UpdateManyAttachments.
type UpdateManyAttachmentsParamsDepth int

// CreateOneAttachmentParams defines parameters for CreateOneAttachment.
type CreateOneAttachmentParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneAttachmentParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneAttachmentParamsDepth defines parameters for CreateOneAttachment.
type CreateOneAttachmentParamsDepth int

// FindAttachmentDuplicatesJSONBody defines parameters for FindAttachmentDuplicates.
type FindAttachmentDuplicatesJSONBody struct {
	Data *[]Attachment         `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindAttachmentDuplicatesParams defines parameters for FindAttachmentDuplicates.
type FindAttachmentDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindAttachmentDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindAttachmentDuplicatesParamsDepth defines parameters for FindAttachmentDuplicates.
type FindAttachmentDuplicatesParamsDepth int

// MergeManyAttachmentsJSONBody defines parameters for MergeManyAttachments.
type MergeManyAttachmentsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyAttachmentsParams defines parameters for MergeManyAttachments.
type MergeManyAttachmentsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyAttachmentsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyAttachmentsParamsDepth defines parameters for MergeManyAttachments.
type MergeManyAttachmentsParamsDepth int

// DeleteOneAttachmentParams defines parameters for DeleteOneAttachment.
type DeleteOneAttachmentParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneAttachmentParams defines parameters for FindOneAttachment.
type FindOneAttachmentParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneAttachmentParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneAttachmentParamsDepth defines parameters for FindOneAttachment.
type FindOneAttachmentParamsDepth int

// UpdateOneAttachmentParams defines parameters for UpdateOneAttachment.
type UpdateOneAttachmentParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneAttachmentParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneAttachmentParamsDepth defines parameters for UpdateOneAttachment.
type UpdateOneAttachmentParamsDepth int

// CreateManyAttachmentsJSONBody defines parameters for CreateManyAttachments.
type CreateManyAttachmentsJSONBody = []Attachment

// CreateManyAttachmentsParams defines parameters for CreateManyAttachments.
type CreateManyAttachmentsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyAttachmentsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyAttachmentsParamsDepth defines parameters for CreateManyAttachments.
type CreateManyAttachmentsParamsDepth int

// CreateManyBlocklistsJSONBody defines parameters for CreateManyBlocklists.
type CreateManyBlocklistsJSONBody = []Blocklist

// CreateManyBlocklistsParams defines parameters for CreateManyBlocklists.
type CreateManyBlocklistsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyBlocklistsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyBlocklistsParamsDepth defines parameters for CreateManyBlocklists.
type CreateManyBlocklistsParamsDepth int

// CreateManyCalendarChannelEventAssociationsJSONBody defines parameters for CreateManyCalendarChannelEventAssociations.
type CreateManyCalendarChannelEventAssociationsJSONBody = []CalendarChannelEventAssociation

// CreateManyCalendarChannelEventAssociationsParams defines parameters for CreateManyCalendarChannelEventAssociations.
type CreateManyCalendarChannelEventAssociationsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyCalendarChannelEventAssociationsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyCalendarChannelEventAssociationsParamsDepth defines parameters for CreateManyCalendarChannelEventAssociations.
type CreateManyCalendarChannelEventAssociationsParamsDepth int

// CreateManyCalendarChannelsJSONBody defines parameters for CreateManyCalendarChannels.
type CreateManyCalendarChannelsJSONBody = []CalendarChannel

// CreateManyCalendarChannelsParams defines parameters for CreateManyCalendarChannels.
type CreateManyCalendarChannelsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyCalendarChannelsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyCalendarChannelsParamsDepth defines parameters for CreateManyCalendarChannels.
type CreateManyCalendarChannelsParamsDepth int

// CreateManyCalendarEventParticipantsJSONBody defines parameters for CreateManyCalendarEventParticipants.
type CreateManyCalendarEventParticipantsJSONBody = []CalendarEventParticipant

// CreateManyCalendarEventParticipantsParams defines parameters for CreateManyCalendarEventParticipants.
type CreateManyCalendarEventParticipantsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyCalendarEventParticipantsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyCalendarEventParticipantsParamsDepth defines parameters for CreateManyCalendarEventParticipants.
type CreateManyCalendarEventParticipantsParamsDepth int

// CreateManyCalendarEventsJSONBody defines parameters for CreateManyCalendarEvents.
type CreateManyCalendarEventsJSONBody = []CalendarEvent

// CreateManyCalendarEventsParams defines parameters for CreateManyCalendarEvents.
type CreateManyCalendarEventsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyCalendarEventsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyCalendarEventsParamsDepth defines parameters for CreateManyCalendarEvents.
type CreateManyCalendarEventsParamsDepth int

// CreateManyCompaniesJSONBody defines parameters for CreateManyCompanies.
type CreateManyCompaniesJSONBody = []Company

// CreateManyCompaniesParams defines parameters for CreateManyCompanies.
type CreateManyCompaniesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyCompaniesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyCompaniesParamsDepth defines parameters for CreateManyCompanies.
type CreateManyCompaniesParamsDepth int

// CreateManyConnectedAccountsJSONBody defines parameters for CreateManyConnectedAccounts.
type CreateManyConnectedAccountsJSONBody = []ConnectedAccount

// CreateManyConnectedAccountsParams defines parameters for CreateManyConnectedAccounts.
type CreateManyConnectedAccountsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyConnectedAccountsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyConnectedAccountsParamsDepth defines parameters for CreateManyConnectedAccounts.
type CreateManyConnectedAccountsParamsDepth int

// CreateManyDashboardsJSONBody defines parameters for CreateManyDashboards.
type CreateManyDashboardsJSONBody = []Dashboard

// CreateManyDashboardsParams defines parameters for CreateManyDashboards.
type CreateManyDashboardsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyDashboardsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyDashboardsParamsDepth defines parameters for CreateManyDashboards.
type CreateManyDashboardsParamsDepth int

// CreateManyFavoriteFoldersJSONBody defines parameters for CreateManyFavoriteFolders.
type CreateManyFavoriteFoldersJSONBody = []FavoriteFolder

// CreateManyFavoriteFoldersParams defines parameters for CreateManyFavoriteFolders.
type CreateManyFavoriteFoldersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyFavoriteFoldersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyFavoriteFoldersParamsDepth defines parameters for CreateManyFavoriteFolders.
type CreateManyFavoriteFoldersParamsDepth int

// CreateManyFavoritesJSONBody defines parameters for CreateManyFavorites.
type CreateManyFavoritesJSONBody = []Favorite

// CreateManyFavoritesParams defines parameters for CreateManyFavorites.
type CreateManyFavoritesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyFavoritesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyFavoritesParamsDepth defines parameters for CreateManyFavorites.
type CreateManyFavoritesParamsDepth int

// CreateManyMessageChannelMessageAssociationsJSONBody defines parameters for CreateManyMessageChannelMessageAssociations.
type CreateManyMessageChannelMessageAssociationsJSONBody = []MessageChannelMessageAssociation

// CreateManyMessageChannelMessageAssociationsParams defines parameters for CreateManyMessageChannelMessageAssociations.
type CreateManyMessageChannelMessageAssociationsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyMessageChannelMessageAssociationsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyMessageChannelMessageAssociationsParamsDepth defines parameters for CreateManyMessageChannelMessageAssociations.
type CreateManyMessageChannelMessageAssociationsParamsDepth int

// CreateManyMessageChannelsJSONBody defines parameters for CreateManyMessageChannels.
type CreateManyMessageChannelsJSONBody = []MessageChannel

// CreateManyMessageChannelsParams defines parameters for CreateManyMessageChannels.
type CreateManyMessageChannelsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyMessageChannelsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyMessageChannelsParamsDepth defines parameters for CreateManyMessageChannels.
type CreateManyMessageChannelsParamsDepth int

// CreateManyMessageFoldersJSONBody defines parameters for CreateManyMessageFolders.
type CreateManyMessageFoldersJSONBody = []MessageFolder

// CreateManyMessageFoldersParams defines parameters for CreateManyMessageFolders.
type CreateManyMessageFoldersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyMessageFoldersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyMessageFoldersParamsDepth defines parameters for CreateManyMessageFolders.
type CreateManyMessageFoldersParamsDepth int

// CreateManyMessageParticipantsJSONBody defines parameters for CreateManyMessageParticipants.
type CreateManyMessageParticipantsJSONBody = []MessageParticipant

// CreateManyMessageParticipantsParams defines parameters for CreateManyMessageParticipants.
type CreateManyMessageParticipantsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyMessageParticipantsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyMessageParticipantsParamsDepth defines parameters for CreateManyMessageParticipants.
type CreateManyMessageParticipantsParamsDepth int

// CreateManyMessageThreadsJSONBody defines parameters for CreateManyMessageThreads.
type CreateManyMessageThreadsJSONBody = []MessageThread

// CreateManyMessageThreadsParams defines parameters for CreateManyMessageThreads.
type CreateManyMessageThreadsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyMessageThreadsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyMessageThreadsParamsDepth defines parameters for CreateManyMessageThreads.
type CreateManyMessageThreadsParamsDepth int

// CreateManyMessagesJSONBody defines parameters for CreateManyMessages.
type CreateManyMessagesJSONBody = []Message

// CreateManyMessagesParams defines parameters for CreateManyMessages.
type CreateManyMessagesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyMessagesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyMessagesParamsDepth defines parameters for CreateManyMessages.
type CreateManyMessagesParamsDepth int

// CreateManyNoteTargetsJSONBody defines parameters for CreateManyNoteTargets.
type CreateManyNoteTargetsJSONBody = []NoteTarget

// CreateManyNoteTargetsParams defines parameters for CreateManyNoteTargets.
type CreateManyNoteTargetsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyNoteTargetsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyNoteTargetsParamsDepth defines parameters for CreateManyNoteTargets.
type CreateManyNoteTargetsParamsDepth int

// CreateManyNotesJSONBody defines parameters for CreateManyNotes.
type CreateManyNotesJSONBody = []Note

// CreateManyNotesParams defines parameters for CreateManyNotes.
type CreateManyNotesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyNotesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyNotesParamsDepth defines parameters for CreateManyNotes.
type CreateManyNotesParamsDepth int

// CreateManyOpportunitiesJSONBody defines parameters for CreateManyOpportunities.
type CreateManyOpportunitiesJSONBody = []Opportunity

// CreateManyOpportunitiesParams defines parameters for CreateManyOpportunities.
type CreateManyOpportunitiesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyOpportunitiesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyOpportunitiesParamsDepth defines parameters for CreateManyOpportunities.
type CreateManyOpportunitiesParamsDepth int

// CreateManyPeopleJSONBody defines parameters for CreateManyPeople.
type CreateManyPeopleJSONBody = []Person

// CreateManyPeopleParams defines parameters for CreateManyPeople.
type CreateManyPeopleParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyPeopleParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyPeopleParamsDepth defines parameters for CreateManyPeople.
type CreateManyPeopleParamsDepth int

// CreateManyTaskTargetsJSONBody defines parameters for CreateManyTaskTargets.
type CreateManyTaskTargetsJSONBody = []TaskTarget

// CreateManyTaskTargetsParams defines parameters for CreateManyTaskTargets.
type CreateManyTaskTargetsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyTaskTargetsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyTaskTargetsParamsDepth defines parameters for CreateManyTaskTargets.
type CreateManyTaskTargetsParamsDepth int

// CreateManyTasksJSONBody defines parameters for CreateManyTasks.
type CreateManyTasksJSONBody = []Task

// CreateManyTasksParams defines parameters for CreateManyTasks.
type CreateManyTasksParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyTasksParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyTasksParamsDepth defines parameters for CreateManyTasks.
type CreateManyTasksParamsDepth int

// CreateManyTimelineActivitiesJSONBody defines parameters for CreateManyTimelineActivities.
type CreateManyTimelineActivitiesJSONBody = []TimelineActivity

// CreateManyTimelineActivitiesParams defines parameters for CreateManyTimelineActivities.
type CreateManyTimelineActivitiesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyTimelineActivitiesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyTimelineActivitiesParamsDepth defines parameters for CreateManyTimelineActivities.
type CreateManyTimelineActivitiesParamsDepth int

// CreateManyWorkflowAutomatedTriggersJSONBody defines parameters for CreateManyWorkflowAutomatedTriggers.
type CreateManyWorkflowAutomatedTriggersJSONBody = []WorkflowAutomatedTrigger

// CreateManyWorkflowAutomatedTriggersParams defines parameters for CreateManyWorkflowAutomatedTriggers.
type CreateManyWorkflowAutomatedTriggersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyWorkflowAutomatedTriggersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyWorkflowAutomatedTriggersParamsDepth defines parameters for CreateManyWorkflowAutomatedTriggers.
type CreateManyWorkflowAutomatedTriggersParamsDepth int

// CreateManyWorkflowRunsJSONBody defines parameters for CreateManyWorkflowRuns.
type CreateManyWorkflowRunsJSONBody = []WorkflowRun

// CreateManyWorkflowRunsParams defines parameters for CreateManyWorkflowRuns.
type CreateManyWorkflowRunsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyWorkflowRunsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyWorkflowRunsParamsDepth defines parameters for CreateManyWorkflowRuns.
type CreateManyWorkflowRunsParamsDepth int

// CreateManyWorkflowVersionsJSONBody defines parameters for CreateManyWorkflowVersions.
type CreateManyWorkflowVersionsJSONBody = []WorkflowVersion

// CreateManyWorkflowVersionsParams defines parameters for CreateManyWorkflowVersions.
type CreateManyWorkflowVersionsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyWorkflowVersionsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyWorkflowVersionsParamsDepth defines parameters for CreateManyWorkflowVersions.
type CreateManyWorkflowVersionsParamsDepth int

// CreateManyWorkflowsJSONBody defines parameters for CreateManyWorkflows.
type CreateManyWorkflowsJSONBody = []Workflow

// CreateManyWorkflowsParams defines parameters for CreateManyWorkflows.
type CreateManyWorkflowsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyWorkflowsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyWorkflowsParamsDepth defines parameters for CreateManyWorkflows.
type CreateManyWorkflowsParamsDepth int

// CreateManyWorkspaceMembersJSONBody defines parameters for CreateManyWorkspaceMembers.
type CreateManyWorkspaceMembersJSONBody = []WorkspaceMember

// CreateManyWorkspaceMembersParams defines parameters for CreateManyWorkspaceMembers.
type CreateManyWorkspaceMembersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateManyWorkspaceMembersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateManyWorkspaceMembersParamsDepth defines parameters for CreateManyWorkspaceMembers.
type CreateManyWorkspaceMembersParamsDepth int

// DeleteManyBlocklistsParams defines parameters for DeleteManyBlocklists.
type DeleteManyBlocklistsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyBlocklistsParams defines parameters for FindManyBlocklists.
type FindManyBlocklistsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyBlocklistsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyBlocklistsParamsDepth defines parameters for FindManyBlocklists.
type FindManyBlocklistsParamsDepth int

// UpdateManyBlocklistsParams defines parameters for UpdateManyBlocklists.
type UpdateManyBlocklistsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyBlocklistsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyBlocklistsParamsDepth defines parameters for UpdateManyBlocklists.
type UpdateManyBlocklistsParamsDepth int

// CreateOneBlocklistParams defines parameters for CreateOneBlocklist.
type CreateOneBlocklistParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneBlocklistParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneBlocklistParamsDepth defines parameters for CreateOneBlocklist.
type CreateOneBlocklistParamsDepth int

// FindBlocklistDuplicatesJSONBody defines parameters for FindBlocklistDuplicates.
type FindBlocklistDuplicatesJSONBody struct {
	Data *[]Blocklist          `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindBlocklistDuplicatesParams defines parameters for FindBlocklistDuplicates.
type FindBlocklistDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindBlocklistDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindBlocklistDuplicatesParamsDepth defines parameters for FindBlocklistDuplicates.
type FindBlocklistDuplicatesParamsDepth int

// MergeManyBlocklistsJSONBody defines parameters for MergeManyBlocklists.
type MergeManyBlocklistsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyBlocklistsParams defines parameters for MergeManyBlocklists.
type MergeManyBlocklistsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyBlocklistsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyBlocklistsParamsDepth defines parameters for MergeManyBlocklists.
type MergeManyBlocklistsParamsDepth int

// DeleteOneBlocklistParams defines parameters for DeleteOneBlocklist.
type DeleteOneBlocklistParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneBlocklistParams defines parameters for FindOneBlocklist.
type FindOneBlocklistParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneBlocklistParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneBlocklistParamsDepth defines parameters for FindOneBlocklist.
type FindOneBlocklistParamsDepth int

// UpdateOneBlocklistParams defines parameters for UpdateOneBlocklist.
type UpdateOneBlocklistParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneBlocklistParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneBlocklistParamsDepth defines parameters for UpdateOneBlocklist.
type UpdateOneBlocklistParamsDepth int

// DeleteManyCalendarChannelEventAssociationsParams defines parameters for DeleteManyCalendarChannelEventAssociations.
type DeleteManyCalendarChannelEventAssociationsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyCalendarChannelEventAssociationsParams defines parameters for FindManyCalendarChannelEventAssociations.
type FindManyCalendarChannelEventAssociationsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyCalendarChannelEventAssociationsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyCalendarChannelEventAssociationsParamsDepth defines parameters for FindManyCalendarChannelEventAssociations.
type FindManyCalendarChannelEventAssociationsParamsDepth int

// UpdateManyCalendarChannelEventAssociationsParams defines parameters for UpdateManyCalendarChannelEventAssociations.
type UpdateManyCalendarChannelEventAssociationsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyCalendarChannelEventAssociationsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyCalendarChannelEventAssociationsParamsDepth defines parameters for UpdateManyCalendarChannelEventAssociations.
type UpdateManyCalendarChannelEventAssociationsParamsDepth int

// CreateOneCalendarChannelEventAssociationParams defines parameters for CreateOneCalendarChannelEventAssociation.
type CreateOneCalendarChannelEventAssociationParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneCalendarChannelEventAssociationParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneCalendarChannelEventAssociationParamsDepth defines parameters for CreateOneCalendarChannelEventAssociation.
type CreateOneCalendarChannelEventAssociationParamsDepth int

// FindCalendarChannelEventAssociationDuplicatesJSONBody defines parameters for FindCalendarChannelEventAssociationDuplicates.
type FindCalendarChannelEventAssociationDuplicatesJSONBody struct {
	Data *[]CalendarChannelEventAssociation `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID              `json:"ids,omitempty"`
}

// FindCalendarChannelEventAssociationDuplicatesParams defines parameters for FindCalendarChannelEventAssociationDuplicates.
type FindCalendarChannelEventAssociationDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindCalendarChannelEventAssociationDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindCalendarChannelEventAssociationDuplicatesParamsDepth defines parameters for FindCalendarChannelEventAssociationDuplicates.
type FindCalendarChannelEventAssociationDuplicatesParamsDepth int

// MergeManyCalendarChannelEventAssociationsJSONBody defines parameters for MergeManyCalendarChannelEventAssociations.
type MergeManyCalendarChannelEventAssociationsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyCalendarChannelEventAssociationsParams defines parameters for MergeManyCalendarChannelEventAssociations.
type MergeManyCalendarChannelEventAssociationsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyCalendarChannelEventAssociationsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyCalendarChannelEventAssociationsParamsDepth defines parameters for MergeManyCalendarChannelEventAssociations.
type MergeManyCalendarChannelEventAssociationsParamsDepth int

// DeleteOneCalendarChannelEventAssociationParams defines parameters for DeleteOneCalendarChannelEventAssociation.
type DeleteOneCalendarChannelEventAssociationParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneCalendarChannelEventAssociationParams defines parameters for FindOneCalendarChannelEventAssociation.
type FindOneCalendarChannelEventAssociationParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneCalendarChannelEventAssociationParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneCalendarChannelEventAssociationParamsDepth defines parameters for FindOneCalendarChannelEventAssociation.
type FindOneCalendarChannelEventAssociationParamsDepth int

// UpdateOneCalendarChannelEventAssociationParams defines parameters for UpdateOneCalendarChannelEventAssociation.
type UpdateOneCalendarChannelEventAssociationParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneCalendarChannelEventAssociationParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneCalendarChannelEventAssociationParamsDepth defines parameters for UpdateOneCalendarChannelEventAssociation.
type UpdateOneCalendarChannelEventAssociationParamsDepth int

// DeleteManyCalendarChannelsParams defines parameters for DeleteManyCalendarChannels.
type DeleteManyCalendarChannelsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyCalendarChannelsParams defines parameters for FindManyCalendarChannels.
type FindManyCalendarChannelsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyCalendarChannelsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyCalendarChannelsParamsDepth defines parameters for FindManyCalendarChannels.
type FindManyCalendarChannelsParamsDepth int

// UpdateManyCalendarChannelsParams defines parameters for UpdateManyCalendarChannels.
type UpdateManyCalendarChannelsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyCalendarChannelsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyCalendarChannelsParamsDepth defines parameters for UpdateManyCalendarChannels.
type UpdateManyCalendarChannelsParamsDepth int

// CreateOneCalendarChannelParams defines parameters for CreateOneCalendarChannel.
type CreateOneCalendarChannelParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneCalendarChannelParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneCalendarChannelParamsDepth defines parameters for CreateOneCalendarChannel.
type CreateOneCalendarChannelParamsDepth int

// FindCalendarChannelDuplicatesJSONBody defines parameters for FindCalendarChannelDuplicates.
type FindCalendarChannelDuplicatesJSONBody struct {
	Data *[]CalendarChannel    `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindCalendarChannelDuplicatesParams defines parameters for FindCalendarChannelDuplicates.
type FindCalendarChannelDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindCalendarChannelDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindCalendarChannelDuplicatesParamsDepth defines parameters for FindCalendarChannelDuplicates.
type FindCalendarChannelDuplicatesParamsDepth int

// MergeManyCalendarChannelsJSONBody defines parameters for MergeManyCalendarChannels.
type MergeManyCalendarChannelsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyCalendarChannelsParams defines parameters for MergeManyCalendarChannels.
type MergeManyCalendarChannelsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyCalendarChannelsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyCalendarChannelsParamsDepth defines parameters for MergeManyCalendarChannels.
type MergeManyCalendarChannelsParamsDepth int

// DeleteOneCalendarChannelParams defines parameters for DeleteOneCalendarChannel.
type DeleteOneCalendarChannelParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneCalendarChannelParams defines parameters for FindOneCalendarChannel.
type FindOneCalendarChannelParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneCalendarChannelParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneCalendarChannelParamsDepth defines parameters for FindOneCalendarChannel.
type FindOneCalendarChannelParamsDepth int

// UpdateOneCalendarChannelParams defines parameters for UpdateOneCalendarChannel.
type UpdateOneCalendarChannelParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneCalendarChannelParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneCalendarChannelParamsDepth defines parameters for UpdateOneCalendarChannel.
type UpdateOneCalendarChannelParamsDepth int

// DeleteManyCalendarEventParticipantsParams defines parameters for DeleteManyCalendarEventParticipants.
type DeleteManyCalendarEventParticipantsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyCalendarEventParticipantsParams defines parameters for FindManyCalendarEventParticipants.
type FindManyCalendarEventParticipantsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyCalendarEventParticipantsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyCalendarEventParticipantsParamsDepth defines parameters for FindManyCalendarEventParticipants.
type FindManyCalendarEventParticipantsParamsDepth int

// UpdateManyCalendarEventParticipantsParams defines parameters for UpdateManyCalendarEventParticipants.
type UpdateManyCalendarEventParticipantsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyCalendarEventParticipantsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyCalendarEventParticipantsParamsDepth defines parameters for UpdateManyCalendarEventParticipants.
type UpdateManyCalendarEventParticipantsParamsDepth int

// CreateOneCalendarEventParticipantParams defines parameters for CreateOneCalendarEventParticipant.
type CreateOneCalendarEventParticipantParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneCalendarEventParticipantParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneCalendarEventParticipantParamsDepth defines parameters for CreateOneCalendarEventParticipant.
type CreateOneCalendarEventParticipantParamsDepth int

// FindCalendarEventParticipantDuplicatesJSONBody defines parameters for FindCalendarEventParticipantDuplicates.
type FindCalendarEventParticipantDuplicatesJSONBody struct {
	Data *[]CalendarEventParticipant `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID       `json:"ids,omitempty"`
}

// FindCalendarEventParticipantDuplicatesParams defines parameters for FindCalendarEventParticipantDuplicates.
type FindCalendarEventParticipantDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindCalendarEventParticipantDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindCalendarEventParticipantDuplicatesParamsDepth defines parameters for FindCalendarEventParticipantDuplicates.
type FindCalendarEventParticipantDuplicatesParamsDepth int

// MergeManyCalendarEventParticipantsJSONBody defines parameters for MergeManyCalendarEventParticipants.
type MergeManyCalendarEventParticipantsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyCalendarEventParticipantsParams defines parameters for MergeManyCalendarEventParticipants.
type MergeManyCalendarEventParticipantsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyCalendarEventParticipantsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyCalendarEventParticipantsParamsDepth defines parameters for MergeManyCalendarEventParticipants.
type MergeManyCalendarEventParticipantsParamsDepth int

// DeleteOneCalendarEventParticipantParams defines parameters for DeleteOneCalendarEventParticipant.
type DeleteOneCalendarEventParticipantParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneCalendarEventParticipantParams defines parameters for FindOneCalendarEventParticipant.
type FindOneCalendarEventParticipantParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneCalendarEventParticipantParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneCalendarEventParticipantParamsDepth defines parameters for FindOneCalendarEventParticipant.
type FindOneCalendarEventParticipantParamsDepth int

// UpdateOneCalendarEventParticipantParams defines parameters for UpdateOneCalendarEventParticipant.
type UpdateOneCalendarEventParticipantParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneCalendarEventParticipantParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneCalendarEventParticipantParamsDepth defines parameters for UpdateOneCalendarEventParticipant.
type UpdateOneCalendarEventParticipantParamsDepth int

// DeleteManyCalendarEventsParams defines parameters for DeleteManyCalendarEvents.
type DeleteManyCalendarEventsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyCalendarEventsParams defines parameters for FindManyCalendarEvents.
type FindManyCalendarEventsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyCalendarEventsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyCalendarEventsParamsDepth defines parameters for FindManyCalendarEvents.
type FindManyCalendarEventsParamsDepth int

// UpdateManyCalendarEventsParams defines parameters for UpdateManyCalendarEvents.
type UpdateManyCalendarEventsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyCalendarEventsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyCalendarEventsParamsDepth defines parameters for UpdateManyCalendarEvents.
type UpdateManyCalendarEventsParamsDepth int

// CreateOneCalendarEventParams defines parameters for CreateOneCalendarEvent.
type CreateOneCalendarEventParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneCalendarEventParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneCalendarEventParamsDepth defines parameters for CreateOneCalendarEvent.
type CreateOneCalendarEventParamsDepth int

// FindCalendarEventDuplicatesJSONBody defines parameters for FindCalendarEventDuplicates.
type FindCalendarEventDuplicatesJSONBody struct {
	Data *[]CalendarEvent      `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindCalendarEventDuplicatesParams defines parameters for FindCalendarEventDuplicates.
type FindCalendarEventDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindCalendarEventDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindCalendarEventDuplicatesParamsDepth defines parameters for FindCalendarEventDuplicates.
type FindCalendarEventDuplicatesParamsDepth int

// MergeManyCalendarEventsJSONBody defines parameters for MergeManyCalendarEvents.
type MergeManyCalendarEventsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyCalendarEventsParams defines parameters for MergeManyCalendarEvents.
type MergeManyCalendarEventsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyCalendarEventsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyCalendarEventsParamsDepth defines parameters for MergeManyCalendarEvents.
type MergeManyCalendarEventsParamsDepth int

// DeleteOneCalendarEventParams defines parameters for DeleteOneCalendarEvent.
type DeleteOneCalendarEventParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneCalendarEventParams defines parameters for FindOneCalendarEvent.
type FindOneCalendarEventParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneCalendarEventParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneCalendarEventParamsDepth defines parameters for FindOneCalendarEvent.
type FindOneCalendarEventParamsDepth int

// UpdateOneCalendarEventParams defines parameters for UpdateOneCalendarEvent.
type UpdateOneCalendarEventParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneCalendarEventParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneCalendarEventParamsDepth defines parameters for UpdateOneCalendarEvent.
type UpdateOneCalendarEventParamsDepth int

// DeleteManyCompaniesParams defines parameters for DeleteManyCompanies.
type DeleteManyCompaniesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyCompaniesParams defines parameters for FindManyCompanies.
type FindManyCompaniesParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyCompaniesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyCompaniesParamsDepth defines parameters for FindManyCompanies.
type FindManyCompaniesParamsDepth int

// UpdateManyCompaniesParams defines parameters for UpdateManyCompanies.
type UpdateManyCompaniesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyCompaniesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyCompaniesParamsDepth defines parameters for UpdateManyCompanies.
type UpdateManyCompaniesParamsDepth int

// CreateOneCompanyParams defines parameters for CreateOneCompany.
type CreateOneCompanyParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneCompanyParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneCompanyParamsDepth defines parameters for CreateOneCompany.
type CreateOneCompanyParamsDepth int

// FindCompanyDuplicatesJSONBody defines parameters for FindCompanyDuplicates.
type FindCompanyDuplicatesJSONBody struct {
	Data *[]Company            `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindCompanyDuplicatesParams defines parameters for FindCompanyDuplicates.
type FindCompanyDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindCompanyDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindCompanyDuplicatesParamsDepth defines parameters for FindCompanyDuplicates.
type FindCompanyDuplicatesParamsDepth int

// MergeManyCompaniesJSONBody defines parameters for MergeManyCompanies.
type MergeManyCompaniesJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyCompaniesParams defines parameters for MergeManyCompanies.
type MergeManyCompaniesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyCompaniesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyCompaniesParamsDepth defines parameters for MergeManyCompanies.
type MergeManyCompaniesParamsDepth int

// DeleteOneCompanyParams defines parameters for DeleteOneCompany.
type DeleteOneCompanyParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneCompanyParams defines parameters for FindOneCompany.
type FindOneCompanyParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneCompanyParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneCompanyParamsDepth defines parameters for FindOneCompany.
type FindOneCompanyParamsDepth int

// UpdateOneCompanyParams defines parameters for UpdateOneCompany.
type UpdateOneCompanyParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneCompanyParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneCompanyParamsDepth defines parameters for UpdateOneCompany.
type UpdateOneCompanyParamsDepth int

// DeleteManyConnectedAccountsParams defines parameters for DeleteManyConnectedAccounts.
type DeleteManyConnectedAccountsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyConnectedAccountsParams defines parameters for FindManyConnectedAccounts.
type FindManyConnectedAccountsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyConnectedAccountsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyConnectedAccountsParamsDepth defines parameters for FindManyConnectedAccounts.
type FindManyConnectedAccountsParamsDepth int

// UpdateManyConnectedAccountsParams defines parameters for UpdateManyConnectedAccounts.
type UpdateManyConnectedAccountsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyConnectedAccountsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyConnectedAccountsParamsDepth defines parameters for UpdateManyConnectedAccounts.
type UpdateManyConnectedAccountsParamsDepth int

// CreateOneConnectedAccountParams defines parameters for CreateOneConnectedAccount.
type CreateOneConnectedAccountParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneConnectedAccountParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneConnectedAccountParamsDepth defines parameters for CreateOneConnectedAccount.
type CreateOneConnectedAccountParamsDepth int

// FindConnectedAccountDuplicatesJSONBody defines parameters for FindConnectedAccountDuplicates.
type FindConnectedAccountDuplicatesJSONBody struct {
	Data *[]ConnectedAccount   `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindConnectedAccountDuplicatesParams defines parameters for FindConnectedAccountDuplicates.
type FindConnectedAccountDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindConnectedAccountDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindConnectedAccountDuplicatesParamsDepth defines parameters for FindConnectedAccountDuplicates.
type FindConnectedAccountDuplicatesParamsDepth int

// MergeManyConnectedAccountsJSONBody defines parameters for MergeManyConnectedAccounts.
type MergeManyConnectedAccountsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyConnectedAccountsParams defines parameters for MergeManyConnectedAccounts.
type MergeManyConnectedAccountsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyConnectedAccountsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyConnectedAccountsParamsDepth defines parameters for MergeManyConnectedAccounts.
type MergeManyConnectedAccountsParamsDepth int

// DeleteOneConnectedAccountParams defines parameters for DeleteOneConnectedAccount.
type DeleteOneConnectedAccountParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneConnectedAccountParams defines parameters for FindOneConnectedAccount.
type FindOneConnectedAccountParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneConnectedAccountParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneConnectedAccountParamsDepth defines parameters for FindOneConnectedAccount.
type FindOneConnectedAccountParamsDepth int

// UpdateOneConnectedAccountParams defines parameters for UpdateOneConnectedAccount.
type UpdateOneConnectedAccountParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneConnectedAccountParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneConnectedAccountParamsDepth defines parameters for UpdateOneConnectedAccount.
type UpdateOneConnectedAccountParamsDepth int

// DeleteManyDashboardsParams defines parameters for DeleteManyDashboards.
type DeleteManyDashboardsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyDashboardsParams defines parameters for FindManyDashboards.
type FindManyDashboardsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyDashboardsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyDashboardsParamsDepth defines parameters for FindManyDashboards.
type FindManyDashboardsParamsDepth int

// UpdateManyDashboardsParams defines parameters for UpdateManyDashboards.
type UpdateManyDashboardsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyDashboardsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyDashboardsParamsDepth defines parameters for UpdateManyDashboards.
type UpdateManyDashboardsParamsDepth int

// CreateOneDashboardParams defines parameters for CreateOneDashboard.
type CreateOneDashboardParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneDashboardParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneDashboardParamsDepth defines parameters for CreateOneDashboard.
type CreateOneDashboardParamsDepth int

// FindDashboardDuplicatesJSONBody defines parameters for FindDashboardDuplicates.
type FindDashboardDuplicatesJSONBody struct {
	Data *[]Dashboard          `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindDashboardDuplicatesParams defines parameters for FindDashboardDuplicates.
type FindDashboardDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindDashboardDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindDashboardDuplicatesParamsDepth defines parameters for FindDashboardDuplicates.
type FindDashboardDuplicatesParamsDepth int

// MergeManyDashboardsJSONBody defines parameters for MergeManyDashboards.
type MergeManyDashboardsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyDashboardsParams defines parameters for MergeManyDashboards.
type MergeManyDashboardsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyDashboardsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyDashboardsParamsDepth defines parameters for MergeManyDashboards.
type MergeManyDashboardsParamsDepth int

// DeleteOneDashboardParams defines parameters for DeleteOneDashboard.
type DeleteOneDashboardParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneDashboardParams defines parameters for FindOneDashboard.
type FindOneDashboardParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneDashboardParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneDashboardParamsDepth defines parameters for FindOneDashboard.
type FindOneDashboardParamsDepth int

// UpdateOneDashboardParams defines parameters for UpdateOneDashboard.
type UpdateOneDashboardParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneDashboardParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneDashboardParamsDepth defines parameters for UpdateOneDashboard.
type UpdateOneDashboardParamsDepth int

// DeleteManyFavoriteFoldersParams defines parameters for DeleteManyFavoriteFolders.
type DeleteManyFavoriteFoldersParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyFavoriteFoldersParams defines parameters for FindManyFavoriteFolders.
type FindManyFavoriteFoldersParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyFavoriteFoldersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyFavoriteFoldersParamsDepth defines parameters for FindManyFavoriteFolders.
type FindManyFavoriteFoldersParamsDepth int

// UpdateManyFavoriteFoldersParams defines parameters for UpdateManyFavoriteFolders.
type UpdateManyFavoriteFoldersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyFavoriteFoldersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyFavoriteFoldersParamsDepth defines parameters for UpdateManyFavoriteFolders.
type UpdateManyFavoriteFoldersParamsDepth int

// CreateOneFavoriteFolderParams defines parameters for CreateOneFavoriteFolder.
type CreateOneFavoriteFolderParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneFavoriteFolderParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneFavoriteFolderParamsDepth defines parameters for CreateOneFavoriteFolder.
type CreateOneFavoriteFolderParamsDepth int

// FindFavoriteFolderDuplicatesJSONBody defines parameters for FindFavoriteFolderDuplicates.
type FindFavoriteFolderDuplicatesJSONBody struct {
	Data *[]FavoriteFolder     `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindFavoriteFolderDuplicatesParams defines parameters for FindFavoriteFolderDuplicates.
type FindFavoriteFolderDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindFavoriteFolderDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindFavoriteFolderDuplicatesParamsDepth defines parameters for FindFavoriteFolderDuplicates.
type FindFavoriteFolderDuplicatesParamsDepth int

// MergeManyFavoriteFoldersJSONBody defines parameters for MergeManyFavoriteFolders.
type MergeManyFavoriteFoldersJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyFavoriteFoldersParams defines parameters for MergeManyFavoriteFolders.
type MergeManyFavoriteFoldersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyFavoriteFoldersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyFavoriteFoldersParamsDepth defines parameters for MergeManyFavoriteFolders.
type MergeManyFavoriteFoldersParamsDepth int

// DeleteOneFavoriteFolderParams defines parameters for DeleteOneFavoriteFolder.
type DeleteOneFavoriteFolderParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneFavoriteFolderParams defines parameters for FindOneFavoriteFolder.
type FindOneFavoriteFolderParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneFavoriteFolderParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneFavoriteFolderParamsDepth defines parameters for FindOneFavoriteFolder.
type FindOneFavoriteFolderParamsDepth int

// UpdateOneFavoriteFolderParams defines parameters for UpdateOneFavoriteFolder.
type UpdateOneFavoriteFolderParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneFavoriteFolderParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneFavoriteFolderParamsDepth defines parameters for UpdateOneFavoriteFolder.
type UpdateOneFavoriteFolderParamsDepth int

// DeleteManyFavoritesParams defines parameters for DeleteManyFavorites.
type DeleteManyFavoritesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyFavoritesParams defines parameters for FindManyFavorites.
type FindManyFavoritesParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyFavoritesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyFavoritesParamsDepth defines parameters for FindManyFavorites.
type FindManyFavoritesParamsDepth int

// UpdateManyFavoritesParams defines parameters for UpdateManyFavorites.
type UpdateManyFavoritesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyFavoritesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyFavoritesParamsDepth defines parameters for UpdateManyFavorites.
type UpdateManyFavoritesParamsDepth int

// CreateOneFavoriteParams defines parameters for CreateOneFavorite.
type CreateOneFavoriteParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneFavoriteParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneFavoriteParamsDepth defines parameters for CreateOneFavorite.
type CreateOneFavoriteParamsDepth int

// FindFavoriteDuplicatesJSONBody defines parameters for FindFavoriteDuplicates.
type FindFavoriteDuplicatesJSONBody struct {
	Data *[]Favorite           `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindFavoriteDuplicatesParams defines parameters for FindFavoriteDuplicates.
type FindFavoriteDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindFavoriteDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindFavoriteDuplicatesParamsDepth defines parameters for FindFavoriteDuplicates.
type FindFavoriteDuplicatesParamsDepth int

// MergeManyFavoritesJSONBody defines parameters for MergeManyFavorites.
type MergeManyFavoritesJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyFavoritesParams defines parameters for MergeManyFavorites.
type MergeManyFavoritesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyFavoritesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyFavoritesParamsDepth defines parameters for MergeManyFavorites.
type MergeManyFavoritesParamsDepth int

// DeleteOneFavoriteParams defines parameters for DeleteOneFavorite.
type DeleteOneFavoriteParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneFavoriteParams defines parameters for FindOneFavorite.
type FindOneFavoriteParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneFavoriteParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneFavoriteParamsDepth defines parameters for FindOneFavorite.
type FindOneFavoriteParamsDepth int

// UpdateOneFavoriteParams defines parameters for UpdateOneFavorite.
type UpdateOneFavoriteParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneFavoriteParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneFavoriteParamsDepth defines parameters for UpdateOneFavorite.
type UpdateOneFavoriteParamsDepth int

// DeleteManyMessageChannelMessageAssociationsParams defines parameters for DeleteManyMessageChannelMessageAssociations.
type DeleteManyMessageChannelMessageAssociationsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyMessageChannelMessageAssociationsParams defines parameters for FindManyMessageChannelMessageAssociations.
type FindManyMessageChannelMessageAssociationsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyMessageChannelMessageAssociationsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyMessageChannelMessageAssociationsParamsDepth defines parameters for FindManyMessageChannelMessageAssociations.
type FindManyMessageChannelMessageAssociationsParamsDepth int

// UpdateManyMessageChannelMessageAssociationsParams defines parameters for UpdateManyMessageChannelMessageAssociations.
type UpdateManyMessageChannelMessageAssociationsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyMessageChannelMessageAssociationsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyMessageChannelMessageAssociationsParamsDepth defines parameters for UpdateManyMessageChannelMessageAssociations.
type UpdateManyMessageChannelMessageAssociationsParamsDepth int

// CreateOneMessageChannelMessageAssociationParams defines parameters for CreateOneMessageChannelMessageAssociation.
type CreateOneMessageChannelMessageAssociationParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneMessageChannelMessageAssociationParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneMessageChannelMessageAssociationParamsDepth defines parameters for CreateOneMessageChannelMessageAssociation.
type CreateOneMessageChannelMessageAssociationParamsDepth int

// FindMessageChannelMessageAssociationDuplicatesJSONBody defines parameters for FindMessageChannelMessageAssociationDuplicates.
type FindMessageChannelMessageAssociationDuplicatesJSONBody struct {
	Data *[]MessageChannelMessageAssociation `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID               `json:"ids,omitempty"`
}

// FindMessageChannelMessageAssociationDuplicatesParams defines parameters for FindMessageChannelMessageAssociationDuplicates.
type FindMessageChannelMessageAssociationDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindMessageChannelMessageAssociationDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindMessageChannelMessageAssociationDuplicatesParamsDepth defines parameters for FindMessageChannelMessageAssociationDuplicates.
type FindMessageChannelMessageAssociationDuplicatesParamsDepth int

// MergeManyMessageChannelMessageAssociationsJSONBody defines parameters for MergeManyMessageChannelMessageAssociations.
type MergeManyMessageChannelMessageAssociationsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyMessageChannelMessageAssociationsParams defines parameters for MergeManyMessageChannelMessageAssociations.
type MergeManyMessageChannelMessageAssociationsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyMessageChannelMessageAssociationsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyMessageChannelMessageAssociationsParamsDepth defines parameters for MergeManyMessageChannelMessageAssociations.
type MergeManyMessageChannelMessageAssociationsParamsDepth int

// DeleteOneMessageChannelMessageAssociationParams defines parameters for DeleteOneMessageChannelMessageAssociation.
type DeleteOneMessageChannelMessageAssociationParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneMessageChannelMessageAssociationParams defines parameters for FindOneMessageChannelMessageAssociation.
type FindOneMessageChannelMessageAssociationParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneMessageChannelMessageAssociationParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneMessageChannelMessageAssociationParamsDepth defines parameters for FindOneMessageChannelMessageAssociation.
type FindOneMessageChannelMessageAssociationParamsDepth int

// UpdateOneMessageChannelMessageAssociationParams defines parameters for UpdateOneMessageChannelMessageAssociation.
type UpdateOneMessageChannelMessageAssociationParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneMessageChannelMessageAssociationParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneMessageChannelMessageAssociationParamsDepth defines parameters for UpdateOneMessageChannelMessageAssociation.
type UpdateOneMessageChannelMessageAssociationParamsDepth int

// DeleteManyMessageChannelsParams defines parameters for DeleteManyMessageChannels.
type DeleteManyMessageChannelsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyMessageChannelsParams defines parameters for FindManyMessageChannels.
type FindManyMessageChannelsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyMessageChannelsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyMessageChannelsParamsDepth defines parameters for FindManyMessageChannels.
type FindManyMessageChannelsParamsDepth int

// UpdateManyMessageChannelsParams defines parameters for UpdateManyMessageChannels.
type UpdateManyMessageChannelsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyMessageChannelsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyMessageChannelsParamsDepth defines parameters for UpdateManyMessageChannels.
type UpdateManyMessageChannelsParamsDepth int

// CreateOneMessageChannelParams defines parameters for CreateOneMessageChannel.
type CreateOneMessageChannelParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneMessageChannelParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneMessageChannelParamsDepth defines parameters for CreateOneMessageChannel.
type CreateOneMessageChannelParamsDepth int

// FindMessageChannelDuplicatesJSONBody defines parameters for FindMessageChannelDuplicates.
type FindMessageChannelDuplicatesJSONBody struct {
	Data *[]MessageChannel     `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindMessageChannelDuplicatesParams defines parameters for FindMessageChannelDuplicates.
type FindMessageChannelDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindMessageChannelDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindMessageChannelDuplicatesParamsDepth defines parameters for FindMessageChannelDuplicates.
type FindMessageChannelDuplicatesParamsDepth int

// MergeManyMessageChannelsJSONBody defines parameters for MergeManyMessageChannels.
type MergeManyMessageChannelsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyMessageChannelsParams defines parameters for MergeManyMessageChannels.
type MergeManyMessageChannelsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyMessageChannelsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyMessageChannelsParamsDepth defines parameters for MergeManyMessageChannels.
type MergeManyMessageChannelsParamsDepth int

// DeleteOneMessageChannelParams defines parameters for DeleteOneMessageChannel.
type DeleteOneMessageChannelParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneMessageChannelParams defines parameters for FindOneMessageChannel.
type FindOneMessageChannelParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneMessageChannelParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneMessageChannelParamsDepth defines parameters for FindOneMessageChannel.
type FindOneMessageChannelParamsDepth int

// UpdateOneMessageChannelParams defines parameters for UpdateOneMessageChannel.
type UpdateOneMessageChannelParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneMessageChannelParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneMessageChannelParamsDepth defines parameters for UpdateOneMessageChannel.
type UpdateOneMessageChannelParamsDepth int

// DeleteManyMessageFoldersParams defines parameters for DeleteManyMessageFolders.
type DeleteManyMessageFoldersParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyMessageFoldersParams defines parameters for FindManyMessageFolders.
type FindManyMessageFoldersParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyMessageFoldersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyMessageFoldersParamsDepth defines parameters for FindManyMessageFolders.
type FindManyMessageFoldersParamsDepth int

// UpdateManyMessageFoldersParams defines parameters for UpdateManyMessageFolders.
type UpdateManyMessageFoldersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyMessageFoldersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyMessageFoldersParamsDepth defines parameters for UpdateManyMessageFolders.
type UpdateManyMessageFoldersParamsDepth int

// CreateOneMessageFolderParams defines parameters for CreateOneMessageFolder.
type CreateOneMessageFolderParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneMessageFolderParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneMessageFolderParamsDepth defines parameters for CreateOneMessageFolder.
type CreateOneMessageFolderParamsDepth int

// FindMessageFolderDuplicatesJSONBody defines parameters for FindMessageFolderDuplicates.
type FindMessageFolderDuplicatesJSONBody struct {
	Data *[]MessageFolder      `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindMessageFolderDuplicatesParams defines parameters for FindMessageFolderDuplicates.
type FindMessageFolderDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindMessageFolderDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindMessageFolderDuplicatesParamsDepth defines parameters for FindMessageFolderDuplicates.
type FindMessageFolderDuplicatesParamsDepth int

// MergeManyMessageFoldersJSONBody defines parameters for MergeManyMessageFolders.
type MergeManyMessageFoldersJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyMessageFoldersParams defines parameters for MergeManyMessageFolders.
type MergeManyMessageFoldersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyMessageFoldersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyMessageFoldersParamsDepth defines parameters for MergeManyMessageFolders.
type MergeManyMessageFoldersParamsDepth int

// DeleteOneMessageFolderParams defines parameters for DeleteOneMessageFolder.
type DeleteOneMessageFolderParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneMessageFolderParams defines parameters for FindOneMessageFolder.
type FindOneMessageFolderParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneMessageFolderParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneMessageFolderParamsDepth defines parameters for FindOneMessageFolder.
type FindOneMessageFolderParamsDepth int

// UpdateOneMessageFolderParams defines parameters for UpdateOneMessageFolder.
type UpdateOneMessageFolderParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneMessageFolderParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneMessageFolderParamsDepth defines parameters for UpdateOneMessageFolder.
type UpdateOneMessageFolderParamsDepth int

// DeleteManyMessageParticipantsParams defines parameters for DeleteManyMessageParticipants.
type DeleteManyMessageParticipantsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyMessageParticipantsParams defines parameters for FindManyMessageParticipants.
type FindManyMessageParticipantsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyMessageParticipantsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyMessageParticipantsParamsDepth defines parameters for FindManyMessageParticipants.
type FindManyMessageParticipantsParamsDepth int

// UpdateManyMessageParticipantsParams defines parameters for UpdateManyMessageParticipants.
type UpdateManyMessageParticipantsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyMessageParticipantsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyMessageParticipantsParamsDepth defines parameters for UpdateManyMessageParticipants.
type UpdateManyMessageParticipantsParamsDepth int

// CreateOneMessageParticipantParams defines parameters for CreateOneMessageParticipant.
type CreateOneMessageParticipantParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneMessageParticipantParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneMessageParticipantParamsDepth defines parameters for CreateOneMessageParticipant.
type CreateOneMessageParticipantParamsDepth int

// FindMessageParticipantDuplicatesJSONBody defines parameters for FindMessageParticipantDuplicates.
type FindMessageParticipantDuplicatesJSONBody struct {
	Data *[]MessageParticipant `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindMessageParticipantDuplicatesParams defines parameters for FindMessageParticipantDuplicates.
type FindMessageParticipantDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindMessageParticipantDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindMessageParticipantDuplicatesParamsDepth defines parameters for FindMessageParticipantDuplicates.
type FindMessageParticipantDuplicatesParamsDepth int

// MergeManyMessageParticipantsJSONBody defines parameters for MergeManyMessageParticipants.
type MergeManyMessageParticipantsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyMessageParticipantsParams defines parameters for MergeManyMessageParticipants.
type MergeManyMessageParticipantsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyMessageParticipantsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyMessageParticipantsParamsDepth defines parameters for MergeManyMessageParticipants.
type MergeManyMessageParticipantsParamsDepth int

// DeleteOneMessageParticipantParams defines parameters for DeleteOneMessageParticipant.
type DeleteOneMessageParticipantParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneMessageParticipantParams defines parameters for FindOneMessageParticipant.
type FindOneMessageParticipantParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneMessageParticipantParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneMessageParticipantParamsDepth defines parameters for FindOneMessageParticipant.
type FindOneMessageParticipantParamsDepth int

// UpdateOneMessageParticipantParams defines parameters for UpdateOneMessageParticipant.
type UpdateOneMessageParticipantParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneMessageParticipantParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneMessageParticipantParamsDepth defines parameters for UpdateOneMessageParticipant.
type UpdateOneMessageParticipantParamsDepth int

// DeleteManyMessageThreadsParams defines parameters for DeleteManyMessageThreads.
type DeleteManyMessageThreadsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyMessageThreadsParams defines parameters for FindManyMessageThreads.
type FindManyMessageThreadsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyMessageThreadsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyMessageThreadsParamsDepth defines parameters for FindManyMessageThreads.
type FindManyMessageThreadsParamsDepth int

// UpdateManyMessageThreadsParams defines parameters for UpdateManyMessageThreads.
type UpdateManyMessageThreadsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyMessageThreadsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyMessageThreadsParamsDepth defines parameters for UpdateManyMessageThreads.
type UpdateManyMessageThreadsParamsDepth int

// CreateOneMessageThreadParams defines parameters for CreateOneMessageThread.
type CreateOneMessageThreadParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneMessageThreadParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneMessageThreadParamsDepth defines parameters for CreateOneMessageThread.
type CreateOneMessageThreadParamsDepth int

// FindMessageThreadDuplicatesJSONBody defines parameters for FindMessageThreadDuplicates.
type FindMessageThreadDuplicatesJSONBody struct {
	Data *[]MessageThread      `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindMessageThreadDuplicatesParams defines parameters for FindMessageThreadDuplicates.
type FindMessageThreadDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindMessageThreadDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindMessageThreadDuplicatesParamsDepth defines parameters for FindMessageThreadDuplicates.
type FindMessageThreadDuplicatesParamsDepth int

// MergeManyMessageThreadsJSONBody defines parameters for MergeManyMessageThreads.
type MergeManyMessageThreadsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyMessageThreadsParams defines parameters for MergeManyMessageThreads.
type MergeManyMessageThreadsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyMessageThreadsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyMessageThreadsParamsDepth defines parameters for MergeManyMessageThreads.
type MergeManyMessageThreadsParamsDepth int

// DeleteOneMessageThreadParams defines parameters for DeleteOneMessageThread.
type DeleteOneMessageThreadParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneMessageThreadParams defines parameters for FindOneMessageThread.
type FindOneMessageThreadParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneMessageThreadParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneMessageThreadParamsDepth defines parameters for FindOneMessageThread.
type FindOneMessageThreadParamsDepth int

// UpdateOneMessageThreadParams defines parameters for UpdateOneMessageThread.
type UpdateOneMessageThreadParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneMessageThreadParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneMessageThreadParamsDepth defines parameters for UpdateOneMessageThread.
type UpdateOneMessageThreadParamsDepth int

// DeleteManyMessagesParams defines parameters for DeleteManyMessages.
type DeleteManyMessagesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyMessagesParams defines parameters for FindManyMessages.
type FindManyMessagesParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyMessagesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyMessagesParamsDepth defines parameters for FindManyMessages.
type FindManyMessagesParamsDepth int

// UpdateManyMessagesParams defines parameters for UpdateManyMessages.
type UpdateManyMessagesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyMessagesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyMessagesParamsDepth defines parameters for UpdateManyMessages.
type UpdateManyMessagesParamsDepth int

// CreateOneMessageParams defines parameters for CreateOneMessage.
type CreateOneMessageParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneMessageParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneMessageParamsDepth defines parameters for CreateOneMessage.
type CreateOneMessageParamsDepth int

// FindMessageDuplicatesJSONBody defines parameters for FindMessageDuplicates.
type FindMessageDuplicatesJSONBody struct {
	Data *[]Message            `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindMessageDuplicatesParams defines parameters for FindMessageDuplicates.
type FindMessageDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindMessageDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindMessageDuplicatesParamsDepth defines parameters for FindMessageDuplicates.
type FindMessageDuplicatesParamsDepth int

// MergeManyMessagesJSONBody defines parameters for MergeManyMessages.
type MergeManyMessagesJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyMessagesParams defines parameters for MergeManyMessages.
type MergeManyMessagesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyMessagesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyMessagesParamsDepth defines parameters for MergeManyMessages.
type MergeManyMessagesParamsDepth int

// DeleteOneMessageParams defines parameters for DeleteOneMessage.
type DeleteOneMessageParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneMessageParams defines parameters for FindOneMessage.
type FindOneMessageParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneMessageParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneMessageParamsDepth defines parameters for FindOneMessage.
type FindOneMessageParamsDepth int

// UpdateOneMessageParams defines parameters for UpdateOneMessage.
type UpdateOneMessageParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneMessageParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneMessageParamsDepth defines parameters for UpdateOneMessage.
type UpdateOneMessageParamsDepth int

// DeleteManyNoteTargetsParams defines parameters for DeleteManyNoteTargets.
type DeleteManyNoteTargetsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyNoteTargetsParams defines parameters for FindManyNoteTargets.
type FindManyNoteTargetsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyNoteTargetsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyNoteTargetsParamsDepth defines parameters for FindManyNoteTargets.
type FindManyNoteTargetsParamsDepth int

// UpdateManyNoteTargetsParams defines parameters for UpdateManyNoteTargets.
type UpdateManyNoteTargetsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyNoteTargetsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyNoteTargetsParamsDepth defines parameters for UpdateManyNoteTargets.
type UpdateManyNoteTargetsParamsDepth int

// CreateOneNoteTargetParams defines parameters for CreateOneNoteTarget.
type CreateOneNoteTargetParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneNoteTargetParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneNoteTargetParamsDepth defines parameters for CreateOneNoteTarget.
type CreateOneNoteTargetParamsDepth int

// FindNoteTargetDuplicatesJSONBody defines parameters for FindNoteTargetDuplicates.
type FindNoteTargetDuplicatesJSONBody struct {
	Data *[]NoteTarget         `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindNoteTargetDuplicatesParams defines parameters for FindNoteTargetDuplicates.
type FindNoteTargetDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindNoteTargetDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindNoteTargetDuplicatesParamsDepth defines parameters for FindNoteTargetDuplicates.
type FindNoteTargetDuplicatesParamsDepth int

// MergeManyNoteTargetsJSONBody defines parameters for MergeManyNoteTargets.
type MergeManyNoteTargetsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyNoteTargetsParams defines parameters for MergeManyNoteTargets.
type MergeManyNoteTargetsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyNoteTargetsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyNoteTargetsParamsDepth defines parameters for MergeManyNoteTargets.
type MergeManyNoteTargetsParamsDepth int

// DeleteOneNoteTargetParams defines parameters for DeleteOneNoteTarget.
type DeleteOneNoteTargetParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneNoteTargetParams defines parameters for FindOneNoteTarget.
type FindOneNoteTargetParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneNoteTargetParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneNoteTargetParamsDepth defines parameters for FindOneNoteTarget.
type FindOneNoteTargetParamsDepth int

// UpdateOneNoteTargetParams defines parameters for UpdateOneNoteTarget.
type UpdateOneNoteTargetParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneNoteTargetParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneNoteTargetParamsDepth defines parameters for UpdateOneNoteTarget.
type UpdateOneNoteTargetParamsDepth int

// DeleteManyNotesParams defines parameters for DeleteManyNotes.
type DeleteManyNotesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyNotesParams defines parameters for FindManyNotes.
type FindManyNotesParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyNotesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyNotesParamsDepth defines parameters for FindManyNotes.
type FindManyNotesParamsDepth int

// UpdateManyNotesParams defines parameters for UpdateManyNotes.
type UpdateManyNotesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyNotesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyNotesParamsDepth defines parameters for UpdateManyNotes.
type UpdateManyNotesParamsDepth int

// CreateOneNoteParams defines parameters for CreateOneNote.
type CreateOneNoteParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneNoteParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneNoteParamsDepth defines parameters for CreateOneNote.
type CreateOneNoteParamsDepth int

// FindNoteDuplicatesJSONBody defines parameters for FindNoteDuplicates.
type FindNoteDuplicatesJSONBody struct {
	Data *[]Note               `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindNoteDuplicatesParams defines parameters for FindNoteDuplicates.
type FindNoteDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindNoteDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindNoteDuplicatesParamsDepth defines parameters for FindNoteDuplicates.
type FindNoteDuplicatesParamsDepth int

// MergeManyNotesJSONBody defines parameters for MergeManyNotes.
type MergeManyNotesJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyNotesParams defines parameters for MergeManyNotes.
type MergeManyNotesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyNotesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyNotesParamsDepth defines parameters for MergeManyNotes.
type MergeManyNotesParamsDepth int

// DeleteOneNoteParams defines parameters for DeleteOneNote.
type DeleteOneNoteParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneNoteParams defines parameters for FindOneNote.
type FindOneNoteParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneNoteParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneNoteParamsDepth defines parameters for FindOneNote.
type FindOneNoteParamsDepth int

// UpdateOneNoteParams defines parameters for UpdateOneNote.
type UpdateOneNoteParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneNoteParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneNoteParamsDepth defines parameters for UpdateOneNote.
type UpdateOneNoteParamsDepth int

// DeleteManyOpportunitiesParams defines parameters for DeleteManyOpportunities.
type DeleteManyOpportunitiesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyOpportunitiesParams defines parameters for FindManyOpportunities.
type FindManyOpportunitiesParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyOpportunitiesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyOpportunitiesParamsDepth defines parameters for FindManyOpportunities.
type FindManyOpportunitiesParamsDepth int

// UpdateManyOpportunitiesParams defines parameters for UpdateManyOpportunities.
type UpdateManyOpportunitiesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyOpportunitiesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyOpportunitiesParamsDepth defines parameters for UpdateManyOpportunities.
type UpdateManyOpportunitiesParamsDepth int

// CreateOneOpportunityParams defines parameters for CreateOneOpportunity.
type CreateOneOpportunityParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneOpportunityParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneOpportunityParamsDepth defines parameters for CreateOneOpportunity.
type CreateOneOpportunityParamsDepth int

// FindOpportunityDuplicatesJSONBody defines parameters for FindOpportunityDuplicates.
type FindOpportunityDuplicatesJSONBody struct {
	Data *[]Opportunity        `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindOpportunityDuplicatesParams defines parameters for FindOpportunityDuplicates.
type FindOpportunityDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOpportunityDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOpportunityDuplicatesParamsDepth defines parameters for FindOpportunityDuplicates.
type FindOpportunityDuplicatesParamsDepth int

// MergeManyOpportunitiesJSONBody defines parameters for MergeManyOpportunities.
type MergeManyOpportunitiesJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyOpportunitiesParams defines parameters for MergeManyOpportunities.
type MergeManyOpportunitiesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyOpportunitiesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyOpportunitiesParamsDepth defines parameters for MergeManyOpportunities.
type MergeManyOpportunitiesParamsDepth int

// DeleteOneOpportunityParams defines parameters for DeleteOneOpportunity.
type DeleteOneOpportunityParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneOpportunityParams defines parameters for FindOneOpportunity.
type FindOneOpportunityParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneOpportunityParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneOpportunityParamsDepth defines parameters for FindOneOpportunity.
type FindOneOpportunityParamsDepth int

// UpdateOneOpportunityParams defines parameters for UpdateOneOpportunity.
type UpdateOneOpportunityParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneOpportunityParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneOpportunityParamsDepth defines parameters for UpdateOneOpportunity.
type UpdateOneOpportunityParamsDepth int

// DeleteManyPeopleParams defines parameters for DeleteManyPeople.
type DeleteManyPeopleParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyPeopleParams defines parameters for FindManyPeople.
type FindManyPeopleParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyPeopleParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyPeopleParamsDepth defines parameters for FindManyPeople.
type FindManyPeopleParamsDepth int

// UpdateManyPeopleParams defines parameters for UpdateManyPeople.
type UpdateManyPeopleParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyPeopleParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyPeopleParamsDepth defines parameters for UpdateManyPeople.
type UpdateManyPeopleParamsDepth int

// CreateOnePersonParams defines parameters for CreateOnePerson.
type CreateOnePersonParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOnePersonParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOnePersonParamsDepth defines parameters for CreateOnePerson.
type CreateOnePersonParamsDepth int

// FindPersonDuplicatesJSONBody defines parameters for FindPersonDuplicates.
type FindPersonDuplicatesJSONBody struct {
	Data *[]Person             `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindPersonDuplicatesParams defines parameters for FindPersonDuplicates.
type FindPersonDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindPersonDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindPersonDuplicatesParamsDepth defines parameters for FindPersonDuplicates.
type FindPersonDuplicatesParamsDepth int

// MergeManyPeopleJSONBody defines parameters for MergeManyPeople.
type MergeManyPeopleJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyPeopleParams defines parameters for MergeManyPeople.
type MergeManyPeopleParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyPeopleParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyPeopleParamsDepth defines parameters for MergeManyPeople.
type MergeManyPeopleParamsDepth int

// DeleteOnePersonParams defines parameters for DeleteOnePerson.
type DeleteOnePersonParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOnePersonParams defines parameters for FindOnePerson.
type FindOnePersonParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOnePersonParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOnePersonParamsDepth defines parameters for FindOnePerson.
type FindOnePersonParamsDepth int

// UpdateOnePersonParams defines parameters for UpdateOnePerson.
type UpdateOnePersonParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOnePersonParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOnePersonParamsDepth defines parameters for UpdateOnePerson.
type UpdateOnePersonParamsDepth int

// RestoreManyAttachmentsParams defines parameters for RestoreManyAttachments.
type RestoreManyAttachmentsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyAttachmentsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyAttachmentsParamsDepth defines parameters for RestoreManyAttachments.
type RestoreManyAttachmentsParamsDepth int

// RestoreOneAttachmentParams defines parameters for RestoreOneAttachment.
type RestoreOneAttachmentParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneAttachmentParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneAttachmentParamsDepth defines parameters for RestoreOneAttachment.
type RestoreOneAttachmentParamsDepth int

// RestoreManyBlocklistsParams defines parameters for RestoreManyBlocklists.
type RestoreManyBlocklistsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyBlocklistsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyBlocklistsParamsDepth defines parameters for RestoreManyBlocklists.
type RestoreManyBlocklistsParamsDepth int

// RestoreOneBlocklistParams defines parameters for RestoreOneBlocklist.
type RestoreOneBlocklistParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneBlocklistParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneBlocklistParamsDepth defines parameters for RestoreOneBlocklist.
type RestoreOneBlocklistParamsDepth int

// RestoreManyCalendarChannelEventAssociationsParams defines parameters for RestoreManyCalendarChannelEventAssociations.
type RestoreManyCalendarChannelEventAssociationsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyCalendarChannelEventAssociationsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyCalendarChannelEventAssociationsParamsDepth defines parameters for RestoreManyCalendarChannelEventAssociations.
type RestoreManyCalendarChannelEventAssociationsParamsDepth int

// RestoreOneCalendarChannelEventAssociationParams defines parameters for RestoreOneCalendarChannelEventAssociation.
type RestoreOneCalendarChannelEventAssociationParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneCalendarChannelEventAssociationParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneCalendarChannelEventAssociationParamsDepth defines parameters for RestoreOneCalendarChannelEventAssociation.
type RestoreOneCalendarChannelEventAssociationParamsDepth int

// RestoreManyCalendarChannelsParams defines parameters for RestoreManyCalendarChannels.
type RestoreManyCalendarChannelsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyCalendarChannelsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyCalendarChannelsParamsDepth defines parameters for RestoreManyCalendarChannels.
type RestoreManyCalendarChannelsParamsDepth int

// RestoreOneCalendarChannelParams defines parameters for RestoreOneCalendarChannel.
type RestoreOneCalendarChannelParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneCalendarChannelParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneCalendarChannelParamsDepth defines parameters for RestoreOneCalendarChannel.
type RestoreOneCalendarChannelParamsDepth int

// RestoreManyCalendarEventParticipantsParams defines parameters for RestoreManyCalendarEventParticipants.
type RestoreManyCalendarEventParticipantsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyCalendarEventParticipantsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyCalendarEventParticipantsParamsDepth defines parameters for RestoreManyCalendarEventParticipants.
type RestoreManyCalendarEventParticipantsParamsDepth int

// RestoreOneCalendarEventParticipantParams defines parameters for RestoreOneCalendarEventParticipant.
type RestoreOneCalendarEventParticipantParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneCalendarEventParticipantParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneCalendarEventParticipantParamsDepth defines parameters for RestoreOneCalendarEventParticipant.
type RestoreOneCalendarEventParticipantParamsDepth int

// RestoreManyCalendarEventsParams defines parameters for RestoreManyCalendarEvents.
type RestoreManyCalendarEventsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyCalendarEventsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyCalendarEventsParamsDepth defines parameters for RestoreManyCalendarEvents.
type RestoreManyCalendarEventsParamsDepth int

// RestoreOneCalendarEventParams defines parameters for RestoreOneCalendarEvent.
type RestoreOneCalendarEventParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneCalendarEventParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneCalendarEventParamsDepth defines parameters for RestoreOneCalendarEvent.
type RestoreOneCalendarEventParamsDepth int

// RestoreManyCompaniesParams defines parameters for RestoreManyCompanies.
type RestoreManyCompaniesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyCompaniesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyCompaniesParamsDepth defines parameters for RestoreManyCompanies.
type RestoreManyCompaniesParamsDepth int

// RestoreOneCompanyParams defines parameters for RestoreOneCompany.
type RestoreOneCompanyParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneCompanyParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneCompanyParamsDepth defines parameters for RestoreOneCompany.
type RestoreOneCompanyParamsDepth int

// RestoreManyConnectedAccountsParams defines parameters for RestoreManyConnectedAccounts.
type RestoreManyConnectedAccountsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyConnectedAccountsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyConnectedAccountsParamsDepth defines parameters for RestoreManyConnectedAccounts.
type RestoreManyConnectedAccountsParamsDepth int

// RestoreOneConnectedAccountParams defines parameters for RestoreOneConnectedAccount.
type RestoreOneConnectedAccountParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneConnectedAccountParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneConnectedAccountParamsDepth defines parameters for RestoreOneConnectedAccount.
type RestoreOneConnectedAccountParamsDepth int

// RestoreManyDashboardsParams defines parameters for RestoreManyDashboards.
type RestoreManyDashboardsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyDashboardsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyDashboardsParamsDepth defines parameters for RestoreManyDashboards.
type RestoreManyDashboardsParamsDepth int

// RestoreOneDashboardParams defines parameters for RestoreOneDashboard.
type RestoreOneDashboardParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneDashboardParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneDashboardParamsDepth defines parameters for RestoreOneDashboard.
type RestoreOneDashboardParamsDepth int

// RestoreManyFavoriteFoldersParams defines parameters for RestoreManyFavoriteFolders.
type RestoreManyFavoriteFoldersParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyFavoriteFoldersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyFavoriteFoldersParamsDepth defines parameters for RestoreManyFavoriteFolders.
type RestoreManyFavoriteFoldersParamsDepth int

// RestoreOneFavoriteFolderParams defines parameters for RestoreOneFavoriteFolder.
type RestoreOneFavoriteFolderParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneFavoriteFolderParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneFavoriteFolderParamsDepth defines parameters for RestoreOneFavoriteFolder.
type RestoreOneFavoriteFolderParamsDepth int

// RestoreManyFavoritesParams defines parameters for RestoreManyFavorites.
type RestoreManyFavoritesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyFavoritesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyFavoritesParamsDepth defines parameters for RestoreManyFavorites.
type RestoreManyFavoritesParamsDepth int

// RestoreOneFavoriteParams defines parameters for RestoreOneFavorite.
type RestoreOneFavoriteParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneFavoriteParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneFavoriteParamsDepth defines parameters for RestoreOneFavorite.
type RestoreOneFavoriteParamsDepth int

// RestoreManyMessageChannelMessageAssociationsParams defines parameters for RestoreManyMessageChannelMessageAssociations.
type RestoreManyMessageChannelMessageAssociationsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyMessageChannelMessageAssociationsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyMessageChannelMessageAssociationsParamsDepth defines parameters for RestoreManyMessageChannelMessageAssociations.
type RestoreManyMessageChannelMessageAssociationsParamsDepth int

// RestoreOneMessageChannelMessageAssociationParams defines parameters for RestoreOneMessageChannelMessageAssociation.
type RestoreOneMessageChannelMessageAssociationParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneMessageChannelMessageAssociationParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneMessageChannelMessageAssociationParamsDepth defines parameters for RestoreOneMessageChannelMessageAssociation.
type RestoreOneMessageChannelMessageAssociationParamsDepth int

// RestoreManyMessageChannelsParams defines parameters for RestoreManyMessageChannels.
type RestoreManyMessageChannelsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyMessageChannelsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyMessageChannelsParamsDepth defines parameters for RestoreManyMessageChannels.
type RestoreManyMessageChannelsParamsDepth int

// RestoreOneMessageChannelParams defines parameters for RestoreOneMessageChannel.
type RestoreOneMessageChannelParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneMessageChannelParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneMessageChannelParamsDepth defines parameters for RestoreOneMessageChannel.
type RestoreOneMessageChannelParamsDepth int

// RestoreManyMessageFoldersParams defines parameters for RestoreManyMessageFolders.
type RestoreManyMessageFoldersParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyMessageFoldersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyMessageFoldersParamsDepth defines parameters for RestoreManyMessageFolders.
type RestoreManyMessageFoldersParamsDepth int

// RestoreOneMessageFolderParams defines parameters for RestoreOneMessageFolder.
type RestoreOneMessageFolderParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneMessageFolderParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneMessageFolderParamsDepth defines parameters for RestoreOneMessageFolder.
type RestoreOneMessageFolderParamsDepth int

// RestoreManyMessageParticipantsParams defines parameters for RestoreManyMessageParticipants.
type RestoreManyMessageParticipantsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyMessageParticipantsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyMessageParticipantsParamsDepth defines parameters for RestoreManyMessageParticipants.
type RestoreManyMessageParticipantsParamsDepth int

// RestoreOneMessageParticipantParams defines parameters for RestoreOneMessageParticipant.
type RestoreOneMessageParticipantParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneMessageParticipantParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneMessageParticipantParamsDepth defines parameters for RestoreOneMessageParticipant.
type RestoreOneMessageParticipantParamsDepth int

// RestoreManyMessageThreadsParams defines parameters for RestoreManyMessageThreads.
type RestoreManyMessageThreadsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyMessageThreadsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyMessageThreadsParamsDepth defines parameters for RestoreManyMessageThreads.
type RestoreManyMessageThreadsParamsDepth int

// RestoreOneMessageThreadParams defines parameters for RestoreOneMessageThread.
type RestoreOneMessageThreadParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneMessageThreadParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneMessageThreadParamsDepth defines parameters for RestoreOneMessageThread.
type RestoreOneMessageThreadParamsDepth int

// RestoreManyMessagesParams defines parameters for RestoreManyMessages.
type RestoreManyMessagesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyMessagesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyMessagesParamsDepth defines parameters for RestoreManyMessages.
type RestoreManyMessagesParamsDepth int

// RestoreOneMessageParams defines parameters for RestoreOneMessage.
type RestoreOneMessageParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneMessageParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneMessageParamsDepth defines parameters for RestoreOneMessage.
type RestoreOneMessageParamsDepth int

// RestoreManyNoteTargetsParams defines parameters for RestoreManyNoteTargets.
type RestoreManyNoteTargetsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyNoteTargetsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyNoteTargetsParamsDepth defines parameters for RestoreManyNoteTargets.
type RestoreManyNoteTargetsParamsDepth int

// RestoreOneNoteTargetParams defines parameters for RestoreOneNoteTarget.
type RestoreOneNoteTargetParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneNoteTargetParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneNoteTargetParamsDepth defines parameters for RestoreOneNoteTarget.
type RestoreOneNoteTargetParamsDepth int

// RestoreManyNotesParams defines parameters for RestoreManyNotes.
type RestoreManyNotesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyNotesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyNotesParamsDepth defines parameters for RestoreManyNotes.
type RestoreManyNotesParamsDepth int

// RestoreOneNoteParams defines parameters for RestoreOneNote.
type RestoreOneNoteParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneNoteParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneNoteParamsDepth defines parameters for RestoreOneNote.
type RestoreOneNoteParamsDepth int

// RestoreManyOpportunitiesParams defines parameters for RestoreManyOpportunities.
type RestoreManyOpportunitiesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyOpportunitiesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyOpportunitiesParamsDepth defines parameters for RestoreManyOpportunities.
type RestoreManyOpportunitiesParamsDepth int

// RestoreOneOpportunityParams defines parameters for RestoreOneOpportunity.
type RestoreOneOpportunityParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneOpportunityParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneOpportunityParamsDepth defines parameters for RestoreOneOpportunity.
type RestoreOneOpportunityParamsDepth int

// RestoreManyPeopleParams defines parameters for RestoreManyPeople.
type RestoreManyPeopleParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyPeopleParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyPeopleParamsDepth defines parameters for RestoreManyPeople.
type RestoreManyPeopleParamsDepth int

// RestoreOnePersonParams defines parameters for RestoreOnePerson.
type RestoreOnePersonParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOnePersonParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOnePersonParamsDepth defines parameters for RestoreOnePerson.
type RestoreOnePersonParamsDepth int

// RestoreManyTaskTargetsParams defines parameters for RestoreManyTaskTargets.
type RestoreManyTaskTargetsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyTaskTargetsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyTaskTargetsParamsDepth defines parameters for RestoreManyTaskTargets.
type RestoreManyTaskTargetsParamsDepth int

// RestoreOneTaskTargetParams defines parameters for RestoreOneTaskTarget.
type RestoreOneTaskTargetParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneTaskTargetParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneTaskTargetParamsDepth defines parameters for RestoreOneTaskTarget.
type RestoreOneTaskTargetParamsDepth int

// RestoreManyTasksParams defines parameters for RestoreManyTasks.
type RestoreManyTasksParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyTasksParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyTasksParamsDepth defines parameters for RestoreManyTasks.
type RestoreManyTasksParamsDepth int

// RestoreOneTaskParams defines parameters for RestoreOneTask.
type RestoreOneTaskParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneTaskParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneTaskParamsDepth defines parameters for RestoreOneTask.
type RestoreOneTaskParamsDepth int

// RestoreManyTimelineActivitiesParams defines parameters for RestoreManyTimelineActivities.
type RestoreManyTimelineActivitiesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyTimelineActivitiesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyTimelineActivitiesParamsDepth defines parameters for RestoreManyTimelineActivities.
type RestoreManyTimelineActivitiesParamsDepth int

// RestoreOneTimelineActivityParams defines parameters for RestoreOneTimelineActivity.
type RestoreOneTimelineActivityParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneTimelineActivityParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneTimelineActivityParamsDepth defines parameters for RestoreOneTimelineActivity.
type RestoreOneTimelineActivityParamsDepth int

// RestoreManyWorkflowAutomatedTriggersParams defines parameters for RestoreManyWorkflowAutomatedTriggers.
type RestoreManyWorkflowAutomatedTriggersParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyWorkflowAutomatedTriggersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyWorkflowAutomatedTriggersParamsDepth defines parameters for RestoreManyWorkflowAutomatedTriggers.
type RestoreManyWorkflowAutomatedTriggersParamsDepth int

// RestoreOneWorkflowAutomatedTriggerParams defines parameters for RestoreOneWorkflowAutomatedTrigger.
type RestoreOneWorkflowAutomatedTriggerParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneWorkflowAutomatedTriggerParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneWorkflowAutomatedTriggerParamsDepth defines parameters for RestoreOneWorkflowAutomatedTrigger.
type RestoreOneWorkflowAutomatedTriggerParamsDepth int

// RestoreManyWorkflowRunsParams defines parameters for RestoreManyWorkflowRuns.
type RestoreManyWorkflowRunsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyWorkflowRunsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyWorkflowRunsParamsDepth defines parameters for RestoreManyWorkflowRuns.
type RestoreManyWorkflowRunsParamsDepth int

// RestoreOneWorkflowRunParams defines parameters for RestoreOneWorkflowRun.
type RestoreOneWorkflowRunParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneWorkflowRunParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneWorkflowRunParamsDepth defines parameters for RestoreOneWorkflowRun.
type RestoreOneWorkflowRunParamsDepth int

// RestoreManyWorkflowVersionsParams defines parameters for RestoreManyWorkflowVersions.
type RestoreManyWorkflowVersionsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyWorkflowVersionsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyWorkflowVersionsParamsDepth defines parameters for RestoreManyWorkflowVersions.
type RestoreManyWorkflowVersionsParamsDepth int

// RestoreOneWorkflowVersionParams defines parameters for RestoreOneWorkflowVersion.
type RestoreOneWorkflowVersionParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneWorkflowVersionParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneWorkflowVersionParamsDepth defines parameters for RestoreOneWorkflowVersion.
type RestoreOneWorkflowVersionParamsDepth int

// RestoreManyWorkflowsParams defines parameters for RestoreManyWorkflows.
type RestoreManyWorkflowsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyWorkflowsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyWorkflowsParamsDepth defines parameters for RestoreManyWorkflows.
type RestoreManyWorkflowsParamsDepth int

// RestoreOneWorkflowParams defines parameters for RestoreOneWorkflow.
type RestoreOneWorkflowParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneWorkflowParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneWorkflowParamsDepth defines parameters for RestoreOneWorkflow.
type RestoreOneWorkflowParamsDepth int

// RestoreManyWorkspaceMembersParams defines parameters for RestoreManyWorkspaceMembers.
type RestoreManyWorkspaceMembersParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreManyWorkspaceMembersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreManyWorkspaceMembersParamsDepth defines parameters for RestoreManyWorkspaceMembers.
type RestoreManyWorkspaceMembersParamsDepth int

// RestoreOneWorkspaceMemberParams defines parameters for RestoreOneWorkspaceMember.
type RestoreOneWorkspaceMemberParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *RestoreOneWorkspaceMemberParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// RestoreOneWorkspaceMemberParamsDepth defines parameters for RestoreOneWorkspaceMember.
type RestoreOneWorkspaceMemberParamsDepth int

// DeleteManyTaskTargetsParams defines parameters for DeleteManyTaskTargets.
type DeleteManyTaskTargetsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyTaskTargetsParams defines parameters for FindManyTaskTargets.
type FindManyTaskTargetsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyTaskTargetsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyTaskTargetsParamsDepth defines parameters for FindManyTaskTargets.
type FindManyTaskTargetsParamsDepth int

// UpdateManyTaskTargetsParams defines parameters for UpdateManyTaskTargets.
type UpdateManyTaskTargetsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyTaskTargetsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyTaskTargetsParamsDepth defines parameters for UpdateManyTaskTargets.
type UpdateManyTaskTargetsParamsDepth int

// CreateOneTaskTargetParams defines parameters for CreateOneTaskTarget.
type CreateOneTaskTargetParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneTaskTargetParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneTaskTargetParamsDepth defines parameters for CreateOneTaskTarget.
type CreateOneTaskTargetParamsDepth int

// FindTaskTargetDuplicatesJSONBody defines parameters for FindTaskTargetDuplicates.
type FindTaskTargetDuplicatesJSONBody struct {
	Data *[]TaskTarget         `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindTaskTargetDuplicatesParams defines parameters for FindTaskTargetDuplicates.
type FindTaskTargetDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindTaskTargetDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindTaskTargetDuplicatesParamsDepth defines parameters for FindTaskTargetDuplicates.
type FindTaskTargetDuplicatesParamsDepth int

// MergeManyTaskTargetsJSONBody defines parameters for MergeManyTaskTargets.
type MergeManyTaskTargetsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyTaskTargetsParams defines parameters for MergeManyTaskTargets.
type MergeManyTaskTargetsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyTaskTargetsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyTaskTargetsParamsDepth defines parameters for MergeManyTaskTargets.
type MergeManyTaskTargetsParamsDepth int

// DeleteOneTaskTargetParams defines parameters for DeleteOneTaskTarget.
type DeleteOneTaskTargetParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneTaskTargetParams defines parameters for FindOneTaskTarget.
type FindOneTaskTargetParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneTaskTargetParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneTaskTargetParamsDepth defines parameters for FindOneTaskTarget.
type FindOneTaskTargetParamsDepth int

// UpdateOneTaskTargetParams defines parameters for UpdateOneTaskTarget.
type UpdateOneTaskTargetParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneTaskTargetParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneTaskTargetParamsDepth defines parameters for UpdateOneTaskTarget.
type UpdateOneTaskTargetParamsDepth int

// DeleteManyTasksParams defines parameters for DeleteManyTasks.
type DeleteManyTasksParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyTasksParams defines parameters for FindManyTasks.
type FindManyTasksParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyTasksParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyTasksParamsDepth defines parameters for FindManyTasks.
type FindManyTasksParamsDepth int

// UpdateManyTasksParams defines parameters for UpdateManyTasks.
type UpdateManyTasksParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyTasksParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyTasksParamsDepth defines parameters for UpdateManyTasks.
type UpdateManyTasksParamsDepth int

// CreateOneTaskParams defines parameters for CreateOneTask.
type CreateOneTaskParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneTaskParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneTaskParamsDepth defines parameters for CreateOneTask.
type CreateOneTaskParamsDepth int

// FindTaskDuplicatesJSONBody defines parameters for FindTaskDuplicates.
type FindTaskDuplicatesJSONBody struct {
	Data *[]Task               `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindTaskDuplicatesParams defines parameters for FindTaskDuplicates.
type FindTaskDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindTaskDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindTaskDuplicatesParamsDepth defines parameters for FindTaskDuplicates.
type FindTaskDuplicatesParamsDepth int

// MergeManyTasksJSONBody defines parameters for MergeManyTasks.
type MergeManyTasksJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyTasksParams defines parameters for MergeManyTasks.
type MergeManyTasksParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyTasksParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyTasksParamsDepth defines parameters for MergeManyTasks.
type MergeManyTasksParamsDepth int

// DeleteOneTaskParams defines parameters for DeleteOneTask.
type DeleteOneTaskParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneTaskParams defines parameters for FindOneTask.
type FindOneTaskParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneTaskParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneTaskParamsDepth defines parameters for FindOneTask.
type FindOneTaskParamsDepth int

// UpdateOneTaskParams defines parameters for UpdateOneTask.
type UpdateOneTaskParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneTaskParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneTaskParamsDepth defines parameters for UpdateOneTask.
type UpdateOneTaskParamsDepth int

// DeleteManyTimelineActivitiesParams defines parameters for DeleteManyTimelineActivities.
type DeleteManyTimelineActivitiesParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyTimelineActivitiesParams defines parameters for FindManyTimelineActivities.
type FindManyTimelineActivitiesParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyTimelineActivitiesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyTimelineActivitiesParamsDepth defines parameters for FindManyTimelineActivities.
type FindManyTimelineActivitiesParamsDepth int

// UpdateManyTimelineActivitiesParams defines parameters for UpdateManyTimelineActivities.
type UpdateManyTimelineActivitiesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyTimelineActivitiesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyTimelineActivitiesParamsDepth defines parameters for UpdateManyTimelineActivities.
type UpdateManyTimelineActivitiesParamsDepth int

// CreateOneTimelineActivityParams defines parameters for CreateOneTimelineActivity.
type CreateOneTimelineActivityParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneTimelineActivityParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneTimelineActivityParamsDepth defines parameters for CreateOneTimelineActivity.
type CreateOneTimelineActivityParamsDepth int

// FindTimelineActivityDuplicatesJSONBody defines parameters for FindTimelineActivityDuplicates.
type FindTimelineActivityDuplicatesJSONBody struct {
	Data *[]TimelineActivity   `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindTimelineActivityDuplicatesParams defines parameters for FindTimelineActivityDuplicates.
type FindTimelineActivityDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindTimelineActivityDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindTimelineActivityDuplicatesParamsDepth defines parameters for FindTimelineActivityDuplicates.
type FindTimelineActivityDuplicatesParamsDepth int

// MergeManyTimelineActivitiesJSONBody defines parameters for MergeManyTimelineActivities.
type MergeManyTimelineActivitiesJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyTimelineActivitiesParams defines parameters for MergeManyTimelineActivities.
type MergeManyTimelineActivitiesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyTimelineActivitiesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyTimelineActivitiesParamsDepth defines parameters for MergeManyTimelineActivities.
type MergeManyTimelineActivitiesParamsDepth int

// DeleteOneTimelineActivityParams defines parameters for DeleteOneTimelineActivity.
type DeleteOneTimelineActivityParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneTimelineActivityParams defines parameters for FindOneTimelineActivity.
type FindOneTimelineActivityParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneTimelineActivityParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneTimelineActivityParamsDepth defines parameters for FindOneTimelineActivity.
type FindOneTimelineActivityParamsDepth int

// UpdateOneTimelineActivityParams defines parameters for UpdateOneTimelineActivity.
type UpdateOneTimelineActivityParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneTimelineActivityParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneTimelineActivityParamsDepth defines parameters for UpdateOneTimelineActivity.
type UpdateOneTimelineActivityParamsDepth int

// DeleteManyWorkflowAutomatedTriggersParams defines parameters for DeleteManyWorkflowAutomatedTriggers.
type DeleteManyWorkflowAutomatedTriggersParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyWorkflowAutomatedTriggersParams defines parameters for FindManyWorkflowAutomatedTriggers.
type FindManyWorkflowAutomatedTriggersParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyWorkflowAutomatedTriggersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyWorkflowAutomatedTriggersParamsDepth defines parameters for FindManyWorkflowAutomatedTriggers.
type FindManyWorkflowAutomatedTriggersParamsDepth int

// UpdateManyWorkflowAutomatedTriggersParams defines parameters for UpdateManyWorkflowAutomatedTriggers.
type UpdateManyWorkflowAutomatedTriggersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyWorkflowAutomatedTriggersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyWorkflowAutomatedTriggersParamsDepth defines parameters for UpdateManyWorkflowAutomatedTriggers.
type UpdateManyWorkflowAutomatedTriggersParamsDepth int

// CreateOneWorkflowAutomatedTriggerParams defines parameters for CreateOneWorkflowAutomatedTrigger.
type CreateOneWorkflowAutomatedTriggerParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneWorkflowAutomatedTriggerParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneWorkflowAutomatedTriggerParamsDepth defines parameters for CreateOneWorkflowAutomatedTrigger.
type CreateOneWorkflowAutomatedTriggerParamsDepth int

// FindWorkflowAutomatedTriggerDuplicatesJSONBody defines parameters for FindWorkflowAutomatedTriggerDuplicates.
type FindWorkflowAutomatedTriggerDuplicatesJSONBody struct {
	Data *[]WorkflowAutomatedTrigger `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID       `json:"ids,omitempty"`
}

// FindWorkflowAutomatedTriggerDuplicatesParams defines parameters for FindWorkflowAutomatedTriggerDuplicates.
type FindWorkflowAutomatedTriggerDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindWorkflowAutomatedTriggerDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindWorkflowAutomatedTriggerDuplicatesParamsDepth defines parameters for FindWorkflowAutomatedTriggerDuplicates.
type FindWorkflowAutomatedTriggerDuplicatesParamsDepth int

// MergeManyWorkflowAutomatedTriggersJSONBody defines parameters for MergeManyWorkflowAutomatedTriggers.
type MergeManyWorkflowAutomatedTriggersJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyWorkflowAutomatedTriggersParams defines parameters for MergeManyWorkflowAutomatedTriggers.
type MergeManyWorkflowAutomatedTriggersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyWorkflowAutomatedTriggersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyWorkflowAutomatedTriggersParamsDepth defines parameters for MergeManyWorkflowAutomatedTriggers.
type MergeManyWorkflowAutomatedTriggersParamsDepth int

// DeleteOneWorkflowAutomatedTriggerParams defines parameters for DeleteOneWorkflowAutomatedTrigger.
type DeleteOneWorkflowAutomatedTriggerParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneWorkflowAutomatedTriggerParams defines parameters for FindOneWorkflowAutomatedTrigger.
type FindOneWorkflowAutomatedTriggerParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneWorkflowAutomatedTriggerParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneWorkflowAutomatedTriggerParamsDepth defines parameters for FindOneWorkflowAutomatedTrigger.
type FindOneWorkflowAutomatedTriggerParamsDepth int

// UpdateOneWorkflowAutomatedTriggerParams defines parameters for UpdateOneWorkflowAutomatedTrigger.
type UpdateOneWorkflowAutomatedTriggerParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneWorkflowAutomatedTriggerParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneWorkflowAutomatedTriggerParamsDepth defines parameters for UpdateOneWorkflowAutomatedTrigger.
type UpdateOneWorkflowAutomatedTriggerParamsDepth int

// DeleteManyWorkflowRunsParams defines parameters for DeleteManyWorkflowRuns.
type DeleteManyWorkflowRunsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyWorkflowRunsParams defines parameters for FindManyWorkflowRuns.
type FindManyWorkflowRunsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyWorkflowRunsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyWorkflowRunsParamsDepth defines parameters for FindManyWorkflowRuns.
type FindManyWorkflowRunsParamsDepth int

// UpdateManyWorkflowRunsParams defines parameters for UpdateManyWorkflowRuns.
type UpdateManyWorkflowRunsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyWorkflowRunsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyWorkflowRunsParamsDepth defines parameters for UpdateManyWorkflowRuns.
type UpdateManyWorkflowRunsParamsDepth int

// CreateOneWorkflowRunParams defines parameters for CreateOneWorkflowRun.
type CreateOneWorkflowRunParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneWorkflowRunParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneWorkflowRunParamsDepth defines parameters for CreateOneWorkflowRun.
type CreateOneWorkflowRunParamsDepth int

// FindWorkflowRunDuplicatesJSONBody defines parameters for FindWorkflowRunDuplicates.
type FindWorkflowRunDuplicatesJSONBody struct {
	Data *[]WorkflowRun        `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindWorkflowRunDuplicatesParams defines parameters for FindWorkflowRunDuplicates.
type FindWorkflowRunDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindWorkflowRunDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindWorkflowRunDuplicatesParamsDepth defines parameters for FindWorkflowRunDuplicates.
type FindWorkflowRunDuplicatesParamsDepth int

// MergeManyWorkflowRunsJSONBody defines parameters for MergeManyWorkflowRuns.
type MergeManyWorkflowRunsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyWorkflowRunsParams defines parameters for MergeManyWorkflowRuns.
type MergeManyWorkflowRunsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyWorkflowRunsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyWorkflowRunsParamsDepth defines parameters for MergeManyWorkflowRuns.
type MergeManyWorkflowRunsParamsDepth int

// DeleteOneWorkflowRunParams defines parameters for DeleteOneWorkflowRun.
type DeleteOneWorkflowRunParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneWorkflowRunParams defines parameters for FindOneWorkflowRun.
type FindOneWorkflowRunParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneWorkflowRunParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneWorkflowRunParamsDepth defines parameters for FindOneWorkflowRun.
type FindOneWorkflowRunParamsDepth int

// UpdateOneWorkflowRunParams defines parameters for UpdateOneWorkflowRun.
type UpdateOneWorkflowRunParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneWorkflowRunParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneWorkflowRunParamsDepth defines parameters for UpdateOneWorkflowRun.
type UpdateOneWorkflowRunParamsDepth int

// DeleteManyWorkflowVersionsParams defines parameters for DeleteManyWorkflowVersions.
type DeleteManyWorkflowVersionsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyWorkflowVersionsParams defines parameters for FindManyWorkflowVersions.
type FindManyWorkflowVersionsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyWorkflowVersionsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyWorkflowVersionsParamsDepth defines parameters for FindManyWorkflowVersions.
type FindManyWorkflowVersionsParamsDepth int

// UpdateManyWorkflowVersionsParams defines parameters for UpdateManyWorkflowVersions.
type UpdateManyWorkflowVersionsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyWorkflowVersionsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyWorkflowVersionsParamsDepth defines parameters for UpdateManyWorkflowVersions.
type UpdateManyWorkflowVersionsParamsDepth int

// CreateOneWorkflowVersionParams defines parameters for CreateOneWorkflowVersion.
type CreateOneWorkflowVersionParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneWorkflowVersionParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneWorkflowVersionParamsDepth defines parameters for CreateOneWorkflowVersion.
type CreateOneWorkflowVersionParamsDepth int

// FindWorkflowVersionDuplicatesJSONBody defines parameters for FindWorkflowVersionDuplicates.
type FindWorkflowVersionDuplicatesJSONBody struct {
	Data *[]WorkflowVersion    `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindWorkflowVersionDuplicatesParams defines parameters for FindWorkflowVersionDuplicates.
type FindWorkflowVersionDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindWorkflowVersionDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindWorkflowVersionDuplicatesParamsDepth defines parameters for FindWorkflowVersionDuplicates.
type FindWorkflowVersionDuplicatesParamsDepth int

// MergeManyWorkflowVersionsJSONBody defines parameters for MergeManyWorkflowVersions.
type MergeManyWorkflowVersionsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyWorkflowVersionsParams defines parameters for MergeManyWorkflowVersions.
type MergeManyWorkflowVersionsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyWorkflowVersionsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyWorkflowVersionsParamsDepth defines parameters for MergeManyWorkflowVersions.
type MergeManyWorkflowVersionsParamsDepth int

// DeleteOneWorkflowVersionParams defines parameters for DeleteOneWorkflowVersion.
type DeleteOneWorkflowVersionParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneWorkflowVersionParams defines parameters for FindOneWorkflowVersion.
type FindOneWorkflowVersionParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneWorkflowVersionParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneWorkflowVersionParamsDepth defines parameters for FindOneWorkflowVersion.
type FindOneWorkflowVersionParamsDepth int

// UpdateOneWorkflowVersionParams defines parameters for UpdateOneWorkflowVersion.
type UpdateOneWorkflowVersionParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneWorkflowVersionParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneWorkflowVersionParamsDepth defines parameters for UpdateOneWorkflowVersion.
type UpdateOneWorkflowVersionParamsDepth int

// DeleteManyWorkflowsParams defines parameters for DeleteManyWorkflows.
type DeleteManyWorkflowsParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyWorkflowsParams defines parameters for FindManyWorkflows.
type FindManyWorkflowsParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyWorkflowsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyWorkflowsParamsDepth defines parameters for FindManyWorkflows.
type FindManyWorkflowsParamsDepth int

// UpdateManyWorkflowsParams defines parameters for UpdateManyWorkflows.
type UpdateManyWorkflowsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyWorkflowsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyWorkflowsParamsDepth defines parameters for UpdateManyWorkflows.
type UpdateManyWorkflowsParamsDepth int

// CreateOneWorkflowParams defines parameters for CreateOneWorkflow.
type CreateOneWorkflowParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneWorkflowParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneWorkflowParamsDepth defines parameters for CreateOneWorkflow.
type CreateOneWorkflowParamsDepth int

// FindWorkflowDuplicatesJSONBody defines parameters for FindWorkflowDuplicates.
type FindWorkflowDuplicatesJSONBody struct {
	Data *[]Workflow           `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindWorkflowDuplicatesParams defines parameters for FindWorkflowDuplicates.
type FindWorkflowDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindWorkflowDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindWorkflowDuplicatesParamsDepth defines parameters for FindWorkflowDuplicates.
type FindWorkflowDuplicatesParamsDepth int

// MergeManyWorkflowsJSONBody defines parameters for MergeManyWorkflows.
type MergeManyWorkflowsJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyWorkflowsParams defines parameters for MergeManyWorkflows.
type MergeManyWorkflowsParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyWorkflowsParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyWorkflowsParamsDepth defines parameters for MergeManyWorkflows.
type MergeManyWorkflowsParamsDepth int

// DeleteOneWorkflowParams defines parameters for DeleteOneWorkflow.
type DeleteOneWorkflowParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneWorkflowParams defines parameters for FindOneWorkflow.
type FindOneWorkflowParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneWorkflowParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneWorkflowParamsDepth defines parameters for FindOneWorkflow.
type FindOneWorkflowParamsDepth int

// UpdateOneWorkflowParams defines parameters for UpdateOneWorkflow.
type UpdateOneWorkflowParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneWorkflowParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneWorkflowParamsDepth defines parameters for UpdateOneWorkflow.
type UpdateOneWorkflowParamsDepth int

// DeleteManyWorkspaceMembersParams defines parameters for DeleteManyWorkspaceMembers.
type DeleteManyWorkspaceMembersParams struct {
	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindManyWorkspaceMembersParams defines parameters for FindManyWorkspaceMembers.
type FindManyWorkspaceMembersParams struct {
	// OrderBy Format: **field_name_1,field_name_2[DIRECTION_2]
	//     Refer to the filter section at the top of the page for more details.
	OrderBy *OrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`

	// Limit Limits the number of objects returned.
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindManyWorkspaceMembersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// StartingAfter Returns objects starting after a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	StartingAfter *StartingAfter `form:"starting_after,omitempty" json:"starting_after,omitempty"`

	// EndingBefore Returns objects ending before a specific cursor. You can find cursors in **startCursor** and **endCursor** in **pageInfo** in response data
	EndingBefore *EndingBefore `form:"ending_before,omitempty" json:"ending_before,omitempty"`
}

// FindManyWorkspaceMembersParamsDepth defines parameters for FindManyWorkspaceMembers.
type FindManyWorkspaceMembersParamsDepth int

// UpdateManyWorkspaceMembersParams defines parameters for UpdateManyWorkspaceMembers.
type UpdateManyWorkspaceMembersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateManyWorkspaceMembersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Filter Format: field[COMPARATOR]:value,field2[COMPARATOR]:value2
	//     Refer to the filter section at the top of the page for more details.
	Filter *Filter `form:"filter,omitempty" json:"filter,omitempty"`
}

// UpdateManyWorkspaceMembersParamsDepth defines parameters for UpdateManyWorkspaceMembers.
type UpdateManyWorkspaceMembersParamsDepth int

// CreateOneWorkspaceMemberParams defines parameters for CreateOneWorkspaceMember.
type CreateOneWorkspaceMemberParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *CreateOneWorkspaceMemberParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`

	// Upsert If true, creates the object or updates it if it already exists.
	Upsert *Upsert `form:"upsert,omitempty" json:"upsert,omitempty"`
}

// CreateOneWorkspaceMemberParamsDepth defines parameters for CreateOneWorkspaceMember.
type CreateOneWorkspaceMemberParamsDepth int

// FindWorkspaceMemberDuplicatesJSONBody defines parameters for FindWorkspaceMemberDuplicates.
type FindWorkspaceMemberDuplicatesJSONBody struct {
	Data *[]WorkspaceMember    `json:"data,omitempty"`
	Ids  *[]openapi_types.UUID `json:"ids,omitempty"`
}

// FindWorkspaceMemberDuplicatesParams defines parameters for FindWorkspaceMemberDuplicates.
type FindWorkspaceMemberDuplicatesParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindWorkspaceMemberDuplicatesParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindWorkspaceMemberDuplicatesParamsDepth defines parameters for FindWorkspaceMemberDuplicates.
type FindWorkspaceMemberDuplicatesParamsDepth int

// MergeManyWorkspaceMembersJSONBody defines parameters for MergeManyWorkspaceMembers.
type MergeManyWorkspaceMembersJSONBody struct {
	// ConflictPriorityIndex The index of the record to use when conflicts occur
	ConflictPriorityIndex float32 `json:"conflictPriorityIndex"`

	// DryRun If true, the merge will not be performed and a preview of the merge will be returned.
	DryRun *bool `json:"dryRun,omitempty"`

	// Ids The IDs of the records to merge
	Ids []openapi_types.UUID `json:"ids"`
}

// MergeManyWorkspaceMembersParams defines parameters for MergeManyWorkspaceMembers.
type MergeManyWorkspaceMembersParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *MergeManyWorkspaceMembersParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// MergeManyWorkspaceMembersParamsDepth defines parameters for MergeManyWorkspaceMembers.
type MergeManyWorkspaceMembersParamsDepth int

// DeleteOneWorkspaceMemberParams defines parameters for DeleteOneWorkspaceMember.
type DeleteOneWorkspaceMemberParams struct {
	// SoftDelete If true, soft deletes the objects. If false, objects are permanently deleted.
	SoftDelete *SoftDelete `form:"soft_delete,omitempty" json:"soft_delete,omitempty"`
}

// FindOneWorkspaceMemberParams defines parameters for FindOneWorkspaceMember.
type FindOneWorkspaceMemberParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *FindOneWorkspaceMemberParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// FindOneWorkspaceMemberParamsDepth defines parameters for FindOneWorkspaceMember.
type FindOneWorkspaceMemberParamsDepth int

// UpdateOneWorkspaceMemberParams defines parameters for UpdateOneWorkspaceMember.
type UpdateOneWorkspaceMemberParams struct {
	// Depth Determines the level of nested related objects to include in the response.
	//     - 0: Primary object only
	//     - 1: Primary object + direct relations
	Depth *UpdateOneWorkspaceMemberParamsDepth `form:"depth,omitempty" json:"depth,omitempty"`
}

// UpdateOneWorkspaceMemberParamsDepth defines parameters for UpdateOneWorkspaceMember.
type UpdateOneWorkspaceMemberParamsDepth int

// UpdateManyAttachmentsJSONRequestBody defines body for UpdateManyAttachments for application/json ContentType.
type UpdateManyAttachmentsJSONRequestBody = AttachmentForUpdate

// CreateOneAttachmentJSONRequestBody defines body for CreateOneAttachment for application/json ContentType.
type CreateOneAttachmentJSONRequestBody = Attachment

// FindAttachmentDuplicatesJSONRequestBody defines body for FindAttachmentDuplicates for application/json ContentType.
type FindAttachmentDuplicatesJSONRequestBody FindAttachmentDuplicatesJSONBody

// MergeManyAttachmentsJSONRequestBody defines body for MergeManyAttachments for application/json ContentType.
type MergeManyAttachmentsJSONRequestBody MergeManyAttachmentsJSONBody

// UpdateOneAttachmentJSONRequestBody defines body for UpdateOneAttachment for application/json ContentType.
type UpdateOneAttachmentJSONRequestBody = AttachmentForUpdate

// CreateManyAttachmentsJSONRequestBody defines body for CreateManyAttachments for application/json ContentType.
type CreateManyAttachmentsJSONRequestBody = CreateManyAttachmentsJSONBody

// CreateManyBlocklistsJSONRequestBody defines body for CreateManyBlocklists for application/json ContentType.
type CreateManyBlocklistsJSONRequestBody = CreateManyBlocklistsJSONBody

// CreateManyCalendarChannelEventAssociationsJSONRequestBody defines body for CreateManyCalendarChannelEventAssociations for application/json ContentType.
type CreateManyCalendarChannelEventAssociationsJSONRequestBody = CreateManyCalendarChannelEventAssociationsJSONBody

// CreateManyCalendarChannelsJSONRequestBody defines body for CreateManyCalendarChannels for application/json ContentType.
type CreateManyCalendarChannelsJSONRequestBody = CreateManyCalendarChannelsJSONBody

// CreateManyCalendarEventParticipantsJSONRequestBody defines body for CreateManyCalendarEventParticipants for application/json ContentType.
type CreateManyCalendarEventParticipantsJSONRequestBody = CreateManyCalendarEventParticipantsJSONBody

// CreateManyCalendarEventsJSONRequestBody defines body for CreateManyCalendarEvents for application/json ContentType.
type CreateManyCalendarEventsJSONRequestBody = CreateManyCalendarEventsJSONBody

// CreateManyCompaniesJSONRequestBody defines body for CreateManyCompanies for application/json ContentType.
type CreateManyCompaniesJSONRequestBody = CreateManyCompaniesJSONBody

// CreateManyConnectedAccountsJSONRequestBody defines body for CreateManyConnectedAccounts for application/json ContentType.
type CreateManyConnectedAccountsJSONRequestBody = CreateManyConnectedAccountsJSONBody

// CreateManyDashboardsJSONRequestBody defines body for CreateManyDashboards for application/json ContentType.
type CreateManyDashboardsJSONRequestBody = CreateManyDashboardsJSONBody

// CreateManyFavoriteFoldersJSONRequestBody defines body for CreateManyFavoriteFolders for application/json ContentType.
type CreateManyFavoriteFoldersJSONRequestBody = CreateManyFavoriteFoldersJSONBody

// CreateManyFavoritesJSONRequestBody defines body for CreateManyFavorites for application/json ContentType.
type CreateManyFavoritesJSONRequestBody = CreateManyFavoritesJSONBody

// CreateManyMessageChannelMessageAssociationsJSONRequestBody defines body for CreateManyMessageChannelMessageAssociations for application/json ContentType.
type CreateManyMessageChannelMessageAssociationsJSONRequestBody = CreateManyMessageChannelMessageAssociationsJSONBody

// CreateManyMessageChannelsJSONRequestBody defines body for CreateManyMessageChannels for application/json ContentType.
type CreateManyMessageChannelsJSONRequestBody = CreateManyMessageChannelsJSONBody

// CreateManyMessageFoldersJSONRequestBody defines body for CreateManyMessageFolders for application/json ContentType.
type CreateManyMessageFoldersJSONRequestBody = CreateManyMessageFoldersJSONBody

// CreateManyMessageParticipantsJSONRequestBody defines body for CreateManyMessageParticipants for application/json ContentType.
type CreateManyMessageParticipantsJSONRequestBody = CreateManyMessageParticipantsJSONBody

// CreateManyMessageThreadsJSONRequestBody defines body for CreateManyMessageThreads for application/json ContentType.
type CreateManyMessageThreadsJSONRequestBody = CreateManyMessageThreadsJSONBody

// CreateManyMessagesJSONRequestBody defines body for CreateManyMessages for application/json ContentType.
type CreateManyMessagesJSONRequestBody = CreateManyMessagesJSONBody

// CreateManyNoteTargetsJSONRequestBody defines body for CreateManyNoteTargets for application/json ContentType.
type CreateManyNoteTargetsJSONRequestBody = CreateManyNoteTargetsJSONBody

// CreateManyNotesJSONRequestBody defines body for CreateManyNotes for application/json ContentType.
type CreateManyNotesJSONRequestBody = CreateManyNotesJSONBody

// CreateManyOpportunitiesJSONRequestBody defines body for CreateManyOpportunities for application/json ContentType.
type CreateManyOpportunitiesJSONRequestBody = CreateManyOpportunitiesJSONBody

// CreateManyPeopleJSONRequestBody defines body for CreateManyPeople for application/json ContentType.
type CreateManyPeopleJSONRequestBody = CreateManyPeopleJSONBody

// CreateManyTaskTargetsJSONRequestBody defines body for CreateManyTaskTargets for application/json ContentType.
type CreateManyTaskTargetsJSONRequestBody = CreateManyTaskTargetsJSONBody

// CreateManyTasksJSONRequestBody defines body for CreateManyTasks for application/json ContentType.
type CreateManyTasksJSONRequestBody = CreateManyTasksJSONBody

// CreateManyTimelineActivitiesJSONRequestBody defines body for CreateManyTimelineActivities for application/json ContentType.
type CreateManyTimelineActivitiesJSONRequestBody = CreateManyTimelineActivitiesJSONBody

// CreateManyWorkflowAutomatedTriggersJSONRequestBody defines body for CreateManyWorkflowAutomatedTriggers for application/json ContentType.
type CreateManyWorkflowAutomatedTriggersJSONRequestBody = CreateManyWorkflowAutomatedTriggersJSONBody

// CreateManyWorkflowRunsJSONRequestBody defines body for CreateManyWorkflowRuns for application/json ContentType.
type CreateManyWorkflowRunsJSONRequestBody = CreateManyWorkflowRunsJSONBody

// CreateManyWorkflowVersionsJSONRequestBody defines body for CreateManyWorkflowVersions for application/json ContentType.
type CreateManyWorkflowVersionsJSONRequestBody = CreateManyWorkflowVersionsJSONBody

// CreateManyWorkflowsJSONRequestBody defines body for CreateManyWorkflows for application/json ContentType.
type CreateManyWorkflowsJSONRequestBody = CreateManyWorkflowsJSONBody

// CreateManyWorkspaceMembersJSONRequestBody defines body for CreateManyWorkspaceMembers for application/json ContentType.
type CreateManyWorkspaceMembersJSONRequestBody = CreateManyWorkspaceMembersJSONBody

// UpdateManyBlocklistsJSONRequestBody defines body for UpdateManyBlocklists for application/json ContentType.
type UpdateManyBlocklistsJSONRequestBody = BlocklistForUpdate

// CreateOneBlocklistJSONRequestBody defines body for CreateOneBlocklist for application/json ContentType.
type CreateOneBlocklistJSONRequestBody = Blocklist

// FindBlocklistDuplicatesJSONRequestBody defines body for FindBlocklistDuplicates for application/json ContentType.
type FindBlocklistDuplicatesJSONRequestBody FindBlocklistDuplicatesJSONBody

// MergeManyBlocklistsJSONRequestBody defines body for MergeManyBlocklists for application/json ContentType.
type MergeManyBlocklistsJSONRequestBody MergeManyBlocklistsJSONBody

// UpdateOneBlocklistJSONRequestBody defines body for UpdateOneBlocklist for application/json ContentType.
type UpdateOneBlocklistJSONRequestBody = BlocklistForUpdate

// UpdateManyCalendarChannelEventAssociationsJSONRequestBody defines body for UpdateManyCalendarChannelEventAssociations for application/json ContentType.
type UpdateManyCalendarChannelEventAssociationsJSONRequestBody = CalendarChannelEventAssociationForUpdate

// CreateOneCalendarChannelEventAssociationJSONRequestBody defines body for CreateOneCalendarChannelEventAssociation for application/json ContentType.
type CreateOneCalendarChannelEventAssociationJSONRequestBody = CalendarChannelEventAssociation

// FindCalendarChannelEventAssociationDuplicatesJSONRequestBody defines body for FindCalendarChannelEventAssociationDuplicates for application/json ContentType.
type FindCalendarChannelEventAssociationDuplicatesJSONRequestBody FindCalendarChannelEventAssociationDuplicatesJSONBody

// MergeManyCalendarChannelEventAssociationsJSONRequestBody defines body for MergeManyCalendarChannelEventAssociations for application/json ContentType.
type MergeManyCalendarChannelEventAssociationsJSONRequestBody MergeManyCalendarChannelEventAssociationsJSONBody

// UpdateOneCalendarChannelEventAssociationJSONRequestBody defines body for UpdateOneCalendarChannelEventAssociation for application/json ContentType.
type UpdateOneCalendarChannelEventAssociationJSONRequestBody = CalendarChannelEventAssociationForUpdate

// UpdateManyCalendarChannelsJSONRequestBody defines body for UpdateManyCalendarChannels for application/json ContentType.
type UpdateManyCalendarChannelsJSONRequestBody = CalendarChannelForUpdate

// CreateOneCalendarChannelJSONRequestBody defines body for CreateOneCalendarChannel for application/json ContentType.
type CreateOneCalendarChannelJSONRequestBody = CalendarChannel

// FindCalendarChannelDuplicatesJSONRequestBody defines body for FindCalendarChannelDuplicates for application/json ContentType.
type FindCalendarChannelDuplicatesJSONRequestBody FindCalendarChannelDuplicatesJSONBody

// MergeManyCalendarChannelsJSONRequestBody defines body for MergeManyCalendarChannels for application/json ContentType.
type MergeManyCalendarChannelsJSONRequestBody MergeManyCalendarChannelsJSONBody

// UpdateOneCalendarChannelJSONRequestBody defines body for UpdateOneCalendarChannel for application/json ContentType.
type UpdateOneCalendarChannelJSONRequestBody = CalendarChannelForUpdate

// UpdateManyCalendarEventParticipantsJSONRequestBody defines body for UpdateManyCalendarEventParticipants for application/json ContentType.
type UpdateManyCalendarEventParticipantsJSONRequestBody = CalendarEventParticipantForUpdate

// CreateOneCalendarEventParticipantJSONRequestBody defines body for CreateOneCalendarEventParticipant for application/json ContentType.
type CreateOneCalendarEventParticipantJSONRequestBody = CalendarEventParticipant

// FindCalendarEventParticipantDuplicatesJSONRequestBody defines body for FindCalendarEventParticipantDuplicates for application/json ContentType.
type FindCalendarEventParticipantDuplicatesJSONRequestBody FindCalendarEventParticipantDuplicatesJSONBody

// MergeManyCalendarEventParticipantsJSONRequestBody defines body for MergeManyCalendarEventParticipants for application/json ContentType.
type MergeManyCalendarEventParticipantsJSONRequestBody MergeManyCalendarEventParticipantsJSONBody

// UpdateOneCalendarEventParticipantJSONRequestBody defines body for UpdateOneCalendarEventParticipant for application/json ContentType.
type UpdateOneCalendarEventParticipantJSONRequestBody = CalendarEventParticipantForUpdate

// UpdateManyCalendarEventsJSONRequestBody defines body for UpdateManyCalendarEvents for application/json ContentType.
type UpdateManyCalendarEventsJSONRequestBody = CalendarEventForUpdate

// CreateOneCalendarEventJSONRequestBody defines body for CreateOneCalendarEvent for application/json ContentType.
type CreateOneCalendarEventJSONRequestBody = CalendarEvent

// FindCalendarEventDuplicatesJSONRequestBody defines body for FindCalendarEventDuplicates for application/json ContentType.
type FindCalendarEventDuplicatesJSONRequestBody FindCalendarEventDuplicatesJSONBody

// MergeManyCalendarEventsJSONRequestBody defines body for MergeManyCalendarEvents for application/json ContentType.
type MergeManyCalendarEventsJSONRequestBody MergeManyCalendarEventsJSONBody

// UpdateOneCalendarEventJSONRequestBody defines body for UpdateOneCalendarEvent for application/json ContentType.
type UpdateOneCalendarEventJSONRequestBody = CalendarEventForUpdate

// UpdateManyCompaniesJSONRequestBody defines body for UpdateManyCompanies for application/json ContentType.
type UpdateManyCompaniesJSONRequestBody = CompanyForUpdate

// CreateOneCompanyJSONRequestBody defines body for CreateOneCompany for application/json ContentType.
type CreateOneCompanyJSONRequestBody = Company

// FindCompanyDuplicatesJSONRequestBody defines body for FindCompanyDuplicates for application/json ContentType.
type FindCompanyDuplicatesJSONRequestBody FindCompanyDuplicatesJSONBody

// MergeManyCompaniesJSONRequestBody defines body for MergeManyCompanies for application/json ContentType.
type MergeManyCompaniesJSONRequestBody MergeManyCompaniesJSONBody

// UpdateOneCompanyJSONRequestBody defines body for UpdateOneCompany for application/json ContentType.
type UpdateOneCompanyJSONRequestBody = CompanyForUpdate

// UpdateManyConnectedAccountsJSONRequestBody defines body for UpdateManyConnectedAccounts for application/json ContentType.
type UpdateManyConnectedAccountsJSONRequestBody = ConnectedAccountForUpdate

// CreateOneConnectedAccountJSONRequestBody defines body for CreateOneConnectedAccount for application/json ContentType.
type CreateOneConnectedAccountJSONRequestBody = ConnectedAccount

// FindConnectedAccountDuplicatesJSONRequestBody defines body for FindConnectedAccountDuplicates for application/json ContentType.
type FindConnectedAccountDuplicatesJSONRequestBody FindConnectedAccountDuplicatesJSONBody

// MergeManyConnectedAccountsJSONRequestBody defines body for MergeManyConnectedAccounts for application/json ContentType.
type MergeManyConnectedAccountsJSONRequestBody MergeManyConnectedAccountsJSONBody

// UpdateOneConnectedAccountJSONRequestBody defines body for UpdateOneConnectedAccount for application/json ContentType.
type UpdateOneConnectedAccountJSONRequestBody = ConnectedAccountForUpdate

// UpdateManyDashboardsJSONRequestBody defines body for UpdateManyDashboards for application/json ContentType.
type UpdateManyDashboardsJSONRequestBody = DashboardForUpdate

// CreateOneDashboardJSONRequestBody defines body for CreateOneDashboard for application/json ContentType.
type CreateOneDashboardJSONRequestBody = Dashboard

// FindDashboardDuplicatesJSONRequestBody defines body for FindDashboardDuplicates for application/json ContentType.
type FindDashboardDuplicatesJSONRequestBody FindDashboardDuplicatesJSONBody

// MergeManyDashboardsJSONRequestBody defines body for MergeManyDashboards for application/json ContentType.
type MergeManyDashboardsJSONRequestBody MergeManyDashboardsJSONBody

// UpdateOneDashboardJSONRequestBody defines body for UpdateOneDashboard for application/json ContentType.
type UpdateOneDashboardJSONRequestBody = DashboardForUpdate

// UpdateManyFavoriteFoldersJSONRequestBody defines body for UpdateManyFavoriteFolders for application/json ContentType.
type UpdateManyFavoriteFoldersJSONRequestBody = FavoriteFolderForUpdate

// CreateOneFavoriteFolderJSONRequestBody defines body for CreateOneFavoriteFolder for application/json ContentType.
type CreateOneFavoriteFolderJSONRequestBody = FavoriteFolder

// FindFavoriteFolderDuplicatesJSONRequestBody defines body for FindFavoriteFolderDuplicates for application/json ContentType.
type FindFavoriteFolderDuplicatesJSONRequestBody FindFavoriteFolderDuplicatesJSONBody

// MergeManyFavoriteFoldersJSONRequestBody defines body for MergeManyFavoriteFolders for application/json ContentType.
type MergeManyFavoriteFoldersJSONRequestBody MergeManyFavoriteFoldersJSONBody

// UpdateOneFavoriteFolderJSONRequestBody defines body for UpdateOneFavoriteFolder for application/json ContentType.
type UpdateOneFavoriteFolderJSONRequestBody = FavoriteFolderForUpdate

// UpdateManyFavoritesJSONRequestBody defines body for UpdateManyFavorites for application/json ContentType.
type UpdateManyFavoritesJSONRequestBody = FavoriteForUpdate

// CreateOneFavoriteJSONRequestBody defines body for CreateOneFavorite for application/json ContentType.
type CreateOneFavoriteJSONRequestBody = Favorite

// FindFavoriteDuplicatesJSONRequestBody defines body for FindFavoriteDuplicates for application/json ContentType.
type FindFavoriteDuplicatesJSONRequestBody FindFavoriteDuplicatesJSONBody

// MergeManyFavoritesJSONRequestBody defines body for MergeManyFavorites for application/json ContentType.
type MergeManyFavoritesJSONRequestBody MergeManyFavoritesJSONBody

// UpdateOneFavoriteJSONRequestBody defines body for UpdateOneFavorite for application/json ContentType.
type UpdateOneFavoriteJSONRequestBody = FavoriteForUpdate

// UpdateManyMessageChannelMessageAssociationsJSONRequestBody defines body for UpdateManyMessageChannelMessageAssociations for application/json ContentType.
type UpdateManyMessageChannelMessageAssociationsJSONRequestBody = MessageChannelMessageAssociationForUpdate

// CreateOneMessageChannelMessageAssociationJSONRequestBody defines body for CreateOneMessageChannelMessageAssociation for application/json ContentType.
type CreateOneMessageChannelMessageAssociationJSONRequestBody = MessageChannelMessageAssociation

// FindMessageChannelMessageAssociationDuplicatesJSONRequestBody defines body for FindMessageChannelMessageAssociationDuplicates for application/json ContentType.
type FindMessageChannelMessageAssociationDuplicatesJSONRequestBody FindMessageChannelMessageAssociationDuplicatesJSONBody

// MergeManyMessageChannelMessageAssociationsJSONRequestBody defines body for MergeManyMessageChannelMessageAssociations for application/json ContentType.
type MergeManyMessageChannelMessageAssociationsJSONRequestBody MergeManyMessageChannelMessageAssociationsJSONBody

// UpdateOneMessageChannelMessageAssociationJSONRequestBody defines body for UpdateOneMessageChannelMessageAssociation for application/json ContentType.
type UpdateOneMessageChannelMessageAssociationJSONRequestBody = MessageChannelMessageAssociationForUpdate

// UpdateManyMessageChannelsJSONRequestBody defines body for UpdateManyMessageChannels for application/json ContentType.
type UpdateManyMessageChannelsJSONRequestBody = MessageChannelForUpdate

// CreateOneMessageChannelJSONRequestBody defines body for CreateOneMessageChannel for application/json ContentType.
type CreateOneMessageChannelJSONRequestBody = MessageChannel

// FindMessageChannelDuplicatesJSONRequestBody defines body for FindMessageChannelDuplicates for application/json ContentType.
type FindMessageChannelDuplicatesJSONRequestBody FindMessageChannelDuplicatesJSONBody

// MergeManyMessageChannelsJSONRequestBody defines body for MergeManyMessageChannels for application/json ContentType.
type MergeManyMessageChannelsJSONRequestBody MergeManyMessageChannelsJSONBody

// UpdateOneMessageChannelJSONRequestBody defines body for UpdateOneMessageChannel for application/json ContentType.
type UpdateOneMessageChannelJSONRequestBody = MessageChannelForUpdate

// UpdateManyMessageFoldersJSONRequestBody defines body for UpdateManyMessageFolders for application/json ContentType.
type UpdateManyMessageFoldersJSONRequestBody = MessageFolderForUpdate

// CreateOneMessageFolderJSONRequestBody defines body for CreateOneMessageFolder for application/json ContentType.
type CreateOneMessageFolderJSONRequestBody = MessageFolder

// FindMessageFolderDuplicatesJSONRequestBody defines body for FindMessageFolderDuplicates for application/json ContentType.
type FindMessageFolderDuplicatesJSONRequestBody FindMessageFolderDuplicatesJSONBody

// MergeManyMessageFoldersJSONRequestBody defines body for MergeManyMessageFolders for application/json ContentType.
type MergeManyMessageFoldersJSONRequestBody MergeManyMessageFoldersJSONBody

// UpdateOneMessageFolderJSONRequestBody defines body for UpdateOneMessageFolder for application/json ContentType.
type UpdateOneMessageFolderJSONRequestBody = MessageFolderForUpdate

// UpdateManyMessageParticipantsJSONRequestBody defines body for UpdateManyMessageParticipants for application/json ContentType.
type UpdateManyMessageParticipantsJSONRequestBody = MessageParticipantForUpdate

// CreateOneMessageParticipantJSONRequestBody defines body for CreateOneMessageParticipant for application/json ContentType.
type CreateOneMessageParticipantJSONRequestBody = MessageParticipant

// FindMessageParticipantDuplicatesJSONRequestBody defines body for FindMessageParticipantDuplicates for application/json ContentType.
type FindMessageParticipantDuplicatesJSONRequestBody FindMessageParticipantDuplicatesJSONBody

// MergeManyMessageParticipantsJSONRequestBody defines body for MergeManyMessageParticipants for application/json ContentType.
type MergeManyMessageParticipantsJSONRequestBody MergeManyMessageParticipantsJSONBody

// UpdateOneMessageParticipantJSONRequestBody defines body for UpdateOneMessageParticipant for application/json ContentType.
type UpdateOneMessageParticipantJSONRequestBody = MessageParticipantForUpdate

// UpdateManyMessageThreadsJSONRequestBody defines body for UpdateManyMessageThreads for application/json ContentType.
type UpdateManyMessageThreadsJSONRequestBody = MessageThreadForUpdate

// CreateOneMessageThreadJSONRequestBody defines body for CreateOneMessageThread for application/json ContentType.
type CreateOneMessageThreadJSONRequestBody = MessageThread

// FindMessageThreadDuplicatesJSONRequestBody defines body for FindMessageThreadDuplicates for application/json ContentType.
type FindMessageThreadDuplicatesJSONRequestBody FindMessageThreadDuplicatesJSONBody

// MergeManyMessageThreadsJSONRequestBody defines body for MergeManyMessageThreads for application/json ContentType.
type MergeManyMessageThreadsJSONRequestBody MergeManyMessageThreadsJSONBody

// UpdateOneMessageThreadJSONRequestBody defines body for UpdateOneMessageThread for application/json ContentType.
type UpdateOneMessageThreadJSONRequestBody = MessageThreadForUpdate

// UpdateManyMessagesJSONRequestBody defines body for UpdateManyMessages for application/json ContentType.
type UpdateManyMessagesJSONRequestBody = MessageForUpdate

// CreateOneMessageJSONRequestBody defines body for CreateOneMessage for application/json ContentType.
type CreateOneMessageJSONRequestBody = Message

// FindMessageDuplicatesJSONRequestBody defines body for FindMessageDuplicates for application/json ContentType.
type FindMessageDuplicatesJSONRequestBody FindMessageDuplicatesJSONBody

// MergeManyMessagesJSONRequestBody defines body for MergeManyMessages for application/json ContentType.
type MergeManyMessagesJSONRequestBody MergeManyMessagesJSONBody

// UpdateOneMessageJSONRequestBody defines body for UpdateOneMessage for application/json ContentType.
type UpdateOneMessageJSONRequestBody = MessageForUpdate

// UpdateManyNoteTargetsJSONRequestBody defines body for UpdateManyNoteTargets for application/json ContentType.
type UpdateManyNoteTargetsJSONRequestBody = NoteTargetForUpdate

// CreateOneNoteTargetJSONRequestBody defines body for CreateOneNoteTarget for application/json ContentType.
type CreateOneNoteTargetJSONRequestBody = NoteTarget

// FindNoteTargetDuplicatesJSONRequestBody defines body for FindNoteTargetDuplicates for application/json ContentType.
type FindNoteTargetDuplicatesJSONRequestBody FindNoteTargetDuplicatesJSONBody

// MergeManyNoteTargetsJSONRequestBody defines body for MergeManyNoteTargets for application/json ContentType.
type MergeManyNoteTargetsJSONRequestBody MergeManyNoteTargetsJSONBody

// UpdateOneNoteTargetJSONRequestBody defines body for UpdateOneNoteTarget for application/json ContentType.
type UpdateOneNoteTargetJSONRequestBody = NoteTargetForUpdate

// UpdateManyNotesJSONRequestBody defines body for UpdateManyNotes for application/json ContentType.
type UpdateManyNotesJSONRequestBody = NoteForUpdate

// CreateOneNoteJSONRequestBody defines body for CreateOneNote for application/json ContentType.
type CreateOneNoteJSONRequestBody = Note

// FindNoteDuplicatesJSONRequestBody defines body for FindNoteDuplicates for application/json ContentType.
type FindNoteDuplicatesJSONRequestBody FindNoteDuplicatesJSONBody

// MergeManyNotesJSONRequestBody defines body for MergeManyNotes for application/json ContentType.
type MergeManyNotesJSONRequestBody MergeManyNotesJSONBody

// UpdateOneNoteJSONRequestBody defines body for UpdateOneNote for application/json ContentType.
type UpdateOneNoteJSONRequestBody = NoteForUpdate

// UpdateManyOpportunitiesJSONRequestBody defines body for UpdateManyOpportunities for application/json ContentType.
type UpdateManyOpportunitiesJSONRequestBody = OpportunityForUpdate

// CreateOneOpportunityJSONRequestBody defines body for CreateOneOpportunity for application/json ContentType.
type CreateOneOpportunityJSONRequestBody = Opportunity

// FindOpportunityDuplicatesJSONRequestBody defines body for FindOpportunityDuplicates for application/json ContentType.
type FindOpportunityDuplicatesJSONRequestBody FindOpportunityDuplicatesJSONBody

// MergeManyOpportunitiesJSONRequestBody defines body for MergeManyOpportunities for application/json ContentType.
type MergeManyOpportunitiesJSONRequestBody MergeManyOpportunitiesJSONBody

// UpdateOneOpportunityJSONRequestBody defines body for UpdateOneOpportunity for application/json ContentType.
type UpdateOneOpportunityJSONRequestBody = OpportunityForUpdate

// UpdateManyPeopleJSONRequestBody defines body for UpdateManyPeople for application/json ContentType.
type UpdateManyPeopleJSONRequestBody = PersonForUpdate

// CreateOnePersonJSONRequestBody defines body for CreateOnePerson for application/json ContentType.
type CreateOnePersonJSONRequestBody = Person

// FindPersonDuplicatesJSONRequestBody defines body for FindPersonDuplicates for application/json ContentType.
type FindPersonDuplicatesJSONRequestBody FindPersonDuplicatesJSONBody

// MergeManyPeopleJSONRequestBody defines body for MergeManyPeople for application/json ContentType.
type MergeManyPeopleJSONRequestBody MergeManyPeopleJSONBody

// UpdateOnePersonJSONRequestBody defines body for UpdateOnePerson for application/json ContentType.
type UpdateOnePersonJSONRequestBody = PersonForUpdate

// UpdateManyTaskTargetsJSONRequestBody defines body for UpdateManyTaskTargets for application/json ContentType.
type UpdateManyTaskTargetsJSONRequestBody = TaskTargetForUpdate

// CreateOneTaskTargetJSONRequestBody defines body for CreateOneTaskTarget for application/json ContentType.
type CreateOneTaskTargetJSONRequestBody = TaskTarget

// FindTaskTargetDuplicatesJSONRequestBody defines body for FindTaskTargetDuplicates for application/json ContentType.
type FindTaskTargetDuplicatesJSONRequestBody FindTaskTargetDuplicatesJSONBody

// MergeManyTaskTargetsJSONRequestBody defines body for MergeManyTaskTargets for application/json ContentType.
type MergeManyTaskTargetsJSONRequestBody MergeManyTaskTargetsJSONBody

// UpdateOneTaskTargetJSONRequestBody defines body for UpdateOneTaskTarget for application/json ContentType.
type UpdateOneTaskTargetJSONRequestBody = TaskTargetForUpdate

// UpdateManyTasksJSONRequestBody defines body for UpdateManyTasks for application/json ContentType.
type UpdateManyTasksJSONRequestBody = TaskForUpdate

// CreateOneTaskJSONRequestBody defines body for CreateOneTask for application/json ContentType.
type CreateOneTaskJSONRequestBody = Task

// FindTaskDuplicatesJSONRequestBody defines body for FindTaskDuplicates for application/json ContentType.
type FindTaskDuplicatesJSONRequestBody FindTaskDuplicatesJSONBody

// MergeManyTasksJSONRequestBody defines body for MergeManyTasks for application/json ContentType.
type MergeManyTasksJSONRequestBody MergeManyTasksJSONBody

// UpdateOneTaskJSONRequestBody defines body for UpdateOneTask for application/json ContentType.
type UpdateOneTaskJSONRequestBody = TaskForUpdate

// UpdateManyTimelineActivitiesJSONRequestBody defines body for UpdateManyTimelineActivities for application/json ContentType.
type UpdateManyTimelineActivitiesJSONRequestBody = TimelineActivityForUpdate

// CreateOneTimelineActivityJSONRequestBody defines body for CreateOneTimelineActivity for application/json ContentType.
type CreateOneTimelineActivityJSONRequestBody = TimelineActivity

// FindTimelineActivityDuplicatesJSONRequestBody defines body for FindTimelineActivityDuplicates for application/json ContentType.
type FindTimelineActivityDuplicatesJSONRequestBody FindTimelineActivityDuplicatesJSONBody

// MergeManyTimelineActivitiesJSONRequestBody defines body for MergeManyTimelineActivities for application/json ContentType.
type MergeManyTimelineActivitiesJSONRequestBody MergeManyTimelineActivitiesJSONBody

// UpdateOneTimelineActivityJSONRequestBody defines body for UpdateOneTimelineActivity for application/json ContentType.
type UpdateOneTimelineActivityJSONRequestBody = TimelineActivityForUpdate

// UpdateManyWorkflowAutomatedTriggersJSONRequestBody defines body for UpdateManyWorkflowAutomatedTriggers for application/json ContentType.
type UpdateManyWorkflowAutomatedTriggersJSONRequestBody = WorkflowAutomatedTriggerForUpdate

// CreateOneWorkflowAutomatedTriggerJSONRequestBody defines body for CreateOneWorkflowAutomatedTrigger for application/json ContentType.
type CreateOneWorkflowAutomatedTriggerJSONRequestBody = WorkflowAutomatedTrigger

// FindWorkflowAutomatedTriggerDuplicatesJSONRequestBody defines body for FindWorkflowAutomatedTriggerDuplicates for application/json ContentType.
type FindWorkflowAutomatedTriggerDuplicatesJSONRequestBody FindWorkflowAutomatedTriggerDuplicatesJSONBody

// MergeManyWorkflowAutomatedTriggersJSONRequestBody defines body for MergeManyWorkflowAutomatedTriggers for application/json ContentType.
type MergeManyWorkflowAutomatedTriggersJSONRequestBody MergeManyWorkflowAutomatedTriggersJSONBody

// UpdateOneWorkflowAutomatedTriggerJSONRequestBody defines body for UpdateOneWorkflowAutomatedTrigger for application/json ContentType.
type UpdateOneWorkflowAutomatedTriggerJSONRequestBody = WorkflowAutomatedTriggerForUpdate

// UpdateManyWorkflowRunsJSONRequestBody defines body for UpdateManyWorkflowRuns for application/json ContentType.
type UpdateManyWorkflowRunsJSONRequestBody = WorkflowRunForUpdate

// CreateOneWorkflowRunJSONRequestBody defines body for CreateOneWorkflowRun for application/json ContentType.
type CreateOneWorkflowRunJSONRequestBody = WorkflowRun

// FindWorkflowRunDuplicatesJSONRequestBody defines body for FindWorkflowRunDuplicates for application/json ContentType.
type FindWorkflowRunDuplicatesJSONRequestBody FindWorkflowRunDuplicatesJSONBody

// MergeManyWorkflowRunsJSONRequestBody defines body for MergeManyWorkflowRuns for application/json ContentType.
type MergeManyWorkflowRunsJSONRequestBody MergeManyWorkflowRunsJSONBody

// UpdateOneWorkflowRunJSONRequestBody defines body for UpdateOneWorkflowRun for application/json ContentType.
type UpdateOneWorkflowRunJSONRequestBody = WorkflowRunForUpdate

// UpdateManyWorkflowVersionsJSONRequestBody defines body for UpdateManyWorkflowVersions for application/json ContentType.
type UpdateManyWorkflowVersionsJSONRequestBody = WorkflowVersionForUpdate

// CreateOneWorkflowVersionJSONRequestBody defines body for CreateOneWorkflowVersion for application/json ContentType.
type CreateOneWorkflowVersionJSONRequestBody = WorkflowVersion

// FindWorkflowVersionDuplicatesJSONRequestBody defines body for FindWorkflowVersionDuplicates for application/json ContentType.
type FindWorkflowVersionDuplicatesJSONRequestBody FindWorkflowVersionDuplicatesJSONBody

// MergeManyWorkflowVersionsJSONRequestBody defines body for MergeManyWorkflowVersions for application/json ContentType.
type MergeManyWorkflowVersionsJSONRequestBody MergeManyWorkflowVersionsJSONBody

// UpdateOneWorkflowVersionJSONRequestBody defines body for UpdateOneWorkflowVersion for application/json ContentType.
type UpdateOneWorkflowVersionJSONRequestBody = WorkflowVersionForUpdate

// UpdateManyWorkflowsJSONRequestBody defines body for UpdateManyWorkflows for application/json ContentType.
type UpdateManyWorkflowsJSONRequestBody = WorkflowForUpdate

// CreateOneWorkflowJSONRequestBody defines body for CreateOneWorkflow for application/json ContentType.
type CreateOneWorkflowJSONRequestBody = Workflow

// FindWorkflowDuplicatesJSONRequestBody defines body for FindWorkflowDuplicates for application/json ContentType.
type FindWorkflowDuplicatesJSONRequestBody FindWorkflowDuplicatesJSONBody

// MergeManyWorkflowsJSONRequestBody defines body for MergeManyWorkflows for application/json ContentType.
type MergeManyWorkflowsJSONRequestBody MergeManyWorkflowsJSONBody

// UpdateOneWorkflowJSONRequestBody defines body for UpdateOneWorkflow for application/json ContentType.
type UpdateOneWorkflowJSONRequestBody = WorkflowForUpdate

// UpdateManyWorkspaceMembersJSONRequestBody defines body for UpdateManyWorkspaceMembers for application/json ContentType.
type UpdateManyWorkspaceMembersJSONRequestBody = WorkspaceMemberForUpdate

// CreateOneWorkspaceMemberJSONRequestBody defines body for CreateOneWorkspaceMember for application/json ContentType.
type CreateOneWorkspaceMemberJSONRequestBody = WorkspaceMember

// FindWorkspaceMemberDuplicatesJSONRequestBody defines body for FindWorkspaceMemberDuplicates for application/json ContentType.
type FindWorkspaceMemberDuplicatesJSONRequestBody FindWorkspaceMemberDuplicatesJSONBody

// MergeManyWorkspaceMembersJSONRequestBody defines body for MergeManyWorkspaceMembers for application/json ContentType.
type MergeManyWorkspaceMembersJSONRequestBody MergeManyWorkspaceMembersJSONBody

// UpdateOneWorkspaceMemberJSONRequestBody defines body for UpdateOneWorkspaceMember for application/json ContentType.
type UpdateOneWorkspaceMemberJSONRequestBody = WorkspaceMemberForUpdate
