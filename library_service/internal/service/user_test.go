package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"onelab/internal/model"
	"onelab/internal/repository"
	mock_repository "onelab/internal/repository/mocks"
	"testing"
)

func TestUserServiceValidID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Success",
			args: args{1},
			want: true,
		},
		{
			name: "Success",
			args: args{0},
			want: false,
		},
		{
			name: "Success",
			args: args{10},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidID(tt.args.id); got != tt.want {
				t.Errorf("ValidID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		prepare func(f *mock_repository.MockIUserRepository)
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{context.Background(), model.User{Email: "test@gmail.com"}},
			want:    1,
			wantErr: false,
			prepare: func(f *mock_repository.MockIUserRepository) {
				f.EXPECT().Create(gomock.Any(), model.User{Email: "test@gmail.com"}).
					Return(uint(1), nil)
			},
		},
		{
			name:    "Success",
			args:    args{context.Background(), model.User{Email: "test"}},
			want:    0,
			wantErr: true,
			prepare: func(f *mock_repository.MockIUserRepository) {
				f.EXPECT().Create(gomock.Any(), model.User{Email: "test"}).
					Return(uint(0), errors.New("no @ in email"))
			},
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		userRepo := mock_repository.NewMockIUserRepository(ctrl)
		tt.prepare(userRepo)
		s := NewUserService(&repository.Manager{User: userRepo})
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
