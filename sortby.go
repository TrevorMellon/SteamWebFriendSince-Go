package steamfriendssince

import "sort"

type SortByType int

const (
	SortByUnixTime SortByType = iota
	SortByYear
)

func SortBy(friends []FriendSince, sorttype SortByType) []FriendSince {
	if sorttype == SortByUnixTime {
		return sortbyunixtime(friends)
	}
	return nil
}

type sortbyunixtimeasc []FriendSince

func (a sortbyunixtimeasc) Less(i, j int) bool {
	if a[i].FriendSince.UnixTime < a[j].FriendSince.UnixTime {
		return true
	}
	if a[i].FriendSince.UnixTime > a[j].FriendSince.UnixTime {
		return false
	}

	return a[i].PersonaName < a[j].PersonaName
}
func (a sortbyunixtimeasc) Len() int {
	return len(a)
}
func (a sortbyunixtimeasc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func sortbyunixtime(friends []FriendSince) []FriendSince {
	sort.Sort(sortbyunixtimeasc(friends))
	return friends
}
