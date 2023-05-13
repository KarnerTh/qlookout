package notify

import (
	"testing"

	"github.com/KarnerTh/query-lookout/config"
	"github.com/KarnerTh/query-lookout/usecase/lookout"
	"github.com/KarnerTh/query-lookout/usecase/review"
	"github.com/stretchr/testify/mock"
)

func createNotifierWithMocks() (NotifyManager, *MockNotifier, *MockNotifier, *lookout.MockLookoutRepo, *config.MockConfig) {
	localNotifier := MockNotifier{}
	mailNotifier := MockNotifier{}
	lookoutRepo := lookout.MockLookoutRepo{}
	config := config.MockConfig{}

	return &notifyManager{
		localNotifier: &localNotifier,
		mailNotifier:  &mailNotifier,
		lookoutRepo:   &lookoutRepo,
		config:        &config,
	}, &localNotifier, &mailNotifier, &lookoutRepo, &config
}

func TestNotify(t *testing.T) {
	t.Run("Do not send anything if the result is a success", func(t *testing.T) {
		// Arrange
		manager, _, _, _, _ := createNotifierWithMocks()
		reviewResult := review.ReviewResult{
			Success: true,
		}

		// Act
		manager.Notify(reviewResult)

		// Assert
		// Testify will fail when methodes are called that are not mocked - so
		// will fail if some Send method is calledd
	})

	t.Run("NotifyLocal should send local notification", func(t *testing.T) {
		// Arrange
		manager, localNotifierMock, _, lookoutRepoMock, configMock := createNotifierWithMocks()
		lookout := lookout.LookoutConfig{Id: 1, NotifyLocal: true}
		lookoutRepoMock.On("GetById", lookout.Id).Return(&lookout, nil).Once()
		configMock.On("BaseUrl").Return("").Once()
		localNotifierMock.On("Send", mock.Anything).Return(nil).Once()
		reviewResult := review.ReviewResult{
			Success: false,
			Rule: review.ReviewRule{
				LookoutId: lookout.Id,
			},
		}

		// Act
		manager.Notify(reviewResult)

		// Assert
		localNotifierMock.AssertExpectations(t)
		lookoutRepoMock.AssertExpectations(t)
	})

	t.Run("NotifyMail should send local notification", func(t *testing.T) {
		// Arrange
		manager, _, mailNotifierMock, lookoutRepoMock, configMock := createNotifierWithMocks()
		lookout := lookout.LookoutConfig{Id: 1, NotifyMail: true}
		lookoutRepoMock.On("GetById", lookout.Id).Return(&lookout, nil).Once()
		configMock.On("BaseUrl").Return("").Once()
		mailNotifierMock.On("Send", mock.Anything).Return(nil).Once()
		reviewResult := review.ReviewResult{
			Success: false,
			Rule: review.ReviewRule{
				LookoutId: lookout.Id,
			},
		}

		// Act
		manager.Notify(reviewResult)

		// Assert
		mailNotifierMock.AssertExpectations(t)
		lookoutRepoMock.AssertExpectations(t)
	})

	t.Run("Send multiple notifications", func(t *testing.T) {
		// Arrange
		manager, localNotifierMock, mailNotifierMock, lookoutRepoMock, configMock := createNotifierWithMocks()
		lookout := lookout.LookoutConfig{Id: 1, NotifyLocal: true, NotifyMail: true}
		lookoutRepoMock.On("GetById", lookout.Id).Return(&lookout, nil).Once()
		configMock.On("BaseUrl").Return("").Once()
		localNotifierMock.On("Send", mock.Anything).Return(nil).Once()
		mailNotifierMock.On("Send", mock.Anything).Return(nil).Once()
		reviewResult := review.ReviewResult{
			Success: false,
			Rule: review.ReviewRule{
				LookoutId: lookout.Id,
			},
		}

		// Act
		manager.Notify(reviewResult)

		// Assert
		localNotifierMock.AssertExpectations(t)
		mailNotifierMock.AssertExpectations(t)
		lookoutRepoMock.AssertExpectations(t)
	})
}
