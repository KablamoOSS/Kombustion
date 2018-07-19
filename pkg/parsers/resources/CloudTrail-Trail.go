package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// CloudTrailTrail Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html
type CloudTrailTrail struct {
	Type       string                    `yaml:"Type"`
	Properties CloudTrailTrailProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

// CloudTrailTrail Properties
type CloudTrailTrailProperties struct {
	CloudWatchLogsLogGroupArn  interface{} `yaml:"CloudWatchLogsLogGroupArn,omitempty"`
	CloudWatchLogsRoleArn      interface{} `yaml:"CloudWatchLogsRoleArn,omitempty"`
	EnableLogFileValidation    interface{} `yaml:"EnableLogFileValidation,omitempty"`
	IncludeGlobalServiceEvents interface{} `yaml:"IncludeGlobalServiceEvents,omitempty"`
	IsLogging                  interface{} `yaml:"IsLogging"`
	IsMultiRegionTrail         interface{} `yaml:"IsMultiRegionTrail,omitempty"`
	KMSKeyId                   interface{} `yaml:"KMSKeyId,omitempty"`
	S3BucketName               interface{} `yaml:"S3BucketName"`
	S3KeyPrefix                interface{} `yaml:"S3KeyPrefix,omitempty"`
	SnsTopicName               interface{} `yaml:"SnsTopicName,omitempty"`
	TrailName                  interface{} `yaml:"TrailName,omitempty"`
	EventSelectors             interface{} `yaml:"EventSelectors,omitempty"`
	Tags                       interface{} `yaml:"Tags,omitempty"`
}

// NewCloudTrailTrail constructor creates a new CloudTrailTrail
func NewCloudTrailTrail(properties CloudTrailTrailProperties, deps ...interface{}) CloudTrailTrail {
	return CloudTrailTrail{
		Type:       "AWS::CloudTrail::Trail",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCloudTrailTrail parses CloudTrailTrail
func ParseCloudTrailTrail(
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
	var resource CloudTrailTrail
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

// ParseCloudTrailTrail validator
func (resource CloudTrailTrail) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCloudTrailTrailProperties validator
func (resource CloudTrailTrailProperties) Validate() []error {
	errors := []error{}
	if resource.IsLogging == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'IsLogging'"))
	}
	if resource.S3BucketName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'S3BucketName'"))
	}
	return errors
}
