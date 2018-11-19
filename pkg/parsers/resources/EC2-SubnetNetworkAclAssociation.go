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

// EC2SubnetNetworkAclAssociation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html
type EC2SubnetNetworkAclAssociation struct {
	Type       string                                   `yaml:"Type"`
	Properties EC2SubnetNetworkAclAssociationProperties `yaml:"Properties"`
	Condition  interface{}                              `yaml:"Condition,omitempty"`
	Metadata   interface{}                              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                              `yaml:"DependsOn,omitempty"`
}

// EC2SubnetNetworkAclAssociation Properties
type EC2SubnetNetworkAclAssociationProperties struct {
	NetworkAclId interface{} `yaml:"NetworkAclId"`
	SubnetId     interface{} `yaml:"SubnetId"`
}

// NewEC2SubnetNetworkAclAssociation constructor creates a new EC2SubnetNetworkAclAssociation
func NewEC2SubnetNetworkAclAssociation(properties EC2SubnetNetworkAclAssociationProperties, deps ...interface{}) EC2SubnetNetworkAclAssociation {
	return EC2SubnetNetworkAclAssociation{
		Type:       "AWS::EC2::SubnetNetworkAclAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2SubnetNetworkAclAssociation parses EC2SubnetNetworkAclAssociation
func ParseEC2SubnetNetworkAclAssociation(
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
	var resource EC2SubnetNetworkAclAssociation
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
					"Fn::Sub": "${AWS::StackName}-EC2SubnetNetworkAclAssociation-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2SubnetNetworkAclAssociation validator
func (resource EC2SubnetNetworkAclAssociation) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2SubnetNetworkAclAssociationProperties validator
func (resource EC2SubnetNetworkAclAssociationProperties) Validate() []error {
	errors := []error{}
	return errors
}
