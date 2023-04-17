package service

import (
	"mygram/model"
	"mygram/repository/mocks"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSocialMediaService_GetAll(t *testing.T) {
	socialMediaRepository := mocks.NewISocialMediaRepository(t)
	tests := []struct {
		name     string
		sms      *SocialMediaService
		want     []model.SocialMediaResponse
		mockFunc func()
		wantErr  bool
	}{
		// TODO: Add test cases.

		{
			name: "Case #1 - Success",
			sms: &SocialMediaService{
				SocialMediaRepository: socialMediaRepository,
			},
			want: []model.SocialMediaResponse{
				{
					ID:             "1",
					UserID:         "1",
					Name:           "Twitter",
					SocialMediaURL: "twitter.com/adi",
				},
				{
					ID:             "2",
					UserID:         "1",
					Name:           "Telegram",
					SocialMediaURL: "@adi",
				},
			},
			mockFunc: func() {
				socialMediaRepository.On("Get").Return([]model.SocialMedia{
					{
						ID:             "1",
						UserID:         "1",
						Name:           "Twitter",
						SocialMediaURL: "twitter.com/adi",
					},
					{
						ID:             "2",
						UserID:         "1",
						Name:           "Telegram",
						SocialMediaURL: "@adi",
					},
				}, nil).Once()
			},
			wantErr: false,
		},
		{
			name: "Case #2 - Success (Empty Data)",
			sms: &SocialMediaService{
				SocialMediaRepository: socialMediaRepository,
			},
			want: []model.SocialMediaResponse{},
			mockFunc: func() {
				socialMediaRepository.On("Get").Return([]model.SocialMedia{}, nil).Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := tt.sms.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("SocialMediaService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SocialMediaService.GetAll() = %v, want %v", got, tt.want)
			}
			tt.sms.SocialMediaRepository = socialMediaRepository
		})
	}
}

func TestSocialMediaService_GetById(t *testing.T) {
	socialMediaRepository := mocks.NewISocialMediaRepository(t)

	type args struct {
		id string
	}
	tests := []struct {
		name     string
		sms      *SocialMediaService
		args     args
		want     model.SocialMediaResponse
		mockFunc func()
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Case #1 - Success",
			sms: &SocialMediaService{
				SocialMediaRepository: socialMediaRepository,
			},
			args: args{
				id: "1",
			},
			want: model.SocialMediaResponse{
				ID:             "1",
				UserID:         "1",
				Name:           "Twitter",
				SocialMediaURL: "twitter.com/adi",
			},
			mockFunc: func() {
				socialMediaRepository.On("GetOne", mock.Anything).Return(model.SocialMedia{
					ID:             "1",
					UserID:         "1",
					Name:           "Twitter",
					SocialMediaURL: "twitter.com/adi",
				}, nil).Once()
			},
			wantErr: false,
		},
		{
			name: "Case #2 - Not Found (Failed)",
			sms: &SocialMediaService{
				SocialMediaRepository: socialMediaRepository,
			},
			args: args{
				id: "1",
			},
			want: model.SocialMediaResponse{},
			mockFunc: func() {
				socialMediaRepository.On("GetOne", mock.Anything).Return(model.SocialMedia{}, model.ErrorNotFound).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := tt.sms.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SocialMediaService.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SocialMediaService.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSocialMediaService_Add(t *testing.T) {
	socialMediaRepository := mocks.NewISocialMediaRepository(t)

	type args struct {
		request model.SocialMediaCreateRequest
		userId  string
	}
	tests := []struct {
		name     string
		sms      *SocialMediaService
		args     args
		want     model.SocialMediaCreateResponse
		mockFunc func()
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Case #1 - Success Create Social Media",
			sms:  &SocialMediaService{SocialMediaRepository: socialMediaRepository},
			args: args{
				userId: "1",
				request: model.SocialMediaCreateRequest{
					Name:           "Twitter",
					SocialMediaURL: "twitter.com/adiwahyudi",
				},
			},
			want: model.SocialMediaCreateResponse{
				ID:             "1",
				UserID:         "1",
				Name:           "Twitter",
				SocialMediaURL: "twitter.com/adiwahyudi",
			},
			mockFunc: func() {
				socialMediaRepository.On("Save", mock.Anything).Return(model.SocialMedia{
					ID:             "1",
					UserID:         "1",
					Name:           "Twitter",
					SocialMediaURL: "twitter.com/adiwahyudi",
				}, nil).Once()
			},
			wantErr: false,
		},

		// {
		// 	name: "Case #2 - Fail (Bad Request)",
		// 	sms:  &SocialMediaService{SocialMediaRepository: socialMediaRepository},
		// 	want: model.SocialMediaCreateResponse{},
		// 	mockFunc: func() {
		// 		socialMediaRepository.On("Save", mock.AnythingOfType("string")).Return(model.SocialMedia{}, nil).Once()
		// 	},
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := tt.sms.Add(tt.args.request, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SocialMediaService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SocialMediaService.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
