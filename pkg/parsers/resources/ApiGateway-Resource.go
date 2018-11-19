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

// ApiGatewayResource Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html
type ApiGatewayResource struct {
	Type       string                       `yaml:"Type"`
	Properties ApiGatewayResourceProperties `yaml:"Properties"`
	Condition  interface{}                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                  `yaml:"DependsOn,omitempty"`
}

// ApiGatewayResource Properties
type ApiGatewayResourceProperties struct {
	ParentId  interface{} `yaml:"ParentId"`
	PathPart  interface{} `yaml:"PathPart"`
	RestApiId interface{} `yaml:"RestApiId"`
}

// NewApiGatewayResource constructor creates a new ApiGatewayResource
func NewApiGatewayResource(properties ApiGatewayResourceProperties, deps ...interface{}) ApiGatewayResource {
	return ApiGatewayResource{
		Type:       "AWS::ApiGateway::Resource",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayResource parses ApiGatewayResource
func ParseApiGatewayResource(
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
	var resource ApiGatewayResource
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
					"Fn::Sub": "${AWS::StackName}-ApiGatewayResource-" + name,
				},
			},
		},
	}

	return
}

// ParseApiGatewayResource validator
func (resource ApiGatewayResource) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayResourceProperties validator
func (resource ApiGatewayResourceProperties) Validate() []error {
	errors := []error{}
	return errors
}
