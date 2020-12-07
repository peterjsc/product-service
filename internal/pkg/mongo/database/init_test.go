package database

import (
	"context"
	"os"
	"testing"
)

func setEnv() {
	os.Setenv("DATABASE_ENDPOINT", "mongodb://localhost:27017")
	os.Setenv("DATABASE_NAME", "news-feed")
	os.Setenv("DATABASE_USER", "root")
	os.Setenv("DATABASE_PASS", "example")
}

func unsetEnv() {
	os.Unsetenv("DATABASE_ENDPOINT")
	os.Unsetenv("DATABASE_NAME")
	os.Unsetenv("DATABASE_USER")
	os.Unsetenv("DATABASE_PASS")
}

func initDBFramework() (DBFramework, error) {
	ctx := context.Background()

	setEnv()
	df, err := New(ctx)

	return df, err
}

func TestDBFramework(t *testing.T) {
	df, err := initDBFramework()

	if err != nil {
		t.Errorf("Expected nil but got %v", err.Error())
	}
	if (DBFramework{}) == df {
		t.Errorf("Failed to create DB")
	}
	unsetEnv()
}
