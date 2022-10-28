package secretservice

func newProvider() Provider {}

type Provider struct {
}

type gcpProvider struct {
	g   githubClient
	gcp gcpClient
	aws awsClient
}

func (p *Provider) GithubToken() string {

}

func (p *Provider) OtherToken() string {}

type githubClient interface{}
type gcpClient interface{}
type awsClient interface{}
