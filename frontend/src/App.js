import React, { useState, useEffect } from "react"
import "./App.css"
import InputField from "./components/InputField"
import Todo from "./components/Todo"

function App() {
  const [todos, settodos] = useState([
    {
      id: 0,
      title: "",
      completed: false,
    },
  ])

  useEffect(() => {
    fetch("http://localhost:8080/api/todos")
      .then((res) => res.json())
      .then((data) => {
        settodos(data)
      })
  }, [])

  const inputHandler = (text) => {
    const todo = {
      title: text,
      completed: false,
    }
    fetch("http://localhost:8080/api/todos", {
      method: "POST",
      body: JSON.stringify(todo),
    })
      .then((res) => res.json())
      .then((data) => {
        settodos([...todos, data])

      })
  }

  const deleteTodo = (id) => {
    fetch(`http://localhost:8080/api/todos/${id}`, {
      method: "DELETE",
    }).then(() => {
      var index = todos.findIndex((obj) => {
        return obj.id === id
      })
      let allTodos = [...todos]
      allTodos.splice(index, 1)
      settodos(allTodos)
    })
  }

  const completeTodo = (id) => {
    var index = todos.findIndex((obj) => {
      return obj.id === id
    })

    let allTodos = [...todos]
    
    allTodos[index].completed = !allTodos[index].completed

    let updatedTodo = allTodos[index]

    fetch(`http://localhost:8080/api/todos/${id}`, {
      method: "PUT",
      body: JSON.stringify(updatedTodo)
    }).then(() => {
      settodos(allTodos)
    })
  }

  return (
    <div className="App">
      <InputField click={inputHandler} />

      {todos.map((todo) => {
        return (
          <Todo
            click={() => {
              deleteTodo(todo.id)
            }}
            complete={() => {
              completeTodo(todo.id)
            }}
            key={todo.id}
            id={todo.id}
            title={todo.title}
            completed={todo.completed}
          />
        )
      })}

    </div>
  )
}

export default App
