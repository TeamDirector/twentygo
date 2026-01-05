package twentygo

import (
	"context"
	"net/http"
)

// The interface specification for the client above.
type ClientInterface interface {
	// DeleteManyAttachments request
	DeleteManyAttachments(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyAttachments request
	FindManyAttachments(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyAttachments(ctx context.Context, params *UpdateManyParams, body NewAttachment) (*http.Response, error)

	CreateOneAttachment(ctx context.Context, params *CreateOneParams, body Attachment) (*http.Response, error)

	FindAttachmentDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyAttachments(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneAttachment request
	DeleteOneAttachment(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneAttachment request
	FindOneAttachment(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneAttachment(ctx context.Context, id string, params *UpdateOneParams, body NewAttachment) (*http.Response, error)

	CreateManyAttachments(ctx context.Context, params *CreateManyParams, body []NewAttachment) (*http.Response, error)

	CreateManyBlocklists(ctx context.Context, params *CreateManyParams, body []NewBlocklist) (*http.Response, error)

	CreateManyCalendarChannelEventAssociations(ctx context.Context, params *CreateManyParams, body []NewCalendarChannelEventAssociation) (*http.Response, error)

	CreateManyCalendarChannels(ctx context.Context, params *CreateManyParams, body []NewCalendarChannel) (*http.Response, error)

	CreateManyCalendarEventParticipants(ctx context.Context, params *CreateManyParams, body []NewCalendarEventParticipant) (*http.Response, error)

	CreateManyCalendarEvents(ctx context.Context, params *CreateManyParams, body []NewCalendarEvent) (*http.Response, error)

	CreateManyCompanies(ctx context.Context, params *CreateManyParams, body []NewCompany) (*http.Response, error)

	CreateManyConnectedAccounts(ctx context.Context, params *CreateManyParams, body []NewConnectedAccount) (*http.Response, error)

	CreateManyDashboards(ctx context.Context, params *CreateManyParams, body []NewDashboard) (*http.Response, error)

	CreateManyFavoriteFolders(ctx context.Context, params *CreateManyParams, body []NewFavoriteFolder) (*http.Response, error)

	CreateManyFavorites(ctx context.Context, params *CreateManyParams, body []NewFavorite) (*http.Response, error)

	CreateManyMessageChannelMessageAssociations(ctx context.Context, params *CreateManyParams, body []NewMessageChannelMessageAssociation) (*http.Response, error)

	CreateManyMessageChannels(ctx context.Context, params *CreateManyParams, body []NewMessageChannel) (*http.Response, error)

	CreateManyMessageFolders(ctx context.Context, params *CreateManyParams, body []NewMessageFolder) (*http.Response, error)

	CreateManyMessageParticipants(ctx context.Context, params *CreateManyParams, body []NewMessageParticipant) (*http.Response, error)

	CreateManyMessageThreads(ctx context.Context, params *CreateManyParams, body []NewMessageThread) (*http.Response, error)

	CreateManyMessages(ctx context.Context, params *CreateManyParams, body []NewMessage) (*http.Response, error)

	CreateManyNoteTargets(ctx context.Context, params *CreateManyParams, body []NewNoteTarget) (*http.Response, error)

	CreateManyNotes(ctx context.Context, params *CreateManyParams, body []NewNote) (*http.Response, error)

	CreateManyOpportunities(ctx context.Context, params *CreateManyParams, body []NewOpportunity) (*http.Response, error)

	CreateManyPeople(ctx context.Context, params *CreateManyParams, body []NewPerson) (*http.Response, error)

	CreateManyTaskTargets(ctx context.Context, params *CreateManyParams, body []NewTaskTarget) (*http.Response, error)

	CreateManyTasks(ctx context.Context, params *CreateManyParams, body []NewTask) (*http.Response, error)

	CreateManyTimelineActivities(ctx context.Context, params *CreateManyParams, body []NewTimelineActivity) (*http.Response, error)

	CreateManyWorkflowAutomatedTriggers(ctx context.Context, params *CreateManyParams, body []NewWorkflowAutomatedTrigger) (*http.Response, error)

	CreateManyWorkflowRuns(ctx context.Context, params *CreateManyParams, body []NewWorkflowRun) (*http.Response, error)

	CreateManyWorkflowVersions(ctx context.Context, params *CreateManyParams, body []NewWorkflowVersion) (*http.Response, error)

	CreateManyWorkflows(ctx context.Context, params *CreateManyParams, body []NewWorkflow) (*http.Response, error)

	CreateManyWorkspaceMembers(ctx context.Context, params *CreateManyParams, body []NewWorkspaceMember) (*http.Response, error)

	// DeleteManyBlocklists request
	DeleteManyBlocklists(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyBlocklists request
	FindManyBlocklists(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyBlocklists(ctx context.Context, params *UpdateManyParams, body NewBlocklist) (*http.Response, error)

	CreateOneBlocklist(ctx context.Context, params *CreateOneParams, body NewBlocklist) (*http.Response, error)

	FindBlocklistDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyBlocklists(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneBlocklist request
	DeleteOneBlocklist(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneBlocklist request
	FindOneBlocklist(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneBlocklist(ctx context.Context, id string, params *UpdateOneParams, body NewBlocklist) (*http.Response, error)

	// DeleteManyCalendarChannelEventAssociations request
	DeleteManyCalendarChannelEventAssociations(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyCalendarChannelEventAssociations request
	FindManyCalendarChannelEventAssociations(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyCalendarChannelEventAssociations(ctx context.Context, params *UpdateManyParams, body NewCalendarChannelEventAssociation) (*http.Response, error)

	CreateOneCalendarChannelEventAssociation(ctx context.Context, params *CreateOneParams, body NewCalendarChannelEventAssociation) (*http.Response, error)

	FindCalendarChannelEventAssociationDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyCalendarChannelEventAssociations(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneCalendarChannelEventAssociation request
	DeleteOneCalendarChannelEventAssociation(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneCalendarChannelEventAssociation request
	FindOneCalendarChannelEventAssociation(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneCalendarChannelEventAssociation(ctx context.Context, id string, params *UpdateOneParams, body NewCalendarChannelEventAssociation) (*http.Response, error)

	// DeleteManyCalendarChannels request
	DeleteManyCalendarChannels(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyCalendarChannels request
	FindManyCalendarChannels(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyCalendarChannels(ctx context.Context, params *UpdateManyParams, body NewCalendarChannel) (*http.Response, error)

	CreateOneCalendarChannel(ctx context.Context, params *CreateOneParams, body NewCalendarChannel) (*http.Response, error)

	FindCalendarChannelDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyCalendarChannels(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneCalendarChannel request
	DeleteOneCalendarChannel(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneCalendarChannel request
	FindOneCalendarChannel(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneCalendarChannel(ctx context.Context, id string, params *UpdateOneParams, body NewCalendarChannel) (*http.Response, error)

	// DeleteManyCalendarEventParticipants request
	DeleteManyCalendarEventParticipants(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyCalendarEventParticipants request
	FindManyCalendarEventParticipants(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyCalendarEventParticipants(ctx context.Context, params *UpdateManyParams, body NewCalendarEventParticipant) (*http.Response, error)

	CreateOneCalendarEventParticipant(ctx context.Context, params *CreateOneParams, body NewCalendarEventParticipant) (*http.Response, error)

	FindCalendarEventParticipantDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyCalendarEventParticipants(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneCalendarEventParticipant request
	DeleteOneCalendarEventParticipant(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneCalendarEventParticipant request
	FindOneCalendarEventParticipant(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneCalendarEventParticipant(ctx context.Context, id string, params *UpdateOneParams, body NewCalendarEventParticipant) (*http.Response, error)

	// DeleteManyCalendarEvents request
	DeleteManyCalendarEvents(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyCalendarEvents request
	FindManyCalendarEvents(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyCalendarEvents(ctx context.Context, params *UpdateManyParams, body NewCalendarEvent) (*http.Response, error)

	CreateOneCalendarEvent(ctx context.Context, params *CreateOneParams, body NewCalendarEvent) (*http.Response, error)

	FindCalendarEventDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyCalendarEvents(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneCalendarEvent request
	DeleteOneCalendarEvent(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneCalendarEvent request
	FindOneCalendarEvent(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneCalendarEvent(ctx context.Context, id string, params *UpdateOneParams, body NewCalendarEvent) (*http.Response, error)

	// DeleteManyCompanies request
	DeleteManyCompanies(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyCompanies request
	FindManyCompanies(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyCompanies(ctx context.Context, params *UpdateManyParams, body NewCompany) (*http.Response, error)

	CreateOneCompany(ctx context.Context, params *CreateOneParams, body NewCompany) (*http.Response, error)

	FindCompanyDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyCompanies(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneCompany request
	DeleteOneCompany(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneCompany request
	FindOneCompany(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneCompany(ctx context.Context, id string, params *UpdateOneParams, body NewCompany) (*http.Response, error)

	// DeleteManyConnectedAccounts request
	DeleteManyConnectedAccounts(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyConnectedAccounts request
	FindManyConnectedAccounts(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyConnectedAccounts(ctx context.Context, params *UpdateManyParams, body NewConnectedAccount) (*http.Response, error)

	CreateOneConnectedAccount(ctx context.Context, params *CreateOneParams, body NewConnectedAccount) (*http.Response, error)

	FindConnectedAccountDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyConnectedAccounts(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneConnectedAccount request
	DeleteOneConnectedAccount(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneConnectedAccount request
	FindOneConnectedAccount(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneConnectedAccount(ctx context.Context, id string, params *UpdateOneParams, body NewConnectedAccount) (*http.Response, error)

	// DeleteManyDashboards request
	DeleteManyDashboards(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyDashboards request
	FindManyDashboards(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyDashboards(ctx context.Context, params *UpdateManyParams, body NewDashboard) (*http.Response, error)

	CreateOneDashboard(ctx context.Context, params *CreateOneParams, body NewDashboard) (*http.Response, error)

	FindDashboardDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyDashboards(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneDashboard request
	DeleteOneDashboard(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneDashboard request
	FindOneDashboard(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneDashboard(ctx context.Context, id string, params *UpdateOneParams, body NewDashboard) (*http.Response, error)

	// DuplicateDashboard request
	DuplicateDashboard(ctx context.Context, id string) (*http.Response, error)

	// DeleteManyFavoriteFolders request
	DeleteManyFavoriteFolders(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyFavoriteFolders request
	FindManyFavoriteFolders(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyFavoriteFolders(ctx context.Context, params *UpdateManyParams, body NewFavoriteFolder) (*http.Response, error)

	CreateOneFavoriteFolder(ctx context.Context, params *CreateOneParams, body NewFavoriteFolder) (*http.Response, error)

	FindFavoriteFolderDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyFavoriteFolders(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneFavoriteFolder request
	DeleteOneFavoriteFolder(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneFavoriteFolder request
	FindOneFavoriteFolder(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneFavoriteFolder(ctx context.Context, id string, params *UpdateOneParams, body NewFavoriteFolder) (*http.Response, error)

	// DeleteManyFavorites request
	DeleteManyFavorites(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyFavorites request
	FindManyFavorites(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyFavorites(ctx context.Context, params *UpdateManyParams, body NewFavorite) (*http.Response, error)

	CreateOneFavorite(ctx context.Context, params *CreateOneParams, body NewFavorite) (*http.Response, error)

	FindFavoriteDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyFavorites(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneFavorite request
	DeleteOneFavorite(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneFavorite request
	FindOneFavorite(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneFavorite(ctx context.Context, id string, params *UpdateOneParams, body NewFavorite) (*http.Response, error)

	// DeleteManyMessageChannelMessageAssociations request
	DeleteManyMessageChannelMessageAssociations(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyMessageChannelMessageAssociations request
	FindManyMessageChannelMessageAssociations(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyMessageChannelMessageAssociations(ctx context.Context, params *UpdateManyParams, body NewMessageChannelMessageAssociation) (*http.Response, error)

	CreateOneMessageChannelMessageAssociation(ctx context.Context, params *CreateOneParams, body NewMessageChannelMessageAssociation) (*http.Response, error)

	FindMessageChannelMessageAssociationDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyMessageChannelMessageAssociations(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneMessageChannelMessageAssociation request
	DeleteOneMessageChannelMessageAssociation(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneMessageChannelMessageAssociation request
	FindOneMessageChannelMessageAssociation(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneMessageChannelMessageAssociation(ctx context.Context, id string, params *UpdateOneParams, body NewMessageChannelMessageAssociation) (*http.Response, error)

	// DeleteManyMessageChannels request
	DeleteManyMessageChannels(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyMessageChannels request
	FindManyMessageChannels(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyMessageChannels(ctx context.Context, params *UpdateManyParams, body NewMessageChannel) (*http.Response, error)

	CreateOneMessageChannel(ctx context.Context, params *CreateOneParams, body NewMessageChannel) (*http.Response, error)

	FindMessageChannelDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyMessageChannels(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneMessageChannel request
	DeleteOneMessageChannel(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneMessageChannel request
	FindOneMessageChannel(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneMessageChannel(ctx context.Context, id string, params *UpdateOneParams, body NewMessageChannel) (*http.Response, error)

	// DeleteManyMessageFolders request
	DeleteManyMessageFolders(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyMessageFolders request
	FindManyMessageFolders(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyMessageFolders(ctx context.Context, params *UpdateManyParams, body NewMessageFolder) (*http.Response, error)

	CreateOneMessageFolder(ctx context.Context, params *CreateOneParams, body NewMessageFolder) (*http.Response, error)

	FindMessageFolderDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyMessageFolders(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneMessageFolder request
	DeleteOneMessageFolder(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneMessageFolder request
	FindOneMessageFolder(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneMessageFolder(ctx context.Context, id string, params *UpdateOneParams, body NewMessageFolder) (*http.Response, error)

	// DeleteManyMessageParticipants request
	DeleteManyMessageParticipants(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyMessageParticipants request
	FindManyMessageParticipants(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyMessageParticipants(ctx context.Context, params *UpdateManyParams, body NewMessageParticipant) (*http.Response, error)

	CreateOneMessageParticipant(ctx context.Context, params *CreateOneParams, body NewMessageParticipant) (*http.Response, error)

	FindMessageParticipantDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyMessageParticipants(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneMessageParticipant request
	DeleteOneMessageParticipant(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneMessageParticipant request
	FindOneMessageParticipant(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneMessageParticipant(ctx context.Context, id string, params *UpdateOneParams, body NewMessageParticipant) (*http.Response, error)

	// DeleteManyMessageThreads request
	DeleteManyMessageThreads(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyMessageThreads request
	FindManyMessageThreads(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyMessageThreads(ctx context.Context, params *UpdateManyParams, body NewMessageThread) (*http.Response, error)

	CreateOneMessageThread(ctx context.Context, params *CreateOneParams, body NewMessageThread) (*http.Response, error)

	FindMessageThreadDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyMessageThreads(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneMessageThread request
	DeleteOneMessageThread(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneMessageThread request
	FindOneMessageThread(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneMessageThread(ctx context.Context, id string, params *UpdateOneParams, body NewMessageThread) (*http.Response, error)

	// DeleteManyMessages request
	DeleteManyMessages(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyMessages request
	FindManyMessages(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyMessages(ctx context.Context, params *UpdateManyParams, body NewMessage) (*http.Response, error)

	CreateOneMessage(ctx context.Context, params *CreateOneParams, body NewMessage) (*http.Response, error)

	FindMessageDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyMessages(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneMessage request
	DeleteOneMessage(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneMessage request
	FindOneMessage(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneMessage(ctx context.Context, id string, params *UpdateOneParams, body NewMessage) (*http.Response, error)

	// DeleteManyNoteTargets request
	DeleteManyNoteTargets(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyNoteTargets request
	FindManyNoteTargets(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyNoteTargets(ctx context.Context, params *UpdateManyParams, body NewNoteTarget) (*http.Response, error)

	CreateOneNoteTarget(ctx context.Context, params *CreateOneParams, body NewNoteTarget) (*http.Response, error)

	FindNoteTargetDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyNoteTargets(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneNoteTarget request
	DeleteOneNoteTarget(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneNoteTarget request
	FindOneNoteTarget(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneNoteTarget(ctx context.Context, id string, params *UpdateOneParams, body NewNoteTarget) (*http.Response, error)

	// DeleteManyNotes request
	DeleteManyNotes(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyNotes request
	FindManyNotes(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyNotes(ctx context.Context, params *UpdateManyParams, body NewNote) (*http.Response, error)

	CreateOneNote(ctx context.Context, params *CreateOneParams, body NewNote) (*http.Response, error)

	FindNoteDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyNotes(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneNote request
	DeleteOneNote(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneNote request
	FindOneNote(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneNote(ctx context.Context, id string, params *UpdateOneParams, body NewNote) (*http.Response, error)

	// GetOpenApiSchema request
	GetOpenApiSchema(ctx context.Context) (*http.Response, error)

	// DeleteManyOpportunities request
	DeleteManyOpportunities(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyOpportunities request
	FindManyOpportunities(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyOpportunities(ctx context.Context, params *UpdateManyParams, body NewOpportunity) (*http.Response, error)

	CreateOneOpportunity(ctx context.Context, params *CreateOneParams, body NewOpportunity) (*http.Response, error)

	FindOpportunityDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyOpportunities(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneOpportunity request
	DeleteOneOpportunity(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneOpportunity request
	FindOneOpportunity(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneOpportunity(ctx context.Context, id string, params *UpdateOneParams, body NewOpportunity) (*http.Response, error)

	// DeleteManyPeople request
	DeleteManyPeople(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyPeople request
	FindManyPeople(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyPeople(ctx context.Context, params *UpdateManyParams, body NewPerson) (*http.Response, error)

	CreateOnePerson(ctx context.Context, params *CreateOneParams, body NewPerson) (*http.Response, error)

	FindPersonDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyPeople(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOnePerson request
	DeleteOnePerson(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOnePerson request
	FindOnePerson(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOnePerson(ctx context.Context, id string, params *UpdateOneParams, body NewPerson) (*http.Response, error)

	// RestoreManyAttachments request
	RestoreManyAttachments(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneAttachment request
	RestoreOneAttachment(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyBlocklists request
	RestoreManyBlocklists(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneBlocklist request
	RestoreOneBlocklist(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyCalendarChannelEventAssociations request
	RestoreManyCalendarChannelEventAssociations(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneCalendarChannelEventAssociation request
	RestoreOneCalendarChannelEventAssociation(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyCalendarChannels request
	RestoreManyCalendarChannels(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneCalendarChannel request
	RestoreOneCalendarChannel(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyCalendarEventParticipants request
	RestoreManyCalendarEventParticipants(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneCalendarEventParticipant request
	RestoreOneCalendarEventParticipant(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyCalendarEvents request
	RestoreManyCalendarEvents(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneCalendarEvent request
	RestoreOneCalendarEvent(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyCompanies request
	RestoreManyCompanies(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneCompany request
	RestoreOneCompany(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyConnectedAccounts request
	RestoreManyConnectedAccounts(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneConnectedAccount request
	RestoreOneConnectedAccount(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyDashboards request
	RestoreManyDashboards(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneDashboard request
	RestoreOneDashboard(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyFavoriteFolders request
	RestoreManyFavoriteFolders(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneFavoriteFolder request
	RestoreOneFavoriteFolder(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyFavorites request
	RestoreManyFavorites(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneFavorite request
	RestoreOneFavorite(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyMessageChannelMessageAssociations request
	RestoreManyMessageChannelMessageAssociations(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneMessageChannelMessageAssociation request
	RestoreOneMessageChannelMessageAssociation(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyMessageChannels request
	RestoreManyMessageChannels(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneMessageChannel request
	RestoreOneMessageChannel(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyMessageFolders request
	RestoreManyMessageFolders(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneMessageFolder request
	RestoreOneMessageFolder(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyMessageParticipants request
	RestoreManyMessageParticipants(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneMessageParticipant request
	RestoreOneMessageParticipant(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyMessageThreads request
	RestoreManyMessageThreads(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneMessageThread request
	RestoreOneMessageThread(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyMessages request
	RestoreManyMessages(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneMessage request
	RestoreOneMessage(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyNoteTargets request
	RestoreManyNoteTargets(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneNoteTarget request
	RestoreOneNoteTarget(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyNotes request
	RestoreManyNotes(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneNote request
	RestoreOneNote(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyOpportunities request
	RestoreManyOpportunities(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneOpportunity request
	RestoreOneOpportunity(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyPeople request
	RestoreManyPeople(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOnePerson request
	RestoreOnePerson(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyTaskTargets request
	RestoreManyTaskTargets(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneTaskTarget request
	RestoreOneTaskTarget(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyTasks request
	RestoreManyTasks(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneTask request
	RestoreOneTask(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyTimelineActivities request
	RestoreManyTimelineActivities(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneTimelineActivity request
	RestoreOneTimelineActivity(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyWorkflowAutomatedTriggers request
	RestoreManyWorkflowAutomatedTriggers(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneWorkflowAutomatedTrigger request
	RestoreOneWorkflowAutomatedTrigger(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyWorkflowRuns request
	RestoreManyWorkflowRuns(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneWorkflowRun request
	RestoreOneWorkflowRun(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyWorkflowVersions request
	RestoreManyWorkflowVersions(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneWorkflowVersion request
	RestoreOneWorkflowVersion(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyWorkflows request
	RestoreManyWorkflows(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneWorkflow request
	RestoreOneWorkflow(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// RestoreManyWorkspaceMembers request
	RestoreManyWorkspaceMembers(ctx context.Context, params *RestoreManyParams) (*http.Response, error)

	// RestoreOneWorkspaceMember request
	RestoreOneWorkspaceMember(ctx context.Context, id string, params *RestoreOneParams) (*http.Response, error)

	// DeleteManyTaskTargets request
	DeleteManyTaskTargets(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyTaskTargets request
	FindManyTaskTargets(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyTaskTargets(ctx context.Context, params *UpdateManyParams, body NewTaskTarget) (*http.Response, error)

	CreateOneTaskTarget(ctx context.Context, params *CreateOneParams, body NewTaskTarget) (*http.Response, error)

	FindTaskTargetDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyTaskTargets(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneTaskTarget request
	DeleteOneTaskTarget(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneTaskTarget request
	FindOneTaskTarget(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneTaskTarget(ctx context.Context, id string, params *UpdateOneParams, body NewTaskTarget) (*http.Response, error)

	// DeleteManyTasks request
	DeleteManyTasks(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyTasks request
	FindManyTasks(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyTasks(ctx context.Context, params *UpdateManyParams, body NewTask) (*http.Response, error)

	CreateOneTask(ctx context.Context, params *CreateOneParams, body NewTask) (*http.Response, error)

	FindTaskDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyTasks(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneTask request
	DeleteOneTask(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneTask request
	FindOneTask(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneTask(ctx context.Context, id string, params *UpdateOneParams, body NewTask) (*http.Response, error)

	// DeleteManyTimelineActivities request
	DeleteManyTimelineActivities(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyTimelineActivities request
	FindManyTimelineActivities(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyTimelineActivities(ctx context.Context, params *UpdateManyParams, body NewTimelineActivity) (*http.Response, error)

	CreateOneTimelineActivity(ctx context.Context, params *CreateOneParams, body NewTimelineActivity) (*http.Response, error)

	FindTimelineActivityDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyTimelineActivities(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneTimelineActivity request
	DeleteOneTimelineActivity(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneTimelineActivity request
	FindOneTimelineActivity(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneTimelineActivity(ctx context.Context, id string, params *UpdateOneParams, body NewTimelineActivity) (*http.Response, error)

	// DeleteManyWorkflowAutomatedTriggers request
	DeleteManyWorkflowAutomatedTriggers(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyWorkflowAutomatedTriggers request
	FindManyWorkflowAutomatedTriggers(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyWorkflowAutomatedTriggers(ctx context.Context, params *UpdateManyParams, body NewWorkflowAutomatedTrigger) (*http.Response, error)

	CreateOneWorkflowAutomatedTrigger(ctx context.Context, params *CreateOneParams, body NewWorkflowAutomatedTrigger) (*http.Response, error)

	FindWorkflowAutomatedTriggerDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyWorkflowAutomatedTriggers(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneWorkflowAutomatedTrigger request
	DeleteOneWorkflowAutomatedTrigger(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneWorkflowAutomatedTrigger request
	FindOneWorkflowAutomatedTrigger(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneWorkflowAutomatedTrigger(ctx context.Context, id string, params *UpdateOneParams, body NewWorkflowAutomatedTrigger) (*http.Response, error)

	// DeleteManyWorkflowRuns request
	DeleteManyWorkflowRuns(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyWorkflowRuns request
	FindManyWorkflowRuns(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyWorkflowRuns(ctx context.Context, params *UpdateManyParams, body NewWorkflowRun) (*http.Response, error)

	CreateOneWorkflowRun(ctx context.Context, params *CreateOneParams, body NewWorkflowRun) (*http.Response, error)

	FindWorkflowRunDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyWorkflowRuns(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneWorkflowRun request
	DeleteOneWorkflowRun(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneWorkflowRun request
	FindOneWorkflowRun(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneWorkflowRun(ctx context.Context, id string, params *UpdateOneParams, body NewWorkflowRun) (*http.Response, error)

	// DeleteManyWorkflowVersions request
	DeleteManyWorkflowVersions(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyWorkflowVersions request
	FindManyWorkflowVersions(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyWorkflowVersions(ctx context.Context, params *UpdateManyParams, body NewWorkflowVersion) (*http.Response, error)

	CreateOneWorkflowVersion(ctx context.Context, params *CreateOneParams, body NewWorkflowVersion) (*http.Response, error)

	FindWorkflowVersionDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyWorkflowVersions(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneWorkflowVersion request
	DeleteOneWorkflowVersion(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneWorkflowVersion request
	FindOneWorkflowVersion(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneWorkflowVersion(ctx context.Context, id string, params *UpdateOneParams, body NewWorkflowVersion) (*http.Response, error)

	// DeleteManyWorkflows request
	DeleteManyWorkflows(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyWorkflows request
	FindManyWorkflows(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyWorkflows(ctx context.Context, params *UpdateManyParams, body NewWorkflow) (*http.Response, error)

	CreateOneWorkflow(ctx context.Context, params *CreateOneParams, body NewWorkflow) (*http.Response, error)

	FindWorkflowDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyWorkflows(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneWorkflow request
	DeleteOneWorkflow(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneWorkflow request
	FindOneWorkflow(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneWorkflow(ctx context.Context, id string, params *UpdateOneParams, body NewWorkflow) (*http.Response, error)

	// DeleteManyWorkspaceMembers request
	DeleteManyWorkspaceMembers(ctx context.Context, params *DeleteManyParams) (*http.Response, error)

	// FindManyWorkspaceMembers request
	FindManyWorkspaceMembers(ctx context.Context, params FindManyParams) (*http.Response, error)

	UpdateManyWorkspaceMembers(ctx context.Context, params *UpdateManyParams, body NewWorkspaceMember) (*http.Response, error)

	CreateOneWorkspaceMember(ctx context.Context, params *CreateOneParams, body NewWorkspaceMember) (*http.Response, error)

	FindWorkspaceMemberDuplicates(ctx context.Context, params *FindDuplicatesParams, body FindDuplicatesBody) (*http.Response, error)

	MergeManyWorkspaceMembers(ctx context.Context, params *MergeManyParams, body MergeManyBody) (*http.Response, error)

	// DeleteOneWorkspaceMember request
	DeleteOneWorkspaceMember(ctx context.Context, id string, params *DeleteOneParams) (*http.Response, error)

	// FindOneWorkspaceMember request
	FindOneWorkspaceMember(ctx context.Context, id string, params FindOneParams) (*http.Response, error)

	UpdateOneWorkspaceMember(ctx context.Context, id string, params *UpdateOneParams, body NewWorkspaceMember) (*http.Response, error)
}
