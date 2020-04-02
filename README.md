# Good Vibes Only

<img src="./good-vibes-only.jpeg" width="400" />

A simple server for authorizing Spotify-based client applications. The official OAuth flow is available [here](https://developer.spotify.com/documentation/general/guides/authorization-guide/). Image by [@carltraw](https://unsplash.com/@carltraw)

## Getting Started

**1) Set up your Spotify Developer account**

Head over to their [dashboard](https://developer.spotify.com/dashboard/login). Log in with your Spotify account, create a Client ID, and register a `Redirect URI` set to http://localhost:3001/oauth.

This URI is where Spotify will return the user upon successful OAuth.

**2) Configure your environment**

Create a `.env` file at the root of this repo, and fill out your Spotify client params:

```
SPOTIFY_CLIENT_ID=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
SPOTIFY_SECRET_KEY=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

**3) Get the server running**

```bash
# start app
$ go run .

# healthcheck
$ curl localhost:3001/alive

# get spotify login URL
$ curl localhost:3001/login
```

_In the future, I'd like to make this whole thing a simple Docker image so you don't need to have the Go toolchain installed locally._

## Web app (Optional)

You can write your own client, but if you'd like to see a live example that uses this server, check out the [client repo](https://github.com/Robinnnnn/good-vibes-only-client).
