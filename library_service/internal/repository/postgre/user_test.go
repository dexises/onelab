package postgre

import (
	"context"
	"log"
	"onelab/internal/config"
	"onelab/internal/model"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name string
		args args
		// не знаю на сколько правильно это будет, но я создал тестовую функцию create которая возвращает email, а не id
		want    string
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{context.Background(), model.User{Name: "test", Email: "test@test.test"}},
			want:    "test@test.test",
			wantErr: false,
		},
	}

	repo := CreateUserRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.TestCreate(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
			email, err := repo.GetByEmail(context.Background(), got)
			if err != nil {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
			log.Print(email)
		})
	}
}

func CreateUserRepo(t *testing.T) *UserRepo {
	cfg := config.New()
	db, _ := NewConnection(cfg.DB)
	return NewUserRepo(db)
}
