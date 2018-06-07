package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type DeliveryStream_ElasticsearchBufferingHints struct {
	IntervalInSeconds interface{} `yaml:"IntervalInSeconds"`
	SizeInMBs         interface{} `yaml:"SizeInMBs"`
}

func (resource DeliveryStream_ElasticsearchBufferingHints) Validate() []error {
	errs := []error{}

	if resource.IntervalInSeconds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IntervalInSeconds'"))
	}
	if resource.SizeInMBs == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SizeInMBs'"))
	}
	return errs
}