package Config

import (
	"os"
)

type auth struct {
	Secret        string
	RefreshSecret string
}

func setupAuth() *auth {
	value := &auth{
		Secret:        os.Getenv("SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}

	return value
}
