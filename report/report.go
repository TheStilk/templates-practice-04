package report

import (
	"fmt"
)

type Report struct{}

func (r *Report) Open() {
	fmt.Println("Opening Report document...")
}
