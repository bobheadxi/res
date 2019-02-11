# res

Package res provides handy primitives for working with JSON in Go HTTP servers
via [`go-chi/render`](https://github.com/go-chi/render). It is designed to be
lightweight and easy to extend.

I originally wrote something similar to this in two
[UBC Launch Pad](https://www.ubclaunchpad.com/) projects that I worked on -
[Inertia](https://github.com/ubclaunchpad/inertia) and
[Pinpoint](https://github.com/ubclaunchpad/pinpoint) - and felt like it might
be useful to have it as a standalone package.

It is currently a work-in-progress - I'm hoping to continue refining the API
and add more robust tests.

## Usage

```sh
go get -u github.com/bobheadxi/res
```

### Clientside

```go
import "github.com/bobheadxi/res"

func main() {
  resp, err := http.Get(os.Getenv("URL"))
  if err != nil {
    log.Fatal(err)
  }
  var info string
  b, err := res.Unmarshal(resp.Body, api.KV{Key: "info", Value: &info})
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

```go
import "github.com/bobheadxi/res"

func Handler(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    render.Render(w, r, res.ErrBadRequest("failed to read request",
      "error", err))
    return
  }

  render.Render(w, r, res.MsgOK("hello world!",
    "details", map[string]string{"world": "hello"}))
}
```
