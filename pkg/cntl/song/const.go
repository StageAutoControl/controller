package song

import "errors"

// Errors which can be thrown during validation
var (
	ErrSongMustHaveABarChangeAtFrame0 = errors.New("song needs to have a bar change at frame 0")
)
