import { useContext, useEffect } from "react"
import toast from "react-hot-toast"
import { Navigate } from "react-router-dom"
import UserContext from "../context/UserContext"

export default function SignOut() {
  const { user, setUser } = useContext(UserContext)

  useEffect(() => {
    if (user) {
      setUser(null)
      toast.success("Signed out")
    }
  }, [])

  return <Navigate to="/" />
}
