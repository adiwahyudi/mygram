package service

import (
	"mygram/helper"
	"mygram/model"
	"mygram/repository/mocks"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestUserService_Add(t *testing.T) {

	userRepository := mocks.NewIUserRepository(t)

	type args struct {
		request model.UserRegisterRequest
	}
	tests := []struct {
		name     string
		us       *UserService
		args     args
		want     model.UserRegisterResponse
		mockFunc func()
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Case #1 - Register Success",
			us: &UserService{
				UserRepository: userRepository,
			},
			args: args{
				model.UserRegisterRequest{
					Email:    "adiwahyudi@mail.com",
					Username: "adiwahyudi",
					Password: "adiwahyudi",
					Age:      22,
				},
			},
			want: model.UserRegisterResponse{
				Username: "adiwahyudi",
			},
			mockFunc: func() {
				userRepository.
					On("Save", mock.Anything).
					Return(
						model.User{
							Username: "adiwahyudi",
							Email:    "adiwahyudi@mail.com",
							Password: "adiwahyudi",
							Age:      22,
						}, nil,
					)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			got, err := tt.us.Add(tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	userRepository := mocks.NewIUserRepository(t)

	type args struct {
		request model.UserLoginRequest
	}
	tests := []struct {
		name     string
		us       *UserService
		args     args
		want     model.UserLoginResponse
		mockFunc func()
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Case #1 - Login Success",
			us: &UserService{
				UserRepository: userRepository,
			},
			args: args{
				model.UserLoginRequest{
					Username: "adiwahyudi",
					Password: "adiwahyudi",
				}},
			want: model.UserLoginResponse{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiIn0.gSVmQwE743HtFlKzIfnHeL8cf4ThKB1RhuAYHnAKdTQ",
			},
			mockFunc: func() {
				hashPassword, _ := helper.HashPassword("adiwahyudi")
				userRepository.
					On("GetByUsername", mock.AnythingOfType("string")).
					Return(
						model.User{
							Email:    "adiwahyudi@mail.com",
							Password: hashPassword,
							Age:      22,
						}, nil,
					)
			},
			wantErr: false,
		},
		{
			name: "Case #2 - Login Failed (Incorrect email or Password)",
			us: &UserService{
				UserRepository: userRepository,
			},
			args: args{
				request: model.UserLoginRequest{
					Username: "adi",
					Password: "adi",
				},
			},
			want: model.UserLoginResponse{},
			mockFunc: func() {
				hashPassword, _ := helper.HashPassword("random_________thing")
				userRepository.
					On("GetByUsername", mock.AnythingOfType("string")).
					Return(
						model.User{
							Email:    "adiwahyudi@mail.com",
							Password: hashPassword,
							Age:      22,
						}, nil,
					)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := tt.us.Login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
