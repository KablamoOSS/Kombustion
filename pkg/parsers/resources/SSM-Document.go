package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	//
	// "fmt"
	//
	//
)

// SSMDocument Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html
type SSMDocument struct {
	Type       string                `yaml:"Type"`
	Properties SSMDocumentProperties `yaml:"Properties"`
	Condition  interface{}           `yaml:"Condition,omitempty"`
	Metadata   interface{}           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}           `yaml:"DependsOn,omitempty"`
}

// SSMDocument Properties
type SSMDocumentProperties struct {
	Content      interface{} `yaml:"Content"`
	DocumentType interface{} `yaml:"DocumentType,omitempty"`
	Tags         interface{} `yaml:"Tags,omitempty"`
}

// NewSSMDocument constructor creates a new SSMDocument
func NewSSMDocument(properties SSMDocumentProperties, deps ...interface{}) SSMDocument {
	return SSMDocument{
		Type:       "AWS::SSM::Document",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSSMDocument parses SSMDocument
func ParseSSMDocument(
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
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"

	// Resources
	var resource SSMDocument
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

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-SSMDocument-" + name,
				},
			},
		},
	}

	return
}

// ParseSSMDocument validator
func (resource SSMDocument) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSSMDocumentProperties validator
func (resource SSMDocumentProperties) Validate() []error {
	errors := []error{}
	return errors
}
