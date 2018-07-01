package steamfriendssince

import "sort"

type bySteamFriendsByDateDesc []FriendSince

func (a bySteamFriendsByDateDesc) Len() int {
	return len(a)
}
func (a bySteamFriendsByDateDesc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a bySteamFriendsByDateDesc) Less(i, j int) bool {
	if a[i].FriendSince.UnixTime > a[j].FriendSince.UnixTime {
		return true
	}
	if a[i].FriendSince.UnixTime < a[j].FriendSince.UnixTime {
		return false
	}

	return a[i].PersonaName > a[j].PersonaName
}

/*
	Sorts a FriendSince struct by DateAdded and then by name
	This function creates and returns a copy of the original data
*/
func SortFriendsByDateDescCopy(friends []FriendSince) []FriendSince {
	f := make([]FriendSince, len(friends))
	copy(friends, f)
	sort.Sort(bySteamFriendsByDateDesc(f))
	return f
}

/*
	Sorts a FriendSince struct by DateAdded and then by name
	Modifies the original data
*/
func SortFriendsByDateDesc(friends []FriendSince) []FriendSince {
	sort.Sort(bySteamFriendsByDateDesc(friends))
	return friends
}
