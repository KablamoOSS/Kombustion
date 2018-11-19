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

// LogsLogStream Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html
type LogsLogStream struct {
	Type       string                  `yaml:"Type"`
	Properties LogsLogStreamProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// LogsLogStream Properties
type LogsLogStreamProperties struct {
	LogGroupName  interface{} `yaml:"LogGroupName"`
	LogStreamName interface{} `yaml:"LogStreamName,omitempty"`
}

// NewLogsLogStream constructor creates a new LogsLogStream
func NewLogsLogStream(properties LogsLogStreamProperties, deps ...interface{}) LogsLogStream {
	return LogsLogStream{
		Type:       "AWS::Logs::LogStream",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseLogsLogStream parses LogsLogStream
func ParseLogsLogStream(
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
	var resource LogsLogStream
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
					"Fn::Sub": "${AWS::StackName}-LogsLogStream-" + name,
				},
			},
		},
	}

	return
}

// ParseLogsLogStream validator
func (resource LogsLogStream) Validate() []error {
	return resource.Properties.Validate()
}

// ParseLogsLogStreamProperties validator
func (resource LogsLogStreamProperties) Validate() []error {
	errors := []error{}
	return errors
}
