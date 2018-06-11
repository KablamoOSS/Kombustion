package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
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
	Description                  interface{}                          `yaml:"Description,omitempty"`
	FunctionName                 interface{}                          `yaml:"FunctionName,omitempty"`
	Handler                      interface{}                          `yaml:"Handler"`
	KmsKeyArn                    interface{}                          `yaml:"KmsKeyArn,omitempty"`
	MemorySize                   interface{}                          `yaml:"MemorySize,omitempty"`
	ReservedConcurrentExecutions interface{}                          `yaml:"ReservedConcurrentExecutions,omitempty"`
	Role                         interface{}                          `yaml:"Role"`
	Runtime                      interface{}                          `yaml:"Runtime"`
	Timeout                      interface{}                          `yaml:"Timeout,omitempty"`
	VpcConfig                    *properties.FunctionVpcConfig        `yaml:"VpcConfig,omitempty"`
	TracingConfig                *properties.FunctionTracingConfig    `yaml:"TracingConfig,omitempty"`
	Tags                         interface{}                          `yaml:"Tags,omitempty"`
	Environment                  *properties.FunctionEnvironment      `yaml:"Environment,omitempty"`
	DeadLetterConfig             *properties.FunctionDeadLetterConfig `yaml:"DeadLetterConfig,omitempty"`
	Code                         *properties.FunctionCode             `yaml:"Code"`
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
func ParseLambdaFunction(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource LambdaFunction
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LambdaFunction - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseLambdaFunction validator
func (resource LambdaFunction) Validate() []error {
	return resource.Properties.Validate()
}

// ParseLambdaFunctionProperties validator
func (resource LambdaFunctionProperties) Validate() []error {
	errs := []error{}
	if resource.Handler == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Handler'"))
	}
	if resource.Role == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Role'"))
	}
	if resource.Runtime == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Runtime'"))
	}
	if resource.Code == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Code'"))
	} else {
		errs = append(errs, resource.Code.Validate()...)
	}
	return errs
}
