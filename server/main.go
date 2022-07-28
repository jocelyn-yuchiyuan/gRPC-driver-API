package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"

	pb "driversapp.com/grpc/protos"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var drivers []*pb.DriverInfo

type driverServer struct {
	pb.UnimplementedDriverServer
}

func main() {
	initDrivers()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterDriverServer(s, &driverServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func initDrivers() {
	driver1 := &pb.DriverInfo{Id: "1", Name: "Jack",
		Tele: "0123456", Point: &pb.Point{Latitude: 1, Longitude: 2}}
	driver2 := &pb.DriverInfo{Id: "2", Name: "Eric",
		Tele:  "2345678",
		Point: &pb.Point{Latitude: 1, Longitude: 2}}

	drivers = append(drivers, driver1)
	drivers = append(drivers, driver2)
}

func (s *driverServer) GetDrivers(in *pb.Empty,
	stream pb.Driver_GetDriversServer) error {
	log.Printf("Received: %v", in)
	for _, driver := range drivers {
		if err := stream.Send(driver); err != nil {
			return err
		}
	}
	return nil
}

func (s *driverServer) GetDriver(ctx context.Context,
	in *pb.Id) (*pb.DriverInfo, error) {
	log.Printf("Received: %v", in)

	res := &pb.DriverInfo{}

	for _, driver := range drivers {
		if driver.GetId() == in.GetValue() {
			res = driver
			break
		}
	}

	return res, nil
}

func (s *driverServer) CreateDriver(ctx context.Context,
	in *pb.DriverInfo) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	res := pb.Id{}
	res.Value = strconv.Itoa(rand.Intn(100000000))
	in.Id = res.GetValue()
	drivers = append(drivers, in)
	return &res, nil
}

func (s *driverServer) UpdateDriver(ctx context.Context,
	in *pb.DriverInfo) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	res := pb.Status{}
	for index, driver := range drivers {
		if driver.GetId() == in.GetId() {
			drivers = append(drivers[:index], drivers[index+1:]...)
			in.Id = driver.GetId()
			drivers = append(drivers, in)
			res.Value = 1
			break
		}
	}

	return &res, nil
}

func (s *driverServer) DeleteDriver(ctx context.Context,
	in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	res := pb.Status{}
	for index, driver := range drivers {
		if driver.GetId() == in.GetValue() {
			drivers = append(drivers[:index], drivers[index+1:]...)
			res.Value = 1
			break
		}
	}

	return &res, nil
}
