import React from "react"

const Todo = (props) => {
  const deleteTodo = () => {
    props.click(props.id)
  }

  const completeTodo = () => {
    props.complete(props.id)
  }

  return (
    <div className="todo">
      {props.completed ? (
        <p>
          <strike>{props.title}</strike>
        </p>
      ) : (
        <p>{props.title}</p>
      )}
      <div>
        {props.completed ? (
          <span onClick={completeTodo}>
            <i className="fa fa-times"></i>
          </span>
        ) : (
          <span onClick={completeTodo}>
            <i className="fa fa-check"></i>
          </span>
        )}

        <span onClick={deleteTodo}>
          <i className="fa fa-trash"></i>
        </span>
      </div>
    </div>
  )
}

export default Todo
