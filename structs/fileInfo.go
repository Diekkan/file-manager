package structs

type FileInfo struct {
	Name      string `json: "Name"`
	Extension string `json: "Extension"`
	Size      int64  `json: "Size"`
	IsDir     bool   `json: "IsDir"`
}
