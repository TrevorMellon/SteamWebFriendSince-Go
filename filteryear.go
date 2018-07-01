package steamfriendssince

func FilterByYear(friends []FriendSince, year int) []FriendSince {
	var count int

	for _, v := range friends {
		if v.FriendSince.Year == year {
			count++
		}
	}

	ret := make([]FriendSince, count)

	idx := 0

	for _, v := range friends {
		if v.FriendSince.Year == year {
			ret[idx] = v
			idx++
		}
	}
	return ret
}
