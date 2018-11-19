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

// EC2VPCEndpoint Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html
type EC2VPCEndpoint struct {
	Type       string                   `yaml:"Type"`
	Properties EC2VPCEndpointProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// EC2VPCEndpoint Properties
type EC2VPCEndpointProperties struct {
	PolicyDocument    interface{} `yaml:"PolicyDocument,omitempty"`
	PrivateDnsEnabled interface{} `yaml:"PrivateDnsEnabled,omitempty"`
	ServiceName       interface{} `yaml:"ServiceName"`
	VpcEndpointType   interface{} `yaml:"VpcEndpointType,omitempty"`
	VpcId             interface{} `yaml:"VpcId"`
	RouteTableIds     interface{} `yaml:"RouteTableIds,omitempty"`
	SecurityGroupIds  interface{} `yaml:"SecurityGroupIds,omitempty"`
	SubnetIds         interface{} `yaml:"SubnetIds,omitempty"`
}

// NewEC2VPCEndpoint constructor creates a new EC2VPCEndpoint
func NewEC2VPCEndpoint(properties EC2VPCEndpointProperties, deps ...interface{}) EC2VPCEndpoint {
	return EC2VPCEndpoint{
		Type:       "AWS::EC2::VPCEndpoint",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2VPCEndpoint parses EC2VPCEndpoint
func ParseEC2VPCEndpoint(
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
	var resource EC2VPCEndpoint
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
					"Fn::Sub": "${AWS::StackName}-EC2VPCEndpoint-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2VPCEndpoint validator
func (resource EC2VPCEndpoint) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2VPCEndpointProperties validator
func (resource EC2VPCEndpointProperties) Validate() []error {
	errors := []error{}
	return errors
}
