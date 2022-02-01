import axios from "axios"
import { useContext, useEffect } from "react"
import toast from "react-hot-toast"
import { Navigate } from "react-router-dom"
import UserContext from "../context/UserContext"

export default function SignOutPage() {
  const { user, setUser } = useContext(UserContext)

  const signOut = async () => {
    try {
      await axios.post("/api/users/logout")
      setUser(null)
      toast.success("Signed out")
    } catch (e) {
      console.error(e)
      toast.error("Error signing out, please try again.")
    }
  }

  useEffect(() => {
    if (user) {
      signOut()
    }
  }, [])

  return <Navigate to="/" />
}
