package main

import (
	"reflect"
	"testing"
)

func TestCalculator_Multiply(t *testing.T) {
	type args struct {
		n int
		d int
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *Calculator
		inspect func(r *Calculator, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args

		want1 int
	}{
		{
			name:"123",
			init: func(t *testing.T) *Calculator {
				return &Calculator{}
			},
			args: func(t *testing.T) args {
				return args{
					n: 1,
					d: 2,
				}
			},
			want1: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			got1 := receiver.Multiply(tArgs.n, tArgs.d)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Calculator.Multiply got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	type args struct {
		n int
		d int
	}
	tests := []struct {
		name    string
		init    func(t *testing.T) *Calculator
		inspect func(r *Calculator, t *testing.T) //inspects receiver after test run

		args func(t *testing.T) args

		want1      int
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			receiver := tt.init(t)
			got1, err := receiver.Divide(tArgs.n, tArgs.d)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Calculator.Divide got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("Calculator.Divide error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
