package main

import (
    "github.com/shoobyban/dashing"
    _ "./jobs"
)

func main() {
    dashing.Start()
}
