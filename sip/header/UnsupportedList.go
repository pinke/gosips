package header

import "github.com/pinke/gosips/core"

/**
*   List of Unsupported headers.
 */
type UnsupportedList struct {
	SIPHeaderList
}

/** Default Constructor
 */
func NewUnsupportedList() *UnsupportedList {
	this := &UnsupportedList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_UNSUPPORTED)
	return this
}
