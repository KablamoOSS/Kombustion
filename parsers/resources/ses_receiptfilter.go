package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type SESReceiptFilter struct {
	Type       string                     `yaml:"Type"`
	Properties SESReceiptFilterProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

type SESReceiptFilterProperties struct {
	Filter *properties.ReceiptFilter_Filter `yaml:"Filter"`
}

func NewSESReceiptFilter(properties SESReceiptFilterProperties, deps ...interface{}) SESReceiptFilter {
	return SESReceiptFilter{
		Type:       "AWS::SES::ReceiptFilter",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSESReceiptFilter(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource SESReceiptFilter
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SESReceiptFilter - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource SESReceiptFilter) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SESReceiptFilterProperties) Validate() []error {
	errs := []error{}
	if resource.Filter == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Filter'"))
	} else {
		errs = append(errs, resource.Filter.Validate()...)
	}
	return errs
}
