package types

import (
	"time"

	ecc "github.com/nbltrust/jadepool-golib-crypto/crypto"
)

// JPOrderResult ...
type JPOrderResult struct {
	ID            string      `json:"id"`
	CoinName      string      `json:"coinName"`
	Txid          string      `json:"txid"`
	Meta          interface{} `json:"meta"`
	State         string      `json:"state"`
	BizType       string      `json:"bizType"`
	Type          string      `json:"type"`
	SubType       string      `json:"subType"`
	CoinType      string      `json:"coinType"`
	To            string      `json:"to"`
	Value         string      `json:"value"`
	Confirmations int         `json:"confirmations"`
	CreateAt      int64       `json:"create_at"`
	UpdateAt      int64       `json:"update_at"`
	From          string      `json:"from"`
	N             int         `json:"n"`
	Fee           string      `json:"fee"`
	Hash          string      `json:"hash"`
	Block         int         `json:"block"`
	ExtraData     string      `json:"extraData"`
	Memo          string      `json:"memo"`
	SendAgain     bool        `json:"sendAgain"`
}

// JPEvent ...
type JPEvent struct {
	Code      int                    `json:"code"`
	Status    int                    `json:"status"`
	Message   string                 `json:"message"`
	Crypto    string                 `json:"crypto"`
	Timestamp int64                  `json:"timestamp"`
	Sig       *ecc.ECCSig            `json:"sig"`
	Result    map[string]interface{} `json:"result"`
}

// JPAddressResult ...
type JPAddressResult struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}

// JPWithdrawResult ...
type JPWithdrawResult JPOrderResult

// JPAddressRequest ...
type JPAddressRequest struct {
	Type string `json:"type"`
	// Address   string `json:"address"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Callback  string `json:"callback,omitempty"`
}

// JPEmptyRequest ...
type JPEmptyRequest struct {
	// Address   string `json:"address"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Callback  string `json:"callback,omitempty"`
}

// JPAuditRequest ...
type JPAuditRequest struct {
	Type      string `json:"type"`
	AuditTime uint   `json:"audittime"`
	// Address   string `json:"address"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Callback  string `json:"callback,omitempty"`
}

// JPWithdrawRequest ...
type JPWithdrawRequest struct {
	Type      string `json:"type"`
	To        string `json:"to"`
	Value     string `json:"value"`
	Sequence  uint   `json:"sequence"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Memo      string `json:"memo"`
	ExtraData string `json:"extraData"`
}

// JPSendData ...
type JPSendData struct {
	Crypto    string      `json:"crypto"`
	Hash      string      `json:"hash"`
	Encode    string      `json:"encode"`
	AppID     string      `json:"appid"`
	Timestamp int64       `json:"timestamp"`
	Sig       *ecc.ECCSig `json:"sig"`
	Data      interface{} `json:"data"`
}

// VerifyRes ...
type VerifyRes struct {
	Address   string `json:"address"`
	Asset     string `json:"asset"`
	Timestamp uint   `json:"timestamp"`
	Valid     bool   `json:"valid"`
}

// JPAuditResult ...
type JPAuditResult struct {
	Current struct {
		ID string `json:"id"`
	} `json:"current"`
}

// JPAudit ...
type JPAudit struct {
	Calculated             bool   `json:"calculated"`
	DepositTotal           string `json:"deposit_total"`
	DepositNum             int    `json:"deposit_num"`
	WithdrawTotal          string `json:"withdraw_total"`
	WithdrawNum            int    `json:"withdraw_num"`
	SweepTotal             string `json:"sweep_total"`
	SweepNum               int    `json:"sweep_num"`
	SweepInternalTotal     string `json:"sweep_internal_total"`
	SweepInternalNum       int    `json:"sweep_internal_num"`
	AirdropTotal           string `json:"airdrop_total"`
	AirdropNum             int    `json:"airdrop_num"`
	RechargeTotal          string `json:"recharge_total"`
	RechargeNum            int    `json:"recharge_num"`
	RechargeInternalTotal  string `json:"recharge_internal_total"`
	RechargeInternalNum    int    `json:"recharge_internal_num"`
	RechargeUnknownTotal   string `json:"recharge_unknown_total"`
	RechargeUnknownNum     int    `json:"recharge_unknown_num"`
	RechargeSpecialTotal   string `json:"recharge_special_total"`
	RechargeSpecialNum     int    `json:"recharge_special_num"`
	FailedWithdrawNum      int    `json:"failed_withdraw_num"`
	FailedSweepNum         int    `json:"failed_sweep_num"`
	FailedSweepInternalNum int    `json:"failed_sweep_internal_num"`
	Fees                   []struct {
		WithdrawFee            string `json:"withdraw_fee"`
		SweepFee               string `json:"sweep_fee"`
		SweepInternalFee       string `json:"sweep_internal_fee"`
		FailedWithdrawFee      string `json:"failed_withdraw_fee"`
		FailedSweepFee         string `json:"failed_sweep_fee"`
		FailedSweepInternalFee string `json:"failed_sweep_internal_fee"`
		ID                     string `json:"_id"`
		FeeType                string `json:"fee_type"`
	} `json:"fees"`
	Type         string    `json:"type"`
	Timestamp    int64     `json:"timestamp"`
	Blocknumber  int       `json:"blocknumber"`
	Appid        string    `json:"appid"`
	CreateAt     time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
	V            int       `json:"__v"`
	CalcOrderNum int       `json:"calc_order_num"`
	Last         string    `json:"last"`
	ID           string    `json:"id"`
}

// JPBalance ...
type JPBalance struct {
	Balance            string `json:"balance"`
	BalanceAvailable   string `json:"balanceAvailable"`
	BalanceUnavailable string `json:"balanceUnavailable"`
}
