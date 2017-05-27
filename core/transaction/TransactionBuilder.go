package transaction

import (
	"DNA/common"
	"DNA/core/asset"
	"DNA/core/contract/program"
	"DNA/core/transaction/payload"
	"DNA/crypto"
	"DNA/core/code"
)

//initial a new transaction with asset registration payload
func NewRegisterAssetTransaction(asset *asset.Asset, amount common.Fixed64, issuer *crypto.PubKey, conroller common.Uint160) (*Transaction, error) {

	//TODO: check arguments

	assetRegPayload := &payload.RegisterAsset{
		Asset:  asset,
		Amount: amount,
		//Precision: precision,
		Issuer:     issuer,
		Controller: conroller,
	}

	return &Transaction{
		//nonce uint64 //TODO: genenrate nonce
		UTXOInputs:    []*UTXOTxInput{},
		BalanceInputs: []*BalanceTxInput{},
		Attributes:    []*TxAttribute{},
		TxType:        RegisterAsset,
		Payload:       assetRegPayload,
		Programs:      []*program.Program{},
	}, nil
}

func NewIssueAssetTransaction(outputs []*TxOutput) (*Transaction, error) {

	assetRegPayload := &payload.IssueAsset{}

	return &Transaction{
		TxType:        IssueAsset,
		Payload:       assetRegPayload,
		Attributes:    []*TxAttribute{},
		BalanceInputs: []*BalanceTxInput{},
		Outputs:       outputs,
		Programs:      []*program.Program{},
	}, nil
}

func NewTransferAssetTransaction(inputs []*UTXOTxInput, outputs []*TxOutput) (*Transaction, error) {

	//TODO: check arguments

	assetRegPayload := &payload.TransferAsset{}

	return &Transaction{
		TxType:        TransferAsset,
		Payload:       assetRegPayload,
		Attributes:    []*TxAttribute{},
		UTXOInputs:    inputs,
		BalanceInputs: []*BalanceTxInput{},
		Outputs:       outputs,
		Programs:      []*program.Program{},
	}, nil
}

//initial a new transaction with record payload
func NewRecordTransaction(recordType string, recordData []byte) (*Transaction, error) {
	//TODO: check arguments
	recordPayload := &payload.Record{
		RecordType: recordType,
		RecordData: recordData,
	}

	return &Transaction{
		TxType:        Record,
		Payload:       recordPayload,
		Attributes:    []*TxAttribute{},
		UTXOInputs:    []*UTXOTxInput{},
		BalanceInputs: []*BalanceTxInput{},
		Programs:      []*program.Program{},
	}, nil
}

//initial a new transaction with publish payload
func NewPublishTransaction(fc *code.FunctionCode, ic []byte, name string,codeversion string, author string,email string,desp string) (*Transaction, error) {
	//TODO: check arguments
	DeployCodePayload := &payload.DeployCode{
		Code:        fc,
		InitCode:    ic,
		Name:        name,
		CodeVersion: codeversion,
		Author:      author,
		Email:       email,
		Description: desp,
	}

	return &Transaction{
		TxType:        DeployCode,
		Payload:       DeployCodePayload,
		Attributes:    []*TxAttribute{},
		UTXOInputs:    []*UTXOTxInput{},
		BalanceInputs: []*BalanceTxInput{},
		Programs:      []*program.Program{},
	}, nil
}


//initial a new transaction with invoke payload
func NewInvokeTransaction(fc []byte) (*Transaction, error) {
	//TODO: check arguments
	InvokeCodePayload := &payload.InvokeCode{
		Code:        fc,
	}

	return &Transaction{
		TxType:        InvokeCode,
		Payload:       InvokeCodePayload,
		Attributes:    []*TxAttribute{},
		UTXOInputs:    []*UTXOTxInput{},
		BalanceInputs: []*BalanceTxInput{},
		Programs:      []*program.Program{},
	}, nil
}