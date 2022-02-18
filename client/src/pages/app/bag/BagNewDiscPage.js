import axios from "axios"
import { useContext } from "react"
import { useForm } from "react-hook-form"
import toast from "react-hot-toast"
import { useNavigate } from "react-router-dom"
import UserContext from "../../../context/UserContext"

export default function BagNewDiscPage() {
  const navigate = useNavigate()
  const { user } = useContext(UserContext)
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm()

  const onSubmit = async (data) => {
    const disc = {
      ...data,
      created_by: user.id,
      in_bag: data.in_bag === "true",
      weight: +data.weight,
      speed: +data.speed,
      glide: +data.glide,
      turn: +data.turn,
      fade: +data.fade,
    }

    try {
      await axios.post("/api/discs", disc)
      navigate("/app/bag")
      toast.success(`${data.name} added to bag.`)
    } catch (e) {
      toast.error("Unable to add new disc. Please try again.")
      console.error(e.response.data)
    }
  }

  console.log(errors)

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <fieldset className="flex flex-col items-center">
          <legend>New Disc</legend>

          {/* Name */}
          <div>
            <label htmlFor="name" className="w-full block">
              Name
            </label>
            <input
              {...register("name", { required: true })}
              type="text"
              name="name"
              id="name"
              className={`border-2 border-gray-300 p-1 ${
                errors.name ? "border-red-600" : ""
              }`}
            />
            <small className="block text-red-600">
              {errors.name && "Disc name is required."}
            </small>
          </div>

          {/* Type */}
          <div>
            <input
              {...register("type", { required: true })}
              type="radio"
              name="type"
              id="driver"
              value="Driver"
              className={`border-2 border-gray-300 p-1 ${
                errors.type ? "border-red-600" : ""
              }`}
            />
            <label htmlFor="driver">Driver</label>

            <input
              {...register("type", { required: true })}
              type="radio"
              name="type"
              id="fairway"
              value="Fairway"
              className={`border-2 border-gray-300 p-1 ${
                errors.type ? "border-red-600" : ""
              }`}
            />
            <label htmlFor="fairway">Fairway</label>

            <input
              {...register("type", { required: true })}
              type="radio"
              name="type"
              id="mid"
              value="Mid"
              className={`border-2 border-gray-300 p-1 ${
                errors.type ? "border-red-600" : ""
              }`}
            />
            <label htmlFor="mid">Mid</label>

            <input
              {...register("type", { required: true })}
              type="radio"
              name="type"
              id="putter"
              value="Putter"
              className={`border-2 border-gray-300 p-1 ${
                errors.type ? "border-red-600" : ""
              }`}
            />
            <label htmlFor="putter">Putter</label>

            <small className="block text-red-600">
              {errors.type && "Disc type is required."}
            </small>
          </div>

          {/* Manufacturer */}
          <div>
            <label htmlFor="manufacturer" className="w-full block">
              Manufacturer
            </label>
            <input
              {...register("manufacturer", { required: true })}
              type="text"
              name="manufacturer"
              id="manufacturer"
              className={`border-2 border-gray-300 p-1 ${
                errors.manufacturer ? "border-red-600" : ""
              }`}
            />
            <small className="block text-red-600">
              {errors.manufacturer && "Disc manufacturer is required."}
            </small>
          </div>

          {/* Plastic */}
          <div>
            <label htmlFor="plastic" className="w-full block">
              Plastic
            </label>
            <input
              {...register("plastic", { required: true })}
              type="text"
              name="plastic"
              id="plastic"
              className={`border-2 border-gray-300 p-1 ${
                errors.plastic ? "border-red-600" : ""
              }`}
            />
            <small className="block text-red-600">{errors.plastic && ""}</small>
          </div>

          {/* Weight */}
          <div>
            <label htmlFor="weight" className="w-full block">
              Weight
            </label>
            <input
              {...register("weight", { min: 100, max: 200 })}
              type="number"
              name="weight"
              id="weight"
              className={`border-2 border-gray-300 p-1 ${
                errors.weight ? "border-red-600" : ""
              }`}
            />
            <small className="block text-red-600">
              {errors.weight && "Weight must be between 100 and 200."}
            </small>
          </div>

          {/* Flight numbers */}
          <div className="flex">
            {/* Speed */}
            <div>
              <label htmlFor="speed">Speed</label>
              <input
                {...register("speed", { min: 1, max: 14 })}
                type="number"
                name="speed"
                id="speed"
                className={`border-2 border-gray-300 p-1 ${
                  errors.speed ? "border-red-600" : ""
                }`}
              />
              <small className="block text-red-600">
                {errors.speed && "Speed must be between 1 and 14."}
              </small>
            </div>

            {/* Glide */}
            <div>
              <label htmlFor="glide" className="w-full block">
                Glide
              </label>
              <input
                {...register("glide", { min: 1, max: 14 })}
                type="number"
                name="glide"
                id="glide"
                className={`border-2 border-gray-300 p-1 ${
                  errors.glide ? "border-red-600" : ""
                }`}
              />
              <small className="block text-red-600">
                {errors.glide && "Glide must be between 1 and 10."}
              </small>
            </div>

            {/* Turn */}
            <div>
              <label htmlFor="turn">Turn</label>
              <input
                {...register("turn", { min: -5, max: 5 })}
                type="number"
                name="turn"
                id="turn"
                className={`border-2 border-gray-300 p-1 ${
                  errors.turn ? "border-red-600" : ""
                }`}
              />
              <small className="block text-red-600">
                {errors.turn && "Turn must be between -5 and 5."}
              </small>
            </div>

            {/* Fade */}
            <div>
              <label htmlFor="fade">Fade</label>
              <input
                {...register("fade", { min: 0, max: 10 })}
                type="number"
                name="fade"
                id="fade"
                className={`border-2 border-gray-300 p-1 ${
                  errors.fade ? "border-red-600" : ""
                }`}
              />
              <small className="block text-red-600">
                {errors.fade && "fade must be between 0 and 10."}
              </small>
            </div>
          </div>

          {/* In bag */}
          <div>
            <input
              {...register("in_bag", { required: true })}
              type="radio"
              name="in_bag"
              id="yes"
              value="true"
              className={`border-2 border-gray-300 p-1 ${
                errors.in_bag ? "border-red-600" : ""
              }`}
            />
            <label htmlFor="yes">Yes</label>

            <input
              {...register("in_bag", { required: true })}
              type="radio"
              name="in_bag"
              id="no"
              value="false"
              className={`border-2 border-gray-300 p-1 ${
                errors.in_bag ? "border-red-600" : ""
              }`}
            />
            <label htmlFor="no">No</label>
          </div>

          <button
            type="submit"
            className="px-4 py-2 text-white bg-gray-700 mt-4"
          >
            Add Disc
          </button>
        </fieldset>
      </form>
    </>
  )
}
