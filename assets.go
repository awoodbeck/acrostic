// +build dev

package acrostic

import "net/http"

var Assets http.FileSystem = http.Dir("assets")
