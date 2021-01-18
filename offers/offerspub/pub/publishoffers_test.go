package pub

import (
	"reflect"
	"testing"

	"github.com/devminnu/assignments/offers/offerspub/config"
	"github.com/devminnu/assignments/offers/offerspub/logger"
	"github.com/devminnu/assignments/offers/offerspub/model"
)

func TestPublishOffers(t *testing.T) {
	config.Load()
	logger.Init()
	Init()
	type args struct {
		offers model.Offers
	}
	tests := []struct {
		name    string
		args    args
		want    model.Offers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "NoOffers",
			args: args{
				offers: model.Offers{},
			},
			want:    model.Offers{},
			wantErr: true,
		},
		{
			name: "NILOffers",
			args: args{
				offers: model.Offers{
					Offers: nil,
				},
			},
			want:    model.Offers{},
			wantErr: true,
		},
		{
			name: "Offer",
			args: args{
				offers: model.Offers{
					Offers: []model.Offer{
						{
							CmOfferID: "8f6995366e854c9faf1d9f3d233702b8",
							Hotel: model.Hotel{
								HotelID:   "BH~46456",
								Name:      "Hawthorn Suites by Wyndham Eagle CO",
								Country:   "US",
								Address:   "0315 Chambers Avenue, 81631",
								Latitude:  39.660193,
								Longitude: -106.824123,
								Telephone: "+1-970-3283000",
								Amenities: []model.Amenity{
									{
										Name: "Business Centre",
									},
									{
										Name: "Fitness Room/Gym",
									},
								},
								Description: "Stay a while in beautiful mountain country at this Hawthorn Suites by Wyndham Eagle CO hotel, just off Interstate 70, only 6 miles from the Vail/Eagle Airport and close to skiing, golfing, Eagle River and great restaurants. Pets are welcome at this h",
								RoomCount:   1,
								Currency:    "USD",
							},
							// RoomID:       "",
							Room: model.Room{
								RoomID:      "S2Q",
								Description: "JUNIOR SUITES WITH 2 QUEEN BEDS",
								Name:        "JUNIOR SUITES WITH 2 QUEEN BED",
								Capacity: model.Capacity{
									MaxAdults:     2,
									ExtraChildren: 2,
								},
							},
							// RatePlanID:   "",
							RatePlan: model.RatePlan{
								RatePlanID: "BAR",
								CancellationPolicy: []model.CancellationPolicy{
									{
										Type:              "Free cancellation",
										ExpiresDaysBefore: 2,
									},
								},
								Name: "BEST AVAILABLE RATE",
								OtherConditions: []model.OtherCondition{
									{
										Name: "CXL BY 2 DAYS PRIOR TO ARRIVAL-FEE 1 NIGHT 2 DAYS PRIOR TO ARRIVAL",
									},
									{
										Name: "BEST AVAILABLE RATE",
									},
								},
								MealPlan: "Room only",
							},
							OriginalData: model.OriginalData{
								GuaranteePolicy: model.GuaranteePolicy{
									Required: true,
								},
							},
							Capacity: model.Capacity{
								MaxAdults:     2,
								ExtraChildren: 2,
							},
							Number:   1,
							Price:    1520,
							Currency: "USD",
							CheckIn:  "2020-11-18",
							CheckOut: "2020-11-20",
							Fees: []model.Fee{
								{
									Type:        "CountyTax",
									Description: "COUNTY TAX PER STAY",
									Included:    true,
									Percent:     17.5,
								},
							},
						},
					},
				},
			},
			want:    model.Offers{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PublishOffers(tt.args.offers)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublishOffers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublishOffers() = %v, want %v", got, tt.want)
			}
		})
	}
}
