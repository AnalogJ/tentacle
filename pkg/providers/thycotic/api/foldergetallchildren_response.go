package api


type FolderGetAllChildrenResponse struct {
	FolderGetAllChildrenResult FolderGetAllChildrenResult
}


type FolderGetAllChildrenResult struct {
	Success bool
	Errors []string `xml:"Errors>string"`
	Folders []Folder `xml:"Folders>Folder"`
}

type Folder struct {
	Id int
	Name string
	TypeId int
	ParentFolderId int
}