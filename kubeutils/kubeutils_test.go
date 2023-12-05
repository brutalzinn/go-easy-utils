package kubeutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadManifest(t *testing.T) {
	manifestFile := "../artefacts/test.yaml"
	manifest, _ := ReadManifest(manifestFile)
	assert.NotEmpty(t, manifest, "empty list?? something goes wrong..")
}
