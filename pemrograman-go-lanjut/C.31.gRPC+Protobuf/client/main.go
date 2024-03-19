package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc/credentials/insecure"

	"C.31.gRPC+Protobuf/common/config"
	"C.31.gRPC+Protobuf/model"
	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.ServiceGaragePort
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.ServiceUserPort
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "N-001",
		Name:     "Febrianty Elizabeth Sailolin",
		Password: "tes123",
		Gender:   model.UserGender(model.UserGender_value["FEMALE"]),
	}

	user2 := model.User{
		Id:       "N-002",
		Name:     "Dede Syarifudin Hidayat",
		Password: "tes123",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	garage1 := model.Garage{
		Id:   "Q-001",
		Name: "Febrianty Elizabeth Sailolin",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	garage2 := model.Garage{
		Id:   "Q-002",
		Name: "Dede Syarifudin Hidayat",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	// User ======================================================================================
	user := serviceUser()

	fmt.Printf("\n %s> user test\n", strings.Repeat("=", 10))

	// register user1
	user.Register(context.Background(), &user1)

	// register user2
	user.Register(context.Background(), &user2)

	// Garage ======================================================================================
	garage := serviceGarage()

	fmt.Printf("\n %s> garage test A\n", strings.Repeat("=", 10))

	// add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	// add garage2 to user2
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user2.Id,
		Garage: &garage2,
	})

	// show all garages of user1
	res2, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))

	fmt.Printf("\n %s> garage test B\n", strings.Repeat("=", 10))

	// add garage3 to user2
	// garage.Add(context.Background(), &model.GarageAndUserId{
	// 	UserId: user2.Id,
	// 	Garage: &garage2,
	// })

	// // show all garages of user2
	// res3, err := garage.List(context.Background(), &model.GarageUserId{UserId: user2.Id})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// res3String, _ := json.Marshal(res3.List)
	// log.Println(string(res3String))
}
