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
