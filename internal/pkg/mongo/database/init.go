package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBFramework struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func New(ctx context.Context) (DBFramework, error) {

	dbEndpoint := os.Getenv("DATABASE_ENDPOINT")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")

	cred := options.Credential{
		Username: dbUser,
		Password: dbPass,
	}

	connStr := addCredentialsToURI(dbEndpoint, dbUser, dbPass)
	if connStr == "" {
		log.Fatalf("Failed to add database credentials to URI %s", dbEndpoint)
		return DBFramework{}, errors.New("Failed to add database credentials")
	}

	clientOptions := options.Client().ApplyURI(connStr).SetAuth(cred)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to database %s", err.Error())
		return DBFramework{}, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed with checking database connection %s", err.Error())
		return DBFramework{}, err
	}

	database := client.Database(dbName)

	dbf := DBFramework{
		Client:   client,
		Database: database,
	}

	return dbf, nil
}

func addCredentialsToURI(endpoint, user, pass string) string {
	uriReg, _ := regexp.Compile(`^(mongodb(?:\+srv)?):\/\/(.*)$`)
	matches := uriReg.FindStringSubmatch(endpoint)

	if matches == nil {
		return ""
	}
	if user == "" || pass == "" {
		return endpoint
	}
	connStr := fmt.Sprintf("%s://%s:%s@%s", matches[1], user, pass, matches[2])
	return connStr
}
