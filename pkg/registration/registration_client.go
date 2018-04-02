package registration

import (
	"github.com/golang/glog"

	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

const (
	TargetIdentifierField string = "targetIdentifier"
	Username              string = "username"
	Password              string = "password"
)

type DemoRegistrationClient struct {
}

func NewRegistrationClient() *DemoRegistrationClient {
	return &DemoRegistrationClient{}
}

func (rClient *DemoRegistrationClient) GetSupplyChainDefinition() []*proto.TemplateDTO {
	supplyChainFactory := NewSupplyChainFactory()
	supplyChain, err := supplyChainFactory.createSupplyChain()
	if err != nil {
		glog.Errorf("Failed to build supply chain: %v", err)
		return []*proto.TemplateDTO{}
	}
	return supplyChain
}

func (rClient *DemoRegistrationClient) GetAccountDefinition() []*proto.AccountDefEntry {
	var acctDefProps []*proto.AccountDefEntry

	// target ID
	targetIDAcctDefEntry := builder.NewAccountDefEntryBuilder(TargetIdentifierField, "TargetID",
		"ID or address of the target", ".*", false, false).Create()
	acctDefProps = append(acctDefProps, targetIDAcctDefEntry)

	// username
	usernameAcctDefEntry := builder.NewAccountDefEntryBuilder(Username, "Username",
		"Username of the target cluster master", ".*", false, false).Create()
	acctDefProps = append(acctDefProps, usernameAcctDefEntry)

	// password
	passwordAcctDefEntry := builder.NewAccountDefEntryBuilder(Password, "Password",
		"Password of the target cluster master", ".*", false, true).Create()
	acctDefProps = append(acctDefProps, passwordAcctDefEntry)

	return acctDefProps
}

func (rClient *DemoRegistrationClient) GetIdentifyingFields() string {
	return TargetIdentifierField
}
