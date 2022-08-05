package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, kubus.Volume(), volumeSeharusnya, "Perhitungan volume salah!")
	/*
		 	t.Logf("Volume\t: %.2f", kubus.Volume())

			if kubus.Volume() != volumeSeharusnya {
				t.Errorf("SALAH! harusnya %.2f", volumeSeharusnya)
			}
	*/
}

func TestHitungLuas(t *testing.T) {
	assert.Equal(t, kubus.Luas(), luasSeharusnya, "Perhitungan luas salah!")
	/*
		t.Logf("Luas\t: %.2f", kubus.Luas())

		if kubus.Luas() != luasSeharusnya {
			t.Errorf("SALAH! harusnya %.2f", luasSeharusnya)
		}
	*/
}

func TestHitungKeliling(t *testing.T) {
	assert.Equal(t, kubus.Keliling(), kelilingSeharusnya, "Perhitungan keliling salah!")
	/* t.Logf("Keliling\t: %.2f", kubus.Keliling())

	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", kelilingSeharusnya)
	} */
}

/* func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
} */
