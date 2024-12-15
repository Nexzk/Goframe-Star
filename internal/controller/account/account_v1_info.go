package account

import (
	"context"
	"star/internal/logic/users"

	"star/api/account/v1"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	user, err := users.Info(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.InfoRes{
		Username: user.Username,
		Password: user.Password,
		CreateAt: user.CreatedAt,
		UpdateAt: user.UpdatedAt,
	}, nil
	return
}
