package service

import (
	"log"

	"github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/usergroupservice/usergroup"
	"github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/userservice/user"
	"google.golang.org/protobuf/encoding/protojson"
)

//	func () {
//		finance.GroupId = 1
//		finance.Description = "Group Finance"
//		finance.Roles = []string{"audit", "report", "account"}
//	}
func AddUserToUserGroup() {
	// create group
	g1 := usergroup.UserGroup{
		GroupId:     1,
		GroupName:   "Hindia",
		Description: "Finance Group",
		Roles:       []string{"audit", "report", "account"},
	}
	// create user
	u1 := user.User{
		Id:       1,
		Username: "rizkysr90",
		IsActive: true,
		Password: []byte("adarizki123"),
		Gender:   user.Gender_GENDER_MALE,
		Emails:   []string{"rizki@gmail.com", "rizkibfi@bfi.co.id"},
		Address:  &user.Address{Street: "Jalan", City: "Bekasi", Country: "Indonesia"},
	}
	u2 := user.User{
		Id:       1,
		Username: "rizkysr901",
		IsActive: true,
		Password: []byte("adarizki123"),
		Gender:   user.Gender_GENDER_MALE,
		Emails:   []string{"rizki@gmail.com", "rizkibfi@bfi.co.id"},
		Address:  &user.Address{Street: "Jalan", City: "Bekasi", Country: "Indonesia"},
	}
	g1.Users = []*user.User{&u1, &u2}
	log.Println("sucesss insert user to user group")
	log.Println(&g1)
	log.Println("proto usergroup to json")
	jsonBytes, _ := protojson.Marshal(&g1)
	log.Println(string(jsonBytes))

}
