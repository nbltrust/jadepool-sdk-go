# jadepool-sdk-go

## example

send request

see lab/apilab.go

```
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
	address, err := jp.NewAddress("ETH")
	fmt.Println(address, err)
	verify, err := jp.VerifyAddress("ETH", "0x7598ed8363d513975f7186f9d880f86febecb8aa")
	fmt.Println(verify, err)
	withdraw, err := jp.Withdraw(1, "ETH", "0.001", "0xC3ABEBBAEf594f9ceCa420e3Bad4a45D457f60fa", "", "")
	fmt.Println(withdraw, err)
	auditID, err := jp.AskAudit("ETH", 1557307620000)
	fmt.Println(auditID, err)
	audit, err := jp.QueryAudit(auditID)
	fmt.Println(audit, err)
	order, err := jp.QueryOrder("1")
	fmt.Println(order, err)
	balance, err := jp.GetBalance("ETH")
	fmt.Println(balance, err)
}

```

parse callback

see lab/callback.go

```
	jp, _ := api.InitWith(types.JPConfig{
		JadePubKey: "04ace32532c90652e1bae916248e427a7ab10aeeea1067949669a3f4da10965ef90d7297f538f23006a31f94fdcfaed9e8dd38c85ba7e285f727430332925aefe5",
	})
	o1 := []byte(order1)
	order, err := jp.OrderCallback(o1)
	fmt.Println(order, err)

	a1 := []byte(audit1)
	audit, err := jp.AuditCallback(a1)
	fmt.Printf("%+v,%v", audit, err)
```