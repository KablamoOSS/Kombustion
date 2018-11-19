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

// InspectorAssessmentTarget Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttarget.html
type InspectorAssessmentTarget struct {
	Type       string                              `yaml:"Type"`
	Properties InspectorAssessmentTargetProperties `yaml:"Properties"`
	Condition  interface{}                         `yaml:"Condition,omitempty"`
	Metadata   interface{}                         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                         `yaml:"DependsOn,omitempty"`
}

// InspectorAssessmentTarget Properties
type InspectorAssessmentTargetProperties struct {
	AssessmentTargetName interface{} `yaml:"AssessmentTargetName,omitempty"`
	ResourceGroupArn     interface{} `yaml:"ResourceGroupArn"`
}

// NewInspectorAssessmentTarget constructor creates a new InspectorAssessmentTarget
func NewInspectorAssessmentTarget(properties InspectorAssessmentTargetProperties, deps ...interface{}) InspectorAssessmentTarget {
	return InspectorAssessmentTarget{
		Type:       "AWS::Inspector::AssessmentTarget",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseInspectorAssessmentTarget parses InspectorAssessmentTarget
func ParseInspectorAssessmentTarget(
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
	var resource InspectorAssessmentTarget
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
					"Fn::Sub": "${AWS::StackName}-InspectorAssessmentTarget-" + name,
				},
			},
		},
	}

	return
}

// ParseInspectorAssessmentTarget validator
func (resource InspectorAssessmentTarget) Validate() []error {
	return resource.Properties.Validate()
}

// ParseInspectorAssessmentTargetProperties validator
func (resource InspectorAssessmentTargetProperties) Validate() []error {
	errors := []error{}
	return errors
}
