package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ConfigurationTagsEntry Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-configuration-tagsentry.html
type ConfigurationTagsEntry struct {
	Key   interface{} `yaml:"Key"`
	Value interface{} `yaml:"Value"`
}

// ConfigurationTagsEntry validation
func (resource ConfigurationTagsEntry) Validate() []error {
	errors := []error{}

	return errors
}
