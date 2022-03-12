package contract

type AttachmentType int

const (
	Dialog AttachmentType = iota
	Album
)

type Attachemt struct {
	AttachmentType AttachmentType
	Type           string
	Name           string
	Url            string
	IsLink         bool
}
