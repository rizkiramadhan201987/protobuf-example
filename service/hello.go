package service

import (
	"log"

	"github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/helloservice/hello"
)

func BasicHello() {
	h := hello.Hello{
		Name: "Rizki Susilo Ramadhan",
	}

	log.Println(&h)
}
