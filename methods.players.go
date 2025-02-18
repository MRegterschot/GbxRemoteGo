package main

import (
	"errors"

	"github.com/MRegterschot/GbxRemoteGo/structs"
)

// Ban the player with the specified login, with an optional message. Only available to Admin.
func (client *GbxClient) Ban(login string, reason string) error {
	_, err := client.Call("Ban", login, reason)
	return err
}

// Ban the player with the specified login, with a message. Add it to the black list, and optionally save the new list. Only available to Admin.
func (client *GbxClient) BanAndBlackList(login string, reason string, save bool) error {
	_, err := client.Call("BanAndBlackList", login, reason, save)
	return err
}

// Ban the player with the specified PlayerId, with an optional message. Only available to Admin.
func (client *GbxClient) BanId(id int, reason string) error {
	_, err := client.Call("BanId", id, reason)
	return err
}

// Unban the player with the specified login. Only available to Admin.
func (client *GbxClient) UnBan(login string) error {
	_, err := client.Call("UnBan", login)
	return err
}

// Returns the list of banned players. This method takes two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the list. The list is an array of structures. Each structure contains the following fields : Login, ClientName and IPAddress.
func (client *GbxClient) GetBanList(max int, start int) ([]structs.TMBanListEntry, error) {
	res, err := client.Call("GetBanList", max, start)
	if err != nil {
		return []structs.TMBanListEntry{}, err
	}

	var banList []structs.TMBanListEntry
	err = convertToStruct(res, &banList)
	if err != nil {
		return []structs.TMBanListEntry{}, err
	}

	return banList, nil
}

// Clean the ban list of the server. Only available to Admin.
func (client *GbxClient) CleanBanList() error {
	_, err := client.Call("CleanBanList")
	return err
}

// Blacklist the player with the specified login. Only available to SuperAdmin.
func (client *GbxClient) BlackList(login string) error {
	_, err := client.Call("BlackList", login)
	return err
}

// Blacklist the player with the specified PlayerId. Only available to SuperAdmin.
func (client *GbxClient) BlackListId(id int) error {
	_, err := client.Call("BlackListId", id)
	return err
}

// UnTMBlackListEntry the player with the specified login. Only available to SuperAdmin.
func (client *GbxClient) UnBlackList(login string) error {
	_, err := client.Call("UnBlackList", login)
	return err
}

// Returns the list of blacklisted players. This method takes two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the list. The list is an array of structures. Each structure contains the following fields : Login.
func (client *GbxClient) GetBlackList(max int, start int) ([]structs.TMBlackListEntry, error) {
	res, err := client.Call("GetBlackList", max, start)
	if err != nil {
		return []structs.TMBlackListEntry{}, err
	}

	var blackList []structs.TMBlackListEntry
	err = convertToStruct(res, &blackList)
	if err != nil {
		return []structs.TMBlackListEntry{}, err
	}

	return blackList, nil
}

// Load the black list file with the specified file name. Only available to Admin.
func (client *GbxClient) LoadBlackList(fileName string) error {
	_, err := client.Call("LoadBlackList", fileName)
	return err
}

// Save the black list in the file with specified file name. Only available to Admin.
func (client *GbxClient) SaveBlackList(fileName string) error {
	_, err := client.Call("SaveBlackList", fileName)
	return err
}

// Clean the blacklist of the server. Only available to SuperAdmin.
func (client *GbxClient) CleanBlackList() error {
	_, err := client.Call("CleanBlackList")
	return err
}

// (debug tool) Connect a fake player to the server. Only available to Admin.
func (client *GbxClient) ConnectFakePlayer() error {
	_, err := client.Send("ConnectFakePlayer")
	return err
}

// (debug tool) Disconnect a fake player, or all the fake players if login is '*'. Only available to Admin.
func (client *GbxClient) DisconnectFakePlayer(login string) error {
	_, err := client.Send("DisconnectFakePlayer", login)
	return err
}

// Add the player with the specified login on the guest list. Only available to Admin.
func (client *GbxClient) AddGuest(login string) error {
	_, err := client.Call("AddGuest", login)
	return err
}

// Add the player with the specified PlayerId on the guest list. Only available to Admin.
func (client *GbxClient) AddGuestId(id int) error {
	_, err := client.Call("AddGuestId", id)
	return err
}

// Remove the player with the specified login from the guest list. Only available to Admin.
func (client *GbxClient) RemoveGuest(login string) error {
	_, err := client.Call("RemoveGuest", login)
	return err
}

// Returns the list of players on the guest list. This method takes two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the list. The list is an array of structures. Each structure contains the following fields : Login.
func (client *GbxClient) GetGuestList(max int, start int) ([]structs.TMGuestListEntry, error) {
	res, err := client.Call("GetGuestList", max, start)
	if err != nil {
		return []structs.TMGuestListEntry{}, err
	}

	var guestList []structs.TMGuestListEntry
	err = convertToStruct(res, &guestList)
	if err != nil {
		return []structs.TMGuestListEntry{}, err
	}

	return guestList, nil
}

// Load the guest list file with the specified file name. Only available to Admin.
func (client *GbxClient) LoadGuestList(fileName string) error {
	_, err := client.Call("LoadGuestList", fileName)
	return err
}

// Remove the player with the specified PlayerId from the guest list. Only available to Admin.
func (client *GbxClient) RemoveGuestId(id int) error {
	_, err := client.Call("RemoveGuestId", id)
	return err
}

// Save the guest list in the file with specified file name. Only available to Admin.
func (client *GbxClient) SaveGuestList(fileName string) error {
	_, err := client.Call("SaveGuestList", fileName)
	return err
}

// Clean the guest list of the server. Only available to Admin.
func (client *GbxClient) CleanGuestList() error {
	_, err := client.Call("CleanGuestList")
	return err
}

// Kick the player with the specified login, with an optional message. Only available to Admin.
func (client *GbxClient) Kick(login string, reason string) error {
	_, err := client.Call("Kick", login, reason)
	return err
}

// Kick the player with the specified PlayerId, with an optional message. Only available to Admin.
func (client *GbxClient) KickId(id int, reason string) error {
	_, err := client.Call("KickId", id, reason)
	return err
}

// Returns the list of players on the server. This method take two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the list, an optional 3rd parameter is used for compatibility: struct version (0 = united, 1 = forever, 2 = forever, including the servers). The list is an array of PlayerInfo structures. Forever PlayerInfo struct is: Login, NickName, PlayerId, TeamId, SpectatorStatus, LadderRanking, and Flags.
// LadderRanking is 0 when not in official mode,
// Flags = ForceSpectator(0,1,2) + StereoDisplayMode * 1000 + IsManagedByAnOtherServer * 10000 + IsServer * 100000 + HasPlayerSlot * 1000000 + IsBroadcasting * 10000000 + HasJoinedGame * 100000000
// SpectatorStatus = Spectator + TemporarySpectator * 10 + PureSpectator * 100 + AutoTarget * 1000 + CurrentTargetId * 10000
func (client *GbxClient) GetPlayerList(max int, start int, version ...int) ([]structs.TMPlayerInfo, error) {
	var param int = -1
	if len(version) > 0 {
		param = version[0]
	}
	res, err := client.Call("GetPlayerList", max, start, param)
	if err != nil {
		return []structs.TMPlayerInfo{}, err
	}

	var playerList []structs.TMPlayerInfo
	err = convertToStruct(res, &playerList)
	if err != nil {
		return []structs.TMPlayerInfo{}, err
	}

	return playerList, nil
}

// Returns a struct containing the infos on the player with the specified login, with an optional parameter for compatibility: struct version (0 = united, 1 = forever). The structure is identical to the ones from GetPlayerList. Forever PlayerInfo struct is: Login, NickName, PlayerId, TeamId, SpectatorStatus, LadderRanking, and Flags.
// LadderRanking is 0 when not in official mode,
// Flags = ForceSpectator(0,1,2) + StereoDisplayMode * 1000 + IsManagedByAnOtherServer * 10000 + IsServer * 100000 + HasPlayerSlot * 1000000 + IsBroadcasting * 10000000 + HasJoinedGame * 100000000
// SpectatorStatus = Spectator + TemporarySpectator * 10 + PureSpectator * 100 + AutoTarget * 1000 + CurrentTargetId * 10000
func (client *GbxClient) GetPlayerInfo(login string, version ...int) (structs.TMPlayerInfo, error) {
	var param int = -1
	if len(version) > 0 {
		param = version[0]
	}
	res, err := client.Call("GetPlayerInfo", login, param)
	if err != nil {
		return structs.TMPlayerInfo{}, err
	}

	var playerInfo structs.TMPlayerInfo
	err = convertToStruct(res, &playerInfo)
	if err != nil {
		return structs.TMPlayerInfo{}, err
	}

	return playerInfo, nil
}

// Returns a struct containing the infos on the player with the specified login. The structure contains the following fields : Login, NickName, PlayerId, TeamId, IPAddress, DownloadRate, UploadRate, Language, IsSpectator, IsInOfficialMode, a structure named Avatar, an array of structures named Skins, a structure named LadderStats, HoursSinceZoneInscription and OnlineRights (0: nations account, 3: united account). Each structure of the array Skins contains two fields Environnement and a struct PackDesc. Each structure PackDesc, as well as the struct Avatar, contains two fields FileName and Checksum.
func (client *GbxClient) GetDetailedPlayerInfo(login string) (structs.TMPlayerDetailedInfo, error) {
	res, err := client.Call("GetDetailedPlayerInfo", login)
	if err != nil {
		return structs.TMPlayerDetailedInfo{}, err
	}

	var playerInfo structs.TMPlayerDetailedInfo
	err = convertToStruct(res, &playerInfo)
	if err != nil {
		return structs.TMPlayerDetailedInfo{}, err
	}

	return playerInfo, nil
}

// Returns a struct containing the player infos of the game server (ie: in case of a basic server, itself; in case of a relay server, the main server), with an optional parameter for compatibility: struct version (0 = united, 1 = forever). The structure is identical to the ones from GetPlayerList. Forever PlayerInfo struct is: Login, NickName, PlayerId, TeamId, SpectatorStatus, LadderRanking, and Flags.
// LadderRanking is 0 when not in official mode,
// Flags = ForceSpectator(0,1,2) + StereoDisplayMode * 1000 + IsManagedByAnOtherServer * 10000 + IsServer * 100000 + HasPlayerSlot * 1000000 + IsBroadcasting * 10000000 + HasJoinedGame * 100000000
// SpectatorStatus = Spectator + TemporarySpectator * 10 + PureSpectator * 100 + AutoTarget * 1000 + CurrentTargetId * 10000
func (client *GbxClient) GetMainServerPlayerInfo(version ...int) (structs.TMPlayerInfo, error) {
	var param int = -1
	if len(version) > 0 {
		param = version[0]
	}
	res, err := client.Call("GetMainServerPlayerInfo", param)
	if err != nil {
		return structs.TMPlayerInfo{}, err
	}

	var playerInfo structs.TMPlayerInfo
	err = convertToStruct(res, &playerInfo)
	if err != nil {
		return structs.TMPlayerInfo{}, err
	}

	return playerInfo, nil
}

// Returns the current rankings for the race in progress. (In trackmania legacy team modes, the scores for the two teams are returned. In other modes, it's the individual players' scores) This method take two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the ranking. The ranking returned is a list of structures. Each structure contains the following fields : Login, NickName, PlayerId and Rank. In addition, for legacy trackmania modes it also contains BestTime, Score, NbrLapsFinished, LadderScore, and an array BestCheckpoints that contains the checkpoint times for the best race.
func (client *GbxClient) GetCurrentRanking(max int, start int) ([]structs.TMPlayerRanking, error) {
	res, err := client.Call("GetCurrentRanking", max, start)
	if err != nil {
		return []structs.TMPlayerRanking{}, err
	}

	var playerList []structs.TMPlayerRanking
	err = convertToStruct(res, &playerList)
	if err != nil {
		return []structs.TMPlayerRanking{}, err
	}

	return playerList, nil
}

// Returns the current ranking for the race in progressof the player with the specified login (or list of comma-separated logins). The ranking returned is a list of structures. Each structure contains the following fields : Login, NickName, PlayerId and Rank. In addition, for legacy trackmania modes it also contains BestTime, Score, NbrLapsFinished, LadderScore, and an array BestCheckpoints that contains the checkpoint times for the best race.
func (client *GbxClient) GetCurrentRankingForLogin(login string) ([]structs.TMPlayerRanking, error) {
	res, err := client.Call("GetCurrentRankingForLogin", login)
	if err != nil {
		return []structs.TMPlayerRanking{}, err
	}

	var playerList []structs.TMPlayerRanking
	err = convertToStruct(res, &playerList)
	if err != nil {
		return []structs.TMPlayerRanking{}, err
	}

	return playerList, nil
}

// Returns the current winning team for the race in progress. (-1: if not in team mode, or draw match)
func (client *GbxClient) GetCurrentWinnerTeam() (int, error) {
	res, err := client.Call("GetCurrentWinnerTeam")
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}