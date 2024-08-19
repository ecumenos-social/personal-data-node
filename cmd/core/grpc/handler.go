package grpc

import (
	"context"

	grpcutils "github.com/ecumenos-social/grpc-utils"
	pbv1 "github.com/ecumenos-social/schemas/proto/gen/personaldatanode/v1"
	"github.com/ecumenos-social/toolkitfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type Handler struct {
	pbv1.CoreServiceServer

	logger *zap.Logger

	networkWardenID int64
}

var _ pbv1.CoreServiceServer = (*Handler)(nil)

type handlerParams struct {
	fx.In

	Config *toolkitfx.PersonalDataNodeAppConfig
	Logger *zap.Logger
}

func NewHandler(params handlerParams) *Handler {
	return &Handler{
		logger: params.Logger,

		networkWardenID: params.Config.NetworkWardenID,
	}
}

func (h *Handler) customizeLogger(ctx context.Context, operationName string) *zap.Logger {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return h.logger
	}

	l := h.logger.With(
		zap.String("operation-name", operationName),
		zap.Int64("network-warden-id", h.networkWardenID),
	)
	if corrID := md.Get("correlation-id"); len(corrID) > 0 {
		l = l.With(zap.String("correlation-id", corrID[0]))
	}
	if ip := grpcutils.ExtractRemoteIPAddress(ctx); ip != "" {
		l = l.With(zap.String("remote-ip-address", ip))
	}

	return l
}
