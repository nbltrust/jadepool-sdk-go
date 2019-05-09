package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	ecc "github.com/nbltrust/jadepool-golib-crypto/crypto"
	"github.com/nbltrust/jadepool-sdk-go/types"
	"github.com/nbltrust/jadepool-sdk-go/utils"
)

// JPApi ...
type JPApi struct {
	Config types.JPConfig
}

// InitWith ...
func InitWith(config types.JPConfig) (*JPApi, error) {
	api := &JPApi{
		Config: config,
	}
	return api, nil
}

// Withdraw ...
func (jp *JPApi) Withdraw(sequence uint, coinID string, value string, to string, memo string, extraData string) (*types.JPWithdrawResult, error) {
	requestObj := &types.JPWithdrawRequest{}
	requestObj.Type = coinID
	requestObj.To = to
	requestObj.Value = value
	requestObj.Sequence = sequence
	requestObj.Memo = memo
	requestObj.ExtraData = extraData
	// 构造ecc部分
	timestamp := time.Now().Unix() * 1000
	requestObj.Timestamp = timestamp
	sendData, err := jp.sendDataEcc(requestObj, timestamp)
	if err != nil {
		return nil, err
	}
	data := types.JPEvent{}
	// fmt.Println(sendData.Data)
	err = jp.postResult("/api/v1/transactions", sendData, &data)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", data)
	err = jp.CheckComing(&data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := types.JPWithdrawResult{}
	err = utils.ResultToStruct(data.Result, &result)
	if err != nil {
		// log.Errorln(err)
		return nil, err
	}
	return &result, err
}

// NewAddress ...
func (jp *JPApi) NewAddress(coinType string) (*types.JPAddressResult, error) {
	requestObj := &types.JPAddressRequest{}
	requestObj.Type = coinType
	// 构造ecc部分
	timestamp := time.Now().Unix() * 1000
	requestObj.Timestamp = timestamp
	sendData, err := jp.sendDataEcc(requestObj, timestamp)
	if err != nil {
		return nil, err
	}
	data := types.JPEvent{}
	// fmt.Println(sendData.Data)
	err = jp.postResult("/api/v1/addresses/new", sendData, &data)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", data)
	err = jp.CheckComing(&data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := types.JPAddressResult{}
	err = utils.ResultToStruct(data.Result, &result)
	if err != nil {
		// log.Errorln(err)
		return nil, err
	}
	return &result, err
}

// VerifyAddress ...
func (jp *JPApi) VerifyAddress(coinType string, address string) (*types.VerifyRes, error) {
	requestObj := &types.JPAddressRequest{}
	requestObj.Type = coinType
	// requestObj.Address = address
	// 构造ecc部分
	timestamp := time.Now().Unix() * 1000
	requestObj.Timestamp = timestamp
	sendData, err := jp.sendDataEcc(requestObj, timestamp)
	if err != nil {
		return nil, err
	}
	data := types.JPEvent{}
	// fmt.Printf("%+v\n", sendData)
	err = jp.postResult("/api/v1/addresses/"+address+"/verify", sendData, &data)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", data)
	err = jp.CheckComing(&data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := types.VerifyRes{}
	err = utils.ResultToStruct(data.Result, &result)
	if err != nil {
		// log.Errorln(err)
		return nil, err
	}
	return &result, err
}

// AskAudit ...
func (jp *JPApi) AskAudit(coinType string, auditTime uint) (string, error) {
	requestObj := &types.JPAuditRequest{}
	requestObj.Type = coinType
	requestObj.AuditTime = auditTime
	// requestObj.Address = address
	// 构造ecc部分
	timestamp := time.Now().Unix() * 1000
	requestObj.Timestamp = timestamp
	sendData, err := jp.sendDataEcc(requestObj, timestamp)
	if err != nil {
		return "", err
	}
	data := types.JPEvent{}
	// fmt.Printf("%+v\n", sendData)
	err = jp.postResult("/api/v1/audits", sendData, &data)
	if err != nil {
		return "", err
	}
	// fmt.Printf("%+v\n", data)
	err = jp.CheckComing(&data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	result := types.JPAuditResult{}
	err = utils.ResultToStruct(data.Result, &result)
	if err != nil {
		// log.Errorln(err)
		return "", err
	}
	return result.Current.ID, err
}

// QueryOrder ...
func (jp *JPApi) QueryOrder(orderID string) (*types.JPOrderResult, error) {
	requestObj := &types.JPEmptyRequest{}
	// requestObj.Address = address
	// 构造ecc部分
	timestamp := time.Now().Unix() * 1000
	requestObj.Timestamp = timestamp
	sendData, err := jp.urlDataEcc(requestObj, timestamp)
	if err != nil {
		return nil, err
	}

	data := types.JPEvent{}
	// fmt.Printf("data %+v\n", sendData)
	err = jp.getResult("/api/v1/transactions/"+orderID, sendData, &data)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", data)
	err = jp.CheckComing(&data)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	result := types.JPOrderResult{}
	err = utils.ResultToStruct(data.Result, &result)
	if err != nil {
		// log.Errorln(err)
		return nil, err
	}
	return &result, err
}

// QueryAudit ...
func (jp *JPApi) QueryAudit(auditID string) (*types.JPAudit, error) {
	requestObj := &types.JPEmptyRequest{}
	// requestObj.Address = address
	// 构造ecc部分
	timestamp := time.Now().Unix() * 1000
	requestObj.Timestamp = timestamp
	sendData, err := jp.urlDataEcc(requestObj, timestamp)
	if err != nil {
		return nil, err
	}

	data := types.JPEvent{}
	// fmt.Printf("data %+v\n", sendData)
	err = jp.getResult("/api/v1/audits/"+auditID, sendData, &data)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", data)
	err = jp.CheckComing(&data)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	result := types.JPAudit{}
	err = utils.ResultToStruct(data.Result, &result)
	if err != nil {
		// log.Errorln(err)
		return nil, err
	}
	return &result, err
}

// GetBalance ...
func (jp *JPApi) GetBalance(coinType string) (*types.JPBalance, error) {
	requestObj := &types.JPEmptyRequest{}
	// requestObj.Address = address
	// 构造ecc部分
	timestamp := time.Now().Unix() * 1000
	requestObj.Timestamp = timestamp
	sendData, err := jp.urlDataEcc(requestObj, timestamp)
	if err != nil {
		return nil, err
	}

	data := types.JPEvent{}
	// fmt.Printf("data %+v\n", sendData)
	err = jp.getResult(`/api/v1/wallet/`+coinType+`/status`, sendData, &data)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", data)
	err = jp.CheckComing(&data)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	result := types.JPBalance{}
	err = utils.ResultToStruct(data.Result, &result)
	if err != nil {
		// log.Errorln(err)
		return nil, err
	}
	return &result, err
}

// AuthCoin ...
func (jp *JPApi) AuthCoin(coinID string, coinType string, chain string, token string, decimal int, contract string) {

}
func (jp *JPApi) getResult(urlPath string, sendData url.Values, v interface{}) (err error) {
	// fmt.Println(string(bs[:]))
	jadepoolAddr := jp.Config.URL
	url := jadepoolAddr + urlPath + "?" + sendData.Encode()
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("post error: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("ReadAll error: %v", err)
	}
	err = json.Unmarshal(bodyBytes, v)
	if err != nil {
		return fmt.Errorf("Unmarshal error: %v, body: %s", err, string(bodyBytes))
	}
	return nil
}
func (jp *JPApi) postResult(urlPath string, sendData *types.JPSendData, v interface{}) (err error) {
	bs, _ := json.Marshal(sendData)
	// fmt.Println(string(bs[:]))
	jadepoolAddr := jp.Config.URL
	url := jadepoolAddr + urlPath
	resp, err := http.Post(url, "application/json", bytes.NewReader(bs))
	if err != nil {
		return fmt.Errorf("post error: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("ReadAll error: %v", err)
	}
	err = json.Unmarshal(bodyBytes, v)
	if err != nil {
		return fmt.Errorf("Unmarshal error: %v, body: %s", err, string(bodyBytes))
	}
	return nil
}

//CheckComing ...
func (jp *JPApi) CheckComing(data *types.JPEvent) (err error) {
	// checkEcc
	if data.Code != 0 || data.Status != 0 || data.Result == nil {
		return fmt.Errorf("BN request failed, data: %#v", data)
	}
	// verify sig
	if _, ok := data.Result["timestamp"]; !ok {
		//do something here
		// fmt.Println(">>>>>>>>>>>>>>>>>")
		data.Result["timestamp"] = data.Timestamp
	}
	pubKey := jp.Config.JadePubKey
	// bs, _ := json.Marshal(data.Result)
	// fmt.Println(string(bs[:]))
	// fmt.Println(data.Result, data.Sig)
	ok, err := ecc.VerifyECCSign(data.Result, data.Sig, pubKey)
	if err != nil {
		return fmt.Errorf("verifySign error: %v, data: %#v", err, data)
	}
	if !ok {
		return fmt.Errorf("verify result: %v, data: %#v", ok, data)
	}
	return nil
}

func (jp *JPApi) sendDataEcc(data interface{}, timestamp int64) (sendData *types.JPSendData, err error) {
	sendData = &types.JPSendData{}
	sendData.Crypto = "ecc"
	sendData.Encode = "base64" //jp.Config.EccKeyEncoder
	sendData.Timestamp = timestamp
	sendData.Hash = "sha3"
	sendData.AppID = jp.Config.AppID
	sendData.Data = data
	//签名
	priKey := jp.Config.EccKey
	sig, err := ecc.SignECCData(priKey, data)
	if err != nil {
		return nil, fmt.Errorf("SignECCData error: %v", err)
	}
	sendData.Sig = sig
	return sendData, nil
}

func (jp *JPApi) urlDataEcc(data interface{}, timestamp int64) (sendData url.Values, err error) {
	sendData = url.Values{}
	sendData.Add("crypto", "ecc")
	sendData.Add("encode", "base64")
	sendData.Add("timestamp", fmt.Sprintf("%d", timestamp))
	sendData.Add("hash", "sha3")
	sendData.Add("appid", jp.Config.AppID)
	//签名
	priKey := jp.Config.EccKey
	sig, err := ecc.SignECCData(priKey, data)
	if err != nil {
		return nil, fmt.Errorf("SignECCData error: %v", err)
	}
	sendData.Add("sig_r", sig.R)
	sendData.Add("sig_s", sig.S)
	sendData.Add("sig_v", fmt.Sprintf("%d", sig.V))
	return sendData, nil
}
