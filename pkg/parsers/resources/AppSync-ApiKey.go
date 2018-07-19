package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// AppSyncApiKey Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apikey.html
type AppSyncApiKey struct {
	Type       string                  `yaml:"Type"`
	Properties AppSyncApiKeyProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// AppSyncApiKey Properties
type AppSyncApiKeyProperties struct {
	ApiId       interface{} `yaml:"ApiId"`
	Description interface{} `yaml:"Description,omitempty"`
	Expires     interface{} `yaml:"Expires,omitempty"`
}

// NewAppSyncApiKey constructor creates a new AppSyncApiKey
func NewAppSyncApiKey(properties AppSyncApiKeyProperties, deps ...interface{}) AppSyncApiKey {
	return AppSyncApiKey{
		Type:       "AWS::AppSync::ApiKey",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAppSyncApiKey parses AppSyncApiKey
func ParseAppSyncApiKey(
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
	var resource AppSyncApiKey
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

// ParseAppSyncApiKey validator
func (resource AppSyncApiKey) Validate() []error {
	return resource.Properties.Validate()
}

// ParseAppSyncApiKeyProperties validator
func (resource AppSyncApiKeyProperties) Validate() []error {
	errors := []error{}
	if resource.ApiId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ApiId'"))
	}
	return errors
}
