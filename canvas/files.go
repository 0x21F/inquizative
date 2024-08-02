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
	Id               int
	Uuid             string
	Folder_Id        int
	Display_Name     string
	Filename         string
	Content_Type     string
	Created_At       time.Time
	Updated_At       time.Time
	Unlock_At        time.Time
	Locked           bool
	Hidden           bool
	Lock_At          time.Time
	Hidden_For_User  bool
	Visibility       Visibility
	Thumbnail_Url    string
	Modified_At      time.Time
	Mime_Class       string
	Media_Entry_Id   string
	Locked_For_User  bool
	Lock_Info        string
	Lock_Explanation string
	Preview_Url      string
}

type CanvasFolder struct {
	Context_Type     string
	Context_Id       int
	Files_Count      int
	Position         int
	Updated_At       time.Time
	Folders_Url      string
	Files_Url        string
	Full_Name        string
	Lock_At          time.Time
	Id               int
	Folders_Count    int
	Name             string
	Parent_Folder_Id int
	Created_At       time.Time
	Unlock_At        time.Time
	Hidden           bool
	Hidden_For_User  bool
	Locked           bool
	Locked_For_User  bool
	For_Submissions  bool
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
	legal_copyright string
	// Justification for using the file in a Canvas course. Valid values are
	// 'own_copyright', 'public_domain', 'used_by_permission', 'fair_use',
	// 'creative_commons'
	use_justification Justification
	// License identifier for the file.
	license string
	// Readable license name
	license_name string
	// Explanation of the action performed
	message string
	// List of ids of files that were updated
	file_ids []int
}

type License struct {
	Id, Name, Url string
}

func listFiles(url string, options ...string) ([]CanvasFile, error) {
	var res []CanvasFile

	return res, nil
}
