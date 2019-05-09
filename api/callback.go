package api

import (
	"encoding/json"
	"fmt"

	"github.com/nbltrust/jadepool-sdk-go/types"
	"github.com/nbltrust/jadepool-sdk-go/utils"
)

// OrderCallback ...
func (jp *JPApi) OrderCallback(callback []byte) (*types.JPOrderResult, error) {
	data := types.JPEvent{}
	err := json.Unmarshal(callback, &data)
	if err != nil {
		return nil, err
	}
	err = jp.CheckComing(&data)
	if err != nil {
		fmt.Println(err)
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

// AuditCallback ...
func (jp *JPApi) AuditCallback(callback []byte) (*types.JPAudit, error) {
	data := types.JPEvent{}
	err := json.Unmarshal(callback, &data)
	if err != nil {
		return nil, err
	}
	err = jp.CheckComing(&data)
	if err != nil {
		fmt.Println(err)
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
