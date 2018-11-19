package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// WAFByteMatchSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html
type WAFByteMatchSet struct {
	Type       string                    `yaml:"Type"`
	Properties WAFByteMatchSetProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

// WAFByteMatchSet Properties
type WAFByteMatchSetProperties struct {
	Name            interface{} `yaml:"Name"`
	ByteMatchTuples interface{} `yaml:"ByteMatchTuples,omitempty"`
}

// NewWAFByteMatchSet constructor creates a new WAFByteMatchSet
func NewWAFByteMatchSet(properties WAFByteMatchSetProperties, deps ...interface{}) WAFByteMatchSet {
	return WAFByteMatchSet{
		Type:       "AWS::WAF::ByteMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFByteMatchSet parses WAFByteMatchSet
func ParseWAFByteMatchSet(
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
	var resource WAFByteMatchSet
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
					"Fn::Sub": "${AWS::StackName}-WAFByteMatchSet-" + name,
				},
			},
		},
	}

	return
}

// ParseWAFByteMatchSet validator
func (resource WAFByteMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFByteMatchSetProperties validator
func (resource WAFByteMatchSetProperties) Validate() []error {
	errors := []error{}
	return errors
}
