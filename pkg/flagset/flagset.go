package flagset

import (
	"github.com/micro/cli"
	"github.com/owncloud/ocis-graph/pkg/config"
)

// RootWithConfig applies cfg to the root flagset
func RootWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "config-file",
			Value:       "",
			Usage:       "Path to config file",
			EnvVar:      "GRAPH_CONFIG_FILE",
			Destination: &cfg.File,
		},
		&cli.StringFlag{
			Name:        "log-level",
			Value:       "info",
			Usage:       "Set logging level",
			EnvVar:      "GRAPH_LOG_LEVEL",
			Destination: &cfg.Log.Level,
		},
		&cli.BoolTFlag{
			Name:        "log-pretty",
			Usage:       "Enable pretty logging",
			EnvVar:      "GRAPH_LOG_PRETTY",
			Destination: &cfg.Log.Pretty,
		},
		&cli.BoolTFlag{
			Name:        "log-color",
			Usage:       "Enable colored logging",
			EnvVar:      "GRAPH_LOG_COLOR",
			Destination: &cfg.Log.Color,
		},
	}
}

// HealthWithConfig applies cfg to the root flagset
func HealthWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       "0.0.0.0:9124",
			Usage:       "Address to debug endpoint",
			EnvVar:      "GRAPH_DEBUG_ADDR",
			Destination: &cfg.Debug.Addr,
		},
	}
}

// ServerWithConfig applies cfg to the root flagset
func ServerWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "tracing-enabled",
			Usage:       "Enable sending traces",
			EnvVar:      "GRAPH_TRACING_ENABLED",
			Destination: &cfg.Tracing.Enabled,
		},
		&cli.StringFlag{
			Name:        "tracing-type",
			Value:       "jaeger",
			Usage:       "Tracing backend type",
			EnvVar:      "GRAPH_TRACING_TYPE",
			Destination: &cfg.Tracing.Type,
		},
		&cli.StringFlag{
			Name:        "tracing-endpoint",
			Value:       "",
			Usage:       "Endpoint for the agent",
			EnvVar:      "GRAPH_TRACING_ENDPOINT",
			Destination: &cfg.Tracing.Endpoint,
		},
		&cli.StringFlag{
			Name:        "tracing-collector",
			Value:       "",
			Usage:       "Endpoint for the collector",
			EnvVar:      "GRAPH_TRACING_COLLECTOR",
			Destination: &cfg.Tracing.Collector,
		},
		&cli.StringFlag{
			Name:        "tracing-service",
			Value:       "graph",
			Usage:       "Service name for tracing",
			EnvVar:      "GRAPH_TRACING_SERVICE",
			Destination: &cfg.Tracing.Service,
		},
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       "0.0.0.0:9124",
			Usage:       "Address to bind debug server",
			EnvVar:      "GRAPH_DEBUG_ADDR",
			Destination: &cfg.Debug.Addr,
		},
		&cli.StringFlag{
			Name:        "debug-token",
			Value:       "",
			Usage:       "Token to grant metrics access",
			EnvVar:      "GRAPH_DEBUG_TOKEN",
			Destination: &cfg.Debug.Token,
		},
		&cli.BoolFlag{
			Name:        "debug-pprof",
			Usage:       "Enable pprof debugging",
			EnvVar:      "GRAPH_DEBUG_PPROF",
			Destination: &cfg.Debug.Pprof,
		},
		&cli.BoolFlag{
			Name:        "debug-zpages",
			Usage:       "Enable zpages debugging",
			EnvVar:      "GRAPH_DEBUG_ZPAGES",
			Destination: &cfg.Debug.Zpages,
		},
		&cli.StringFlag{
			Name:        "http-addr",
			Value:       "0.0.0.0:9120",
			Usage:       "Address to bind http server",
			EnvVar:      "GRAPH_HTTP_ADDR",
			Destination: &cfg.HTTP.Addr,
		},
		&cli.StringFlag{
			Name:        "http-root",
			Value:       "/",
			Usage:       "Root path of http server",
			EnvVar:      "GRAPH_HTTP_ROOT",
			Destination: &cfg.HTTP.Root,
		},
		&cli.StringFlag{
			Name:        "http-namespace",
			Value:       "com.owncloud.web",
			Usage:       "Set the base namespace for the http service for service discovery",
			EnvVar:      "GRAPH_HTTP_NAMESPACE",
			Destination: &cfg.HTTP.Namespace,
		},
		&cli.StringFlag{
			Name:        "ldap-network",
			Value:       "tcp",
			Usage:       "Network protocol to use to connect to the Ldap server",
			EnvVar:      "GRAPH_LDAP_NETWORK",
			Destination: &cfg.Ldap.Network,
		},
		&cli.StringFlag{
			Name:        "ldap-address",
			Value:       "localhost:9125",
			Usage:       "Address to connect to the Ldap server",
			EnvVar:      "GRAPH_LDAP_ADDRESS",
			Destination: &cfg.Ldap.Address,
		},
		&cli.StringFlag{
			Name:        "ldap-username",
			Value:       "cn=admin,dc=example,dc=org",
			Usage:       "User to bind to the Ldap server",
			EnvVar:      "GRAPH_LDAP_USERNAME",
			Destination: &cfg.Ldap.UserName,
		},
		&cli.StringFlag{
			Name:        "ldap-password",
			Value:       "admin",
			Usage:       "Password to bind to the Ldap server",
			EnvVar:      "GRAPH_LDAP_PASSWORD",
			Destination: &cfg.Ldap.Password,
		},
		&cli.StringFlag{
			Name:        "ldap-basedn-users",
			Value:       "ou=users,dc=example,dc=org",
			Usage:       "BaseDN to look for users",
			EnvVar:      "GRAPH_LDAP_BASEDN_USERS",
			Destination: &cfg.Ldap.BaseDNUsers,
		},
		&cli.StringFlag{
			Name:        "ldap-basedn-groups",
			Value:       "ou=groups,dc=example,dc=org",
			Usage:       "BaseDN to look for users",
			EnvVar:      "GRAPH_LDAP_BASEDN_GROUPS",
			Destination: &cfg.Ldap.BaseDNGroups,
		},
		&cli.StringFlag{
			Name:        "oidc-endpoint",
			Value:       "",
			Usage:       "OpenIDConnect endpoint",
			EnvVar:      "GRAPH_OIDC_ENDPOINT",
			Destination: &cfg.OpenIDConnect.Endpoint,
		},
		&cli.BoolFlag{
			Name:        "oidc-insecure",
			Usage:       "OpenIDConnect endpoint",
			EnvVar:      "GRAPH_OIDC_INSECURE",
			Destination: &cfg.OpenIDConnect.Insecure,
		},
		&cli.StringFlag{
			Name:        "oidc-realm",
			Value:       "",
			Usage:       "OpenIDConnect realm",
			EnvVar:      "GRAPH_OIDC_REALM",
			Destination: &cfg.OpenIDConnect.Realm,
		},
	}
}
