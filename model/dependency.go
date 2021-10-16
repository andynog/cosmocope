package model

// Dependency
type Dependency struct {
	CosmosSDK   string    `json:"cosmos_sdk"`
	Tendermint  string    `json:"tendermint"`
	IBC         string    `json:"ibc"`
}
