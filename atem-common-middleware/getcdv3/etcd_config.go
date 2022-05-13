package getcdv3

import "errors"

type EtcdConfig struct {
	ServiceName string   `mapstructure:"service-name" yaml:"service-name"`
	EtcdSchema  string   `mapstructure:"schema" yaml:"schema"`
	EtcdAddrs   []string `mapstructure:"addrs" yaml:"addrs"`
	EtcdEnable  bool     `mapstructure:"enable" json:"enable"`
	EtcdUsePool bool     `mapstructure:"use-pool" json:"use-pool"`
}

var (
	// ErrClosed is the error when the client pool is closed
	ErrClosed = errors.New("grpc pool: client pool is closed")
	// ErrTimeout is the error when the client pool timed out
	ErrTimeout = errors.New("grpc pool: client pool timed out")
	// ErrAlreadyClosed is the error when the client conn was already closed
	ErrAlreadyClosed = errors.New("grpc pool: the connection was already closed")
	// ErrFullPool is the error when the pool is already full
	ErrFullPool = errors.New("grpc pool: closing a ClientConn into a full pool")
)
