package pkg_test

import (
	"fmt"
	"testing"

	assignments "github.com/91diego/platform-challenge/pkg"
	"github.com/91diego/platform-challenge/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestShipmentAssignemts(t *testing.T) {

	type args struct {
		ShipmentFile string
		DriverFile   string
	}

	tests := []struct {
		name           string
		args           args
		expectedResult []models.AssignmentResponse
		err            error
	}{
		{
			name: "failed_test_shipment_file_not_valid",
			args: args{
				ShipmentFile: "file://shipmentfile.txt",
				DriverFile:   "test_files/txt/driverFile.txt",
			},
			expectedResult: nil,
			err:            fmt.Errorf("shipment file is not valid"),
		},
		{
			name: "failed_test_driver_file_not_valid",
			args: args{
				ShipmentFile: "test_files/txt/shipmentFile.txt",
				DriverFile:   "ile://driverfile.txt",
			},
			expectedResult: nil,
			err:            fmt.Errorf("driver file is not valid"),
		},
		{
			name: "Successful_Test_CSV_file",
			args: args{
				ShipmentFile: "test_files/csv/shipmentFile.csv",
				DriverFile:   "test_files/csv/driverFile.csv",
			},
			expectedResult: []models.AssignmentResponse{
				{
					ShipmentAddress: "Diego Rivera",
					Driver: models.Driver{
						Name: "Diego Rivera",
						SS:   11.25,
					},
				},
				{
					ShipmentAddress: "Lisboa",
					Driver: models.Driver{
						Name: "Fernando Ramirez",
						SS:   7.5,
					},
				},
				{
					ShipmentAddress: "Hilarion Romero Gil",
					Driver: models.Driver{
						Name: "Carlos Gonzalez",
						SS:   14,
					},
				},
				{
					ShipmentAddress: "Av Vallarta",
					Driver: models.Driver{
						Name: "Juan Mendoza",
						SS:   11,
					},
				},
				{
					ShipmentAddress: "Volcan San Francisco",
					Driver: models.Driver{
						Name: "Julio Mendez",
						SS:   6.5,
					},
				},
				{
					ShipmentAddress: "Av Fidel Velazquez",
					Driver: models.Driver{
						Name: "Juan Perez",
						SS:   5.5,
					},
				},
			},
			err: nil,
		},
		{
			name: "Successful_Test_TXT_file",
			args: args{
				ShipmentFile: "test_files/txt/shipmentFile.txt",
				DriverFile:   "test_files/txt/driverFile.txt",
			},
			expectedResult: []models.AssignmentResponse{
				{
					ShipmentAddress: "Diego Rivera",
					Driver: models.Driver{
						Name: "Diego Rivera",
						SS:   11.25,
					},
				},
				{
					ShipmentAddress: "Lisboa",
					Driver: models.Driver{
						Name: "Fernando Ramirez",
						SS:   7.5,
					},
				},
				{
					ShipmentAddress: "Hilarion Romero Gil",
					Driver: models.Driver{
						Name: "Carlos Gonzalez",
						SS:   14,
					},
				},
				{
					ShipmentAddress: "Av Vallarta",
					Driver: models.Driver{
						Name: "Juan Mendoza",
						SS:   11,
					},
				},
				{
					ShipmentAddress: "Volcan San Francisco",
					Driver: models.Driver{
						Name: "Julio Mendez",
						SS:   6.5,
					},
				},
				{
					ShipmentAddress: "Av Fidel Velazquez",
					Driver: models.Driver{
						Name: "Juan Perez",
						SS:   5.5,
					},
				},
			},
			err: nil,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			res, err := assignments.ShipmentAssignments(&tc.args.ShipmentFile, &tc.args.DriverFile)

			if err != nil {
				assert.EqualError(t, tc.err, err.Error())
			}
			assert.EqualValues(t, tc.expectedResult, res)
		})

	}
}
