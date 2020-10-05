package controllers

import (
	"beeGo/models"
	"encoding/json"
	"io/ioutil"

	uuid "github.com/satori/go.uuid"
)

type TodoController struct {
	MainController
}

func (c *TodoController) Get() {
	idParam := uuid.FromStringOrNil(c.Ctx.Input.Param(":id"))

	todo := models.Todo{}
	result, err := models.Todo.GetByID(todo, idParam)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  err,
			"success": false,
		}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{
		"result":  result,
		"success": true,
	}
	c.ServeJSON()
}

func (c *TodoController) Create() {

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	todo := models.Todo{}
	err = json.Unmarshal(body, &todo)

	result, err := models.Todo.PostTodo(todo, todo)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  err,
			"success": false,
		}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{
		"result":  result,
		"success": true,
	}
	c.ServeJSON()
}

func (c *TodoController) Update() {

	idParam := uuid.FromStringOrNil(c.Ctx.Input.Param(":id"))

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	todo := models.Todo{}
	err = json.Unmarshal(body, &todo)

	todo.ID = idParam

	result, err := models.Todo.UpdateActivity(todo, todo)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  err,
			"success": false,
		}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{
		"result":  result,
		"success": true,
	}
	c.ServeJSON()
}

func (c *TodoController) Delete() {

	idParam := uuid.FromStringOrNil(c.Ctx.Input.Param(":id"))
	todo := models.Todo{}
	result, err := models.Todo.DeleteActivity(todo, idParam)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  err,
			"success": false,
		}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{
		"result":  result,
		"success": true,
	}
	c.ServeJSON()
}
