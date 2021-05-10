package lark

import (
	"context"
)

// DeletePublicMailboxMember 删除公共邮箱单个成员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/delete
func (r *MailAPI) DeletePublicMailboxMember(ctx context.Context, request *DeletePublicMailboxMemberReq) (*DeletePublicMailboxMemberResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deletePublicMailboxMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "DeletePublicMailboxMember", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeletePublicMailboxMemberReq struct {
	PublicMailboxID string `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值："xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
	MemberID        string `path:"member_id" json:"-"`         // 公共邮箱内成员唯一标识, 示例值："xxxxxxxxxxxxxxx"
}

type deletePublicMailboxMemberResp struct {
	Code int                            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *DeletePublicMailboxMemberResp `json:"data,omitempty"`
}

type DeletePublicMailboxMemberResp struct{}
