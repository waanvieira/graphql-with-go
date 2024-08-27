#### GraphQL com GOLANG

Utilizaremos a lib https://gqlgen.com/

Iniciando

Criamos um arquivo tools.go (pode ser qualquer nome, usamos tools.go como convenção) ou executamos o comando

printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

Esse arquivo irá conter as nossas depêndencias da lib

Depois de criado o tool.go vamos executar o comando

go run github.com/99designs/gqlgen init

go mod tidy

Esse comando irá criar o nosso esqueleto da aplicação

Para executar o nosso servidor GraphQL rodar o arquivo server.go

go run server.go

Por padrão irá abrir na porta :8080


##### Criando nossos Schemas (Inputs e Outputs do GraphQL)
Para criar os nossos schemas alterar e criar os arquivos em

graph/model

criamos algum arquivo nessa pasta nome_do_arquivo.graphqls nesse caso criamos

schema.graphqls

Aqui podemos verificar as opções e como montar os nossos schemas https://gqlgen.com/getting-started/

Os schemas serão os nossos "controllers" nossas informações dos inputs do GraphQL como campos, tipos de campos e seus retornos

#### Criando nossos métodos

Para criarmos nossos métodos e injetar nossas ações alteramos o arquivo graph/schema.resolvers.go nesse arquivo vamos indicar nossos métodos quando alguma ação é executada
quando for para fazer uma consulta, criar, alterar e deletar algo


