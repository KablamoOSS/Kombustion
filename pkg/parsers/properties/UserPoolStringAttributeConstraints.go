package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// UserPoolStringAttributeConstraints Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html
type UserPoolStringAttributeConstraints struct {
	MaxLength interface{} `yaml:"MaxLength,omitempty"`
	MinLength interface{} `yaml:"MinLength,omitempty"`
}

// UserPoolStringAttributeConstraints validation
func (resource UserPoolStringAttributeConstraints) Validate() []error {
	errs := []error{}

	return errs
}
