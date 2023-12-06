package kubeutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadExternalSecret(t *testing.T) {
	manifestFile := "../artefacts/test.yaml"
	manifest, _ := ReadExternalSecret(manifestFile)
	assert.NotEmpty(t, manifest, "empty list?? something goes wrong..")
}
