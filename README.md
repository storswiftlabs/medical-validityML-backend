# medical-zkML-backend

This repository contains all code for the medical validityML backend.

## Getting Started

### Prerequisites

- Install [Golang](https://golang.google.cn/dl/), golang version >= 1.20
- Install [Docker](https://www.docker.com/get-started/), docker version >= v0.11.2



### Get the code

```shell
git clone  https://github.com/storswiftlabs/medical-validityML-backend.git
```

### Server startup

```shell
medical-validityML-backend
make build
```

Introducing the required libraries for Python

```shell
pip install -r requirements.txt
```

Configure database connections, IPFS information, and agreed addresses in the configuration file under the project root path (config.yaml).
database:

```yaml
database:
  mysql:
    driverName: MYSQL
    host: LOCALHOST
    port: '3306'
    user: ROOT
    password: EXAMPLE
    schema: 
    database: MEDICAL-ZKML
```

IPFS:

```yaml
ipfs:
  url: https://api.nft.storage/upload
  auth: IPFS_AUTH
```

GIZA:
```yaml
giza:
  user: GIZA_USER
  passwd: GIZA_PASSWORD
  email: GIZA_EMAIL
```

Start service

```shell
bash medical &
```