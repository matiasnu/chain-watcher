package models

type ContractWatch struct {
	Contract string   `json:"contract"`
	Topics   []string `json:"topics"`
}
