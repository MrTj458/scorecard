import { Outlet } from "react-router-dom"
import NavBar from "./NavBar"
import TitleBar from "./TitleBar"

export default function AppLayout() {
  return (
    <>
      <TitleBar />
      <Outlet />
      <NavBar />
    </>
  )
}
