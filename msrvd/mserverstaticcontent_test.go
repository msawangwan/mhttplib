package msrvd

import "testing"

const (
	PASSMARK    = "\u2713"
	FAILMARK    = "\u2717"
	LISTEN_PORT = "127.0.0.1:1337"
)

func TestSetupAndTearDown(t *testing.T) { // skip -- need to figure out how to 'tear down'
	t.Logf("executing test: spool up the server and listen on port %s %s", LISTEN_PORT, PASSMARK)
	t.Skipf("skipping test %s", FAILMARK)

	server := NewStaticContentHandler()
	server.ListenAndServeStaticContent(LISTEN_PORT)

	t.Logf("%s", PASSMARK)
}
