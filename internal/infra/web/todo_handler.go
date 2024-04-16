package web

import (
	"net/http"
	"todoList/internal/entity"
	"todoList/internal/infra/web/webserver"
	"todoList/internal/usecase"
	"todoList/internal/usecase/complete"
	"todoList/internal/usecase/create"
	"todoList/internal/usecase/delete"
	"todoList/internal/usecase/retrieve"
	"todoList/internal/usecase/undo"
	"todoList/internal/usecase/update"
	"todoList/utils"
)

type TodoHandler struct {
	todoRepository entity.TodoRepositoryInterface
}

func NewTodoHandler(todoRepository entity.TodoRepositoryInterface) *TodoHandler {
	return &TodoHandler{todoRepository: todoRepository}
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, _ *http.Request) {
	useCase := retrieve.NewGetTodos(h.todoRepository)
	todos, err := useCase.Execute()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, todos)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var input create.CreateTodoInput
	if err := utils.DecodeJSONBody(r.Body, &input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	useCase := create.NewCreateTodo(h.todoRepository)
	if err := useCase.Execute(input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, nil)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var input update.UpdateTodoInput
	if err := utils.DecodeJSONBody(r.Body, &input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	id := r.PathValue("id")
	input.ID = id
	useCase := update.NewUpdateTodo(h.todoRepository)
	if err := useCase.Execute(input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}

func (h *TodoHandler) Undo(w http.ResponseWriter, r *http.Request) {
	var input undo.UndoTodoInput
	id := r.PathValue("id")
	input.ID = id

	useCase := undo.NewUndoTodo(h.todoRepository)
	if err := useCase.Execute(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}

func (h *TodoHandler) Complete(w http.ResponseWriter, r *http.Request) {
	var input complete.CompleteTodoInput
	id := r.PathValue("id")
	input.ID = id
	useCase := complete.NewCompleteTodo(h.todoRepository)
	if err := useCase.Execute(input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	var input retrieve.GetTodoInput
	id := r.PathValue("id")
	input.ID = id
	useCase := retrieve.NewGetTodo(h.todoRepository)
	todo, err := useCase.Execute(input)
	if err != nil {
		if err.Error() == usecase.ErrNoResultRows.Error() {
			utils.RespondWithError(w, http.StatusNotFound, usecase.ErrTodoNotFound.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var input delete.DeleteTodoInput
	id := r.PathValue("id")
	input.ID = id
	useCase := delete.NewDeleteTodo(h.todoRepository)
	if err := useCase.Execute(input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}

func (h *TodoHandler) RegisterRoutes(webServer *webserver.WebServer) {
	webServer.AddHandler("GET /todos", h.GetTodos)
	webServer.AddHandler("POST /todos", h.CreateTodo)
	webServer.AddHandler("PUT /todos/{id}", h.UpdateTodo)
	webServer.AddHandler("DELETE /todos/{id}", h.DeleteTodo)
	webServer.AddHandler("PUT /todos/{id}/complete", h.Complete)
	webServer.AddHandler("PUT /todos/{id}/undo", h.Undo)
	webServer.AddHandler("GET /todos/{id}", h.GetTodo)
}
