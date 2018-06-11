package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// CloudFormationWaitConditionHandle Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitconditionhandle.html
type CloudFormationWaitConditionHandle struct {
	Type       string                                      `yaml:"Type"`
	Properties CloudFormationWaitConditionHandleProperties `yaml:"Properties"`
	Condition  interface{}                                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                 `yaml:"DependsOn,omitempty"`
}

// CloudFormationWaitConditionHandle Properties
type CloudFormationWaitConditionHandleProperties struct {
}

// NewCloudFormationWaitConditionHandle constructor creates a new CloudFormationWaitConditionHandle
func NewCloudFormationWaitConditionHandle(properties CloudFormationWaitConditionHandleProperties, deps ...interface{}) CloudFormationWaitConditionHandle {
	return CloudFormationWaitConditionHandle{
		Type:       "AWS::CloudFormation::WaitConditionHandle",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCloudFormationWaitConditionHandle parses CloudFormationWaitConditionHandle
func ParseCloudFormationWaitConditionHandle(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource CloudFormationWaitConditionHandle
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudFormationWaitConditionHandle - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseCloudFormationWaitConditionHandle validator
func (resource CloudFormationWaitConditionHandle) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCloudFormationWaitConditionHandleProperties validator
func (resource CloudFormationWaitConditionHandleProperties) Validate() []error {
	errs := []error{}
	return errs
}
