package grpc

import (
	"context"

	grpcutils "github.com/ecumenos-social/grpc-utils"
	pbv1 "github.com/ecumenos-social/schemas/proto/gen/personaldatanode/v1"
	"github.com/ecumenos-social/toolkitfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pbv1.AdminServiceServer

	logger *zap.Logger

	networkWardenID int64
}

var _ pbv1.AdminServiceServer = (*Handler)(nil)

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

func (h *Handler) LoginAdmin(ctx context.Context, _ *pbv1.AdminServiceLoginAdminRequest) (*pbv1.AdminServiceLoginAdminResponse, error) {
	logger := h.customizeLogger(ctx, "LoginAdmin")
	defer logger.Info("request processed")

	return &pbv1.AdminServiceLoginAdminResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) RefreshAdminToken(ctx context.Context, _ *pbv1.AdminServiceRefreshAdminTokenRequest) (*pbv1.AdminServiceRefreshAdminTokenResponse, error) {
	logger := h.customizeLogger(ctx, "RefreshAdminToken")
	defer logger.Info("request processed")

	return &pbv1.AdminServiceRefreshAdminTokenResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) LogoutAdmin(ctx context.Context, _ *pbv1.AdminServiceLogoutAdminRequest) (*pbv1.AdminServiceLogoutAdminResponse, error) {
	logger := h.customizeLogger(ctx, "LogoutAdmin")
	defer logger.Info("request processed")

	return &pbv1.AdminServiceLogoutAdminResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) ChangeAdminPassword(ctx context.Context, _ *pbv1.AdminServiceChangeAdminPasswordRequest) (*pbv1.AdminServiceChangeAdminPasswordResponse, error) {
	logger := h.customizeLogger(ctx, "ChangeAdminPassword")
	defer logger.Info("request processed")

	return &pbv1.AdminServiceChangeAdminPasswordResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}
