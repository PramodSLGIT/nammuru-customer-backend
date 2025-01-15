package database

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDb struct {
	c            *mongo.Client
	db           *mongo.Database
	rt           time.Duration
	wg           sync.WaitGroup
	shutDownflag int32
}

var (
	_initMongoDBCtx sync.Once
	mdb             *MongoDb
)

