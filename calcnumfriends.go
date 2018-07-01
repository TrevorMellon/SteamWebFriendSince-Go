package steamfriendssince

import steamapi "github.com/Philipp15b/go-steamapi"

/*
FriendsStatistics
The basic friends states
*/
type FriendsStatistics struct {
	FriendsNumber         uint64
	FriendsOnline         uint64
	FriendsOffline        uint64
	FriendsAway           uint64
	FriendsBusy           uint64
	FriendsSnoozing       uint64
	FriendsLookingToTrade uint64
	FriendsLookingToPlay  uint64
	FriendsPublic         uint64
	FriendsPrivate        uint64
	FriendsOnly           uint64
}

/*
CalcNumberOfFriends
Calculates the states and numbers of Steam friends
*/
func CalcNumberOfFriends(friends []FriendSince) FriendsStatistics {
	n := len(friends)
	var f FriendsStatistics

	f.FriendsNumber = uint64(n)

	for _, v := range friends {
		summ := v.Summary
		if summ.PersonaState == steamapi.Online {
			f.FriendsOnline++
		}
		if summ.PersonaState == steamapi.Offline {
			f.FriendsOffline++
		}
		if summ.PersonaState == steamapi.Away {
			f.FriendsAway++
		}
		if summ.PersonaState == steamapi.Busy {
			f.FriendsBusy++
		}
		if summ.PersonaState == steamapi.Snooze {
			f.FriendsSnoozing++
		}
		if summ.PersonaState == steamapi.LookingToTrade {
			f.FriendsLookingToTrade++
		}
		if summ.PersonaState == steamapi.LookingToPlay {
			f.FriendsLookingToPlay++
		}
		if summ.CommunityVisibilityState == steamapi.Private {
			f.FriendsPrivate++
		}
		if summ.CommunityVisibilityState == steamapi.FriendsOnly {
			f.FriendsOnly++
		}
		if summ.CommunityVisibilityState == steamapi.Public {
			f.FriendsPublic++
		}
	}
	return f
}
