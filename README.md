# A simple URL shortener for Go

# Introdation

This's only a **simple go-URL-shortener** project.

# How to install

```bash
go get github.com/hy-shine/go-simple-url-shortener
```

# How to use
## api

you should make sure you machine that it had install `docker-compose` .

```bash
cd /path/to/the-project
docker-compose up -d .
```

Visit the `api` address `localhost:9100` by browser.

## client
```
cd /path/to/the-project/client
go build client.go
```

Run the application with `./client`

```bash
./client
Usage: ./client [OPTION] [URL]
	for example:
		./client -g https://www.google.com

	-h --help	show the client help info
	-g		generate a shorter url by the client
	-v --version	show the client version info
```