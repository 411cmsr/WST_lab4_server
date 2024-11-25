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
		databaseURL = "host=127.0.0.1 dbname=restapi_test user=pguser password=pgpassword sslmode=disable"
	}
	os.Exit(m.Run())
}
