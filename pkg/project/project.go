package project

var (
	description = "Restart pods that need IRSA settings injected."
	gitSHA      = "n/a"
	name        = "aws-pod-identity-webhook"
	source      = "https://github.com/giantswarm/aws-pod-identity-webhook"
	version     = "1.19.1-dev"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
