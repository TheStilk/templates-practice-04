package letter

import (
	"fmt"
)

type Letter struct{}

func (l *Letter) Open() {
	fmt.Println("Opening Letter document...")
}
