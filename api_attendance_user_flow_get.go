// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetUserFlow
//
// 通过打卡记录 ID 获取用户的打卡流水记录。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//GetCardSwipeHistory
func (r *AttendanceAPI) GetUserFlow(ctx context.Context, request *GetUserFlowReq, options ...MethodOptionFunc) (*GetUserFlowResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetUserFlow != nil {
		return r.cli.mock.mockAttendanceGetUserFlow(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_flows/:user_flow_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getUserFlowResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Attendance", "GetUserFlow", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

func (r *Mock) MockAttendanceGetUserFlow(f func(ctx context.Context, request *GetUserFlowReq, options ...MethodOptionFunc) (*GetUserFlowResp, *Response, error)) {
	r.mockAttendanceGetUserFlow = f
}

func (r *Mock) UnMockAttendanceGetUserFlow() {
	r.mockAttendanceGetUserFlow = nil
}

type GetUserFlowReq struct {
	EmployeeType EmployeeType `query:"employee_type" json:"-"` // 请求体中的 user_id 和 creator_id 的员工工号类型，可用值：【employee_id（员工的 employeeId），employee_no（员工工号）】，示例值：“employee_id”
	UserFlowID   string       `path:"user_flow_id" json:"-"`   // 打卡流水记录 ID，示例值："6708236686834352397"
}

type getUserFlowResp struct {
	Code int              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *GetUserFlowResp `json:"data,omitempty"` // -
}

type GetUserFlowResp struct {
	UserID       string   `json:"user_id,omitempty"`       // 员工工号
	CreatorID    string   `json:"creator_id,omitempty"`    // 打卡记录创建者的 employee_no
	LocationName string   `json:"location_name,omitempty"` // 打卡位置名称信息
	CheckTime    string   `json:"check_time,omitempty"`    // 打卡时间，精确到秒的时间戳
	Comment      string   `json:"comment,omitempty"`       // 打卡备注
	RecordID     string   `json:"record_id,omitempty"`     // 打卡记录 ID
	Longitude    float64  `json:"longitude,omitempty"`     // 打卡经度
	Latitude     float64  `json:"latitude,omitempty"`      // 打卡纬度
	Ssid         string   `json:"ssid,omitempty"`          // 打卡 Wi-Fi 的 SSID
	Bssid        string   `json:"bssid,omitempty"`         // 打卡 Wi-Fi 的 MAC 地址
	IsField      bool     `json:"is_field,omitempty"`      // 是否为外勤打卡
	IsWifi       bool     `json:"is_wifi,omitempty"`       // 是否为 Wi-Fi 打卡
	Type         int      `json:"type,omitempty"`          // 记录生成方式，可用值：【0（用户自己打卡），1（管理员修改），2（用户补卡），3（系统自动生成），4（下班免打卡），5（考勤机打卡），6（极速打卡），7（考勤开放平台导入）】
	PhotoUrls    []string `json:"photo_urls,omitempty"`    // 打卡照片列表
}
