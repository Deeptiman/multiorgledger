package main

import (
	"encoding/json"
	"fmt"
	"multiorgledger/chaincode/model"
	"strconv"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func (t *MultiOrgChaincode) createUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(" ******** Invoke Create User ******** ")

	var queryCreatorOrg string
	var email, name, mobile, age, salary string
	var eventID string
	var owner string

	var needHistory bool

	var timestamp *timestamp.Timestamp

	/* User Data Parameter */
	name = args[1]
	email = args[2]
	mobile = args[3]
	age = args[4]
	salary = args[5]

	eventID = args[6]
	queryCreatorOrg = args[7]
	needHistory, _ = strconv.ParseBool(args[8])

	role, err := t.getRole(stub)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to get roles from the account: %v", err))
	}

	userID, err := cid.GetID(stub)

	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to identify the ID of the request owner: %v", err))
	}

	timestamp, err = stub.GetTxTimestamp()

	if err != nil {
		return shim.Error("Timestamp Error " + err.Error())
	}

	tm := model.GetTime(timestamp)

	user := &model.User{
		ID:     userID,
		Name:   name,
		Email:  email,
		Mobile: mobile,
		Age:    age,
		Salary: salary,
		Owner:  queryCreatorOrg,
		Role:   role,
		Time:   tm,
	}

	userJSONasBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}

	indexName := model.COLLECTION_KEY
	userNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{user.Email})

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(userNameIndexKey, userJSONasBytes)

	if err != nil {
		return shim.Error("###### Error Put Private Create User Data Failed " + err.Error())
	}

	err = stub.SetEvent(eventID, []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(" ###### Create Data Parameters ###### ")
	fmt.Println(" ID 			= " + userID)
	fmt.Println(" Email			= " + email)
	fmt.Println(" Name 			= " + name)
	fmt.Println(" Mobile 			= " + mobile)
	fmt.Println(" Age			= " + age)
	fmt.Println(" Salary 			= " + salary)
	fmt.Println(" Owner 			= " + owner)
	fmt.Println(" Role			= " + role)
	fmt.Println(" Time			= " + tm)
	fmt.Println(" ################################## ")

	/*	Created History for Create user Transaction */

	if needHistory {
		query := args[0]
		queryCreator := email
		remarks := email + " user created"
		t.createHistory(stub, queryCreator, queryCreatorOrg, email, query, remarks)
	}

	fmt.Println("User Invoked into the Ledger Successfully")

	return shim.Success(nil)
}
