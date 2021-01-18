package service

import (
	"testing"

	"github.com/devminnu/assignments/offers/offerspub/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (mock *MockService) PublishOffers(model.Offers) (model.Offers, error) {
	args := mock.Called()
	return args.Get(0).(model.Offers), args.Error(1)
}

func Test_offersService_PublishOffers(t *testing.T) {
	mocksvc := new(MockService)
	type fields struct {
		Offers model.Offers
	}
	//
	offers := model.Offers{}
	mocksvc.On("PublishOffers", offers).Return(offers, nil)

	//
	testsvc := New()
	_, err := testsvc.PublishOffers(offers)
	mocksvc.AssertExpectations(t)

	assert.Equal(t, err.Error(), "NO_OFFERS")
	type args struct {
		offers model.Offers
	}
	// tests := []struct {
	// 	name    string
	// 	fields  fields
	// 	args    args
	// 	want    model.Offers
	// 	wantErr bool
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		service := &offersService{
	// 			Offers: tt.fields.Offers,
	// 		}
	// 		got, err := service.PublishOffers(tt.args.offers)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("offersService.PublishOffers() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("offersService.PublishOffers() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
