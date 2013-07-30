package util

import "os"

const MODE_DEV = "dev"

// Current runing mode.
func GetMode() string {
  if len(os.Args) < 2 {
    return "dev"
  } else {
    return os.Args[1]
  }
}
