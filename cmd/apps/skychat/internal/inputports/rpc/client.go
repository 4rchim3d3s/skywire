// Package rpc contains code of the rpc handler for inputports
package rpc

import (
	"log"
	"net/rpc"

	"github.com/skycoin/skywire-utilities/pkg/cipher"
	"github.com/skycoin/skywire/cmd/apps/skychat/internal/app"
	"github.com/skycoin/skywire/cmd/apps/skychat/internal/inputports/rpc/chat"
)

// Client represents the rpc client running for this service
type Client struct {
	appServices app.Services
	rpcPort     string
}

// NewClient RPC Client constructor
func NewClient(appServices app.Services, rpcPort string) *Client {
	rc := &Client{appServices: appServices, rpcPort: rpcPort}
	return rc
}

// SendTextMessage sends the command to send a message via rpc
func (c *Client) SendTextMessage(VisorPk cipher.PubKey, ServerPk cipher.PubKey, RoomPk cipher.PubKey, Message string) error {

	rpcClient, err := rpc.DialHTTP("tcp", c.rpcPort)
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	stmrm := chat.SendTextMessageRequestModel{
		VisorPk:  VisorPk,
		ServerPk: ServerPk,
		RoomPk:   RoomPk,
		Msg:      Message,
	}

	err = rpcClient.Call(chat.SendTextMessageRPCParam, stmrm, nil)
	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}

	return nil
}
