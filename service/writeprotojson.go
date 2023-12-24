package service

import (
	"log"
	"os"

	"github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/userservice/user"
	"google.golang.org/protobuf/encoding/protojson"
)

func WriteProtoToJsonBytes() {
	u := user.User{
		Id:       1,
		Username: "rizkiramadhan201987",
		IsActive: true,
		Password: []byte("rizkikece"),
		Gender:   user.Gender_GENDER_MALE,
		Emails:   []string{"rizkysr90@gmail.com", "rizki.ramadhan@bfi.co.id"},
		Address: &user.Address{
			Street:  "Jalan-jalan santai",
			City:    "Bekasi",
			Country: "Indonesia",
		},
		Skills: map[string]uint32{"editing": 100},
	}
	jsonBytes, err := protojson.Marshal(&u)
	if err != nil {
		log.Fatalln("failed to marshal proto to json")
	}
	if err := os.WriteFile("protoJSON.bin", jsonBytes, 0644); err != nil {
		log.Fatalln("Can not write to file", err)
	}
}
func ReadJsonBytesToProto() {
	u := user.User{}
	in, err := os.ReadFile("protoJSON.bin")
	if err != nil {
		log.Fatalln("can not read file", err)
	}

	protojson.Unmarshal(in, &u)
	log.Println(&u)
}
