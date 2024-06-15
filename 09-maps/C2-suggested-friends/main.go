package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	// Collect the users direct friends
	directFriends, ok := friendships[username]
	if !ok {
		return nil
	}

	// Track suggested friends to avoid duplicates
	suggestedFriendsMap := make(map[string]bool)

	// Set of direct friends for quick lookup
	directFriendsSet := make(map[string]bool)

	for _, friend := range directFriends {
		directFriendsSet[friend] = true
	}

	// locate friends of friends who are not direct friends
	for _, directFriend := range directFriends{
		friendsOfFriend, ok := friendships[directFriend]
		if !ok {
			continue
		}
		for _, friendOfFriend := range friendsOfFriend {
			if friendOfFriend != username && !directFriendsSet[friendOfFriend] {
				suggestedFriendsMap[friendOfFriend] = true
			}
		}
	}

	// Convert map keys to slice
	var suggestedFriends[]string
	for friend := range suggestedFriendsMap {
		suggestedFriends = append(suggestedFriends, friend)
	}
	if len(suggestedFriends) == 0 {
		return nil
	}
	return suggestedFriends
}
