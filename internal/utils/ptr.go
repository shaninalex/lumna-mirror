// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package utils

// Pointer return pointer of a provided value
func Pointer[T any](v T) *T {
	return &v
}
