// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateUpdateUserDailyShift
//
// 班表是用来描述考勤组内人员每天按哪个班次进行上班。目前班表支持按一个整月对一位或多位人员进行排班。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//CreateandEditShifts
func (r *AttendanceAPI) CreateUpdateUserDailyShift(ctx context.Context, request *CreateUpdateUserDailyShiftReq, options ...MethodOptionFunc) (*CreateUpdateUserDailyShiftResp, *Response, error) {
	if r.cli.mock.mockAttendanceCreateUpdateUserDailyShift != nil {
		return r.cli.mock.mockAttendanceCreateUpdateUserDailyShift(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_daily_shifts/batch_create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createUpdateUserDailyShiftResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Attendance", "CreateUpdateUserDailyShift", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

func (r *Mock) MockAttendanceCreateUpdateUserDailyShift(f func(ctx context.Context, request *CreateUpdateUserDailyShiftReq, options ...MethodOptionFunc) (*CreateUpdateUserDailyShiftResp, *Response, error)) {
	r.mockAttendanceCreateUpdateUserDailyShift = f
}

func (r *Mock) UnMockAttendanceCreateUpdateUserDailyShift() {
	r.mockAttendanceCreateUpdateUserDailyShift = nil
}

type CreateUpdateUserDailyShiftReq struct {
	EmployeeType    EmployeeType                                   `query:"employee_type" json:"-"`     // 请求体中的 user_id 的员工工号类型可用值：【employee_id（员工的 employeeId），employee_no（员工工号）】，示例值："employee_id"
	UserDailyShifts []*CreateUpdateUserDailyShiftReqUserDailyShift `json:"user_daily_shifts,omitempty"` // 班表信息列表
}

type CreateUpdateUserDailyShiftReqUserDailyShift struct {
	GroupID string `json:"group_id,omitempty"` // 考勤组 ID
	ShiftID string `json:"shift_id,omitempty"` // 班次 ID，休息为 0
	Month   int    `json:"month,omitempty"`    // 月份
	UserID  string `json:"user_id,omitempty"`  // 用户
	DayNo   int    `json:"day_no,omitempty"`   // 日期
}

type createUpdateUserDailyShiftResp struct {
	Code int                             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *CreateUpdateUserDailyShiftResp `json:"data,omitempty"` // -
}

type CreateUpdateUserDailyShiftResp struct {
	UserDailyShifts []*CreateUpdateUserDailyShiftRespUserDailyShift `json:"user_daily_shifts,omitempty"` // 班表信息列表
}

type CreateUpdateUserDailyShiftRespUserDailyShift struct {
	GroupID string `json:"group_id,omitempty"` // 考勤组 ID
	ShiftID string `json:"shift_id,omitempty"` // 班次 ID
	Month   int    `json:"month,omitempty"`    // 月份
	UserID  string `json:"user_id,omitempty"`  // 用户
	DayNo   int    `json:"day_no,omitempty"`   // 日期
}
