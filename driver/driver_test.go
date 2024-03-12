package driver_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/rferrazz/sqinn-go/driver"
	"github.com/stretchr/testify/require"
)

const test_name = "pippo"

func TestDriver(t *testing.T) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("?sqinnpath=%s", os.Getenv("SQINN_PATH")))
	require.Nil(t, err, "error should be nil")
	defer db.Close()

	_, err = db.Exec("CREATE TABLE users (name VARCHAR)")
	require.Nil(t, err, "error should be nil")

	var name string
	row := db.QueryRow("SELECT name FROM users WHERE name = ?", test_name)
	if err := row.Scan(&name); err != nil {
		require.Equal(t, err, sql.ErrNoRows)
	}

	_, err = db.Exec("insert into users(name) values(?)", test_name)
	require.Nil(t, err, "error should be nil")

	row = db.QueryRow("SELECT name FROM users WHERE name = ?", test_name)
	err = row.Scan(&name)
	require.Nil(t, err, "error should be nil")
	require.Equal(t, name, test_name)
}
