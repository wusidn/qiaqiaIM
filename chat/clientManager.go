package chat


import(
	"io"
	"log"
	"net"
)

type ClientManager struct{
	Clients		map[*Client]struct{}
	ClientIds	map[string]*Client

	Register	chan *Client
	Unregister	chan *Client
}

type Client struct{
	ID		string
	Conn 	net.Conn
	Send	chan []byte
}

type Message struct{
	Sender		string		`json:"sender,omitempty"`
	Recipient	string		`json:"recipient,omitempty"`
	Content		string		`json:"content,omitempty"`
}

func NewClientManager()*ClientManager{
	return &ClientManager{
		Clients: 	make(map[*Client]struct{}),
		ClientIds: 	make(map[string]*Client),
		Register: 	make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (manager *ClientManager)Start(){

	for{
		select{
		case client := <-manager.Register:
			manager.Clients[client] = struct{}{}
			manager.ClientIds[client.ID] = client

			//init client.Send channel
			if client.Send == nil{
				client.Send = make(chan []byte)
			}

			go client.Read()
			go client.Write()
		case client := <-manager.Unregister:
			close(client.Send)
			delete(manager.Clients, client)
			delete(manager.ClientIds, client.ID)
			log.Printf("Unregister socket client")
		}
	}
}

func (client *Client)Read(){
	defer func(){
		clientManager.Unregister <- client
		client.Conn.Close()
		log.Printf("read close")
	}()

	buf := make([]byte, 2048)

	for{
		n, err := client.Conn.Read(buf)

		if err == io.EOF{
			break
		}

		if err != nil{
			log.Printf("clinet read fail [%v]", err.Error())
			break
		}

		// 
		log.Printf("clinet read: %s", string(buf[:n]))
	}
}

func (client *Client)Write(){

	defer func(){
		client.Conn.Close()
		log.Printf("write close")
	}()

	MainLoop:
	for{
		select{
		case message, ok:= <-client.Send:
			if !ok {
				break MainLoop
			}
			client.Conn.Write(message)
		}
	}
}