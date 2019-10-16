package mock

import (
	"time"
)

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
	StartTime time.Time
	EndTime   time.Time
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

func (p *Pipeline) GetDeploymentTime() time.Duration {
	return p.EndTime.Sub(p.StartTime)
}
