package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ServiceDiscoveryInstance Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-instance.html
type ServiceDiscoveryInstance struct {
	Type       string                             `yaml:"Type"`
	Properties ServiceDiscoveryInstanceProperties `yaml:"Properties"`
	Condition  interface{}                        `yaml:"Condition,omitempty"`
	Metadata   interface{}                        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                        `yaml:"DependsOn,omitempty"`
}

// ServiceDiscoveryInstance Properties
type ServiceDiscoveryInstanceProperties struct {
	InstanceAttributes interface{} `yaml:"InstanceAttributes"`
	InstanceId         interface{} `yaml:"InstanceId,omitempty"`
	ServiceId          interface{} `yaml:"ServiceId"`
}

// NewServiceDiscoveryInstance constructor creates a new ServiceDiscoveryInstance
func NewServiceDiscoveryInstance(properties ServiceDiscoveryInstanceProperties, deps ...interface{}) ServiceDiscoveryInstance {
	return ServiceDiscoveryInstance{
		Type:       "AWS::ServiceDiscovery::Instance",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceDiscoveryInstance parses ServiceDiscoveryInstance
func ParseServiceDiscoveryInstance(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ServiceDiscoveryInstance
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ServiceDiscoveryInstance - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseServiceDiscoveryInstance validator
func (resource ServiceDiscoveryInstance) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceDiscoveryInstanceProperties validator
func (resource ServiceDiscoveryInstanceProperties) Validate() []error {
	errs := []error{}
	if resource.InstanceAttributes == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceAttributes'"))
	}
	if resource.ServiceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceId'"))
	}
	return errs
}
