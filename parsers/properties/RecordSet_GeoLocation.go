package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type RecordSet_GeoLocation struct {
	ContinentCode   interface{} `yaml:"ContinentCode,omitempty"`
	CountryCode     interface{} `yaml:"CountryCode,omitempty"`
	SubdivisionCode interface{} `yaml:"SubdivisionCode,omitempty"`
}

func (resource RecordSet_GeoLocation) Validate() []error {
	errs := []error{}

	return errs
}
