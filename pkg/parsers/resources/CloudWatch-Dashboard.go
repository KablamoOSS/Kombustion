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

// CloudWatchDashboard Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-dashboard.html
type CloudWatchDashboard struct {
	Type       string                        `yaml:"Type"`
	Properties CloudWatchDashboardProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

// CloudWatchDashboard Properties
type CloudWatchDashboardProperties struct {
	DashboardBody interface{} `yaml:"DashboardBody"`
	DashboardName interface{} `yaml:"DashboardName,omitempty"`
}

// NewCloudWatchDashboard constructor creates a new CloudWatchDashboard
func NewCloudWatchDashboard(properties CloudWatchDashboardProperties, deps ...interface{}) CloudWatchDashboard {
	return CloudWatchDashboard{
		Type:       "AWS::CloudWatch::Dashboard",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCloudWatchDashboard parses CloudWatchDashboard
func ParseCloudWatchDashboard(
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
	var resource CloudWatchDashboard
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
					"Fn::Sub": "${AWS::StackName}-CloudWatchDashboard-" + name,
				},
			},
		},
	}

	return
}

// ParseCloudWatchDashboard validator
func (resource CloudWatchDashboard) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCloudWatchDashboardProperties validator
func (resource CloudWatchDashboardProperties) Validate() []error {
	errors := []error{}
	return errors
}
