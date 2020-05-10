package chat


type UserManager struct{
	Users	map[string]*User
	//ClientId -> UserId
	Clients map[string]string
	Send	chan []byte
}

type User struct{
	ID 		string
	Clients	[]string
}


func NewUserManager() *UserManager{
	return &UserManager{
		Users: make(map[string]*User),
		Clients: make(map[string]string),
		Send: make(chan []byte, 100),
	}
}

func (manager *UserManager)Start(){
	
}