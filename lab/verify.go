package main

import (
	"encoding/json"
	"fmt"

	"github.com/nbltrust/jadepool-sdk-go/api"
	"github.com/nbltrust/jadepool-sdk-go/types"
)

const re string = `{"code":0,"status":0,"message":"OK","crypto":"ecc","timestamp":1557307620000,"sig":{"r":"wS058aaoD0BLmw7RF7Fvgj+1/6LGnQcQeeF/vj2CthI=","s":"TtXeNpISv4mZ5U49nJnN0YNSa11zbFmXPjqAue46T0Q=","v":28},"result":{"calculated":true,"deposit_total":"0","deposit_num":0,"withdraw_total":"0.001","withdraw_num":1,"sweep_total":"0","sweep_num":0,"sweep_internal_total":"0","sweep_internal_num":0,"airdrop_total":"0","airdrop_num":0,"recharge_total":"0","recharge_num":0,"recharge_internal_total":"0","recharge_internal_num":0,"recharge_unknown_total":"0","recharge_unknown_num":0,"recharge_special_total":"0","recharge_special_num":0,"failed_withdraw_num":0,"failed_sweep_num":0,"failed_sweep_internal_num":0,"fees":[{"withdraw_fee":"0.00021","sweep_fee":"0","sweep_internal_fee":"0","failed_withdraw_fee":"0","failed_sweep_fee":"0","failed_sweep_internal_fee":"0","_id":"5cd3a02cb4a0e400a1238865","fee_type":"ETH"}],"type":"ETH","timestamp":1557307620000,"blocknumber":10909989,"appid":"bbb","create_at":"2019-05-08T10:30:38.623Z","update_at":"2019-05-09T03:36:12.009Z","__v":17,"calc_order_num":1,"last":"5cd3a02cb4a0e400a1238864","id":"5cd2afceb4a0e400a1235931"}}`
const re1 string = `{"code":0,"status":0,"message":"OK","crypto":"ecc","timestamp":1557373751666,"sig":{"r":"4j/+HQGmWvuBdt0YaKqCkGH6dk38UXALy67tlR05sls=","s":"RDJ7Gp63X+nvYIgTPU3sdNKAuBpxKUJaMvKmICLliQU=","v":28},"result":{"type":"ETH","current":{"id":"5cd2afceb4a0e400a1235931","type":"ETH","blocknumber":10909989,"timestamp":1557307620000},"last":null,"namespace":"ETH","sid":"fDJXJg6hEOpSHuncAAAE"}}`
const re2 string = `{"code":0,"status":0,"message":"OK","crypto":"ecc","timestamp":1557380342000,"sig":{"r":"X2BcfvTGzbKvxXfHO1fNc2Xofzd8Uls+vGtmkAKiN+k=","s":"ZkMt89VRWgVq7WSYBk3IETJ7OslGJVxxfdRIxfayejY=","v":28},"result":{"current":{"id":"5cd3bd56ed62f505b359d2c0","type":"NASH","blocknumber":8354383,"timestamp":1557380342000,"deposit_total":"0","withdraw_total":"0","fee_total":"0","internal_fee":"0","internal_num":0},"calculated":true,"deposit_total":"0","deposit_num":0,"withdraw_total":"0","withdraw_num":0,"sweep_total":"0","sweep_num":0,"sweep_internal_total":"0","sweep_internal_num":0,"airdrop_total":"0","airdrop_num":0,"recharge_total":"0","recharge_num":0,"recharge_internal_total":"0","recharge_internal_num":0,"recharge_unknown_total":"0","recharge_unknown_num":0,"recharge_special_total":"0","recharge_special_num":0,"failed_withdraw_num":0,"failed_sweep_num":0,"failed_sweep_internal_num":0,"fees":[],"type":"NASH","timestamp":1557380342000,"blocknumber":8354383,"appid":"test","create_at":"2019-05-09T05:40:38.477Z","update_at":"2019-05-09T05:40:45.087Z","__v":0,"calc_order_num":0,"last":"5cd3bd43ed62f505b359d27e","id":"5cd3bd56ed62f505b359d2c0"}}`
const re3 string = `{"code":0,"status":0,"message":"OK","crypto":"ecc","timestamp":1557380968731,"sig":{"r":"ao0CEphZzA/Ig34pvZF22zNrfBfPuOkAQ37kfr2OARI=","s":"VsJ62S5E/k9hwXw+soZEEdA4QN5OTJoUSd2gsq7+uL0=","v":27},"result":{"id":"5cd3bf682ac13d0814363e37","type":"NASH","blocknumber":8354965,"timestamp":1557380968731,"namespace":"ETH","sid":"Eki2-HQNhYHqeCOUAAAF"}}`

func main() {
	jp, _ := api.InitWith(types.JPConfig{
		AppID:          "bbb",
		URL:            "http://127.0.0.1:7001",
		JadePubKey:     "04ace32532c90652e1bae916248e427a7ab10aeeea1067949669a3f4da10965ef90d7297f538f23006a31f94fdcfaed9e8dd38c85ba7e285f727430332925aefe5",
		EccKey:         "bf12996feeaa2977b6ca0d33a0e8bd2ccfc4844c6f8a7e6d15c099f8da4a255d",
		EccKeyEncoder:  "hex",
		AuthKey:        "",
		AuthKeyEncoder: "",
		Language:       "cn",
	})
	event := &types.JPEvent{}
	json.Unmarshal([]byte(re3), event)
	// fmt.Println(event)
	err := jp.CheckComing(event)
	fmt.Println(err)
}
