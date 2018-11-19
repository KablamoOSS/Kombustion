package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// AppStreamStackFleetAssociation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html
type AppStreamStackFleetAssociation struct {
	Type       string                                   `yaml:"Type"`
	Properties AppStreamStackFleetAssociationProperties `yaml:"Properties"`
	Condition  interface{}                              `yaml:"Condition,omitempty"`
	Metadata   interface{}                              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                              `yaml:"DependsOn,omitempty"`
}

// AppStreamStackFleetAssociation Properties
type AppStreamStackFleetAssociationProperties struct {
	FleetName interface{} `yaml:"FleetName"`
	StackName interface{} `yaml:"StackName"`
}

// NewAppStreamStackFleetAssociation constructor creates a new AppStreamStackFleetAssociation
func NewAppStreamStackFleetAssociation(properties AppStreamStackFleetAssociationProperties, deps ...interface{}) AppStreamStackFleetAssociation {
	return AppStreamStackFleetAssociation{
		Type:       "AWS::AppStream::StackFleetAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAppStreamStackFleetAssociation parses AppStreamStackFleetAssociation
func ParseAppStreamStackFleetAssociation(
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
	var resource AppStreamStackFleetAssociation
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
					"Fn::Sub": "${AWS::StackName}-AppStreamStackFleetAssociation-" + name,
				},
			},
		},
	}

	return
}

// ParseAppStreamStackFleetAssociation validator
func (resource AppStreamStackFleetAssociation) Validate() []error {
	return resource.Properties.Validate()
}

// ParseAppStreamStackFleetAssociationProperties validator
func (resource AppStreamStackFleetAssociationProperties) Validate() []error {
	errors := []error{}
	return errors
}
