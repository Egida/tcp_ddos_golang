package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/ypapax/tcp_ddos_golang/common"
	"testing"
)

func TestTcpServe(t *testing.T) {
	r := require.New(t)
	h, err := common.HashcashObjFromEnv()
	r.NoError(err)
	stamp, err := h.Mint("test")
	r.NoError(err)
	type testCase struct {
		stamp string
		expSuccess bool
	}
	cases := []testCase{
		{stamp: stamp, expSuccess: true},
		{stamp: stamp, expSuccess: false}, // same stamp two times is not valid
		{stamp: "some random stamp", expSuccess: false},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			ri := require.New(t)
			a, errR := common.ReqWisdom(fmt.Sprintf("localhost:%+v", testPort), c.stamp)
			ri.NoError(errR)
			t.Logf("a: %+v", a)
			ri.NotEmpty(a)
			if !c.expSuccess {
				ri.Contains(a, "is not verified")
			}
		})
	}
}
