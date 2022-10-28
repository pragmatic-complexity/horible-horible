package awesomethingy

import (
	"testing"
)

type fakeSecret struct{ token string }

func (f *fakeSecret) GithubToken() string {
	return f.token
}

func TestServe(t *testing.T) {
	tests := []struct {
		name string
		ts   githubSecretProvider
	}{
		{name: "moo", ts: &fakeSecret{token: "bah"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Serve(tt.ts)
		})
	}
}
