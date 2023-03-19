package filesystem

import (
	"os"
)

type System interface {
	GetFile(name string) (os.File, error)
	SaveFile(name string, file os.File) error
	DeleteFile(name string) error
}
