## Planes

Hello.

The following routes exist:

* '/airports'
* '/airports/{code}'
* '/routes'
* '/routes/{originCode}/{destinationCode}'

Each API is a GET request and returns either a 200 OK with JSON, or a 400 Bad Request.

### How do I run this?

```shell
git clone github.com/tombuildsstuff/planes
go build . && ./planes
```

Then check `http://localhost:2021`.

### What is this repo?

Wanderlust, I guess. But it's handy having an API kicking around in case you need to demo something.

### Is there a licence?

MIT, I guess.
