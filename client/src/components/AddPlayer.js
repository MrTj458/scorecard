import axios from "axios"
import { useState } from "react"
import toast from "react-hot-toast"
import { FaUserCircle } from "react-icons/fa"

export default function AddPlayer({ addPlayer }) {
  const [username, setUsername] = useState("")
  const [users, setUsers] = useState([])

  const handleSubmit = async (e) => {
    e.preventDefault()

    try {
      const res = await axios.get(`/api/users?username=${username}`)
      setUsers(res.data)
    } catch (e) {
      console.error(e.response.data)
      toast.error("Unable to search users")
    }
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <legend className="text-center">Add Player</legend>
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          className="border-2 p-1 w-full"
        />
        <button className="w-full my-2 bg-gray-500 text-white text-center py-2">
          Search
        </button>
      </form>

      <ul>
        {users.map((user) => (
          <li key={user.id}>
            <div className="w-full flex justify-between my-2 bg-gray-100 p-3">
              <div className="flex gap-2 items-center">
                <FaUserCircle size="2em" />
                {user.username}
              </div>
              <button
                className="bg-orange-500 text-white py-2 px-4"
                onClick={() =>
                  addPlayer({ id: user.id, username: user.username })
                }
              >
                Add
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  )
}
