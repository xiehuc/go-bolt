package goBolt

import (
	"github.com/mindstand/go-bolt/errors"
	"time"
)

type Opt func(*Client) error

// WithConnectionString provides client option with connection string
func WithConnectionString(connString string) Opt {
	return func(client *Client) error {
		if client == nil {
			return errors.Wrap(errors.ErrConfiguration, "client can not be nil")
		}

		if client.host != "" || client.port != 0 {
			return errors.Wrap(errors.ErrConfiguration, "can not call WithConnectionString and WithHostPort")
		}

		client.connStr = connString
		return nil
	}
}

// allows setting the host and port of neo4j
func WithHostPort(host string, port int) Opt {
	return func(client *Client) error {
		if client == nil {
			return errors.Wrap(errors.ErrConfiguration, "client can not be nil")
		}

		if client.connStr != "" {
			return errors.Wrap(errors.ErrConfiguration, "can not call WithHostPort and WithConnectionString")
		}

		client.host = host
		client.port = port
		return nil
	}
}

// allows setting protocol to bolt+routing
func WithRouting() Opt {
	return func(client *Client) error {
		if client == nil {
			return errors.Wrap(errors.ErrConfiguration, "client can not be nil")
		}

		if client.connStr != "" {
			return errors.Wrap(errors.ErrConfiguration, "can not call WithRouting and WithConnectionString")
		}

		client.routing = true
		return nil
	}
}

// allows setting chunk size
func WithChunkSize(size uint16) Opt {
	return func(client *Client) error {
		if client == nil {
			return errors.Wrap(errors.ErrConfiguration, "client can not be nil")
		}

		client.chunkSize = size
		return nil
	}
}

// allows authentication with basic auth
func WithBasicAuth(username, password string) Opt {
	return func(client *Client) error {
		if client == nil {
			return errors.Wrap(errors.ErrConfiguration, "client can not be nil")
		}

		client.user = username
		client.password = password

		return nil
	}
}

// allows authentication with tls
func WithTLS(cacertPath, certPath, keyPath string, tlsNoVerify bool) Opt {
	return func(client *Client) error {
		if client == nil {
			return errors.Wrap(errors.ErrConfiguration, "client can not be nil")
		}

		client.caCertFile = cacertPath
		client.certFile = certPath
		client.keyFile = keyPath
		client.useTLS = true
		client.tlsNoVerify = tlsNoVerify
		return nil
	}
}

// tells client to negotiate version
func WithProtocolVersionNegotiation() Opt {
	return func(client *Client) error {
		if client == nil {
			return errors.Wrap(errors.ErrConfiguration, "client can not be nil")
		}

		client.negotiateVersion = true
		return nil
	}
}

// tells client what timeout it should use
func WithTimeout(timeout time.Duration) Opt {
	return func(client *Client) error {
		if client == nil {
			return errors.Wrap(errors.ErrConfiguration, "client can not be nil")
		}

		client.timeout = timeout
		return nil
	}
}

// tells client which bolt version to use
//func WithVersion(version int) Opt {
//	return func(client *Client) error {
//		if client == nil {
//			return errors.Wrap(errors.ErrConfiguration, "client can not be nil")
//		}
//
//		if client.negotiateVersion {
//			return errors.Wrap(errors.ErrConfiguration, "can not set client version and negotiate version")
//		}
//
//		if strings.Contains(version, "4") {
//			client.serverVersion = make([]byte, 4)
//			client.supportsV4 = true
//		} else {
//			client.serverVersion = make([]byte, 4)
//		}
//
//		return nil
//	}
//}

//func WithPool(max int) Opt {
//	return func(client *Client) error {
//		if client == nil {
//			return errors.Wrap(errors.ErrConfiguration, "client can not be nil")
//		}
//
//		if max < 1 {
//			return errors.Wrap(errors.ErrConfiguration, "client pool has to have at least 1 connection")
//		}
//
//		client.pooled = true
//		client.maxConnections = max
//		return nil
//	}
//}
//
//func WithV4CreateDBIfNotExists() Opt {
//	return func(client *Client) error {
//		if client == nil {
//			return errors.Wrap(errors.ErrConfiguration, "client can not be nil")
//		}
//
//		client.createDbIfNotExists = true
//
//		return nil
//	}
//}
