package utils

import (
	"crud/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var Settings config.Settings
var MongoClient *mongo.Client
