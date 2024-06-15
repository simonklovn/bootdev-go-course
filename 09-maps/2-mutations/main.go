package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	existingUser, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}	
	if existingUser.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}
type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
