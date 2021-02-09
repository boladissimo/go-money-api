package stock

//Entity is the entity of the stock domain
type Entity struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	FantasyName string `json:"fantasyName"`
}
