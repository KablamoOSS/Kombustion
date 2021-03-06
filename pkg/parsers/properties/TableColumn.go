package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TableColumn Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-column.html
type TableColumn struct {
	Comment interface{} `yaml:"Comment,omitempty"`
	Name    interface{} `yaml:"Name"`
	Type    interface{} `yaml:"Type,omitempty"`
}

// TableColumn validation
func (resource TableColumn) Validate() []error {
	errors := []error{}

	return errors
}
