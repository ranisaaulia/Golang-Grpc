package app

import (
	"golang-grpc/app/master-data/user"
	"golang-grpc/config"

	grpc "github.com/asim/go-micro/plugins/server/grpc/v4"
	"go-micro.dev/v4"
)

type serverApp struct {
	Service micro.Service
}

var Server *serverApp

func (s *serverApp) Register() {
	user.Register(s.Service)
}

func (s *serverApp) New() {
	Server = &serverApp{
		Service: micro.NewService(
			micro.Server(grpc.NewServer()),
			micro.Name(config.Application.Name),
			micro.Version(config.Application.Version),
			micro.Address(":"+config.Application.Port),
		),
	}
}
func (s *serverApp) Run() error {
	return s.Service.Run()
}
