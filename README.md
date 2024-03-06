# API-GORM-REDIS
Service-Account
Service-Journal

# Running in Docker
```console
docker-compose up -d
```

# Endpoint
```console
POST    - /account/daftar
POST    - /v1/account/tabung
POST    - /v1/account/tarik
POST    - /v1/account/transfer
GET     - /v1/account/saldo/:no_rekening
GET     - /v1/account/mutasi/:no_rekenig
```

## Core library

Library | Usage
-- | --
fiber | Base framework
postgres | Database
GORM | ORM
redis | Stream message
logrus | Logger library

And others library are listed on `go.mod` file