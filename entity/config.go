package entity

type Config struct {
	Credential Credential `json:"credential"`
	Aliases []Alias `json:"aliases"`
}
