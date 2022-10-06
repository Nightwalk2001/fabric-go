package docs

type Arrears struct {
	Client   string `json:"client" bson:"_id"`
	Amount   int    `json:"amount"`
	Money    int    `json:"money"`
	Received int    `json:"received"`
}
