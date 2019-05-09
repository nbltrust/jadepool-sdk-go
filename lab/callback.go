package main

import (
	"fmt"

	"github.com/nbltrust/jadepool-sdk-go/api"
	"github.com/nbltrust/jadepool-sdk-go/types"
)

const order1 string = `{"code":0,"status":0,"message":"OK","crypto":"ecc","timestamp":1557385283017,"sig":{"r":"DTvEoSAy43FwyKDxfae8GziRGE3ew7A2POhWOBkin1s=","s":"OpE0/nIQojjxNA29uB8T4mA+vK5wyPQ6QCCS3GkgmPk=","v":28},"result":{"_id":"5cd2ac1eb4a0e400a123557d","id":"1","coinName":"ETH","txid":"0x508dfe84a257dfda49b88f122e8aafca79b6765517944074ff2b252a745e24d2","meta":null,"state":"done","bizType":"WITHDRAW","type":"ETH","coinType":"ETH","from":"0xb42edcca0fa411e972d172d7f5d18c59ea9084ad","to":"0xC3ABEBBAEf594f9ceCa420e3Bad4a45D457f60fa","value":"0.001","sequence":1,"confirmations":74247,"create_at":1557310494641,"update_at":1557381905938,"actionArgs":[],"actionResults":[],"n":0,"fee":"0.00021","fees":[{"_id":"5cd2ac36b4a0e400a1235598","amount":"0.00021","coinName":"ETH","nativeAmount":"0","nativeName":""}],"data":{"timestampBegin":1557310497919,"timestampFinish":1557310518217,"nonce":26,"type":"Ethereum","hash":"0x508dfe84a257dfda49b88f122e8aafca79b6765517944074ff2b252a745e24d2","blockNumber":8285075,"blockHash":"0x00df94b6bddf746228b1e509c0313472140c712679516f25498deb0348d2fc1f","from":[{"address":"0xb42edcca0fa411e972d172d7f5d18c59ea9084ad","value":"0.001","asset":"ETH"}],"to":[{"address":"0xC3ABEBBAEf594f9ceCa420e3Bad4a45D457f60fa","value":"0.001","asset":"ETH"}],"fee":"0.00021","confirmations":74247,"state":"done"},"hash":"0x508dfe84a257dfda49b88f122e8aafca79b6765517944074ff2b252a745e24d2","block":8285075,"extraData":"","memo":"","sendAgain":false,"namespace":"ETH","sid":"fDJXJg6hEOpSHuncAAAE"}}`
const audit1 string = `{"code":0,"status":0,"message":"OK","crypto":"ecc","timestamp":1557385458209,"sig":{"r":"OLOEz4fxUxTqZ82TX3sStiq4r0RYBK7iVvGlalXCMYQ=","s":"LJ37xGP17XNSNHzMCF5LcH7xmKmtzUjcBYOIcrq2JQU=","v":28},"result":{"current":{"id":"5cd3d0f260e73f0f534b163f","type":"ETH","blocknumber":8359463,"timestamp":1557385458209,"deposit_total":"0","withdraw_total":"0","fee_total":"0","internal_fee":"0","internal_num":0},"calculated":true,"deposit_total":"0","deposit_num":0,"withdraw_total":"0","withdraw_num":0,"sweep_total":"0","sweep_num":0,"sweep_internal_total":"0","sweep_internal_num":0,"airdrop_total":"0","airdrop_num":0,"recharge_total":"0","recharge_num":0,"recharge_internal_total":"0","recharge_internal_num":0,"recharge_unknown_total":"0","recharge_unknown_num":0,"recharge_special_total":"0","recharge_special_num":0,"failed_withdraw_num":0,"failed_sweep_num":0,"failed_sweep_internal_num":0,"fees":[],"type":"ETH","timestamp":1557385458209,"blocknumber":8359463,"appid":"pri","create_at":"2019-05-09T07:04:18.256Z","update_at":"2019-05-09T07:04:23.349Z","__v":0,"last":"5cd3c0402ac13d0814364029","calc_order_num":0,"id":"5cd3d0f260e73f0f534b163f"}}`

func main() {
	jp, _ := api.InitWith(types.JPConfig{
		JadePubKey: "04ace32532c90652e1bae916248e427a7ab10aeeea1067949669a3f4da10965ef90d7297f538f23006a31f94fdcfaed9e8dd38c85ba7e285f727430332925aefe5",
	})
	o1 := []byte(order1)
	order, err := jp.OrderCallback(o1)
	fmt.Println(order, err)

	a1 := []byte(audit1)
	audit, err := jp.AuditCallback(a1)
	fmt.Printf("%+v,%v", audit, err)
}
