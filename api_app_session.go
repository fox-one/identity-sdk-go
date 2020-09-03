package identity

import (
	"context"
	"fmt"
)

// Login Login
func (ir AppRequest) Login(ctx context.Context, req *LoginRequest) (*User, *AppError) {
	var user User
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s/v1/app/session/lgoin", ir.ServerURL), req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// TwoFactorVerify TwoFactorVerify
func (ir AppRequest) TwoFactorVerify(ctx context.Context, req *TwoFactorRequest) (*User, *AppError) {
	var user User
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s/v1/app/session/two_factor", ir.ServerURL), req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
