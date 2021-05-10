package lark

import (
	"context"
)

// ClearPublicMailboxMember 删除公共邮箱所有成员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/clear
func (r *MailAPI) ClearPublicMailboxMember(ctx context.Context, request *ClearPublicMailboxMemberReq) (*ClearPublicMailboxMemberResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/clear",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(clearPublicMailboxMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "ClearPublicMailboxMember", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type ClearPublicMailboxMemberReq struct {
	PublicMailboxID string `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值："xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
}

type clearPublicMailboxMemberResp struct {
	Code int                           `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *ClearPublicMailboxMemberResp `json:"data,omitempty"`
}

type ClearPublicMailboxMemberResp struct{}
