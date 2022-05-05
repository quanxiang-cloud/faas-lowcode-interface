package user

import "context"

type GetProfileReq struct{}

type GetProfileResp struct{}

type User interface {
	// GetProfile
	GetProfile(ctx context.Context, req *GetProfileReq) (*GetProfileResp, error)
}
