package fs

import (
	"fmt"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("Failed read file %s: is not exist", path)
	} else if err != nil {
		return nil, fmt.Errorf("Failed read file %s: permmition denied", path)
	} else {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("Failed read file %s: unexpected error", path)
		}
		return data, nil
	}
}
