package ttk

import (
	"fmt"

	"github.com/ybbus/jsonrpc/v2"
)

type TTKConfig struct {
	Username  string
	Password  string
	Endpoint  string
	SSLVerify bool
	SessionID string
}

type TTKClient struct {
	cli jsonrpc.RPCClient
}

type TTKCmd struct {
	Command string
	Params  string
}

type TTKLogin struct {
	Username string
	Password string
}

type TTKResponse_Login struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  []TTKRLogin `json:"result"`
}

var (
	ttk *TTKConfig
)

type TTKRLogin struct {
	UbusRPCSession string `json:"ubus_rpc_session"`
	Timeout        int    `json:"timeout"`
	Expires        int    `json:"expires"`
	Acls           struct {
		AccessGroup struct {
			Superuser       []string `json:"superuser"`
			Unauthenticated []string `json:"unauthenticated"`
		} `json:"access-group"`
		Ubus struct {
			NAMING_FAILED []string `json:"*"`
			Session       []string `json:"session"`
		} `json:"ubus"`
		Uci struct {
			NAMING_FAILED []string `json:"*"`
		} `json:"uci"`
	} `json:"acls"`
	Data struct {
		Username string `json:"username"`
	} `json:"data"`
}

func Client(ttk *TTKConfig) (ttkC *TTKClient, err error) {

	rpcClient := jsonrpc.NewClient(ttk.Endpoint + "/ubus")
	ttkC = &TTKClient{
		cli: rpcClient,
	}

	var resp *jsonrpc.RPCResponse
	if resp, err = ttkC.cli.Call("call", "00000000000000000000000000000000", "session", "login", &TTKLogin{Username: ttk.Username, Password: ttk.Password}); err != nil {
		return nil, err
	}

	var TTKResponse TTKResponse_Login
	if err = resp.GetObject(&TTKResponse); err != nil {
		return nil, err
	}

	fmt.Print(TTKResponse.Result[0].UbusRPCSession)

	return ttkC, err
}
