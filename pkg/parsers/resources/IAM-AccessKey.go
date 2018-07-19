package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// IAMAccessKey Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html
type IAMAccessKey struct {
	Type       string                 `yaml:"Type"`
	Properties IAMAccessKeyProperties `yaml:"Properties"`
	Condition  interface{}            `yaml:"Condition,omitempty"`
	Metadata   interface{}            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}            `yaml:"DependsOn,omitempty"`
}

// IAMAccessKey Properties
type IAMAccessKeyProperties struct {
	Serial   interface{} `yaml:"Serial,omitempty"`
	Status   interface{} `yaml:"Status,omitempty"`
	UserName interface{} `yaml:"UserName"`
}

// NewIAMAccessKey constructor creates a new IAMAccessKey
func NewIAMAccessKey(properties IAMAccessKeyProperties, deps ...interface{}) IAMAccessKey {
	return IAMAccessKey{
		Type:       "AWS::IAM::AccessKey",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseIAMAccessKey parses IAMAccessKey
func ParseIAMAccessKey(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource IAMAccessKey
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	return
}

// ParseIAMAccessKey validator
func (resource IAMAccessKey) Validate() []error {
	return resource.Properties.Validate()
}

// ParseIAMAccessKeyProperties validator
func (resource IAMAccessKeyProperties) Validate() []error {
	errors := []error{}
	if resource.UserName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'UserName'"))
	}
	return errors
}
