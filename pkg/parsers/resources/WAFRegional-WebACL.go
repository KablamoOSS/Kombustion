package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// WAFRegionalWebACL Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html
type WAFRegionalWebACL struct {
	Type       string                      `yaml:"Type"`
	Properties WAFRegionalWebACLProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

// WAFRegionalWebACL Properties
type WAFRegionalWebACLProperties struct {
	MetricName    interface{}              `yaml:"MetricName"`
	Name          interface{}              `yaml:"Name"`
	Rules         interface{}              `yaml:"Rules,omitempty"`
	DefaultAction *properties.WebACLAction `yaml:"DefaultAction"`
}

// NewWAFRegionalWebACL constructor creates a new WAFRegionalWebACL
func NewWAFRegionalWebACL(properties WAFRegionalWebACLProperties, deps ...interface{}) WAFRegionalWebACL {
	return WAFRegionalWebACL{
		Type:       "AWS::WAFRegional::WebACL",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFRegionalWebACL parses WAFRegionalWebACL
func ParseWAFRegionalWebACL(
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
	var resource WAFRegionalWebACL
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

// ParseWAFRegionalWebACL validator
func (resource WAFRegionalWebACL) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFRegionalWebACLProperties validator
func (resource WAFRegionalWebACLProperties) Validate() []error {
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
