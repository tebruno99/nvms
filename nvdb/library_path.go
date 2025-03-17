package nvdb

type LibraryPath struct {
	ID        int64  `json:"id"`
	Path      string `json:"path"`
	LibraryID int64  `json:"libraryId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
