package kodirpc

import (
	"github.com/KeisukeYamashita/jsonrpc"
	"github.com/pkg/errors"
)

func (c *Client) ScanVideoLibrary() (*jsonrpc.RPCResponse, error) {
	response, err := c.client.Call("VideoLibrary.Scan")
	if err != nil {
		return nil, errors.WithMessagef(err, "error calling service VideoLibrary.Scan")
	}
	return response, nil
}

func (c *Client) GetTVShows() (*jsonrpc.RPCResponse, error) {
	response, err := c.client.Call("VideoLibrary.GetTvShows")
	if err != nil {
		return nil, errors.WithMessagef(err, "error calling service VideoLibrary.GetTvShows")
	}
	return response, nil
}
