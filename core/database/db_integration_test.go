package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type columnTestResult struct {
	Name       string
	isPrimary  bool
	isForeign  bool
	isUnique   bool
	isNullable bool
}

func TestDatabaseIntegrations(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	connectorFactory := NewConnectorFactory()

	for _, testCase := range GetTestDbConnections() {
		t.Run(testCase.DbType.String(), func(t *testing.T) {
			t.Run("Connect", func(t *testing.T) {
				// Act
				db, err := connectorFactory.NewConnector(testCase.ConnectionString)

				// Assert
				assert.Nil(t, err)
				assert.NotNil(t, db)
			})
		})
	}
}
