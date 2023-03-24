package user

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	pkgErr "github.com/pkg/errors"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	DB_HOST     = "database"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_DATABASE = "backend"
	DB_PORT     = "5432"
	dsn         = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDB *sql.DB
var testRepo Repository

func TestMain(m *testing.M) {
	// connect to docker; fail if docker not running
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker; is it running? %s", err)
	}

	pool = p

	// set up our docker options, specifying the image and so forth
	opts := dockertest.RunOptions{
		Repository: "postgres",
		Env: []string{
			"POSTGRES_USER=" + DB_USER,
			"POSTGRES_PASSWORD=" + DB_PASSWORD,
			"POSTGRES_DB=" + DB_DATABASE,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: DB_PORT},
			},
		},
	}

	// get a resource (docker image)
	resource, err = pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("could not start resource: %s", err)
		_ = pool.Purge(resource)
	}

	// start the image and wait until it's ready
	if err := pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_DATABASE))
		if err != nil {
			log.Fatal(pkgErr.WithStack(err))
			return err
		}
		return testDB.Ping()
	}); err != nil {
		log.Fatalf("could not connect to database: %s", err)
		_ = pool.Purge(resource)
	}

	// populate the database with empty tables
	err = createTables()
	if err != nil {
		log.Fatalf("error creating tables: %s", err)
	}

	testRepo = &UserRepository{db: testDB}

	// run tests
	code := m.Run()

	// clean up
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge resource: %s", err)
	}

	os.Exit(code)
}

func createTables() error {
	tableSQL, err := os.ReadFile("./testdb/init.sql")
	if err != nil {
		log.Fatal(pkgErr.WithStack(err))
		return err
	}

	_, err = testDB.Exec(string(tableSQL))
	if err != nil {
		log.Fatal(pkgErr.WithStack(err))
		return err
	}

	return nil
}

func Test_pingDB(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Error("can't ping database")
	}
}
