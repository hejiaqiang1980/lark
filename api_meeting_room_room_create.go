// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateRoom 该接口用于创建会议室。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uITNwYjLyUDM24iM1AjN
func (r *MeetingRoomAPI) CreateRoom(ctx context.Context, request *CreateRoomReq, options ...MethodOptionFunc) (*CreateRoomResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomCreateRoom != nil {
		return r.cli.mock.mockMeetingRoomCreateRoom(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/room/create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createRoomResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "CreateRoom", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

func (r *Mock) MockMeetingRoomCreateRoom(f func(ctx context.Context, request *CreateRoomReq, options ...MethodOptionFunc) (*CreateRoomResp, *Response, error)) {
	r.mockMeetingRoomCreateRoom = f
}

func (r *Mock) UnMockMeetingRoomCreateRoom() {
	r.mockMeetingRoomCreateRoom = nil
}

type CreateRoomReq struct {
	BuildingID   string  `json:"building_id,omitempty"`    // 会议室所在的建筑ID
	Floor        string  `json:"floor,omitempty"`          // 会议室所在的建筑楼层
	Name         string  `json:"name,omitempty"`           // 会议室名称
	Capacity     int     `json:"capacity,omitempty"`       // 容量
	IsDisabled   bool    `json:"is_disabled,omitempty"`    // 是否禁用
	CustomRoomID *string `json:"custom_room_id,omitempty"` // 租户自定义会议室ID
}

type createRoomResp struct {
	Code int             `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *CreateRoomResp `json:"data,omitempty"` // 返回业务信息
}

type CreateRoomResp struct {
	RoomID string `json:"room_id,omitempty"` // 成功创建的会议室ID
}
