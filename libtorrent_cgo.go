//go:build !android
// +build !android

package libtorrent

// #cgo linux pkg-config: libtorrent-rasterbar openssl
// #cgo windows pkg-config: --static libtorrent-rasterbar openssl
// #cgo darwin CXXFLAGS: -I/usr/local/include -I/usr/local/include/libtorrent
// #cgo darwin LDFLAGS: -lm
// #cgo linux CFLAGS: -Wno-deprecated-declarations -Wno-strict-aliasing
// #cgo linux CXXFLAGS: -Wno-deprecated-declarations -Wno-strict-aliasing
//// #cgo linux CXXFLAGS: -I/usr/include/libtorrent -I/usr/include -Wno-deprecated-declarations
//// #cgo linux LDFLAGS: -lm -ldl -lrt
// #cgo windows CXXFLAGS: -DSWIGWIN -D_WIN32_WINNT=0x0501 -D__MINGW32__ -I/usr/i686-w64-mingw32/usr/include -I/usr/i686-w64-mingw32/usr/include/libtorrent -Wno-deprecated-declarations
// #cgo windows LDFLAGS: -L/usr/i686-w64-mingw32/usr/lib
import "C"
