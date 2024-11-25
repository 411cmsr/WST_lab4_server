package sqldb_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=192.168.253.229 dbname=restapi_test user=postgres password=postgres sslmode=disable"
	}
	os.Exit(m.Run())
}
