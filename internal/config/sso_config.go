package config

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"os"
)

func SetupIdentityProviders() {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), os.Getenv("GOOGLE_CALLBACK_URL"), "profile", "email"),
	)
}
