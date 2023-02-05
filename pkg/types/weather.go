package types

type Weather struct {
	Temperature string `json:"temperature"`
	Conditions  string `json:"conditions"`
	Alerts      string `json:"alerts"`
}
