package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// LambdaFunction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html
type LambdaFunction struct {
	Type       string                   `yaml:"Type"`
	Properties LambdaFunctionProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// LambdaFunction Properties
type LambdaFunctionProperties struct {
	Description                  interface{} `yaml:"Description,omitempty"`
	FunctionName                 interface{} `yaml:"FunctionName,omitempty"`
	Handler                      interface{} `yaml:"Handler"`
	KmsKeyArn                    interface{} `yaml:"KmsKeyArn,omitempty"`
	MemorySize                   interface{} `yaml:"MemorySize,omitempty"`
	ReservedConcurrentExecutions interface{} `yaml:"ReservedConcurrentExecutions,omitempty"`
	Role                         interface{} `yaml:"Role"`
	Runtime                      interface{} `yaml:"Runtime"`
	Timeout                      interface{} `yaml:"Timeout,omitempty"`
	VpcConfig                    interface{} `yaml:"VpcConfig,omitempty"`
	TracingConfig                interface{} `yaml:"TracingConfig,omitempty"`
	Layers                       interface{} `yaml:"Layers,omitempty"`
	Tags                         interface{} `yaml:"Tags,omitempty"`
	Environment                  interface{} `yaml:"Environment,omitempty"`
	DeadLetterConfig             interface{} `yaml:"DeadLetterConfig,omitempty"`
	Code                         interface{} `yaml:"Code"`
}

// NewLambdaFunction constructor creates a new LambdaFunction
func NewLambdaFunction(properties LambdaFunctionProperties, deps ...interface{}) LambdaFunction {
	return LambdaFunction{
		Type:       "AWS::Lambda::Function",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseLambdaFunction parses LambdaFunction
func ParseLambdaFunction(
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
	var resource LambdaFunction
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
					"Fn::Sub": "${AWS::StackName}-LambdaFunction-" + name,
				},
			},
		},
	}

	return
}

// ParseLambdaFunction validator
func (resource LambdaFunction) Validate() []error {
	return resource.Properties.Validate()
}

// ParseLambdaFunctionProperties validator
func (resource LambdaFunctionProperties) Validate() []error {
	errors := []error{}
	return errors
}
