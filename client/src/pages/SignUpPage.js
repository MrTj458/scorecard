import axios from "axios"
import { useContext, useState } from "react"
import toast from "react-hot-toast"
import { useNavigate } from "react-router-dom"
import UserContext from "../context/UserContext"

export default function SignUpPage() {
  const navigate = useNavigate()
  const { setUser } = useContext(UserContext)

  const [username, setUsername] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [error, setError] = useState("")

  const handleSubmit = async (e) => {
    e.preventDefault()

    try {
      const res = await axios.post("/api/users", { username, email, password })
      setUser(res.data)
      navigate("/rounds", { replace: true })
    } catch (e) {
      if (e.response.status === 422) {
        setError(e.response.data.detail)
      } else {
        console.error(e)
        toast.error("Error creating new user, please try again.")
      }
    }
  }

  return (
    <div className="mt-5">
      <form onSubmit={handleSubmit} className="border-2 border-gray-300">
        <fieldset className="flex flex-col items-center p-4">
          <legend className="w-full mt-2 text-xl font-bold text-center">
            Sign Up
          </legend>

          <div>
            <label htmlFor="username" className="w-full block">
              Username
            </label>
            <input
              required
              name="username"
              type="text"
              className="border-2 border-gray-300 p-1"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>

          <div>
            <label htmlFor="email" className="w-full block">
              Email
            </label>
            <input
              required
              name="email"
              type="email"
              className="border-2 border-gray-300 p-1"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>

          <div>
            <label htmlFor="password" className="w-full block">
              Password
            </label>
            <input
              required
              name="password"
              type="password"
              className="border-2 border-gray-300 p-1"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>

          {error && <p className="mt-2 text-red-600">{error}</p>}

          <button
            type="submit"
            className="px-4 py-2 text-white bg-gray-700 mt-4"
          >
            Sign Up
          </button>
        </fieldset>
      </form>
    </div>
  )
}
