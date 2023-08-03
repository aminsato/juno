//go:generate protoc --go_out=./ --proto_path=./ --go_opt=Mp2p/proto/requests.proto=./p2p-spec --go_opt=Mp2p/proto/transaction.proto=./p2p-spec --go_opt=Mp2p/proto/state.proto=./p2p-spec --go_opt=Mp2p/proto/snapshot.proto=./p2p-spec --go_opt=Mp2p/proto/receipt.proto=./p2p-spec --go_opt=Mp2p/proto/mempool.proto=./p2p-spec --go_opt=Mp2p/proto/event.proto=./p2p-spec --go_opt=Mp2p/proto/block.proto=./p2p-spec --go_opt=Mp2p/proto/common.proto=./p2p-spec p2p/proto/transaction.proto p2p/proto/state.proto p2p/proto/snapshot.proto p2p/proto/common.proto p2p/proto/block.proto p2p/proto/event.proto p2p/proto/receipt.proto p2p/proto/requests.proto
package p2p

import (
	"bytes"
	"errors"
	"sync"

	"github.com/NethermindEth/juno/blockchain"
	p2p_spec "github.com/NethermindEth/juno/p2p/p2p-spec"
	"github.com/NethermindEth/juno/utils"
	"github.com/libp2p/go-libp2p/core/network"
	"google.golang.org/protobuf/proto"
)

type Handler struct {
	bcReader blockchain.Reader
	log      utils.Logger
}

func NewHandler(bcReader blockchain.Reader, log utils.Logger) *Handler {
	return &Handler{
		bcReader: bcReader,
		log:      log,
	}
}

func (h *Handler) NewBlock(msg *p2p_spec.NewBlock) {
	h.log.Infof("Got NewBlock", msg.String())
}

// bufferPool caches unused buffer objects for later reuse.
var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func getBuffer() *bytes.Buffer {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	return buffer
}

func (h *Handler) StreamHandler(stream network.Stream) {
	defer func() {
		if err := stream.Close(); err != nil {
			h.log.Debugw("error closing stream", stream.ID(), stream.Protocol(), err)
		}
	}()

	buffer := getBuffer()
	defer bufferPool.Put(buffer)

	_, err := buffer.ReadFrom(stream)
	if err != nil {
		h.log.Debugw("error reading from stream", stream.ID(), stream.Protocol(), err)
		return
	}

	var req p2p_spec.Request
	if err := proto.Unmarshal(buffer.Bytes(), &req); err != nil {
		h.log.Debugw("error unmarshaling message", stream.ID(), stream.Protocol(), err)
		return
	}

	response, err := h.reqHandler(&req)
	if err != nil {
		h.log.Debugw("error handling request", stream.ID(), stream.Protocol(), req.String(), err)
		return
	}

	responseBytes, err := proto.Marshal(response)
	if err != nil {
		h.log.Debugw("error marshaling response", stream.ID(), stream.Protocol(), response, err)
		return
	}

	_, err = stream.Write(responseBytes)
	if err != nil {
		h.log.Debugw("error writing response", stream.ID(), stream.Protocol(), err)
		return
	}
}

func (h *Handler) reqHandler(req *p2p_spec.Request) (proto.Message, error) {
	switch req.GetReq().(type) {
	case *p2p_spec.Request_NewBlock:
		h.NewBlock(req.GetNewBlock())
		return nil, nil
	default:
		return nil, errors.New("unhandled request")
	}
}
