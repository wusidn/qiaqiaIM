package main


import (
	"net"
	"log"
)

func main(){

	client, err := net.Dial("tcp", "127.0.0.1:9962")

	if err != nil{
		log.Printf("client connent fail [%v]", err.Error())
		return
	}

	defer client.Close()

	_, err = client.Write([]byte("hello world!"))

	if err != nil{
		log.Printf("clinet write fail [%v]", err.Error())
		return
	}

}