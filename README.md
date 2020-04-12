# Good Vibes Only

<img src="./good-vibes-only.jpeg" width="400" />

A simple server for authorizing Spotify-based client applications. The official OAuth flow is available [here](https://developer.spotify.com/documentation/general/guides/authorization-guide/). Image by [@carltraw](https://unsplash.com/@carltraw)

## Getting Started

### 1) Set up your Spotify Developer account

Head over to their [dashboard](https://developer.spotify.com/dashboard/login). Log in with your Spotify account, create a Client ID, and register a `Redirect URI` set to http://localhost:3001/oauth.

This URI is where Spotify will return the user upon successful OAuth.

### 2) Configure your environment

Create a `.env` file at the root of this repo, and fill out your Spotify client params:

```
SPOTIFY_CLIENT_ID=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
SPOTIFY_SECRET_KEY=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

### 3) Get the server running

You can do this using either Docker or Go:

#### With Docker

To use this method, you'll need to have [Docker installed](https://docs.docker.com/get-docker/) on your machine.

```bash
# build the container
$ docker build -t [NAME] .

# run the built image
$ docker run -p [PORT]:3001 [NAME]
```

You can use any name you want, as well as the port you want it to run on your machine. For example:

```bash
# build the container
$ docker build -t good-vibes-only-server .

# run the built image
$ docker run -p 4444:3001 good-vibes-only-server

# test that it's responsive
$ curl localhost:4444/alive
=> "ok"
```



#### With Go

To use this method, you'll need have the [Golang toolchain installed](https://golang.org/doc/install) on your machine.

```bash
# start app
$ go run .

# healthcheck
$ curl localhost:3001/alive

# get spotify login URL
$ curl localhost:3001/login
```

## Web app (Optional)

You can write your own client, but if you'd like to see a live example that uses this server, check out the [client repo](https://github.com/Robinnnnn/good-vibes-only-client).
