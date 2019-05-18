package probe

import "testing"

func TestProbe(t *testing.T) {
	if err := Create(); err != nil {
		t.Errorf("should creates a live file to liveness probe %s", err)
	}
	if !Exists() {
		t.Errorf("should exists a live file to liveness probe")
	}

	if err := Remove(); err != nil {
		t.Errorf("should removes a live file to liveness probe %s", err)
	}

	if Exists() {
		t.Errorf("after remove the file, it should not exists a live file to liveness probe")
	}
}
