package configurations

import (
	"time"

	cli "github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.Int64Flag{
		Name:    "nw-app-id",
		Usage:   "it is application ID",
		Value:   0,
		EnvVars: []string{"PERSONAL_DATA_NODE_APP_ID"},
	},
	&cli.Int64Flag{
		Name:    "nw-app-id-gen-node",
		Usage:   "it is application ID generation node",
		Value:   0,
		EnvVars: []string{"PERSONAL_DATA_NODE_APP_ID_GEN_NODE"},
	},
	&cli.StringFlag{
		Name:    "nw-app-name",
		Usage:   "it is unique application name",
		Value:   "name",
		EnvVars: []string{"PERSONAL_DATA_NODE_APP_NAME"},
	},
	&cli.StringFlag{
		Name:    "nw-app-description",
		Usage:   "it is application description",
		Value:   "it is personal data node",
		EnvVars: []string{"PERSONAL_DATA_NODE_APP_DESCRIPTION"},
	},
	&cli.StringFlag{
		Name:    "nw-app-address-suffix",
		Usage:   "it is unique address suffix specific for this personal data node",
		Value:   "nw",
		EnvVars: []string{"PERSONAL_DATA_NODE_APP_ADDRESS_SUFFIX"},
	},
	&cli.Int64Flag{
		Name:    "nw-app-rate-limit-max-requests",
		Usage:   "it is rate limit max requests amount",
		Value:   1000,
		EnvVars: []string{"PERSONAL_DATA_NODE_APP_RATE_LIMIT_MAX_REQUESTS"},
	},
	&cli.DurationFlag{
		Name:    "nw-app-rate-limit-interval",
		Usage:   "it is rate limit interval",
		Value:   time.Minute,
		EnvVars: []string{"PERSONAL_DATA_NODE_APP_RATE_LIMIT_INTERVAL"},
	},
	&cli.BoolFlag{
		Name:    "nw-logger-production",
		Usage:   "make it true if you need logging on production environment",
		Value:   false,
		EnvVars: []string{"PERSONAL_DATA_NODE_LOGGER_PRODUCTION"},
	},
	&cli.StringFlag{
		Name:    "nw-grpc-host",
		Usage:   "it is gRPC server host",
		Value:   "0.0.0.0",
		EnvVars: []string{"PERSONAL_DATA_NODE_GRPC_HOST"},
	},
	&cli.StringFlag{
		Name:    "nw-grpc-port",
		Usage:   "it is gRPC server port",
		Value:   "8080",
		EnvVars: []string{"PERSONAL_DATA_NODE_GRPC_PORT"},
	},
	&cli.StringFlag{
		Name:    "nw-http-gateway-host",
		Usage:   "it is HTTP gateway host",
		Value:   "0.0.0.0",
		EnvVars: []string{"PERSONAL_DATA_NODE_HTTP_GATEWAY_HOST"},
	},
	&cli.StringFlag{
		Name:    "nw-http-gateway-port",
		Usage:   "it is HTTP gateway port",
		Value:   "9090",
		EnvVars: []string{"PERSONAL_DATA_NODE_HTTP_GATEWAY_PORT"},
	},
	&cli.BoolFlag{
		Name:    "nw-enabled-health-server",
		Usage:   "make it true if you need to enable health server",
		Value:   false,
		EnvVars: []string{"PERSONAL_DATA_NODE_ENABLED_HEALTH_SERVER"},
	},
	&cli.StringFlag{
		Name:    "nw-health-server-host",
		Usage:   "it is health server host",
		Value:   "0.0.0.0",
		EnvVars: []string{"PERSONAL_DATA_NODE_HEALTH_SERVER_HOST"},
	},
	&cli.StringFlag{
		Name:    "nw-health-server-port",
		Usage:   "it is health server port",
		Value:   "10010",
		EnvVars: []string{"PERSONAL_DATA_NODE_HEALTH_SERVER_PORT"},
	},
	&cli.StringFlag{
		Name:    "nw-liveness-gateway-host",
		Usage:   "it is liveness gateway host",
		Value:   "0.0.0.0",
		EnvVars: []string{"PERSONAL_DATA_NODE_LIVENESS_GATEWAY_HOST"},
	},
	&cli.StringFlag{
		Name:    "nw-liveness-gateway-port",
		Usage:   "it is liveness gateway port",
		Value:   "8086",
		EnvVars: []string{"PERSONAL_DATA_NODE_LIVENESS_GATEWAY_PORT"},
	},
	&cli.DurationFlag{
		Name:    "nw-grpc-max-conn-age",
		Usage:   "it is max age of connection with gRPC server",
		Value:   5 * time.Minute,
		EnvVars: []string{"PERSONAL_DATA_NODE_GRPC_MAX_CONNECTION_AGE"},
	},
	&cli.DurationFlag{
		Name:    "nw-grpc-keep-alive-enforcement-min-time",
		Usage:   "it is minimal time of keep alive enforcement gRPC server",
		Value:   time.Minute,
		EnvVars: []string{"PERSONAL_DATA_NODE_GRPC_KEEP_ALIVE_ENFORCEMENT_MIN_TIME"},
	},
	&cli.BoolFlag{
		Name:    "nw-grpc-keep-alive-enforcement-permit-without-stream",
		Usage:   "",
		Value:   true,
		EnvVars: []string{"PERSONAL_DATA_NODE_GRPC_KEEP_ALIVE_ENFORCEMENT_PERMIT_WITHOUT_STREAM"},
	},
	&cli.StringFlag{
		Name:    "nw-postgres-url",
		Usage:   "it is URL of postgres database connected to the app",
		Value:   `postgresql://ecumenosuser:rootpassword@localhost:5432/ecumenos_personal_data_node_db`,
		EnvVars: []string{"PERSONAL_DATA_NODE_POSTGRES_URL"},
	},
	&cli.StringFlag{
		Name:    "nw-postgres-migrations-path",
		Usage:   "it is path to directory with postgres migrations",
		Value:   `file://cmd/core/pgmigrations`,
		EnvVars: []string{"PERSONAL_DATA_NODE_POSTGRES_MIGRATIONS_PATH"},
	},
}
