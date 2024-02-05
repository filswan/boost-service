package service

import boost "github.com/filswan/swan-boost-lib/client"

type Boost struct {
	client *boost.Client
}

func NewBoost(apiNode string, repo string) (*Boost, error) {
	client, err := new(boost.Client).WithRepo(repo).WithUrl(apiNode)
	if err != nil {
		return nil, err
	}
	return &Boost{client: client}, nil
}

func (b *Boost) ProviderStorageAsk(provider string) (*boost.AskInfo, error) {
	return b.client.StorageAsk(provider, 0, 0)
}

var BoostService *Boost

func Init(apiNode string, repo string) error {
	var err error
	BoostService, err = NewBoost(apiNode, repo)
	return err
}
