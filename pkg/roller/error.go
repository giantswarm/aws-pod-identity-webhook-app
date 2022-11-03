package roller

import "github.com/giantswarm/microerror"

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

var unsupportedKindError = &microerror.Error{
	Kind: "unsupportedKindError",
}
