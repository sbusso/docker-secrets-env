# Docker Secrets Env

A small utility to load docker secrets as ENV Variables.

Explicit load:


```go
import dockersecretsenv "github.com/sbusso/docker-secrets-env"

func main() {
 dockersecretsenv.LoadSecrets()
 ...
}
```

or automatic loading:

```go
import _ "github.com/sbusso/docker-secrets-env/autoload"
```
