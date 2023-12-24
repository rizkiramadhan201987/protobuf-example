package service

import (
	"log"
	"math/rand"
	"time"

	"github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/userservice/user"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func SayHelloUser() {
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
	}
	log.Println(&u)
}

func UserProtoToJson() {
	comm := randomCommunicationChannel()
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
			Coordinate: &user.Address_Coordinate{
				Lat: 99999,
				Lng: 11111,
			},
		},
		CommunicationChannel: comm,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func UserJsonToProto() {
	var protoFromJson user.User
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
	}
	jsonBytes, _ := protojson.Marshal(&u)
	protojson.Unmarshal(jsonBytes, &protoFromJson)
	log.Println(&protoFromJson)
}

func BasicUnmarshall() {
	sm := user.SocialMedia{SocialMediaPlatform: "Instagram", SocialMediaUsername: "rizkysr90"}
	var a anypb.Any
	anypb.MarshalFrom(&a, &sm, proto.MarshalOptions{})

	newSm := user.SocialMedia{}

	if err := a.UnmarshalTo(&newSm); err != nil {
		return
	}
	jsonBytes, _ := protojson.Marshal(&newSm)
	log.Println(string(jsonBytes))

}
func BasicMap() {
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
		Skills: map[string]uint32{"editing": 100, "cooking": 120},
	}
	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}
func BasicOneOf() {
	sm := user.SocialMedia{SocialMediaPlatform: "instagram", SocialMediaUsername: "rizkysr90"}
	instanceSM := user.User_SocialMedia{SocialMedia: &sm}

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
		ElectronicCommChannel: &instanceSM,
	}
	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}
func randomCommunicationChannel() *anypb.Any {
	rand.Seed(time.Now().UnixNano())

	pm := user.PaperMail{PaperMailAddress: "STRING"}
	sm := user.SocialMedia{SocialMediaPlatform: "Instagram", SocialMediaUsername: "rizkysr90"}
	ins := user.InstantMessaging{InstantMessagingProduct: "Skype", InstantMessagingUsername: "rizkysr90"}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &pm, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &sm, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &ins, proto.MarshalOptions{})
	}
	return &a

}
