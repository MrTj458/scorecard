import axios from "axios"
import { useContext, useState } from "react"
import toast from "react-hot-toast"
import { FaUserCircle } from "react-icons/fa"
import { useNavigate } from "react-router-dom"
import UserContext from "../../context/UserContext"

export default function NewScorecard() {
  const navigate = useNavigate()
  const { user } = useContext(UserContext)

  const [courseName, setCourseName] = useState("")
  const [courseState, setCourseState] = useState("")
  const [numHoles, setNumHoles] = useState(18)
  const [players, setPlayers] = useState([
    { id: user.id, username: user.username },
  ])

  const handleSubmit = async (e) => {
    e.preventDefault()

    const data = {
      created_by: user.id,
      course_name: courseName,
      course_state: courseState,
      num_holes: numHoles,
      players,
    }

    try {
      const res = await axios.post("/api/scorecards", data)
      navigate(`/scorecards/${res.data.id}`)
      toast.success("New Scorecard Created")
    } catch (e) {
      console.error(e.response.data)
      toast.error("Unable to create new scorecard")
    }
  }

  return (
    <>
      <form onSubmit={handleSubmit} className="w-full">
        <fieldset className="w-full flex flex-col items-center">
          <legend className="w-full text-center mt-2 text-xl font-bold">
            New Scorecard
          </legend>

          <div className="bg-gray-100 w-full p-2 m-2">
            <label htmlFor="course_name" className="w-full block">
              Course Name
            </label>
            <input
              required
              type="text"
              name="course_name"
              className="border-2 p-1 w-full"
              value={courseName}
              onChange={(e) => setCourseName(e.target.value)}
            />
          </div>

          <div className="bg-gray-100 w-full p-2 m-2">
            <label htmlFor="course_state" className="w-full p-2">
              Course State
            </label>
            <input
              required
              type="text"
              name="course_state"
              className="border-2 p-1 w-full"
              value={courseState}
              onChange={(e) => setCourseState(e.target.value)}
            />
          </div>

          <div className="bg-gray-100 w-full p-2 m-2">
            <label htmlFor="num_holes" className="w-full block">
              Number of Holes
            </label>
            <input
              required
              type="number"
              min="1"
              max="36"
              name="num_holes"
              className="border-2 p-1 w-full"
              value={numHoles}
              onChange={(e) => setNumHoles(+e.target.value)}
            />
          </div>

          <p>Players</p>

          <button
            type="button"
            className="w-full m-2 bg-gray-500 text-white text-center py-2"
          >
            Add Player
          </button>

          <ul className="w-full">
            {players.map((player) => (
              <li
                key={player.id}
                className="w-full bg-gray-100 flex items-center gap-2 p-2 my-2"
              >
                <FaUserCircle size="2em" />
                <p>{player.username}</p>
              </li>
            ))}
          </ul>

          <button className="w-full m-2 bg-orange-500 text-white text-center py-2">
            Create Scorecard
          </button>
        </fieldset>
      </form>
    </>
  )
}
