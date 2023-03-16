package auth

import (
	"github.com/amirhnajafiz/DJaaS/pkg/enum"
)

// Claims are for token value.
type Claims struct {
	Username string
	Role     enum.Role
	State    enum.State
}
