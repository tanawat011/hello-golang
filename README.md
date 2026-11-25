# hello-golang

Small Go programs from my first weeks with Go. Each folder is one idea:

| Folder | What it shows |
|---|---|
| `helloworld` | variables, a loop, functions, two return values |
| `hello` + `school` | importing a package from the same module |
| `sampledefer` | `defer` runs when the function returns |
| `samplestruc` | a struct turned into JSON |
| `gotoshop` | goroutines and a channel: five errands in the time of two (comments in Thai) |
| `sendmail` | composing an HTML e-mail and sending it over SMTP |

Originally written in 2018 while learning Go; refreshed in Nov 2026
(Go modules, current Go, tests, CI, no settings in the code).

## Quick start

Needs Go 1.26 or newer.

```sh
git clone https://github.com/tanawat011/hello-golang.git
cd hello-golang
go run ./helloworld
go run ./hello
go run ./sampledefer
go run ./samplestruc
go run ./gotoshop        # takes about 2 seconds
```

### sendmail

Every setting comes from environment variables; `.env.example` lists them
with placeholders. Without `-send` it only prints the message it would send.

```sh
cp .env.example .env      # then put your own values in .env (git-ignored)
set -a; . ./.env; set +a
go run ./sendmail         # prints the message
go run ./sendmail -send   # sends it (TLS required)
```

Gmail and most providers need an app password, not your account password.

## Tests

```sh
go vet ./...
go test -race ./...
```

CI runs gofmt, vet, build and the tests on the two supported Go releases,
and `govulncheck`.
