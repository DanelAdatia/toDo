import React, { useState } from "react"

const InputField = (props) => {
  const [text, settext] = useState("")

  const inputTextHandler = (e) => {
    settext(e.target.value)
  }

  const addButtonHandler = () => {
    if (text.length > 1) {
      props.click(text)
    }
    settext("")
  }

  return (
    <div className="header">
      <input
        onChange={inputTextHandler}
        type="text"
        value={text}
        placeholder="Enter title"
      />
      <button onClick={addButtonHandler}>Add</button>
    </div>
  )
}

export default InputField
