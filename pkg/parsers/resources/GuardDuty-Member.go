package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// GuardDutyMember Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html
type GuardDutyMember struct {
	Type       string                    `yaml:"Type"`
	Properties GuardDutyMemberProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

// GuardDutyMember Properties
type GuardDutyMemberProperties struct {
	DetectorId               interface{} `yaml:"DetectorId"`
	DisableEmailNotification interface{} `yaml:"DisableEmailNotification,omitempty"`
	Email                    interface{} `yaml:"Email"`
	MemberId                 interface{} `yaml:"MemberId"`
	Message                  interface{} `yaml:"Message,omitempty"`
	Status                   interface{} `yaml:"Status,omitempty"`
}

// NewGuardDutyMember constructor creates a new GuardDutyMember
func NewGuardDutyMember(properties GuardDutyMemberProperties, deps ...interface{}) GuardDutyMember {
	return GuardDutyMember{
		Type:       "AWS::GuardDuty::Member",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseGuardDutyMember parses GuardDutyMember
func ParseGuardDutyMember(
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
	var resource GuardDutyMember
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

// ParseGuardDutyMember validator
func (resource GuardDutyMember) Validate() []error {
	return resource.Properties.Validate()
}

// ParseGuardDutyMemberProperties validator
func (resource GuardDutyMemberProperties) Validate() []error {
	errors := []error{}
	if resource.DetectorId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DetectorId'"))
	}
	if resource.Email == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Email'"))
	}
	if resource.MemberId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'MemberId'"))
	}
	return errors
}
