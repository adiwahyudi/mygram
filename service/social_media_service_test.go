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
			name: "Case #1 - Success (Empty Data)",
			sms: &SocialMediaService{
				SocialMediaRepository: socialMediaRepository,
			},
			want: []model.SocialMediaResponse{},
			mockFunc: func() {
				socialMediaRepository.On("Get").Return([]model.SocialMedia{}, nil)
			},
			wantErr: false,
		},
		{
			name: "Case #2 - Success",
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
				}, nil)
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
	// socialMediaRepository2 := mocks.NewISocialMediaRepository(t)

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
				}, nil)
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
				socialMediaRepository.On("GetOne", mock.Anything).Return(model.SocialMedia{}, model.ErrorNotFound)
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
