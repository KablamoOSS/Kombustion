package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// BatchJobDefinition Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html
type BatchJobDefinition struct {
	Type       string                       `yaml:"Type"`
	Properties BatchJobDefinitionProperties `yaml:"Properties"`
	Condition  interface{}                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                  `yaml:"DependsOn,omitempty"`
}

// BatchJobDefinition Properties
type BatchJobDefinitionProperties struct {
	JobDefinitionName   interface{} `yaml:"JobDefinitionName,omitempty"`
	Parameters          interface{} `yaml:"Parameters,omitempty"`
	Type                interface{} `yaml:"Type"`
	Timeout             interface{} `yaml:"Timeout,omitempty"`
	RetryStrategy       interface{} `yaml:"RetryStrategy,omitempty"`
	ContainerProperties interface{} `yaml:"ContainerProperties"`
}

// NewBatchJobDefinition constructor creates a new BatchJobDefinition
func NewBatchJobDefinition(properties BatchJobDefinitionProperties, deps ...interface{}) BatchJobDefinition {
	return BatchJobDefinition{
		Type:       "AWS::Batch::JobDefinition",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseBatchJobDefinition parses BatchJobDefinition
func ParseBatchJobDefinition(
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
	var resource BatchJobDefinition
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
					"Fn::Sub": "${AWS::StackName}-BatchJobDefinition-" + name,
				},
			},
		},
	}

	return
}

// ParseBatchJobDefinition validator
func (resource BatchJobDefinition) Validate() []error {
	return resource.Properties.Validate()
}

// ParseBatchJobDefinitionProperties validator
func (resource BatchJobDefinitionProperties) Validate() []error {
	errors := []error{}
	return errors
}
