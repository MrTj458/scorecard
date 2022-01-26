import React from "react"
import ReactDOM from "react-dom"
import { Toaster } from "react-hot-toast"
import { HashRouter } from "react-router-dom"
import App from "./App"
import "./styles/index.css"

ReactDOM.render(
  <React.StrictMode>
    <HashRouter>
      <Toaster />
      <App />
    </HashRouter>
  </React.StrictMode>,
  document.getElementById("root")
)
