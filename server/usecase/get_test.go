package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/gogapopp/gophkeeper/models"
	pb "github.com/gogapopp/gophkeeper/proto"
	"github.com/gogapopp/gophkeeper/server/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetDatas(t *testing.T) {
	testcases := []struct {
		name    string
		keys    map[string][]string
		mockErr error
		wantErr bool
	}{
		{
			name: "Success",
			keys: map[string][]string{
				"key1": {"value1", "value2"},
				"key2": {"value3", "value4"},
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "Error",
			keys: map[string][]string{
				"key2": {"value3", "value4"},
			},
			mockErr: errors.New("error"),
			wantErr: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockGetter := new(mocks.MockRepo)
			ctx := context.TODO()
			mockGetter.On("GetDatas", tc.keys).Return(models.SyncData{}, tc.mockErr)
			getUsecase := GetUsecase{get: mockGetter}
			syncRequest := &pb.SyncRequest{
				Keys: make(map[string]*pb.RepeatedUniqueKeys),
			}
			for k, v := range tc.keys {
				syncRequest.Keys[k] = &pb.RepeatedUniqueKeys{Values: v}
			}
			datas, err := getUsecase.GetDatas(ctx, syncRequest)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, datas)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, datas)
			}
			mockGetter.AssertExpectations(t)
		})
	}
}
