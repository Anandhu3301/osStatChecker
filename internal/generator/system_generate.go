package generator

import (
	"fmt"

	"github.com/Anandhu3301/osStatChecker/internal/info"
	"github.com/Anandhu3301/osStatChecker/models"
)

func DataGenerator() {
	sv := &models.SystemValues{}

	generateArray := []models.SystemGenerator{
		info.MemoryValues(),
		info.OsData(),
	}

	for _, gen := range generateArray {
		gen(sv)
	}

	fmt.Printf("%+v\n", sv)

}
