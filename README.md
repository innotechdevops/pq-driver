# pq-driver

## Install

```
$ go get github.com/innotechdevops/pq-driver
```

## How to use

- Wtih env

```golang
driver := pqdriver.New(pqdriver.ConfigEnv())
```

- With config

```golang
driver := pqdriver.New(pqdriver.Config{
    User:         os.Getenv("POSTGRES_USER"),
    Pass:         os.Getenv("POSTGRES_PASS"),
    Host:         os.Getenv("POSTGRES_HOST"),
    DatabaseName: os.Getenv("POSTGRES_DATABASE"),
    Port:         pqdriver.DefaultPort,
    SSLMode:      pqdriver.SSLModeDisable,
    MaxLifetime:  os.Getenv("MARIA_MAX_LIFETIME"),
    MaxIdleConns: os.Getenv("MARIA_MAX_IDLE_CONNS"),
    MaxOpenConns: os.Getenv("MARIA_MAX_OPEN_CONNS"),
})
```