package grpc

import (
	context "context"
	"fmt"

	"github.com/exepirit/todo-app/internal/services/todolist"
	"github.com/exepirit/todo-app/proto"
	"github.com/google/uuid"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	domain "github.com/exepirit/todo-app/internal/domain/todolist"
)

type server struct {
	proto.UnimplementedTodoListsServer
	service todolist.IService
}

func (s *server) GetUserLists(ctx context.Context, request *proto.GetListsRequest) (*proto.TodoListArray, error) {
	userId, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to decode user ID: %w", err)
	}

	lists, err := s.service.GetUserLists(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get users lists: %w", err)
	}

	response := &proto.TodoListArray{
		Lists: make([]*proto.TodoList, len(lists)),
	}
	for i, todoList := range lists {
		response.Lists[i] = &proto.TodoList{
			Id: todoList.ID().String(),
			Owner: &proto.User{
				Id:   todoList.User().ID.String(),
				Name: todoList.User().Name,
			},
			Items: make([]*proto.TodoItem, len(todoList.Items())),
		}
		for i, item := range todoList.Items() {
			response.Lists[i].Items[i] = &proto.TodoItem{
				Text: item.Text,
			}
		}

	}
	return response, nil
}

func (s *server) Create(ctx context.Context, request *proto.CreateRequest) (*proto.TodoList, error) {
	userId, err := uuid.Parse(request.Owner.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to decode user ID: %w", err)
	}
	owner := domain.User{
		ID:   userId,
		Name: request.Owner.Name,
	}

	result, err := s.service.Create(ctx, owner)
	if err != nil {
		return nil, fmt.Errorf("failed to create new list: %w", err)
	}
	return &proto.TodoList{
		Id: result.ID().String(),
		Owner: &proto.User{
			Id:   result.User().ID.String(),
			Name: result.User().Name,
		},
		Items: make([]*proto.TodoItem, 0),
	}, nil
}

func (s *server) PutItem(ctx context.Context, request *proto.PutItemRequest) (*emptypb.Empty, error) {
	item := &domain.TodoItem{
		Text: request.Item.Text,
	}
	listId, err := uuid.Parse(request.ListId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse list ID: %w", err)
	}

	return &emptypb.Empty{}, s.service.PutItem(ctx, listId, item)
}

func (s *server) mustEmbedUnimplementedTodoListsServer() {
	panic("not implemented") // TODO: Implement
}
