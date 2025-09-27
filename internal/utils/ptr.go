// Copyright © 2025 Lumna. All rights reserved.

package utils

// Pointer return pointer of a provided value
func Pointer[T any](v T) *T {
	return &v
}
