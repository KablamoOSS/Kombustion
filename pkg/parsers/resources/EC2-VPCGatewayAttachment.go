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

// EC2VPCGatewayAttachment Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html
type EC2VPCGatewayAttachment struct {
	Type       string                            `yaml:"Type"`
	Properties EC2VPCGatewayAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// EC2VPCGatewayAttachment Properties
type EC2VPCGatewayAttachmentProperties struct {
	InternetGatewayId interface{} `yaml:"InternetGatewayId,omitempty"`
	VpcId             interface{} `yaml:"VpcId"`
	VpnGatewayId      interface{} `yaml:"VpnGatewayId,omitempty"`
}

// NewEC2VPCGatewayAttachment constructor creates a new EC2VPCGatewayAttachment
func NewEC2VPCGatewayAttachment(properties EC2VPCGatewayAttachmentProperties, deps ...interface{}) EC2VPCGatewayAttachment {
	return EC2VPCGatewayAttachment{
		Type:       "AWS::EC2::VPCGatewayAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2VPCGatewayAttachment parses EC2VPCGatewayAttachment
func ParseEC2VPCGatewayAttachment(
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
	var resource EC2VPCGatewayAttachment
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
					"Fn::Sub": "${AWS::StackName}-EC2VPCGatewayAttachment-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2VPCGatewayAttachment validator
func (resource EC2VPCGatewayAttachment) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2VPCGatewayAttachmentProperties validator
func (resource EC2VPCGatewayAttachmentProperties) Validate() []error {
	errors := []error{}
	return errors
}
