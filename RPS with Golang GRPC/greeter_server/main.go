/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"math/rand"
	"strconv"
	"fmt"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	clientHand, _ := strconv.Atoi(in.Dice)

	rand.Seed(time.Now().Unix())
	serverHand := rand.Intn(3-1) + 1

	clientHandType := ""
	serverHandType := ""

	if(clientHand == 1){
		clientHandType = "Batu"
	}else if (clientHand == 2){
		clientHandType = "Gunting"
	}else{
		clientHandType = "Kertas"
	}

	if(serverHand == 1){
		serverHandType = "Batu"
	}else if(serverHand == 2){
		serverHandType = "Gunting"
	}else{
		serverHandType = "Kertas"
	}


	message := "| Client Use : " + clientHandType + " | Server Use : " + serverHandType + " | "


	if(clientHand == 1 && serverHand == 1){
		fmt.Println("Game Info : ", message, "Draw |")
		return &pb.HelloReply{Message: message + " Draw |"}, nil
	}else if(clientHand == 1 && serverHand == 2){
		fmt.Println("Game Info : ", message, "You Lose |")
		return &pb.HelloReply{Message: message + " You Win |"}, nil
	}else if(clientHand == 1 && serverHand == 3){
		fmt.Println("Game Info : ", message, "You Win |")
		return &pb.HelloReply{Message: message + " You Lose |"}, nil
	}else if(clientHand == 2 && serverHand == 1){
		fmt.Println("Game Info : ", message, "You Win |")
		return &pb.HelloReply{Message: message + " You Lose |"}, nil
	}else if(clientHand == 2 && serverHand == 2){
		fmt.Println("Game Info : ", message, "Draw |")
		return &pb.HelloReply{Message: message + " Draw |"}, nil
	}else if(clientHand == 2 && serverHand == 3){
		fmt.Println("Game Info : ", message, "You Lose |")
		return &pb.HelloReply{Message: message + " You Win |"}, nil
	}else if(clientHand == 3 && serverHand == 1){
		fmt.Println("Game Info : ", message, "You Lose |")
		return &pb.HelloReply{Message: message + " You Win |"}, nil
	}else if(clientHand == 3 && serverHand == 2){
		fmt.Println("Game Info : ", message, "You Win |")
		return &pb.HelloReply{Message: message + " You Lose |"}, nil
	}else if(clientHand == 3 && serverHand == 3){
		fmt.Println("Game Info : ", message, "Draw |")
		return &pb.HelloReply{Message: message + " Draw |"}, nil
	} else {
		return &pb.HelloReply{Message: "Hehe"}, nil
	}

	return &pb.HelloReply{Message: message + "Hehe2"}, nil
}








func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
