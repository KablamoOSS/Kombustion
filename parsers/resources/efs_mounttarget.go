package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type EFSMountTarget struct {
	Type       string                   `yaml:"Type"`
	Properties EFSMountTargetProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

type EFSMountTargetProperties struct {
	FileSystemId   interface{} `yaml:"FileSystemId"`
	IpAddress      interface{} `yaml:"IpAddress,omitempty"`
	SubnetId       interface{} `yaml:"SubnetId"`
	SecurityGroups interface{} `yaml:"SecurityGroups"`
}

func NewEFSMountTarget(properties EFSMountTargetProperties, deps ...interface{}) EFSMountTarget {
	return EFSMountTarget{
		Type:       "AWS::EFS::MountTarget",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEFSMountTarget(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource EFSMountTarget
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EFSMountTarget - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource EFSMountTarget) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EFSMountTargetProperties) Validate() []error {
	errs := []error{}
	if resource.FileSystemId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FileSystemId'"))
	}
	if resource.SubnetId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetId'"))
	}
	if resource.SecurityGroups == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SecurityGroups'"))
	}
	return errs
}
