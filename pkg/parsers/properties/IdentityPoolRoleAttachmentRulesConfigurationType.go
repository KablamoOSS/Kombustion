package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// IdentityPoolRoleAttachmentRulesConfigurationType Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rulesconfigurationtype.html
type IdentityPoolRoleAttachmentRulesConfigurationType struct {
	Rules interface{} `yaml:"Rules"`
}

// IdentityPoolRoleAttachmentRulesConfigurationType validation
func (resource IdentityPoolRoleAttachmentRulesConfigurationType) Validate() []error {
	errors := []error{}

	return errors
}
