package awesomethingy

import "fmt"

type githubSecretProvider interface {
	GithubToken() string
}

func Serve(g githubSecretProvider) {

	blah := g.GithubToken()
	fmt.Println(blah)
}
