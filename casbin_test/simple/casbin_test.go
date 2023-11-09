package simple

import (
	"github.com/casbin/casbin/v2"
	scas "github.com/qiangmzsx/string-adapter/v2"
	"testing"
)

func Test_ACL(t *testing.T) {
	tests := []struct {
		name    string
		request []string
		want    bool
	}{
		{
			name:    "test1",
			request: []string{"alice", "data1", "read"},
			want:    true,
		},
		{
			name:    "test2",
			request: []string{"alice", "data1", "write"},
			want:    false,
		},
		{
			name:    "test3",
			request: []string{"bob", "data2", "write"},
			want:    true,
		},
		{
			name:    "test4",
			request: []string{"bob", "data2", "read"},
			want:    false,
		},
	}

	m, _ := geACLModel()
	line := `
p, alice, data1, read
p, bob, data2, write
`
	sa := scas.NewAdapter(line)
	e, _ := casbin.NewEnforcer(m, sa)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := e.Enforce(tt.request[0], tt.request[1], tt.request[2])

			if err != nil {
				t.Errorf("getEnforcer() error = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("cardNumberHidden() = %v, want %v", got, tt.want)
			}
		})
	}
}
