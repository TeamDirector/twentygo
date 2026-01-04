package main

import (
	"context"
	"io"
	"net/http"
)

// The interface specification for the client above.
type ClientInterface interface {
	// DeleteManyAttachments request
	DeleteManyAttachments(ctx context.Context, params *DeleteManyAttachmentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyAttachments request
	FindManyAttachments(ctx context.Context, params *FindManyAttachmentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyAttachmentsWithBody request with any body
	UpdateManyAttachmentsWithBody(ctx context.Context, params *UpdateManyAttachmentsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyAttachments(ctx context.Context, params *UpdateManyAttachmentsParams, body UpdateManyAttachmentsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneAttachmentWithBody request with any body
	CreateOneAttachmentWithBody(ctx context.Context, params *CreateOneAttachmentParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneAttachment(ctx context.Context, params *CreateOneAttachmentParams, body CreateOneAttachmentJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindAttachmentDuplicatesWithBody request with any body
	FindAttachmentDuplicatesWithBody(ctx context.Context, params *FindAttachmentDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindAttachmentDuplicates(ctx context.Context, params *FindAttachmentDuplicatesParams, body FindAttachmentDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyAttachmentsWithBody request with any body
	MergeManyAttachmentsWithBody(ctx context.Context, params *MergeManyAttachmentsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyAttachments(ctx context.Context, params *MergeManyAttachmentsParams, body MergeManyAttachmentsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneAttachment request
	DeleteOneAttachment(ctx context.Context, id IdPath, params *DeleteOneAttachmentParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneAttachment request
	FindOneAttachment(ctx context.Context, id IdPath, params *FindOneAttachmentParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneAttachmentWithBody request with any body
	UpdateOneAttachmentWithBody(ctx context.Context, id IdPath, params *UpdateOneAttachmentParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneAttachment(ctx context.Context, id IdPath, params *UpdateOneAttachmentParams, body UpdateOneAttachmentJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyAttachmentsWithBody request with any body
	CreateManyAttachmentsWithBody(ctx context.Context, params *CreateManyAttachmentsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyAttachments(ctx context.Context, params *CreateManyAttachmentsParams, body CreateManyAttachmentsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyBlocklistsWithBody request with any body
	CreateManyBlocklistsWithBody(ctx context.Context, params *CreateManyBlocklistsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyBlocklists(ctx context.Context, params *CreateManyBlocklistsParams, body CreateManyBlocklistsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyCalendarChannelEventAssociationsWithBody request with any body
	CreateManyCalendarChannelEventAssociationsWithBody(ctx context.Context, params *CreateManyCalendarChannelEventAssociationsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyCalendarChannelEventAssociations(ctx context.Context, params *CreateManyCalendarChannelEventAssociationsParams, body CreateManyCalendarChannelEventAssociationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyCalendarChannelsWithBody request with any body
	CreateManyCalendarChannelsWithBody(ctx context.Context, params *CreateManyCalendarChannelsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyCalendarChannels(ctx context.Context, params *CreateManyCalendarChannelsParams, body CreateManyCalendarChannelsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyCalendarEventParticipantsWithBody request with any body
	CreateManyCalendarEventParticipantsWithBody(ctx context.Context, params *CreateManyCalendarEventParticipantsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyCalendarEventParticipants(ctx context.Context, params *CreateManyCalendarEventParticipantsParams, body CreateManyCalendarEventParticipantsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyCalendarEventsWithBody request with any body
	CreateManyCalendarEventsWithBody(ctx context.Context, params *CreateManyCalendarEventsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyCalendarEvents(ctx context.Context, params *CreateManyCalendarEventsParams, body CreateManyCalendarEventsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyCompaniesWithBody request with any body
	CreateManyCompaniesWithBody(ctx context.Context, params *CreateManyCompaniesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyCompanies(ctx context.Context, params *CreateManyCompaniesParams, body CreateManyCompaniesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyConnectedAccountsWithBody request with any body
	CreateManyConnectedAccountsWithBody(ctx context.Context, params *CreateManyConnectedAccountsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyConnectedAccounts(ctx context.Context, params *CreateManyConnectedAccountsParams, body CreateManyConnectedAccountsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyDashboardsWithBody request with any body
	CreateManyDashboardsWithBody(ctx context.Context, params *CreateManyDashboardsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyDashboards(ctx context.Context, params *CreateManyDashboardsParams, body CreateManyDashboardsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyFavoriteFoldersWithBody request with any body
	CreateManyFavoriteFoldersWithBody(ctx context.Context, params *CreateManyFavoriteFoldersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyFavoriteFolders(ctx context.Context, params *CreateManyFavoriteFoldersParams, body CreateManyFavoriteFoldersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyFavoritesWithBody request with any body
	CreateManyFavoritesWithBody(ctx context.Context, params *CreateManyFavoritesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyFavorites(ctx context.Context, params *CreateManyFavoritesParams, body CreateManyFavoritesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyMessageChannelMessageAssociationsWithBody request with any body
	CreateManyMessageChannelMessageAssociationsWithBody(ctx context.Context, params *CreateManyMessageChannelMessageAssociationsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyMessageChannelMessageAssociations(ctx context.Context, params *CreateManyMessageChannelMessageAssociationsParams, body CreateManyMessageChannelMessageAssociationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyMessageChannelsWithBody request with any body
	CreateManyMessageChannelsWithBody(ctx context.Context, params *CreateManyMessageChannelsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyMessageChannels(ctx context.Context, params *CreateManyMessageChannelsParams, body CreateManyMessageChannelsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyMessageFoldersWithBody request with any body
	CreateManyMessageFoldersWithBody(ctx context.Context, params *CreateManyMessageFoldersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyMessageFolders(ctx context.Context, params *CreateManyMessageFoldersParams, body CreateManyMessageFoldersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyMessageParticipantsWithBody request with any body
	CreateManyMessageParticipantsWithBody(ctx context.Context, params *CreateManyMessageParticipantsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyMessageParticipants(ctx context.Context, params *CreateManyMessageParticipantsParams, body CreateManyMessageParticipantsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyMessageThreadsWithBody request with any body
	CreateManyMessageThreadsWithBody(ctx context.Context, params *CreateManyMessageThreadsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyMessageThreads(ctx context.Context, params *CreateManyMessageThreadsParams, body CreateManyMessageThreadsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyMessagesWithBody request with any body
	CreateManyMessagesWithBody(ctx context.Context, params *CreateManyMessagesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyMessages(ctx context.Context, params *CreateManyMessagesParams, body CreateManyMessagesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyNoteTargetsWithBody request with any body
	CreateManyNoteTargetsWithBody(ctx context.Context, params *CreateManyNoteTargetsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyNoteTargets(ctx context.Context, params *CreateManyNoteTargetsParams, body CreateManyNoteTargetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyNotesWithBody request with any body
	CreateManyNotesWithBody(ctx context.Context, params *CreateManyNotesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyNotes(ctx context.Context, params *CreateManyNotesParams, body CreateManyNotesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyOpportunitiesWithBody request with any body
	CreateManyOpportunitiesWithBody(ctx context.Context, params *CreateManyOpportunitiesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyOpportunities(ctx context.Context, params *CreateManyOpportunitiesParams, body CreateManyOpportunitiesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyPeopleWithBody request with any body
	CreateManyPeopleWithBody(ctx context.Context, params *CreateManyPeopleParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyPeople(ctx context.Context, params *CreateManyPeopleParams, body CreateManyPeopleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyTaskTargetsWithBody request with any body
	CreateManyTaskTargetsWithBody(ctx context.Context, params *CreateManyTaskTargetsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyTaskTargets(ctx context.Context, params *CreateManyTaskTargetsParams, body CreateManyTaskTargetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyTasksWithBody request with any body
	CreateManyTasksWithBody(ctx context.Context, params *CreateManyTasksParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyTasks(ctx context.Context, params *CreateManyTasksParams, body CreateManyTasksJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyTimelineActivitiesWithBody request with any body
	CreateManyTimelineActivitiesWithBody(ctx context.Context, params *CreateManyTimelineActivitiesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyTimelineActivities(ctx context.Context, params *CreateManyTimelineActivitiesParams, body CreateManyTimelineActivitiesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyWorkflowAutomatedTriggersWithBody request with any body
	CreateManyWorkflowAutomatedTriggersWithBody(ctx context.Context, params *CreateManyWorkflowAutomatedTriggersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyWorkflowAutomatedTriggers(ctx context.Context, params *CreateManyWorkflowAutomatedTriggersParams, body CreateManyWorkflowAutomatedTriggersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyWorkflowRunsWithBody request with any body
	CreateManyWorkflowRunsWithBody(ctx context.Context, params *CreateManyWorkflowRunsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyWorkflowRuns(ctx context.Context, params *CreateManyWorkflowRunsParams, body CreateManyWorkflowRunsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyWorkflowVersionsWithBody request with any body
	CreateManyWorkflowVersionsWithBody(ctx context.Context, params *CreateManyWorkflowVersionsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyWorkflowVersions(ctx context.Context, params *CreateManyWorkflowVersionsParams, body CreateManyWorkflowVersionsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyWorkflowsWithBody request with any body
	CreateManyWorkflowsWithBody(ctx context.Context, params *CreateManyWorkflowsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyWorkflows(ctx context.Context, params *CreateManyWorkflowsParams, body CreateManyWorkflowsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateManyWorkspaceMembersWithBody request with any body
	CreateManyWorkspaceMembersWithBody(ctx context.Context, params *CreateManyWorkspaceMembersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateManyWorkspaceMembers(ctx context.Context, params *CreateManyWorkspaceMembersParams, body CreateManyWorkspaceMembersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyBlocklists request
	DeleteManyBlocklists(ctx context.Context, params *DeleteManyBlocklistsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyBlocklists request
	FindManyBlocklists(ctx context.Context, params *FindManyBlocklistsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyBlocklistsWithBody request with any body
	UpdateManyBlocklistsWithBody(ctx context.Context, params *UpdateManyBlocklistsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyBlocklists(ctx context.Context, params *UpdateManyBlocklistsParams, body UpdateManyBlocklistsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneBlocklistWithBody request with any body
	CreateOneBlocklistWithBody(ctx context.Context, params *CreateOneBlocklistParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneBlocklist(ctx context.Context, params *CreateOneBlocklistParams, body CreateOneBlocklistJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindBlocklistDuplicatesWithBody request with any body
	FindBlocklistDuplicatesWithBody(ctx context.Context, params *FindBlocklistDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindBlocklistDuplicates(ctx context.Context, params *FindBlocklistDuplicatesParams, body FindBlocklistDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyBlocklistsWithBody request with any body
	MergeManyBlocklistsWithBody(ctx context.Context, params *MergeManyBlocklistsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyBlocklists(ctx context.Context, params *MergeManyBlocklistsParams, body MergeManyBlocklistsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneBlocklist request
	DeleteOneBlocklist(ctx context.Context, id IdPath, params *DeleteOneBlocklistParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneBlocklist request
	FindOneBlocklist(ctx context.Context, id IdPath, params *FindOneBlocklistParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneBlocklistWithBody request with any body
	UpdateOneBlocklistWithBody(ctx context.Context, id IdPath, params *UpdateOneBlocklistParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneBlocklist(ctx context.Context, id IdPath, params *UpdateOneBlocklistParams, body UpdateOneBlocklistJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyCalendarChannelEventAssociations request
	DeleteManyCalendarChannelEventAssociations(ctx context.Context, params *DeleteManyCalendarChannelEventAssociationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyCalendarChannelEventAssociations request
	FindManyCalendarChannelEventAssociations(ctx context.Context, params *FindManyCalendarChannelEventAssociationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyCalendarChannelEventAssociationsWithBody request with any body
	UpdateManyCalendarChannelEventAssociationsWithBody(ctx context.Context, params *UpdateManyCalendarChannelEventAssociationsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyCalendarChannelEventAssociations(ctx context.Context, params *UpdateManyCalendarChannelEventAssociationsParams, body UpdateManyCalendarChannelEventAssociationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneCalendarChannelEventAssociationWithBody request with any body
	CreateOneCalendarChannelEventAssociationWithBody(ctx context.Context, params *CreateOneCalendarChannelEventAssociationParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneCalendarChannelEventAssociation(ctx context.Context, params *CreateOneCalendarChannelEventAssociationParams, body CreateOneCalendarChannelEventAssociationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindCalendarChannelEventAssociationDuplicatesWithBody request with any body
	FindCalendarChannelEventAssociationDuplicatesWithBody(ctx context.Context, params *FindCalendarChannelEventAssociationDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindCalendarChannelEventAssociationDuplicates(ctx context.Context, params *FindCalendarChannelEventAssociationDuplicatesParams, body FindCalendarChannelEventAssociationDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyCalendarChannelEventAssociationsWithBody request with any body
	MergeManyCalendarChannelEventAssociationsWithBody(ctx context.Context, params *MergeManyCalendarChannelEventAssociationsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyCalendarChannelEventAssociations(ctx context.Context, params *MergeManyCalendarChannelEventAssociationsParams, body MergeManyCalendarChannelEventAssociationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneCalendarChannelEventAssociation request
	DeleteOneCalendarChannelEventAssociation(ctx context.Context, id IdPath, params *DeleteOneCalendarChannelEventAssociationParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneCalendarChannelEventAssociation request
	FindOneCalendarChannelEventAssociation(ctx context.Context, id IdPath, params *FindOneCalendarChannelEventAssociationParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneCalendarChannelEventAssociationWithBody request with any body
	UpdateOneCalendarChannelEventAssociationWithBody(ctx context.Context, id IdPath, params *UpdateOneCalendarChannelEventAssociationParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneCalendarChannelEventAssociation(ctx context.Context, id IdPath, params *UpdateOneCalendarChannelEventAssociationParams, body UpdateOneCalendarChannelEventAssociationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyCalendarChannels request
	DeleteManyCalendarChannels(ctx context.Context, params *DeleteManyCalendarChannelsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyCalendarChannels request
	FindManyCalendarChannels(ctx context.Context, params *FindManyCalendarChannelsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyCalendarChannelsWithBody request with any body
	UpdateManyCalendarChannelsWithBody(ctx context.Context, params *UpdateManyCalendarChannelsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyCalendarChannels(ctx context.Context, params *UpdateManyCalendarChannelsParams, body UpdateManyCalendarChannelsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneCalendarChannelWithBody request with any body
	CreateOneCalendarChannelWithBody(ctx context.Context, params *CreateOneCalendarChannelParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneCalendarChannel(ctx context.Context, params *CreateOneCalendarChannelParams, body CreateOneCalendarChannelJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindCalendarChannelDuplicatesWithBody request with any body
	FindCalendarChannelDuplicatesWithBody(ctx context.Context, params *FindCalendarChannelDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindCalendarChannelDuplicates(ctx context.Context, params *FindCalendarChannelDuplicatesParams, body FindCalendarChannelDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyCalendarChannelsWithBody request with any body
	MergeManyCalendarChannelsWithBody(ctx context.Context, params *MergeManyCalendarChannelsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyCalendarChannels(ctx context.Context, params *MergeManyCalendarChannelsParams, body MergeManyCalendarChannelsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneCalendarChannel request
	DeleteOneCalendarChannel(ctx context.Context, id IdPath, params *DeleteOneCalendarChannelParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneCalendarChannel request
	FindOneCalendarChannel(ctx context.Context, id IdPath, params *FindOneCalendarChannelParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneCalendarChannelWithBody request with any body
	UpdateOneCalendarChannelWithBody(ctx context.Context, id IdPath, params *UpdateOneCalendarChannelParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneCalendarChannel(ctx context.Context, id IdPath, params *UpdateOneCalendarChannelParams, body UpdateOneCalendarChannelJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyCalendarEventParticipants request
	DeleteManyCalendarEventParticipants(ctx context.Context, params *DeleteManyCalendarEventParticipantsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyCalendarEventParticipants request
	FindManyCalendarEventParticipants(ctx context.Context, params *FindManyCalendarEventParticipantsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyCalendarEventParticipantsWithBody request with any body
	UpdateManyCalendarEventParticipantsWithBody(ctx context.Context, params *UpdateManyCalendarEventParticipantsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyCalendarEventParticipants(ctx context.Context, params *UpdateManyCalendarEventParticipantsParams, body UpdateManyCalendarEventParticipantsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneCalendarEventParticipantWithBody request with any body
	CreateOneCalendarEventParticipantWithBody(ctx context.Context, params *CreateOneCalendarEventParticipantParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneCalendarEventParticipant(ctx context.Context, params *CreateOneCalendarEventParticipantParams, body CreateOneCalendarEventParticipantJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindCalendarEventParticipantDuplicatesWithBody request with any body
	FindCalendarEventParticipantDuplicatesWithBody(ctx context.Context, params *FindCalendarEventParticipantDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindCalendarEventParticipantDuplicates(ctx context.Context, params *FindCalendarEventParticipantDuplicatesParams, body FindCalendarEventParticipantDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyCalendarEventParticipantsWithBody request with any body
	MergeManyCalendarEventParticipantsWithBody(ctx context.Context, params *MergeManyCalendarEventParticipantsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyCalendarEventParticipants(ctx context.Context, params *MergeManyCalendarEventParticipantsParams, body MergeManyCalendarEventParticipantsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneCalendarEventParticipant request
	DeleteOneCalendarEventParticipant(ctx context.Context, id IdPath, params *DeleteOneCalendarEventParticipantParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneCalendarEventParticipant request
	FindOneCalendarEventParticipant(ctx context.Context, id IdPath, params *FindOneCalendarEventParticipantParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneCalendarEventParticipantWithBody request with any body
	UpdateOneCalendarEventParticipantWithBody(ctx context.Context, id IdPath, params *UpdateOneCalendarEventParticipantParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneCalendarEventParticipant(ctx context.Context, id IdPath, params *UpdateOneCalendarEventParticipantParams, body UpdateOneCalendarEventParticipantJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyCalendarEvents request
	DeleteManyCalendarEvents(ctx context.Context, params *DeleteManyCalendarEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyCalendarEvents request
	FindManyCalendarEvents(ctx context.Context, params *FindManyCalendarEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyCalendarEventsWithBody request with any body
	UpdateManyCalendarEventsWithBody(ctx context.Context, params *UpdateManyCalendarEventsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyCalendarEvents(ctx context.Context, params *UpdateManyCalendarEventsParams, body UpdateManyCalendarEventsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneCalendarEventWithBody request with any body
	CreateOneCalendarEventWithBody(ctx context.Context, params *CreateOneCalendarEventParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneCalendarEvent(ctx context.Context, params *CreateOneCalendarEventParams, body CreateOneCalendarEventJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindCalendarEventDuplicatesWithBody request with any body
	FindCalendarEventDuplicatesWithBody(ctx context.Context, params *FindCalendarEventDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindCalendarEventDuplicates(ctx context.Context, params *FindCalendarEventDuplicatesParams, body FindCalendarEventDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyCalendarEventsWithBody request with any body
	MergeManyCalendarEventsWithBody(ctx context.Context, params *MergeManyCalendarEventsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyCalendarEvents(ctx context.Context, params *MergeManyCalendarEventsParams, body MergeManyCalendarEventsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneCalendarEvent request
	DeleteOneCalendarEvent(ctx context.Context, id IdPath, params *DeleteOneCalendarEventParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneCalendarEvent request
	FindOneCalendarEvent(ctx context.Context, id IdPath, params *FindOneCalendarEventParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneCalendarEventWithBody request with any body
	UpdateOneCalendarEventWithBody(ctx context.Context, id IdPath, params *UpdateOneCalendarEventParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneCalendarEvent(ctx context.Context, id IdPath, params *UpdateOneCalendarEventParams, body UpdateOneCalendarEventJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyCompanies request
	DeleteManyCompanies(ctx context.Context, params *DeleteManyCompaniesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyCompanies request
	FindManyCompanies(ctx context.Context, params *FindManyCompaniesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyCompaniesWithBody request with any body
	UpdateManyCompaniesWithBody(ctx context.Context, params *UpdateManyCompaniesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyCompanies(ctx context.Context, params *UpdateManyCompaniesParams, body UpdateManyCompaniesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneCompanyWithBody request with any body
	CreateOneCompanyWithBody(ctx context.Context, params *CreateOneCompanyParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneCompany(ctx context.Context, params *CreateOneCompanyParams, body CreateOneCompanyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindCompanyDuplicatesWithBody request with any body
	FindCompanyDuplicatesWithBody(ctx context.Context, params *FindCompanyDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindCompanyDuplicates(ctx context.Context, params *FindCompanyDuplicatesParams, body FindCompanyDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyCompaniesWithBody request with any body
	MergeManyCompaniesWithBody(ctx context.Context, params *MergeManyCompaniesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyCompanies(ctx context.Context, params *MergeManyCompaniesParams, body MergeManyCompaniesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneCompany request
	DeleteOneCompany(ctx context.Context, id IdPath, params *DeleteOneCompanyParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneCompany request
	FindOneCompany(ctx context.Context, id IdPath, params *FindOneCompanyParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneCompanyWithBody request with any body
	UpdateOneCompanyWithBody(ctx context.Context, id IdPath, params *UpdateOneCompanyParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneCompany(ctx context.Context, id IdPath, params *UpdateOneCompanyParams, body UpdateOneCompanyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyConnectedAccounts request
	DeleteManyConnectedAccounts(ctx context.Context, params *DeleteManyConnectedAccountsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyConnectedAccounts request
	FindManyConnectedAccounts(ctx context.Context, params *FindManyConnectedAccountsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyConnectedAccountsWithBody request with any body
	UpdateManyConnectedAccountsWithBody(ctx context.Context, params *UpdateManyConnectedAccountsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyConnectedAccounts(ctx context.Context, params *UpdateManyConnectedAccountsParams, body UpdateManyConnectedAccountsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneConnectedAccountWithBody request with any body
	CreateOneConnectedAccountWithBody(ctx context.Context, params *CreateOneConnectedAccountParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneConnectedAccount(ctx context.Context, params *CreateOneConnectedAccountParams, body CreateOneConnectedAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindConnectedAccountDuplicatesWithBody request with any body
	FindConnectedAccountDuplicatesWithBody(ctx context.Context, params *FindConnectedAccountDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindConnectedAccountDuplicates(ctx context.Context, params *FindConnectedAccountDuplicatesParams, body FindConnectedAccountDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyConnectedAccountsWithBody request with any body
	MergeManyConnectedAccountsWithBody(ctx context.Context, params *MergeManyConnectedAccountsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyConnectedAccounts(ctx context.Context, params *MergeManyConnectedAccountsParams, body MergeManyConnectedAccountsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneConnectedAccount request
	DeleteOneConnectedAccount(ctx context.Context, id IdPath, params *DeleteOneConnectedAccountParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneConnectedAccount request
	FindOneConnectedAccount(ctx context.Context, id IdPath, params *FindOneConnectedAccountParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneConnectedAccountWithBody request with any body
	UpdateOneConnectedAccountWithBody(ctx context.Context, id IdPath, params *UpdateOneConnectedAccountParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneConnectedAccount(ctx context.Context, id IdPath, params *UpdateOneConnectedAccountParams, body UpdateOneConnectedAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyDashboards request
	DeleteManyDashboards(ctx context.Context, params *DeleteManyDashboardsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyDashboards request
	FindManyDashboards(ctx context.Context, params *FindManyDashboardsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyDashboardsWithBody request with any body
	UpdateManyDashboardsWithBody(ctx context.Context, params *UpdateManyDashboardsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyDashboards(ctx context.Context, params *UpdateManyDashboardsParams, body UpdateManyDashboardsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneDashboardWithBody request with any body
	CreateOneDashboardWithBody(ctx context.Context, params *CreateOneDashboardParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneDashboard(ctx context.Context, params *CreateOneDashboardParams, body CreateOneDashboardJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindDashboardDuplicatesWithBody request with any body
	FindDashboardDuplicatesWithBody(ctx context.Context, params *FindDashboardDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindDashboardDuplicates(ctx context.Context, params *FindDashboardDuplicatesParams, body FindDashboardDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyDashboardsWithBody request with any body
	MergeManyDashboardsWithBody(ctx context.Context, params *MergeManyDashboardsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyDashboards(ctx context.Context, params *MergeManyDashboardsParams, body MergeManyDashboardsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneDashboard request
	DeleteOneDashboard(ctx context.Context, id IdPath, params *DeleteOneDashboardParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneDashboard request
	FindOneDashboard(ctx context.Context, id IdPath, params *FindOneDashboardParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneDashboardWithBody request with any body
	UpdateOneDashboardWithBody(ctx context.Context, id IdPath, params *UpdateOneDashboardParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneDashboard(ctx context.Context, id IdPath, params *UpdateOneDashboardParams, body UpdateOneDashboardJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DuplicateDashboard request
	DuplicateDashboard(ctx context.Context, id IdPath, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyFavoriteFolders request
	DeleteManyFavoriteFolders(ctx context.Context, params *DeleteManyFavoriteFoldersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyFavoriteFolders request
	FindManyFavoriteFolders(ctx context.Context, params *FindManyFavoriteFoldersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyFavoriteFoldersWithBody request with any body
	UpdateManyFavoriteFoldersWithBody(ctx context.Context, params *UpdateManyFavoriteFoldersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyFavoriteFolders(ctx context.Context, params *UpdateManyFavoriteFoldersParams, body UpdateManyFavoriteFoldersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneFavoriteFolderWithBody request with any body
	CreateOneFavoriteFolderWithBody(ctx context.Context, params *CreateOneFavoriteFolderParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneFavoriteFolder(ctx context.Context, params *CreateOneFavoriteFolderParams, body CreateOneFavoriteFolderJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindFavoriteFolderDuplicatesWithBody request with any body
	FindFavoriteFolderDuplicatesWithBody(ctx context.Context, params *FindFavoriteFolderDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindFavoriteFolderDuplicates(ctx context.Context, params *FindFavoriteFolderDuplicatesParams, body FindFavoriteFolderDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyFavoriteFoldersWithBody request with any body
	MergeManyFavoriteFoldersWithBody(ctx context.Context, params *MergeManyFavoriteFoldersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyFavoriteFolders(ctx context.Context, params *MergeManyFavoriteFoldersParams, body MergeManyFavoriteFoldersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneFavoriteFolder request
	DeleteOneFavoriteFolder(ctx context.Context, id IdPath, params *DeleteOneFavoriteFolderParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneFavoriteFolder request
	FindOneFavoriteFolder(ctx context.Context, id IdPath, params *FindOneFavoriteFolderParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneFavoriteFolderWithBody request with any body
	UpdateOneFavoriteFolderWithBody(ctx context.Context, id IdPath, params *UpdateOneFavoriteFolderParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneFavoriteFolder(ctx context.Context, id IdPath, params *UpdateOneFavoriteFolderParams, body UpdateOneFavoriteFolderJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyFavorites request
	DeleteManyFavorites(ctx context.Context, params *DeleteManyFavoritesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyFavorites request
	FindManyFavorites(ctx context.Context, params *FindManyFavoritesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyFavoritesWithBody request with any body
	UpdateManyFavoritesWithBody(ctx context.Context, params *UpdateManyFavoritesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyFavorites(ctx context.Context, params *UpdateManyFavoritesParams, body UpdateManyFavoritesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneFavoriteWithBody request with any body
	CreateOneFavoriteWithBody(ctx context.Context, params *CreateOneFavoriteParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneFavorite(ctx context.Context, params *CreateOneFavoriteParams, body CreateOneFavoriteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindFavoriteDuplicatesWithBody request with any body
	FindFavoriteDuplicatesWithBody(ctx context.Context, params *FindFavoriteDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindFavoriteDuplicates(ctx context.Context, params *FindFavoriteDuplicatesParams, body FindFavoriteDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyFavoritesWithBody request with any body
	MergeManyFavoritesWithBody(ctx context.Context, params *MergeManyFavoritesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyFavorites(ctx context.Context, params *MergeManyFavoritesParams, body MergeManyFavoritesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneFavorite request
	DeleteOneFavorite(ctx context.Context, id IdPath, params *DeleteOneFavoriteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneFavorite request
	FindOneFavorite(ctx context.Context, id IdPath, params *FindOneFavoriteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneFavoriteWithBody request with any body
	UpdateOneFavoriteWithBody(ctx context.Context, id IdPath, params *UpdateOneFavoriteParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneFavorite(ctx context.Context, id IdPath, params *UpdateOneFavoriteParams, body UpdateOneFavoriteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyMessageChannelMessageAssociations request
	DeleteManyMessageChannelMessageAssociations(ctx context.Context, params *DeleteManyMessageChannelMessageAssociationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyMessageChannelMessageAssociations request
	FindManyMessageChannelMessageAssociations(ctx context.Context, params *FindManyMessageChannelMessageAssociationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyMessageChannelMessageAssociationsWithBody request with any body
	UpdateManyMessageChannelMessageAssociationsWithBody(ctx context.Context, params *UpdateManyMessageChannelMessageAssociationsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyMessageChannelMessageAssociations(ctx context.Context, params *UpdateManyMessageChannelMessageAssociationsParams, body UpdateManyMessageChannelMessageAssociationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneMessageChannelMessageAssociationWithBody request with any body
	CreateOneMessageChannelMessageAssociationWithBody(ctx context.Context, params *CreateOneMessageChannelMessageAssociationParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneMessageChannelMessageAssociation(ctx context.Context, params *CreateOneMessageChannelMessageAssociationParams, body CreateOneMessageChannelMessageAssociationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindMessageChannelMessageAssociationDuplicatesWithBody request with any body
	FindMessageChannelMessageAssociationDuplicatesWithBody(ctx context.Context, params *FindMessageChannelMessageAssociationDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindMessageChannelMessageAssociationDuplicates(ctx context.Context, params *FindMessageChannelMessageAssociationDuplicatesParams, body FindMessageChannelMessageAssociationDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyMessageChannelMessageAssociationsWithBody request with any body
	MergeManyMessageChannelMessageAssociationsWithBody(ctx context.Context, params *MergeManyMessageChannelMessageAssociationsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyMessageChannelMessageAssociations(ctx context.Context, params *MergeManyMessageChannelMessageAssociationsParams, body MergeManyMessageChannelMessageAssociationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneMessageChannelMessageAssociation request
	DeleteOneMessageChannelMessageAssociation(ctx context.Context, id IdPath, params *DeleteOneMessageChannelMessageAssociationParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneMessageChannelMessageAssociation request
	FindOneMessageChannelMessageAssociation(ctx context.Context, id IdPath, params *FindOneMessageChannelMessageAssociationParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneMessageChannelMessageAssociationWithBody request with any body
	UpdateOneMessageChannelMessageAssociationWithBody(ctx context.Context, id IdPath, params *UpdateOneMessageChannelMessageAssociationParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneMessageChannelMessageAssociation(ctx context.Context, id IdPath, params *UpdateOneMessageChannelMessageAssociationParams, body UpdateOneMessageChannelMessageAssociationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyMessageChannels request
	DeleteManyMessageChannels(ctx context.Context, params *DeleteManyMessageChannelsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyMessageChannels request
	FindManyMessageChannels(ctx context.Context, params *FindManyMessageChannelsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyMessageChannelsWithBody request with any body
	UpdateManyMessageChannelsWithBody(ctx context.Context, params *UpdateManyMessageChannelsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyMessageChannels(ctx context.Context, params *UpdateManyMessageChannelsParams, body UpdateManyMessageChannelsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneMessageChannelWithBody request with any body
	CreateOneMessageChannelWithBody(ctx context.Context, params *CreateOneMessageChannelParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneMessageChannel(ctx context.Context, params *CreateOneMessageChannelParams, body CreateOneMessageChannelJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindMessageChannelDuplicatesWithBody request with any body
	FindMessageChannelDuplicatesWithBody(ctx context.Context, params *FindMessageChannelDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindMessageChannelDuplicates(ctx context.Context, params *FindMessageChannelDuplicatesParams, body FindMessageChannelDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyMessageChannelsWithBody request with any body
	MergeManyMessageChannelsWithBody(ctx context.Context, params *MergeManyMessageChannelsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyMessageChannels(ctx context.Context, params *MergeManyMessageChannelsParams, body MergeManyMessageChannelsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneMessageChannel request
	DeleteOneMessageChannel(ctx context.Context, id IdPath, params *DeleteOneMessageChannelParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneMessageChannel request
	FindOneMessageChannel(ctx context.Context, id IdPath, params *FindOneMessageChannelParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneMessageChannelWithBody request with any body
	UpdateOneMessageChannelWithBody(ctx context.Context, id IdPath, params *UpdateOneMessageChannelParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneMessageChannel(ctx context.Context, id IdPath, params *UpdateOneMessageChannelParams, body UpdateOneMessageChannelJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyMessageFolders request
	DeleteManyMessageFolders(ctx context.Context, params *DeleteManyMessageFoldersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyMessageFolders request
	FindManyMessageFolders(ctx context.Context, params *FindManyMessageFoldersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyMessageFoldersWithBody request with any body
	UpdateManyMessageFoldersWithBody(ctx context.Context, params *UpdateManyMessageFoldersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyMessageFolders(ctx context.Context, params *UpdateManyMessageFoldersParams, body UpdateManyMessageFoldersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneMessageFolderWithBody request with any body
	CreateOneMessageFolderWithBody(ctx context.Context, params *CreateOneMessageFolderParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneMessageFolder(ctx context.Context, params *CreateOneMessageFolderParams, body CreateOneMessageFolderJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindMessageFolderDuplicatesWithBody request with any body
	FindMessageFolderDuplicatesWithBody(ctx context.Context, params *FindMessageFolderDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindMessageFolderDuplicates(ctx context.Context, params *FindMessageFolderDuplicatesParams, body FindMessageFolderDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyMessageFoldersWithBody request with any body
	MergeManyMessageFoldersWithBody(ctx context.Context, params *MergeManyMessageFoldersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyMessageFolders(ctx context.Context, params *MergeManyMessageFoldersParams, body MergeManyMessageFoldersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneMessageFolder request
	DeleteOneMessageFolder(ctx context.Context, id IdPath, params *DeleteOneMessageFolderParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneMessageFolder request
	FindOneMessageFolder(ctx context.Context, id IdPath, params *FindOneMessageFolderParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneMessageFolderWithBody request with any body
	UpdateOneMessageFolderWithBody(ctx context.Context, id IdPath, params *UpdateOneMessageFolderParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneMessageFolder(ctx context.Context, id IdPath, params *UpdateOneMessageFolderParams, body UpdateOneMessageFolderJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyMessageParticipants request
	DeleteManyMessageParticipants(ctx context.Context, params *DeleteManyMessageParticipantsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyMessageParticipants request
	FindManyMessageParticipants(ctx context.Context, params *FindManyMessageParticipantsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyMessageParticipantsWithBody request with any body
	UpdateManyMessageParticipantsWithBody(ctx context.Context, params *UpdateManyMessageParticipantsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyMessageParticipants(ctx context.Context, params *UpdateManyMessageParticipantsParams, body UpdateManyMessageParticipantsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneMessageParticipantWithBody request with any body
	CreateOneMessageParticipantWithBody(ctx context.Context, params *CreateOneMessageParticipantParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneMessageParticipant(ctx context.Context, params *CreateOneMessageParticipantParams, body CreateOneMessageParticipantJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindMessageParticipantDuplicatesWithBody request with any body
	FindMessageParticipantDuplicatesWithBody(ctx context.Context, params *FindMessageParticipantDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindMessageParticipantDuplicates(ctx context.Context, params *FindMessageParticipantDuplicatesParams, body FindMessageParticipantDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyMessageParticipantsWithBody request with any body
	MergeManyMessageParticipantsWithBody(ctx context.Context, params *MergeManyMessageParticipantsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyMessageParticipants(ctx context.Context, params *MergeManyMessageParticipantsParams, body MergeManyMessageParticipantsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneMessageParticipant request
	DeleteOneMessageParticipant(ctx context.Context, id IdPath, params *DeleteOneMessageParticipantParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneMessageParticipant request
	FindOneMessageParticipant(ctx context.Context, id IdPath, params *FindOneMessageParticipantParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneMessageParticipantWithBody request with any body
	UpdateOneMessageParticipantWithBody(ctx context.Context, id IdPath, params *UpdateOneMessageParticipantParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneMessageParticipant(ctx context.Context, id IdPath, params *UpdateOneMessageParticipantParams, body UpdateOneMessageParticipantJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyMessageThreads request
	DeleteManyMessageThreads(ctx context.Context, params *DeleteManyMessageThreadsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyMessageThreads request
	FindManyMessageThreads(ctx context.Context, params *FindManyMessageThreadsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyMessageThreadsWithBody request with any body
	UpdateManyMessageThreadsWithBody(ctx context.Context, params *UpdateManyMessageThreadsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyMessageThreads(ctx context.Context, params *UpdateManyMessageThreadsParams, body UpdateManyMessageThreadsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneMessageThreadWithBody request with any body
	CreateOneMessageThreadWithBody(ctx context.Context, params *CreateOneMessageThreadParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneMessageThread(ctx context.Context, params *CreateOneMessageThreadParams, body CreateOneMessageThreadJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindMessageThreadDuplicatesWithBody request with any body
	FindMessageThreadDuplicatesWithBody(ctx context.Context, params *FindMessageThreadDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindMessageThreadDuplicates(ctx context.Context, params *FindMessageThreadDuplicatesParams, body FindMessageThreadDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyMessageThreadsWithBody request with any body
	MergeManyMessageThreadsWithBody(ctx context.Context, params *MergeManyMessageThreadsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyMessageThreads(ctx context.Context, params *MergeManyMessageThreadsParams, body MergeManyMessageThreadsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneMessageThread request
	DeleteOneMessageThread(ctx context.Context, id IdPath, params *DeleteOneMessageThreadParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneMessageThread request
	FindOneMessageThread(ctx context.Context, id IdPath, params *FindOneMessageThreadParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneMessageThreadWithBody request with any body
	UpdateOneMessageThreadWithBody(ctx context.Context, id IdPath, params *UpdateOneMessageThreadParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneMessageThread(ctx context.Context, id IdPath, params *UpdateOneMessageThreadParams, body UpdateOneMessageThreadJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyMessages request
	DeleteManyMessages(ctx context.Context, params *DeleteManyMessagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyMessages request
	FindManyMessages(ctx context.Context, params *FindManyMessagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyMessagesWithBody request with any body
	UpdateManyMessagesWithBody(ctx context.Context, params *UpdateManyMessagesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyMessages(ctx context.Context, params *UpdateManyMessagesParams, body UpdateManyMessagesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneMessageWithBody request with any body
	CreateOneMessageWithBody(ctx context.Context, params *CreateOneMessageParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneMessage(ctx context.Context, params *CreateOneMessageParams, body CreateOneMessageJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindMessageDuplicatesWithBody request with any body
	FindMessageDuplicatesWithBody(ctx context.Context, params *FindMessageDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindMessageDuplicates(ctx context.Context, params *FindMessageDuplicatesParams, body FindMessageDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyMessagesWithBody request with any body
	MergeManyMessagesWithBody(ctx context.Context, params *MergeManyMessagesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyMessages(ctx context.Context, params *MergeManyMessagesParams, body MergeManyMessagesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneMessage request
	DeleteOneMessage(ctx context.Context, id IdPath, params *DeleteOneMessageParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneMessage request
	FindOneMessage(ctx context.Context, id IdPath, params *FindOneMessageParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneMessageWithBody request with any body
	UpdateOneMessageWithBody(ctx context.Context, id IdPath, params *UpdateOneMessageParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneMessage(ctx context.Context, id IdPath, params *UpdateOneMessageParams, body UpdateOneMessageJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyNoteTargets request
	DeleteManyNoteTargets(ctx context.Context, params *DeleteManyNoteTargetsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyNoteTargets request
	FindManyNoteTargets(ctx context.Context, params *FindManyNoteTargetsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyNoteTargetsWithBody request with any body
	UpdateManyNoteTargetsWithBody(ctx context.Context, params *UpdateManyNoteTargetsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyNoteTargets(ctx context.Context, params *UpdateManyNoteTargetsParams, body UpdateManyNoteTargetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneNoteTargetWithBody request with any body
	CreateOneNoteTargetWithBody(ctx context.Context, params *CreateOneNoteTargetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneNoteTarget(ctx context.Context, params *CreateOneNoteTargetParams, body CreateOneNoteTargetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindNoteTargetDuplicatesWithBody request with any body
	FindNoteTargetDuplicatesWithBody(ctx context.Context, params *FindNoteTargetDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindNoteTargetDuplicates(ctx context.Context, params *FindNoteTargetDuplicatesParams, body FindNoteTargetDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyNoteTargetsWithBody request with any body
	MergeManyNoteTargetsWithBody(ctx context.Context, params *MergeManyNoteTargetsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyNoteTargets(ctx context.Context, params *MergeManyNoteTargetsParams, body MergeManyNoteTargetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneNoteTarget request
	DeleteOneNoteTarget(ctx context.Context, id IdPath, params *DeleteOneNoteTargetParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneNoteTarget request
	FindOneNoteTarget(ctx context.Context, id IdPath, params *FindOneNoteTargetParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneNoteTargetWithBody request with any body
	UpdateOneNoteTargetWithBody(ctx context.Context, id IdPath, params *UpdateOneNoteTargetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneNoteTarget(ctx context.Context, id IdPath, params *UpdateOneNoteTargetParams, body UpdateOneNoteTargetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyNotes request
	DeleteManyNotes(ctx context.Context, params *DeleteManyNotesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyNotes request
	FindManyNotes(ctx context.Context, params *FindManyNotesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyNotesWithBody request with any body
	UpdateManyNotesWithBody(ctx context.Context, params *UpdateManyNotesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyNotes(ctx context.Context, params *UpdateManyNotesParams, body UpdateManyNotesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneNoteWithBody request with any body
	CreateOneNoteWithBody(ctx context.Context, params *CreateOneNoteParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneNote(ctx context.Context, params *CreateOneNoteParams, body CreateOneNoteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindNoteDuplicatesWithBody request with any body
	FindNoteDuplicatesWithBody(ctx context.Context, params *FindNoteDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindNoteDuplicates(ctx context.Context, params *FindNoteDuplicatesParams, body FindNoteDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyNotesWithBody request with any body
	MergeManyNotesWithBody(ctx context.Context, params *MergeManyNotesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyNotes(ctx context.Context, params *MergeManyNotesParams, body MergeManyNotesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneNote request
	DeleteOneNote(ctx context.Context, id IdPath, params *DeleteOneNoteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneNote request
	FindOneNote(ctx context.Context, id IdPath, params *FindOneNoteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneNoteWithBody request with any body
	UpdateOneNoteWithBody(ctx context.Context, id IdPath, params *UpdateOneNoteParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneNote(ctx context.Context, id IdPath, params *UpdateOneNoteParams, body UpdateOneNoteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetOpenApiSchema request
	GetOpenApiSchema(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyOpportunities request
	DeleteManyOpportunities(ctx context.Context, params *DeleteManyOpportunitiesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyOpportunities request
	FindManyOpportunities(ctx context.Context, params *FindManyOpportunitiesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyOpportunitiesWithBody request with any body
	UpdateManyOpportunitiesWithBody(ctx context.Context, params *UpdateManyOpportunitiesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyOpportunities(ctx context.Context, params *UpdateManyOpportunitiesParams, body UpdateManyOpportunitiesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneOpportunityWithBody request with any body
	CreateOneOpportunityWithBody(ctx context.Context, params *CreateOneOpportunityParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneOpportunity(ctx context.Context, params *CreateOneOpportunityParams, body CreateOneOpportunityJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOpportunityDuplicatesWithBody request with any body
	FindOpportunityDuplicatesWithBody(ctx context.Context, params *FindOpportunityDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindOpportunityDuplicates(ctx context.Context, params *FindOpportunityDuplicatesParams, body FindOpportunityDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyOpportunitiesWithBody request with any body
	MergeManyOpportunitiesWithBody(ctx context.Context, params *MergeManyOpportunitiesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyOpportunities(ctx context.Context, params *MergeManyOpportunitiesParams, body MergeManyOpportunitiesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneOpportunity request
	DeleteOneOpportunity(ctx context.Context, id IdPath, params *DeleteOneOpportunityParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneOpportunity request
	FindOneOpportunity(ctx context.Context, id IdPath, params *FindOneOpportunityParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneOpportunityWithBody request with any body
	UpdateOneOpportunityWithBody(ctx context.Context, id IdPath, params *UpdateOneOpportunityParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneOpportunity(ctx context.Context, id IdPath, params *UpdateOneOpportunityParams, body UpdateOneOpportunityJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyPeople request
	DeleteManyPeople(ctx context.Context, params *DeleteManyPeopleParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyPeople request
	FindManyPeople(ctx context.Context, params *FindManyPeopleParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyPeopleWithBody request with any body
	UpdateManyPeopleWithBody(ctx context.Context, params *UpdateManyPeopleParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyPeople(ctx context.Context, params *UpdateManyPeopleParams, body UpdateManyPeopleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOnePersonWithBody request with any body
	CreateOnePersonWithBody(ctx context.Context, params *CreateOnePersonParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOnePerson(ctx context.Context, params *CreateOnePersonParams, body CreateOnePersonJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindPersonDuplicatesWithBody request with any body
	FindPersonDuplicatesWithBody(ctx context.Context, params *FindPersonDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindPersonDuplicates(ctx context.Context, params *FindPersonDuplicatesParams, body FindPersonDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyPeopleWithBody request with any body
	MergeManyPeopleWithBody(ctx context.Context, params *MergeManyPeopleParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyPeople(ctx context.Context, params *MergeManyPeopleParams, body MergeManyPeopleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOnePerson request
	DeleteOnePerson(ctx context.Context, id IdPath, params *DeleteOnePersonParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOnePerson request
	FindOnePerson(ctx context.Context, id IdPath, params *FindOnePersonParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOnePersonWithBody request with any body
	UpdateOnePersonWithBody(ctx context.Context, id IdPath, params *UpdateOnePersonParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOnePerson(ctx context.Context, id IdPath, params *UpdateOnePersonParams, body UpdateOnePersonJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyAttachments request
	RestoreManyAttachments(ctx context.Context, params *RestoreManyAttachmentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneAttachment request
	RestoreOneAttachment(ctx context.Context, id IdPath, params *RestoreOneAttachmentParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyBlocklists request
	RestoreManyBlocklists(ctx context.Context, params *RestoreManyBlocklistsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneBlocklist request
	RestoreOneBlocklist(ctx context.Context, id IdPath, params *RestoreOneBlocklistParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyCalendarChannelEventAssociations request
	RestoreManyCalendarChannelEventAssociations(ctx context.Context, params *RestoreManyCalendarChannelEventAssociationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneCalendarChannelEventAssociation request
	RestoreOneCalendarChannelEventAssociation(ctx context.Context, id IdPath, params *RestoreOneCalendarChannelEventAssociationParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyCalendarChannels request
	RestoreManyCalendarChannels(ctx context.Context, params *RestoreManyCalendarChannelsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneCalendarChannel request
	RestoreOneCalendarChannel(ctx context.Context, id IdPath, params *RestoreOneCalendarChannelParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyCalendarEventParticipants request
	RestoreManyCalendarEventParticipants(ctx context.Context, params *RestoreManyCalendarEventParticipantsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneCalendarEventParticipant request
	RestoreOneCalendarEventParticipant(ctx context.Context, id IdPath, params *RestoreOneCalendarEventParticipantParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyCalendarEvents request
	RestoreManyCalendarEvents(ctx context.Context, params *RestoreManyCalendarEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneCalendarEvent request
	RestoreOneCalendarEvent(ctx context.Context, id IdPath, params *RestoreOneCalendarEventParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyCompanies request
	RestoreManyCompanies(ctx context.Context, params *RestoreManyCompaniesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneCompany request
	RestoreOneCompany(ctx context.Context, id IdPath, params *RestoreOneCompanyParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyConnectedAccounts request
	RestoreManyConnectedAccounts(ctx context.Context, params *RestoreManyConnectedAccountsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneConnectedAccount request
	RestoreOneConnectedAccount(ctx context.Context, id IdPath, params *RestoreOneConnectedAccountParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyDashboards request
	RestoreManyDashboards(ctx context.Context, params *RestoreManyDashboardsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneDashboard request
	RestoreOneDashboard(ctx context.Context, id IdPath, params *RestoreOneDashboardParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyFavoriteFolders request
	RestoreManyFavoriteFolders(ctx context.Context, params *RestoreManyFavoriteFoldersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneFavoriteFolder request
	RestoreOneFavoriteFolder(ctx context.Context, id IdPath, params *RestoreOneFavoriteFolderParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyFavorites request
	RestoreManyFavorites(ctx context.Context, params *RestoreManyFavoritesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneFavorite request
	RestoreOneFavorite(ctx context.Context, id IdPath, params *RestoreOneFavoriteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyMessageChannelMessageAssociations request
	RestoreManyMessageChannelMessageAssociations(ctx context.Context, params *RestoreManyMessageChannelMessageAssociationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneMessageChannelMessageAssociation request
	RestoreOneMessageChannelMessageAssociation(ctx context.Context, id IdPath, params *RestoreOneMessageChannelMessageAssociationParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyMessageChannels request
	RestoreManyMessageChannels(ctx context.Context, params *RestoreManyMessageChannelsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneMessageChannel request
	RestoreOneMessageChannel(ctx context.Context, id IdPath, params *RestoreOneMessageChannelParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyMessageFolders request
	RestoreManyMessageFolders(ctx context.Context, params *RestoreManyMessageFoldersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneMessageFolder request
	RestoreOneMessageFolder(ctx context.Context, id IdPath, params *RestoreOneMessageFolderParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyMessageParticipants request
	RestoreManyMessageParticipants(ctx context.Context, params *RestoreManyMessageParticipantsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneMessageParticipant request
	RestoreOneMessageParticipant(ctx context.Context, id IdPath, params *RestoreOneMessageParticipantParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyMessageThreads request
	RestoreManyMessageThreads(ctx context.Context, params *RestoreManyMessageThreadsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneMessageThread request
	RestoreOneMessageThread(ctx context.Context, id IdPath, params *RestoreOneMessageThreadParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyMessages request
	RestoreManyMessages(ctx context.Context, params *RestoreManyMessagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneMessage request
	RestoreOneMessage(ctx context.Context, id IdPath, params *RestoreOneMessageParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyNoteTargets request
	RestoreManyNoteTargets(ctx context.Context, params *RestoreManyNoteTargetsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneNoteTarget request
	RestoreOneNoteTarget(ctx context.Context, id IdPath, params *RestoreOneNoteTargetParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyNotes request
	RestoreManyNotes(ctx context.Context, params *RestoreManyNotesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneNote request
	RestoreOneNote(ctx context.Context, id IdPath, params *RestoreOneNoteParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyOpportunities request
	RestoreManyOpportunities(ctx context.Context, params *RestoreManyOpportunitiesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneOpportunity request
	RestoreOneOpportunity(ctx context.Context, id IdPath, params *RestoreOneOpportunityParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyPeople request
	RestoreManyPeople(ctx context.Context, params *RestoreManyPeopleParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOnePerson request
	RestoreOnePerson(ctx context.Context, id IdPath, params *RestoreOnePersonParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyTaskTargets request
	RestoreManyTaskTargets(ctx context.Context, params *RestoreManyTaskTargetsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneTaskTarget request
	RestoreOneTaskTarget(ctx context.Context, id IdPath, params *RestoreOneTaskTargetParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyTasks request
	RestoreManyTasks(ctx context.Context, params *RestoreManyTasksParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneTask request
	RestoreOneTask(ctx context.Context, id IdPath, params *RestoreOneTaskParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyTimelineActivities request
	RestoreManyTimelineActivities(ctx context.Context, params *RestoreManyTimelineActivitiesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneTimelineActivity request
	RestoreOneTimelineActivity(ctx context.Context, id IdPath, params *RestoreOneTimelineActivityParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyWorkflowAutomatedTriggers request
	RestoreManyWorkflowAutomatedTriggers(ctx context.Context, params *RestoreManyWorkflowAutomatedTriggersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneWorkflowAutomatedTrigger request
	RestoreOneWorkflowAutomatedTrigger(ctx context.Context, id IdPath, params *RestoreOneWorkflowAutomatedTriggerParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyWorkflowRuns request
	RestoreManyWorkflowRuns(ctx context.Context, params *RestoreManyWorkflowRunsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneWorkflowRun request
	RestoreOneWorkflowRun(ctx context.Context, id IdPath, params *RestoreOneWorkflowRunParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyWorkflowVersions request
	RestoreManyWorkflowVersions(ctx context.Context, params *RestoreManyWorkflowVersionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneWorkflowVersion request
	RestoreOneWorkflowVersion(ctx context.Context, id IdPath, params *RestoreOneWorkflowVersionParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyWorkflows request
	RestoreManyWorkflows(ctx context.Context, params *RestoreManyWorkflowsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneWorkflow request
	RestoreOneWorkflow(ctx context.Context, id IdPath, params *RestoreOneWorkflowParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreManyWorkspaceMembers request
	RestoreManyWorkspaceMembers(ctx context.Context, params *RestoreManyWorkspaceMembersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RestoreOneWorkspaceMember request
	RestoreOneWorkspaceMember(ctx context.Context, id IdPath, params *RestoreOneWorkspaceMemberParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyTaskTargets request
	DeleteManyTaskTargets(ctx context.Context, params *DeleteManyTaskTargetsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyTaskTargets request
	FindManyTaskTargets(ctx context.Context, params *FindManyTaskTargetsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyTaskTargetsWithBody request with any body
	UpdateManyTaskTargetsWithBody(ctx context.Context, params *UpdateManyTaskTargetsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyTaskTargets(ctx context.Context, params *UpdateManyTaskTargetsParams, body UpdateManyTaskTargetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneTaskTargetWithBody request with any body
	CreateOneTaskTargetWithBody(ctx context.Context, params *CreateOneTaskTargetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneTaskTarget(ctx context.Context, params *CreateOneTaskTargetParams, body CreateOneTaskTargetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindTaskTargetDuplicatesWithBody request with any body
	FindTaskTargetDuplicatesWithBody(ctx context.Context, params *FindTaskTargetDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindTaskTargetDuplicates(ctx context.Context, params *FindTaskTargetDuplicatesParams, body FindTaskTargetDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyTaskTargetsWithBody request with any body
	MergeManyTaskTargetsWithBody(ctx context.Context, params *MergeManyTaskTargetsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyTaskTargets(ctx context.Context, params *MergeManyTaskTargetsParams, body MergeManyTaskTargetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneTaskTarget request
	DeleteOneTaskTarget(ctx context.Context, id IdPath, params *DeleteOneTaskTargetParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneTaskTarget request
	FindOneTaskTarget(ctx context.Context, id IdPath, params *FindOneTaskTargetParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneTaskTargetWithBody request with any body
	UpdateOneTaskTargetWithBody(ctx context.Context, id IdPath, params *UpdateOneTaskTargetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneTaskTarget(ctx context.Context, id IdPath, params *UpdateOneTaskTargetParams, body UpdateOneTaskTargetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyTasks request
	DeleteManyTasks(ctx context.Context, params *DeleteManyTasksParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyTasks request
	FindManyTasks(ctx context.Context, params *FindManyTasksParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyTasksWithBody request with any body
	UpdateManyTasksWithBody(ctx context.Context, params *UpdateManyTasksParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyTasks(ctx context.Context, params *UpdateManyTasksParams, body UpdateManyTasksJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneTaskWithBody request with any body
	CreateOneTaskWithBody(ctx context.Context, params *CreateOneTaskParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneTask(ctx context.Context, params *CreateOneTaskParams, body CreateOneTaskJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindTaskDuplicatesWithBody request with any body
	FindTaskDuplicatesWithBody(ctx context.Context, params *FindTaskDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindTaskDuplicates(ctx context.Context, params *FindTaskDuplicatesParams, body FindTaskDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyTasksWithBody request with any body
	MergeManyTasksWithBody(ctx context.Context, params *MergeManyTasksParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyTasks(ctx context.Context, params *MergeManyTasksParams, body MergeManyTasksJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneTask request
	DeleteOneTask(ctx context.Context, id IdPath, params *DeleteOneTaskParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneTask request
	FindOneTask(ctx context.Context, id IdPath, params *FindOneTaskParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneTaskWithBody request with any body
	UpdateOneTaskWithBody(ctx context.Context, id IdPath, params *UpdateOneTaskParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneTask(ctx context.Context, id IdPath, params *UpdateOneTaskParams, body UpdateOneTaskJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyTimelineActivities request
	DeleteManyTimelineActivities(ctx context.Context, params *DeleteManyTimelineActivitiesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyTimelineActivities request
	FindManyTimelineActivities(ctx context.Context, params *FindManyTimelineActivitiesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyTimelineActivitiesWithBody request with any body
	UpdateManyTimelineActivitiesWithBody(ctx context.Context, params *UpdateManyTimelineActivitiesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyTimelineActivities(ctx context.Context, params *UpdateManyTimelineActivitiesParams, body UpdateManyTimelineActivitiesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneTimelineActivityWithBody request with any body
	CreateOneTimelineActivityWithBody(ctx context.Context, params *CreateOneTimelineActivityParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneTimelineActivity(ctx context.Context, params *CreateOneTimelineActivityParams, body CreateOneTimelineActivityJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindTimelineActivityDuplicatesWithBody request with any body
	FindTimelineActivityDuplicatesWithBody(ctx context.Context, params *FindTimelineActivityDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindTimelineActivityDuplicates(ctx context.Context, params *FindTimelineActivityDuplicatesParams, body FindTimelineActivityDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyTimelineActivitiesWithBody request with any body
	MergeManyTimelineActivitiesWithBody(ctx context.Context, params *MergeManyTimelineActivitiesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyTimelineActivities(ctx context.Context, params *MergeManyTimelineActivitiesParams, body MergeManyTimelineActivitiesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneTimelineActivity request
	DeleteOneTimelineActivity(ctx context.Context, id IdPath, params *DeleteOneTimelineActivityParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneTimelineActivity request
	FindOneTimelineActivity(ctx context.Context, id IdPath, params *FindOneTimelineActivityParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneTimelineActivityWithBody request with any body
	UpdateOneTimelineActivityWithBody(ctx context.Context, id IdPath, params *UpdateOneTimelineActivityParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneTimelineActivity(ctx context.Context, id IdPath, params *UpdateOneTimelineActivityParams, body UpdateOneTimelineActivityJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyWorkflowAutomatedTriggers request
	DeleteManyWorkflowAutomatedTriggers(ctx context.Context, params *DeleteManyWorkflowAutomatedTriggersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyWorkflowAutomatedTriggers request
	FindManyWorkflowAutomatedTriggers(ctx context.Context, params *FindManyWorkflowAutomatedTriggersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyWorkflowAutomatedTriggersWithBody request with any body
	UpdateManyWorkflowAutomatedTriggersWithBody(ctx context.Context, params *UpdateManyWorkflowAutomatedTriggersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyWorkflowAutomatedTriggers(ctx context.Context, params *UpdateManyWorkflowAutomatedTriggersParams, body UpdateManyWorkflowAutomatedTriggersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneWorkflowAutomatedTriggerWithBody request with any body
	CreateOneWorkflowAutomatedTriggerWithBody(ctx context.Context, params *CreateOneWorkflowAutomatedTriggerParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneWorkflowAutomatedTrigger(ctx context.Context, params *CreateOneWorkflowAutomatedTriggerParams, body CreateOneWorkflowAutomatedTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindWorkflowAutomatedTriggerDuplicatesWithBody request with any body
	FindWorkflowAutomatedTriggerDuplicatesWithBody(ctx context.Context, params *FindWorkflowAutomatedTriggerDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindWorkflowAutomatedTriggerDuplicates(ctx context.Context, params *FindWorkflowAutomatedTriggerDuplicatesParams, body FindWorkflowAutomatedTriggerDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyWorkflowAutomatedTriggersWithBody request with any body
	MergeManyWorkflowAutomatedTriggersWithBody(ctx context.Context, params *MergeManyWorkflowAutomatedTriggersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyWorkflowAutomatedTriggers(ctx context.Context, params *MergeManyWorkflowAutomatedTriggersParams, body MergeManyWorkflowAutomatedTriggersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneWorkflowAutomatedTrigger request
	DeleteOneWorkflowAutomatedTrigger(ctx context.Context, id IdPath, params *DeleteOneWorkflowAutomatedTriggerParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneWorkflowAutomatedTrigger request
	FindOneWorkflowAutomatedTrigger(ctx context.Context, id IdPath, params *FindOneWorkflowAutomatedTriggerParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneWorkflowAutomatedTriggerWithBody request with any body
	UpdateOneWorkflowAutomatedTriggerWithBody(ctx context.Context, id IdPath, params *UpdateOneWorkflowAutomatedTriggerParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneWorkflowAutomatedTrigger(ctx context.Context, id IdPath, params *UpdateOneWorkflowAutomatedTriggerParams, body UpdateOneWorkflowAutomatedTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyWorkflowRuns request
	DeleteManyWorkflowRuns(ctx context.Context, params *DeleteManyWorkflowRunsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyWorkflowRuns request
	FindManyWorkflowRuns(ctx context.Context, params *FindManyWorkflowRunsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyWorkflowRunsWithBody request with any body
	UpdateManyWorkflowRunsWithBody(ctx context.Context, params *UpdateManyWorkflowRunsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyWorkflowRuns(ctx context.Context, params *UpdateManyWorkflowRunsParams, body UpdateManyWorkflowRunsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneWorkflowRunWithBody request with any body
	CreateOneWorkflowRunWithBody(ctx context.Context, params *CreateOneWorkflowRunParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneWorkflowRun(ctx context.Context, params *CreateOneWorkflowRunParams, body CreateOneWorkflowRunJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindWorkflowRunDuplicatesWithBody request with any body
	FindWorkflowRunDuplicatesWithBody(ctx context.Context, params *FindWorkflowRunDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindWorkflowRunDuplicates(ctx context.Context, params *FindWorkflowRunDuplicatesParams, body FindWorkflowRunDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyWorkflowRunsWithBody request with any body
	MergeManyWorkflowRunsWithBody(ctx context.Context, params *MergeManyWorkflowRunsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyWorkflowRuns(ctx context.Context, params *MergeManyWorkflowRunsParams, body MergeManyWorkflowRunsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneWorkflowRun request
	DeleteOneWorkflowRun(ctx context.Context, id IdPath, params *DeleteOneWorkflowRunParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneWorkflowRun request
	FindOneWorkflowRun(ctx context.Context, id IdPath, params *FindOneWorkflowRunParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneWorkflowRunWithBody request with any body
	UpdateOneWorkflowRunWithBody(ctx context.Context, id IdPath, params *UpdateOneWorkflowRunParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneWorkflowRun(ctx context.Context, id IdPath, params *UpdateOneWorkflowRunParams, body UpdateOneWorkflowRunJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyWorkflowVersions request
	DeleteManyWorkflowVersions(ctx context.Context, params *DeleteManyWorkflowVersionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyWorkflowVersions request
	FindManyWorkflowVersions(ctx context.Context, params *FindManyWorkflowVersionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyWorkflowVersionsWithBody request with any body
	UpdateManyWorkflowVersionsWithBody(ctx context.Context, params *UpdateManyWorkflowVersionsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyWorkflowVersions(ctx context.Context, params *UpdateManyWorkflowVersionsParams, body UpdateManyWorkflowVersionsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneWorkflowVersionWithBody request with any body
	CreateOneWorkflowVersionWithBody(ctx context.Context, params *CreateOneWorkflowVersionParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneWorkflowVersion(ctx context.Context, params *CreateOneWorkflowVersionParams, body CreateOneWorkflowVersionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindWorkflowVersionDuplicatesWithBody request with any body
	FindWorkflowVersionDuplicatesWithBody(ctx context.Context, params *FindWorkflowVersionDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindWorkflowVersionDuplicates(ctx context.Context, params *FindWorkflowVersionDuplicatesParams, body FindWorkflowVersionDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyWorkflowVersionsWithBody request with any body
	MergeManyWorkflowVersionsWithBody(ctx context.Context, params *MergeManyWorkflowVersionsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyWorkflowVersions(ctx context.Context, params *MergeManyWorkflowVersionsParams, body MergeManyWorkflowVersionsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneWorkflowVersion request
	DeleteOneWorkflowVersion(ctx context.Context, id IdPath, params *DeleteOneWorkflowVersionParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneWorkflowVersion request
	FindOneWorkflowVersion(ctx context.Context, id IdPath, params *FindOneWorkflowVersionParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneWorkflowVersionWithBody request with any body
	UpdateOneWorkflowVersionWithBody(ctx context.Context, id IdPath, params *UpdateOneWorkflowVersionParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneWorkflowVersion(ctx context.Context, id IdPath, params *UpdateOneWorkflowVersionParams, body UpdateOneWorkflowVersionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyWorkflows request
	DeleteManyWorkflows(ctx context.Context, params *DeleteManyWorkflowsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyWorkflows request
	FindManyWorkflows(ctx context.Context, params *FindManyWorkflowsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyWorkflowsWithBody request with any body
	UpdateManyWorkflowsWithBody(ctx context.Context, params *UpdateManyWorkflowsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyWorkflows(ctx context.Context, params *UpdateManyWorkflowsParams, body UpdateManyWorkflowsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneWorkflowWithBody request with any body
	CreateOneWorkflowWithBody(ctx context.Context, params *CreateOneWorkflowParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneWorkflow(ctx context.Context, params *CreateOneWorkflowParams, body CreateOneWorkflowJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindWorkflowDuplicatesWithBody request with any body
	FindWorkflowDuplicatesWithBody(ctx context.Context, params *FindWorkflowDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindWorkflowDuplicates(ctx context.Context, params *FindWorkflowDuplicatesParams, body FindWorkflowDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyWorkflowsWithBody request with any body
	MergeManyWorkflowsWithBody(ctx context.Context, params *MergeManyWorkflowsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyWorkflows(ctx context.Context, params *MergeManyWorkflowsParams, body MergeManyWorkflowsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneWorkflow request
	DeleteOneWorkflow(ctx context.Context, id IdPath, params *DeleteOneWorkflowParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneWorkflow request
	FindOneWorkflow(ctx context.Context, id IdPath, params *FindOneWorkflowParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneWorkflowWithBody request with any body
	UpdateOneWorkflowWithBody(ctx context.Context, id IdPath, params *UpdateOneWorkflowParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneWorkflow(ctx context.Context, id IdPath, params *UpdateOneWorkflowParams, body UpdateOneWorkflowJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteManyWorkspaceMembers request
	DeleteManyWorkspaceMembers(ctx context.Context, params *DeleteManyWorkspaceMembersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindManyWorkspaceMembers request
	FindManyWorkspaceMembers(ctx context.Context, params *FindManyWorkspaceMembersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateManyWorkspaceMembersWithBody request with any body
	UpdateManyWorkspaceMembersWithBody(ctx context.Context, params *UpdateManyWorkspaceMembersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateManyWorkspaceMembers(ctx context.Context, params *UpdateManyWorkspaceMembersParams, body UpdateManyWorkspaceMembersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateOneWorkspaceMemberWithBody request with any body
	CreateOneWorkspaceMemberWithBody(ctx context.Context, params *CreateOneWorkspaceMemberParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateOneWorkspaceMember(ctx context.Context, params *CreateOneWorkspaceMemberParams, body CreateOneWorkspaceMemberJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindWorkspaceMemberDuplicatesWithBody request with any body
	FindWorkspaceMemberDuplicatesWithBody(ctx context.Context, params *FindWorkspaceMemberDuplicatesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FindWorkspaceMemberDuplicates(ctx context.Context, params *FindWorkspaceMemberDuplicatesParams, body FindWorkspaceMemberDuplicatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MergeManyWorkspaceMembersWithBody request with any body
	MergeManyWorkspaceMembersWithBody(ctx context.Context, params *MergeManyWorkspaceMembersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MergeManyWorkspaceMembers(ctx context.Context, params *MergeManyWorkspaceMembersParams, body MergeManyWorkspaceMembersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteOneWorkspaceMember request
	DeleteOneWorkspaceMember(ctx context.Context, id IdPath, params *DeleteOneWorkspaceMemberParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindOneWorkspaceMember request
	FindOneWorkspaceMember(ctx context.Context, id IdPath, params *FindOneWorkspaceMemberParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateOneWorkspaceMemberWithBody request with any body
	UpdateOneWorkspaceMemberWithBody(ctx context.Context, id IdPath, params *UpdateOneWorkspaceMemberParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateOneWorkspaceMember(ctx context.Context, id IdPath, params *UpdateOneWorkspaceMemberParams, body UpdateOneWorkspaceMemberJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}
