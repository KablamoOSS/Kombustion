package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type KMSKey struct {
	Type       string           `yaml:"Type"`
	Properties KMSKeyProperties `yaml:"Properties"`
	Condition  interface{}      `yaml:"Condition,omitempty"`
	Metadata   interface{}      `yaml:"Metadata,omitempty"`
	DependsOn  interface{}      `yaml:"DependsOn,omitempty"`
}

type KMSKeyProperties struct {
	Description       interface{} `yaml:"Description,omitempty"`
	EnableKeyRotation interface{} `yaml:"EnableKeyRotation,omitempty"`
	Enabled           interface{} `yaml:"Enabled,omitempty"`
	KeyPolicy         interface{} `yaml:"KeyPolicy"`
	KeyUsage          interface{} `yaml:"KeyUsage,omitempty"`
	Tags              interface{} `yaml:"Tags,omitempty"`
}

func NewKMSKey(properties KMSKeyProperties, deps ...interface{}) KMSKey {
	return KMSKey{
		Type:       "AWS::KMS::Key",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseKMSKey(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource KMSKey
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: KMSKey - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource KMSKey) Validate() []error {
	return resource.Properties.Validate()
}

func (resource KMSKeyProperties) Validate() []error {
	errs := []error{}
	if resource.KeyPolicy == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeyPolicy'"))
	}
	return errs
}
