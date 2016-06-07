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
	})

	os.Exit(m.Run())
}

func TestNeverBounceGetsAccessToken(t *testing.T) {
	Assert(t, neverbounce.GetAccessToken() != "")
}

func TestNeverBounceVerifiesEmail(t *testing.T) {
	Assert(t, neverbounce.VerifyEmail("andre@sliideapp.com").Result == 0)
	Assert(t, neverbounce.VerifyEmail("asasd@allls.com").Result != 0)
}
