package simple

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseRepository(t *testing.T) {
	databaseRepository := InisializeDatabaseRepository()
	assert.NotNil(t, databaseRepository)
	assert.NotNil(t, databaseRepository.DatabaseMongoDB)
	assert.NotNil(t, databaseRepository.DatabasePostgreSQL)

	fmt.Println("pg: " + databaseRepository.DatabasePostgreSQL.Name)
	fmt.Println("md: " + databaseRepository.DatabaseMongoDB.Name)
}
