package main

import (
	pb "github.com/ruprict/evalentine/consignment-service/proto/consignment"
	mgo "gopkg.in/mgo.v2"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

// Repository is the interface we want
type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

// ConsignmentRepository is our MongoDB
type ConsignmentRepository struct {
	session *mgo.Session
}

// Create a new consignment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) {
	return repo.collection().Insert(consignment)
}

// GetAll consignments
func (repo *ConsignmentRepository) GetAll() ([]*pb.Collection, error) {
	var consignments []*pb.Consignments

	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}

// Close the session
func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}

func (repo *ConsignmentRepository) colleciton() *mgo.Collection {
	return repo.session.DB(dbName).C(consignmentCollection)
}
