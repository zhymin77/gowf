package framework

import (
  "strings"
)

// Helper template operation helper.
type Helper string

var helper *Helper

// GetHelperInstance  find helper instance.
func GetHelperInstance() *Helper {
  if helper == nil {
    helper = new(Helper)
  }
  return helper
}

// Add addition.
func (h *Helper)Add(a, b int) int {
  return a + b
}

// Sub substraction.
func (h *Helper)Sub(a, b int) int {
  return a - b
}

// Mul multiplication.
func (h *Helper)Mul(a, b int) int {
  return a * b
}

// Div division.
func (h *Helper)Div(a, b int) int {
  return a / b
}

// Mod module.
func (h *Helper)Mod(a, b int) int {
  return a % b
}

// IsE is equals.
func (h *Helper)IsE(a, b interface{}) bool {
  return a == b
}

// IsOrE is equals.
func (h *Helper)IsOrE(a interface{}, bs ...interface{}) bool {
  for _, b := range bs {
    if a == b {
      return true
    }
  }
  return false
}

// Gtn greater than.
func (h *Helper)Gtn(a, b int) bool {
  return a > b
}

// Gte greater than or equal.
func (h *Helper)Gte(a, b int) bool {
  return a >= b
}

// Ltn less than.
func (h *Helper)Ltn(a, b int) bool {
  return a < b
}

// Lte less than or equal.
func (h *Helper)Lte(a, b int) bool {
  return a <= b
}

// NonEmpty check no empty.
func (h *Helper)NonEmpty(s string) bool {
  return len(s) > 0
}

// Contains string contains.
func (h *Helper)Contains(s, substr string) bool {
  return strings.Contains(s, substr)
}

// StrEq string equal.
func (h *Helper)StrEq(s, substr string) bool {
  return strings.EqualFold(s, substr)
}
