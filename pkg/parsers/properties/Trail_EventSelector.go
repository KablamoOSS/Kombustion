package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type Trail_EventSelector struct {
	IncludeManagementEvents interface{} `yaml:"IncludeManagementEvents,omitempty"`
	ReadWriteType           interface{} `yaml:"ReadWriteType,omitempty"`
	DataResources           interface{} `yaml:"DataResources,omitempty"`
}

func (resource Trail_EventSelector) Validate() []error {
	errs := []error{}

	return errs
}