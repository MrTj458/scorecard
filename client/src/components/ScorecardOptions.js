import axios from "axios"
import { useContext } from "react"
import toast from "react-hot-toast"
import { useNavigate } from "react-router-dom"
import UserContext from "../context/UserContext"

export default function ScorecardOptions({ card }) {
  const navigate = useNavigate()
  const { user } = useContext(UserContext)

  const handleDelete = async (e) => {
    e.preventDefault()

    try {
      await axios.delete(`/api/scorecards/${card.id}`)
      toast.success("Scorecard deleted.")
      navigate("/rounds", { replace: true })
    } catch (e) {
      console.error(e.response.value)
      toast.error("Error deleting scorecard, please try again.")
    }
  }

  return (
    <div>
      <h1 className="font-bold text-center my-2 text-lg">Options</h1>
      {user.id === card.created_by && (
        <form onSubmit={handleDelete} className="my-2">
          <button className="p-2 w-full text-white bg-red-500">
            Delete Scorecard
          </button>
        </form>
      )}
    </div>
  )
}
