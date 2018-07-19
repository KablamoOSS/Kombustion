package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// CodePipelineWebhook Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html
type CodePipelineWebhook struct {
	Type       string                        `yaml:"Type"`
	Properties CodePipelineWebhookProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

// CodePipelineWebhook Properties
type CodePipelineWebhookProperties struct {
	Authentication              interface{}                                 `yaml:"Authentication"`
	Name                        interface{}                                 `yaml:"Name,omitempty"`
	RegisterWithThirdParty      interface{}                                 `yaml:"RegisterWithThirdParty,omitempty"`
	TargetAction                interface{}                                 `yaml:"TargetAction"`
	TargetPipeline              interface{}                                 `yaml:"TargetPipeline"`
	TargetPipelineVersion       interface{}                                 `yaml:"TargetPipelineVersion"`
	AuthenticationConfiguration *properties.WebhookWebhookAuthConfiguration `yaml:"AuthenticationConfiguration"`
	Filters                     interface{}                                 `yaml:"Filters"`
}

// NewCodePipelineWebhook constructor creates a new CodePipelineWebhook
func NewCodePipelineWebhook(properties CodePipelineWebhookProperties, deps ...interface{}) CodePipelineWebhook {
	return CodePipelineWebhook{
		Type:       "AWS::CodePipeline::Webhook",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCodePipelineWebhook parses CodePipelineWebhook
func ParseCodePipelineWebhook(
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
	var resource CodePipelineWebhook
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

// ParseCodePipelineWebhook validator
func (resource CodePipelineWebhook) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCodePipelineWebhookProperties validator
func (resource CodePipelineWebhookProperties) Validate() []error {
	errors := []error{}
	if resource.Authentication == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Authentication'"))
	}
	if resource.TargetAction == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'TargetAction'"))
	}
	if resource.TargetPipeline == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'TargetPipeline'"))
	}
	if resource.TargetPipelineVersion == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'TargetPipelineVersion'"))
	}
	if resource.AuthenticationConfiguration == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'AuthenticationConfiguration'"))
	} else {
		errors = append(errors, resource.AuthenticationConfiguration.Validate()...)
	}
	if resource.Filters == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Filters'"))
	}
	return errors
}
