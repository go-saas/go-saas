package makeless_go_model

import "sync"

type Team struct {
	Model
	Name *string `gorm:"not null" json:"name"`

	UserId *uint `gorm:"not null" json:"userId"`
	User   *User `json:"user"`

	TeamUsers       []*TeamUser       `json:"teamUsers"`
	TeamInvitations []*TeamInvitation `json:"teamInvitations"`
	Tokens          []*Token          `json:"tokens"`

	*sync.RWMutex
}

func (team *Team) GetId() uint {
	team.RLock()
	defer team.RUnlock()

	return team.Id
}

func (team *Team) GetName() *string {
	team.RLock()
	defer team.RUnlock()

	return team.Name
}

func (team *Team) GetUserId() *uint {
	team.RLock()
	defer team.RUnlock()

	return team.UserId
}

func (team *Team) GetUser() *User {
	team.RLock()
	defer team.RUnlock()

	return team.User
}

func (team *Team) GetTeamUsers() []*TeamUser {
	team.RLock()
	defer team.RUnlock()

	return team.TeamUsers
}

func (team *Team) GetTeamInvitations() []*TeamInvitation {
	team.RLock()
	defer team.RUnlock()

	return team.TeamInvitations
}
