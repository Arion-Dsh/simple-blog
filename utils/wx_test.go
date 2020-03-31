package utils

import "testing"

func TestWXToken(t *testing.T) {
	NewWXToken("wxe473a50da2ee4873", "1c7d54d7ddc19be96541973d0d30dee1")
	WX.GetToken()
	WX.getQR()
}
