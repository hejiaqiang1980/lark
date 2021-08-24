// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetBitableRecordList 该接口用于列出数据表中的现有记录
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/list
func (r *BitableService) GetBitableRecordList(ctx context.Context, request *GetBitableRecordListReq, options ...MethodOptionFunc) (*GetBitableRecordListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableRecordList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableRecordList mock enable")
		return r.cli.mock.mockBitableGetBitableRecordList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableRecordList",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableRecordListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockBitableGetBitableRecordList(f func(ctx context.Context, request *GetBitableRecordListReq, options ...MethodOptionFunc) (*GetBitableRecordListResp, *Response, error)) {
	r.mockBitableGetBitableRecordList = f
}

func (r *Mock) UnMockBitableGetBitableRecordList() {
	r.mockBitableGetBitableRecordList = nil
}

type GetBitableRecordListReq struct {
	ViewID     *string `query:"view_id" json:"-"`      // 视图 id, 如filter或sort有值, view_id会被忽略, 示例值："vewqhz51lk"
	Filter     *string `query:"filter" json:"-"`       // filter, 不超过2000个字符, 不支持对带特殊字段(关联和公式)的表的使用, 示例值："AND(CurrentValue.[身高]>180, CurrentValue.[体重]>150)"
	Sort       *string `query:"sort" json:"-"`         // sort, 不超过1000字符, 不支持对带特殊字段(关联和公式)的表的使用, 示例值："["字段1 DESC","字段2 ASC"]"
	FieldNames *string `query:"field_names" json:"-"`  // field_names, 示例值："["字段1"]"
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："recn0hoyXL"
	PageSize   *int64  `query:"page_size" json:"-"`    // 分页大小, 示例值：10, 最大值：`100`
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求:  获取用户 userid
	AppToken   string  `path:"app_token" json:"-"`     // bitable app token, 示例值："bascnCMII2ORej2RItqpZZUNMIe"
	TableID    string  `path:"table_id" json:"-"`      // table id, 示例值："tblxI2tWaxP5dG7p"
}

type getBitableRecordListResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableRecordListResp `json:"data,omitempty"`
}

type GetBitableRecordListResp struct {
	HasMore   bool                            `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                          `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Total     int64                           `json:"total,omitempty"`      // 总数
	Items     []*GetBitableRecordListRespItem `json:"items,omitempty"`      // 记录信息
}

type GetBitableRecordListRespItem struct {
	RecordID string                 `json:"record_id,omitempty"` // 记录 id
	Fields   map[string]interface{} `json:"fields,omitempty"`    // 记录字段
}
