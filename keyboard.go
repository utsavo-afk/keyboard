package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	buffer, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error readString: %#v", err)
	}
	buffer = strings.TrimSpace(buffer)
	input, err := strconv.ParseFloat(buffer, 64)
	if err != nil {
		return 0, fmt.Errorf("error parseFloat: %#v", err)
	}
	return input, nil
}
