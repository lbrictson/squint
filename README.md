# squint
A status page for B2B and SaaS providers

Features
- [ ] Simple setup and administration
- [ ] Several email backends
- [ ] Unlimited pages
- [ ] Per page branding
- [ ] Metric display
- [ ] JSON management API
- [ ] Per page security

### Development Dependencies you need to install

- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [docker](https://docs.docker.com/get-docker/)
- [air](https://github.com/cosmtrek/air)

### Configuring squint

All runtime configuration is done through environment variables. Below are available variables and their defaults.

```
# Database config options
SQUINT_DBUSER="postgres"
SQUINT_DBPASSWORD="postgres"
SQUINT_DBNAME="squint"
SQUINT_DBPORT=5432
SQUINT_DBHOST="localhost"
# Logger config options
SQUINT_LOGLEVEL="info" // info,warn,error,debug
SQUINT_LOGFILE="" // Set file location to disable stdout logging
SQUINT_LOGJSON=true // Set to false to disable json logging
```