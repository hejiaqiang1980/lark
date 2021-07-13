// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetApplicationUsageOverview
//
// 查询应用在指定时间段内企业员工的使用概览信息。
// :::warning
// 此接口目前仅支持小程序的使用情况查询，不支持网页应用和机器人应用的使用情况查询；
// 仅支持查询自建应用，不支持查询商店应用
// :::
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uETN0YjLxUDN24SM1QjN
func (r *ApplicationService) GetApplicationUsageOverview(ctx context.Context, request *GetApplicationUsageOverviewReq, options ...MethodOptionFunc) (*GetApplicationUsageOverviewResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationUsageOverview != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#GetApplicationUsageOverview mock enable")
		return r.cli.mock.mockApplicationGetApplicationUsageOverview(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationUsageOverview",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/application/v1/app_usage_overview",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationUsageOverviewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockApplicationGetApplicationUsageOverview(f func(ctx context.Context, request *GetApplicationUsageOverviewReq, options ...MethodOptionFunc) (*GetApplicationUsageOverviewResp, *Response, error)) {
	r.mockApplicationGetApplicationUsageOverview = f
}

func (r *Mock) UnMockApplicationGetApplicationUsageOverview() {
	r.mockApplicationGetApplicationUsageOverview = nil
}

type GetApplicationUsageOverviewReq struct {
	AppID     string                                  `json:"app_id,omitempty"`     // 目标应用的 ID，支持自建应用
	Ability   string                                  `json:"ability,omitempty"`    // 应用能力，mp：小程序
	TimeStart int64                                   `json:"time_start,omitempty"` // 起始时间戳（秒），时间跨度最长支持180天
	TimeEnd   int64                                   `json:"time_end,omitempty"`   // 截止时间戳（秒），时间跨度最长支持180天
	Filters   []*GetApplicationUsageOverviewReqFilter `json:"filters,omitempty"`    // 过滤条件
}

type GetApplicationUsageOverviewReqFilter struct {
	Key   string `json:"key,omitempty"`   // 过滤字段，支持`department_id`
	Op    string `json:"op,omitempty"`    // 过滤操作，支持`in`、`=`
	Value string `json:"value,omitempty"` // 过滤字段值，多个使用英文逗号分隔
}

type getApplicationUsageOverviewResp struct {
	Code int64                            `json:"code,omitempty"` // 返回码，非0表示失败
	Msg  string                           `json:"msg,omitempty"`  // 返回码的描述
	Data *GetApplicationUsageOverviewResp `json:"data,omitempty"` // 返回的业务信息，仅code = 0时有效
}

type GetApplicationUsageOverviewResp struct {
	Item map[string]*GetApplicationUsageOverviewRespItem `json:"item,omitempty"` // 返回项
}

type GetApplicationUsageOverviewRespItem struct {
	Pv int64 `json:"pv,omitempty"` // 应用使用pv
	Uv int64 `json:"uv,omitempty"` // 应用使用uv
}
