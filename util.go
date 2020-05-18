package libduitku

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

const (
	// CreditCard is a default value of credit card (VISA/Master) from duitku.
	CreditCard string = "VC"
	// BCAKlikPay is a default value of BCA KlikPay from duitku.
	BCAKlikPay string = "BK"
	// MandiriVirtualAccount is a default value of Mandiri Virtual Account from duitku.
	MandiriVirtualAccount string = "M1"
	// PermataBankVirtualAccount is a default value of Permata Bank Virtual Account from duitku.
	PermataBankVirtualAccount string = "BT"
	// ATMBersama is a default value of ATM Bersama from duitku.
	ATMBersama string = "A1"
	// CIMBNiagaVirtualAccount is a default value of CIMB Niaga Virtual Account from duitku.
	CIMBNiagaVirtualAccount string = "B1"
	// BNIVirtualAccount is a default value of BNI Virtual Account from duitku.
	BNIVirtualAccount string = "I1"
	// MaybankVirtualAccount is a default value of Maybank Virtual Account from duitku.
	MaybankVirtualAccount string = "VA"
	// Ritel is a default value of Ritel from duitku.
	Ritel string = "FT"
	// OVO is a default value of OVO from duitku.
	OVO string = "OV"
)

// HashToMD5 is used to generate signature
func (core *CoreDuitku) hashToMD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func (core *CoreDuitku) isSanbox(prod, sandbox string) string {
	if strings.Contains(core.Client.Host, "sandbox.duitku") {
		return sandbox
	}
	return prod
}

func (core *CoreDuitku) hashToSHA256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (core *CoreDuitku) makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
