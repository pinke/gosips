package header

import (
	"gitee.com/xppinke/gosips/core"
)

/**
*  Content Encoding SIP header List. Keeps a list of ContentEncoding headers
 */
type ContentEncodingList struct {
	SIPHeaderList
}

/** Default constructor.
 */
func NewContentEncodingList() *ContentEncodingList {
	this := &ContentEncodingList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_CONTENT_ENCODING)
	return this
}
