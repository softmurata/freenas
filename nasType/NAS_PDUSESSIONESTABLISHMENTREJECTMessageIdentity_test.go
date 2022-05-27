package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	nas "github.com/softmurata/freenas"
	"github.com/softmurata/freenas/nasType"
)

func TestNasTypeNewPDUSESSIONESTABLISHMENTREJECTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTREJECTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionEstablishmentReject, nas.MsgTypePDUSessionEstablishmentReject},
}

func TestNasTypeGetSetPDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTREJECTMessageIdentity()
	for _, table := range nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
