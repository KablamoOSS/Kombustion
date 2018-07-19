package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// DMSEventSubscription Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-eventsubscription.html
type DMSEventSubscription struct {
	Type       string                         `yaml:"Type"`
	Properties DMSEventSubscriptionProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// DMSEventSubscription Properties
type DMSEventSubscriptionProperties struct {
	Enabled          interface{} `yaml:"Enabled,omitempty"`
	SnsTopicArn      interface{} `yaml:"SnsTopicArn"`
	SourceType       interface{} `yaml:"SourceType,omitempty"`
	SubscriptionName interface{} `yaml:"SubscriptionName,omitempty"`
	EventCategories  interface{} `yaml:"EventCategories,omitempty"`
	SourceIds        interface{} `yaml:"SourceIds,omitempty"`
	Tags             interface{} `yaml:"Tags,omitempty"`
}

// NewDMSEventSubscription constructor creates a new DMSEventSubscription
func NewDMSEventSubscription(properties DMSEventSubscriptionProperties, deps ...interface{}) DMSEventSubscription {
	return DMSEventSubscription{
		Type:       "AWS::DMS::EventSubscription",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDMSEventSubscription parses DMSEventSubscription
func ParseDMSEventSubscription(
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
	var resource DMSEventSubscription
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

// ParseDMSEventSubscription validator
func (resource DMSEventSubscription) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDMSEventSubscriptionProperties validator
func (resource DMSEventSubscriptionProperties) Validate() []error {
	errors := []error{}
	if resource.SnsTopicArn == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'SnsTopicArn'"))
	}
	return errors
}
