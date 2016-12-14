package common

// aws ec2 describe-regions --query 'Regions[].{Name:RegionName}' --output text | sed 's/\(.*\)/"\1",/' | sort
func listEC2Regions() []string {
	return []string{
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-south-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"ca-central-1",
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"sa-east-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		// not part of autogenerated list
		"us-gov-west-1",
		"cn-north-1",
	}
}

// ValidateRegion returns true if the supplied region is a valid AWS
// region and false if it's not.
func ValidateRegion(region string) bool {
	for _, valid := range listEC2Regions() {
		if region == valid {
			return true
		}
	}
	return false
}
