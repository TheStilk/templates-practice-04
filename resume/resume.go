package resume

import (
	"fmt"
)

type Resume struct{}

func (r *Resume) Open() {
	fmt.Println("Opening Resume document...")
}
