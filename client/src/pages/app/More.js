import axios from "axios"
import { useContext } from "react"
import { useNavigate } from "react-router-dom"
import UserContext from "../../context/UserContext"

export default function More() {
  const { setUser } = useContext(UserContext)

  const navigate = useNavigate()

  const logout = async () => {
    await axios.post("/api/users/logout")
    navigate("/")
    setUser(null)
  }

  return (
    <>
      <button
        onClick={logout}
        className="w-full bg-gray-200 p-3 text-red-500 m-2"
      >
        Log Out
      </button>
    </>
  )
}
