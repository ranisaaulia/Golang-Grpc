package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type (
	DBConfig struct {
		HOST          string `envconfig:"DB_HOST" default:"localhost"`
		PORT          string `envconfig:"DB_PORT" default:"6543"`
		USER          string `envconfig:"DB_USER" default:"user"`
		NAME          string `envconfig:"DB_NAME" default:"go_grpc"`
		PASS          string `envconfig:"DB_PASSWORD" default:"P@ssw0rd"`
		MAX_OPEN_CONN int    `envconfig:"DB_MAX_OPEN_CONN" default:"100"`
		MAX_IDLE_CONN int    `envconfig:"DB_MAX_IDDLE_CONN" default:"10"`
	}
	Log struct {
		Level string `envconfig:"LOG_LEVEL" default:"debug"`
		DSN   string `envconfig:"LOG_DSN" default:""`
	}
	GrpcClient struct {
		MasterData string `envconfig:"GRPC_CLIENT" default:"localhost:8081"`
		Context    *context.Context
		Cancel     context.CancelFunc
		GrpcOpts   []grpc.DialOption
	}
	application struct {
		Name     string `envconfig:"APP_NAME" default:"user-service"`
		Version  string `envconfig:"APP_VERSION" default:"1.0.0"`
		Host     string `envconfig:"APP_HOST" default:"localhost"`
		Port     string `envconfig:"APP_PORT" default:"8080"`
		Log      Log
		DB       *sqlx.DB
		DBConfig DBConfig
		GRPC     GrpcClient
	}
)

var Application *application

func (a *application) InitConfig() error {
	Application = &application{}
	var err error

	Application.initENV()
	Application.initGRPCClient()
	if err = Application.initDatabase(); err != nil {
		log.Println("Database connection Error", err)
		return err
	}
	return nil
}
func (a *application) initENV() error {
	var err error
	if err = godotenv.Load(".env"); err != nil {
		return err
	}
	if err = envconfig.Process("", Application); err != nil {
		return err
	}
	return err
}

func (a *application) initDatabase() error {
	log.Println("connecting to database", a.DBConfig.HOST)
	var connString string = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		a.DBConfig.HOST, a.DBConfig.PORT, a.DBConfig.USER, a.DBConfig.PASS, a.DBConfig.NAME,
	)
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Println("Error connecting to database host:", a.DBConfig.HOST, a.DBConfig.PORT, err)
		return err
	}
	if err = db.Ping(); err != nil {
		log.Println("Error connecting to database host:", a.DBConfig.HOST, a.DBConfig.PORT, err)
		return err
	}
	db.SetMaxOpenConns(a.DBConfig.MAX_OPEN_CONN)
	db.SetMaxIdleConns(a.DBConfig.MAX_IDLE_CONN)
	a.DB = db
	log.Println("Database connection success")
	return nil
}
func (a *application) initGRPCClient() {
	ctx, cancel := customContext()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	a.GRPC.Context = &ctx
	a.GRPC.Cancel = cancel
	a.GRPC.GrpcOpts = opts
}

func customContext() (context.Context, context.CancelFunc) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	md := metadata.New(map[string]string{"x-request-from": "direct-grpc"})
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx, cancel
}
