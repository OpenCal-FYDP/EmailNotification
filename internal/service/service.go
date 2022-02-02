package service

import (
	"context"
	"github.com/OpenCal-FYDP/EmailNotification/rpc"
)

type EmailNotificationService struct{}

func (e EmailNotificationService) CreateNotification(ctx context.Context, req *rpc.CreateNotificationReq) (*rpc.CreateNotificationRes, error) {
	panic("implement me")
}

func New() *EmailNotificationService {
	return &EmailNotificationService{}
}
