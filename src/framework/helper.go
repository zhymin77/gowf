package framework

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

// Gte less than or equal.
func (h *Helper)Lte(a, b int) bool {
  return a <= b
}
