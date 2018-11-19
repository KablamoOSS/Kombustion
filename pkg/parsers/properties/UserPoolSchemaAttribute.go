package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// UserPoolSchemaAttribute Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html
type UserPoolSchemaAttribute struct {
	AttributeDataType          interface{} `yaml:"AttributeDataType,omitempty"`
	DeveloperOnlyAttribute     interface{} `yaml:"DeveloperOnlyAttribute,omitempty"`
	Mutable                    interface{} `yaml:"Mutable,omitempty"`
	Name                       interface{} `yaml:"Name,omitempty"`
	Required                   interface{} `yaml:"Required,omitempty"`
	StringAttributeConstraints interface{} `yaml:"StringAttributeConstraints,omitempty"`
	NumberAttributeConstraints interface{} `yaml:"NumberAttributeConstraints,omitempty"`
}

// UserPoolSchemaAttribute validation
func (resource UserPoolSchemaAttribute) Validate() []error {
	errors := []error{}

	return errors
}
