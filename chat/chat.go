package chat

import (
	"sync"
	"net"
	"log"
	"errors"
	// "fmt"
	"github.com/wusidn/qiaqia/dao"
)

type ChatService struct{
	db 				dao.Dao
	done			chan struct{}
}

func New(db dao.Dao)*ChatService{
	return &ChatService{
		db: db,
		done: make(chan struct{}),
	}
}

var clientManager *ClientManager = NewClientManager()
var userManager *UserManager = NewUserManager()

func getLocalHostAddress() (string, error){

	result := ""

	netInterfaces, err := net.Interfaces()
    if err != nil {
        return "", err
	}

	OuterLoop:
	for i := 0; i < len(netInterfaces); i++ {
        if (netInterfaces[i].Flags & net.FlagUp) != 0 {
            addrs, _ := netInterfaces[i].Addrs()
 
            for _, address := range addrs {
                if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
                    if ipnet.IP.To4() != nil {
						result = ipnet.IP.String()
                        break OuterLoop
                    }
                }
            }
        }
	}

	if len(result) <= 0{
		return "", errors.New("address not find")
	}

	return result, nil
}

func (service *ChatService)Run(wg *sync.WaitGroup){

	defer wg.Done()

	addr, err := getLocalHostAddress()
	if err != nil{
		log.Printf("getLocalHostAddress fail [%v]", err.Error())
		return
	}

	server, err := net.Listen("tcp", addr + ":9962")

	if err != nil{
		log.Printf("tcp listen fail [%v]", err.Error())
		return
	}

	defer server.Close()

	log.Printf("tcp server listen " + addr +  ":9962")


	go clientManager.Start()
	go userManager.Start()

	go func(){
		for{
			client, err := server.Accept()
			if err != nil{
				log.Printf("accept fail %v", err.Error())
				break
			}
	
			clientManager.Register <- &Client{
				ID: "",
				Conn: client,
			}
		}
	}()

	MainLoop:
	for{
		select{
		case <- service.done:
			break MainLoop
		}
	}
}


func (service *ChatService)Done(){
	service.done <- struct{}{}
}
