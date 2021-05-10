package lark

import (
	"context"
)

// GetMailGroupMemberList 分页批量获取邮件组成员列表
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/list
func (r *MailAPI) GetMailGroupMemberList(ctx context.Context, request *GetMailGroupMemberListReq) (*GetMailGroupMemberListResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/members",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getMailGroupMemberListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "GetMailGroupMemberList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetMailGroupMemberListReq struct {
	UserIDType       *IDType `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	DepartmentIDType *IDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值："open_department_id", 可选值有: `department_id`：以自定义department_id来标识部门, `open_department_id`：以open_department_id来标识部门
	PageToken        *string `query:"page_token" json:"-"`         // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："xxx"
	PageSize         *int    `query:"page_size" json:"-"`          // 分页大小, 示例值：10, 最大值：`200`
	MailgroupID      string  `path:"mailgroup_id" json:"-"`        // The unique ID or email address of a mail group, 示例值："xxxxxxxxxxxxxxx or test_mail_group@xxx.xx"
}

type getMailGroupMemberListResp struct {
	Code int                         `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *GetMailGroupMemberListResp `json:"data,omitempty"` //
}

type GetMailGroupMemberListResp struct {
	HasMore   bool                              `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                            `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*GetMailGroupMemberListRespItem `json:"items,omitempty"`      // 邮件组成员列表
}

type GetMailGroupMemberListRespItem struct {
	MemberID     string `json:"member_id,omitempty"`     // 邮件组内成员唯一标识
	Email        string `json:"email,omitempty"`         // 成员邮箱地址（当成员类型是EXTERNAL_USER/MAIL_GROUP/OTHER_MEMBER时有值）
	UserID       string `json:"user_id,omitempty"`       // 租户内用户的唯一标识（当成员类型是USER时有值）
	DepartmentID string `json:"department_id,omitempty"` // 租户内部门的唯一标识（当成员类型是DEPARTMENT时有值）
	Type         string `json:"type,omitempty"`          // 成员类型, 可选值有: `USER`：内部用户, `DEPARTMENT`：部门, `COMPANY`：全员, `EXTERNAL_USER`：外部用户, `MAIL_GROUP`：邮件组, `OTHER_MEMBER`：内部成员
}
