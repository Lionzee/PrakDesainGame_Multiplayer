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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"time"
	"bufio"
	"os"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50052"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Apa yang ingin anda gunakan ?")
	fmt.Println("1. Batu ")
	fmt.Println("2. Gunting ")
	fmt.Println("3. Kertas ")

	handInString, _ := reader.ReadString('\n');
	replaceHandInString := strings.Replace(handInString, "\n", "", -1)

	// handInString := strconv.Itoa(clientHand)


	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	// request function yang mau di run di server disini
	r, err := c.SayHello(ctx, &pb.HelloRequest{Dice: replaceHandInString})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Game Info : %s", r.Message)
}
