package 
info

import (
	"runtime"
	"github.com/Anandhu3301/osStatChecker/models"
)

func OsData() models.SystemGenerator{
	 return func(sv *models.SystemValues) {
          sv.OperatingSystem = runtime.GOOS
		  sv.Architecture = runtime.GOARCH
	 };
}