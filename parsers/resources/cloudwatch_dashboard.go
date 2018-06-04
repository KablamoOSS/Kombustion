package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type CloudWatchDashboard struct {
	Type       string                        `yaml:"Type"`
	Properties CloudWatchDashboardProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

type CloudWatchDashboardProperties struct {
	DashboardBody interface{} `yaml:"DashboardBody"`
	DashboardName interface{} `yaml:"DashboardName,omitempty"`
}

func NewCloudWatchDashboard(properties CloudWatchDashboardProperties, deps ...interface{}) CloudWatchDashboard {
	return CloudWatchDashboard{
		Type:       "AWS::CloudWatch::Dashboard",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCloudWatchDashboard(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource CloudWatchDashboard
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudWatchDashboard - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource CloudWatchDashboard) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CloudWatchDashboardProperties) Validate() []error {
	errs := []error{}
	if resource.DashboardBody == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DashboardBody'"))
	}
	return errs
}
