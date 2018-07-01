package steamfriendssince

import steamapi "github.com/Philipp15b/go-steamapi"

type Filter int

const (
	FilterNone Filter = iota
	FilterOnlineFriends
	FilterOfflineFriends
	FilterAwayFriends
	FilterBusyFriends
	FilterPublicFriends
	FilterPrivateFriends
	FilterFriendsOnlyFriends
)

func FilterFriends(friends []FriendSince, filter Filter) []FriendSince {
	//n := len(friends)
	stats := CalcNumberOfFriends(friends)

	if filter == FilterOnlineFriends {
		f := make([]FriendSince, stats.FriendsOnline)
		idx := 0
		for _, v := range friends {
			if v.Summary.PersonaState == steamapi.Online {
				f[idx] = v
				idx++
			}
		}
		return f
	}
	if filter == FilterOfflineFriends {
		f := make([]FriendSince, stats.FriendsOffline)
		idx := 0
		for _, v := range friends {
			if v.Summary.PersonaState == steamapi.Offline {
				f[idx] = v
				idx++
			}
		}
		return f
	}
	if filter == FilterAwayFriends {
		f := make([]FriendSince, stats.FriendsAway)
		idx := 0
		for _, v := range friends {
			if v.Summary.PersonaState == steamapi.Away {
				f[idx] = v
				idx++
			}
		}
		return f
	}
	if filter == FilterBusyFriends {
		f := make([]FriendSince, stats.FriendsBusy)
		idx := 0
		for _, v := range friends {
			if v.Summary.PersonaState == steamapi.Busy {
				f[idx] = v
				idx++
			}
		}
		return f
	}
	if filter == FilterPublicFriends {
		f := make([]FriendSince, stats.FriendsPublic)
		idx := 0
		for _, v := range friends {
			if v.Summary.CommunityVisibilityState == steamapi.Public {
				f[idx] = v
				idx++
			}
		}
		return f
	}
	if filter == FilterPrivateFriends {
		f := make([]FriendSince, stats.FriendsPrivate)
		idx := 0
		for _, v := range friends {
			if v.Summary.CommunityVisibilityState == steamapi.Private {
				f[idx] = v
				idx++
			}
		}
		return f
	}
	if filter == FilterFriendsOnlyFriends {
		f := make([]FriendSince, stats.FriendsOnly)
		idx := 0
		for _, v := range friends {
			if v.Summary.CommunityVisibilityState == steamapi.FriendsOnly {
				f[idx] = v
				idx++
			}
		}
		return f
	}
	return nil
}
