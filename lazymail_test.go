package eazymail

import (
	"fmt"
	"testing"
	"time"
)

// Sending 3 Message in 1 Second and have time so Send 1 Email
func TestQueueSendInBackground(t *testing.T) {
	t.Parallel()
	want := make([]string, 0)
	lm := NewLazymail("smtp.foo.de", "user@foo.de", "secret", lazyUnitTest(800, &want))
	lm.Send("m4rc0 k1tt3l", "marco@foo.de", "Topic", "Email Text")
	lm.Send("m4rc0 k1tt3l", "marco@foo.de", "Topic", "Email Text")
	lm.Send("m4rc0 k1tt3l", "marco@foo.de", "Topic", "Email Text")
	time.Sleep(1000 * time.Millisecond)
	if len(want) != 1 {
		t.Fatalf("Want must have one sended message.")
	}
}

// Sending 3 Message in 1 Second and all got send
func TestQueuedContentSendContent(t *testing.T) {
	t.Parallel()
	want := make([]string, 0)
	lm := NewLazymail("smtp.foo.de", "user@foo.de", "secret", lazyUnitTest(50, &want))
	lm.Send("m4rc0 k1tt3l", "marco@foo.de", "Topic", "Email Text")
	lm.Send("m4rc0 k1tt3l", "franci@foo.de", "Topic2", "Email Text")
	lm.Send("m4rc0 k1tt3l", "marco@foo.de", "Topic3", "Email Text")
	time.Sleep(1000 * time.Millisecond)
	if len(want) != 3 {
		t.Fatalf("Want must have three sended message.")
	}

	m := fmt.Sprintf("Empfänger: %s Betreff: %s Inhalt: %s", "marco@foo.de", "Topic", "Email Text")
	msg := want[0]
	if m != msg {
		t.Fatalf("Email Text should be: %s \n\n but is %s", m, msg)
	}

	m = fmt.Sprintf("Empfänger: %s Betreff: %s Inhalt: %s", "franci@foo.de", "Topic2", "Email Text")
	msg = want[1]
	if m != msg {
		t.Fatalf("Email Text should be: %s \n\n but is %s", m, msg)
	}
}
