package data

type Tutorial struct {
	ID       int
	Title    string
	Author   string
	Comments []Comment
}

type Comment struct {
	Body string
}

func Populate() []Tutorial {
	tutorial := Tutorial{
		ID:     1,
		Title:  "House of the dragons",
		Author: "Keerthana",
		Comments: []Comment{
			Comment{Body: "First Comment"},
			Comment{Body: "Second Comment"},
		},
	}
	tutorial2 := Tutorial{
		ID:     2,
		Title:  "Game of thrones",
		Author: "Keerthana",
		Comments: []Comment{
			Comment{Body: "Second Comment"},
		},
	}

	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)
	tutorials = append(tutorials, tutorial2)

	return tutorials
}
