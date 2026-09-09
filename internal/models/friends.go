package models

type BlockListResponse struct {
	BlockedUsers []string `json:"blockedUsers"`
}

type RecentPlayersResponse struct {
	RecentPlayers []string `json:"recentplayers"` // i assume this is an arr of accountIds
}
