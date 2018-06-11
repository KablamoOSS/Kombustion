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

// ElasticBeanstalkConfigurationTemplate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html
type ElasticBeanstalkConfigurationTemplate struct {
	Type       string                                          `yaml:"Type"`
	Properties ElasticBeanstalkConfigurationTemplateProperties `yaml:"Properties"`
	Condition  interface{}                                     `yaml:"Condition,omitempty"`
	Metadata   interface{}                                     `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                     `yaml:"DependsOn,omitempty"`
}

// ElasticBeanstalkConfigurationTemplate Properties
type ElasticBeanstalkConfigurationTemplateProperties struct {
	ApplicationName     interface{}                                          `yaml:"ApplicationName"`
	Description         interface{}                                          `yaml:"Description,omitempty"`
	EnvironmentId       interface{}                                          `yaml:"EnvironmentId,omitempty"`
	PlatformArn         interface{}                                          `yaml:"PlatformArn,omitempty"`
	SolutionStackName   interface{}                                          `yaml:"SolutionStackName,omitempty"`
	SourceConfiguration *properties.ConfigurationTemplateSourceConfiguration `yaml:"SourceConfiguration,omitempty"`
	OptionSettings      interface{}                                          `yaml:"OptionSettings,omitempty"`
}

// NewElasticBeanstalkConfigurationTemplate constructor creates a new ElasticBeanstalkConfigurationTemplate
func NewElasticBeanstalkConfigurationTemplate(properties ElasticBeanstalkConfigurationTemplateProperties, deps ...interface{}) ElasticBeanstalkConfigurationTemplate {
	return ElasticBeanstalkConfigurationTemplate{
		Type:       "AWS::ElasticBeanstalk::ConfigurationTemplate",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElasticBeanstalkConfigurationTemplate parses ElasticBeanstalkConfigurationTemplate
func ParseElasticBeanstalkConfigurationTemplate(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ElasticBeanstalkConfigurationTemplate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticBeanstalkConfigurationTemplate - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseElasticBeanstalkConfigurationTemplate validator
func (resource ElasticBeanstalkConfigurationTemplate) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElasticBeanstalkConfigurationTemplateProperties validator
func (resource ElasticBeanstalkConfigurationTemplateProperties) Validate() []error {
	errs := []error{}
	if resource.ApplicationName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApplicationName'"))
	}
	return errs
}
