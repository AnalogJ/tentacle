package api


type DownloadFileAttachmentByItemIdResponse struct {
	DownloadFileAttachmentByItemIdResult DownloadFileAttachmentByItemIdResult
}


type DownloadFileAttachmentByItemIdResult struct {
	FileName string
	Errors []string `xml:"Errors>string"`
	FileAttachment string
}
