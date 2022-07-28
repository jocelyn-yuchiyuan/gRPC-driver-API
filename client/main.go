package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "driversapp.com/grpc/protos"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewDriverClient(conn)

	// runGetMovies(client)
	// runGetMovie(client, "1")
	// runCreateMovie(client, "24325645", "Spiderman Spiderverse",
	// 	"Stan", "Lee")
	// runUpdateMovie(client, "98498081", "24325645", "Spiderman Spiderverse",
	// 	"Peter", "Parker")
	// runDeleteMovie(client, "98498081")
}

func runGetDrivers(client pb.DriverClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Empty{}
	stream, err := client.GetDrivers(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetDrivers(_) = _, %v", client, err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetDrivers(_) = _, %v", client, err)
		}
		log.Printf("DriverInfo: %v", row)
	}
}

func runGetDriver(client pb.DriverClient, driverid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: driverid}
	res, err := client.GetDriver(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetDriver(_) = _, %v", client, err)
	}
	log.Printf("DriverInfo: %v", res)
}

func runCreateDriver(client pb.DriverClient, driverid string, name string,
	tele string, latitude int32, longitude int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.DriverInfo{Id: driverid, Name: name, Tele: tele,
		Point: &pb.Point{Latitude: latitude, Longitude: longitude}}
	res, err := client.CreateDriver(ctx, req)
	if err != nil {
		log.Fatalf("%v.CreateDriver(_) = _, %v", client, err)
	}
	if res.GetValue() != "" {
		log.Printf("CreateDriver Id: %v", res)
	} else {
		log.Printf("CreateDriver Failed")
	}

}

func runUpdateDriver(client pb.DriverClient, driverid string, name string,
	tele string, latitude int32, longitude int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.DriverInfo{Id: driverid, Name: name, Tele: tele,
		Point: &pb.Point{Latitude: latitude, Longitude: longitude}}
	res, err := client.UpdateDriver(ctx, req)
	if err != nil {
		log.Fatalf("%v.UpdateDriver(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("UpdateDriver Success")
	} else {
		log.Printf("UpdateDriver Failed")
	}
}

func runDeleteMovie(client pb.DriverClient, driverid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: driverid}
	res, err := client.DeleteDriver(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteMovie(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("DeleteDriver Success")
	} else {
		log.Printf("Deletedriver Failed")
	}
}
