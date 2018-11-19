package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayMethod Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html
type ApiGatewayMethod struct {
	Type       string                     `yaml:"Type"`
	Properties ApiGatewayMethodProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// ApiGatewayMethod Properties
type ApiGatewayMethodProperties struct {
	ApiKeyRequired      interface{} `yaml:"ApiKeyRequired,omitempty"`
	AuthorizationType   interface{} `yaml:"AuthorizationType,omitempty"`
	AuthorizerId        interface{} `yaml:"AuthorizerId,omitempty"`
	HttpMethod          interface{} `yaml:"HttpMethod"`
	OperationName       interface{} `yaml:"OperationName,omitempty"`
	RequestValidatorId  interface{} `yaml:"RequestValidatorId,omitempty"`
	ResourceId          interface{} `yaml:"ResourceId"`
	RestApiId           interface{} `yaml:"RestApiId"`
	RequestModels       interface{} `yaml:"RequestModels,omitempty"`
	RequestParameters   interface{} `yaml:"RequestParameters,omitempty"`
	AuthorizationScopes interface{} `yaml:"AuthorizationScopes,omitempty"`
	MethodResponses     interface{} `yaml:"MethodResponses,omitempty"`
	Integration         interface{} `yaml:"Integration,omitempty"`
}

// NewApiGatewayMethod constructor creates a new ApiGatewayMethod
func NewApiGatewayMethod(properties ApiGatewayMethodProperties, deps ...interface{}) ApiGatewayMethod {
	return ApiGatewayMethod{
		Type:       "AWS::ApiGateway::Method",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayMethod parses ApiGatewayMethod
func ParseApiGatewayMethod(
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
	var resource ApiGatewayMethod
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
					"Fn::Sub": "${AWS::StackName}-ApiGatewayMethod-" + name,
				},
			},
		},
	}

	return
}

// ParseApiGatewayMethod validator
func (resource ApiGatewayMethod) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayMethodProperties validator
func (resource ApiGatewayMethodProperties) Validate() []error {
	errors := []error{}
	return errors
}
