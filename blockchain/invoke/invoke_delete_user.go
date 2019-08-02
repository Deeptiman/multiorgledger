package invoke

import (
	"fmt"
	"strconv"
)

func (s *OrgInvoke) DeleteUserFromLedger(email, role string) error {

	fmt.Println(" ############## Invoke Delete User ################")
  
	eventID := "deleteInvoke"
	queryCreatorOrg := s.User.Setup.OrgName
	queryCreatorRole := s.Role
	needHistory := strconv.FormatBool(true)

	_, err := s.User.Setup.ExecuteChaincodeTranctionEvent(eventID, "invoke",
		[][]byte{
			[]byte("deleteUser"),
			[]byte(email),			
			[]byte(eventID),
			[]byte(role),
			[]byte(queryCreatorOrg),
			[]byte(queryCreatorRole),
			[]byte(needHistory),						
		}, s.User.Setup.ChaincodeId, s.User.Setup.ChannelClient, s.User.Setup.Event)

	if err != nil {
		fmt.Errorf("Error - DeleteUserFromLedger : %s", err.Error())
	} 

	return nil
}
