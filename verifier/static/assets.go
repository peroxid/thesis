// +build dev

package static

import "net/http"

var Assets http.FileSystem = http.Dir("../static/files")