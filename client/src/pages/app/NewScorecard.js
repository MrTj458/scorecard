import axios from "axios"
import { useContext, useState } from "react"
import toast from "react-hot-toast"
import { FaUserCircle } from "react-icons/fa"
import { useNavigate } from "react-router-dom"
import Modal from "../../components/Modal"
import PlayerForm from "../../components/PlayerForm"
import UserContext from "../../context/UserContext"

export default function NewScorecard() {
  const navigate = useNavigate()
  const { user } = useContext(UserContext)

  const [showModal, setShowModal] = useState(false)

  const [courseName, setCourseName] = useState("")
  const [courseState, setCourseState] = useState("")
  const [players, setPlayers] = useState([
    { id: user.id, username: user.username },
  ])

  const addPlayer = (newPlayer) => {
    if (players.length === 4) {
      toast.error("Can't add more than 4 players.")
      return
    }

    if (players.some((p) => p.id === newPlayer.id)) {
      toast.error("Can't add the same player twice.")
      return
    }

    setPlayers([...players, newPlayer])
    setShowModal(false)
    toast.success("Player added")
  }

  const handleSubmit = async (e) => {
    e.preventDefault()

    const data = {
      created_by: user.id,
      course_name: courseName,
      course_state: courseState,
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

          <div className="bg-gray-100 w-full p-4 m-2">
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

          <div className="bg-gray-100 w-full p-4 m-2">
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

          <p>Players</p>

          <button
            type="button"
            className="w-full m-2 bg-gray-500 text-white text-center py-2"
            onClick={() => setShowModal(!showModal)}
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
      <Modal open={showModal} close={() => setShowModal(false)}>
        <PlayerForm addPlayer={addPlayer} />
      </Modal>
    </>
  )
}
