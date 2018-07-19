package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2VPNGateway Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gateway.html
type EC2VPNGateway struct {
	Type       string                  `yaml:"Type"`
	Properties EC2VPNGatewayProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// EC2VPNGateway Properties
type EC2VPNGatewayProperties struct {
	AmazonSideAsn interface{} `yaml:"AmazonSideAsn,omitempty"`
	Type          interface{} `yaml:"Type"`
	Tags          interface{} `yaml:"Tags,omitempty"`
}

// NewEC2VPNGateway constructor creates a new EC2VPNGateway
func NewEC2VPNGateway(properties EC2VPNGatewayProperties, deps ...interface{}) EC2VPNGateway {
	return EC2VPNGateway{
		Type:       "AWS::EC2::VPNGateway",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2VPNGateway parses EC2VPNGateway
func ParseEC2VPNGateway(
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
	var resource EC2VPNGateway
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

// ParseEC2VPNGateway validator
func (resource EC2VPNGateway) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2VPNGatewayProperties validator
func (resource EC2VPNGatewayProperties) Validate() []error {
	errors := []error{}
	if resource.Type == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Type'"))
	}
	return errors
}
