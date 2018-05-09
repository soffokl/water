// +build !linux,!darwin,!windows,!openbsd,!freebsd,!netbsd,!dragonfly

package water

// PlatformSpeficParams
type PlatformSpecificParams struct {
}

func defaultPlatformSpecificParams() PlatformSpecificParams {
	return PlatformSpecificParams{}
}
