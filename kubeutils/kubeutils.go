package kubeutils

import (
	"os"

	"gopkg.in/yaml.v2"
)

func ReadManifest(manifestFile string) (map[string]map[string]string, error) {
	yfile, errReader := os.ReadFile(manifestFile)
	if errReader != nil {
		return nil, errReader
	}
	secrets := make(map[string]map[string]string)
	var manifest ManifestSchema
	errYaml := yaml.Unmarshal(yfile, &manifest)
	if errYaml != nil {
		return nil, errYaml
	}
	for _, data := range manifest.Spec.Data {
		key := data.RemoteRef.Key
		property := data.RemoteRef.Property
		secretKey := data.SecretKey
		if _, ok := secrets[key]; !ok {
			secrets[key] = make(map[string]string)
		}
		secrets[key][secretKey] = property
	}
	return secrets, nil
}
