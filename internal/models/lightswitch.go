package models

type LightswitchServiceStatus struct {
	ServiceInstanceID  string                                  `json:"serviceInstanceId"`
	Status             string                                  `json:"status"`
	Message            string                                  `json:"message"`
	MaintenanceURI     string                                  `json:"maintenanceUri"`
	OverrideCatalogIDs []string                                `json:"overrideCatalogIds"`
	AllowedActions     []string                                `json:"allowedActions"`
	Banned             bool                                    `json:"banned"`
	LauncherInfoDTO    LightswitchServiceStatusLauncherInfoDTO `json:"launcherInfoDTO"`
}

type LightswitchServiceStatusLauncherInfoDTO struct {
	AppName       string `json:"appName"`
	CatalogItemID string `json:"catalogItemId"`
	Namespace     string `json:"namespace"`
}
