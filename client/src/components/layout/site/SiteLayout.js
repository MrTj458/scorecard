import { Outlet } from "react-router-dom"
import NavBar from "./NavBar"

export default function SiteLayout() {
  return (
    <>
      <NavBar />
      <Outlet />
    </>
  )
}
