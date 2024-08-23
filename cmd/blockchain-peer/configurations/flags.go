package configurations

import (
	"time"

	cli "github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.Int64Flag{
		Name:    "pdn-app-id",
		Usage:   "it is application ID",
		Value:   0,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_APP_ID"},
	},
	&cli.Int64Flag{
		Name:    "pdn-app-id-gen-node",
		Usage:   "it is application ID generation node",
		Value:   0,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_APP_ID_GEN_NODE"},
	},
	&cli.StringFlag{
		Name:    "pdn-app-name",
		Usage:   "it is unique application name",
		Value:   "name",
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_APP_NAME"},
	},
	&cli.StringFlag{
		Name:    "pdn-app-description",
		Usage:   "it is application description",
		Value:   "it is personal data node",
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_APP_DESCRIPTION"},
	},
	&cli.StringFlag{
		Name:    "pdn-app-label",
		Usage:   "it is unique label specific for this personal data node",
		Value:   "pdn",
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_APP_LABEL"},
	},
	&cli.Int64Flag{
		Name:    "pdn-app-network-warden-id",
		Usage:   "it is ID of the warden of the network",
		Value:   0,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_APP_NW_ID"},
	},
	&cli.Int64Flag{
		Name:    "pdn-app-accounts-capacity",
		Usage:   "it is the account's capacity of the personal data node",
		Value:   0,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_APP_ACCOUNTS_CAPACITY"},
	},
	&cli.Int64Flag{
		Name:    "pdn-app-crawl-rate-limit-max-requests",
		Usage:   "it is rate limit max crawl requests amount",
		Value:   1000,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_APP_CRAWL_RATE_LIMIT_MAX_REQUESTS"},
	},
	&cli.DurationFlag{
		Name:    "pdn-app-crawl-rate-limit-interval",
		Usage:   "it is rate limit interval for crawl requests",
		Value:   time.Minute,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_APP_CRAWL_RATE_LIMIT_INTERVAL"},
	},
	&cli.Int64Flag{
		Name:    "pdn-app-rate-limit-max-requests",
		Usage:   "it is rate limit max requests amount",
		Value:   1000,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_APP_RATE_LIMIT_MAX_REQUESTS"},
	},
	&cli.DurationFlag{
		Name:    "pdn-app-rate-limit-interval",
		Usage:   "it is rate limit interval",
		Value:   time.Minute,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_APP_RATE_LIMIT_INTERVAL"},
	},
	&cli.BoolFlag{
		Name:    "pdn-logger-production",
		Usage:   "make it true if you need logging on production environment",
		Value:   false,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_LOGGER_PRODUCTION"},
	},
	&cli.StringFlag{
		Name:    "pdn-grpc-host",
		Usage:   "it is gRPC server host",
		Value:   "0.0.0.0",
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_GRPC_HOST"},
	},
	&cli.StringFlag{
		Name:    "pdn-grpc-port",
		Usage:   "it is gRPC server port",
		Value:   "8080",
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_GRPC_PORT"},
	},
	&cli.StringFlag{
		Name:    "pdn-http-gateway-host",
		Usage:   "it is HTTP gateway host",
		Value:   "0.0.0.0",
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_HTTP_GATEWAY_HOST"},
	},
	&cli.StringFlag{
		Name:    "pdn-http-gateway-port",
		Usage:   "it is HTTP gateway port",
		Value:   "9090",
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_HTTP_GATEWAY_PORT"},
	},
	&cli.BoolFlag{
		Name:    "pdn-enabled-health-server",
		Usage:   "make it true if you need to enable health server",
		Value:   false,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_ENABLED_HEALTH_SERVER"},
	},
	&cli.StringFlag{
		Name:    "pdn-health-server-host",
		Usage:   "it is health server host",
		Value:   "0.0.0.0",
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_HEALTH_SERVER_HOST"},
	},
	&cli.StringFlag{
		Name:    "pdn-health-server-port",
		Usage:   "it is health server port",
		Value:   "10010",
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_HEALTH_SERVER_PORT"},
	},
	&cli.StringFlag{
		Name:    "pdn-liveness-gateway-host",
		Usage:   "it is liveness gateway host",
		Value:   "0.0.0.0",
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_LIVENESS_GATEWAY_HOST"},
	},
	&cli.StringFlag{
		Name:    "pdn-liveness-gateway-port",
		Usage:   "it is liveness gateway port",
		Value:   "8086",
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_LIVENESS_GATEWAY_PORT"},
	},
	&cli.DurationFlag{
		Name:    "pdn-grpc-max-conn-age",
		Usage:   "it is max age of connection with gRPC server",
		Value:   5 * time.Minute,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_GRPC_MAX_CONNECTION_AGE"},
	},
	&cli.DurationFlag{
		Name:    "pdn-grpc-keep-alive-enforcement-min-time",
		Usage:   "it is minimal time of keep alive enforcement gRPC server",
		Value:   time.Minute,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_GRPC_KEEP_ALIVE_ENFORCEMENT_MIN_TIME"},
	},
	&cli.BoolFlag{
		Name:    "pdn-grpc-keep-alive-enforcement-permit-without-stream",
		Usage:   "",
		Value:   true,
		EnvVars: []string{"PERSONAL_DATA_NODE_BLOCKCHAIN_PEER_GRPC_KEEP_ALIVE_ENFORCEMENT_PERMIT_WITHOUT_STREAM"},
	},
}
