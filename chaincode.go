/*
Copyright IBM Corp 2016 All Rights Reserved.
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

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var i int
//var entityPrefix = "entity:"
//Entity - Structure for an entity like Lender, Borrower
type Entity struct {
	Type   		  string  `json:"type"`
	Name    		string  `json:"name"`
	CashBalance float64 `json:"balance"`
	Bond 				float64 `json:"bond"`

}

//Product - Structure for products used in create and transfer
type Security struct {
	Name    string  `json:"name"`
	Entity  string  `json:"entity"`
	Value 	float64 `json:"value"`
}

//TxnLinkCollaterals - Entity transactions for linking collaterlas
type TxnCollaterals struct {
	Initiator string `json:"initiator"`
	Remarks   string `json:"remarks"`
	ID        string `json:"id"`
	Time      string `json:"time"`
	Value     float64 `json:"value"`
	Asset     string `json:"asset"`
}

//TxnTopup - Entity transactions for linking securities
type TxnSecurities struct {
	Initiator string `json:"initiator"`
	Remarks   string `json:"remarks"`
	ID        string `json:"id"`
	Time      string `json:"time"`
	Value     float64 `json:"value"`
	Security     string `json:"security"`
}

//TxnTransfer - Entity transactions for transfer of Product
type TxnBorrowTrade struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Offer  string `json:"offer"`
	Ask  string `json:"ask"`
	ID       string `json:"id"`
	Time     string `json:"time"`
	AssetValue      float64 `json:"assetValue"`
	Asset    string `json:"Asset"`
	SecurityQty float64 `json:"securityQty"`
	Security string `json:"security"`
	Action string `json:"action"`
}

type TxnLocateTrade struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Remarks  string `json:"remarks"`
	ID       string `json:"id"`
	Time     string `json:"time"`
	AssetValue      float64 `json:"assetValue"`
	Asset    string `json:"asset"`
	SecurityQty float64 `json:"securityQty"`
	Security string `json:"security"`
	Action string `json:"action"`
}

//TxnGoods - User transaction details for buying goods
//type TxnGoods struct {
	//Sender   string `json:"sender"`
	//Receiver string `json:"receiver"`
	//Remarks  string `json:"remarks"`
	//ID       string `json:"id"`
	//Time     string `json:"time"`
	//Value    string `json:"value"`
	//Asset    string `json:"asset"`
//}

//TxnEncash - details of requests from merchant to encash points
//type TxnEncash struct {
	//Key       string `json:"key"`
	//ID        string `json:"id"`
	//Initiator string `json:"initiator"`
	//Bank      string `json:"bank"`
	//Points    int    `json:"points"`
	//Amount    int    `json:"amount"`
	//Remarks   string `json:"remarks"`
	//Time      string `json:"time"`
//}

// LoyaltyChaincode example simple Chaincode implementation
type SecurityLendingChaincode struct {
}

func main() {
	err := shim.Start(new(SecurityLendingChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SecurityLendingChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	key1 := args[0] //Name of Lender
	key2 := args[1] //Name of Borrower


	Lender := Entity{
		Type:    "Lender",
		Name:    key1,
		CashBalance: 50000,
		Bond: 50000,
	}
	fmt.Println(Lender)
	bytes, err := json.Marshal(Lender)
	if err != nil {
		fmt.Println("Error marsalling")
		return nil, errors.New("Error marshalling")
	}
	fmt.Println(bytes)
	err = stub.PutState(key1, bytes)
	if err != nil {
		fmt.Println("Error writing state")
		return nil, err
	}

	Borrower := Entity{
		Type:    "Borrower",
		Name:    key2,
		CashBalance: 50000,
		Bond: 50000,
	}
	fmt.Println(Borrower)
	bytes, err = json.Marshal(Borrower)
	if err != nil {
		fmt.Println("Error marsalling")
		return nil, errors.New("Error marshalling")
	}
	fmt.Println(bytes)
	err = stub.PutState(key2, bytes)
	if err != nil {
		fmt.Println("Error writing state")
		return nil, err
	}


	Infosys := Security{
		Name:    "Infosys Ltd",
		Entity: "Lender",
		Value: 20000,
	}
	fmt.Println(Infosys)
	bytes, err = json.Marshal(Infosys)
	if err != nil {
		fmt.Println("Error marsalling")
		return nil, errors.New("Error marshalling")
	}
	fmt.Println(bytes)
	err = stub.PutState("LenderInfosys", bytes)
	if err != nil {
		fmt.Println("Error writing state")
		return nil, err
	}

	Infosys1 := Security{
		Name:    "Infosys Ltd",
		Entity: "Borrower",
		Value: 20000,
	}
	fmt.Println(Infosys1)
	bytes, err = json.Marshal(Infosys1)
	if err != nil {
		fmt.Println("Error marsalling")
		return nil, errors.New("Error marshalling")
	}
	fmt.Println(bytes)
	err = stub.PutState("BorrowerInfosys", bytes)
	if err != nil {
		fmt.Println("Error writing state")
		return nil, err
	}

	Apple := Security{
		Name:    "Apple Inc",
		Entity: "Lender",
		Value: 20000,
	}
	fmt.Println(Apple)
	bytes, err = json.Marshal(Apple)
	if err != nil {
		fmt.Println("Error marsalling")
		return nil, errors.New("Error marshalling")
	}
	fmt.Println(bytes)
	err = stub.PutState("LenderApple", bytes)
	if err != nil {
		fmt.Println("Error writing state")
		return nil, err
	}

	Apple1 := Security{
		Name:    "Apple Inc",
		Entity: "Borrower",
		Value: 20000,
	}
	fmt.Println(Apple1)
	bytes, err = json.Marshal(Apple1)
	if err != nil {
		fmt.Println("Error marsalling")
		return nil, errors.New("Error marshalling")
	}
	fmt.Println(bytes)
	err = stub.PutState("BorrowerApple", bytes)
	if err != nil {
		fmt.Println("Error writing state")
		return nil, err
	}

	// Initialize the collection of  keys for products and various transactions
	fmt.Println("Initializing keys collection")
	var blank []string
	blankBytes, _ := json.Marshal(&blank)
	//err = stub.PutState("Security", blankBytes)
	//if err != nil {
	//	fmt.Println("Failed to initialize Products key collection")
	//}
	err = stub.PutState("TxnCollaterals", blankBytes)
	if err != nil {
		fmt.Println("Failed to initialize TxnTopUp key collection")
	}
	err = stub.PutState("TxnSecurities", blankBytes)
	if err != nil {
		fmt.Println("Failed to initialize TxnTopUp key collection")
	}
	err = stub.PutState("TxnBorrowTrade", blankBytes)
	if err != nil {
		fmt.Println("Failed to initialize TxnGoods key collection")
	}
	err = stub.PutState("TxnLocateTrade", blankBytes)
	if err != nil {
		fmt.Println("Failed to initialize TxnGoods key collection")
	}
	//err = stub.PutState("TxnEncash", blankBytes)
	//if err != nil {
		//fmt.Println("Failed to initialize TxnEncash key collection")
	//}
	//err = stub.PutState("TxnTransfer", blankBytes)
	//if err != nil {
		//fmt.Println("Failed to initialize TxnTransfer key collection")
	//}

	fmt.Println("Initialization complete")


	//t.addProduct(stub, []string{"Cappuccino", "295", "2.95", key2, "500"})

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SecurityLendingChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions/transactions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	} else if function == "CreateCollaterals" {
		return t.CreateCollaterals(stub, args)
	}else if function == "CreateSecurities" {
	 return t.CreateSecurities(stub, args)
  }else if function == "BorrowTrade" {
	 return t.BorrowTrade(stub, args)
  }else if function == "approveBorrow" {
	 return t.approveBorrow(stub, args)
  }

	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SecurityLendingChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" {
		return t.read(stub, args)
	} else if function == "getTxnCollaterals" {
		return t.getTxnCollaterals(stub)
	} else if function == "getTxnSecurities" {
		return t.getTxnSecurities(stub)
	}else if function == "getTxnBorrowTrade" {
		return t.getTxnBorrowTrade(stub)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}
func (t *SecurityLendingChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("running write()")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. expecting 4")
	}

	//writing a new customer to blockchain
	typeOf := args[0]
	name := args[1]
	balance, err := strconv.ParseFloat(args[2], 64)
	bond, err := strconv.ParseFloat(args[3], 64)
	entity := Entity{
		Type:    typeOf,
		Name:    name,
		CashBalance: balance,
		Bond: bond,
	}
	fmt.Println(entity)
	bytes, err := json.Marshal(entity)
	if err != nil {
		fmt.Println("Error marsalling")
		return nil, errors.New("Error marshalling")
	}
	fmt.Println(bytes)
	err = stub.PutState(name, bytes)
	if err != nil {
		fmt.Println("Error writing state")
		return nil, err
	}

	return nil, nil
}
func (t *SecurityLendingChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("read is running")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. expecting 1")
	}

	key := args[0] // name of Entity

	bytes, err := stub.GetState(key)
	if err != nil {
		fmt.Println("Error retrieving " + key)
		return nil, errors.New("Error retrieving " + key)
	}
	lender := Entity{}
	err = json.Unmarshal(bytes, &lender)
	if err != nil {
		fmt.Println("Error Unmarshaling customerBytes")
		return nil, errors.New("Error Unmarshaling customerBytes")
	}
	bytes, err = json.Marshal(lender)
	if err != nil {
		fmt.Println("Error marshaling lender")
		return nil, errors.New("Error marshaling lender")
	}

	fmt.Println(bytes)
	return bytes, nil
}



func (t *SecurityLendingChaincode) CreateCollaterals(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("add is running ")

	if len(args) != 3 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 3 for add")
	}


	key := args[0]   //Entity ex: customer


	// GET the state of entity from the ledger
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, errors.New("Failed to get state of " + key)
	}

	entity := Entity{}
	err = json.Unmarshal(bytes, &entity)
	if err != nil {
		fmt.Println("Error Unmarshaling entity Bytes")
		return nil, errors.New("Error Unmarshaling entity Bytes")
	}
		if entity.Type == "Borrower" {
        value, err := strconv.ParseFloat(args[1], 64)

		if err == nil {
			entity.CashBalance = entity.CashBalance + value
			fmt.Println("entity Points = ", entity.CashBalance)
		}
        } else {
        return nil, errors.New("Error in link Collaterals")

        }


	// Write the state back to the ledger
	bytes, err = json.Marshal(entity)
	if err != nil {
		fmt.Println("Error marshaling entity")
		return nil, errors.New("Error marshaling entity")
	}
	err = stub.PutState(key, bytes)
	if err != nil {
		return nil, err
	}


	ID := stub.GetTxID()
	blockTime, err := stub.GetTxTimestamp()
	args = append(args, ID)
	args = append(args, blockTime.String())
	t.putTxnCollaterals(stub, args)

	return nil, nil
}

func (t *SecurityLendingChaincode) putTxnCollaterals(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("putTxnCollaterals is running ")

	if len(args) != 5 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 6 for putTxnCollaterals")
	}
	  assetValue, err := strconv.ParseFloat(args[1], 64)
	txn := TxnCollaterals{

		Initiator: args[0],
		Remarks:   "Value addedd",
		ID:        args[3],
		Time:      args[4],
		Value:     assetValue,
		Asset:     args[2],
	}

	bytes, err := json.Marshal(txn)
	if err != nil {
		fmt.Println("Error marshaling TxnCollaterals")
		return nil, errors.New("Error marshaling TxnCollaterals")
	}

	err = stub.PutState(txn.ID, bytes)
	if err != nil {
		return nil, err
	}

	return t.appendKey(stub, "TxnCollaterals", txn.ID)
}

func (t *SecurityLendingChaincode) CreateSecurities(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("add is running ")

	if len(args) != 3 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 3 for add")
	}
	key := args[0]


	if args[2] == "Infosys Ltd" {
			key = "Infosys"
		}

	bytes, err := stub.GetState("Lender"+key)
	if err != nil {
		return nil, errors.New("Failed to get state of " + key)
	}

	security := Security{}
	err = json.Unmarshal(bytes, &security)
	if err != nil {
		fmt.Println("Error Unmarshaling security Bytes")
		return nil, errors.New("Error Unmarshaling security Bytes")
	}
		if security.Entity == "Lender" {
        value, err := strconv.ParseFloat(args[1], 64)

		if err == nil {
			security.Value = security.Value + value
			fmt.Println("security value = ", security.Value)
		}
        } else {
        return nil, errors.New("Error in link Securities")

        }


	// Write the state back to the ledger
	bytes, err = json.Marshal(security)
	if err != nil {
		fmt.Println("Error marshaling security")
		return nil, errors.New("Error marshaling security")
	}
	err = stub.PutState("Lender"+key, bytes)
	if err != nil {
		return nil, err
	}


	ID := stub.GetTxID()
	blockTime, err := stub.GetTxTimestamp()
	args = append(args, ID)
	args = append(args, blockTime.String())
	t.putTxnSecurities(stub, args)

	return nil, nil
}

func (t *SecurityLendingChaincode) putTxnSecurities(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("putTxnSecurities is running ")

	if len(args) != 5 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 6 for putTxnSecurities")
	}
	securityValue, err := strconv.ParseFloat(args[1], 64)
	txn := TxnSecurities{

		Initiator: args[0],
		Remarks:   "Value addedd",
		ID:        args[3],
		Time:      args[4],
		Value:     securityValue,
		Security:  args[2],

	}

	bytes, err := json.Marshal(txn)
	if err != nil {
		fmt.Println("Error marshaling TxnSecurities")
		return nil, errors.New("Error marshaling TxnSecurities")
	}

	err = stub.PutState(txn.ID, bytes)
	if err != nil {
		return nil, err
	}

	return t.appendKey(stub, "TxnSecurities", txn.ID)
}

func (t *SecurityLendingChaincode) BorrowTrade(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("BorrowTrade is running ")

	if len(args) != 7 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 7 for transfer")
	}


	key := args[0]   // fromEntity
	key2 := args[1]  // toEntity
	key3 := "Apple"

	if args[5] == "Infosys Ltd"{
		key3 = "Infosys"
	}

	// GET the state of fromEntity from the ledger
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, errors.New("Failed to get state of " + key)
	}

	fromEntity := Entity{}
	err = json.Unmarshal(bytes, &fromEntity)
	if err != nil {
		fmt.Println("Error Unmarshaling entity Bytes")
		return nil, errors.New("Error Unmarshaling entity Bytes")
	}

	// GET the state of toEntity from the ledger
	bytes, err = stub.GetState(key2)
	if err != nil {
		return nil, errors.New("Failed to get state of " + key)
	}

	toEntity := Entity{}
	err = json.Unmarshal(bytes, &toEntity)
	if err != nil {
		fmt.Println("Error Unmarshaling entity Bytes")
		return nil, errors.New("Error Unmarshaling entity Bytes")
	}

	bytes, err = stub.GetState(args[1]+key3)
	if err != nil {
		return nil, errors.New("Failed to get state of " + key)
	}

	lenderSecurity := Security{}
	err = json.Unmarshal(bytes, &lenderSecurity)
	if err != nil {
		fmt.Println("Error Unmarshaling lenderSecurity Bytes")
		return nil, errors.New("Error Unmarshaling lenderSecurity Bytes")
	}

      assetValue, err := strconv.ParseFloat(args[2], 64)
			securityQty,err := strconv.ParseFloat(args[4],64)
			if lenderSecurity.Value > securityQty{
			if args[3] == "Cash"{
				if fromEntity.CashBalance > assetValue  {
				if err == nil {
					fromEntity.CashBalance = fromEntity.CashBalance - assetValue
				}
				}
			}	else{
				if fromEntity.Bond > assetValue {
				if err == nil {
					fromEntity.Bond = fromEntity.Bond - assetValue
				}
				}
			}
			}else{
			return nil, errors.New("Security units should less than or equal to avialable security units")
		}
	// Write the state back to the ledger
	bytes, err = json.Marshal(fromEntity)
	if err != nil {
		fmt.Println("Error marshaling fromEntity")
		return nil, errors.New("Error marshaling fromEntity")
	}
	err = stub.PutState(key, bytes)
	if err != nil {
		return nil, err
	}

	bytes, err = json.Marshal(toEntity)
	if err != nil {
		fmt.Println("Error marshaling toEntity")
		return nil, errors.New("Error marshaling toEntity")
	}
	err = stub.PutState(key2, bytes)
	if err != nil {
		return nil, err
	}

	ID := stub.GetTxID()
	blockTime, err := stub.GetTxTimestamp()
	args = append(args, ID)
	args = append(args, blockTime.String())
	t.putTxnBorrowTrade(stub, args)

	return nil, nil

}

func (t *SecurityLendingChaincode) putTxnBorrowTrade(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("putTxnBorrowTrade is running ")

	if len(args) != 9 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 9 for putTxnBorrowTrade")
	}
	assetValue, err := strconv.ParseFloat(args[2], 64)
	securityQty,err := strconv.ParseFloat(args[4],64)
	txn := TxnBorrowTrade{
		Sender:     args[0],
		Receiver:   args[1],
		Offer: args[2] + " " + args[3],
		Ask:   args[4] + " " + args[5],
		ID:        args[7],
		Time:      args[8],
		AssetValue:     assetValue,
		Asset:     args[3],
		SecurityQty: securityQty,
		Security:    args[5],
		Action:   args[6],

	}

	bytes, err := json.Marshal(txn)
	if err != nil {
		fmt.Println("Error marshaling TxnTopup")
		return nil, errors.New("Error marshaling TxnTopup")
	}

	err = stub.PutState(txn.ID, bytes)
	if err != nil {
		return nil, err
	}

	return t.appendKey(stub, "TxnTransfer", txn.ID)
}

func (t *SecurityLendingChaincode) appendKey(stub shim.ChaincodeStubInterface, primeKey string, key string) ([]byte, error) {
	fmt.Println("appendKey is running " + primeKey + " " + key)

	bytes, err := stub.GetState(primeKey)
	if err != nil {
		return nil, err
	}
	var keys []string
	err = json.Unmarshal(bytes, &keys)
	if err != nil {
		return nil, err
	}
	keys = append(keys, key)
	bytes, err = json.Marshal(keys)
	if err != nil {
		fmt.Println("Error marshaling " + primeKey)
		return nil, errors.New("Error marshaling keys" + primeKey)
	}
	err = stub.PutState(primeKey, bytes)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (t *SecurityLendingChaincode) getTxnCollaterals(stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("getTxnCollaterals is running ")

	var txns []TxnCollaterals

	// Get list of all the keys - TxnCollaterals
	keysBytes, err := stub.GetState("TxnCollaterals")
	if err != nil {
		fmt.Println("Error retrieving TxnCollaterals keys")
		return nil, errors.New("Error retrieving TxnCollaterals keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling TxnCollaterals key")
		return nil, errors.New("Error unmarshalling TxnCollaterals keys")
	}

	// Get each txn "TxnCollaterals" keys
	for _, value := range keys {
		bytes, err := stub.GetState(value)

		var txn TxnCollaterals
		err = json.Unmarshal(bytes, &txn)
		if err != nil {
			fmt.Println("Error retrieving txn " + value)
			return nil, errors.New("Error retrieving txn " + value)
		}

		fmt.Println("Appending txn" + value)
		txns = append(txns, txn)
	}

	bytes, err := json.Marshal(txns)
	if err != nil {
		fmt.Println("Error marshaling txns Collaterals")
		return nil, errors.New("Error marshaling txns TxnCollaterals")
	}
	return bytes, nil
}

func (t *SecurityLendingChaincode) getTxnSecurities(stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("getTxnSecurities is running ")

	var txns []TxnSecurities

	// Get list of all the keys - TxnSecurities
	keysBytes, err := stub.GetState("TxnSecurities")
	if err != nil {
		fmt.Println("Error retrieving TxnSecurities keys")
		return nil, errors.New("Error retrieving TxnSecurities keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling TxnSecurities key")
		return nil, errors.New("Error unmarshalling TxnSecurities keys")
	}

	// Get each txn "TxnSecurities" keys
	for _, value := range keys {
		bytes, err := stub.GetState(value)

		var txn TxnSecurities
		err = json.Unmarshal(bytes, &txn)
		if err != nil {
			fmt.Println("Error retrieving txn " + value)
			return nil, errors.New("Error retrieving txn " + value)
		}

		fmt.Println("Appending txn" + value)
		txns = append(txns, txn)
	}

	bytes, err := json.Marshal(txns)
	if err != nil {
		fmt.Println("Error marshaling txns Securities")
		return nil, errors.New("Error marshaling txns Securities")
	}
	return bytes, nil
}

func (t *SecurityLendingChaincode) getTxnBorrowTrade(stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("getTxnBorrowTrade is running ")

	var txns []TxnBorrowTrade

	// Get list of all the keys - TxnSecurities
	keysBytes, err := stub.GetState("TxnBorrowTrade")
	if err != nil {
		fmt.Println("Error retrieving TxnBorrowTrade keys")
		return nil, errors.New("Error retrieving TxnBorrowTrade keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling TxnBorrowTrade key")
		return nil, errors.New("Error unmarshalling TxnBorrowTrade keys")
	}

	// Get each txn "TxnBorrowTrade" keys
	for _, value := range keys {
		bytes, err := stub.GetState(value)

		var txn TxnBorrowTrade
		err = json.Unmarshal(bytes, &txn)
		if err != nil {
			fmt.Println("Error retrieving txn " + value)
			return nil, errors.New("Error retrieving txn " + value)
		}

		fmt.Println("Appending txn" + value)
		txns = append(txns, txn)
	}

	bytes, err := json.Marshal(txns)
	if err != nil {
		fmt.Println("Error marshaling txns BorrowTrade")
		return nil, errors.New("Error marshaling txns BorrowTrade")
	}
	return bytes, nil
}

func (t *SecurityLendingChaincode) approveBorrow(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("approveBorrow is running ")

	if len(args) != 7 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 7 for approveBorrow")
	}

	assetValue, err := strconv.ParseFloat(args[2], 64)
	securityQty, err := strconv.ParseFloat(args[4], 64)

	bytes, err := stub.GetState(args[0])
	if err != nil {
		return nil, errors.New("Failed to get state of " + args[0])
	}
	if bytes == nil {
		return nil, errors.New("Entity not found")
	}
	lender := Entity{}
	err = json.Unmarshal(bytes, &lender)
	if err != nil {
		fmt.Println("Error Unmarshaling lender")
		return nil, errors.New("Error Unmarshaling lender")
	}
	keySecurity := "Apple"
	if args[5] == "Infosys Ltd"{
		keySecurity = "Infosys"
	}
	bytes, err = stub.GetState(args[0]+keySecurity)
	if err != nil {
		return nil, errors.New("Failed to get state of " + args[1])
	}
	if bytes == nil {
		return nil, errors.New("Entity not found")
	}
	lenderSecurity := Security{}
	err = json.Unmarshal(bytes, &lenderSecurity)
	if err != nil {
		fmt.Println("Error Unmarshaling lenderSecurity")
		return nil, errors.New("Error Unmarshaling lenderSecurity")
	}

	bytes, err = stub.GetState(args[1]+keySecurity)
	if err != nil {
		return nil, errors.New("Failed to get state of " + args[1])
	}
	if bytes == nil {
		return nil, errors.New("Entity not found")
	}
	borrowerSecurity := Security{}
	err = json.Unmarshal(bytes, &borrowerSecurity)
	if err != nil {
		fmt.Println("Error Unmarshaling borrowerSecurity")
		return nil, errors.New("Error Unmarshaling borrowerSecurity")
	}


	// Perform approve
		if args[3] == "Cash"{
			lender.CashBalance = lender.CashBalance + assetValue
		}else{
			lender.Bond = lender.Bond + assetValue
		}


			lenderSecurity.Value = lenderSecurity.Value - securityQty
		  borrowerSecurity.Value = borrowerSecurity.Value + securityQty


	// Write the merchant/entity1 state back to the ledger
	bytes, err = json.Marshal(lender)
	if err != nil {
		fmt.Println("Error marshaling lender")
		return nil, errors.New("Error marshaling lender")
	}
	err = stub.PutState(args[0], bytes)
	if err != nil {
		return nil, err
	}

	// Write the bank/entity2 state back to the ledger]
	bytes, err = json.Marshal(lenderSecurity)
	if err != nil {
		fmt.Println("Error marshaling lenderSecurity")
		return nil, errors.New("Error marshaling lenderSecurity")
	}
	err = stub.PutState(args[0]+keySecurity, bytes)
	if err != nil {
		return nil, err
	}
	bytes, err = json.Marshal(borrowerSecurity)
	if err != nil {
		fmt.Println("Error marshaling borrowerSecurity")
		return nil, errors.New("Error marshaling borrowerSecurity")
	}
	err = stub.PutState(args[1]+keySecurity, bytes)
	if err != nil {
		return nil, err
	}


	ID := stub.GetTxID()
	blockTime, err := stub.GetTxTimestamp()
	args = append(args, ID)
	args = append(args, blockTime.String())
	t.putTxnBorrowTrade(stub, args)

	return nil, nil
}
