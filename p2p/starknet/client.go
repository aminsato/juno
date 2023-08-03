package p2p

import (
	"io"

	p2p_spec "github.com/NethermindEth/juno/p2p/p2p-spec"
	"github.com/NethermindEth/juno/utils"
	"github.com/libp2p/go-libp2p/core/network"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	io.Closer

	stream network.Stream
	log    utils.Logger
}

func NewClient(stream network.Stream, log utils.Logger) *Client {
	return &Client{
		stream: stream,
		log:    log,
	}
}

func (c *Client) Close() error {
	return c.stream.Close()
}

func (c *Client) send(msg proto.Message) error {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = c.stream.Write(msgBytes)
	return err
}

func (c *Client) receive(res proto.Message) error {
	buffer := getBuffer()
	defer bufferPool.Put(buffer)

	_, err := buffer.ReadFrom(c.stream)
	if err != nil {
		return err
	}

	return proto.Unmarshal(buffer.Bytes(), res)
}

func (c *Client) SendNewBlock(newBlock *p2p_spec.NewBlock) error {
	return c.send(newBlock)
}
