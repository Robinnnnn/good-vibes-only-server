# Good Vibes Only

<img src="./good-vibes-only.jpeg" width="400" />

A simple server for authorizing Spotify-based client applications. The official OAuth flow is available [here](https://developer.spotify.com/documentation/general/guides/authorization-guide/). Image by [@carltraw](https://unsplash.com/@carltraw)

## Local development

```bash
# start app
$ go run .

# healthcheck
$ curl localhost:3001/alive

# get spotify login URL
$ curl localhost:3001/login
```

## Web app

TODO
