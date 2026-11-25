package main

import (
	"bytes"
	"strings"
	"sync"
	"testing"
	"time"
)

// lockedBuffer lets several goroutines write to one buffer safely.
type lockedBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *lockedBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func TestShopRunsErrandsAtTheSameTime(t *testing.T) {
	const step = 100 * time.Millisecond
	var out lockedBuffer
	took := shop(&out, step)

	// Five errands of one step each: one after another would take 5 steps.
	// Three run in goroutines beside the two done in turn, so about 2.
	if took < 2*step || took >= 4*step {
		t.Errorf("took %v, want between %v and %v", took, 2*step, 4*step)
	}
	for _, want := range []string{"ซื้อแว่น", "ซื้อนาฬิกา", "ซื้อผลไม้", "ซื้อรถ", "นาย A ได้รับของแล้ว"} {
		if !strings.Contains(out.buf.String(), want) {
			t.Errorf("output lacks %q:\n%s", want, out.buf.String())
		}
	}
}
