package usecase

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcPos(t *testing.T) {
	// 一番使うことが多いパターン
	// w, h = 144
	assert.Equal(t, image.Pt(0, 0), calcPos(0, 2, 4, 144, 144, 8))
	assert.Equal(t, image.Pt(144, 0), calcPos(1, 2, 4, 144, 144, 8))
	assert.Equal(t, image.Pt(432, 0), calcPos(3, 2, 4, 144, 144, 8))
	assert.Equal(t, image.Pt(0, 144), calcPos(4, 2, 4, 144, 144, 8))
	assert.Equal(t, image.Pt(432, 144), calcPos(7, 2, 4, 144, 144, 8))
	assert.Equal(t, image.Pt(0, 0), calcPos(8, 2, 4, 144, 144, 8))
	assert.Equal(t, image.Pt(144, 0), calcPos(9, 2, 4, 144, 144, 8))
	// w, h = 96
	assert.Equal(t, image.Pt(288, 96), calcPos(7, 2, 4, 96, 96, 8))
	// 異常系
	assert.Equal(t, image.Pt(0, 0), calcPos(1, 1, 1, 144, 144, 1))
}
