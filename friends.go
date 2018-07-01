package steamfriendssince

import (
	"time"

	steamapi "github.com/Philipp15b/go-steamapi"
)

type Friends struct {
	LowestYear  int
	HighestYear int
	Friends     []FriendSince
}

/*
Combine PlayerSummary with Date Added
*/
type FriendSince struct {
	SteamID     uint64
	PersonaName string
	RealName    string
	//FriendSince       int64
	FriendSince DateSince

	FriendSinceString string
	Summary           steamapi.PlayerSummary
	OnlineColor       int
}

type DateSince struct {
	UnixTime int64
	Year     int
	Month    int
	Day      int
	Hour     int
	Minute   int
}

/*
	returns a Friends since struct
	Merging Player Summary (names etc) with date added
*/
func GetMySteamFriends(apikey string, steamid uint64) Friends {

	ret, _ := steamapi.GetFriendsList(steamid, steamapi.Friend, apikey)

	var number_of_ids uint64 = uint64(len(ret))

	ids := make([]uint64, number_of_ids)
	friends := make([]FriendSince, number_of_ids)

	var friend Friends

	friend.Friends = friends

	for key, value := range ret {

		var v steamapi.SteamFriend
		v = value
		fid := v.SteamID

		ids[key] = fid
		friends[key].SteamID = fid
		friends[key].FriendSince.UnixTime = v.FriendSince
		t := time.Unix(v.FriendSince, 0)
		friends[key].FriendSinceString = t.Format("1" + " Jan" + " 2006")

		friends[key].FriendSince.Year = t.Year()
		friends[key].FriendSince.Month = int(t.Month())
		friends[key].FriendSince.Day = t.Day()
		friends[key].FriendSince.Hour = t.Hour()
		friends[key].FriendSince.Minute = t.Minute()

		if key == 0 {
			friend.LowestYear = t.Year()
			friend.HighestYear = t.Year()
		} else {
			if t.Year() > friend.HighestYear {
				friend.HighestYear = t.Year()
			}
			if t.Year() < friend.LowestYear {
				friend.LowestYear = t.Year()
			}
		}

		iarr := make([]uint64, 1)
		iarr[0] = fid

	}
	var i, last int
	i = 0

	last = 100
	if last > len(ids) {
		last = len(ids)
	}
	for i = 0; i < len(ids); i += 100 {
		sendids := ids[i:last]
		rr, _ := steamapi.GetPlayerSummaries(sendids, apikey)
		for _, value := range rr {
			var ps steamapi.PlayerSummary
			ps = value
			//fmt.Println(ps.PersonaName)
			for k, v := range friends {
				if v.SteamID == value.SteamID {
					friends[k].PersonaName = ps.PersonaName
					friends[k].RealName = ps.RealName
					friends[k].Summary = ps
				}
			}

		}
		last += 100
		if last > len(ids) {
			last = len(ids)
		}
	}

	/*for _, value := range rr {
		var ps steamapi.PlayerSummary
		ps = value
		//fmt.Println(ps.PersonaName)
		for k, v := range friends {
			if v.SteamID == value.SteamID {
				friends[k].PersonaName = ps.PersonaName
				friends[k].RealName = ps.RealName
				friends[k].Summary = ps
			}
		}

	}*/

	for k, v := range friends {
		if v.PersonaName == "" {
			friends[k].PersonaName = "* Private Profile *"
		}
	}

	/*for _, value := range friends {
		fmt.Println(value.PersonaName)
		t := time.Unix(value.FriendSince, 0)
		fmt.Println(t)
	}*/

	return friend
}
