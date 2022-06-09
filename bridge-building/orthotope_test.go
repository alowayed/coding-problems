package orth

import (
	"reflect"
	"testing"
)

var (
	twoD = map[string]bool{
		"0-0": true,
		"0-1": true,
		"0-2": true,
		"0-3": true,

		"1-0": true,
		"1-1": true,
		"1-2": true,
		"1-3": true,

		"2-0": true,
		"2-1": true,
		"2-2": true,
		"2-3": true,
	}

	threeD = map[string]bool{
		"0-0-0": true,
		"0-0-1": true,

		"0-1-0": true,
		"0-1-1": true,

		"1-0-0": true,
		"1-0-1": true,

		"1-1-0": true,
		"1-1-1": true,
	}
)

func TestNew(t *testing.T) {
	type args struct {
		lengths []int
	}
	tests := []struct {
		name    string
		args    args
		want    *Orthotope
		wantErr bool
	}{
		{
			name: "0-D",
			args: args{
				lengths: []int{},
			},
			want: &Orthotope{
				Lengths:    []int{},
				bridges:    map[string]bool{},
				nonBridges: map[string]bool{},
			},
			wantErr: false,
		},
		{
			name: "1-D",
			args: args{
				lengths: []int{3},
			},
			want: &Orthotope{
				Lengths: []int{3},
				bridges: map[string]bool{},
				nonBridges: map[string]bool{
					"0": true,
					"1": true,
					"2": true,
				},
			},
			wantErr: false,
		},
		{
			name: "2-D",
			args: args{
				lengths: []int{3, 4},
			},
			want: &Orthotope{
				Lengths:    []int{3, 4},
				bridges:    map[string]bool{},
				nonBridges: twoD,
			},
			wantErr: false,
		},
		{
			name: "3-D",
			args: args{
				lengths: []int{2, 2, 2},
			},
			want: &Orthotope{
				Lengths:    []int{2, 2, 2},
				bridges:    map[string]bool{},
				nonBridges: threeD,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.lengths)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %+v, want %+v", *got, *tt.want)
			}
		})
	}
}

func TestOrthotope_Build(t *testing.T) {

	type fields struct {
		Lengths    []int
		bridges    map[string]bool
		nonBridges map[string]bool
	}
	type args struct {
		locs []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Orthotope
		wantErr bool
	}{
		{
			name: "2D build",
			fields: fields{
				Lengths:    []int{3, 4},
				bridges:    map[string]bool{},
				nonBridges: twoD,
			},
			args: args{
				locs: []int{1, 2},
			},
			want: &Orthotope{
				Lengths: []int{3, 4},
				bridges: map[string]bool{
					"1-2": true,
				},
				nonBridges: map[string]bool{
					"0-0": true,
					"0-1": true,
					"0-2": true,
					"0-3": true,

					"1-0": true,
					"1-1": true,
					// "1-2": true, REMOVED
					"1-3": true,

					"2-0": true,
					"2-1": true,
					"2-2": true,
					"2-3": true,
				},
			},
			wantErr: false,
		},
		{
			name: "3D build",
			fields: fields{
				Lengths:    []int{2, 2, 2},
				bridges:    map[string]bool{},
				nonBridges: threeD,
			},
			args: args{
				locs: []int{0, 1, 0},
			},
			want: &Orthotope{
				Lengths: []int{2, 2, 2},
				bridges: map[string]bool{
					"0-1-0": true,
				},
				nonBridges: map[string]bool{
					"0-0-0": true,
					"0-0-1": true,

					// "0-1-0": true, REMOVED
					"0-1-1": true,

					"1-0-0": true,
					"1-0-1": true,

					"1-1-0": true,
					"1-1-1": true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Orthotope{
				Lengths:    tt.fields.Lengths,
				bridges:    tt.fields.bridges,
				nonBridges: tt.fields.nonBridges,
			}
			if err := o.Build(tt.args.locs...); (err != nil) != tt.wantErr {
				t.Errorf("Orthotope.Build() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(o, tt.want) {
				t.Errorf("Orthotope.Build() -> %+v, want %+v", *o, *tt.want)
			}
		})
	}
}

func TestOrthotope_BuildRandom(t *testing.T) {

	type fields struct {
		Lengths    []int
		bridges    map[string]bool
		nonBridges map[string]bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Orthotope
		wantErr bool
	}{
		{
			name: "single 1D",
			fields: fields{
				Lengths:    []int{1},
				bridges:    map[string]bool{},
				nonBridges: map[string]bool{"0": true},
			},
			want: &Orthotope{
				Lengths:    []int{1},
				bridges:    map[string]bool{"0": true},
				nonBridges: map[string]bool{},
			},
			wantErr: false,
		},
		{
			name: "double 1D",
			fields: fields{
				Lengths:    []int{2},
				bridges:    map[string]bool{"1": true},
				nonBridges: map[string]bool{"0": true},
			},
			want: &Orthotope{
				Lengths: []int{2},
				bridges: map[string]bool{
					"0": true,
					"1": true,
				},
				nonBridges: map[string]bool{},
			},
			wantErr: false,
		},
		{
			name: "all occupied",
			fields: fields{
				Lengths: []int{2},
				bridges: map[string]bool{
					"0": true,
					"1": true,
				},
				nonBridges: map[string]bool{},
			},
			want: &Orthotope{
				Lengths: []int{2},
				bridges: map[string]bool{
					"0": true,
					"1": true,
				},
				nonBridges: map[string]bool{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Orthotope{
				Lengths:    tt.fields.Lengths,
				bridges:    tt.fields.bridges,
				nonBridges: tt.fields.nonBridges,
			}
			if err := o.BuildRandom(); (err != nil) != tt.wantErr {
				t.Errorf("Orthotope.BuildRandom() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(o, tt.want) {
				t.Errorf("Orthotope.BuildRandom() -> %+v, want %+v", *o, *tt.want)
			}
		})
	}
}

func TestOrthotope_Built(t *testing.T) {
	type fields struct {
		Lengths    []int
		bridges    map[string]bool
		nonBridges map[string]bool
	}
	type args struct {
		locs []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Orthotope{
				Lengths:    tt.fields.Lengths,
				bridges:    tt.fields.bridges,
				nonBridges: tt.fields.nonBridges,
			}
			got, err := o.Built(tt.args.locs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Orthotope.Built() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Orthotope.Built() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrthotope_BridgeComplete(t *testing.T) {
	type fields struct {
		Lengths    []int
		bridges    map[string]bool
		nonBridges map[string]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty",
			fields: fields{
				Lengths:    []int{3, 4},
				bridges:    map[string]bool{},
				nonBridges: twoD,
			},
			want: false,
		},
		{
			name: "disconnected bridges",
			fields: fields{
				Lengths: []int{3, 4},
				bridges: map[string]bool{
					"0-0": true,
					"0-2": true,

					"1-1": true,

					"2-1": true,
				},
				nonBridges: map[string]bool{
					"0-1": true,
					"0-3": true,

					"1-0": true,
					"1-2": true,
					"1-3": true,

					"2-0": true,
					"2-2": true,
					"2-3": true,
				},
			},
			want: false,
		},
		{
			name: "connected bridges",
			fields: fields{
				Lengths: []int{3, 4},
				bridges: map[string]bool{
					"0-0": true,
					"0-2": true,

					"1-0": true,
					"1-1": true,

					"2-1": true,
				},
				nonBridges: map[string]bool{
					"0-1": true,
					"0-3": true,

					"1-2": true,
					"1-3": true,

					"2-0": true,
					"2-2": true,
					"2-3": true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Orthotope{
				Lengths:    tt.fields.Lengths,
				bridges:    tt.fields.bridges,
				nonBridges: tt.fields.nonBridges,
			}
			if got := o.BridgeComplete(); got != tt.want {
				t.Errorf("Orthotope.BridgeComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrthotope_inBound(t *testing.T) {
	type fields struct {
		Lengths    []int
		bridges    map[string]bool
		nonBridges map[string]bool
	}
	type args struct {
		locs []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Orthotope{
				Lengths:    tt.fields.Lengths,
				bridges:    tt.fields.bridges,
				nonBridges: tt.fields.nonBridges,
			}
			if got := o.inBound(tt.args.locs...); got != tt.want {
				t.Errorf("Orthotope.inBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key(t *testing.T) {
	type args struct {
		locs []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := key(tt.args.locs...); got != tt.want {
				t.Errorf("key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_locations(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := locations(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("locations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("locations() = %v, want %v", got, tt.want)
			}
		})
	}
}
