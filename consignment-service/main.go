package main

import (
	"log"

	"golang.org/x/net/context"

	micro "github.com/micro/go-micro"
	pb "github.com/ruprict/evalentine/consignment-service/proto/consignment"
)

const (
	port = ":50051"
)

// Repository is the interface for our service
type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// ConsignmentRepository is a dummy implementation for now
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

// Create is the repository method that stores consignment in
// an array in memory
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll is the interface method
func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service is, well, our service
type service struct {
	repo Repository
}

// CreateConsignment is the service method. Saves to the repo
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

// GetConsignments is the service method. gets all from repo
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Consignment = consignment
	res.Created = true
	return nil
}
func main() {
	repo := &ConsignmentRepository{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
