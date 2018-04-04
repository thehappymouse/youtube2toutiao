package downloader

// 一个可以获取到的视频信息
type VideoFile struct {
	FilePath, Title, Desc, ID, Md5 string
	FileSize                       int64
}
