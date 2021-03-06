package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// IoT1ClickProject Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html
type IoT1ClickProject struct {
	Type       string                     `yaml:"Type"`
	Properties IoT1ClickProjectProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// IoT1ClickProject Properties
type IoT1ClickProjectProperties struct {
	Description       interface{} `yaml:"Description,omitempty"`
	ProjectName       interface{} `yaml:"ProjectName,omitempty"`
	PlacementTemplate interface{} `yaml:"PlacementTemplate"`
}

// NewIoT1ClickProject constructor creates a new IoT1ClickProject
func NewIoT1ClickProject(properties IoT1ClickProjectProperties, deps ...interface{}) IoT1ClickProject {
	return IoT1ClickProject{
		Type:       "AWS::IoT1Click::Project",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseIoT1ClickProject parses IoT1ClickProject
func ParseIoT1ClickProject(
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
	var resource IoT1ClickProject
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
					"Fn::Sub": "${AWS::StackName}-IoT1ClickProject-" + name,
				},
			},
		},
	}

	return
}

// ParseIoT1ClickProject validator
func (resource IoT1ClickProject) Validate() []error {
	return resource.Properties.Validate()
}

// ParseIoT1ClickProjectProperties validator
func (resource IoT1ClickProjectProperties) Validate() []error {
	errors := []error{}
	return errors
}
