package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestIf1DollarIs82Rs(t *testing.T){
	want:= false
	
	got:= CheckInrToUsdConRate(82)

    assert.Equal(t,want,got)
}

func TestIf1DollarIs82point47Rs(t *testing.T){
	want:= true
	
	got:= CheckInrToUsdConRate(82.47)

    assert.Equal(t,want,got)
}