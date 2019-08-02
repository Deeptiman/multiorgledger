package invoke

import (
	"fmt"
	"strconv"
	"github.com/multiorgledger/blockchain/org"
)

type OrgInvoke struct {
	User  *org.OrgUser
	Role  string
}

func(s *OrgInvoke) InvokeCreateUser(name, age, mobile, salary string) error {

	fmt.Println(" ############## Invoke Create User ################")

	var queryCreatorOrg string
	
	queryCreatorOrg = s.User.Setup.OrgName
	email := s.User.Username
	eventID := "userInvoke"
	needHistory := strconv.FormatBool(true)
 

	fmt.Println(" ###### Create Data Parameters ###### ")
	fmt.Println("	Email 			= "+email)
	fmt.Println(" 	Name 			= "+name)
	fmt.Println(" 	Mobile 			= "+mobile)
	fmt.Println(" 	Age 			= "+age)
	fmt.Println(" 	Salary 			= "+salary)
	fmt.Println(" 	Owner 			= "+queryCreatorOrg)
	fmt.Println(" ################################## ")

	_, err := s.User.Setup.ExecuteChaincodeTranctionEvent(eventID, "invoke",
		[][]byte{
			[]byte("createUser"), 
			[]byte(name), 
			[]byte(email),
			[]byte(mobile), 
			[]byte(age), 
			[]byte(salary),			
			[]byte(eventID), 
			[]byte(queryCreatorOrg),  
			[]byte(needHistory), 

	}, s.User.Setup.ChaincodeId, s.User.ChannelClient,s.User.Event)
	
	if err != nil {
		return fmt.Errorf("Error - addUserToLedger : %s", err.Error())
	} 


	fmt.Println("#### User added Successfully ####")

	return nil	
}