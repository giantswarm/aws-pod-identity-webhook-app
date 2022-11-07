package ownerfinder

import "github.com/giantswarm/microerror"

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

var unsupportedOwnerKindError = &microerror.Error{
	Kind: "unsupportedOwnerKindError",
}
