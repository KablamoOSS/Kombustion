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

// RDSOptionGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html
type RDSOptionGroup struct {
	Type       string                   `yaml:"Type"`
	Properties RDSOptionGroupProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// RDSOptionGroup Properties
type RDSOptionGroupProperties struct {
	EngineName             interface{} `yaml:"EngineName"`
	MajorEngineVersion     interface{} `yaml:"MajorEngineVersion"`
	OptionGroupDescription interface{} `yaml:"OptionGroupDescription"`
	OptionConfigurations   interface{} `yaml:"OptionConfigurations"`
	Tags                   interface{} `yaml:"Tags,omitempty"`
}

// NewRDSOptionGroup constructor creates a new RDSOptionGroup
func NewRDSOptionGroup(properties RDSOptionGroupProperties, deps ...interface{}) RDSOptionGroup {
	return RDSOptionGroup{
		Type:       "AWS::RDS::OptionGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseRDSOptionGroup parses RDSOptionGroup
func ParseRDSOptionGroup(
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
	var resource RDSOptionGroup
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
					"Fn::Sub": "${AWS::StackName}-RDSOptionGroup-" + name,
				},
			},
		},
	}

	return
}

// ParseRDSOptionGroup validator
func (resource RDSOptionGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseRDSOptionGroupProperties validator
func (resource RDSOptionGroupProperties) Validate() []error {
	errors := []error{}
	return errors
}
