package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	//
	//
	//"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	//
)

// EC2LaunchTemplate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html
type EC2LaunchTemplate struct {
	Type       string                      `yaml:"Type"`
	Properties EC2LaunchTemplateProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

// EC2LaunchTemplate Properties
type EC2LaunchTemplateProperties struct {
	LaunchTemplateName interface{} `yaml:"LaunchTemplateName,omitempty"`
	LaunchTemplateData interface{} `yaml:"LaunchTemplateData,omitempty"`
}

// NewEC2LaunchTemplate constructor creates a new EC2LaunchTemplate
func NewEC2LaunchTemplate(properties EC2LaunchTemplateProperties, deps ...interface{}) EC2LaunchTemplate {
	return EC2LaunchTemplate{
		Type:       "AWS::EC2::LaunchTemplate",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2LaunchTemplate parses EC2LaunchTemplate
func ParseEC2LaunchTemplate(
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
	var resource EC2LaunchTemplate
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
					"Fn::Sub": "${AWS::StackName}-EC2LaunchTemplate-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2LaunchTemplate validator
func (resource EC2LaunchTemplate) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2LaunchTemplateProperties validator
func (resource EC2LaunchTemplateProperties) Validate() []error {
	errors := []error{}
	return errors
}
