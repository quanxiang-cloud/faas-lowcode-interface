package form

import (
	"context"
	"github.com/quanxiang-cloud/faas-lowcode-interface/lowcode/consensus"
)

type Entity map[string]interface{}

type Entities []map[string]interface{}

type ListOpt struct {
	Page  int64                  `json:"page"`
	Size  int64                  `json:"size"`
	Sort  []string               `json:"sort"`
	Query map[string]interface{} `json:"query"`
}

type Universal struct {
	TableID string `json:"tableID,omitempty"`
	AppID   string `json:"appID,omitempty"`
}
type ListReq struct {
	Universal
	ListOpt
}

type ListResp struct {
	consensus.UniversalResp
	Data struct {
		Total    int64    `json:"total,omitempty"`
		Entities Entities `json:"entities,omitempty"`
	} `json:"data,omitempty"`
}

type GetReq struct {
	Universal
	ID string `json:"id,omitempty"`
}

type GetResp struct {
	consensus.UniversalResp
	Data struct {
		Total  int64  `json:"total,omitempty"`
		Entity Entity `json:"entity,omitempty"`
	} `json:"data,omitempty"`
}

type DeleteReq struct {
	Universal
	ID string `json:"id"`
}

type DeleteResp struct {
	consensus.UniversalResp
	Data struct {
		Total int64 `json:"total,omitempty"`
	} `json:"data,omitempty"`
}

type UpdateReq struct {
	Universal
	Entity Entity `json:"entity,omitempty"`
	ID     string `json:"id,omitempty"`
}

type UpdateResp struct {
	consensus.UniversalResp
	Data struct {
		Total  int64  `json:"total,omitempty"`
		Entity Entity `json:"entity,omitempty"`
	} `json:"data,omitempty"`
}

type CreateReq struct {
	Universal
	Entity Entity `json:"entity,omitempty"`
}

type CreateRsp struct {
	consensus.UniversalResp
	Data struct {
		Total  int64  `json:"total,omitempty"`
		Entity Entity `json:"entity,omitempty"`
	} `json:"data,omitempty"`
}

type Form interface {
	ListForm(ctx context.Context, req *ListReq) (*ListResp, error)
	GetForm(ctx context.Context, req *GetReq) (*GetResp, error)
	DeleteForm(ctx context.Context, req *DeleteReq) (*DeleteResp, error)
	UpdateForm(ctx context.Context, req *UpdateReq) (*UpdateResp, error)
	CreateForm(ctx context.Context, req *CreateReq) (*CreateRsp, error)
}
