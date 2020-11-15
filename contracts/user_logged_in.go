package contracts


type UserLoggedInEvent struct {
	ID string `json:"id"`
}


func (c *UserLoggedInEvent) EventName() string {
	return "userLoggedIn"
}