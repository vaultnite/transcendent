package models

import "encoding/json"

type APIVersionModule struct {
    CLN       string `json:"cln"`
    Build     string `json:"build"`
    BuildDate string `json:"buildDate"`
    Version   string `json:"version"`
    Branch    string `json:"branch"`
}

type APIVersionResponse struct {
    App                       string                      `json:"app"`
    ServerDate                string                      `json:"serverDate"`
    OverridePropertiesVersion string                      `json:"overridePropertiesVersion"`
    CLN                       string                      `json:"cln"`
    Build                     string                      `json:"build"`
    ModuleName                string                      `json:"moduleName"`
    BuildDate                 string                      `json:"buildDate"`
    Version                   string                      `json:"version"`
    Branch                    string                      `json:"branch"`
    Modules                   map[string]APIVersionModule `json:"modules"`
}

type VersionCheckResponseType string

const (
    VersionCheckResponseTypeAppRedirect VersionCheckResponseType = "APP_REDIRECT" // needs appRedirect to be set
    VersionCheckResponseTypeHardUpdate  VersionCheckResponseType = "HARD_UPDATE"
    VersionCheckResponseTypeNotEnabled  VersionCheckResponseType = "NOT_ENABLED"
    VersionCheckResponseTypeNoUpdate    VersionCheckResponseType = "NO_UPDATE"
    VersionCheckResponseTypeSoftUpdate  VersionCheckResponseType = "SOFT_UPDATE"
)

type VersionCheckResponse struct {
    Type        VersionCheckResponseType `json:"type"`
    AppRedirect string                   `json:"appRedirect,omitempty"`
}

type Receipt struct {
    AppStore    string `json:"appStore"`
    AppStoreID  string `json:"appStoreId"`
    ReceiptID   string `json:"receiptId"`
    ReceiptInfo string `json:"receiptInfo"`
}

type RankedTeamInfo struct {
    Members         []RankedTeamMemberInfo `json:"members"`
    SocialPartySize int32                  `json:"socialPartySize"`
}

type RankedTeamMemberInfo struct {
    AccountID     string `json:"accountId"`
    PartyLeaderID string `json:"partyLeaderId"`
}

type GetAccountMMRResponse struct {
    Rating         int32 `json:"rating"`
    NumGamesPlayed int32 `json:"numGamesPlayed"`
}

type GetTeamEloResponse struct {
    Rating int32 `json:"rating"`
}

type ReportWaitTimesPayload struct {
    BucketID            string  `json:"bucketId"`
    SessionID           string  `json:"sessionId"`
    PlayerMatchWaitSecs float64 `json:"playerMatchWaitSecs"`
    TeamMatchWaitSecs   float64 `json:"teamatchWaitSecs"`
}

type WaitTimeEstimate struct {
    AverageWaitTimeSecs float64 `json:"averageWaitTimeSecs"`
    NumSamples          int     `json:"numSamples"`
}

type WaitTimeEstimateOrion struct {
    WaitTimeEstimate
    /* these appear to be part of the response orion expects from reading the string dump, ida is being annoying rn so cant rlly be sure */
    BucketID string `json:"bucketId"`
    HeroID   string `json:"heroId"`
}

type WaitTimeEstimateUT struct {
    WaitTimeEstimate
    RatingType string `json:"ratingType"` // appears to be only present in the ut '/estimate' subpath, and not expected by orion
}

// MatchmakingRequestResponseWIP yoinked from UT4MasterServer, will need to compare with fn and try to get a universal structure sorted.
type MatchmakingRequestResponseWIP struct {
    ID                              string                                  `json:"id"`
    OwnerID                         string                                  `json:"ownerId"`
    OwnerName                       string                                  `json:"ownerName"`
    ServerName                      string                                  `json:"serverName"`
    ServerAddress                   string                                  `json:"serverAddress"`
    ServerPort                      int                                     `json:"serverPort"`
    MaxPublicPlayers                int                                     `json:"maxPublicPlayers"`
    OpenPublicPlayers               int                                     `json:"openPublicPlayers"`
    MaxPrivatePlayers               int                                     `json:"maxPrivatePlayers"`
    OpenPrivatePlayers              int                                     `json:"openPrivatePlayers"`
    Attributes                      MatchmakingRequestResponseWIPAttributes `json:"attributes"`
    PublicPlayers                   []interface{}                           `json:"publicPlayers"`
    PrivatePlayers                  []interface{}                           `json:"privatePlayers"`
    TotalPlayers                    int                                     `json:"totalPlayers"`
    AllowJoinInProgress             bool                                    `json:"allowJoinInProgress"`
    ShouldAdvertise                 bool                                    `json:"shouldAdvertise"`
    IsDedicated                     bool                                    `json:"isDedicated"`
    UsesStats                       bool                                    `json:"usesStats"`
    AllowInvites                    bool                                    `json:"allowInvites"`
    UsesPresence                    bool                                    `json:"usesPresence"`
    AllowJoinViaPresence            bool                                    `json:"allowJoinViaPresence"`
    AllowJoinViaPresenceFriendsOnly bool                                    `json:"allowJoinViaPresenceFriendsOnly"`
    BuildUniqueID                   string                                  `json:"buildUniqueId"`
    LastUpdated                     string                                  `json:"lastUpdated"`
    Started                         bool                                    `json:"started"`
}

type MatchmakingRequestResponseWIPAttributes = json.RawMessage

//    Attributes         struct {
//        UTSERVERNAMES         string `json:"UT_SERVERNAME_s"`
//        UTREDTEAMSIZEI        int    `json:"UT_REDTEAMSIZE_i"`
//        UTNUMMATCHESI         int    `json:"UT_NUMMATCHES_i"`
//        UTGAMEINSTANCEI       int    `json:"UT_GAMEINSTANCE_i"`
//        UTMAXSPECTATORSI      int    `json:"UT_MAXSPECTATORS_i"`
//        BEACONPORTI           int    `json:"BEACONPORT_i"`
//        UTPLAYERONLINEI       int    `json:"UT_PLAYERONLINE_i"`
//        UTSERVERVERSIONS      string `json:"UT_SERVERVERSION_s"`
//        GAMEMODES             string `json:"GAMEMODE_s"`
//        UTHUBGUIDS            string `json:"UT_HUBGUID_s"`
//        UTBLUETEAMSIZEI       int    `json:"UT_BLUETEAMSIZE_i"`
//        UTMATCHSTATES         string `json:"UT_MATCHSTATE_s"`
//        UTSERVERTRUSTLEVELI   int    `json:"UT_SERVERTRUSTLEVEL_i"`
//        UTSERVERINSTANCEGUIDS string `json:"UT_SERVERINSTANCEGUID_s"`
//        UTTRAININGGROUNDB     bool   `json:"UT_TRAININGGROUND_b"`
//        UTMINELOI             int    `json:"UT_MINELO_i"`
//        UTMAXELOI             int    `json:"UT_MAXELO_i"`
//        UTSPECTATORSONLINEI   int    `json:"UT_SPECTATORSONLINE_i"`
//        UTMAXPLAYERSI         int    `json:"UT_MAXPLAYERS_i"`
//        UTSERVERMOTDS         string `json:"UT_SERVERMOTD_s"`
//        MAPNAMES              string `json:"MAPNAME_s"`
//        UTMATCHDURATIONI      int    `json:"UT_MATCHDURATION_i"`
//        UTSERVERFLAGSI        int    `json:"UT_SERVERFLAGS_i"`
//    } `json:"attributes"`
