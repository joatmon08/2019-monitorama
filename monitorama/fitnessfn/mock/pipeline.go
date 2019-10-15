package mock

const (
	RoleDeveloper    = "Developer"
	RoleTeamLead     = "Team Lead"
	RoleProductOwner = "Product Owner"
)

type Approver struct {
	Name string
	Role string
}

type Pipeline struct {
	Approvers []*Approver
}

func (p *Pipeline) GetApproversForPipeline() []*Approver {
	return []*Approver{
		&Approver{
			Name: "Dani Lee",
			Role: RoleDeveloper,
		},
		&Approver{
			Name: "MJ Smith",
			Role: RoleDeveloper,
		},
	}
}
