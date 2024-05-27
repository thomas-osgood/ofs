package common

import "os"

type FileSpec interface {
	string | *os.File
}
