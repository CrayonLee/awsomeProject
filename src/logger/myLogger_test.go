package main

import "testing"

func TestConstLevel(t *testing.T){
	t.Logf("%v %T\n", DeubgLevel, DeubgLevel)
	t.Logf("%v %T\n", FatalLevel, FatalLevel)
}