# ChaTwit

Do realtime chat with your twitter friends.

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

* Add better logging
* be paranoid with ongoing errors
* check if session is already running when logged out
* redirect to somewhere when trying to login with twitter but user is already
  logged in
* Show current user names (a modal or a sidebar)
