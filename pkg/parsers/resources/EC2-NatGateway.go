package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2NatGateway Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html
type EC2NatGateway struct {
	Type       string                  `yaml:"Type"`
	Properties EC2NatGatewayProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// EC2NatGateway Properties
type EC2NatGatewayProperties struct {
	AllocationId interface{} `yaml:"AllocationId"`
	SubnetId     interface{} `yaml:"SubnetId"`
	Tags         interface{} `yaml:"Tags,omitempty"`
}

// NewEC2NatGateway constructor creates a new EC2NatGateway
func NewEC2NatGateway(properties EC2NatGatewayProperties, deps ...interface{}) EC2NatGateway {
	return EC2NatGateway{
		Type:       "AWS::EC2::NatGateway",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2NatGateway parses EC2NatGateway
func ParseEC2NatGateway(
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
	var resource EC2NatGateway
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
					"Fn::Sub": "${AWS::StackName}-EC2NatGateway-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2NatGateway validator
func (resource EC2NatGateway) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2NatGatewayProperties validator
func (resource EC2NatGatewayProperties) Validate() []error {
	errors := []error{}
	return errors
}
