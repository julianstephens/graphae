package utils

import (
	"github.com/op/go-logging"
)

var LOG = logging.MustGetLogger("image-processor")

var format = logging.MustStringFormatter( //nolint:golint,unused
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)
