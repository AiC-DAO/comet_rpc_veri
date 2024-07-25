package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/cometbft/cometbft/rpc/jsonrpc/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/gogoproto/proto"
)

const testJsonStr = `
					{
						"jsonrpc": "2.0",
						"id": 0,
						"method": "broadcast_tx_sync",
						"params": {
							"tx_bytes":"Cv0BCvoBCikvaWJjLmFwcGxpY2F0aW9ucy50cmFuc2Zlci52MS5Nc2dUcmFuc2ZlchLMAQoIdHJhbnNmZXISDGNoYW5uZWwtMjQ5NBpPCkRpYmMvMjA4QjJGMTM3Q0RFNTEwQjQ0QzQxOTQ3QzA0NUNGREMyN0Y5OTZBOUQ5OTBFQTY0NDYwQkRENUIzREJFQjJFRBIHMTAwMDAwMCIrb3NtbzE0aDhmem53N3RuYTNmeXg0bGhydnp2a2FjMjhseWZrMHlmZ2s1cSorcGFzZzE0aDhmem53N3RuYTNmeXg0bGhydnp2a2FjMjhseWZrMDAyenUwZDIHCAIQp4O6BBJnClEKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiECPhXxDoDT/qErK5k9PJiHVbQc/0zysoUvwj2y1tsx6WQSBAoCCH8YyQMSEgoMCgV1b3NtbxIDNDY4EMLBCRpAUxPsgWQXuzK+OohdGx1HLRxwkTRkk2AjN0wO7ipkkXo70SFdy3mn8dbF3d6uShC3hemdtf9j7m1PqO0oyrRCpA==",
							"mode":"BROADCAST_MODE_SYNC"
						}
					}
					`
const testTxBytes = "Cv0BCvoBCikvaWJjLmFwcGxpY2F0aW9ucy50cmFuc2Zlci52MS5Nc2dUcmFuc2ZlchLMAQoIdHJhbnNmZXISDGNoYW5uZWwtMjQ5NBpPCkRpYmMvMjA4QjJGMTM3Q0RFNTEwQjQ0QzQxOTQ3QzA0NUNGREMyN0Y5OTZBOUQ5OTBFQTY0NDYwQkRENUIzREJFQjJFRBIHMTAwMDAwMCIrb3NtbzE0aDhmem53N3RuYTNmeXg0bGhydnp2a2FjMjhseWZrMHlmZ2s1cSorcGFzZzE0aDhmem53N3RuYTNmeXg0bGhydnp2a2FjMjhseWZrMDAyenUwZDIHCAIQp4O6BBJnClEKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiECPhXxDoDT/qErK5k9PJiHVbQc/0zysoUvwj2y1tsx6WQSBAoCCH8YyQMSEgoMCgV1b3NtbxIDNDY4EMLBCRpAUxPsgWQXuzK+OohdGx1HLRxwkTRkk2AjN0wO7ipkkXo70SFdy3mn8dbF3d6uShC3hemdtf9j7m1PqO0oyrRCpA=="

func main() {

	// 首先节点接收到jsonrpc请求，json解析jsonrpc请求为RPCRequest
	var request types.RPCRequest
	if err := json.Unmarshal([]byte(testJsonStr), &request); err != nil {
		print(err)
	}

	var m map[string]json.RawMessage
	if err := json.Unmarshal(request.Params, &m); err != nil {
		print(err)
	}

	// base64解析tx_bytes为字节数组decodedBytes
	decodedBytes, err := base64.StdEncoding.DecodeString(testTxBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// proto解析字节数组decodedBytes为tx
	tx := txtypes.Tx{}
	err = proto.Unmarshal(decodedBytes, &tx)
	if err != nil {
		log.Fatal("Unmarshal error: ", err)
	}

	println("done.")

}
