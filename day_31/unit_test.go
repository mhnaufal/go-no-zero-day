package day_31

import "testing"

var (
	kubus        Kubus   = Kubus{4}
	volumeAwal   float64 = 64
	luasAwal     float64 = 96
	kelilingAwal float64 = 48
)

func TestHitungVolume(t *testing.T) {
	// show up the log
	t.Logf("Volume : %.2f", kubus.Volume())

	if kubus.Volume() != volumeAwal {
		t.Errorf("INCORRECT -- It should be %.2f", volumeAwal)
	}

}
