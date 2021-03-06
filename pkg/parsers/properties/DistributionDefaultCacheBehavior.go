package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DistributionDefaultCacheBehavior Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-defaultcachebehavior.html
type DistributionDefaultCacheBehavior struct {
	Compress                   interface{} `yaml:"Compress,omitempty"`
	DefaultTTL                 interface{} `yaml:"DefaultTTL,omitempty"`
	FieldLevelEncryptionId     interface{} `yaml:"FieldLevelEncryptionId,omitempty"`
	MaxTTL                     interface{} `yaml:"MaxTTL,omitempty"`
	MinTTL                     interface{} `yaml:"MinTTL,omitempty"`
	SmoothStreaming            interface{} `yaml:"SmoothStreaming,omitempty"`
	TargetOriginId             interface{} `yaml:"TargetOriginId"`
	ViewerProtocolPolicy       interface{} `yaml:"ViewerProtocolPolicy"`
	AllowedMethods             interface{} `yaml:"AllowedMethods,omitempty"`
	CachedMethods              interface{} `yaml:"CachedMethods,omitempty"`
	LambdaFunctionAssociations interface{} `yaml:"LambdaFunctionAssociations,omitempty"`
	TrustedSigners             interface{} `yaml:"TrustedSigners,omitempty"`
	ForwardedValues            interface{} `yaml:"ForwardedValues"`
}

// DistributionDefaultCacheBehavior validation
func (resource DistributionDefaultCacheBehavior) Validate() []error {
	errors := []error{}

	return errors
}
