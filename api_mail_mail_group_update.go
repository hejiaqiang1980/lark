package lark

import (
	"context"
)

// UpdateMailGroup 更新邮件组所有信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/update
func (r *MailAPI) UpdateMailGroup(ctx context.Context, request *UpdateMailGroupReq) (*UpdateMailGroupResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(updateMailGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "UpdateMailGroup", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UpdateMailGroupReq struct {
	MailgroupID    string  `path:"mailgroup_id" json:"-"`       // 邮件组ID或者邮件组地址, 示例值："xxxxxxxxxxxxxxx 或 test_mail_group@xxx.xx"
	Email          *string `json:"email,omitempty"`             // 邮件组地址, 示例值："test_mail_group@xxx.xx"
	Name           *string `json:"name,omitempty"`              // 邮件组名称, 示例值："test mail group"
	Description    *string `json:"description,omitempty"`       // 邮件组描述, 示例值："mail group for testing"
	WhoCanSendMail *string `json:"who_can_send_mail,omitempty"` // 谁可发送邮件到此邮件组, 示例值："ALL_INTERNAL_USERS", 可选值有: `ANYONE`：任何人, `ALL_INTERNAL_USERS`：仅组织内部成员, `ALL_GROUP_MEMBERS`：仅邮件组成员, `CUSTOM_MEMBERS`：自定义成员
}

type updateMailGroupResp struct {
	Code int                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *UpdateMailGroupResp `json:"data,omitempty"` //
}

type UpdateMailGroupResp struct {
	MailgroupID             string `json:"mailgroup_id,omitempty"`               // 邮件组ID
	Email                   string `json:"email,omitempty"`                      // 邮件组地址
	Name                    string `json:"name,omitempty"`                       // 邮件组名称
	Description             string `json:"description,omitempty"`                // 邮件组描述
	DirectMembersCount      string `json:"direct_members_count,omitempty"`       // 邮件组成员数量
	IncludeExternalMember   bool   `json:"include_external_member,omitempty"`    // 是否包含外部成员
	IncludeAllCompanyMember bool   `json:"include_all_company_member,omitempty"` // 是否是全员邮件组
	WhoCanSendMail          string `json:"who_can_send_mail,omitempty"`          // 谁可发送邮件到此邮件组, 可选值有: `ANYONE`：任何人, `ALL_INTERNAL_USERS`：仅组织内部成员, `ALL_GROUP_MEMBERS`：仅邮件组成员, `CUSTOM_MEMBERS`：自定义成员
}
