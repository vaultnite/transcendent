package models

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

type GetTeamEloResponse struct {
    Rating int32 `json:"rating"`
}

type WaitTimeEstimate struct {
    /* these appear to be part of the response orion expects from reading the string dump, ida is being annoying rn so cant rlly be sure */
    BucketID string `json:"bucketId,omitempty"`
    HerotID  string `json:"heroId,omitempty"`

    RatingType string `json:"ratingType,omitempty"` // appears to be only present in the ut '/estimate' subpath, and not expected by orion

    AverageWaitTimeSecs float64 `json:"averageWaitTimeSecs"`
    NumSamples          int     `json:"numSamples"`
}
