package stocks

//DTO represents the stock entity, but with no internal ID
type DTO struct {
	Code        string `json:"code"`
	FantasyName string `json:"fantasyName"`
}
