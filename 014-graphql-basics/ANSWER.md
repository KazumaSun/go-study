# 014 正解例（課題別）

## 課題1: 最小Schema

```graphql
type Todo {
  id: ID!
  text: String!
}

type Query {
  todos: [Todo!]!
}

type Mutation {
  createTodo(text: String!): Todo!
}
```

## 課題2: Resolver最小例（擬似コード）

```go
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{{ID: "1", Text: "study graphql"}}, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, text string) (*model.Todo, error) {
	if strings.TrimSpace(text) == "" {
		return nil, fmt.Errorf("text is required")
	}
	return &model.Todo{ID: "2", Text: text}, nil
}
```

## 課題3: DataLoader導入の方向

```go
// resolver内で都度DBアクセスするのではなく、IDを収集して一括取得する
user, err := loaders.For(ctx).UserByID.Load(todo.UserID)
if err != nil {
	return nil, err
}
```
