package fetch

type PodLog struct {
	PodName   string
	Namespace string
	LogPath   string
	FetchDest string
	Container string
}
