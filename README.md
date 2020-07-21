# feedgomail
parses rss feeds and sends them to a dedicated email address

# Build

- install deps:
```
make deps_install
```
- build:
```
make build
```
see output in `/build`

# Config

```
./build/feedgomail -from <from-mail> -to <to-mail> -host <smtp-host> -port <smtp-port> -password <smtp-password> -feeds <feed-a> -feeds <feed-b> -feeds <feed-c>
```
