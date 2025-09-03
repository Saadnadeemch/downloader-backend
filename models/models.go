package models

type Request struct {
	URL     string `json:"url"`
	Quality string `json:"quality"`
}

type AudioRequest struct {
	URL       string `json:"url"`
	RequestID string `json:"request_id"`
	Title     string
}

type DownloadVideoRequest struct {
	URL       string `json:"url"`
	Quality   string `json:"quality"`
	RequestID string `json:"request_id"`
	FormatID  string `json:"format_id"`
	Title     string `json:"title"`
}

type Format struct {
	FormatID   string `json:"format_id"`
	Extension  string `json:"ext"`
	Resolution string `json:"resolution"`
	Vcodec     string `json:"vcodec"`
	Acodec     string `json:"acodec"`
	Height     int    `json:"height"`
	Width      int    `json:"width"`
	Protocol   string
}

type VideoInfo struct {
	Title       string  `json:"title"`
	Thumbnail   string  `json:"thumbnail"`
	Uploader    string  `json:"uploader"`
	Views       int64   `json:"views"`
	Source      string  `json:"source"`
	DownloadURL *string `json:"downloadUrl,omitempty"`
	Description *string `json:"description,omitempty"`
	UploadDate  *string `json:"upload_date,omitempty"`
	LikeCount   *int64  `json:"like_count,omitempty"`
	VideoPage   string  `json:"url"`
}

type VideoDownloadResult struct {
	RequestID   string `json:"request_id"`
	FilePath    string `json:"file_path"`
	Title       string `json:"title"`
	FileName    string `json:"file_name"`
	DownloadURL string `json:"download_url"`
	CleanupAt   int64  `json:"cleanup_at"`
}

type StreamVideoDownloadRequest struct {
	URL string `json:"url"`
}

type PlatformInfo struct {
	Platform string
	IsValid  bool
	IsDirect string
}
