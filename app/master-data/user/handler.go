package user

import (
	"context"
	user_proto "golang-grpc/app/master-data/user/proto"

	"go-micro.dev/v4"
)

type handler struct {
	usecase *usecase
}

var userHandler *handler

func Register(service micro.Service) {
	userHandler = &handler{
		usecase: NewUserUsecase(),
	}
	user_proto.RegisterUserServiceHandler(service.Server(), userHandler)
}
func (h *handler) GetAllUser(ctx context.Context, req *user_proto.GetAllUserRequest, res *user_proto.GetAllUserResponse) error {
	return h.usecase.GetAllUser(ctx, req, res)
}
func (h *handler) GetOneUser(ctx context.Context, req *user_proto.GetOneUserRequest, res *user_proto.GetOneUserResponse) error {
	return h.usecase.GetOneUser(ctx, req, res)
}
