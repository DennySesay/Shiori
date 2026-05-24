package registry

type Entry struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Editor string `json:"editor"`
}
