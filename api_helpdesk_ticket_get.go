// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetHelpdeskTicket 该接口用于获取单个服务台工单详情。仅支持自建应用。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/get
func (r *HelpdeskService) GetHelpdeskTicket(ctx context.Context, request *GetHelpdeskTicketReq, options ...MethodOptionFunc) (*GetHelpdeskTicketResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskTicket != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskTicket mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskTicket(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskTicket",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskTicketResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHelpdeskGetHelpdeskTicket(f func(ctx context.Context, request *GetHelpdeskTicketReq, options ...MethodOptionFunc) (*GetHelpdeskTicketResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskTicket = f
}

func (r *Mock) UnMockHelpdeskGetHelpdeskTicket() {
	r.mockHelpdeskGetHelpdeskTicket = nil
}

type GetHelpdeskTicketReq struct {
	TicketID string `path:"ticket_id" json:"-"` // ticket id, 示例值："123456"
}

type getHelpdeskTicketResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskTicketResp `json:"data,omitempty"`
}

type GetHelpdeskTicketResp struct {
	Ticket *GetHelpdeskTicketRespTicket `json:"ticket,omitempty"` // 工单详情
}

type GetHelpdeskTicketRespTicket struct {
	TicketID                   string                                        `json:"ticket_id,omitempty"`                     // 工单ID,[可以从工单列表里面取](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/list),[也可以订阅工单创建事件获取](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/events/created)
	HelpdeskID                 string                                        `json:"helpdesk_id,omitempty"`                   // 服务台ID
	Guest                      *GetHelpdeskTicketRespTicketGuest             `json:"guest,omitempty"`                         // 工单创建用户
	Stage                      int64                                         `json:"stage,omitempty"`                         // 工单阶段，1：bot，2：人工
	Status                     int64                                         `json:"status,omitempty"`                        // 工单状态，1：已创建 2: 处理中 3: 排队中 4：待定 5：待用户响应 50: 被机器人关闭 51: 被人工关闭
	Score                      int64                                         `json:"score,omitempty"`                         // 工单评分，1：不满意，2:一般，3:满意
	CreatedAt                  int64                                         `json:"created_at,omitempty"`                    // 工单创建时间
	UpdatedAt                  int64                                         `json:"updated_at,omitempty"`                    // 工单更新时间，没有值时为-1
	ClosedAt                   int64                                         `json:"closed_at,omitempty"`                     // 工单结束时间
	Agents                     []*GetHelpdeskTicketRespTicketAgent           `json:"agents,omitempty"`                        // 工单客服
	Channel                    int64                                         `json:"channel,omitempty"`                       // 工单渠道
	Solve                      int64                                         `json:"solve,omitempty"`                         // 工单是否解决 1:没解决 2:已解决
	ClosedBy                   *GetHelpdeskTicketRespTicketClosedBy          `json:"closed_by,omitempty"`                     // 关单用户ID
	Collaborators              []*GetHelpdeskTicketRespTicketCollaborator    `json:"collaborators,omitempty"`                 // 工单协作者
	CustomizedFields           []*GetHelpdeskTicketRespTicketCustomizedField `json:"customized_fields,omitempty"`             // 自定义字段列表，没有值时不设置
	AgentServiceDuration       float64                                       `json:"agent_service_duration,omitempty"`        // 客服服务时长minutes
	AgentFirstResponseDuration int64                                         `json:"agent_first_response_duration,omitempty"` // 客服第一反应时间 seconds
	BotServiceDuration         int64                                         `json:"bot_service_duration,omitempty"`          // 机器人服务时长 seconds
	AgentResolutionTime        int64                                         `json:"agent_resolution_time,omitempty"`         // 解决时长(秒)
	ActualProcessingTime       int64                                         `json:"actual_processing_time,omitempty"`        // 处理时长(秒)
}

type GetHelpdeskTicketRespTicketGuest struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Email      string `json:"email,omitempty"`      // 用户邮箱
	Department string `json:"department,omitempty"` // 所在部门名称
}

type GetHelpdeskTicketRespTicketAgent struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Email      string `json:"email,omitempty"`      // 用户邮箱
	Department string `json:"department,omitempty"` // 所在部门名称
}

type GetHelpdeskTicketRespTicketClosedBy struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Email      string `json:"email,omitempty"`      // 用户邮箱
	Department string `json:"department,omitempty"` // 所在部门名称
}

type GetHelpdeskTicketRespTicketCollaborator struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Email      string `json:"email,omitempty"`      // 用户邮箱
	Department string `json:"department,omitempty"` // 所在部门名称
}

type GetHelpdeskTicketRespTicketCustomizedField struct {
	ID          string `json:"id,omitempty"`           // 自定义字段ID
	Value       string `json:"value,omitempty"`        // 自定义字段值
	KeyName     string `json:"key_name,omitempty"`     // 键名
	DisplayName string `json:"display_name,omitempty"` // 展示名称
	Position    int64  `json:"position,omitempty"`     // 展示位置
	Required    bool   `json:"required,omitempty"`     // 是否必填
	Editable    bool   `json:"editable,omitempty"`     // 是否可修改
}
