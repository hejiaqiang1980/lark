// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

type Mock struct {
	mockGetTenantAccessToken                           func(ctx context.Context) (*TokenExpire, *Response, error)
	mockContactCreateUser                              func(ctx context.Context, request *CreateUserReq, options ...MethodOptionFunc) (*CreateUserResp, *Response, error)
	mockContactDeleteUser                              func(ctx context.Context, request *DeleteUserReq, options ...MethodOptionFunc) (*DeleteUserResp, *Response, error)
	mockContactGetUser                                 func(ctx context.Context, request *GetUserReq, options ...MethodOptionFunc) (*GetUserResp, *Response, error)
	mockContactGetUserList                             func(ctx context.Context, request *GetUserListReq, options ...MethodOptionFunc) (*GetUserListResp, *Response, error)
	mockContactUpdateUserPatch                         func(ctx context.Context, request *UpdateUserPatchReq, options ...MethodOptionFunc) (*UpdateUserPatchResp, *Response, error)
	mockContactUpdateUser                              func(ctx context.Context, request *UpdateUserReq, options ...MethodOptionFunc) (*UpdateUserResp, *Response, error)
	mockContactCreateDepartment                        func(ctx context.Context, request *CreateDepartmentReq, options ...MethodOptionFunc) (*CreateDepartmentResp, *Response, error)
	mockContactGetDepartment                           func(ctx context.Context, request *GetDepartmentReq, options ...MethodOptionFunc) (*GetDepartmentResp, *Response, error)
	mockContactGetDepartmentList                       func(ctx context.Context, request *GetDepartmentListReq, options ...MethodOptionFunc) (*GetDepartmentListResp, *Response, error)
	mockContactGetParentDepartment                     func(ctx context.Context, request *GetParentDepartmentReq, options ...MethodOptionFunc) (*GetParentDepartmentResp, *Response, error)
	mockContactSearchDepartment                        func(ctx context.Context, request *SearchDepartmentReq, options ...MethodOptionFunc) (*SearchDepartmentResp, *Response, error)
	mockContactUpdateDepartmentPatch                   func(ctx context.Context, request *UpdateDepartmentPatchReq, options ...MethodOptionFunc) (*UpdateDepartmentPatchResp, *Response, error)
	mockContactUpdateDepartment                        func(ctx context.Context, request *UpdateDepartmentReq, options ...MethodOptionFunc) (*UpdateDepartmentResp, *Response, error)
	mockContactDeleteDepartment                        func(ctx context.Context, request *DeleteDepartmentReq, options ...MethodOptionFunc) (*DeleteDepartmentResp, *Response, error)
	mockMessageSendRawMessage                          func(ctx context.Context, request *SendRawMessageReq, options ...MethodOptionFunc) (*SendRawMessageResp, *Response, error)
	mockMessageReplyRawMessage                         func(ctx context.Context, request *ReplyRawMessageReq, options ...MethodOptionFunc) (*ReplyRawMessageResp, *Response, error)
	mockMessageDeleteMessage                           func(ctx context.Context, request *DeleteMessageReq, options ...MethodOptionFunc) (*DeleteMessageResp, *Response, error)
	mockMessageUpdateMessage                           func(ctx context.Context, request *UpdateMessageReq, options ...MethodOptionFunc) (*UpdateMessageResp, *Response, error)
	mockMessageGetMessageReadUserList                  func(ctx context.Context, request *GetMessageReadUserListReq, options ...MethodOptionFunc) (*GetMessageReadUserListResp, *Response, error)
	mockMessageGetMessageList                          func(ctx context.Context, request *GetMessageListReq, options ...MethodOptionFunc) (*GetMessageListResp, *Response, error)
	mockMessageGetMessageFile                          func(ctx context.Context, request *GetMessageFileReq, options ...MethodOptionFunc) (*GetMessageFileResp, *Response, error)
	mockMessageGetMessage                              func(ctx context.Context, request *GetMessageReq, options ...MethodOptionFunc) (*GetMessageResp, *Response, error)
	mockChatCreateChat                                 func(ctx context.Context, request *CreateChatReq, options ...MethodOptionFunc) (*CreateChatResp, *Response, error)
	mockChatGetChat                                    func(ctx context.Context, request *GetChatReq, options ...MethodOptionFunc) (*GetChatResp, *Response, error)
	mockChatUpdateChat                                 func(ctx context.Context, request *UpdateChatReq, options ...MethodOptionFunc) (*UpdateChatResp, *Response, error)
	mockChatDeleteChat                                 func(ctx context.Context, request *DeleteChatReq, options ...MethodOptionFunc) (*DeleteChatResp, *Response, error)
	mockChatGetChatListOfSelf                          func(ctx context.Context, request *GetChatListOfSelfReq, options ...MethodOptionFunc) (*GetChatListOfSelfResp, *Response, error)
	mockChatSearchChat                                 func(ctx context.Context, request *SearchChatReq, options ...MethodOptionFunc) (*SearchChatResp, *Response, error)
	mockChatGetMemberList                              func(ctx context.Context, request *GetMemberListReq, options ...MethodOptionFunc) (*GetMemberListResp, *Response, error)
	mockChatIsInChat                                   func(ctx context.Context, request *IsInChatReq, options ...MethodOptionFunc) (*IsInChatResp, *Response, error)
	mockChatAddMember                                  func(ctx context.Context, request *AddMemberReq, options ...MethodOptionFunc) (*AddMemberResp, *Response, error)
	mockChatDeleteMember                               func(ctx context.Context, request *DeleteMemberReq, options ...MethodOptionFunc) (*DeleteMemberResp, *Response, error)
	mockChatGetAnnouncement                            func(ctx context.Context, request *GetAnnouncementReq, options ...MethodOptionFunc) (*GetAnnouncementResp, *Response, error)
	mockChatUpdateAnnouncement                         func(ctx context.Context, request *UpdateAnnouncementReq, options ...MethodOptionFunc) (*UpdateAnnouncementResp, *Response, error)
	mockBotGetBotInfo                                  func(ctx context.Context, request *GetBotInfoReq, options ...MethodOptionFunc) (*GetBotInfoResp, *Response, error)
	mockCalendarCreateCalendarACL                      func(ctx context.Context, request *CreateCalendarACLReq, options ...MethodOptionFunc) (*CreateCalendarACLResp, *Response, error)
	mockCalendarDeleteCalendarACL                      func(ctx context.Context, request *DeleteCalendarACLReq, options ...MethodOptionFunc) (*DeleteCalendarACLResp, *Response, error)
	mockCalendarGetCalendarACLList                     func(ctx context.Context, request *GetCalendarACLListReq, options ...MethodOptionFunc) (*GetCalendarACLListResp, *Response, error)
	mockCalendarSubscribeCalendarACL                   func(ctx context.Context, request *SubscribeCalendarACLReq, options ...MethodOptionFunc) (*SubscribeCalendarACLResp, *Response, error)
	mockCalendarCreateCalendar                         func(ctx context.Context, request *CreateCalendarReq, options ...MethodOptionFunc) (*CreateCalendarResp, *Response, error)
	mockCalendarDeleteCalendar                         func(ctx context.Context, request *DeleteCalendarReq, options ...MethodOptionFunc) (*DeleteCalendarResp, *Response, error)
	mockCalendarGetCalendar                            func(ctx context.Context, request *GetCalendarReq, options ...MethodOptionFunc) (*GetCalendarResp, *Response, error)
	mockCalendarGetCalendarList                        func(ctx context.Context, request *GetCalendarListReq, options ...MethodOptionFunc) (*GetCalendarListResp, *Response, error)
	mockCalendarUpdateCalendar                         func(ctx context.Context, request *UpdateCalendarReq, options ...MethodOptionFunc) (*UpdateCalendarResp, *Response, error)
	mockCalendarSearchCalendar                         func(ctx context.Context, request *SearchCalendarReq, options ...MethodOptionFunc) (*SearchCalendarResp, *Response, error)
	mockCalendarSubscribeCalendar                      func(ctx context.Context, request *SubscribeCalendarReq, options ...MethodOptionFunc) (*SubscribeCalendarResp, *Response, error)
	mockCalendarUnsubscribeCalendar                    func(ctx context.Context, request *UnsubscribeCalendarReq, options ...MethodOptionFunc) (*UnsubscribeCalendarResp, *Response, error)
	mockCalendarCreateCalendarEvent                    func(ctx context.Context, request *CreateCalendarEventReq, options ...MethodOptionFunc) (*CreateCalendarEventResp, *Response, error)
	mockCalendarDeleteCalendarEvent                    func(ctx context.Context, request *DeleteCalendarEventReq, options ...MethodOptionFunc) (*DeleteCalendarEventResp, *Response, error)
	mockCalendarGetCalendarEvent                       func(ctx context.Context, request *GetCalendarEventReq, options ...MethodOptionFunc) (*GetCalendarEventResp, *Response, error)
	mockCalendarGetCalendarEventList                   func(ctx context.Context, request *GetCalendarEventListReq, options ...MethodOptionFunc) (*GetCalendarEventListResp, *Response, error)
	mockCalendarUpdateCalendarEvent                    func(ctx context.Context, request *UpdateCalendarEventReq, options ...MethodOptionFunc) (*UpdateCalendarEventResp, *Response, error)
	mockCalendarSearchCalendarEvent                    func(ctx context.Context, request *SearchCalendarEventReq, options ...MethodOptionFunc) (*SearchCalendarEventResp, *Response, error)
	mockCalendarSubscribeCalendarEvent                 func(ctx context.Context, request *SubscribeCalendarEventReq, options ...MethodOptionFunc) (*SubscribeCalendarEventResp, *Response, error)
	mockCalendarCreateCalendarEventAttendee            func(ctx context.Context, request *CreateCalendarEventAttendeeReq, options ...MethodOptionFunc) (*CreateCalendarEventAttendeeResp, *Response, error)
	mockCalendarGetCalendarEventAttendeeList           func(ctx context.Context, request *GetCalendarEventAttendeeListReq, options ...MethodOptionFunc) (*GetCalendarEventAttendeeListResp, *Response, error)
	mockCalendarDeleteCalendarEventAttendee            func(ctx context.Context, request *DeleteCalendarEventAttendeeReq, options ...MethodOptionFunc) (*DeleteCalendarEventAttendeeResp, *Response, error)
	mockCalendarGetCalendarEventAttendeeChatMemberList func(ctx context.Context, request *GetCalendarEventAttendeeChatMemberListReq, options ...MethodOptionFunc) (*GetCalendarEventAttendeeChatMemberListResp, *Response, error)
	mockCalendarGetCalendarFreeBusyList                func(ctx context.Context, request *GetCalendarFreeBusyListReq, options ...MethodOptionFunc) (*GetCalendarFreeBusyListResp, *Response, error)
	mockCalendarCreateCalendarTimeoffEvent             func(ctx context.Context, request *CreateCalendarTimeoffEventReq, options ...MethodOptionFunc) (*CreateCalendarTimeoffEventResp, *Response, error)
	mockCalendarDeleteCalendarTimeoffEvent             func(ctx context.Context, request *DeleteCalendarTimeoffEventReq, options ...MethodOptionFunc) (*DeleteCalendarTimeoffEventResp, *Response, error)
	mockCalendarGenerateCaldavConf                     func(ctx context.Context, request *GenerateCaldavConfReq, options ...MethodOptionFunc) (*GenerateCaldavConfResp, *Response, error)
	mockMeetingRoomBatchGetSummary                     func(ctx context.Context, request *BatchGetSummaryReq, options ...MethodOptionFunc) (*BatchGetSummaryResp, *Response, error)
	mockMeetingRoomGetBuildingList                     func(ctx context.Context, request *GetBuildingListReq, options ...MethodOptionFunc) (*GetBuildingListResp, *Response, error)
	mockMeetingRoomBatchGetBuilding                    func(ctx context.Context, request *BatchGetBuildingReq, options ...MethodOptionFunc) (*BatchGetBuildingResp, *Response, error)
	mockMeetingRoomGetRoomList                         func(ctx context.Context, request *GetRoomListReq, options ...MethodOptionFunc) (*GetRoomListResp, *Response, error)
	mockMeetingRoomBatchGetRoom                        func(ctx context.Context, request *BatchGetRoomReq, options ...MethodOptionFunc) (*BatchGetRoomResp, *Response, error)
	mockMeetingRoomBatchGetFreebusy                    func(ctx context.Context, request *BatchGetFreebusyReq, options ...MethodOptionFunc) (*BatchGetFreebusyResp, *Response, error)
	mockMeetingRoomReplyInstance                       func(ctx context.Context, request *ReplyInstanceReq, options ...MethodOptionFunc) (*ReplyInstanceResp, *Response, error)
	mockMeetingRoomCreateBuilding                      func(ctx context.Context, request *CreateBuildingReq, options ...MethodOptionFunc) (*CreateBuildingResp, *Response, error)
	mockMeetingRoomUpdateBuilding                      func(ctx context.Context, request *UpdateBuildingReq, options ...MethodOptionFunc) (*UpdateBuildingResp, *Response, error)
	mockMeetingRoomDeleteBuilding                      func(ctx context.Context, request *DeleteBuildingReq, options ...MethodOptionFunc) (*DeleteBuildingResp, *Response, error)
	mockMeetingRoomBatchGetBuildingID                  func(ctx context.Context, request *BatchGetBuildingIDReq, options ...MethodOptionFunc) (*BatchGetBuildingIDResp, *Response, error)
	mockMeetingRoomCreateRoom                          func(ctx context.Context, request *CreateRoomReq, options ...MethodOptionFunc) (*CreateRoomResp, *Response, error)
	mockMeetingRoomUpdateRoom                          func(ctx context.Context, request *UpdateRoomReq, options ...MethodOptionFunc) (*UpdateRoomResp, *Response, error)
	mockMeetingRoomDeleteRoom                          func(ctx context.Context, request *DeleteRoomReq, options ...MethodOptionFunc) (*DeleteRoomResp, *Response, error)
	mockMeetingRoomBatchGetRoomID                      func(ctx context.Context, request *BatchGetRoomIDReq, options ...MethodOptionFunc) (*BatchGetRoomIDResp, *Response, error)
	mockMeetingRoomGetCountryList                      func(ctx context.Context, request *GetCountryListReq, options ...MethodOptionFunc) (*GetCountryListResp, *Response, error)
	mockMeetingRoomGetDistrictList                     func(ctx context.Context, request *GetDistrictListReq, options ...MethodOptionFunc) (*GetDistrictListResp, *Response, error)
	mockVCApplyReserve                                 func(ctx context.Context, request *ApplyReserveReq, options ...MethodOptionFunc) (*ApplyReserveResp, *Response, error)
	mockVCUpdateReserve                                func(ctx context.Context, request *UpdateReserveReq, options ...MethodOptionFunc) (*UpdateReserveResp, *Response, error)
	mockVCDeleteReserve                                func(ctx context.Context, request *DeleteReserveReq, options ...MethodOptionFunc) (*DeleteReserveResp, *Response, error)
	mockVCGetReserveActiveMeeting                      func(ctx context.Context, request *GetReserveActiveMeetingReq, options ...MethodOptionFunc) (*GetReserveActiveMeetingResp, *Response, error)
	mockVCGetMeeting                                   func(ctx context.Context, request *GetMeetingReq, options ...MethodOptionFunc) (*GetMeetingResp, *Response, error)
	mockVCInviteMeeting                                func(ctx context.Context, request *InviteMeetingReq, options ...MethodOptionFunc) (*InviteMeetingResp, *Response, error)
	mockVCSetHostMeeting                               func(ctx context.Context, request *SetHostMeetingReq, options ...MethodOptionFunc) (*SetHostMeetingResp, *Response, error)
	mockVCEndMeeting                                   func(ctx context.Context, request *EndMeetingReq, options ...MethodOptionFunc) (*EndMeetingResp, *Response, error)
	mockVCStartMeetingRecording                        func(ctx context.Context, request *StartMeetingRecordingReq, options ...MethodOptionFunc) (*StartMeetingRecordingResp, *Response, error)
	mockVCStopMeetingRecording                         func(ctx context.Context, request *StopMeetingRecordingReq, options ...MethodOptionFunc) (*StopMeetingRecordingResp, *Response, error)
	mockVCGetMeetingRecording                          func(ctx context.Context, request *GetMeetingRecordingReq, options ...MethodOptionFunc) (*GetMeetingRecordingResp, *Response, error)
	mockVCSetPermissionMeetingRecording                func(ctx context.Context, request *SetPermissionMeetingRecordingReq, options ...MethodOptionFunc) (*SetPermissionMeetingRecordingResp, *Response, error)
	mockVCGetDailyReport                               func(ctx context.Context, request *GetDailyReportReq, options ...MethodOptionFunc) (*GetDailyReportResp, *Response, error)
	mockVCGetTopUserReport                             func(ctx context.Context, request *GetTopUserReportReq, options ...MethodOptionFunc) (*GetTopUserReportResp, *Response, error)
	mockVCQueryRoomConfig                              func(ctx context.Context, request *QueryRoomConfigReq, options ...MethodOptionFunc) (*QueryRoomConfigResp, *Response, error)
	mockVCSetRoomConfig                                func(ctx context.Context, request *SetRoomConfigReq, options ...MethodOptionFunc) (*SetRoomConfigResp, *Response, error)
	mockMailCreateMailGroup                            func(ctx context.Context, request *CreateMailGroupReq, options ...MethodOptionFunc) (*CreateMailGroupResp, *Response, error)
	mockMailGetMailGroup                               func(ctx context.Context, request *GetMailGroupReq, options ...MethodOptionFunc) (*GetMailGroupResp, *Response, error)
	mockMailGetMailGroupList                           func(ctx context.Context, request *GetMailGroupListReq, options ...MethodOptionFunc) (*GetMailGroupListResp, *Response, error)
	mockMailUpdateMailGroupPatch                       func(ctx context.Context, request *UpdateMailGroupPatchReq, options ...MethodOptionFunc) (*UpdateMailGroupPatchResp, *Response, error)
	mockMailUpdateMailGroup                            func(ctx context.Context, request *UpdateMailGroupReq, options ...MethodOptionFunc) (*UpdateMailGroupResp, *Response, error)
	mockMailDeleteMailGroup                            func(ctx context.Context, request *DeleteMailGroupReq, options ...MethodOptionFunc) (*DeleteMailGroupResp, *Response, error)
	mockMailCreateMailGroupMember                      func(ctx context.Context, request *CreateMailGroupMemberReq, options ...MethodOptionFunc) (*CreateMailGroupMemberResp, *Response, error)
	mockMailGetMailGroupMember                         func(ctx context.Context, request *GetMailGroupMemberReq, options ...MethodOptionFunc) (*GetMailGroupMemberResp, *Response, error)
	mockMailGetMailGroupMemberList                     func(ctx context.Context, request *GetMailGroupMemberListReq, options ...MethodOptionFunc) (*GetMailGroupMemberListResp, *Response, error)
	mockMailDeleteMailGroupMember                      func(ctx context.Context, request *DeleteMailGroupMemberReq, options ...MethodOptionFunc) (*DeleteMailGroupMemberResp, *Response, error)
	mockMailCreateMailGroupPermissionMember            func(ctx context.Context, request *CreateMailGroupPermissionMemberReq, options ...MethodOptionFunc) (*CreateMailGroupPermissionMemberResp, *Response, error)
	mockMailGetMailGroupPermissionMember               func(ctx context.Context, request *GetMailGroupPermissionMemberReq, options ...MethodOptionFunc) (*GetMailGroupPermissionMemberResp, *Response, error)
	mockMailGetMailGroupPermissionMemberList           func(ctx context.Context, request *GetMailGroupPermissionMemberListReq, options ...MethodOptionFunc) (*GetMailGroupPermissionMemberListResp, *Response, error)
	mockMailDeleteMailGroupPermissionMember            func(ctx context.Context, request *DeleteMailGroupPermissionMemberReq, options ...MethodOptionFunc) (*DeleteMailGroupPermissionMemberResp, *Response, error)
	mockMailCreatePublicMailbox                        func(ctx context.Context, request *CreatePublicMailboxReq, options ...MethodOptionFunc) (*CreatePublicMailboxResp, *Response, error)
	mockMailGetPublicMailbox                           func(ctx context.Context, request *GetPublicMailboxReq, options ...MethodOptionFunc) (*GetPublicMailboxResp, *Response, error)
	mockMailGetPublicMailboxList                       func(ctx context.Context, request *GetPublicMailboxListReq, options ...MethodOptionFunc) (*GetPublicMailboxListResp, *Response, error)
	mockMailUpdatePublicMailboxPatch                   func(ctx context.Context, request *UpdatePublicMailboxPatchReq, options ...MethodOptionFunc) (*UpdatePublicMailboxPatchResp, *Response, error)
	mockMailUpdatePublicMailbox                        func(ctx context.Context, request *UpdatePublicMailboxReq, options ...MethodOptionFunc) (*UpdatePublicMailboxResp, *Response, error)
	mockMailCreatePublicMailboxMember                  func(ctx context.Context, request *CreatePublicMailboxMemberReq, options ...MethodOptionFunc) (*CreatePublicMailboxMemberResp, *Response, error)
	mockMailGetPublicMailboxMember                     func(ctx context.Context, request *GetPublicMailboxMemberReq, options ...MethodOptionFunc) (*GetPublicMailboxMemberResp, *Response, error)
	mockMailGetPublicMailboxMemberList                 func(ctx context.Context, request *GetPublicMailboxMemberListReq, options ...MethodOptionFunc) (*GetPublicMailboxMemberListResp, *Response, error)
	mockMailDeletePublicMailboxMember                  func(ctx context.Context, request *DeletePublicMailboxMemberReq, options ...MethodOptionFunc) (*DeletePublicMailboxMemberResp, *Response, error)
	mockMailClearPublicMailboxMember                   func(ctx context.Context, request *ClearPublicMailboxMemberReq, options ...MethodOptionFunc) (*ClearPublicMailboxMemberResp, *Response, error)
	mockApprovalGetInstanceList                        func(ctx context.Context, request *GetInstanceListReq, options ...MethodOptionFunc) (*GetInstanceListResp, *Response, error)
	mockHelpdeskStartService                           func(ctx context.Context, request *StartServiceReq, options ...MethodOptionFunc) (*StartServiceResp, *Response, error)
	mockHelpdeskGetTicket                              func(ctx context.Context, request *GetTicketReq, options ...MethodOptionFunc) (*GetTicketResp, *Response, error)
	mockHelpdeskUpdateTicket                           func(ctx context.Context, request *UpdateTicketReq, options ...MethodOptionFunc) (*UpdateTicketResp, *Response, error)
	mockHelpdeskGetTicketList                          func(ctx context.Context, request *GetTicketListReq, options ...MethodOptionFunc) (*GetTicketListResp, *Response, error)
	mockHelpdeskDownloadTicketImage                    func(ctx context.Context, request *DownloadTicketImageReq, options ...MethodOptionFunc) (*DownloadTicketImageResp, *Response, error)
	mockHelpdeskGetTicketMessageList                   func(ctx context.Context, request *GetTicketMessageListReq, options ...MethodOptionFunc) (*GetTicketMessageListResp, *Response, error)
	mockHelpdeskSendTicketMessage                      func(ctx context.Context, request *SendTicketMessageReq, options ...MethodOptionFunc) (*SendTicketMessageResp, *Response, error)
	mockHelpdeskGetTicketCustomizedFieldList           func(ctx context.Context, request *GetTicketCustomizedFieldListReq, options ...MethodOptionFunc) (*GetTicketCustomizedFieldListResp, *Response, error)
	mockHelpdeskDeleteTicketCustomizedField            func(ctx context.Context, request *DeleteTicketCustomizedFieldReq, options ...MethodOptionFunc) (*DeleteTicketCustomizedFieldResp, *Response, error)
	mockHelpdeskUpdateTicketCustomizedField            func(ctx context.Context, request *UpdateTicketCustomizedFieldReq, options ...MethodOptionFunc) (*UpdateTicketCustomizedFieldResp, *Response, error)
	mockHelpdeskCreateTicketCustomizedField            func(ctx context.Context, request *CreateTicketCustomizedFieldReq, options ...MethodOptionFunc) (*CreateTicketCustomizedFieldResp, *Response, error)
	mockHelpdeskGetTicketCustomizedField               func(ctx context.Context, request *GetTicketCustomizedFieldReq, options ...MethodOptionFunc) (*GetTicketCustomizedFieldResp, *Response, error)
	mockHelpdeskCreateCategory                         func(ctx context.Context, request *CreateCategoryReq, options ...MethodOptionFunc) (*CreateCategoryResp, *Response, error)
	mockHelpdeskGetCategory                            func(ctx context.Context, request *GetCategoryReq, options ...MethodOptionFunc) (*GetCategoryResp, *Response, error)
	mockHelpdeskUpdateCategory                         func(ctx context.Context, request *UpdateCategoryReq, options ...MethodOptionFunc) (*UpdateCategoryResp, *Response, error)
	mockHelpdeskDeleteCategory                         func(ctx context.Context, request *DeleteCategoryReq, options ...MethodOptionFunc) (*DeleteCategoryResp, *Response, error)
	mockHelpdeskGetCategoryList                        func(ctx context.Context, request *GetCategoryListReq, options ...MethodOptionFunc) (*GetCategoryListResp, *Response, error)
	mockHelpdeskCreateFAQ                              func(ctx context.Context, request *CreateFAQReq, options ...MethodOptionFunc) (*CreateFAQResp, *Response, error)
	mockHelpdeskGetFAQ                                 func(ctx context.Context, request *GetFAQReq, options ...MethodOptionFunc) (*GetFAQResp, *Response, error)
	mockHelpdeskUpdateFAQ                              func(ctx context.Context, request *UpdateFAQReq, options ...MethodOptionFunc) (*UpdateFAQResp, *Response, error)
	mockHelpdeskDeleteFAQ                              func(ctx context.Context, request *DeleteFAQReq, options ...MethodOptionFunc) (*DeleteFAQResp, *Response, error)
	mockHelpdeskGetFAQList                             func(ctx context.Context, request *GetFAQListReq, options ...MethodOptionFunc) (*GetFAQListResp, *Response, error)
	mockHelpdeskGetFAQImage                            func(ctx context.Context, request *GetFAQImageReq, options ...MethodOptionFunc) (*GetFAQImageResp, *Response, error)
	mockHelpdeskSearchFAQ                              func(ctx context.Context, request *SearchFAQReq, options ...MethodOptionFunc) (*SearchFAQResp, *Response, error)
	mockAdminGetAdminDeptStats                         func(ctx context.Context, request *GetAdminDeptStatsReq, options ...MethodOptionFunc) (*GetAdminDeptStatsResp, *Response, error)
	mockAdminGetAdminUserStats                         func(ctx context.Context, request *GetAdminUserStatsReq, options ...MethodOptionFunc) (*GetAdminUserStatsResp, *Response, error)
	mockHumanAuthGetFaceVerifyAuthResult               func(ctx context.Context, request *GetFaceVerifyAuthResultReq, options ...MethodOptionFunc) (*GetFaceVerifyAuthResultResp, *Response, error)
	mockHumanAuthUploadFaceVerifyImage                 func(ctx context.Context, request *UploadFaceVerifyImageReq, options ...MethodOptionFunc) (*UploadFaceVerifyImageResp, *Response, error)
	mockHumanAuthCropFaceVerifyImage                   func(ctx context.Context, request *CropFaceVerifyImageReq, options ...MethodOptionFunc) (*CropFaceVerifyImageResp, *Response, error)
	mockHumanAuthCreateIdentity                        func(ctx context.Context, request *CreateIdentityReq, options ...MethodOptionFunc) (*CreateIdentityResp, *Response, error)
	mockAIRecognizeBasicImage                          func(ctx context.Context, request *RecognizeBasicImageReq, options ...MethodOptionFunc) (*RecognizeBasicImageResp, *Response, error)
	mockAIRecognizeSpeechStream                        func(ctx context.Context, request *RecognizeSpeechStreamReq, options ...MethodOptionFunc) (*RecognizeSpeechStreamResp, *Response, error)
	mockAIRecognizeSpeechFile                          func(ctx context.Context, request *RecognizeSpeechFileReq, options ...MethodOptionFunc) (*RecognizeSpeechFileResp, *Response, error)
	mockAITranslateText                                func(ctx context.Context, request *TranslateTextReq, options ...MethodOptionFunc) (*TranslateTextResp, *Response, error)
	mockAIDetectTextLanguage                           func(ctx context.Context, request *DetectTextLanguageReq, options ...MethodOptionFunc) (*DetectTextLanguageResp, *Response, error)
	mockAttendanceUpdateUserSettings                   func(ctx context.Context, request *UpdateUserSettingsReq, options ...MethodOptionFunc) (*UpdateUserSettingsResp, *Response, error)
	mockAttendanceUploadAttendanceFile                 func(ctx context.Context, request *UploadAttendanceFileReq, options ...MethodOptionFunc) (*UploadAttendanceFileResp, *Response, error)
	mockAttendanceCreateUpdateGroup                    func(ctx context.Context, request *CreateUpdateGroupReq, options ...MethodOptionFunc) (*CreateUpdateGroupResp, *Response, error)
	mockAttendanceDeleteGroup                          func(ctx context.Context, request *DeleteGroupReq, options ...MethodOptionFunc) (*DeleteGroupResp, *Response, error)
	mockAttendanceGetGroup                             func(ctx context.Context, request *GetGroupReq, options ...MethodOptionFunc) (*GetGroupResp, *Response, error)
	mockAttendanceCreateShift                          func(ctx context.Context, request *CreateShiftReq, options ...MethodOptionFunc) (*CreateShiftResp, *Response, error)
	mockAttendanceDeleteShift                          func(ctx context.Context, request *DeleteShiftReq, options ...MethodOptionFunc) (*DeleteShiftResp, *Response, error)
	mockAttendanceGetShiftByID                         func(ctx context.Context, request *GetShiftByIDReq, options ...MethodOptionFunc) (*GetShiftByIDResp, *Response, error)
	mockAttendanceGetShiftByName                       func(ctx context.Context, request *GetShiftByNameReq, options ...MethodOptionFunc) (*GetShiftByNameResp, *Response, error)
	mockAttendanceGetStatisticsData                    func(ctx context.Context, request *GetStatisticsDataReq, options ...MethodOptionFunc) (*GetStatisticsDataResp, *Response, error)
	mockAttendanceGetStatisticsHeader                  func(ctx context.Context, request *GetStatisticsHeaderReq, options ...MethodOptionFunc) (*GetStatisticsHeaderResp, *Response, error)
	mockAttendanceUpdateUserStatisticsSettings         func(ctx context.Context, request *UpdateUserStatisticsSettingsReq, options ...MethodOptionFunc) (*UpdateUserStatisticsSettingsResp, *Response, error)
	mockAttendanceGetUserStatisticsSettings            func(ctx context.Context, request *GetUserStatisticsSettingsReq, options ...MethodOptionFunc) (*GetUserStatisticsSettingsResp, *Response, error)
	mockAttendanceGetUserDailyShift                    func(ctx context.Context, request *GetUserDailyShiftReq, options ...MethodOptionFunc) (*GetUserDailyShiftResp, *Response, error)
	mockAttendanceGetUserTask                          func(ctx context.Context, request *GetUserTaskReq, options ...MethodOptionFunc) (*GetUserTaskResp, *Response, error)
	mockAttendanceGetUserFlow                          func(ctx context.Context, request *GetUserFlowReq, options ...MethodOptionFunc) (*GetUserFlowResp, *Response, error)
	mockAttendanceBatchGetUserFlow                     func(ctx context.Context, request *BatchGetUserFlowReq, options ...MethodOptionFunc) (*BatchGetUserFlowResp, *Response, error)
	mockAttendanceBatchCreateUserFlow                  func(ctx context.Context, request *BatchCreateUserFlowReq, options ...MethodOptionFunc) (*BatchCreateUserFlowResp, *Response, error)
	mockAttendanceGetUserTaskRemedy                    func(ctx context.Context, request *GetUserTaskRemedyReq, options ...MethodOptionFunc) (*GetUserTaskRemedyResp, *Response, error)
	mockAttendanceCreateUpdateUserDailyShift           func(ctx context.Context, request *CreateUpdateUserDailyShiftReq, options ...MethodOptionFunc) (*CreateUpdateUserDailyShiftResp, *Response, error)
	mockAttendanceGetUserApproval                      func(ctx context.Context, request *GetUserApprovalReq, options ...MethodOptionFunc) (*GetUserApprovalResp, *Response, error)
	mockAttendanceCreateUserApproval                   func(ctx context.Context, request *CreateUserApprovalReq, options ...MethodOptionFunc) (*CreateUserApprovalResp, *Response, error)
	mockFileUploadImage                                func(ctx context.Context, request *UploadImageReq, options ...MethodOptionFunc) (*UploadImageResp, *Response, error)
	mockFileDownloadImage                              func(ctx context.Context, request *DownloadImageReq, options ...MethodOptionFunc) (*DownloadImageResp, *Response, error)
	mockFileUploadFile                                 func(ctx context.Context, request *UploadFileReq, options ...MethodOptionFunc) (*UploadFileResp, *Response, error)
	mockFileDownloadFile                               func(ctx context.Context, request *DownloadFileReq, options ...MethodOptionFunc) (*DownloadFileResp, *Response, error)
	mockEHRGetEmployeeList                             func(ctx context.Context, request *GetEmployeeListReq, options ...MethodOptionFunc) (*GetEmployeeListResp, *Response, error)
	mockEHRDownloadAttachments                         func(ctx context.Context, request *DownloadAttachmentsReq, options ...MethodOptionFunc) (*DownloadAttachmentsResp, *Response, error)
	mockTenantQueryTenant                              func(ctx context.Context, request *QueryTenantReq, options ...MethodOptionFunc) (*QueryTenantResp, *Response, error)
}

func (r *Lark) Mock() *Mock {
	return r.mock
}
