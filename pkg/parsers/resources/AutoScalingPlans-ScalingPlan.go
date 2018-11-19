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
	//"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	//
)

// AutoScalingPlansScalingPlan Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html
type AutoScalingPlansScalingPlan struct {
	Type       string                                `yaml:"Type"`
	Properties AutoScalingPlansScalingPlanProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// AutoScalingPlansScalingPlan Properties
type AutoScalingPlansScalingPlanProperties struct {
	ScalingInstructions interface{} `yaml:"ScalingInstructions"`
	ApplicationSource   interface{} `yaml:"ApplicationSource"`
}

// NewAutoScalingPlansScalingPlan constructor creates a new AutoScalingPlansScalingPlan
func NewAutoScalingPlansScalingPlan(properties AutoScalingPlansScalingPlanProperties, deps ...interface{}) AutoScalingPlansScalingPlan {
	return AutoScalingPlansScalingPlan{
		Type:       "AWS::AutoScalingPlans::ScalingPlan",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAutoScalingPlansScalingPlan parses AutoScalingPlansScalingPlan
func ParseAutoScalingPlansScalingPlan(
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
	var resource AutoScalingPlansScalingPlan
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
					"Fn::Sub": "${AWS::StackName}-AutoScalingPlansScalingPlan-" + name,
				},
			},
		},
	}

	return
}

// ParseAutoScalingPlansScalingPlan validator
func (resource AutoScalingPlansScalingPlan) Validate() []error {
	return resource.Properties.Validate()
}

// ParseAutoScalingPlansScalingPlanProperties validator
func (resource AutoScalingPlansScalingPlanProperties) Validate() []error {
	errors := []error{}
	return errors
}
