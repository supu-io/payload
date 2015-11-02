package payload

type Payload struct {
	Config     *Config     `json:"config"`
	Issue      *Issue      `json:"issue"`
	Transition *Transition `json:"transition"`
	Status     *[]string   `json:"status"`
}

type Transition struct {
	From *string `json:"from"`
	To   *string `json:"to"`
}

type Config struct {
	Github *Github `json:"github"`
}

type Github struct {
	Token *string `json:"token"`
}

type Issue struct {
	ID *string `json:"id"`
}
