package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/RamyChaabane/VoteApp/backend/internal/mocks"
	service "github.com/RamyChaabane/VoteApp/backend/internal/usecase/vote"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestVoteService_Vote(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		option    string
		mockError error
		wantErr   bool
	}{
		{
			name:    "successful vote",
			option:  "Cats",
			wantErr: false,
		},
		{
			name:      "repository error",
			option:    "Dogs",
			mockError: errors.New("redis error"),
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockVoteRepository(ctrl)
			svc := service.NewService(mockRepo)

			mockRepo.EXPECT().IncrementVote(ctx, tt.option).
				Return(tt.mockError).Times(1)

			err := svc.Vote(ctx, tt.option)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
