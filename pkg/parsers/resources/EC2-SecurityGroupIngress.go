package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2SecurityGroupIngress Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html
type EC2SecurityGroupIngress struct {
	Type       string                            `yaml:"Type"`
	Properties EC2SecurityGroupIngressProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// EC2SecurityGroupIngress Properties
type EC2SecurityGroupIngressProperties struct {
	CidrIp                     interface{} `yaml:"CidrIp,omitempty"`
	CidrIpv6                   interface{} `yaml:"CidrIpv6,omitempty"`
	Description                interface{} `yaml:"Description,omitempty"`
	FromPort                   interface{} `yaml:"FromPort,omitempty"`
	GroupId                    interface{} `yaml:"GroupId,omitempty"`
	GroupName                  interface{} `yaml:"GroupName,omitempty"`
	IpProtocol                 interface{} `yaml:"IpProtocol"`
	SourceSecurityGroupId      interface{} `yaml:"SourceSecurityGroupId,omitempty"`
	SourceSecurityGroupName    interface{} `yaml:"SourceSecurityGroupName,omitempty"`
	SourceSecurityGroupOwnerId interface{} `yaml:"SourceSecurityGroupOwnerId,omitempty"`
	ToPort                     interface{} `yaml:"ToPort,omitempty"`
}

// NewEC2SecurityGroupIngress constructor creates a new EC2SecurityGroupIngress
func NewEC2SecurityGroupIngress(properties EC2SecurityGroupIngressProperties, deps ...interface{}) EC2SecurityGroupIngress {
	return EC2SecurityGroupIngress{
		Type:       "AWS::EC2::SecurityGroupIngress",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2SecurityGroupIngress parses EC2SecurityGroupIngress
func ParseEC2SecurityGroupIngress(
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
	var resource EC2SecurityGroupIngress
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

// ParseEC2SecurityGroupIngress validator
func (resource EC2SecurityGroupIngress) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2SecurityGroupIngressProperties validator
func (resource EC2SecurityGroupIngressProperties) Validate() []error {
	errors := []error{}
	if resource.IpProtocol == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'IpProtocol'"))
	}
	return errors
}
