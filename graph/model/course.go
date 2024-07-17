package model

// Precisamos adicionar no arquivo gqlgen.yml em model o nosso modelo, para poder ser reconhecido e gerado
//
//	  Course:
//	  model:
//		-  github.com/waanvieira/graphqlgen/model.Course
type Course struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
