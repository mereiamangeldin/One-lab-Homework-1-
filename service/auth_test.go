package service

import (
	"github.com/golang/mock/gomock"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/repository"
	mock_repository "github.com/mereiamangeldin/One-lab-Homework-1/repository/mocks"
	"testing"
)

func TestAuthorizationService_CreateUser(t *testing.T) {
	type fields struct {
		Repo *repository.Repository
	}
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
		prepare func(f *mock_repository.MockIAuthorizationRepository)
	}{
		{
			name:    "success",
			args:    args{model.User{Name: "Merei"}},
			want:    1,
			wantErr: false,
			prepare: func(f *mock_repository.MockIAuthorizationRepository) {
				f.EXPECT().CreateUser(gomock.Any()).Return(uint(1), nil)
			},
		},
		{
			name:    "fail",
			args:    args{model.User{Name: "Ais"}},
			want:    2,
			wantErr: true,
			prepare: func(f *mock_repository.MockIAuthorizationRepository) {
				f.EXPECT().CreateUser(gomock.Any()).Return(uint(3), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			authRepo := mock_repository.NewMockIAuthorizationRepository(ctrl)
			tt.prepare(authRepo)
			s := NewAuthorizationService(&repository.Repository{Auth: authRepo})
			got, err := s.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
