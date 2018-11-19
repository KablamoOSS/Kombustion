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

// RDSDBSubnetGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html
type RDSDBSubnetGroup struct {
	Type       string                     `yaml:"Type"`
	Properties RDSDBSubnetGroupProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// RDSDBSubnetGroup Properties
type RDSDBSubnetGroupProperties struct {
	DBSubnetGroupDescription interface{} `yaml:"DBSubnetGroupDescription"`
	DBSubnetGroupName        interface{} `yaml:"DBSubnetGroupName,omitempty"`
	SubnetIds                interface{} `yaml:"SubnetIds"`
	Tags                     interface{} `yaml:"Tags,omitempty"`
}

// NewRDSDBSubnetGroup constructor creates a new RDSDBSubnetGroup
func NewRDSDBSubnetGroup(properties RDSDBSubnetGroupProperties, deps ...interface{}) RDSDBSubnetGroup {
	return RDSDBSubnetGroup{
		Type:       "AWS::RDS::DBSubnetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseRDSDBSubnetGroup parses RDSDBSubnetGroup
func ParseRDSDBSubnetGroup(
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
	var resource RDSDBSubnetGroup
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
					"Fn::Sub": "${AWS::StackName}-RDSDBSubnetGroup-" + name,
				},
			},
		},
	}

	return
}

// ParseRDSDBSubnetGroup validator
func (resource RDSDBSubnetGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseRDSDBSubnetGroupProperties validator
func (resource RDSDBSubnetGroupProperties) Validate() []error {
	errors := []error{}
	return errors
}
