package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// RuleSqsParameters Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-sqsparameters.html
type RuleSqsParameters struct {
	MessageGroupId interface{} `yaml:"MessageGroupId"`
}

// RuleSqsParameters validation
func (resource RuleSqsParameters) Validate() []error {
	errors := []error{}

	return errors
}
