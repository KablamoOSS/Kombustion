package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2NetworkAcl Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html
type EC2NetworkAcl struct {
	Type       string                  `yaml:"Type"`
	Properties EC2NetworkAclProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// EC2NetworkAcl Properties
type EC2NetworkAclProperties struct {
	VpcId interface{} `yaml:"VpcId"`
	Tags  interface{} `yaml:"Tags,omitempty"`
}

// NewEC2NetworkAcl constructor creates a new EC2NetworkAcl
func NewEC2NetworkAcl(properties EC2NetworkAclProperties, deps ...interface{}) EC2NetworkAcl {
	return EC2NetworkAcl{
		Type:       "AWS::EC2::NetworkAcl",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2NetworkAcl parses EC2NetworkAcl
func ParseEC2NetworkAcl(
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
	var resource EC2NetworkAcl
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

// ParseEC2NetworkAcl validator
func (resource EC2NetworkAcl) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2NetworkAclProperties validator
func (resource EC2NetworkAclProperties) Validate() []error {
	errors := []error{}
	if resource.VpcId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errors
}
