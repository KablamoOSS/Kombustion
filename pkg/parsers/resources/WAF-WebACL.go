package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// WAFWebACL Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html
type WAFWebACL struct {
	Type       string              `yaml:"Type"`
	Properties WAFWebACLProperties `yaml:"Properties"`
	Condition  interface{}         `yaml:"Condition,omitempty"`
	Metadata   interface{}         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}         `yaml:"DependsOn,omitempty"`
}

// WAFWebACL Properties
type WAFWebACLProperties struct {
	MetricName    interface{}                 `yaml:"MetricName"`
	Name          interface{}                 `yaml:"Name"`
	DefaultAction *properties.WebACLWafAction `yaml:"DefaultAction"`
	Rules         interface{}                 `yaml:"Rules,omitempty"`
}

// NewWAFWebACL constructor creates a new WAFWebACL
func NewWAFWebACL(properties WAFWebACLProperties, deps ...interface{}) WAFWebACL {
	return WAFWebACL{
		Type:       "AWS::WAF::WebACL",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFWebACL parses WAFWebACL
func ParseWAFWebACL(
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
	var resource WAFWebACL
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

// ParseWAFWebACL validator
func (resource WAFWebACL) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFWebACLProperties validator
func (resource WAFWebACLProperties) Validate() []error {
	errors := []error{}
	if resource.MetricName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.DefaultAction == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DefaultAction'"))
	} else {
		errors = append(errors, resource.DefaultAction.Validate()...)
	}
	return errors
}
