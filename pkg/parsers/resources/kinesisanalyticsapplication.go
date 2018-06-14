package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// KinesisAnalyticsApplication Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html
type KinesisAnalyticsApplication struct {
	Type       string                                `yaml:"Type"`
	Properties KinesisAnalyticsApplicationProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// KinesisAnalyticsApplication Properties
type KinesisAnalyticsApplicationProperties struct {
	ApplicationCode        interface{} `yaml:"ApplicationCode,omitempty"`
	ApplicationDescription interface{} `yaml:"ApplicationDescription,omitempty"`
	ApplicationName        interface{} `yaml:"ApplicationName,omitempty"`
	Inputs                 interface{} `yaml:"Inputs"`
}

// NewKinesisAnalyticsApplication constructor creates a new KinesisAnalyticsApplication
func NewKinesisAnalyticsApplication(properties KinesisAnalyticsApplicationProperties, deps ...interface{}) KinesisAnalyticsApplication {
	return KinesisAnalyticsApplication{
		Type:       "AWS::KinesisAnalytics::Application",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseKinesisAnalyticsApplication parses KinesisAnalyticsApplication
func ParseKinesisAnalyticsApplication(name string, data string) (cf types.TemplateObject, err error) {
	var resource KinesisAnalyticsApplication
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: KinesisAnalyticsApplication - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseKinesisAnalyticsApplication validator
func (resource KinesisAnalyticsApplication) Validate() []error {
	return resource.Properties.Validate()
}

// ParseKinesisAnalyticsApplicationProperties validator
func (resource KinesisAnalyticsApplicationProperties) Validate() []error {
	errs := []error{}
	if resource.Inputs == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Inputs'"))
	}
	return errs
}
