package securitygroup

type SecurityGroupsCreator struct {
	SecurityGroups []List `yaml:"security-groups"`
}

type List struct {
	Name  string `yaml:"name"`
	Desc  string `yaml:"desc"`
	VpcID string `yaml:"vpc-id"`
}
