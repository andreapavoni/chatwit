# ChaTwit

Do realtime chat with your twitter friends. This is a simple experimenti with Go and websockets, it's not intended for production!

## Setup

### Install missing packages

```make setup```

### Build binary

```make build```

## Compile assets (for deploy)

### Check support for SASS & Co

```make assets-check-support```

### Build assets

```make assets-bundle```

# Todo

* Add tests
* Add better logging
* Be paranoid with ongoing errors
* Check if session is already running when logged out
* Redirect to somewhere when trying to login with twitter but user is already
  logged in
