package pas2

import (
	"strings"
	"testing"
)

func TestKnownDigest(t *testing.T) {
	got, err := Digest(strings.NewReader("abc"))
	if err != nil || got != "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad" {
		t.Fatal(got, err)
	}
}
func TestVerify(t *testing.T) {
	digest, _ := Digest(strings.NewReader("abc"))
	ok, e := Verify(strings.NewReader("abc"), strings.ToUpper(digest))
	if e != nil || !ok {
		t.Fatal(ok, e)
	}
	ok, _ = Verify(strings.NewReader("def"), digest)
	if ok {
		t.Fatal("mismatch")
	}
}
func TestInvalid(t *testing.T) {
	if _, err := Verify(strings.NewReader(""), "bad"); err == nil {
		t.Fatal("invalid expected")
	}
}
