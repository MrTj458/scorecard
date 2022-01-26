import { useContext } from "react"
import { Navigate, useLocation } from "react-router-dom"
import UserContext from "../context/UserContext"

export default function RequireAuth({ children }) {
  const { user } = useContext(UserContext)

  const location = useLocation()

  if (!user) {
    return <Navigate to="/login" state={{ from: location }} />
  }

  return children
}
