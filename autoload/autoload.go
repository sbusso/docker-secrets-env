package autoload

import dockersecretsenv "github.com/sbusso/docker-secrets-env"

func init() {
	dockersecretsenv.LoadSecrets("/run/secrets")
}
