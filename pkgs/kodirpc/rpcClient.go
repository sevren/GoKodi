package kodirpc

import (
	"net/url"

	"github.com/KeisukeYamashita/jsonrpc"
	"github.com/pkg/errors"
)

type Client struct {
	client *jsonrpc.RPCClient
}

type service struct {
	client *Client
}

const (
	defaultURL = "http://localhost:8080/jsonrpc"
)

//New creates a new JSON RPC client
func New(addr string) (*Client, error) {

	if addr == "" {
		addr = defaultURL
	}
	u, err := url.Parse(addr)
	if err != nil {
		return nil, errors.WithMessagef(err, "url can't be parsed %s", addr)
	}

	rpcClient := jsonrpc.NewRPCClient(addr)

	user := u.User.Username()
	pass, passSet := u.User.Password()

	if user != "" && passSet {
		rpcClient.SetBasicAuth(user, pass)
	}

	return &Client{rpcClient}, nil

}
