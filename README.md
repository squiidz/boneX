boneX [![GoDoc](https://godoc.org/github.com/squiidz/bonex?status.png)](http://godoc.org/github.com/go-zoo/bonex) [![Build Status](https://travis-ci.org/go-zoo/boneX.svg?branch=master)](https://travis-ci.org/go-zoo/bonex)
=======

## What is boneX ?

BoneX is a derivation of the go-zoo/bone multiplexer. The major changes are the use of a third parameters
in the handler to get the URL params (increasing speed). Also bonex.Route implement a Eval() method which is use to evaluate the URL params of a request (Check example for more about that). BoneX is just for more complex web server which need more options.

![alt tag](https://blog.adafruit.com/wp-content/uploads/2012/10/405911_282111641849431_1466890904_n.jpg)

## Speed

#### With URL Params

```
-BenchmarkBoneXMux        2000000               697 ns/op
-BenchmarkHttpRouterMux   5000000               304 ns/op
-BenchmarkZeusMux         1000000              1232 ns/op
-BenchmarkGorillaMux      1000000              2071 ns/op
-BenchmarkGorillaPatMux    500000              2182 ns/op

```

#### Without URL Params

```
-BenchmarkBoneXMux       10000000               144 ns/op
-BenchmarkHttpRouterMux  10000000               152 ns/op
-BenchmarkZeusMux         2000000               826 ns/op
-BenchmarkNetHttpMux      2000000               736 ns/op
-BenchmarkGorillaMux       300000              4396 ns/op
-BenchmarkGorillaPatMux    500000              2166 ns/op

```

 These test are just for fun, all these router are great and really efficient. 
 BoneX do not pretend to be the best router for every job. 

## Example

``` go

package main

import(
  "net/http"
  "strconv"

  "github.com/go-zoo/bonex"
)

func main () {
  mux := bonex.New()
  
  // Method takes func (rw Http.ResponseWriter, req *http.Request, args bonex.Args)
  mux.Get("/home/:id", Handler).Eval(isANumber)

  // Eval bind the params with the first provided function,
  // the second params with the second function, etc ...
  mux.Get("/profil/:id/:var", ProfilHandler).Eval(isANumber, biggerThan3)

  http.ListenAndServe(":8080", mux)
}

func Handler(rw http.ResponseWriter, req *http.Request, args bonex.Args) {
  // Get the value of the "id" parameters.
  val := args.GetValue("id")

  rw.Write([]byte(val))
}

func isANumber(str string) bool {
  if _, err := strconv.AtoI(str); err == nil {
    return true
  }
  return false
}

func biggerThan3(str string) bool {
  if len(str) > 3 {
    return true
  }
  return false
}

```
## TODO

- DOC
- More Testing
- Debugging
- Optimisation

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Write Tests!
4. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request

## License
MIT

## Links

Middleware Chaining module : [Claw](https://github.com/go-zoo/claw)
