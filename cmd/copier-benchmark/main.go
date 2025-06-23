package main

import (
	"log"
	"runtime"
	"time"

	"github.com/jinzhu/copier"
)

type ModelIn struct {
	Name      string
	Email     string
	FirstName string
	LastName  string
	Username  string
	Password  string
	Iden      int `copier:"Identity"`
	Extra     Embedded
}

func transfrom(s ModelIn) Model {
	return Model{
		s.Name,
		s.Email,
		s.FirstName,
		s.LastName,
		s.Username,
		s.Password,
		s.Extra,
		s.Iden,
	}
}

type Model struct {
	Name      string
	Email     string
	FirstName string
	LastName  string
	Username  string
	Password  string
	Extra     Embedded
	Identity  int
}

type Embedded struct {
	Name      string
	Email     string
	FirstName string
	LastName  string
	Username  string
	Password  string
}

func main() {
	log.Println("loop 1 result")
	benchmark(1)

	log.Println("loop 10 result")
	benchmark(10)

	log.Println("loop 100 result")
	benchmark(100)

	log.Println("loop 1000 result")
	benchmark(1000)

	log.Println("loop 10000 result")
	benchmark(10000)

	log.Println("loop 100000 result")
	benchmark(100000)
}

func benchmark(cycle int) {
	mi := ModelIn{
		Name:      "Fucker",
		Email:     "natholdallas@gmail.com",
		FirstName: "Nathol",
		LastName:  "Dallas",
		Username:  "natholdallas",
		Password:  "1234567890",
		Extra: Embedded{
			Name:      "Fucker",
			Email:     "natholdallas@gmail.com",
			FirstName: "Nathol",
			LastName:  "Dallas",
			Username:  "natholdallas",
			Password:  "1234567890",
		},
		Iden: 1234,
	}

	// test origin transfrom
	now := time.Now()
	var res Model
	for range cycle {
		res = transfrom(mi)
	}
	log.Println("origin transfrom duration: ", time.Since(now))

	// test copier transfrom
	now = time.Now()
	var res1 Model
	for range cycle {
		copier.Copy(&res1, mi)
	}
	log.Println("copier transfrom duration: ", time.Since(now))

	// test copier transfrom
	now = time.Now()
	var res2 Model
	for range cycle {
		copier.CopyWithOption(&res2, mi, copier.Option{DeepCopy: true})
	}
	log.Println("copier deepcopy transfrom duration: ", time.Since(now))

	runtime.KeepAlive(res)
	runtime.KeepAlive(res1)
	runtime.KeepAlive(res2)
}
