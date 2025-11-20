package entities

// ============================
// DOCUMENTOS
// ============================

type DocumentExtension string

const (
	DocPDF  DocumentExtension = ".pdf"
	DocDOC  DocumentExtension = ".doc"
	DocDOCX DocumentExtension = ".docx"
	DocXLS  DocumentExtension = ".xls"
	DocXLSX DocumentExtension = ".xlsx"
	DocPPT  DocumentExtension = ".ppt"
	DocPPTX DocumentExtension = ".pptx"
	DocTXT  DocumentExtension = ".txt"
	DocCSV  DocumentExtension = ".csv"
	DocODT  DocumentExtension = ".odt"
	DocODS  DocumentExtension = ".ods"
	DocODP  DocumentExtension = ".odp"
	DocEPUB DocumentExtension = ".epub"
)

var DocumentExtensions = []DocumentExtension{
	DocPDF, DocDOC, DocDOCX, DocXLS, DocXLSX, DocPPT, DocPPTX,
	DocTXT, DocCSV, DocODT, DocODS, DocODP, DocEPUB,
}

// ============================
// VÍDEOS
// ============================

type VideoExtension string

const (
	VideoMP4  VideoExtension = ".mp4"
	VideoMKV  VideoExtension = ".mkv"
	VideoAVI  VideoExtension = ".avi"
	VideoMOV  VideoExtension = ".mov"
	VideoWMV  VideoExtension = ".wmv"
	VideoFLV  VideoExtension = ".flv"
	VideoWEBM VideoExtension = ".webm"
)

var VideoExtensions = []VideoExtension{
	VideoMP4, VideoMKV, VideoAVI, VideoMOV,
	VideoWMV, VideoFLV, VideoWEBM,
}

// ============================
// MÚSICAS
// ============================

type MusicExtension string

const (
	MusicMP3  MusicExtension = ".mp3"
	MusicWAV  MusicExtension = ".wav"
	MusicFLAC MusicExtension = ".flac"
	MusicAAC  MusicExtension = ".aac"
	MusicOGG  MusicExtension = ".ogg"
	MusicM4A  MusicExtension = ".m4a"
)

var MusicExtensions = []MusicExtension{
	MusicMP3, MusicWAV, MusicFLAC, MusicAAC,
	MusicOGG, MusicM4A,
}

// ============================
// IMAGENS
// ============================

type ImageExtension string

const (
	ImageJPG  ImageExtension = ".jpg"
	ImageJPEG ImageExtension = ".jpeg"
	ImagePNG  ImageExtension = ".png"
	ImageGIF  ImageExtension = ".gif"
	ImageBMP  ImageExtension = ".bmp"
	ImageSVG  ImageExtension = ".svg"
	ImageWEBP ImageExtension = ".webp"
	ImageTIFF ImageExtension = ".tiff"
)

var ImageExtensions = []ImageExtension{
	ImageJPG, ImageJPEG, ImagePNG, ImageGIF,
	ImageBMP, ImageSVG, ImageWEBP, ImageTIFF,
}
