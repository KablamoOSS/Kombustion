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

// InspectorResourceGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-resourcegroup.html
type InspectorResourceGroup struct {
	Type       string                           `yaml:"Type"`
	Properties InspectorResourceGroupProperties `yaml:"Properties"`
	Condition  interface{}                      `yaml:"Condition,omitempty"`
	Metadata   interface{}                      `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                      `yaml:"DependsOn,omitempty"`
}

// InspectorResourceGroup Properties
type InspectorResourceGroupProperties struct {
	ResourceGroupTags interface{} `yaml:"ResourceGroupTags"`
}

// NewInspectorResourceGroup constructor creates a new InspectorResourceGroup
func NewInspectorResourceGroup(properties InspectorResourceGroupProperties, deps ...interface{}) InspectorResourceGroup {
	return InspectorResourceGroup{
		Type:       "AWS::Inspector::ResourceGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseInspectorResourceGroup parses InspectorResourceGroup
func ParseInspectorResourceGroup(
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
	var resource InspectorResourceGroup
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
					"Fn::Sub": "${AWS::StackName}-InspectorResourceGroup-" + name,
				},
			},
		},
	}

	return
}

// ParseInspectorResourceGroup validator
func (resource InspectorResourceGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseInspectorResourceGroupProperties validator
func (resource InspectorResourceGroupProperties) Validate() []error {
	errors := []error{}
	return errors
}
