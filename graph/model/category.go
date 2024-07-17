package model

// Precisamos adicionar no arquivo gqlgen.yml em model o nosso modelo, para poder ser reconhecido e gerado
//
//	  Category:
//	  model:
//		-  github.com/waanvieira/graphqlgen/model.Category
type Category struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
