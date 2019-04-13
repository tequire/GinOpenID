package tokens

import (
	"context"

	"github.com/coreos/go-oidc"
)

var verifier *oidc.IDTokenVerifier

// Verifier is a klødafjøoaslkjdf
func Verifier() *oidc.IDTokenVerifier {
	return verifier
}

func init() {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "")
	if err != nil {
		panic(err.Error())
	}
	verifier = provider.Verifier(&oidc.Config{
		ClientID: "",
	})
}
