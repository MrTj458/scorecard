import React from "react"
import ReactDOM from "react-dom"
import { Toaster } from "react-hot-toast"
import { BrowserRouter } from "react-router-dom"
import App from "./App"
import "./styles/index.css"

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
      <Toaster />
      <App />
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById("root")
)
