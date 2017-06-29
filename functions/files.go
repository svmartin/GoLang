package files

import (
	"fmt"
	"os"
)

func Size(fileName string) (int64, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		// RETURN THE VALUES 0 AND err
		return 0, fmt.Errorf(err)
	}
	// RETURN THE VALUES fileInfo.Size() AND nil
	return fileInfo.Size(), err
}
