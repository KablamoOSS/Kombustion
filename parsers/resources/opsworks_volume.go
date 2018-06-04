package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type OpsWorksVolume struct {
	Type       string                   `yaml:"Type"`
	Properties OpsWorksVolumeProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

type OpsWorksVolumeProperties struct {
	Ec2VolumeId interface{} `yaml:"Ec2VolumeId"`
	MountPoint  interface{} `yaml:"MountPoint,omitempty"`
	Name        interface{} `yaml:"Name,omitempty"`
	StackId     interface{} `yaml:"StackId"`
}

func NewOpsWorksVolume(properties OpsWorksVolumeProperties, deps ...interface{}) OpsWorksVolume {
	return OpsWorksVolume{
		Type:       "AWS::OpsWorks::Volume",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseOpsWorksVolume(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource OpsWorksVolume
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: OpsWorksVolume - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource OpsWorksVolume) Validate() []error {
	return resource.Properties.Validate()
}

func (resource OpsWorksVolumeProperties) Validate() []error {
	errs := []error{}
	if resource.Ec2VolumeId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Ec2VolumeId'"))
	}
	if resource.StackId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StackId'"))
	}
	return errs
}
