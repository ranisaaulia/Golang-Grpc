package user

import (
	"context"
	user_proto "golang-grpc/app/master-data/user/proto"
	user_sql "golang-grpc/app/master-data/user/sqlc"
	"golang-grpc/config"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type usecase struct {
	repository *user_sql.Queries
}

func NewUserUsecase() *usecase {
	return &usecase{
		repository: user_sql.New(config.Application.DB),
	}
}
func (u *usecase) GetAllUser(ctx context.Context, req *user_proto.GetAllUserRequest, res *user_proto.GetAllUserResponse) error {
	var err error
	var dataReturn []*user_proto.UserItem
	userData, err := u.repository.SelectAllUser(ctx)
	if err != nil {
		return err
	}
	for _, item := range userData {
		itemInsert := user_proto.UserItem{}
		itemInsert.ID = item.ID
		itemInsert.Username = item.Username
		itemInsert.Fullname = item.Fullname.String
		itemInsert.IsActive = item.IsActive
		itemInsert.CreatedBy = item.CreatedBy
		itemInsert.CreatedDate = timestamppb.New(item.CreatedDate)
		itemInsert.ModifiedBy = item.ModifiedBy
		itemInsert.ModifiedDate = timestamppb.New(item.ModifiedDate)
		dataReturn = append(dataReturn, &itemInsert)
	}
	res.Data = dataReturn
	res.IsError = false
	return err
}
func (u *usecase) GetOneUser(ctx context.Context, req *user_proto.GetOneUserRequest, res *user_proto.GetOneUserResponse) error {
	var err error
	userData, err := u.repository.SelectOneUser(ctx, req.Username)
	if err != nil {
		return err
	}
	res.Data = &user_proto.UserItem{
		ID:           userData.ID,
		Username:     userData.Username,
		Password:     userData.Password,
		Fullname:     userData.Fullname.String,
		IsActive:     userData.IsActive,
		CreatedBy:    userData.CreatedBy,
		CreatedDate:  timestamppb.New(userData.CreatedDate),
		ModifiedBy:   userData.ModifiedBy,
		ModifiedDate: timestamppb.New(userData.ModifiedDate),
	}
	res.IsError = false
	return err
}
