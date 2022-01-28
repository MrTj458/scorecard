import { Outlet } from "react-router-dom"
import NavBar from "./NavBar"
import TitleBar from "./TitleBar"

export default function AppLayout() {
  return (
    <>
      <TitleBar />
      <div className="w-screen max-w-2xl flex flex-col items-center">
        <Outlet />
      </div>
      <NavBar />
    </>
  )
}
