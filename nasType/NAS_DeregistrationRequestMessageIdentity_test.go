package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	nas "github.com/softmurata/freenas"
	"github.com/softmurata/freenas/nasType"
)

type nasTypeDeregistrationRequestMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeDeregistrationRequestMessageIdentityTable = []nasTypeDeregistrationRequestMessageIdentityData{
	{nas.MsgTypeDeregistrationRequestUETerminatedDeregistration, nas.MsgTypeDeregistrationRequestUETerminatedDeregistration},
}

func TestNasTypeNewDeregistrationRequestMessageIdentity(t *testing.T) {
	a := nasType.NewDeregistrationRequestMessageIdentity()
	assert.NotNil(t, a)
}

func TestNasTypeGetSetDeregistrationRequestMessageIdentity(t *testing.T) {
	a := nasType.NewDeregistrationRequestMessageIdentity()
	for _, table := range nasTypeDeregistrationRequestMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
