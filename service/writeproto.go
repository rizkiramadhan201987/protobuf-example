package service

import (
	"log"
	"os"

	"github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/userservice/user"
	"google.golang.org/protobuf/proto"
)

// Write proto to biner
func WriteProto() {
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
	bytes, err := proto.Marshal(&u)

	if err != nil {
		return
	}
	err = os.WriteFile("test", bytes, 0644)
	if err != nil {
		return
	}

}
func ReadProto() {
	u := user.User{}
	bytes, err := os.ReadFile("test")
	if err != nil {
		return
	}
	proto.Unmarshal(bytes, &u)
	log.Println(&u)
}
