package microfiber

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConnConfig struct {
	*pgxpool.Config

	//重试次数
	Retry uint

	//重试间隔时间
	RetryInterval time.Duration

	RetryCallback func(err error)
}

func NewPostgresPool(config *PostgresConnConfig) (*pgxpool.Pool, error) {
	var dbpool *pgxpool.Pool
	pc, _ := pgxpool.ParseConfig("")
	if config.BeforeConnect != nil {
		pc.BeforeConnect = config.BeforeConnect
	}
	if config.AfterConnect != nil {
		pc.AfterConnect = config.AfterConnect
	}
	if config.BeforeAcquire != nil {
		pc.BeforeAcquire = config.BeforeAcquire
	}
	if config.AfterRelease != nil {
		pc.AfterRelease = config.AfterRelease
	}
	if config.BeforeClose != nil {
		pc.BeforeClose = config.BeforeClose
	}
	if config.MaxConnLifetime != 0 {
		pc.MaxConnLifetime = config.MaxConnLifetime
	}
	if config.MaxConnLifetimeJitter != 0 {
		pc.MaxConnLifetimeJitter = config.MaxConnLifetimeJitter
	}
	if config.MaxConnIdleTime != 0 {
		pc.MaxConnIdleTime = config.MaxConnIdleTime
	}
	if config.MaxConns != 0 {
		pc.MaxConns = config.MaxConns
	}
	if config.MinConns != 0 {
		pc.MinConns = config.MinConns
	}
	if config.HealthCheckPeriod != 0 {
		pc.HealthCheckPeriod = config.HealthCheckPeriod
	}

	pc.ConnConfig, _ = pgx.ParseConfig("")
	pc.ConnConfig.Host = config.ConnConfig.Host
	pc.ConnConfig.Port = config.ConnConfig.Port
	pc.ConnConfig.Database = config.ConnConfig.Database
	pc.ConnConfig.User = config.ConnConfig.User
	pc.ConnConfig.Password = config.ConnConfig.Password
	if config.ConnConfig.TLSConfig != nil {
		pc.ConnConfig.TLSConfig = config.ConnConfig.TLSConfig
	}
	if config.ConnConfig.ConnectTimeout != 0 {
		pc.ConnConfig.ConnectTimeout = config.ConnConfig.ConnectTimeout
	}
	if config.ConnConfig.DialFunc != nil {
		pc.ConnConfig.DialFunc = config.ConnConfig.DialFunc
	}
	if config.ConnConfig.LookupFunc != nil {
		pc.ConnConfig.LookupFunc = config.ConnConfig.LookupFunc
	}
	if config.ConnConfig.BuildFrontend != nil {
		pc.ConnConfig.BuildFrontend = config.ConnConfig.BuildFrontend
	}
	if config.ConnConfig.RuntimeParams != nil {
		pc.ConnConfig.RuntimeParams = config.ConnConfig.RuntimeParams
	}
	if config.ConnConfig.KerberosSrvName != "" {
		pc.ConnConfig.KerberosSrvName = config.ConnConfig.KerberosSrvName
	}
	if config.ConnConfig.KerberosSpn != "" {
		pc.ConnConfig.KerberosSpn = config.ConnConfig.KerberosSpn
	}
	if config.ConnConfig.Fallbacks != nil {
		pc.ConnConfig.Fallbacks = config.ConnConfig.Fallbacks
	}
	if config.ConnConfig.ValidateConnect != nil {
		pc.ConnConfig.ValidateConnect = config.ConnConfig.ValidateConnect
	}
	if config.ConnConfig.AfterConnect != nil {
		pc.ConnConfig.AfterConnect = config.ConnConfig.AfterConnect
	}
	if config.ConnConfig.OnNotice != nil {
		pc.ConnConfig.OnNotice = config.ConnConfig.OnNotice
	}
	if config.ConnConfig.OnNotification != nil {
		pc.ConnConfig.OnNotification = config.ConnConfig.OnNotification
	}
	err := Retry(
		func() (err error) {
			dbpool, _ = pgxpool.NewWithConfig(context.Background(), pc)
			err = dbpool.Ping(context.Background())
			if err != nil {
				config.RetryCallback(err)
				return
			}
			return
		},
		config.Retry,
		config.RetryInterval,
	)
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}
