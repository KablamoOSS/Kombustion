package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// EnvironmentOptionSetting Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html
type EnvironmentOptionSetting struct {
	Namespace    interface{} `yaml:"Namespace"`
	OptionName   interface{} `yaml:"OptionName"`
	ResourceName interface{} `yaml:"ResourceName,omitempty"`
	Value        interface{} `yaml:"Value,omitempty"`
}

// EnvironmentOptionSetting validation
func (resource EnvironmentOptionSetting) Validate() []error {
	errors := []error{}

	return errors
}
