package framework

import "os"

// Current runing mode.
func GetMode() string {
  if len(os.Args) < 2 {
    return "dev"
  } else {
    return os.Args[1]
  }
}
