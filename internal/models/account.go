package models

type ClientTokenResponse struct {
	AccessToken    string `json:"access_token"`
	ExpiresIn      int64  `json:"expires_in"`
	ExpiresAt      string `json:"expires_at"`
	TokenType      string `json:"token_type"`
	ClientID       string `json:"client_id"`
	InternalClient bool   `json:"internal_client"`
	ClientService  string `json:"client_service"`
	ProductID      string `json:"product_id"`
	ApplicationID  string `json:"application_id"`
}

type TokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	ExpiresAt        string `json:"expires_at"`
	TokenType        string `json:"token_type"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpires   int64  `json:"refresh_expires"`
	RefreshExpiresAt string `json:"refresh_expires_at"`
	AccountID        string `json:"account_id"`
	ClientID         string `json:"client_id"`
	InternalClient   bool   `json:"internal_client"`
	ClientService    string `json:"client_service"`
	DisplayName      string `json:"displayName"`
	App              string `json:"app"`
	InAppID          string `json:"in_app_id"`
	ProductID        string `json:"product_id"`
	ApplicationID    string `json:"application_id"`
}

type AccountLookupResponse struct {
	ID                           string `json:"id"`
	DisplayName                  string `json:"displayName"`
	Name                         string `json:"name"`
	Email                        string `json:"email"`
	FailedLoginAttempts          int    `json:"failedLoginAttempts"`
	LastLogin                    string `json:"lastLogin"`
	NumberOfDisplayNameChanges   int    `json:"numberOfDisplayNameChanges"`
	AgeGroup                     string `json:"ageGroup"`
	Headless                     bool   `json:"headless"`
	Country                      string `json:"country"`
	LastName                     string `json:"lastName"`
	PhoneNumber                  string `json:"phoneNumber"`
	Company                      string `json:"company"`
	PreferredLanguage            string `json:"preferredLanguage"`
	LastDisplayNameChange        string `json:"lastDisplayNameChange"`
	CanUpdateDisplayName         bool   `json:"canUpdateDisplayName"`
	TfaEnabled                   bool   `json:"tfaEnabled"`
	EmailVerified                bool   `json:"emailVerified"`
	MinorVerified                bool   `json:"minorVerified"`
	MinorExpected                bool   `json:"minorExpected"`
	MinorStatus                  string `json:"minorStatus"`
	GuardianChallengeTimestamp   string `json:"guardianChallengeTimestamp"`
	SiweNotificationEnabled      bool   `json:"siweNotificationEnabled"`
	CabinedMode                  bool   `json:"cabinedMode"`
	HasHashedEmail               bool   `json:"hasHashedEmail"`
	LastReviewedSecuritySettings string `json:"lastReviewedSecuritySettings"`
}

type PublicAccountLookupResponse struct {
	ID            string              `json:"id"`
	DisplayName   string              `json:"displayName"`
	ExternalAuths map[string]struct{} `json:"externalAuths"`
}

type AccountExternalAuth struct {
	AccountID           string `json:"accountId"`
	Type                string `json:"type"`
	ExternalAuthID      string `json:"externalAuthId"`
	ExternalAuthIDType  string `json:"externalAuthIdType"`
	ExternalDisplayName string `json:"externalDisplayName"`
	AuthIds             []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"authIds"`
	DateAdded string `json:"dateAdded"`
}
