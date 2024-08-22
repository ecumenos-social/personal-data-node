package configurations

import (
	"github.com/ecumenos-social/toolkit/types"
	"github.com/ecumenos-social/toolkitfx"
	"github.com/ecumenos-social/toolkitfx/fxgrpc"
	"github.com/ecumenos-social/toolkitfx/fxlogger"
	"github.com/ecumenos-social/toolkitfx/fxpostgres"
	cli "github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

type fxConfig struct {
	fx.Out

	App         *toolkitfx.GenericAppConfig
	AppSpecific *toolkitfx.PersonalDataNodeAppConfig
	Logger      *fxlogger.Config
	GRPC        *fxgrpc.Config
	Postgres    *fxpostgres.Config
}

var Module = func(cctx *cli.Context) fx.Option {
	return fx.Options(
		fx.Provide(func() fxConfig {
			return fxConfig{
				App: &toolkitfx.GenericAppConfig{
					ID:          cctx.Int64("pdn-app-id"),
					IDGenNode:   cctx.Int64("pdn-app-id-gen-node"),
					Name:        cctx.String("pdn-app-name"),
					Description: cctx.String("pdn-app-description"),
					RateLimit: &types.RateLimit{
						MaxRequests: cctx.Int64("pdn-app-rate-limit-max-requests"),
						Interval:    cctx.Duration("pdn-app-rate-limit-interval"),
					},
				},
				AppSpecific: &toolkitfx.PersonalDataNodeAppConfig{
					NetworkWardenID:  cctx.Int64("pdn-app-network-warden-id"),
					Label:            cctx.String("pdn-app-label"),
					AccountsCapacity: cctx.Int64("pdn-app-accounts-capacity"),
					CrawlRateLimit: &types.RateLimit{
						MaxRequests: cctx.Int64("pdn-app-crawl-rate-limit-max-requests"),
						Interval:    cctx.Duration("pdn-app-crawl-rate-limit-interval"),
					},
				},
				Logger: &fxlogger.Config{
					Production: cctx.Bool("pdn-logger-production"),
				},
				GRPC: &fxgrpc.Config{
					GRPC: fxgrpc.GRPCConfig{
						Host:                                    cctx.String("pdn-grpc-host"),
						Port:                                    cctx.String("pdn-grpc-port"),
						MaxConnectionAge:                        cctx.Duration("pdn-grpc-max-conn-age"),
						KeepAliveEnforcementMinTime:             cctx.Duration("pdn-grpc-keep-alive-enforcement-min-time"),
						KeepAliveEnforcementPermitWithoutStream: cctx.Bool("pdn-grpc-keep-alive-enforcement-permit-without-stream"),
					},
					Health: fxgrpc.HealthConfig{
						Enabled: cctx.Bool("pdn-enabled-health-server"),
						Host:    cctx.String("pdn-health-server-host"),
						Port:    cctx.String("pdn-health-server-port"),
					},
					HTTPGateway: fxgrpc.HTTPGatewayConfig{
						Host: cctx.String("pdn-http-gateway-host"),
						Port: cctx.String("pdn-http-gateway-port"),
					},
					LivenessGateway: fxgrpc.LivenessGatewayConfig{
						Host: cctx.String("pdn-liveness-gateway-host"),
						Port: cctx.String("pdn-liveness-gateway-port"),
					},
				},
				Postgres: &fxpostgres.Config{
					URL:            cctx.String("pdn-postgres-url"),
					MigrationsPath: cctx.String("pdn-postgres-migrations-path"),
				},
			}
		}),
	)
}
