package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// LambdaLayerVersionPermission Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html
type LambdaLayerVersionPermission struct {
	Type       string                                 `yaml:"Type"`
	Properties LambdaLayerVersionPermissionProperties `yaml:"Properties"`
	Condition  interface{}                            `yaml:"Condition,omitempty"`
	Metadata   interface{}                            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                            `yaml:"DependsOn,omitempty"`
}

// LambdaLayerVersionPermission Properties
type LambdaLayerVersionPermissionProperties struct {
	Action          interface{} `yaml:"Action"`
	LayerVersionArn interface{} `yaml:"LayerVersionArn"`
	OrganizationId  interface{} `yaml:"OrganizationId,omitempty"`
	Principal       interface{} `yaml:"Principal"`
}

// NewLambdaLayerVersionPermission constructor creates a new LambdaLayerVersionPermission
func NewLambdaLayerVersionPermission(properties LambdaLayerVersionPermissionProperties, deps ...interface{}) LambdaLayerVersionPermission {
	return LambdaLayerVersionPermission{
		Type:       "AWS::Lambda::LayerVersionPermission",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseLambdaLayerVersionPermission parses LambdaLayerVersionPermission
func ParseLambdaLayerVersionPermission(
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
	var resource LambdaLayerVersionPermission
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
					"Fn::Sub": "${AWS::StackName}-LambdaLayerVersionPermission-" + name,
				},
			},
		},
	}

	return
}

// ParseLambdaLayerVersionPermission validator
func (resource LambdaLayerVersionPermission) Validate() []error {
	return resource.Properties.Validate()
}

// ParseLambdaLayerVersionPermissionProperties validator
func (resource LambdaLayerVersionPermissionProperties) Validate() []error {
	errors := []error{}
	return errors
}
