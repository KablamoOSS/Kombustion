package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EMRStep Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html
type EMRStep struct {
	Type       string            `yaml:"Type"`
	Properties EMRStepProperties `yaml:"Properties"`
	Condition  interface{}       `yaml:"Condition,omitempty"`
	Metadata   interface{}       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}       `yaml:"DependsOn,omitempty"`
}

// EMRStep Properties
type EMRStepProperties struct {
	ActionOnFailure interface{}                         `yaml:"ActionOnFailure"`
	JobFlowId       interface{}                         `yaml:"JobFlowId"`
	Name            interface{}                         `yaml:"Name"`
	HadoopJarStep   *properties.StepHadoopJarStepConfig `yaml:"HadoopJarStep"`
}

// NewEMRStep constructor creates a new EMRStep
func NewEMRStep(properties EMRStepProperties, deps ...interface{}) EMRStep {
	return EMRStep{
		Type:       "AWS::EMR::Step",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEMRStep parses EMRStep
func ParseEMRStep(
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
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource EMRStep
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

	return
}

// ParseEMRStep validator
func (resource EMRStep) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEMRStepProperties validator
func (resource EMRStepProperties) Validate() []error {
	errors := []error{}
	if resource.ActionOnFailure == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ActionOnFailure'"))
	}
	if resource.JobFlowId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'JobFlowId'"))
	}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.HadoopJarStep == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'HadoopJarStep'"))
	} else {
		errors = append(errors, resource.HadoopJarStep.Validate()...)
	}
	return errors
}
