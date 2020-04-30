# Clusauth

Lakelady OpenID/OAuth Authentication service

> Fork of the SSO and OAuth / OIDC login solution for Nginx using the auth_request module [vouch-proxy](https://github.com/vouch/vouch-proxy)

## Configuration

First, vouch-proxy uses dedicated configuration file

Then, environment variables are used to define docker containers

### Vouch configuration file

configuration files are 
- `config/config.yml` in production mode (copy `config/config_sample.yml` if config is missing)
- `config/config.dev.yml` in dev mode

According to this [example](https://github.com/vouch/vouch-proxy/blob/master/config/config.yml_example), you have to specify `oauth`.

> Remember to restart the projects after modifying configuration files (`make docker-[dev-]restart`)

### Environment variables

- CLUSAUTH_IMAGE: image name. Default is `clusauth[:dev]`
- CLUSAUTH_CONTAINER: container name. Default is `clusauth[-dev]`
- CLUSAUTH_PORT: external port. Default is `8003`
- CLUSAUTH_CONFIG: path to config directory. Default is `../config/`
- CLUSAUTH_DATA: path to db data directory. Default is `../data/`
- CLUSAUTH_CONTEXT: path to root project directory. Default is `../`
- VOUCH_CONFIG: path to vouch config file in the container. Default is `/config/config.[dev.]yml`

## Usage

### Routes

- /login: redirects to provider authentication page and returns the header `X-Clusauth-User` and the authorization cookie (httpOnly,secured,lax by default) `ClusauthCookie` if authorization succeed 
- /validate: if authorization header is `Bearer {X-Clusauth-Token}` or the authorization cookie (httpOnly,secure,lax by default) `ClusauthCookie` is validated, response status code is 200 and the header `X-Clusauth-User` is retrieved
- /auth: try `/validate`, and in case of failure, try `/login`
- /logout: delete authorisation cookie in vouch-proxy database, `ClusauthCookie`
- /token: if `/validate` succeed, then generate a JWT for using in Authorization Bearer header
- /ping: check if clusauth is available

### Headers

- `X-Clusauth-User`: user email
- `X-Clusauth-Token`: 3 months expiration JWT to reuse such as Bearer authorization given by the route `/token`
- `ClusauthCookie`: httpOnly,secure,lax cookie which contains a JWT of duration 3hr

### Makefile

> Require [docker-compose](https://docs.docker.com/compose/install/)

Mount the configuration file such as a volume and expose service at http://localhost:8003

```bash
make docker-up
# or with args : for example, for building before launching the container in detached mode => docker-compose up --build -d
ARGS="--build" make docker-up
# or in development mode (hot-reload)
make docker-dev-up
```

> Other commands are docker-\[dev-](build|stop|down|logs|restart|config|tty|cmd) where cmd permits to custom docker-compose command with env var CMD: `CMD=events make docker-cmd` equivalent to docker-compose -p lakelady -f deployments/docker-compose.yml events

Enjoy !
