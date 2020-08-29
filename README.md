# feedgomail
service, which parses rss feeds and sends them to a dedicated email address

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

put a `feedgomail.json` file in the directory from which you are executing feedgomail, looking like this:
```json
{
	"From":     "<from-mail>",
	"To":       "<to-mail>",
	"Host":     "<smtp-host>",
	"Port":     1234,
	"Password": "<smtp-password",
	"Feeds":    ["<feed-a>", "<feed-b>", "<feed-c>"]
}
```
