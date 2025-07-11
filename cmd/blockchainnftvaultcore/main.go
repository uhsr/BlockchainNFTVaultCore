// cmd/blockchainnftvaultcore/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftvaultcore/internal/blockchainnftvaultcore"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftvaultcore.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
