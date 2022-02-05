import axios from "axios"
import { useContext, useState } from "react"
import { useForm } from "react-hook-form"
import toast from "react-hot-toast"
import { useLocation, useNavigate } from "react-router-dom"
import UserContext from "../context/UserContext"

export default function LoginPage() {
  const navigate = useNavigate()
  const location = useLocation()
  const { setUser } = useContext(UserContext)
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm()

  const [authError, setAuthError] = useState("")

  let from = location.state?.from?.pathname || "/app/scorecards"

  const onSubmit = async (data) => {
    try {
      const res = await axios.post("/api/users/login", {
        email: data.email,
        password: data.password,
      })
      setUser(res.data)
      navigate(from, { replace: true })
      toast.success(`Signed in as ${res.data.username}`)
    } catch (e) {
      if (e.response.status === 401) {
        setAuthError("Incorrect email or password")
      } else {
        console.error(e)
        toast.error("Error signing in, please try again.")
      }
    }
  }

  return (
    <div className="mt-5">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="border-2 border-gray-300"
      >
        <fieldset className="flex flex-col items-center p-4">
          <legend className="w-full mt-2 text-xl font-bold text-center">
            Log In
          </legend>

          <div>
            <label htmlFor="email" className="w-full block">
              Email
            </label>
            <input
              {...register("email", { required: true })}
              name="email"
              className={`border-2 border-gray-300 p-1 ${
                errors.email ? "border-red-600" : ""
              }`}
            />
            <small className="block text-red-600">
              {errors.email && "Email is required."}
            </small>
          </div>

          <div>
            <label htmlFor="password" className="w-full block">
              Password
            </label>
            <input
              {...register("password", { required: true })}
              name="password"
              type="password"
              className={`border-2 border-gray-300 p-1 ${
                errors.email ? "border-red-600" : ""
              }`}
            />
            <small className="block text-red-600">
              {errors.password && "Password is required."}
            </small>
          </div>

          {authError && <p className="mt-2 text-red-600">{authError}</p>}

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
