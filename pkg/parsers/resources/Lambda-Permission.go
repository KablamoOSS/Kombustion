package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// LambdaPermission Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html
type LambdaPermission struct {
	Type       string                     `yaml:"Type"`
	Properties LambdaPermissionProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// LambdaPermission Properties
type LambdaPermissionProperties struct {
	Action           interface{} `yaml:"Action"`
	EventSourceToken interface{} `yaml:"EventSourceToken,omitempty"`
	FunctionName     interface{} `yaml:"FunctionName"`
	Principal        interface{} `yaml:"Principal"`
	SourceAccount    interface{} `yaml:"SourceAccount,omitempty"`
	SourceArn        interface{} `yaml:"SourceArn,omitempty"`
}

// NewLambdaPermission constructor creates a new LambdaPermission
func NewLambdaPermission(properties LambdaPermissionProperties, deps ...interface{}) LambdaPermission {
	return LambdaPermission{
		Type:       "AWS::Lambda::Permission",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseLambdaPermission parses LambdaPermission
func ParseLambdaPermission(
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
	var resource LambdaPermission
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
					"Fn::Sub": "${AWS::StackName}-LambdaPermission-" + name,
				},
			},
		},
	}

	return
}

// ParseLambdaPermission validator
func (resource LambdaPermission) Validate() []error {
	return resource.Properties.Validate()
}

// ParseLambdaPermissionProperties validator
func (resource LambdaPermissionProperties) Validate() []error {
	errors := []error{}
	return errors
}
