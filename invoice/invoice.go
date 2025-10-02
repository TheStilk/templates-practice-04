package invoice

import (
	"fmt"
)

type Invoice struct{}

func (i *Invoice) Open() {
	fmt.Println("Opening Invoice document...")
}
