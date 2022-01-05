import axios from "axios"
import { useEffect, useState } from "react"
import toast from "react-hot-toast"
import PlayerStrokes from "./PlayerStrokes"

export default function ScorecardScoring({ card: initialCard, complete }) {
  const [card, setCard] = useState(initialCard)
  const hole = card.holes.length + 1

  const [distance, setDistance] = useState(300)
  const [par, setPar] = useState(3)
  const [scores, setScores] = useState([])

  useEffect(() => {
    const initialScores = card.players.map((player) => {
      return {
        id: player.id,
        username: player.username,
        strokes: 1,
      }
    })
    setScores(initialScores)
    setDistance(300)
    setPar(3)
  }, [card])

  const setStrokes = (id, strokes) => {
    setScores(
      scores.map((player) => {
        if (player.id === id) {
          return {
            id: player.id,
            username: player.username,
            strokes,
          }
        }
        return player
      })
    )
  }

  const handleSubmit = async (e) => {
    e.preventDefault()

    const data = {
      number: hole,
      distance,
      par,
      scores,
    }

    try {
      const res = await axios.post(`/api/scorecards/${card.id}/hole`, data)
      setCard(res.data)
      toast.success("Hole Saved")
    } catch (e) {
      console.error(e.response.data)
      toast.error("Unable to save hole, please try again")
    }
  }

  return (
    <form onSubmit={handleSubmit} className="w-full text-center">
      <p className="text-gray-500">
        {card.course_name}, {card.course_state}
      </p>
      <fieldset>
        <legend className="w-full font-bold mt-2 text-2xl">Hole {hole}</legend>
        <div className="bg-gray-100 w-full p-2 my-2">
          <label htmlFor="distance" className="w-full block">
            Distance
          </label>
          <input
            required
            type="number"
            min="1"
            max="5000"
            name="distance"
            className="border-2 p-1 w-full"
            value={distance}
            onChange={(e) => setDistance(+e.target.value)}
          />
        </div>

        <div className="w-full flex items-center justify-between p-4 my-2 bg-gray-100">
          <p>Par</p>
          <div className=" flex items-center gap-3">
            <button
              type="button"
              onClick={() => par > 1 && setPar(par - 1)}
              className="bg-orange-400 rounded rounded-full p-2"
            >
              -
            </button>
            <p>{par}</p>
            <button
              type="button"
              onClick={() => setPar(par + 1)}
              className="bg-orange-400 rounded rounded-full p-2"
            >
              +
            </button>
          </div>
        </div>

        <p>Player's Strokes</p>

        <ul className="w-full">
          {scores.map((player) => (
            <li key={player.id}>
              <PlayerStrokes
                player={player}
                card={card}
                setStrokes={setStrokes}
              />
            </li>
          ))}
        </ul>

        <button
          type="submit"
          className="w-full my-2 bg-orange-500 text-white text-center py-2"
        >
          Next Hole
        </button>

        <button
          type="button"
          className="w-full my-2 bg-gray-500 text-white text-center py-2"
          onClick={complete}
        >
          Finish Scorecard
        </button>
      </fieldset>
    </form>
  )
}
