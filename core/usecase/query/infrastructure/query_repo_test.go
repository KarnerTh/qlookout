package infrastructure

import (
	"log/slog"
	"testing"

	"github.com/KarnerTh/qlookout/core/database"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	connectorFactory := database.NewConnectorFactory()

	for _, testCase := range database.GetTestDbConnections() {
		t.Run(testCase.DbType.String(), func(t *testing.T) {
			db, err := connectorFactory.NewConnector(testCase.ConnectionString)
			if err != nil {
				slog.Error("Could not get test db connection", slog.Any("error", err), slog.String("cs", testCase.ConnectionString))
				t.FailNow()
			}

			t.Run("Simple count", func(t *testing.T) {
				// Arrange
				repo := NewQueryRepo(db)
				query := "select count(*) as \"count\" from test_table;"

				// Act
				result, err := repo.Query(query)

				// Assert
				assert.Nil(t, err)
				assert.Len(t, result.Columns, 1)
				assert.Len(t, result.Rows, 1)
				assert.Equal(t, "count", result.Columns[0])
				assert.Equal(t, int64(1), result.Rows[0]["count"])
			})
		})
	}
}
