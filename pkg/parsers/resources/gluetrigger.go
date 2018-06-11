package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// GlueTrigger Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html
type GlueTrigger struct {
	Type       string                `yaml:"Type"`
	Properties GlueTriggerProperties `yaml:"Properties"`
	Condition  interface{}           `yaml:"Condition,omitempty"`
	Metadata   interface{}           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}           `yaml:"DependsOn,omitempty"`
}

// GlueTrigger Properties
type GlueTriggerProperties struct {
	Description interface{}                  `yaml:"Description,omitempty"`
	Name        interface{}                  `yaml:"Name,omitempty"`
	Schedule    interface{}                  `yaml:"Schedule,omitempty"`
	Type        interface{}                  `yaml:"Type"`
	Predicate   *properties.TriggerPredicate `yaml:"Predicate,omitempty"`
	Actions     interface{}                  `yaml:"Actions"`
}

// NewGlueTrigger constructor creates a new GlueTrigger
func NewGlueTrigger(properties GlueTriggerProperties, deps ...interface{}) GlueTrigger {
	return GlueTrigger{
		Type:       "AWS::Glue::Trigger",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseGlueTrigger parses GlueTrigger
func ParseGlueTrigger(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource GlueTrigger
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueTrigger - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseGlueTrigger validator
func (resource GlueTrigger) Validate() []error {
	return resource.Properties.Validate()
}

// ParseGlueTriggerProperties validator
func (resource GlueTriggerProperties) Validate() []error {
	errs := []error{}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.Actions == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Actions'"))
	}
	return errs
}
