package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	//
	//
)

// SESReceiptRuleSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptruleset.html
type SESReceiptRuleSet struct {
	Type       string                      `yaml:"Type"`
	Properties SESReceiptRuleSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

// SESReceiptRuleSet Properties
type SESReceiptRuleSetProperties struct {
	RuleSetName interface{} `yaml:"RuleSetName,omitempty"`
}

// NewSESReceiptRuleSet constructor creates a new SESReceiptRuleSet
func NewSESReceiptRuleSet(properties SESReceiptRuleSetProperties, deps ...interface{}) SESReceiptRuleSet {
	return SESReceiptRuleSet{
		Type:       "AWS::SES::ReceiptRuleSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSESReceiptRuleSet parses SESReceiptRuleSet
func ParseSESReceiptRuleSet(
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
	var resource SESReceiptRuleSet
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
					"Fn::Sub": "${AWS::StackName}-SESReceiptRuleSet-" + name,
				},
			},
		},
	}

	return
}

// ParseSESReceiptRuleSet validator
func (resource SESReceiptRuleSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSESReceiptRuleSetProperties validator
func (resource SESReceiptRuleSetProperties) Validate() []error {
	errors := []error{}
	return errors
}
