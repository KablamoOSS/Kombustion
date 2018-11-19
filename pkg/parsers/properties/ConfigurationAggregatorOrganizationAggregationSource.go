package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ConfigurationAggregatorOrganizationAggregationSource Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-organizationaggregationsource.html
type ConfigurationAggregatorOrganizationAggregationSource struct {
	AllAwsRegions interface{} `yaml:"AllAwsRegions,omitempty"`
	RoleArn       interface{} `yaml:"RoleArn"`
	AwsRegions    interface{} `yaml:"AwsRegions,omitempty"`
}

// ConfigurationAggregatorOrganizationAggregationSource validation
func (resource ConfigurationAggregatorOrganizationAggregationSource) Validate() []error {
	errors := []error{}

	return errors
}
