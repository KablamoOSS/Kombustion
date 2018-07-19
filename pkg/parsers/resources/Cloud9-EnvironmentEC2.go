package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// Cloud9EnvironmentEC2 Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html
type Cloud9EnvironmentEC2 struct {
	Type       string                         `yaml:"Type"`
	Properties Cloud9EnvironmentEC2Properties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// Cloud9EnvironmentEC2 Properties
type Cloud9EnvironmentEC2Properties struct {
	AutomaticStopTimeMinutes interface{} `yaml:"AutomaticStopTimeMinutes,omitempty"`
	Description              interface{} `yaml:"Description,omitempty"`
	InstanceType             interface{} `yaml:"InstanceType"`
	Name                     interface{} `yaml:"Name,omitempty"`
	OwnerArn                 interface{} `yaml:"OwnerArn,omitempty"`
	SubnetId                 interface{} `yaml:"SubnetId,omitempty"`
	Repositories             interface{} `yaml:"Repositories,omitempty"`
}

// NewCloud9EnvironmentEC2 constructor creates a new Cloud9EnvironmentEC2
func NewCloud9EnvironmentEC2(properties Cloud9EnvironmentEC2Properties, deps ...interface{}) Cloud9EnvironmentEC2 {
	return Cloud9EnvironmentEC2{
		Type:       "AWS::Cloud9::EnvironmentEC2",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCloud9EnvironmentEC2 parses Cloud9EnvironmentEC2
func ParseCloud9EnvironmentEC2(
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
	var resource Cloud9EnvironmentEC2
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

// ParseCloud9EnvironmentEC2 validator
func (resource Cloud9EnvironmentEC2) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCloud9EnvironmentEC2Properties validator
func (resource Cloud9EnvironmentEC2Properties) Validate() []error {
	errors := []error{}
	if resource.InstanceType == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'InstanceType'"))
	}
	return errors
}
