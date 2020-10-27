package booking

import (
	"testing"

	"gorm.io/gorm"

	"commerceiq.ai/ticketing/internal/testhelpers"
	"commerceiq.ai/ticketing/pkgs/models"
)

func TestBookSeatsInput_Validate(t *testing.T) {
	db := testhelpers.SetupDB()

	type fields struct {
		ShowID      int
		SeatNumbers []int
		UserID      int
		SeatType    models.SeatType
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestValidate",
			fields: fields{
				ShowID:      1,
				SeatNumbers: []int{1, 2, 3},
				UserID:      1,
				SeatType:    models.Recliner,
			},
			args: args{
				db: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BookSeatsInput{
				ShowID:      tt.fields.ShowID,
				SeatNumbers: tt.fields.SeatNumbers,
				UserID:      tt.fields.UserID,
				SeatType:    tt.fields.SeatType,
			}
			if err := bs.Validate(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
