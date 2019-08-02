package invoke

import (
	"fmt"
	"strconv"
)


func(s *OrgInvoke) UpdateUserFromLedger(email , name, mobile, age, salary, role string) error {

	fmt.Println(" ############## Invoke Update Data ################")
		
	eventID := "updateInvoke"
	queryCreatorOrg := s.User.Setup.OrgName
	queryCreatorRole := s.Role
	needHistory := strconv.FormatBool(true)
 
		_, err := s.User.Setup.ExecuteChaincodeTranctionEvent(eventID, "invoke",
		[][]byte{
			[]byte("updateUserData"),
			[]byte(name),
			[]byte(email),
			[]byte(mobile),
			[]byte(age),
			[]byte(salary),
			[]byte(eventID),
			[]byte(queryCreatorOrg),
			[]byte(queryCreatorRole),
			[]byte(needHistory),
		}, s.User.Setup.ChaincodeId, s.User.Setup.ChannelClient, s.User.Setup.Event)

	if err != nil {
		fmt.Errorf(" Error - Update User Data From Ledger : %s ",err.Error())
	} 

	fmt.Println(" ###################################################### ")

	return nil
}
