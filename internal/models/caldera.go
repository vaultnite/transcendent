package models

type CalderaAnticheatProviderRequest struct {
	AccountID    string `json:"account_id"`
	ExchangeCode string `json:"exchange_code"`
	TestMode     bool   `json:"test_mode"`
	EpicApp      string `json:"epic_app"`
	Nvidia       bool   `json:"nvidia"`
	Luna         bool   `json:"luna"`
	Salmon       bool   `json:"salmon"`
	GDKMode      bool   `json:"gdk_mode"`
}

type CalderaAnticheatProviderResponse struct {
	Provider     string `json:"provider"`
	JWT          string `json:"jwt"`
	DeploymentID string `json:"deployment_id"`
	SandboxID    string `json:"sandbox_id"`
}
