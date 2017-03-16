# The Component world


## 1. The Hello module

This module is a component package, 
to use it in your project:

1. you should add `liuggio/hello:v1` in your require section of your `component.json`
2. run `component install`
3. use in your code like a library:

```
// your_script.go
package main

import (
    "fmt"
    "liuggio/hello"
)

func main() {
	hi, err := hello()
    fmt.Print(hi, err)
}
```
and finally run your script "go run my_script.go"

## 2. The Component tool

Component is a tool that helps you to handle and create web modules.

Is like a RPC but the good one with RESTFul apis.

Is like serverless but with no lockin with the infrastructure.

Is like npm/bundler/composer but for microservices and in any language.

*Component is an easy way to connect with versioned web dependencies.*

## `component create [elixir|golang|php|ruby] hello ...`

the skeleton creates

- A. An almost empty component.json
- B. The Readme.md files
- C. A functional test file hello_test.go
- D. The function file hello.go

## `component install`

1. The component creates the client api code,
so you can import the dependencies in your code with a simple “import component-hello“.

2. the component in local environment run the dependencies and sub-dependencies in separated web servers.