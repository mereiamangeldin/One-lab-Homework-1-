package postgre

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/config"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestAuthorizationRepository_CreateUser(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{user: model.User{Name: "Merei"}},
			want:    1,
			wantErr: false,
		},
		{
			name:    "fail",
			args:    args{user: model.User{Name: "Ais"}},
			want:    3,
			wantErr: true,
		},
	}
	repo := AuthRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AuthorizationRepository{
				db: repo.db,
			}
			got, err := r.CreateUser(tt.args.user)
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

func AuthRepo(t *testing.T) *AuthorizationRepository {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	db, err := Dial(cfg.PgURL)
	return NewAuthorizationRepository(db)
}
