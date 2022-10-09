package oauth

import "golang.org/x/oauth2"

var _ oauth2.TokenSource = (*CachedTokenSource)(nil)

type CachedTokenSource struct {
	ts oauth2.TokenSource
}

func (ts *CachedTokenSource) Token() (*oauth2.Token, error) {
	token, err := ts.ts.Token()
	// if err == nil &&ts.store!=nil{
	// }
	return token, err
}
