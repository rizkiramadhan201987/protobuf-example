package main

import (
	"fmt"
	"log"
	"time"

	"github.com/rizkiramadhan201987/learn-grpc/service"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05" + " " + string(bytes)))
}
func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// service.UserProtoToJson()
	// service.BasicUnmarshall()
	// service.BasicOneOf()
	// service.BasicMap()
	// service.WriteProto()
	// service.WriteProtoToJsonBytes()
	// service.ReadJsonBytesToProto()
	// service.ReadProto()
	// service.BasicHello()
	// service.SayHelloUser()
	service.UserProtoToJson()
	// service.UserJsonToProto()
	// service.AddUserToUserGroup()
	// service.JobSearch()
	// service.JobSearchSoftware()
}
