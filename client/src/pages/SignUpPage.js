import axios from "axios"
import { useContext, useState } from "react"
import { useForm } from "react-hook-form"
import toast from "react-hot-toast"
import { useNavigate } from "react-router-dom"
import UserContext from "../context/UserContext"

export default function SignUpPage() {
  const navigate = useNavigate()
  const { setUser } = useContext(UserContext)
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm()

  const [registerError, setRegisterError] = useState("")

  const onSubmit = async (data) => {
    try {
      const res = await axios.post("/api/users", {
        username: data.username,
        email: data.email,
        password: data.password,
      })
      setUser(res.data)
      navigate("/app/scorecards", { replace: true })
    } catch (e) {
      if (e.response.status === 422) {
        setRegisterError(e.response.data.detail)
      } else {
        console.error(e)
        toast.error("Error creating new user, please try again.")
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
            Sign Up
          </legend>

          <div>
            <label htmlFor="username" className="w-full block">
              Username
            </label>
            <input
              {...register("username", { required: true })}
              name="username"
              type="text"
              className={`border-2 border-gray-300 p-1 ${
                errors.username ? "border-red-600" : ""
              }`}
            />
            <small className="block text-red-600">
              {errors.username && "Username is required."}
            </small>
          </div>

          <div>
            <label htmlFor="email" className="w-full block">
              Email
            </label>
            <input
              {...register("email", { required: true })}
              name="email"
              type="email"
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
                errors.password ? "border-red-600" : ""
              }`}
            />
            <small className="block text-red-600">
              {errors.password && "Password is required."}
            </small>
          </div>

          {registerError && (
            <p className="mt-2 text-red-600">{registerError}</p>
          )}

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
