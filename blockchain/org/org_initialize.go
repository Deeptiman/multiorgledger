package org

import (
	"fmt"
	"strings"
	fabAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	contextAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	cb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/pkg/errors"
)

const (
	collCfgBlockToLive       = 1000
	collCfgRequiredPeerCount = 0
	collCfgMaximumPeerCount  = 3
)

var Orderer OrgSetup
var OrgList []OrgSetup
var OrgNames []string
var totalOrg = 5

var orgPeers []fabAPI.Peer
var channelCtx contextAPI.ChannelProvider
var collCfg *cb.CollectionConfig

var signIdentities = make([]msp.SigningIdentity, 0, totalOrg-1)
var collConfigs = make([]*cb.CollectionConfig, 0, totalOrg-1)


func(s *OrgSetup) Init(processAll bool) error {

	OrgList = make([]OrgSetup, 0, totalOrg-1)
	OrgNames = []string{"org1","org2","org3","org4"}

	if processAll {
		s.InitializeAllOrgs()
	}

	return nil
}

func(s *OrgSetup) GetOrgNames() []string {
	return OrgNames
}

func(s *OrgSetup) FilteredOrgNames() []string {
	var filteredOrg []string
	for _, org := range OrgNames {

		if !strings.EqualFold(s.OrgName, org){
			filteredOrg = append(filteredOrg, org)
		}
	}
	return filteredOrg
}

func(s *OrgSetup) InitializeOrg(org string) (OrgSetup,error) {

	var obj OrgSetup

	switch name := org; name {
			
		case "org1":

			obj = OrgSetup {
				OrgAdmin: 			"Admin",
				OrgName:  			"org1",
				ConfigFile: 			"config-org1.yaml",
				OrgCaID: 			"ca.org1.multi.org.ledger.com",
				ChannelConfig: 			"Org1MSPanchors.tx",
			}
		
			break

		case "org2":

			obj = OrgSetup {
				OrgAdmin: 			"Admin",
				OrgName:  			"org2",
				ConfigFile: 			"config-org2.yaml",
				OrgCaID: 			"ca.org2.multi.org.ledger.com",
				ChannelConfig: 			"Org2MSPanchors.tx",
			}
			 
			break

		case "org3":

			obj = OrgSetup {
				OrgAdmin: 			"Admin",
				OrgName:  			"org3",
				ConfigFile: 			"config-org3.yaml",
				OrgCaID: 			"ca.org3.multi.org.ledger.com",
				ChannelConfig: 			"Org3MSPanchors.tx",
			}
			 
			break

		case "org4":

			obj = OrgSetup {
				OrgAdmin: 			"Admin",
				OrgName:  			"org4",
				ConfigFile: 			"config-org4.yaml",
				OrgCaID: 			"ca.org4.multi.org.ledger.com",
				ChannelConfig: 			"Org4MSPanchors.tx",
			} 

			break	 
	}
	
	orgSetup, err := InitiateOrg(obj)
	if err != nil {
		return OrgSetup{}, errors.WithMessage(err, " failed to initiate Org")
	}
				
	return orgSetup, nil

}

func InitiateOrderer() (OrgSetup,error) {

	obj := OrgSetup {
		OrgAdmin: 			"Admin",
		OrgName:  			"OrdererOrg",
		ConfigFile: 			"config-org1.yaml",
		ChannelConfig: 			"multiorgledger.channel.tx",
	}

	orderer, err  := initializeOrg(obj)

	if err != nil {
		return OrgSetup{}, fmt.Errorf("  failed to setup Org - "+obj.OrgName+" - "+err.Error())
	}

	if orderer == nil {
		return OrgSetup{}, fmt.Errorf("  failed to setup Org - "+obj.OrgName)
	}

	Orderer = *orderer

	Orderer.OrdererID = "orderer.multi.org.ledger.com"

	fmt.Println(" **** Setup Created for "+Orderer.OrgName+" **** ")

	return Orderer, nil
}


func(s *OrgSetup) InitializeAllOrgs() error {

	ordererSetup, err := InitiateOrderer()
	if err != nil {
		return errors.WithMessage(err, " failed to initiate Orderer")
	}
	OrgList = append(OrgList, ordererSetup)

	for _, org := range OrgNames {

		orgSetup, err := s.InitializeOrg(org)
		if err != nil {
			return errors.WithMessage(err, " failed to initiate Org - "+org)
		}
		OrgList = append(OrgList, orgSetup)
	}

	Orderer.SigningIdentities = getSigningIdentities()
	OrgList[0] = Orderer

	return nil
}

func InitiateOrg(obj OrgSetup) (OrgSetup,error) {

	org, err := initializeOrg(obj)
	if err != nil {
		return OrgSetup{}, fmt.Errorf("  failed to setup Org - "+obj.OrgName+" - "+err.Error())
	}

	if org == nil {
		return OrgSetup{}, fmt.Errorf("  failed to setup Org - "+obj.OrgName)
	}
	
	fmt.Println(" **** Setup Created for "+org.OrgName+" **** ")

	return *org, nil
}
 
