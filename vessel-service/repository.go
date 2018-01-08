package main

import (
	pb "github.com/ruprict/evalentine/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

// Repository is the interface we want
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
	Close()
}

// VesselRepository is our MongoDB
type VesselRepository struct {
	session *mgo.Session
}

// FindAvailable finds the first vessel
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}

	return vessel, err
}

// Create a new vessel
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// Close the session
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
