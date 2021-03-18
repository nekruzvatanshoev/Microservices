package handlers

//type CreateItemToStoreHandler struct {
//	eventEmitter msgqueue.EventEmitter
//}
//
//
//func (h *CreateItemToStoreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application")
//	//var db domain.User
//	//db := []domain.User{
//	//	{
//	//		ID: "AAABBBCCC",
//	//		Email: "first.last@company.com",
//	//		Password: "thelonestarstate",
//	//		Admin: true,
//	//	},
//	//	{
//	//		ID: "BBBCCCDDD",
//	//		Email: "first.last@outlook.com",
//	//		Password: "rockymountainhigh",
//	//		Admin: false,
//	//	},
//	//}
//
//	var user domain.User
//
//	err := json.NewDecoder(r.Body).Decode(&user)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	//encodedUser, err := sarama.ByteEncoder.Encode(marshalUser)
//
//
//	msg := contracts.UserLoggedInEvent{ID: user.ID}
//	h.eventEmitter.Emit(&msg, "user")
//}

