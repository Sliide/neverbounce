package neverbounce_test

import (
	"os"
	"testing"

	"github.com/sliide/neverbounce"
)

func Assert(t *testing.T, a bool) {
	if !a {
		t.Fail()
	}
}

func TestMain(m *testing.M) {

	neverbounce.Init(&neverbounce.NeverBounceCli{
		ApiUsername: os.Getenv("NEVERBOUNCE_USERNAME"),
		ApiPassword: os.Getenv("NEVERBOUNCE_PASSWORD"),
		TestMode:    true,
	})

	os.Exit(m.Run())
}

func TestNeverBounceGetsAccessToken(t *testing.T) {
	Assert(t, neverbounce.GetAccessToken() != "")
}

func TestNeverBounceVerifiesEmail(t *testing.T) {
	Assert(t, neverbounce.VerifyEmail("andre@valid.com").Result == 0)
	Assert(t, neverbounce.VerifyEmail("asasd@invalid.com").Result != 0)
}
