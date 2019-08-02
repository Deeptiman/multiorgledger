package html

import (
	"net/http"
	"github.com/multiorgledger/blockchain/org"
	"github.com/multiorgledger/blockchain/invoke"
	"github.com/multiorgledger/chaincode/model"
)


type Data struct {

	Error        			bool
	ErrorMsg     			string
	Success      			bool
	Response     			bool

	Admin        			bool
	User         			bool

	SessionOrg	 			string	
	SessionUserData     	*model.User

	AllUsersData 			[]model.User

	AllHistoryData 			[]model.HistoryData
	HistoryUser				string
	History					bool

	UpdateUser				string
	Update					bool

	DeleteUser				string
	Delete					bool

	Logout					bool

	CustomOrg1				string
	CustomOrg2				string
	CustomOrg3				string
	CustomOrg4				string

}


func (app *HtmlApp) IndexPageHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		data := &Data {}
		
		orgUser := app.Org.GetOrgUser()

		if orgUser == nil {

			data.Error = true
			data.ErrorMsg = "No session available"

		} else {

			data, err := data.Setup(orgUser, true)
			if err != nil && data !=nil {
				data.Response = true
				data.Error = true
				data.ErrorMsg = err.Error()
			}  
						 
			renderTemplate(w, r, "index.html", data)

		}
		
	})
}

func(data *Data) Setup(orgUser *org.OrgUser, needHistory bool) (*Data, error){

	orgInvoke := invoke.OrgInvoke {
		User: orgUser,
	}

	/* Session User Data */

	SessionUserData, err := orgInvoke.GetUserFromLedger(orgUser.Username, needHistory)

	if err != nil {
		return nil, err
	} 
	
	data.SessionUserData = SessionUserData
	data.SessionOrg = orgUser.Setup.OrgName

	/* Is Logged In User is Admin? */
	if model.IsAdmin(SessionUserData.Role){

		allUsersData, err := orgInvoke.GetAllUsersFromLedger()
		if err != nil {
			return nil, err
		}

		data.Admin = true
		data.AllUsersData = allUsersData
		
	} else {
		data.User = true
	}
	
	data.Success = true
	data.Response = true

	return &Data {

		Success:      		data.Success,
		Response:    		data.Response,
		Admin:        		data.Admin,
		User:         		data.User,
		SessionOrg:	  		data.SessionOrg,
		SessionUserData: 	data.SessionUserData,			
		AllUsersData: 		data.AllUsersData,
		CustomOrg1:			model.GetCustomOrgName("org1"),
		CustomOrg2:			model.GetCustomOrgName("org2"),
		CustomOrg3:			model.GetCustomOrgName("org3"),
		CustomOrg4:			model.GetCustomOrgName("org4"),
	}, nil

}
 
