package twentygo

import (
	"encoding/json"

	"github.com/oapi-codegen/runtime"
)

// AsWorkspaceMemberForResponse returns the union data inside the AttachmentForResponse_Author as a WorkspaceMemberForResponse
func (t AttachmentForResponse_Author) AsWorkspaceMemberForResponse() (WorkspaceMemberForResponse, error) {
	var body WorkspaceMemberForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkspaceMemberForResponse overwrites any union data inside the AttachmentForResponse_Author as the provided WorkspaceMemberForResponse
func (t *AttachmentForResponse_Author) FromWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkspaceMemberForResponse performs a merge with any union data inside the AttachmentForResponse_Author, using the provided WorkspaceMemberForResponse
func (t *AttachmentForResponse_Author) MergeWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AttachmentForResponse_Author) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AttachmentForResponse_Author) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCompanyForResponse returns the union data inside the AttachmentForResponse_Company as a CompanyForResponse
func (t AttachmentForResponse_Company) AsCompanyForResponse() (CompanyForResponse, error) {
	var body CompanyForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCompanyForResponse overwrites any union data inside the AttachmentForResponse_Company as the provided CompanyForResponse
func (t *AttachmentForResponse_Company) FromCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCompanyForResponse performs a merge with any union data inside the AttachmentForResponse_Company, using the provided CompanyForResponse
func (t *AttachmentForResponse_Company) MergeCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AttachmentForResponse_Company) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AttachmentForResponse_Company) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsDashboardForResponse returns the union data inside the AttachmentForResponse_Dashboard as a DashboardForResponse
func (t AttachmentForResponse_Dashboard) AsDashboardForResponse() (DashboardForResponse, error) {
	var body DashboardForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDashboardForResponse overwrites any union data inside the AttachmentForResponse_Dashboard as the provided DashboardForResponse
func (t *AttachmentForResponse_Dashboard) FromDashboardForResponse(v DashboardForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDashboardForResponse performs a merge with any union data inside the AttachmentForResponse_Dashboard, using the provided DashboardForResponse
func (t *AttachmentForResponse_Dashboard) MergeDashboardForResponse(v DashboardForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AttachmentForResponse_Dashboard) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AttachmentForResponse_Dashboard) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsNoteForResponse returns the union data inside the AttachmentForResponse_Note as a NoteForResponse
func (t AttachmentForResponse_Note) AsNoteForResponse() (NoteForResponse, error) {
	var body NoteForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNoteForResponse overwrites any union data inside the AttachmentForResponse_Note as the provided NoteForResponse
func (t *AttachmentForResponse_Note) FromNoteForResponse(v NoteForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNoteForResponse performs a merge with any union data inside the AttachmentForResponse_Note, using the provided NoteForResponse
func (t *AttachmentForResponse_Note) MergeNoteForResponse(v NoteForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AttachmentForResponse_Note) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AttachmentForResponse_Note) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOpportunityForResponse returns the union data inside the AttachmentForResponse_Opportunity as a OpportunityForResponse
func (t AttachmentForResponse_Opportunity) AsOpportunityForResponse() (OpportunityForResponse, error) {
	var body OpportunityForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpportunityForResponse overwrites any union data inside the AttachmentForResponse_Opportunity as the provided OpportunityForResponse
func (t *AttachmentForResponse_Opportunity) FromOpportunityForResponse(v OpportunityForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpportunityForResponse performs a merge with any union data inside the AttachmentForResponse_Opportunity, using the provided OpportunityForResponse
func (t *AttachmentForResponse_Opportunity) MergeOpportunityForResponse(v OpportunityForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AttachmentForResponse_Opportunity) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AttachmentForResponse_Opportunity) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPersonForResponse returns the union data inside the AttachmentForResponse_Person as a PersonForResponse
func (t AttachmentForResponse_Person) AsPersonForResponse() (PersonForResponse, error) {
	var body PersonForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPersonForResponse overwrites any union data inside the AttachmentForResponse_Person as the provided PersonForResponse
func (t *AttachmentForResponse_Person) FromPersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePersonForResponse performs a merge with any union data inside the AttachmentForResponse_Person, using the provided PersonForResponse
func (t *AttachmentForResponse_Person) MergePersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AttachmentForResponse_Person) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AttachmentForResponse_Person) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsTaskForResponse returns the union data inside the AttachmentForResponse_Task as a TaskForResponse
func (t AttachmentForResponse_Task) AsTaskForResponse() (TaskForResponse, error) {
	var body TaskForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTaskForResponse overwrites any union data inside the AttachmentForResponse_Task as the provided TaskForResponse
func (t *AttachmentForResponse_Task) FromTaskForResponse(v TaskForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTaskForResponse performs a merge with any union data inside the AttachmentForResponse_Task, using the provided TaskForResponse
func (t *AttachmentForResponse_Task) MergeTaskForResponse(v TaskForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AttachmentForResponse_Task) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AttachmentForResponse_Task) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkflowForResponse returns the union data inside the AttachmentForResponse_Workflow as a WorkflowForResponse
func (t AttachmentForResponse_Workflow) AsWorkflowForResponse() (WorkflowForResponse, error) {
	var body WorkflowForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkflowForResponse overwrites any union data inside the AttachmentForResponse_Workflow as the provided WorkflowForResponse
func (t *AttachmentForResponse_Workflow) FromWorkflowForResponse(v WorkflowForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkflowForResponse performs a merge with any union data inside the AttachmentForResponse_Workflow, using the provided WorkflowForResponse
func (t *AttachmentForResponse_Workflow) MergeWorkflowForResponse(v WorkflowForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t AttachmentForResponse_Workflow) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AttachmentForResponse_Workflow) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkspaceMemberForResponse returns the union data inside the BlocklistForResponse_WorkspaceMember as a WorkspaceMemberForResponse
func (t BlocklistForResponse_WorkspaceMember) AsWorkspaceMemberForResponse() (WorkspaceMemberForResponse, error) {
	var body WorkspaceMemberForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkspaceMemberForResponse overwrites any union data inside the BlocklistForResponse_WorkspaceMember as the provided WorkspaceMemberForResponse
func (t *BlocklistForResponse_WorkspaceMember) FromWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkspaceMemberForResponse performs a merge with any union data inside the BlocklistForResponse_WorkspaceMember, using the provided WorkspaceMemberForResponse
func (t *BlocklistForResponse_WorkspaceMember) MergeWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t BlocklistForResponse_WorkspaceMember) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *BlocklistForResponse_WorkspaceMember) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCalendarChannelForResponse returns the union data inside the CalendarChannelEventAssociationForResponse_CalendarChannel as a CalendarChannelForResponse
func (t CalendarChannelEventAssociationForResponse_CalendarChannel) AsCalendarChannelForResponse() (CalendarChannelForResponse, error) {
	var body CalendarChannelForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCalendarChannelForResponse overwrites any union data inside the CalendarChannelEventAssociationForResponse_CalendarChannel as the provided CalendarChannelForResponse
func (t *CalendarChannelEventAssociationForResponse_CalendarChannel) FromCalendarChannelForResponse(v CalendarChannelForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCalendarChannelForResponse performs a merge with any union data inside the CalendarChannelEventAssociationForResponse_CalendarChannel, using the provided CalendarChannelForResponse
func (t *CalendarChannelEventAssociationForResponse_CalendarChannel) MergeCalendarChannelForResponse(v CalendarChannelForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CalendarChannelEventAssociationForResponse_CalendarChannel) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CalendarChannelEventAssociationForResponse_CalendarChannel) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCalendarEventForResponse returns the union data inside the CalendarChannelEventAssociationForResponse_CalendarEvent as a CalendarEventForResponse
func (t CalendarChannelEventAssociationForResponse_CalendarEvent) AsCalendarEventForResponse() (CalendarEventForResponse, error) {
	var body CalendarEventForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCalendarEventForResponse overwrites any union data inside the CalendarChannelEventAssociationForResponse_CalendarEvent as the provided CalendarEventForResponse
func (t *CalendarChannelEventAssociationForResponse_CalendarEvent) FromCalendarEventForResponse(v CalendarEventForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCalendarEventForResponse performs a merge with any union data inside the CalendarChannelEventAssociationForResponse_CalendarEvent, using the provided CalendarEventForResponse
func (t *CalendarChannelEventAssociationForResponse_CalendarEvent) MergeCalendarEventForResponse(v CalendarEventForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CalendarChannelEventAssociationForResponse_CalendarEvent) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CalendarChannelEventAssociationForResponse_CalendarEvent) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsConnectedAccountForResponse returns the union data inside the CalendarChannelForResponse_ConnectedAccount as a ConnectedAccountForResponse
func (t CalendarChannelForResponse_ConnectedAccount) AsConnectedAccountForResponse() (ConnectedAccountForResponse, error) {
	var body ConnectedAccountForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromConnectedAccountForResponse overwrites any union data inside the CalendarChannelForResponse_ConnectedAccount as the provided ConnectedAccountForResponse
func (t *CalendarChannelForResponse_ConnectedAccount) FromConnectedAccountForResponse(v ConnectedAccountForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeConnectedAccountForResponse performs a merge with any union data inside the CalendarChannelForResponse_ConnectedAccount, using the provided ConnectedAccountForResponse
func (t *CalendarChannelForResponse_ConnectedAccount) MergeConnectedAccountForResponse(v ConnectedAccountForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CalendarChannelForResponse_ConnectedAccount) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CalendarChannelForResponse_ConnectedAccount) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCalendarEventForResponse returns the union data inside the CalendarEventParticipantForResponse_CalendarEvent as a CalendarEventForResponse
func (t CalendarEventParticipantForResponse_CalendarEvent) AsCalendarEventForResponse() (CalendarEventForResponse, error) {
	var body CalendarEventForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCalendarEventForResponse overwrites any union data inside the CalendarEventParticipantForResponse_CalendarEvent as the provided CalendarEventForResponse
func (t *CalendarEventParticipantForResponse_CalendarEvent) FromCalendarEventForResponse(v CalendarEventForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCalendarEventForResponse performs a merge with any union data inside the CalendarEventParticipantForResponse_CalendarEvent, using the provided CalendarEventForResponse
func (t *CalendarEventParticipantForResponse_CalendarEvent) MergeCalendarEventForResponse(v CalendarEventForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CalendarEventParticipantForResponse_CalendarEvent) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CalendarEventParticipantForResponse_CalendarEvent) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPersonForResponse returns the union data inside the CalendarEventParticipantForResponse_Person as a PersonForResponse
func (t CalendarEventParticipantForResponse_Person) AsPersonForResponse() (PersonForResponse, error) {
	var body PersonForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPersonForResponse overwrites any union data inside the CalendarEventParticipantForResponse_Person as the provided PersonForResponse
func (t *CalendarEventParticipantForResponse_Person) FromPersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePersonForResponse performs a merge with any union data inside the CalendarEventParticipantForResponse_Person, using the provided PersonForResponse
func (t *CalendarEventParticipantForResponse_Person) MergePersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CalendarEventParticipantForResponse_Person) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CalendarEventParticipantForResponse_Person) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkspaceMemberForResponse returns the union data inside the CalendarEventParticipantForResponse_WorkspaceMember as a WorkspaceMemberForResponse
func (t CalendarEventParticipantForResponse_WorkspaceMember) AsWorkspaceMemberForResponse() (WorkspaceMemberForResponse, error) {
	var body WorkspaceMemberForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkspaceMemberForResponse overwrites any union data inside the CalendarEventParticipantForResponse_WorkspaceMember as the provided WorkspaceMemberForResponse
func (t *CalendarEventParticipantForResponse_WorkspaceMember) FromWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkspaceMemberForResponse performs a merge with any union data inside the CalendarEventParticipantForResponse_WorkspaceMember, using the provided WorkspaceMemberForResponse
func (t *CalendarEventParticipantForResponse_WorkspaceMember) MergeWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CalendarEventParticipantForResponse_WorkspaceMember) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CalendarEventParticipantForResponse_WorkspaceMember) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkspaceMemberForResponse returns the union data inside the CompanyForResponse_AccountOwner as a WorkspaceMemberForResponse
func (t CompanyForResponse_AccountOwner) AsWorkspaceMemberForResponse() (WorkspaceMemberForResponse, error) {
	var body WorkspaceMemberForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkspaceMemberForResponse overwrites any union data inside the CompanyForResponse_AccountOwner as the provided WorkspaceMemberForResponse
func (t *CompanyForResponse_AccountOwner) FromWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkspaceMemberForResponse performs a merge with any union data inside the CompanyForResponse_AccountOwner, using the provided WorkspaceMemberForResponse
func (t *CompanyForResponse_AccountOwner) MergeWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CompanyForResponse_AccountOwner) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CompanyForResponse_AccountOwner) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkspaceMemberForResponse returns the union data inside the ConnectedAccountForResponse_AccountOwner as a WorkspaceMemberForResponse
func (t ConnectedAccountForResponse_AccountOwner) AsWorkspaceMemberForResponse() (WorkspaceMemberForResponse, error) {
	var body WorkspaceMemberForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkspaceMemberForResponse overwrites any union data inside the ConnectedAccountForResponse_AccountOwner as the provided WorkspaceMemberForResponse
func (t *ConnectedAccountForResponse_AccountOwner) FromWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkspaceMemberForResponse performs a merge with any union data inside the ConnectedAccountForResponse_AccountOwner, using the provided WorkspaceMemberForResponse
func (t *ConnectedAccountForResponse_AccountOwner) MergeWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ConnectedAccountForResponse_AccountOwner) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ConnectedAccountForResponse_AccountOwner) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCompanyForResponse returns the union data inside the FavoriteForResponse_Company as a CompanyForResponse
func (t FavoriteForResponse_Company) AsCompanyForResponse() (CompanyForResponse, error) {
	var body CompanyForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCompanyForResponse overwrites any union data inside the FavoriteForResponse_Company as the provided CompanyForResponse
func (t *FavoriteForResponse_Company) FromCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCompanyForResponse performs a merge with any union data inside the FavoriteForResponse_Company, using the provided CompanyForResponse
func (t *FavoriteForResponse_Company) MergeCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FavoriteForResponse_Company) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FavoriteForResponse_Company) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsDashboardForResponse returns the union data inside the FavoriteForResponse_Dashboard as a DashboardForResponse
func (t FavoriteForResponse_Dashboard) AsDashboardForResponse() (DashboardForResponse, error) {
	var body DashboardForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDashboardForResponse overwrites any union data inside the FavoriteForResponse_Dashboard as the provided DashboardForResponse
func (t *FavoriteForResponse_Dashboard) FromDashboardForResponse(v DashboardForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDashboardForResponse performs a merge with any union data inside the FavoriteForResponse_Dashboard, using the provided DashboardForResponse
func (t *FavoriteForResponse_Dashboard) MergeDashboardForResponse(v DashboardForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FavoriteForResponse_Dashboard) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FavoriteForResponse_Dashboard) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFavoriteFolderForResponse returns the union data inside the FavoriteForResponse_FavoriteFolder as a FavoriteFolderForResponse
func (t FavoriteForResponse_FavoriteFolder) AsFavoriteFolderForResponse() (FavoriteFolderForResponse, error) {
	var body FavoriteFolderForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFavoriteFolderForResponse overwrites any union data inside the FavoriteForResponse_FavoriteFolder as the provided FavoriteFolderForResponse
func (t *FavoriteForResponse_FavoriteFolder) FromFavoriteFolderForResponse(v FavoriteFolderForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFavoriteFolderForResponse performs a merge with any union data inside the FavoriteForResponse_FavoriteFolder, using the provided FavoriteFolderForResponse
func (t *FavoriteForResponse_FavoriteFolder) MergeFavoriteFolderForResponse(v FavoriteFolderForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FavoriteForResponse_FavoriteFolder) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FavoriteForResponse_FavoriteFolder) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkspaceMemberForResponse returns the union data inside the FavoriteForResponse_ForWorkspaceMember as a WorkspaceMemberForResponse
func (t FavoriteForResponse_ForWorkspaceMember) AsWorkspaceMemberForResponse() (WorkspaceMemberForResponse, error) {
	var body WorkspaceMemberForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkspaceMemberForResponse overwrites any union data inside the FavoriteForResponse_ForWorkspaceMember as the provided WorkspaceMemberForResponse
func (t *FavoriteForResponse_ForWorkspaceMember) FromWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkspaceMemberForResponse performs a merge with any union data inside the FavoriteForResponse_ForWorkspaceMember, using the provided WorkspaceMemberForResponse
func (t *FavoriteForResponse_ForWorkspaceMember) MergeWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FavoriteForResponse_ForWorkspaceMember) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FavoriteForResponse_ForWorkspaceMember) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsNoteForResponse returns the union data inside the FavoriteForResponse_Note as a NoteForResponse
func (t FavoriteForResponse_Note) AsNoteForResponse() (NoteForResponse, error) {
	var body NoteForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNoteForResponse overwrites any union data inside the FavoriteForResponse_Note as the provided NoteForResponse
func (t *FavoriteForResponse_Note) FromNoteForResponse(v NoteForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNoteForResponse performs a merge with any union data inside the FavoriteForResponse_Note, using the provided NoteForResponse
func (t *FavoriteForResponse_Note) MergeNoteForResponse(v NoteForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FavoriteForResponse_Note) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FavoriteForResponse_Note) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOpportunityForResponse returns the union data inside the FavoriteForResponse_Opportunity as a OpportunityForResponse
func (t FavoriteForResponse_Opportunity) AsOpportunityForResponse() (OpportunityForResponse, error) {
	var body OpportunityForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpportunityForResponse overwrites any union data inside the FavoriteForResponse_Opportunity as the provided OpportunityForResponse
func (t *FavoriteForResponse_Opportunity) FromOpportunityForResponse(v OpportunityForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpportunityForResponse performs a merge with any union data inside the FavoriteForResponse_Opportunity, using the provided OpportunityForResponse
func (t *FavoriteForResponse_Opportunity) MergeOpportunityForResponse(v OpportunityForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FavoriteForResponse_Opportunity) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FavoriteForResponse_Opportunity) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPersonForResponse returns the union data inside the FavoriteForResponse_Person as a PersonForResponse
func (t FavoriteForResponse_Person) AsPersonForResponse() (PersonForResponse, error) {
	var body PersonForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPersonForResponse overwrites any union data inside the FavoriteForResponse_Person as the provided PersonForResponse
func (t *FavoriteForResponse_Person) FromPersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePersonForResponse performs a merge with any union data inside the FavoriteForResponse_Person, using the provided PersonForResponse
func (t *FavoriteForResponse_Person) MergePersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FavoriteForResponse_Person) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FavoriteForResponse_Person) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsTaskForResponse returns the union data inside the FavoriteForResponse_Task as a TaskForResponse
func (t FavoriteForResponse_Task) AsTaskForResponse() (TaskForResponse, error) {
	var body TaskForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTaskForResponse overwrites any union data inside the FavoriteForResponse_Task as the provided TaskForResponse
func (t *FavoriteForResponse_Task) FromTaskForResponse(v TaskForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTaskForResponse performs a merge with any union data inside the FavoriteForResponse_Task, using the provided TaskForResponse
func (t *FavoriteForResponse_Task) MergeTaskForResponse(v TaskForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FavoriteForResponse_Task) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FavoriteForResponse_Task) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkflowForResponse returns the union data inside the FavoriteForResponse_Workflow as a WorkflowForResponse
func (t FavoriteForResponse_Workflow) AsWorkflowForResponse() (WorkflowForResponse, error) {
	var body WorkflowForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkflowForResponse overwrites any union data inside the FavoriteForResponse_Workflow as the provided WorkflowForResponse
func (t *FavoriteForResponse_Workflow) FromWorkflowForResponse(v WorkflowForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkflowForResponse performs a merge with any union data inside the FavoriteForResponse_Workflow, using the provided WorkflowForResponse
func (t *FavoriteForResponse_Workflow) MergeWorkflowForResponse(v WorkflowForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FavoriteForResponse_Workflow) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FavoriteForResponse_Workflow) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkflowRunForResponse returns the union data inside the FavoriteForResponse_WorkflowRun as a WorkflowRunForResponse
func (t FavoriteForResponse_WorkflowRun) AsWorkflowRunForResponse() (WorkflowRunForResponse, error) {
	var body WorkflowRunForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkflowRunForResponse overwrites any union data inside the FavoriteForResponse_WorkflowRun as the provided WorkflowRunForResponse
func (t *FavoriteForResponse_WorkflowRun) FromWorkflowRunForResponse(v WorkflowRunForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkflowRunForResponse performs a merge with any union data inside the FavoriteForResponse_WorkflowRun, using the provided WorkflowRunForResponse
func (t *FavoriteForResponse_WorkflowRun) MergeWorkflowRunForResponse(v WorkflowRunForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FavoriteForResponse_WorkflowRun) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FavoriteForResponse_WorkflowRun) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkflowVersionForResponse returns the union data inside the FavoriteForResponse_WorkflowVersion as a WorkflowVersionForResponse
func (t FavoriteForResponse_WorkflowVersion) AsWorkflowVersionForResponse() (WorkflowVersionForResponse, error) {
	var body WorkflowVersionForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkflowVersionForResponse overwrites any union data inside the FavoriteForResponse_WorkflowVersion as the provided WorkflowVersionForResponse
func (t *FavoriteForResponse_WorkflowVersion) FromWorkflowVersionForResponse(v WorkflowVersionForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkflowVersionForResponse performs a merge with any union data inside the FavoriteForResponse_WorkflowVersion, using the provided WorkflowVersionForResponse
func (t *FavoriteForResponse_WorkflowVersion) MergeWorkflowVersionForResponse(v WorkflowVersionForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FavoriteForResponse_WorkflowVersion) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FavoriteForResponse_WorkflowVersion) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsConnectedAccountForResponse returns the union data inside the MessageChannelForResponse_ConnectedAccount as a ConnectedAccountForResponse
func (t MessageChannelForResponse_ConnectedAccount) AsConnectedAccountForResponse() (ConnectedAccountForResponse, error) {
	var body ConnectedAccountForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromConnectedAccountForResponse overwrites any union data inside the MessageChannelForResponse_ConnectedAccount as the provided ConnectedAccountForResponse
func (t *MessageChannelForResponse_ConnectedAccount) FromConnectedAccountForResponse(v ConnectedAccountForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeConnectedAccountForResponse performs a merge with any union data inside the MessageChannelForResponse_ConnectedAccount, using the provided ConnectedAccountForResponse
func (t *MessageChannelForResponse_ConnectedAccount) MergeConnectedAccountForResponse(v ConnectedAccountForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t MessageChannelForResponse_ConnectedAccount) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *MessageChannelForResponse_ConnectedAccount) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsMessageForResponse returns the union data inside the MessageChannelMessageAssociationForResponse_Message as a MessageForResponse
func (t MessageChannelMessageAssociationForResponse_Message) AsMessageForResponse() (MessageForResponse, error) {
	var body MessageForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMessageForResponse overwrites any union data inside the MessageChannelMessageAssociationForResponse_Message as the provided MessageForResponse
func (t *MessageChannelMessageAssociationForResponse_Message) FromMessageForResponse(v MessageForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMessageForResponse performs a merge with any union data inside the MessageChannelMessageAssociationForResponse_Message, using the provided MessageForResponse
func (t *MessageChannelMessageAssociationForResponse_Message) MergeMessageForResponse(v MessageForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t MessageChannelMessageAssociationForResponse_Message) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *MessageChannelMessageAssociationForResponse_Message) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsMessageChannelForResponse returns the union data inside the MessageChannelMessageAssociationForResponse_MessageChannel as a MessageChannelForResponse
func (t MessageChannelMessageAssociationForResponse_MessageChannel) AsMessageChannelForResponse() (MessageChannelForResponse, error) {
	var body MessageChannelForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMessageChannelForResponse overwrites any union data inside the MessageChannelMessageAssociationForResponse_MessageChannel as the provided MessageChannelForResponse
func (t *MessageChannelMessageAssociationForResponse_MessageChannel) FromMessageChannelForResponse(v MessageChannelForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMessageChannelForResponse performs a merge with any union data inside the MessageChannelMessageAssociationForResponse_MessageChannel, using the provided MessageChannelForResponse
func (t *MessageChannelMessageAssociationForResponse_MessageChannel) MergeMessageChannelForResponse(v MessageChannelForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t MessageChannelMessageAssociationForResponse_MessageChannel) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *MessageChannelMessageAssociationForResponse_MessageChannel) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsMessageChannelForResponse returns the union data inside the MessageFolderForResponse_MessageChannel as a MessageChannelForResponse
func (t MessageFolderForResponse_MessageChannel) AsMessageChannelForResponse() (MessageChannelForResponse, error) {
	var body MessageChannelForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMessageChannelForResponse overwrites any union data inside the MessageFolderForResponse_MessageChannel as the provided MessageChannelForResponse
func (t *MessageFolderForResponse_MessageChannel) FromMessageChannelForResponse(v MessageChannelForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMessageChannelForResponse performs a merge with any union data inside the MessageFolderForResponse_MessageChannel, using the provided MessageChannelForResponse
func (t *MessageFolderForResponse_MessageChannel) MergeMessageChannelForResponse(v MessageChannelForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t MessageFolderForResponse_MessageChannel) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *MessageFolderForResponse_MessageChannel) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsMessageThreadForResponse returns the union data inside the MessageForResponse_MessageThread as a MessageThreadForResponse
func (t MessageForResponse_MessageThread) AsMessageThreadForResponse() (MessageThreadForResponse, error) {
	var body MessageThreadForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMessageThreadForResponse overwrites any union data inside the MessageForResponse_MessageThread as the provided MessageThreadForResponse
func (t *MessageForResponse_MessageThread) FromMessageThreadForResponse(v MessageThreadForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMessageThreadForResponse performs a merge with any union data inside the MessageForResponse_MessageThread, using the provided MessageThreadForResponse
func (t *MessageForResponse_MessageThread) MergeMessageThreadForResponse(v MessageThreadForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t MessageForResponse_MessageThread) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *MessageForResponse_MessageThread) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsMessageForResponse returns the union data inside the MessageParticipantForResponse_Message as a MessageForResponse
func (t MessageParticipantForResponse_Message) AsMessageForResponse() (MessageForResponse, error) {
	var body MessageForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMessageForResponse overwrites any union data inside the MessageParticipantForResponse_Message as the provided MessageForResponse
func (t *MessageParticipantForResponse_Message) FromMessageForResponse(v MessageForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMessageForResponse performs a merge with any union data inside the MessageParticipantForResponse_Message, using the provided MessageForResponse
func (t *MessageParticipantForResponse_Message) MergeMessageForResponse(v MessageForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t MessageParticipantForResponse_Message) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *MessageParticipantForResponse_Message) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPersonForResponse returns the union data inside the MessageParticipantForResponse_Person as a PersonForResponse
func (t MessageParticipantForResponse_Person) AsPersonForResponse() (PersonForResponse, error) {
	var body PersonForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPersonForResponse overwrites any union data inside the MessageParticipantForResponse_Person as the provided PersonForResponse
func (t *MessageParticipantForResponse_Person) FromPersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePersonForResponse performs a merge with any union data inside the MessageParticipantForResponse_Person, using the provided PersonForResponse
func (t *MessageParticipantForResponse_Person) MergePersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t MessageParticipantForResponse_Person) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *MessageParticipantForResponse_Person) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkspaceMemberForResponse returns the union data inside the MessageParticipantForResponse_WorkspaceMember as a WorkspaceMemberForResponse
func (t MessageParticipantForResponse_WorkspaceMember) AsWorkspaceMemberForResponse() (WorkspaceMemberForResponse, error) {
	var body WorkspaceMemberForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkspaceMemberForResponse overwrites any union data inside the MessageParticipantForResponse_WorkspaceMember as the provided WorkspaceMemberForResponse
func (t *MessageParticipantForResponse_WorkspaceMember) FromWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkspaceMemberForResponse performs a merge with any union data inside the MessageParticipantForResponse_WorkspaceMember, using the provided WorkspaceMemberForResponse
func (t *MessageParticipantForResponse_WorkspaceMember) MergeWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t MessageParticipantForResponse_WorkspaceMember) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *MessageParticipantForResponse_WorkspaceMember) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCompanyForResponse returns the union data inside the NoteTargetForResponse_Company as a CompanyForResponse
func (t NoteTargetForResponse_Company) AsCompanyForResponse() (CompanyForResponse, error) {
	var body CompanyForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCompanyForResponse overwrites any union data inside the NoteTargetForResponse_Company as the provided CompanyForResponse
func (t *NoteTargetForResponse_Company) FromCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCompanyForResponse performs a merge with any union data inside the NoteTargetForResponse_Company, using the provided CompanyForResponse
func (t *NoteTargetForResponse_Company) MergeCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t NoteTargetForResponse_Company) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *NoteTargetForResponse_Company) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsNoteForResponse returns the union data inside the NoteTargetForResponse_Note as a NoteForResponse
func (t NoteTargetForResponse_Note) AsNoteForResponse() (NoteForResponse, error) {
	var body NoteForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNoteForResponse overwrites any union data inside the NoteTargetForResponse_Note as the provided NoteForResponse
func (t *NoteTargetForResponse_Note) FromNoteForResponse(v NoteForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNoteForResponse performs a merge with any union data inside the NoteTargetForResponse_Note, using the provided NoteForResponse
func (t *NoteTargetForResponse_Note) MergeNoteForResponse(v NoteForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t NoteTargetForResponse_Note) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *NoteTargetForResponse_Note) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOpportunityForResponse returns the union data inside the NoteTargetForResponse_Opportunity as a OpportunityForResponse
func (t NoteTargetForResponse_Opportunity) AsOpportunityForResponse() (OpportunityForResponse, error) {
	var body OpportunityForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpportunityForResponse overwrites any union data inside the NoteTargetForResponse_Opportunity as the provided OpportunityForResponse
func (t *NoteTargetForResponse_Opportunity) FromOpportunityForResponse(v OpportunityForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpportunityForResponse performs a merge with any union data inside the NoteTargetForResponse_Opportunity, using the provided OpportunityForResponse
func (t *NoteTargetForResponse_Opportunity) MergeOpportunityForResponse(v OpportunityForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t NoteTargetForResponse_Opportunity) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *NoteTargetForResponse_Opportunity) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPersonForResponse returns the union data inside the NoteTargetForResponse_Person as a PersonForResponse
func (t NoteTargetForResponse_Person) AsPersonForResponse() (PersonForResponse, error) {
	var body PersonForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPersonForResponse overwrites any union data inside the NoteTargetForResponse_Person as the provided PersonForResponse
func (t *NoteTargetForResponse_Person) FromPersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePersonForResponse performs a merge with any union data inside the NoteTargetForResponse_Person, using the provided PersonForResponse
func (t *NoteTargetForResponse_Person) MergePersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t NoteTargetForResponse_Person) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *NoteTargetForResponse_Person) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCompanyForResponse returns the union data inside the OpportunityForResponse_Company as a CompanyForResponse
func (t OpportunityForResponse_Company) AsCompanyForResponse() (CompanyForResponse, error) {
	var body CompanyForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCompanyForResponse overwrites any union data inside the OpportunityForResponse_Company as the provided CompanyForResponse
func (t *OpportunityForResponse_Company) FromCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCompanyForResponse performs a merge with any union data inside the OpportunityForResponse_Company, using the provided CompanyForResponse
func (t *OpportunityForResponse_Company) MergeCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t OpportunityForResponse_Company) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OpportunityForResponse_Company) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPersonForResponse returns the union data inside the OpportunityForResponse_PointOfContact as a PersonForResponse
func (t OpportunityForResponse_PointOfContact) AsPersonForResponse() (PersonForResponse, error) {
	var body PersonForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPersonForResponse overwrites any union data inside the OpportunityForResponse_PointOfContact as the provided PersonForResponse
func (t *OpportunityForResponse_PointOfContact) FromPersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePersonForResponse performs a merge with any union data inside the OpportunityForResponse_PointOfContact, using the provided PersonForResponse
func (t *OpportunityForResponse_PointOfContact) MergePersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t OpportunityForResponse_PointOfContact) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OpportunityForResponse_PointOfContact) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCompanyForResponse returns the union data inside the PersonForResponse_Company as a CompanyForResponse
func (t PersonForResponse_Company) AsCompanyForResponse() (CompanyForResponse, error) {
	var body CompanyForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCompanyForResponse overwrites any union data inside the PersonForResponse_Company as the provided CompanyForResponse
func (t *PersonForResponse_Company) FromCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCompanyForResponse performs a merge with any union data inside the PersonForResponse_Company, using the provided CompanyForResponse
func (t *PersonForResponse_Company) MergeCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t PersonForResponse_Company) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PersonForResponse_Company) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkspaceMemberForResponse returns the union data inside the TaskForResponse_Assignee as a WorkspaceMemberForResponse
func (t TaskForResponse_Assignee) AsWorkspaceMemberForResponse() (WorkspaceMemberForResponse, error) {
	var body WorkspaceMemberForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkspaceMemberForResponse overwrites any union data inside the TaskForResponse_Assignee as the provided WorkspaceMemberForResponse
func (t *TaskForResponse_Assignee) FromWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkspaceMemberForResponse performs a merge with any union data inside the TaskForResponse_Assignee, using the provided WorkspaceMemberForResponse
func (t *TaskForResponse_Assignee) MergeWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t TaskForResponse_Assignee) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *TaskForResponse_Assignee) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCompanyForResponse returns the union data inside the TaskTargetForResponse_Company as a CompanyForResponse
func (t TaskTargetForResponse_Company) AsCompanyForResponse() (CompanyForResponse, error) {
	var body CompanyForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCompanyForResponse overwrites any union data inside the TaskTargetForResponse_Company as the provided CompanyForResponse
func (t *TaskTargetForResponse_Company) FromCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCompanyForResponse performs a merge with any union data inside the TaskTargetForResponse_Company, using the provided CompanyForResponse
func (t *TaskTargetForResponse_Company) MergeCompanyForResponse(v CompanyForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t TaskTargetForResponse_Company) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *TaskTargetForResponse_Company) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOpportunityForResponse returns the union data inside the TaskTargetForResponse_Opportunity as a OpportunityForResponse
func (t TaskTargetForResponse_Opportunity) AsOpportunityForResponse() (OpportunityForResponse, error) {
	var body OpportunityForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpportunityForResponse overwrites any union data inside the TaskTargetForResponse_Opportunity as the provided OpportunityForResponse
func (t *TaskTargetForResponse_Opportunity) FromOpportunityForResponse(v OpportunityForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpportunityForResponse performs a merge with any union data inside the TaskTargetForResponse_Opportunity, using the provided OpportunityForResponse
func (t *TaskTargetForResponse_Opportunity) MergeOpportunityForResponse(v OpportunityForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t TaskTargetForResponse_Opportunity) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *TaskTargetForResponse_Opportunity) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPersonForResponse returns the union data inside the TaskTargetForResponse_Person as a PersonForResponse
func (t TaskTargetForResponse_Person) AsPersonForResponse() (PersonForResponse, error) {
	var body PersonForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPersonForResponse overwrites any union data inside the TaskTargetForResponse_Person as the provided PersonForResponse
func (t *TaskTargetForResponse_Person) FromPersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePersonForResponse performs a merge with any union data inside the TaskTargetForResponse_Person, using the provided PersonForResponse
func (t *TaskTargetForResponse_Person) MergePersonForResponse(v PersonForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t TaskTargetForResponse_Person) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *TaskTargetForResponse_Person) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsTaskForResponse returns the union data inside the TaskTargetForResponse_Task as a TaskForResponse
func (t TaskTargetForResponse_Task) AsTaskForResponse() (TaskForResponse, error) {
	var body TaskForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTaskForResponse overwrites any union data inside the TaskTargetForResponse_Task as the provided TaskForResponse
func (t *TaskTargetForResponse_Task) FromTaskForResponse(v TaskForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTaskForResponse performs a merge with any union data inside the TaskTargetForResponse_Task, using the provided TaskForResponse
func (t *TaskTargetForResponse_Task) MergeTaskForResponse(v TaskForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t TaskTargetForResponse_Task) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *TaskTargetForResponse_Task) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkspaceMemberForResponse returns the union data inside the TimelineActivityForResponse_WorkspaceMember as a WorkspaceMemberForResponse
func (t TimelineActivityForResponse_WorkspaceMember) AsWorkspaceMemberForResponse() (WorkspaceMemberForResponse, error) {
	var body WorkspaceMemberForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkspaceMemberForResponse overwrites any union data inside the TimelineActivityForResponse_WorkspaceMember as the provided WorkspaceMemberForResponse
func (t *TimelineActivityForResponse_WorkspaceMember) FromWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkspaceMemberForResponse performs a merge with any union data inside the TimelineActivityForResponse_WorkspaceMember, using the provided WorkspaceMemberForResponse
func (t *TimelineActivityForResponse_WorkspaceMember) MergeWorkspaceMemberForResponse(v WorkspaceMemberForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t TimelineActivityForResponse_WorkspaceMember) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *TimelineActivityForResponse_WorkspaceMember) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkflowForResponse returns the union data inside the WorkflowAutomatedTriggerForResponse_Workflow as a WorkflowForResponse
func (t WorkflowAutomatedTriggerForResponse_Workflow) AsWorkflowForResponse() (WorkflowForResponse, error) {
	var body WorkflowForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkflowForResponse overwrites any union data inside the WorkflowAutomatedTriggerForResponse_Workflow as the provided WorkflowForResponse
func (t *WorkflowAutomatedTriggerForResponse_Workflow) FromWorkflowForResponse(v WorkflowForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkflowForResponse performs a merge with any union data inside the WorkflowAutomatedTriggerForResponse_Workflow, using the provided WorkflowForResponse
func (t *WorkflowAutomatedTriggerForResponse_Workflow) MergeWorkflowForResponse(v WorkflowForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t WorkflowAutomatedTriggerForResponse_Workflow) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WorkflowAutomatedTriggerForResponse_Workflow) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkflowForResponse returns the union data inside the WorkflowRunForResponse_Workflow as a WorkflowForResponse
func (t WorkflowRunForResponse_Workflow) AsWorkflowForResponse() (WorkflowForResponse, error) {
	var body WorkflowForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkflowForResponse overwrites any union data inside the WorkflowRunForResponse_Workflow as the provided WorkflowForResponse
func (t *WorkflowRunForResponse_Workflow) FromWorkflowForResponse(v WorkflowForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkflowForResponse performs a merge with any union data inside the WorkflowRunForResponse_Workflow, using the provided WorkflowForResponse
func (t *WorkflowRunForResponse_Workflow) MergeWorkflowForResponse(v WorkflowForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t WorkflowRunForResponse_Workflow) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WorkflowRunForResponse_Workflow) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkflowVersionForResponse returns the union data inside the WorkflowRunForResponse_WorkflowVersion as a WorkflowVersionForResponse
func (t WorkflowRunForResponse_WorkflowVersion) AsWorkflowVersionForResponse() (WorkflowVersionForResponse, error) {
	var body WorkflowVersionForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkflowVersionForResponse overwrites any union data inside the WorkflowRunForResponse_WorkflowVersion as the provided WorkflowVersionForResponse
func (t *WorkflowRunForResponse_WorkflowVersion) FromWorkflowVersionForResponse(v WorkflowVersionForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkflowVersionForResponse performs a merge with any union data inside the WorkflowRunForResponse_WorkflowVersion, using the provided WorkflowVersionForResponse
func (t *WorkflowRunForResponse_WorkflowVersion) MergeWorkflowVersionForResponse(v WorkflowVersionForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t WorkflowRunForResponse_WorkflowVersion) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WorkflowRunForResponse_WorkflowVersion) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWorkflowForResponse returns the union data inside the WorkflowVersionForResponse_Workflow as a WorkflowForResponse
func (t WorkflowVersionForResponse_Workflow) AsWorkflowForResponse() (WorkflowForResponse, error) {
	var body WorkflowForResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWorkflowForResponse overwrites any union data inside the WorkflowVersionForResponse_Workflow as the provided WorkflowForResponse
func (t *WorkflowVersionForResponse_Workflow) FromWorkflowForResponse(v WorkflowForResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWorkflowForResponse performs a merge with any union data inside the WorkflowVersionForResponse_Workflow, using the provided WorkflowForResponse
func (t *WorkflowVersionForResponse_Workflow) MergeWorkflowForResponse(v WorkflowForResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t WorkflowVersionForResponse_Workflow) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WorkflowVersionForResponse_Workflow) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
