package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ClusterParameterGroupParameter Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-property-redshift-clusterparametergroup-parameter.html
type ClusterParameterGroupParameter struct {
	ParameterName  interface{} `yaml:"ParameterName"`
	ParameterValue interface{} `yaml:"ParameterValue"`
}

// ClusterParameterGroupParameter validation
func (resource ClusterParameterGroupParameter) Validate() []error {
	errors := []error{}

	return errors
}
