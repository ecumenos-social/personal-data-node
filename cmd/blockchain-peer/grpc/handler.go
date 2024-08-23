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
	pbv1.BlockchainPeerServiceServer

	logger *zap.Logger

	networkWardenID int64
}

var _ pbv1.BlockchainPeerServiceServer = (*Handler)(nil)

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

func (h *Handler) CreateTransaction(ctx context.Context, _ *pbv1.BlockchainPeerServiceCreateTransactionRequest) (*pbv1.BlockchainPeerServiceCreateTransactionResponse, error) {
	logger := h.customizeLogger(ctx, "CreateTransaction")
	defer logger.Info("request processed")

	return &pbv1.BlockchainPeerServiceCreateTransactionResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) Validate(ctx context.Context, _ *pbv1.BlockchainPeerServiceValidateRequest) (*pbv1.BlockchainPeerServiceValidateResponse, error) {
	logger := h.customizeLogger(ctx, "Validate")
	defer logger.Info("request processed")

	return &pbv1.BlockchainPeerServiceValidateResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) GetBlocks(ctx context.Context, _ *pbv1.BlockchainPeerServiceGetBlocksRequest) (*pbv1.BlockchainPeerServiceGetBlocksResponse, error) {
	logger := h.customizeLogger(ctx, "GetBlocks")
	defer logger.Info("request processed")

	return &pbv1.BlockchainPeerServiceGetBlocksResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) GetEntities(ctx context.Context, _ *pbv1.BlockchainPeerServiceGetEntitiesRequest) (*pbv1.BlockchainPeerServiceGetEntitiesResponse, error) {
	logger := h.customizeLogger(ctx, "GetEntities")
	defer logger.Info("request processed")

	return &pbv1.BlockchainPeerServiceGetEntitiesResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}
