package panic

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPanic(t *testing.T) {
	t.Parallel()

	sleep := time.Duration(rand.Intn(1000))*time.Millisecond + 100*time.Millisecond
	t.Logf("Panic:Sleeping for %s", sleep)
	time.Sleep(sleep)

	panic(fmt.Sprintf("I slept for %s and panicked", sleep))
}

func TestPass(t *testing.T) {
	t.Parallel()

	for i := range 10 {
		t.Run(fmt.Sprintf("pass-%d", i), func(t *testing.T) {
			// Sleep to make sure the panic happens in the middle of test executions
			sleep := time.Duration(rand.Intn(100)) * time.Millisecond
			t.Logf("Pass: Sleeping for %s", sleep)
			time.Sleep(sleep)
			t.Log("I pass")
		})
	}
}

func TestFail(t *testing.T) {
	t.Parallel()

	for i := range 10 {
		t.Run(fmt.Sprintf("fail-%d", i), func(t *testing.T) {
			// Sleep to make sure the panic happens in the middle of test executions
			sleep := time.Duration(rand.Intn(100)) * time.Millisecond
			t.Logf("Fail: Sleeping for %s", sleep)
			time.Sleep(sleep)
			t.Log("I fail")
			t.FailNow()
		})
	}
}
