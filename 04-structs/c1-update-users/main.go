package main

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type membershipType
	MessageCharLimit int
}

func newUser(name string, membershipType membershipType) User {
	var charLimit int
	if membershipType == TypePremium {
		charLimit = 1000
	} else {
		charLimit = 100
	}
		
	return User{
		Name: name,
		Membership: Membership{
			Type: membershipType,
			MessageCharLimit: charLimit,
		},
	}
}
