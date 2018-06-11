package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ReceiptRuleWorkmailAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html
type ReceiptRuleWorkmailAction struct {
	OrganizationArn interface{} `yaml:"OrganizationArn"`
	TopicArn        interface{} `yaml:"TopicArn,omitempty"`
}

// ReceiptRuleWorkmailAction validation
func (resource ReceiptRuleWorkmailAction) Validate() []error {
	errs := []error{}

	if resource.OrganizationArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'OrganizationArn'"))
	}
	return errs
}
