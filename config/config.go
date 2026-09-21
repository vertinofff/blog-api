package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Password PasswordConfig
	Logger   LoggerConfig
	JWT      JWTConfig
}
type ServerConfig struct{ InternalPort, ExternalPort, RunMode string }
type LoggerConfig struct{ FilePath, Encoding, Level, Logger string }
type PostgresConfig struct {
	Host, Port, User, Password, DbName, SSLMode, TimeZone string
	MaxIdleConns, MaxOpenConns                            int
	ConnMaxLifetime                                       time.Duration
}
type PasswordConfig struct {
	IncludeChars, IncludeDigits, IncludeUppercase, IncludeLowercase bool
	MinLength, MaxLength                                            int
}
type JWTConfig struct {
	AccessTokenExpireDuration, RefreshTokenExpireDuration time.Duration
	Secret, RefreshSecret                                 string
}

var (
	once    sync.Once
	loaded  *Config
	loadErr error
)

// Load reads configuration once. Environment variables use names such as
// JWT_SECRET, JWT_REFRESH_SECRET, POSTGRES_PASSWORD, and PORT.
func Load() (*Config, error) { once.Do(func() { loaded, loadErr = load() }); return loaded, loadErr }

// GetConfig remains for compatibility. Application code should inject Config.
func GetConfig() *Config {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}
	return cfg
}

func load() (*Config, error) {
	env := strings.ToLower(strings.TrimSpace(os.Getenv("APP_ENV")))
	if env == "" {
		env = "development"
	}
	if env != "development" && env != "production" && env != "test" {
		return nil, fmt.Errorf("unsupported APP_ENV %q", env)
	}
	v := viper.New()
	v.SetConfigFile("config/config-" + env + ".yml")
	v.SetConfigType("yml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return nil, fmt.Errorf("read config: %w", err)
		}
	}
	for key, name := range map[string]string{"server.internalport": "PORT", "postgres.password": "POSTGRES_PASSWORD", "jwt.secret": "JWT_SECRET", "jwt.refreshsecret": "JWT_REFRESH_SECRET"} {
		_ = v.BindEnv(key, name)
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	if cfg.Server.InternalPort == "" {
		return nil, errors.New("server internalPort is required")
	}
	cfg.Server.ExternalPort = cfg.Server.InternalPort
	if cfg.Postgres.TimeZone == "" {
		cfg.Postgres.TimeZone = "UTC"
	}
	if cfg.Postgres.MaxIdleConns < 0 || cfg.Postgres.MaxOpenConns <= 0 {
		return nil, errors.New("invalid postgres pool configuration")
	}
	if cfg.JWT.Secret == "" || cfg.JWT.RefreshSecret == "" {
		return nil, errors.New("JWT_SECRET and JWT_REFRESH_SECRET are required")
	}
	if cfg.JWT.Secret == cfg.JWT.RefreshSecret {
		return nil, errors.New("access and refresh JWT secrets must differ")
	}
	if env == "production" {
		if cfg.Postgres.SSLMode == "disable" || cfg.Postgres.SSLMode == "" {
			return nil, errors.New("production PostgreSQL requires TLS")
		}
		if len(cfg.JWT.Secret) < 32 || len(cfg.JWT.RefreshSecret) < 32 {
			return nil, errors.New("production JWT secrets must be at least 32 bytes")
		}
	}
	return &cfg, nil
}
