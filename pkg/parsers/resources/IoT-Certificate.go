package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// IoTCertificate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificate.html
type IoTCertificate struct {
	Type       string                   `yaml:"Type"`
	Properties IoTCertificateProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// IoTCertificate Properties
type IoTCertificateProperties struct {
	CertificateSigningRequest interface{} `yaml:"CertificateSigningRequest"`
	Status                    interface{} `yaml:"Status"`
}

// NewIoTCertificate constructor creates a new IoTCertificate
func NewIoTCertificate(properties IoTCertificateProperties, deps ...interface{}) IoTCertificate {
	return IoTCertificate{
		Type:       "AWS::IoT::Certificate",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseIoTCertificate parses IoTCertificate
func ParseIoTCertificate(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"

	// Resources
	var resource IoTCertificate
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-IoTCertificate-" + name,
				},
			},
		},
	}

	return
}

// ParseIoTCertificate validator
func (resource IoTCertificate) Validate() []error {
	return resource.Properties.Validate()
}

// ParseIoTCertificateProperties validator
func (resource IoTCertificateProperties) Validate() []error {
	errors := []error{}
	return errors
}
