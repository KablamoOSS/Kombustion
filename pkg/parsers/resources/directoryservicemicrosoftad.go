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

// DirectoryServiceMicrosoftAD Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html
type DirectoryServiceMicrosoftAD struct {
	Type       string                                `yaml:"Type"`
	Properties DirectoryServiceMicrosoftADProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// DirectoryServiceMicrosoftAD Properties
type DirectoryServiceMicrosoftADProperties struct {
	CreateAlias interface{}                        `yaml:"CreateAlias,omitempty"`
	Edition     interface{}                        `yaml:"Edition,omitempty"`
	EnableSso   interface{}                        `yaml:"EnableSso,omitempty"`
	Name        interface{}                        `yaml:"Name"`
	Password    interface{}                        `yaml:"Password"`
	ShortName   interface{}                        `yaml:"ShortName,omitempty"`
	VpcSettings *properties.MicrosoftADVpcSettings `yaml:"VpcSettings"`
}

// NewDirectoryServiceMicrosoftAD constructor creates a new DirectoryServiceMicrosoftAD
func NewDirectoryServiceMicrosoftAD(properties DirectoryServiceMicrosoftADProperties, deps ...interface{}) DirectoryServiceMicrosoftAD {
	return DirectoryServiceMicrosoftAD{
		Type:       "AWS::DirectoryService::MicrosoftAD",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDirectoryServiceMicrosoftAD parses DirectoryServiceMicrosoftAD
func ParseDirectoryServiceMicrosoftAD(name string, data string) (cf types.TemplateObject, err error) {
	var resource DirectoryServiceMicrosoftAD
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DirectoryServiceMicrosoftAD - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseDirectoryServiceMicrosoftAD validator
func (resource DirectoryServiceMicrosoftAD) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDirectoryServiceMicrosoftADProperties validator
func (resource DirectoryServiceMicrosoftADProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Password == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Password'"))
	}
	if resource.VpcSettings == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcSettings'"))
	} else {
		errs = append(errs, resource.VpcSettings.Validate()...)
	}
	return errs
}
