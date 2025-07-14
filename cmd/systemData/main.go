package main

import (
	"github.com/Anandhu3301/osStatChecker/internal/generator"
	"github.com/Anandhu3301/osStatChecker/pkg/newer"
)

// "fmt"
// "runtime"

func main() {
    generator.DataGenerator();
    newer.Latest();
}

