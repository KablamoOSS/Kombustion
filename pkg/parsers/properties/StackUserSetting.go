package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// StackUserSetting Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-usersetting.html
type StackUserSetting struct {
	Action     interface{} `yaml:"Action"`
	Permission interface{} `yaml:"Permission"`
}

// StackUserSetting validation
func (resource StackUserSetting) Validate() []error {
	errors := []error{}

	return errors
}
