package simple

import (
	"testing"
)

// 測試兩個角色 還有 admin最高權限
func Test_RBAC(t *testing.T) {

	r := []role{
		{"角色A",
			"功能A",
		},
		{"角色B",
			"功能B",
		},
	}

	f := []feature{
		{"功能A",
			"權限1",
		},
		{"功能A",
			"權限2",
		},
		{"功能B",
			"權限3",
		},
		{"功能B",
			"權限4",
		},
	}

	p := []permission{
		{"權限1",
			"/p1",
			"GET",
		},
		{"權限2",
			"/p1",
			"POST",
		},
		{"權限3",
			"/p2",
			"GET",
		},
		{"權限4",
			"/p2",
			"POST",
		},
	}

	tests := []struct {
		name    string
		request []string
		want    bool
	}{
		{
			name:    "test1",
			request: []string{"角色A", "/p1", "GET"},
			want:    true,
		},
		{
			name:    "test2",
			request: []string{"角色A", "/p1", "POST"},
			want:    true,
		},
		{
			name:    "test3",
			request: []string{"角色B", "/p1", "GET"},
			want:    false,
		},
		{
			name:    "test4",
			request: []string{"角色B", "/p1", "POST"},
			want:    false,
		},
		{
			name:    "test5",
			request: []string{"admin", "/p1", "GET"},
			want:    true,
		},
		{
			name:    "test6",
			request: []string{"admin", "/p1", "POST"},
			want:    true,
		},
		{
			name:    "test7",
			request: []string{"角色A", "/p2", "GET"},
			want:    false,
		},
		{
			name:    "test8",
			request: []string{"角色A", "/p2", "POST"},
			want:    false,
		},
		{
			name:    "test9",
			request: []string{"角色B", "/p2", "GET"},
			want:    true,
		},
		{
			name:    "test10",
			request: []string{"角色B", "/p2", "POST"},
			want:    true,
		},
		{
			name:    "test11",
			request: []string{"admin", "/p2", "GET"},
			want:    true,
		},
		{
			name:    "test12",
			request: []string{"admin", "/p2", "POST"},
			want:    true,
		},
	}

	e, _ := geRBACModel(r, f, p)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := e.Enforce(tt.request[0], tt.request[1], tt.request[2])

			if err != nil {
				t.Errorf("getEnforcer() error = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("%v got %v, want %v", tt.request, got, tt.want)
			}
		})
	}
}
