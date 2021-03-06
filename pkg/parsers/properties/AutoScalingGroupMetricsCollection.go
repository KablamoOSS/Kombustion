package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// AutoScalingGroupMetricsCollection Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-metricscollection.html
type AutoScalingGroupMetricsCollection struct {
	Granularity interface{} `yaml:"Granularity"`
	Metrics     interface{} `yaml:"Metrics,omitempty"`
}

// AutoScalingGroupMetricsCollection validation
func (resource AutoScalingGroupMetricsCollection) Validate() []error {
	errors := []error{}

	return errors
}
