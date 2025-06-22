package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/RamyChaabane/VoteApp/backend/internal/domain/vote"
	"github.com/RamyChaabane/VoteApp/backend/internal/handler"
	"github.com/RamyChaabane/VoteApp/backend/internal/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestVoteHandler_ServeHTTP(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		method         string
		body           string
		mockServiceErr error
		expectedCode   int
		expectedBody   string
	}{
		{
			name:         "OPTIONS request",
			method:       http.MethodOptions,
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "Invalid vote option",
			method:       http.MethodPost,
			body:         "vote=Birds",
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid vote option",
		},
		{
			name:         "Valid vote - Cats",
			method:       http.MethodPost,
			body:         "vote=Cats",
			expectedCode: http.StatusOK,
			expectedBody: "Vote for Cats recorded!",
		},
		{
			name:           "Service error",
			method:         http.MethodPost,
			body:           "vote=Dogs",
			mockServiceErr: errors.New("db error"),
			expectedCode:   http.StatusInternalServerError,
			expectedBody:   "Vote failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockService(ctrl)
			h := handler.NewVoteHandler(mockService)

			req := httptest.NewRequest(tt.method, "/", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()

			if tt.method == http.MethodPost && vote.IsValidOption(strings.TrimPrefix(tt.body, "vote=")) {
				mockService.EXPECT().Vote(context.Background(), strings.TrimPrefix(tt.body, "vote=")).
					Return(tt.mockServiceErr).Times(1)
			}

			h.ServeHTTP(w, req)

			res := w.Result()
			assert.Equal(t, tt.expectedCode, res.StatusCode)

			if tt.expectedBody != "" {
				defer res.Body.Close()
				body := make([]byte, res.ContentLength)
				res.Body.Read(body)
				assert.Equal(t, tt.expectedBody, string(body))
			}
		})
	}
}
