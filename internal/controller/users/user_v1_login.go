package users

import (
	"context"

	"star/api/users/v1"
	"star/internal/logic/users"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := users.Login(ctx, req.Username, req.Password)
	if err != nil {
		return
	}
	return &v1.LoginRes{Token: token}, nil
}
