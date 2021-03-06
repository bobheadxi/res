# res

Package res provides handy primitives for working with JSON in Go HTTP servers
and clients via [`go-chi/render`](https://github.com/go-chi/render). It is
designed to be lightweight and easy to extend.

[![GoDoc](https://img.shields.io/badge/go.pkg.dev-reference-5272B4)](https://pkg.go.dev/go.bobheadxi.dev/res)
[![CI Status](https://dev.azure.com/bobheadxi/bobheadxi/_apis/build/status/bobheadxi.res?branchName=master)](https://dev.azure.com/bobheadxi/bobheadxi/_build/latest?definitionId=1&branchName=master)
[![Go Report Card](https://goreportcard.com/badge/go.bobheadxi.dev/res)](https://goreportcard.com/report/go.bobheadxi.dev/res)
![Sourcegraph for Repo Reference Count](https://img.shields.io/sourcegraph/rrc/github.com/bobheadxi/res.svg)

I originally wrote something similar to this in two
[UBC Launch Pad](https://www.ubclaunchpad.com/) projects that I worked on -
[Inertia](https://github.com/ubclaunchpad/inertia) and
[Pinpoint](https://github.com/ubclaunchpad/pinpoint) - and felt like it might
be useful to have it as a standalone package.

It is currently a work-in-progress - I'm hoping to continue refining the API
and add more robust tests.

## Usage

```sh
go get -u go.bobheadxi.dev/res
```

### Clientside

I implemented something similar to `res` in [Inertia](https://github.com/ubclaunchpad/inertia).
It has a client that shows how you might leverage this library:
[`inertia/client.Client`](https://github.com/ubclaunchpad/inertia/blob/master/client/client.go)

```go
import "go.bobheadxi.dev/res"

func main() {
  resp, err := http.Get(os.Getenv("URL"))
  if err != nil {
    log.Fatal(err)
  }
  var info string
  b, err := res.Unmarshal(resp.Body, res.KV{Key: "info", Value: &info})
  if err != nil {
    log.Fatal(err)
  }
  if err := b.Error(); err != nil {
    log.Fatal(err)
  }
  println(info)
}
```

### Serverside

#### OK

```go
import "go.bobheadxi.dev/res"

func Handler(w http.ResponseWriter, r *http.Request) {
  res.R(w, r, res.MsgOK("hello world!",
    "stuff", "amazing",
    "details", res.M{"world": "hello"}))
}
```

Will render something like:

```js
{
  "code": 200,
  "message": "hello world",
  "request_id": "12345",
  "body": {
    "stuff": "amazing",
    "details": {
      "world": "hello",
    }
  }
}
```

#### Error

```go
import "go.bobheadxi.dev/res"

func Handler(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    res.R(w, r, res.ErrBadRequest("failed to read request",
      "error", err,
      "details", "something"))
    return
  }
}
```

Will render something like:

```js
{
  "code": 400,
  "message": "failed to read request",
  "request_id": "12345",
  "error": "could not read body",
  "body": {
    "details": "something",
  }
}
```
