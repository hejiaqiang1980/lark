package lark

import (
	"context"
)

// GetPublicMailboxMemberList 分页批量获取公共邮箱成员列表
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/list
func (r *MailAPI) GetPublicMailboxMemberList(ctx context.Context, request *GetPublicMailboxMemberListReq) (*GetPublicMailboxMemberListResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getPublicMailboxMemberListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "GetPublicMailboxMemberList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetPublicMailboxMemberListReq struct {
	UserIDType      *IDType `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	PageToken       *string `query:"page_token" json:"-"`       // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："xxx"
	PageSize        *int    `query:"page_size" json:"-"`        // 分页大小, 示例值：10, 最大值：`200`
	PublicMailboxID string  `path:"public_mailbox_id" json:"-"` // The unique ID or email address of a public mailbox, 示例值："xxxxxxxxxxxxxxx or test_public_mailbox@xxx.xx"
}

type getPublicMailboxMemberListResp struct {
	Code int                             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *GetPublicMailboxMemberListResp `json:"data,omitempty"` //
}

type GetPublicMailboxMemberListResp struct {
	HasMore   bool                                  `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*GetPublicMailboxMemberListRespItem `json:"items,omitempty"`      // 公共邮箱成员列表
}

type GetPublicMailboxMemberListRespItem struct {
	MemberID string `json:"member_id,omitempty"` // 公共邮箱内成员唯一标识
	UserID   string `json:"user_id,omitempty"`   // 租户内用户的唯一标识（当成员类型是USER时有值）
	Type     string `json:"type,omitempty"`      // 成员类型, 可选值有: `USER`：内部用户
}
