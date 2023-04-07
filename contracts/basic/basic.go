/*
Copyright 2023 The Bestchains Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package basic

import (
	"github.com/bestchains/bestchains-contracts/contracts/access"
	"github.com/bestchains/bestchains-contracts/library/context"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var _ IBasic = new(BasicContract)

// BasicContract provides simple key-value Get/Put
type BasicContract struct {
	contractapi.Contract
	// Ownable
	access.IOwnable
}

func NewBasicContract(ownable access.IOwnable) *BasicContract {
	basicContract := new(BasicContract)
	basicContract.Name = "org.bestchains.com.BasicContract"
	basicContract.IOwnable = ownable
	basicContract.TransactionContextHandler = new(context.Context)
	basicContract.BeforeTransaction = context.BeforeTransaction

	return basicContract
}

func (bc *BasicContract) PutValue(ctx context.ContextInterface, key string, value string) error {

	return ctx.GetStub().PutState(key, []byte(value))
}
func (bc *BasicContract) GetValue(ctx context.ContextInterface, key string) (string, error) {
	bytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
