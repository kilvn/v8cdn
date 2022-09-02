// Package svc_auth
package svc_auth

import "context"

type authService struct {
    ctx context.Context
}

func Instance(ctx context.Context) *authService {
    return &authService{
        ctx,
    }
}
