package kodirpc

import "fmt"

type KodiUser struct {
	User string `env:"KODI_USER"`
	Pass string `env:"KODI_PASS"` //TODO: Find a better more secure way to supply this..
}

//New creates a new User object for use with RPC
func New(user, pass string) (*KodiUser, error) {
	if user == "" {
		return nil, fmt.Errorf("Kodi user name not supplied!")
	}

	if pass == "" {
		return nil, fmt.Errorf("Kodi password not supplied!")
	}

	return &KodiUser{User: user, Pass: pass}, nil
}

func (ku *KodiUser) GetUser() string {
	return ku.User
}

func (ku *KodiUser) GetPass() string {
	return ku.Pass
}
