package main

import (
	"fmt"
	"multiorgledger/chaincode/model"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func (t *MultiOrgChaincode) deleteUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println(" ******** Invoke Delete User ******** ")

	var email, role, eventID string
	var queryCreatorOrg string
	var queryCreatorRole string
	var queryCreator string

	var needHistory bool

	email = args[1]
	eventID = args[2]
	role = args[3]
	queryCreatorOrg = args[4]
	queryCreatorRole = args[5]
	needHistory, _ = strconv.ParseBool(args[6])

	fmt.Println(" ###### Delete Data Parameters ###### ")
	fmt.Println(" Email	= " + email)
	fmt.Println(" Role	= " + role)
	fmt.Println(" EventID	= " + eventID)
	fmt.Println(" ################################## ")

	indexName := model.COLLECTION_KEY
	userNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{email})

	if err != nil {
		return shim.Error(err.Error())
	}

	userNameIndexKey, err = stub.CreateCompositeKey(indexName, []string{args[1]})

	err = deleteFromLedger(stub, userNameIndexKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to delete the user in the ledger: %v", err))
	}

	err = stub.SetEvent(eventID, []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	/*	Created History for Delete by email Transaction */

	if needHistory {
		var remarks string
		if strings.EqualFold(queryCreatorRole, model.ADMIN) {
			queryCreator = model.GetCustomOrgName(queryCreatorOrg) + " Admin"
			remarks = queryCreator + " has deleted user - " + email
		} else {
			queryCreator = email
			remarks = queryCreator + " has deleted the account"
		}

		fmt.Println(" ###### Query Access Details ###### ")
		fmt.Println(" queryCreatorRole = " + queryCreatorRole)
		fmt.Println(" queryCreator = " + queryCreator)
		fmt.Println(" ################################## ")

		query := args[0]
		t.createHistory(stub, queryCreator, queryCreatorOrg, email, query, remarks)
	}

	return shim.Success(nil)
}
