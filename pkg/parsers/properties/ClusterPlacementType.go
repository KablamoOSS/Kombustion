package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ClusterPlacementType Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-placementtype.html
type ClusterPlacementType struct {
	AvailabilityZone interface{} `yaml:"AvailabilityZone"`
}

// ClusterPlacementType validation
func (resource ClusterPlacementType) Validate() []error {
	errors := []error{}

	return errors
}
