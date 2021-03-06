package entities

type (
	CachedApiKey struct {
		UserID            int64  `json:"user_id"`
		AccountSecretHash string `json:"secret"`
		DlrURL            string `json:"dlr_url"`
	}
)
