package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// DocDBDBSubnetGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbsubnetgroup.html
type DocDBDBSubnetGroup struct {
	Type       string                       `yaml:"Type"`
	Properties DocDBDBSubnetGroupProperties `yaml:"Properties"`
	Condition  interface{}                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                  `yaml:"DependsOn,omitempty"`
}

// DocDBDBSubnetGroup Properties
type DocDBDBSubnetGroupProperties struct {
	DBSubnetGroupDescription interface{} `yaml:"DBSubnetGroupDescription"`
	DBSubnetGroupName        interface{} `yaml:"DBSubnetGroupName,omitempty"`
	SubnetIds                interface{} `yaml:"SubnetIds"`
	Tags                     interface{} `yaml:"Tags,omitempty"`
}

// NewDocDBDBSubnetGroup constructor creates a new DocDBDBSubnetGroup
func NewDocDBDBSubnetGroup(properties DocDBDBSubnetGroupProperties, deps ...interface{}) DocDBDBSubnetGroup {
	return DocDBDBSubnetGroup{
		Type:       "AWS::DocDB::DBSubnetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDocDBDBSubnetGroup parses DocDBDBSubnetGroup
func ParseDocDBDBSubnetGroup(
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
	var resource DocDBDBSubnetGroup
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
					"Fn::Sub": "${AWS::StackName}-DocDBDBSubnetGroup-" + name,
				},
			},
		},
	}

	return
}

// ParseDocDBDBSubnetGroup validator
func (resource DocDBDBSubnetGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDocDBDBSubnetGroupProperties validator
func (resource DocDBDBSubnetGroupProperties) Validate() []error {
	errors := []error{}
	return errors
}
