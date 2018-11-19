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

// ApplicationAutoScalingScalingPolicy Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html
type ApplicationAutoScalingScalingPolicy struct {
	Type       string                                        `yaml:"Type"`
	Properties ApplicationAutoScalingScalingPolicyProperties `yaml:"Properties"`
	Condition  interface{}                                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                   `yaml:"DependsOn,omitempty"`
}

// ApplicationAutoScalingScalingPolicy Properties
type ApplicationAutoScalingScalingPolicyProperties struct {
	PolicyName                               interface{} `yaml:"PolicyName"`
	PolicyType                               interface{} `yaml:"PolicyType"`
	ResourceId                               interface{} `yaml:"ResourceId,omitempty"`
	ScalableDimension                        interface{} `yaml:"ScalableDimension,omitempty"`
	ScalingTargetId                          interface{} `yaml:"ScalingTargetId,omitempty"`
	ServiceNamespace                         interface{} `yaml:"ServiceNamespace,omitempty"`
	TargetTrackingScalingPolicyConfiguration interface{} `yaml:"TargetTrackingScalingPolicyConfiguration,omitempty"`
	StepScalingPolicyConfiguration           interface{} `yaml:"StepScalingPolicyConfiguration,omitempty"`
}

// NewApplicationAutoScalingScalingPolicy constructor creates a new ApplicationAutoScalingScalingPolicy
func NewApplicationAutoScalingScalingPolicy(properties ApplicationAutoScalingScalingPolicyProperties, deps ...interface{}) ApplicationAutoScalingScalingPolicy {
	return ApplicationAutoScalingScalingPolicy{
		Type:       "AWS::ApplicationAutoScaling::ScalingPolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApplicationAutoScalingScalingPolicy parses ApplicationAutoScalingScalingPolicy
func ParseApplicationAutoScalingScalingPolicy(
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
	var resource ApplicationAutoScalingScalingPolicy
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
					"Fn::Sub": "${AWS::StackName}-ApplicationAutoScalingScalingPolicy-" + name,
				},
			},
		},
	}

	return
}

// ParseApplicationAutoScalingScalingPolicy validator
func (resource ApplicationAutoScalingScalingPolicy) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApplicationAutoScalingScalingPolicyProperties validator
func (resource ApplicationAutoScalingScalingPolicyProperties) Validate() []error {
	errors := []error{}
	return errors
}
