package canvas

import (
	"time"
)

type Visibility string

const (
	Inherit     Visibility = "inherit"
	Course                 = "course"
	Institution            = "institution"
	Public                 = "public"
)

type CanvasFile struct {
	Id              int        `json:"id"`
	Uuid            string     `json:"uuid"`
	FolderId        int        `json:"folder_id"`
	DisplayName     string     `json:"display_name"`
	Filename        string     `json:"filename"`
	ContentType     string     `json:"content_type"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	UnlockAt        time.Time  `json:"unlock_at"`
	Locked          bool       `json:"locked"`
	Hidden          bool       `json:"hidden"`
	LockAt          time.Time  `json:"lock_at"`
	HiddenForUser   bool       `json:"hidden_for_user"`
	Visibility      Visibility `json:"visibility"`
	ThumbnailUrl    string     `json:"thumbnail_url"`
	ModifiedAt      time.Time  `json:"modified_at"`
	MimeClass       string     `json:"mime_class"`
	MediaEntryId    string     `json:"media_entry_id"`
	LockedForUser   bool       `json:"locked_for_user"`
	LockInfo        string     `json:"lock_info"`
	LockExplanation string     `json:"lock_explanation"`
	PreviewUrl      string     `json:"preview_url"`
}

type CanvasFolder struct {
	ContextType    string    `json:"context_type"`
	ContextId      int       `json:"context_id"`
	FilesCount     int       `json:"files_count"`
	Position       int       `json:"position"`
	UpdatedAt      time.Time `json:"updated_at"`
	FoldersUrl     string    `json:"folders_url"`
	FilesUrl       string    `json:"files_url"`
	FullName       string    `json:"full_name"`
	LockAt         time.Time `json:"lock_at"`
	Id             int       `json:"id"`
	FoldersCount   int       `json:"folders_count"`
	Name           string    `json:"name"`
	ParentFolderId int       `json:"parent_folder_id"`
	CreatedAt      time.Time `json:"created_at"`
	UnlockAt       time.Time `json:"unlock_at"`
	Hidden         bool      `json:"hidden"`
	HiddenForUser  bool      `json:"hidden_for_user"`
	Locked         bool      `json:"locked"`
	LockedForUser  bool      `json:"locked_for_user"`
	ForSubmissions bool      `json:"for_submissions"`
}

type Justification string

const (
	Own_Copyright      Justification = "own_copyright"
	Public_Domain                    = "public_domain"
	Used_By_Permission               = "used_by_permission"
	Fair_Use                         = "fair_use"
	Creative_Commons                 = "creative_commons"
)

type UsageRights struct {
	LegalCopyright   string        `json:"legal_copyright"`
	UseJustification Justification `json:"use_justification"`
	License          string        `json:"license"`
	LicenseName      string        `json:"license_name"`
	Message          string        `json:"message"`
	FileIDs          []int         `json:"file_ids"`
}

type License struct {
	Id, Name, Url string
}

func listFiles(url string, options ...string) ([]CanvasFile, error) {
	var res []CanvasFile

	return res, nil
}
