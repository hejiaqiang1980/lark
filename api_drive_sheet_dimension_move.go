// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// MoveSheetDimension 该接口用于移动行列，行列被移动到目标位置后，原本在目标位置的行列会对应右移或下移。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/move_dimension
func (r *DriveService) MoveSheetDimension(ctx context.Context, request *MoveSheetDimensionReq, options ...MethodOptionFunc) (*MoveSheetDimensionResp, *Response, error) {
	if r.cli.mock.mockDriveMoveSheetDimension != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#MoveSheetDimension mock enable")
		return r.cli.mock.mockDriveMoveSheetDimension(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "MoveSheetDimension",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/move_dimension",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(moveSheetDimensionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveMoveSheetDimension(f func(ctx context.Context, request *MoveSheetDimensionReq, options ...MethodOptionFunc) (*MoveSheetDimensionResp, *Response, error)) {
	r.mockDriveMoveSheetDimension = f
}

func (r *Mock) UnMockDriveMoveSheetDimension() {
	r.mockDriveMoveSheetDimension = nil
}

type MoveSheetDimensionReq struct {
	SpreadSheetToken string                       `path:"spreadsheet_token" json:"-"`  // 表格 token, 示例值："shtcnmBA\*****yGehy8"
	SheetID          string                       `path:"sheet_id" json:"-"`           // 子表 id, 示例值："0b\**12"
	Source           *MoveSheetDimensionReqSource `json:"source,omitempty"`            // 移动源位置参数
	DestinationIndex *int64                       `json:"destination_index,omitempty"` // 移动的目标位置行或者列号, 示例值：4
}

type MoveSheetDimensionReqSource struct {
	MajorDimension *string `json:"major_dimension,omitempty"` // 操作行还是列，取值：ROWS、COLUMNS, 示例值："ROWS"
	StartIndex     *int64  `json:"start_index,omitempty"`     // 起始行或者列号, 示例值：0
	EndIndex       *int64  `json:"end_index,omitempty"`       // 结束行或者列号, 示例值：1
}

type moveSheetDimensionResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *MoveSheetDimensionResp `json:"data,omitempty"`
}

type MoveSheetDimensionResp struct{}
