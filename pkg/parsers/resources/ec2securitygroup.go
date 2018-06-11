package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2SecurityGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html
type EC2SecurityGroup struct {
	Type       string                     `yaml:"Type"`
	Properties EC2SecurityGroupProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// EC2SecurityGroup Properties
type EC2SecurityGroupProperties struct {
	GroupDescription     interface{} `yaml:"GroupDescription"`
	GroupName            interface{} `yaml:"GroupName,omitempty"`
	VpcId                interface{} `yaml:"VpcId,omitempty"`
	SecurityGroupEgress  interface{} `yaml:"SecurityGroupEgress,omitempty"`
	SecurityGroupIngress interface{} `yaml:"SecurityGroupIngress,omitempty"`
	Tags                 interface{} `yaml:"Tags,omitempty"`
}

// NewEC2SecurityGroup constructor creates a new EC2SecurityGroup
func NewEC2SecurityGroup(properties EC2SecurityGroupProperties, deps ...interface{}) EC2SecurityGroup {
	return EC2SecurityGroup{
		Type:       "AWS::EC2::SecurityGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2SecurityGroup parses EC2SecurityGroup
func ParseEC2SecurityGroup(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2SecurityGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2SecurityGroup - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEC2SecurityGroup validator
func (resource EC2SecurityGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2SecurityGroupProperties validator
func (resource EC2SecurityGroupProperties) Validate() []error {
	errs := []error{}
	if resource.GroupDescription == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'GroupDescription'"))
	}
	return errs
}
