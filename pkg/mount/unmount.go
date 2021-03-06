package mount

import "github.com/u-root/u-root/pkg/mount"

// Unmount is a wrapper around u-root unmount for convenience
func Unmount(path string, force, lazy bool) error {
	return mount.Unmount(path, force, lazy)
}
