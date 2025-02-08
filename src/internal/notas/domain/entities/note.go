package entities

type Note struct {
	ID      int
	Title   string
	Content string
}

func (n Note) Validate() any {
	panic("unimplemented")
}

func NewNote(id int, title, content string) *Note {
	return &Note{
		ID:      id,
		Title:   title,
		Content: content,
	}
}
