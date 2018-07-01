package steamfriendssince

import "sort"

type bySteamFriendsByDateAsc []FriendSince

func (a bySteamFriendsByDateAsc) Len() int {
	return len(a)
}
func (a bySteamFriendsByDateAsc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a bySteamFriendsByDateAsc) Less(i, j int) bool {
	if a[i].FriendSince.UnixTime < a[j].FriendSince.UnixTime {
		return true
	}
	if a[i].FriendSince.UnixTime > a[j].FriendSince.UnixTime {
		return false
	}

	return a[i].PersonaName < a[j].PersonaName
}

/*
	Sorts a FriendSince struct by DateAdded and then by name
	This function creates and returns a copy of the original data
*/
func SortFriendsByDateAscCopy(friends []FriendSince) []FriendSince {
	f := make([]FriendSince, len(friends))
	copy(friends, f)
	sort.Sort(bySteamFriendsByDateAsc(f))
	return f
}

/*
	Sorts a FriendSince struct by DateAdded and then by name
	Modifies the original data
*/
func SortFriendsByDateAsc(friends []FriendSince) []FriendSince {
	sort.Sort(bySteamFriendsByDateAsc(friends))
	return friends
}
