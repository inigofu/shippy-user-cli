package main

import (
	"log"
	"os"

	pb "github.com/inigofu/shippy-user-service/proto/auth"
	micro "github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

func main() {

	srv := micro.NewService(

		micro.Name("go.micro.srv.user-cli"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	client := pb.NewAuthClient("go.micro.srv.user", microclient.DefaultClient)

	r, err := client.CreateMenu(context.TODO(), &pb.Menu{
		Id:   "1",
		Name: "inicio",
		Url:  "/dashboard",
	})
	r1, err := client.CreateMenu(context.TODO(), &pb.Menu{
		Id:   "2",
		Name: "inicio",
		Url:  "/dashboard",
	})
	r2, err := client.CreateRole(context.TODO(), &pb.Role{
		Id:     "2",
		Name:   "admin",
		Menues: []*pb.Menu{r.Menu, r1.Menu},
	})
	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "test123"
	company := "BBC"
	r3, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
		Roles:    []*pb.Role{r2.Role},
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r3.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	// let's just exit because
	os.Exit(0)
}
