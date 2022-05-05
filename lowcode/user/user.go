package user

import "context"

type GetProfileReq struct {
	ID string `json:"id,omitempty"`
}

type GetProfileResp struct {
	Code int64  `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data struct {
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	} `json:"data,omitempty"`
}

type User interface {
	// GetProfile
	GetProfile(ctx context.Context, req *GetProfileReq) (*GetProfileResp, error)
}
