package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// StackApplicationSettings Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-applicationsettings.html
type StackApplicationSettings struct {
	Enabled       interface{} `yaml:"Enabled"`
	SettingsGroup interface{} `yaml:"SettingsGroup,omitempty"`
}

// StackApplicationSettings validation
func (resource StackApplicationSettings) Validate() []error {
	errors := []error{}

	return errors
}
