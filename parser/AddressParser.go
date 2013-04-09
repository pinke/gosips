package parser

import (
	"strings"
	"gosip/core"
	"gosip/address"
)

/** Parser for addresses.
*/

type AddressParser struct{
 	ParserImpl
}

	

	
	func NewAddressParser(addr string) *AddressParser {
		this := &AddressParser{}
		
		this.ParserImpl.super(addr);
    	this.ParserImpl.GetLexer().SetLexerName("charLexer");
    	
    	return this;
	}
	
	func NewAddressParserFromLexer(lexer core.Lexer) *AddressParser {
		this := &AddressParser{};
		
		this.SetLexer(lexer);
		this.GetLexer().SelectLexer("charLexer");
		
		return this;
	}
	
	func (this *AddressParser) super(addr string) {
    	this.ParserImpl.super(addr);
    	this.ParserImpl.GetLexer().SetLexerName("charLexer");
    }
    
    func (this *AddressParser) superFromLexer(lexer core.Lexer) {
    	this.SetLexer(lexer);
		this.GetLexer().SelectLexer("charLexer");
    }	

	func (this *AddressParser) NameAddr() (addr *address.AddressImpl, ParseException error) {
	    //if (debug) dbg_enter("nameAddr");
	    //try {
	    var ch byte;
	    var err error;
	    
	    lexer := this.GetLexer();
		if ch, _ = lexer.LookAheadK(0); ch == '<' {
		   lexer.Match('<');
		   lexer.SelectLexer("sip_urlLexer");
		   lexer.SPorHT();
		   uriParser := NewURLParser(lexer);
		   uri := uriParser.UriReference();
		   addr = address.NewAddressImpl();
		   addr.SetAddressType(address.NAME_ADDR);
		   addr.SetURI(uri);
		   lexer.SPorHT();
		   lexer.Match('>');
		   return addr, nil;
		} else {
		   addr = address.NewAddressImpl();
		   addr.SetAddressType(address.NAME_ADDR);
		   var name string;
		   if ch, _ = lexer.LookAheadK(0); ch == '"'  {
		      name,_ = lexer.QuotedString();
		      lexer.SPorHT();
		   } else {
		      name = lexer.GetNextTokenByDelim('<');
		   }
		   addr.SetDisplayName(strings.TrimSpace(name));
		   lexer.Match('<');
		   lexer.SPorHT();
		   uriParser := NewURLParser(lexer);
		   uri := uriParser.UriReference();
		   addr = address.NewAddressImpl();
		   addr.SetAddressType(address.NAME_ADDR);
		   addr.SetURI(uri);
		   lexer.SPorHT();
		   lexer.Match('>');
		   return addr, nil;
		}
	    // } finally {
		//if (debug) dbg_leave("nameAddr");
	    //}

	}

	func (this *AddressParser) Address() (retval *address.AddressImpl, ParseException error)  {
	    //if (debug) dbg_enter("address");
	    //AddressImpl retval = null;
	    //try {
	    var ch byte;
	    var err error
	    lexer := this.GetLexer();
	 	k := 0;
		for lexer.HasMoreChars() {
		   if ch, err = lexer.LookAheadK(k); 
		    ch == '<' || 
			ch == '"' ||
			ch == ':' ||
			ch == '/' {
			break;
		   }else if ch == 0{//'\0' 
				return nil, this.CreateParseException("unexpected EOL");
		   }else{ 
		   	k++;
		   }
		}
		if ch, err = lexer.LookAheadK(k); 
			ch == '<' || 
		    ch == '"'  {
			retval,_ = this.NameAddr();
		} else if ch == ':' || 
				  ch == '/' {
			retval = address.NewAddressImpl();
			uriParser := NewURLParser(lexer);
		 	uri := uriParser.UriReference();
			retval.SetAddressType(address.ADDRESS_SPEC);
			retval.SetURI(uri);
		} else {
			return nil, this.CreateParseException("Bad address spec");
		}
		return retval, nil;
	    // } finally {
		//if (debug) dbg_leave("address");
	    // }

	}


