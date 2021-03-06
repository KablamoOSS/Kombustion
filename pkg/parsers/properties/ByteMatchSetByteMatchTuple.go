package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ByteMatchSetByteMatchTuple Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html
type ByteMatchSetByteMatchTuple struct {
	PositionalConstraint interface{} `yaml:"PositionalConstraint"`
	TargetString         interface{} `yaml:"TargetString,omitempty"`
	TargetStringBase64   interface{} `yaml:"TargetStringBase64,omitempty"`
	TextTransformation   interface{} `yaml:"TextTransformation"`
	FieldToMatch         interface{} `yaml:"FieldToMatch"`
}

// ByteMatchSetByteMatchTuple validation
func (resource ByteMatchSetByteMatchTuple) Validate() []error {
	errors := []error{}

	return errors
}
