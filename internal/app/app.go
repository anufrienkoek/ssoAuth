package app

import (
	"log/slog"
	grpcapp "ssoAuth/internal/app/grpc"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: инициализировать хранилище

	// TODO : инициализировать auth service

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
