// +build dev

package acrostic

import "net/http"

// Assets implements file system access to the "assets" directory.
var Assets http.FileSystem = http.Dir("assets")
