package header

import (
	"bytes"
	"errors"
	"github.com/pinke/gosips/core"
	"github.com/pinke/gosips/sip/address"
	"strings"
)

/**
 *From SIP Header.
 */
type From struct {
	AddressParameters
}

/** Default constructor
 */
func NewFrom() *From {
	this := &From{}
	this.AddressParameters.super(core.SIPHeaderNames_FROM)
	return this
}

func (this *From) super(name string) {
	this.AddressParameters.super(name)
}

/** Generate a FROM header from a TO header
 */
func (this *From) CloneTo(to *To) {
	this.super(core.SIPHeaderNames_FROM)
	this.addr = to.addr
	this.parameters = to.parameters
}

/**
 * Compare two To headers for equality.
 * @param otherHeader Object to set
 * @return true if the two headers are the same.
 */
/*public boolean equals(Object otherHeader) {
    try {
        if (!otherHeader.getClass().equals(this.getClass())){
            return false;
        }

        From otherTo = (From) otherHeader;
        if (! otherTo.getAddress().equals(address)) {
            return false;
        }
        return true;
        // exitpoint = 3;
        // return parms.equals(otherTo.parms);
    } finally {
        // System.out.println("equals " + retval + exitpoint);
    }
}*/

/**
 * Encode the header into a String.
 *
 * @return String
 */
func (this *From) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON + core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode the header content into a String.
 *
 * @return String
 */
func (this *From) EncodeBody() string {
	var retval bytes.Buffer
	addr, _ := this.addr.(*address.AddressImpl)
	if addr.GetAddressType() == address.ADDRESS_SPEC {
		retval.WriteString(core.SIPSeparatorNames_LESS_THAN)
	}
	retval.WriteString(this.addr.String())
	if addr.GetAddressType() == address.ADDRESS_SPEC {
		retval.WriteString(core.SIPSeparatorNames_GREATER_THAN)
	}
	if this.parameters.Len() > 0 {
		retval.WriteString(core.SIPSeparatorNames_SEMICOLON)
		retval.WriteString(this.parameters.String())
	}
	return retval.String()
}

/**
 * Conveniance accessor function to get the hostPort field from the address.
 * Warning -- this assumes that the embedded URI is a SipURL.
 *
 * @return hostport field
 */
func (this *From) GetHostPort() (*core.HostPort, error) {
	if this.addr == nil {
		return nil, errors.New("Address is nil")
	}
	addr, _ := this.addr.(*address.AddressImpl)
	return addr.GetHostPort()
}

/**
 * Get the display name from the address.
 * @return Display name
 */
func (this *From) GetDisplayName() string {
	return this.addr.GetDisplayName()
}

/**
 * Get the tag parameter from the address parm list.
 * @return tag field
 */
func (this *From) GetTag() string {
	if this.parameters == nil {
		return ""
	}
	return this.GetParameter(ParameterNames_TAG)
}

/** Boolean function
 * @return true if the Tag exist
 */
func (this *From) HasTag() bool {
	return this.HasParameter(ParameterNames_TAG)
}

/** remove Tag member
 */
func (this *From) RemoveTag() {
	this.parameters.Delete(ParameterNames_TAG)
}

/**
 * Set the address member
 * @param address Address to set
 */
func (this *From) SetAddress(addr address.Address) {
	this.addr = addr
}

/**
 * Set the tag member
 * @param t tag to set. From tags are mandatory.
 */
func (this *From) SetTag(t string) (ParseException error) {
	if t == "" {
		return errors.New("NullPointerException: null tag ")
	} else if strings.TrimSpace(t) == "" {
		return errors.New("ParseException: bad tag")
	}
	/*if (LogWrite.needsLogging) {
		LogWrite.logMessage("From:setTag " + t);
		LogWrite.logStackTrace();
	}*/
	this.SetParameter(ParameterNames_TAG, t)
	return nil
}

/** Get the user@host port string.
 */
func (this *From) GetUserAtHostPort() string {
	addr, _ := this.addr.(*address.AddressImpl)
	return addr.GetUserAtHostPort()
}
