import axios from "axios"
import { useContext, useState } from "react"
import { useNavigate } from "react-router-dom"
import UserContext from "../context/UserContext"

export default function Login() {
  const navigate = useNavigate()
  const { setUser } = useContext(UserContext)

  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [error, setError] = useState("")

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      const res = await axios.post("/api/users/login", { email, password })
      setUser(res.data)
      navigate("/rounds", { replace: true })
    } catch (e) {
      if (e.response.status === 401) {
        setError("Incorrect email or password")
      } else {
        console.error(e)
      }
    }
  }

  return (
    <div className="mt-5">
      <form onSubmit={handleSubmit} className="border-2 border-gray-300">
        <fieldset className="flex flex-col items-center p-4">
          <legend className="w-full mt-2 text-xl font-bold text-center">
            Log In
          </legend>

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
            Log In
          </button>
        </fieldset>
      </form>
    </div>
  )
}
