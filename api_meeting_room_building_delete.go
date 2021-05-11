// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteBuilding 该接口用于删除建筑物（办公大楼）。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzMxYjLzMTM24yMzEjN
func (r *MeetingRoomAPI) DeleteBuilding(ctx context.Context, request *DeleteBuildingReq, options ...MethodOptionFunc) (*DeleteBuildingResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomDeleteBuilding != nil {
		return r.cli.mock.mockMeetingRoomDeleteBuilding(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/building/delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteBuildingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "DeleteBuilding", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

func (r *Mock) MockMeetingRoomDeleteBuilding(f func(ctx context.Context, request *DeleteBuildingReq, options ...MethodOptionFunc) (*DeleteBuildingResp, *Response, error)) {
	r.mockMeetingRoomDeleteBuilding = f
}

func (r *Mock) UnMockMeetingRoomDeleteBuilding() {
	r.mockMeetingRoomDeleteBuilding = nil
}

type DeleteBuildingReq struct {
	BuildingID string `json:"building_id,omitempty"` // 要删除的建筑ID
}

type deleteBuildingResp struct {
	Code int                 `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *DeleteBuildingResp `json:"data,omitempty"`
}

type DeleteBuildingResp struct{}
