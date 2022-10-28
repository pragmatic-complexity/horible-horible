package secretservice

import "testing"

func TestProvider_GithubToken(t *testing.T) {
	tests := []struct {
		name string
		p    *Provider
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{}
			if got := p.GithubToken(); got != tt.want {
				t.Errorf("Provider.GithubToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
