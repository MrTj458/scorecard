import axios from "axios"
import { useNavigate } from "react-router-dom"

export default function More() {
  const navigate = useNavigate()

  const logout = async () => {
    await axios.post("/api/users/logout")
    navigate("/")
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
