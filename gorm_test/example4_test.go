package main

import (
	"testing"
)

func Test_findWithCount2(t *testing.T) {
	type args struct {
		page *Page
		cond any
		data []Users
	}
	tests := []struct {
		name    string
		args    args
		want    []int64
		wantErr bool
	}{
		{"test1",
			args{
				page: nil,
				cond: nil,
				data: make([]Users, 0),
			},
			[]int64{7, 7},
			false,
		},
		{"test2",
			args{
				page: nil,
				cond: &UserCond{
					ID: 2,
				},
				data: make([]Users, 0),
			},
			[]int64{1, 1},
			false,
		},
		{"test3",
			args{
				page: nil,
				cond: UserCond{
					ID: 2,
				},
				data: make([]Users, 0),
			},
			[]int64{1, 1},
			false,
		},
		{"test4",
			args{
				page: &Page{
					PageIndex: 1,
					Size:      3,
				},
				cond: UserCond{
					ID: 2,
				},
				data: make([]Users, 0),
			},
			[]int64{1, 1},
			false,
		},
		{"test5",
			args{
				page: &Page{
					PageIndex: 1,
					Size:      3,
				},
				cond: nil,
				data: make([]Users, 0),
			},
			[]int64{7, 3},
			false,
		},
		{"test6",
			args{
				page: &Page{
					PageIndex: 1,
					Size:      2,
				},
				cond: "id in(1,2,3,4,5)",
				data: make([]Users, 0),
			},
			[]int64{5, 2},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := getDB()
			got, err := findWithCount2(db, tt.args.page, tt.args.cond, &tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("findWithCount2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want[0] {
				t.Errorf("findWithCount2() got = %v, want %v", got, tt.want)
			}

			length := int64(len(tt.args.data))

			if length != tt.want[1] {
				t.Errorf("findWithCount2() got = %v, want %v", got, tt.want)
			}

		})
	}
}
