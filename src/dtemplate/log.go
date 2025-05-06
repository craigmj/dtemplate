package dtemplate

import (
	`log`
	`os`
)

var ilog = log.New(os.Stderr, `I `, log.LstdFlags|log.Lshortfile)
var elog = log.New(os.Stderr, `E `, log.LstdFlags|log.Lshortfile)