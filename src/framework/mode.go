package framework

import "os"

// GetMode get current runing mode.
func GetMode() (mode string) {
  if len(os.Args) < 2 {
    mode = "dev"
  } else {
    mode = os.Args[1]
  }
  return
}
