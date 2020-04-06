package haha

import "log"

type Haha struct{
	Name string
}

func (a *Haha) SayHi() {
	log.Printf("---> hello my name is %s", a.Name)
}