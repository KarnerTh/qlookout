package lookout

import (
	"fmt"
	"testing"

	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/KarnerTh/qlookout/core/usecase/watch"
)

func createManagerWithMocks() (lookoutManager, *MockLookoutRepo, *watch.MockWatcher) {
	lookoutRepo := MockLookoutRepo{}
	watcher := watch.MockWatcher{}
	manager := lookoutManager{watcherIds: make(map[int]cron.EntryID), lookoutRepo: &lookoutRepo, watcher: &watcher}
	return manager, &lookoutRepo, &watcher
}

func TestStart(t *testing.T) {
	t.Run("Return error if lookouts could not be loaded", func(t *testing.T) {
		// Arrange
		manager, repoMock, _ := createManagerWithMocks()
		repoMock.On("Get").Return(nil, fmt.Errorf("error")).Once()

		// Act
		err := manager.Start()

		// Assert
		repoMock.AssertExpectations(t)
		assert.NotNil(t, err)
	})

	t.Run("All lookouts should be added to the watcher", func(t *testing.T) {
		// Arrange
		manager, repoMock, watcherMock := createManagerWithMocks()
		lookouts := []LookoutConfig{{Id: 1}, {Id: 2}}
		repoMock.On("Get").Return(lookouts, nil).Once()
		watcherMock.On("Watch", mock.Anything).Return(cron.EntryID(1)).Times(len(lookouts))

		// Act
		err := manager.Start()

		// Assert
		repoMock.AssertExpectations(t)
		watcherMock.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestWatch(t *testing.T) {
	t.Run("Do not add lookout to watch if already running", func(t *testing.T) {
		// Arrange
		manager, _, _ := createManagerWithMocks()
		manager.watcherIds[1] = cron.EntryID(1)

		// Act
		err := manager.Watch(1)

		// Assert
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "already running")
	})

	t.Run("Do not add lookout if lookout could not be found", func(t *testing.T) {
		// Arrange
		manager, repoMock, _ := createManagerWithMocks()
		repoMock.On("GetById", 1).Return(nil, fmt.Errorf("NOK")).Once()

		// Act
		err := manager.Watch(1)

		// Assert
		repoMock.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "NOK")
	})

	// Arrange
	t.Run("Add watcherId to list on success", func(t *testing.T) {
		manager, repoMock, watcherMock := createManagerWithMocks()
		lookoutId, watcherId := 1, cron.EntryID(7)
		repoMock.On("GetById", lookoutId).Return(&LookoutConfig{}, nil).Once()
		watcherMock.On("Watch", mock.Anything).Return(watcherId).Once()

		// Act
		err := manager.Watch(lookoutId)

		// Assert
		repoMock.AssertExpectations(t)
		watcherMock.AssertExpectations(t)
		assert.Nil(t, err)

		assert.Contains(t, manager.watcherIds, lookoutId)
		resultWatcherId := manager.watcherIds[lookoutId]
		assert.Equal(t, watcherId, resultWatcherId)
	})
}

func TestRemove(t *testing.T) {
	t.Run("Do not remove watch when lookoutId is not present", func(t *testing.T) {
		// Arrange
		manager, _, _ := createManagerWithMocks()

		// Act
		err := manager.Remove(1)

		// Assert
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "was not found")
	})

	t.Run("Remove lookout from watchlist on successfull removal", func(t *testing.T) {
		// Arrange
		manager, _, watcherMock := createManagerWithMocks()
		lookoutId, watcherId := 1, cron.EntryID(7)
		manager.watcherIds[lookoutId] = watcherId
		watcherMock.On("StopWatching", watcherId).Once()

		// Act
		err := manager.Remove(lookoutId)

		// Assert
		assert.Nil(t, err)
		assert.NotContains(t, manager.watcherIds, lookoutId)
	})
}
