package gomavlib

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/aler9/gomavlib/pkg/autoreconnector"
)

type endpointClientConf interface {
	isUDP() bool
	getAddress() string
	init() (Endpoint, error)
}

// EndpointTCPClient sets up a endpoint that works with a TCP client.
// TCP is fit for routing frames through the internet, but is not the most
// appropriate way for transferring frames from a UAV to a GCS, since it does
// not allow frame losses.
type EndpointTCPClient struct {
	// domain name or IP of the server to connect to, example: 1.2.3.4:5600
	Address string
}

func (EndpointTCPClient) isUDP() bool {
	return false
}

func (conf EndpointTCPClient) getAddress() string {
	return conf.Address
}

func (conf EndpointTCPClient) init() (Endpoint, error) {
	return initEndpointClient(conf)
}

// EndpointUDPClient sets up a endpoint that works with a UDP client.
type EndpointUDPClient struct {
	// domain name or IP of the server to connect to, example: 1.2.3.4:5600
	Address string
}

func (EndpointUDPClient) isUDP() bool {
	return true
}

func (conf EndpointUDPClient) getAddress() string {
	return conf.Address
}

func (conf EndpointUDPClient) init() (Endpoint, error) {
	return initEndpointClient(conf)
}

type endpointClient struct {
	conf endpointClientConf
	io.ReadWriteCloser
}

func initEndpointClient(conf endpointClientConf) (Endpoint, error) {
	_, _, err := net.SplitHostPort(conf.getAddress())
	if err != nil {
		return nil, fmt.Errorf("invalid address")
	}

	t := &endpointClient{
		conf: conf,
		ReadWriteCloser: autoreconnector.New(
			func(ctx context.Context) (io.ReadWriteCloser, error) {
				// solve address and connect
				// in UDP, the only possible error is a DNS failure
				// in TCP, the handshake must be completed
				network := func() string {
					if conf.isUDP() {
						return "udp4"
					}
					return "tcp4"
				}()
				timedContext, timedContextClose := context.WithTimeout(ctx, netConnectTimeout)
				nconn, err := (&net.Dialer{}).DialContext(timedContext, network, conf.getAddress())
				timedContextClose()
				if err != nil {
					return nil, err
				}

				return &netTimedConn{nconn}, nil
			},
		),
	}

	return t, nil
}

func (t *endpointClient) isEndpoint() {}

func (t *endpointClient) Conf() EndpointConf {
	return t.conf
}

func (t *endpointClient) label() string {
	return fmt.Sprintf("%s:%s", func() string {
		if t.conf.isUDP() {
			return "udp"
		}
		return "tcp"
	}(), t.conf.getAddress())
}
