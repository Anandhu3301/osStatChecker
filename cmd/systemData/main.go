package main

import (
	"github.com/Anandhu3301/osStatChecker/internal/generator"
	"github.com/Anandhu3301/osStatChecker/internal/info"
)

// "fmt"
// "runtime"

func main() {
    generator.DataGenerator();
    ViewInternalImplementation();
	
}

func ViewInternalImplementation() {
	info.Get_Process_List();
}