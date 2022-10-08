package cursor

import (
	"fmt"
	"runtime"
	"strings"
)

// Area displays content which can be updated on the fly.
// You can use this to create live output, charts, dropdowns, etc.
type Area struct {
	height int
}

// NewArea returns a new Area.
func NewArea() Area {
	return Area{}
}

// Clear clears the content of the Area.
func (area *Area) Clear() {
	Bottom()
	if area.height > 0 {
		ClearLinesUp(area.height)
	}
}

// Update overwrites the content of the Area.
func (area *Area) Update(content string) {
	area.UpdateWithFixed(content, false, 0)
}

// UpdateWithFixed overwrites the content of the Area.
func (area *Area) UpdateWithFixed(content string, updateHeight bool, fixed int) {
	area.Clear()
	lines := strings.Split(content, "\n")

	fmt.Println(strings.Repeat("\n", len(lines)-1)) // This appends space if the terminal is at the bottom
	n := len(lines) + fixed
	Up(n)

	if runtime.GOOS == "windows" {
		for _, line := range lines {
			fmt.Print(line)
			StartOfLineDown(1)
		}
	} else {
		for _, line := range lines {
			fmt.Println(line)
		}
	}
	height = 0

	if updateHeight {
		area.height = n
	} else {
		area.height = n - fixed
	}
}
