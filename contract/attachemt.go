package contract

type AttachmentSection int

const (
	Dialog AttachmentSection = iota
	Album
)

type Attachemt struct {
	Section AttachmentSection
	Type    string
	Name    string
	Url     string
	IsLink  bool
}

var AttachmentSectionNames = map[AttachmentSection]string{
	Dialog: "Dialog",
	Album:  "Album",
}
