package jsonrpc_test

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"testing"

	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"nhooyr.io/websocket"
)

// The caller is responsible for closing the connection.
func testConnection(t *testing.T, rpc *jsonrpc.Server) *websocket.Conn {
	l, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	ctx := context.Background()

	ws := jsonrpc.NewWebsocket(l, rpc, utils.NewNopZapLogger())
	go func() {
		require.NoError(t, ws.Run(context.Background()))
	}()

	remote := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("localhost:%d", l.Addr().(*net.TCPAddr).Port),
	}
	conn, resp, err := websocket.Dial(ctx, remote.String(), nil) //nolint:bodyclose // websocket package closes resp.Body for us.
	require.NoError(t, err)
	require.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode)

	return conn
}

func TestHandler(t *testing.T) {
	rpc := jsonrpc.NewServer(utils.NewNopZapLogger())
	method := jsonrpc.Method{
		Name:   "test_echo",
		Params: []jsonrpc.Parameter{{Name: "msg"}},
		Handler: func(msg string) (string, *jsonrpc.Error) {
			return msg, nil
		},
	}
	require.NoError(t, rpc.RegisterMethod(method))
	conn := testConnection(t, rpc)

	msg := `{"jsonrpc" : "2.0", "method" : "test_echo", "params" : [ "abc123" ], "id" : 1}`
	err := conn.Write(context.Background(), websocket.MessageText, []byte(msg))
	require.NoError(t, err)

	want := `{"jsonrpc":"2.0","result":"abc123","id":1}`
	_, got, err := conn.Read(context.Background())
	require.NoError(t, err)
	assert.Equal(t, want, string(got))

	require.NoError(t, conn.Close(websocket.StatusNormalClosure, ""))
}

func TestWebsocketSubscribeUnsubscribe(t *testing.T) {
	rpc := jsonrpc.NewServer(utils.NewNopZapLogger())
	id := 1
	rpc.WithIDGen(func() (uint64, error) {
		return 1, nil
	})
	values := []int{1, 2, 3}
	submethod := jsonrpc.SubscribeMethod{
		Name: "test_subscribe",
		Handler: func(subServer *jsonrpc.SubscriptionServer) (event.Subscription, *jsonrpc.Error) {
			return event.NewSubscription(func(quit <-chan struct{}) error {
				for _, value := range values {
					select {
					case <-quit:
						return nil
					default:
						if err := subServer.Send(value); err != nil {
							return err
						}
					}
				}
				<-quit
				return nil
			}), nil
		},
		UnsubMethodName: "test_unsubscribe",
	}
	rpc.RegisterSubscribeMethod(submethod)
	conn := testConnection(t, rpc)

	// Initial subscription handshake.
	req := `{"jsonrpc": "2.0", "method": "test_subscribe", "params": [], "id": 1}`
	wswrite(t, conn, req)
	want := fmt.Sprintf(`{"jsonrpc":"2.0","result":%d,"id":1}`, id)
	got := wsread(t, conn)
	require.Equal(t, want, got)

	// All values are received.
	for _, value := range values {
		want = fmt.Sprintf(`{"jsonrpc":"2.0","method":"test_subscribe","params":{"subscription":%d,"result":%d}}`, id, value)
		got = wsread(t, conn)
		assert.Equal(t, want, got)
	}

	// Unsubscribe.
	req = `{"jsonrpc": "2.0", "method": "test_unsubscribe", "params": [%d], "id": 1}`
	wswrite(t, conn, fmt.Sprintf(req, id))
	want = `{"jsonrpc":"2.0","result":true,"id":1}`
	got = wsread(t, conn)
	assert.Equal(t, want, got)

	require.NoError(t, conn.Close(websocket.StatusNormalClosure, ""))
}

func wswrite(t *testing.T, conn *websocket.Conn, req string) {
	err := conn.Write(context.Background(), websocket.MessageText, []byte(req))
	require.NoError(t, err)
}

func wsread(t *testing.T, conn *websocket.Conn) string {
	_, res, err := conn.Read(context.Background())
	require.NoError(t, err)
	return string(res)
}
