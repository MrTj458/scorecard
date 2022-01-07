import axios from "axios"
import { useContext, useEffect, useState } from "react"
import { Link } from "react-router-dom"
import ScoreCard from "../../components/ScoreCard"
import UserContext from "../../context/UserContext"

export default function Rounds() {
  const { user } = useContext(UserContext)

  const [cards, setCards] = useState([])

  const fetchCards = async () => {
    try {
      const res = await axios.get(`/api/scorecards?user=${user.id}`)
      setCards(res.data)
    } catch (e) {
      console.error(e)
    }
  }

  useEffect(() => {
    fetchCards()
  }, [])

  return (
    <>
      <Link
        to="/scorecards/new"
        className="w-full m-2 bg-orange-500 text-white text-center py-2"
      >
        Create New Scorecard
      </Link>

      <ul className="w-full">
        {cards.map((card) => (
          <li key={card.id}>
            <ScoreCard card={card} />
          </li>
        ))}
      </ul>

      <div className="h-60"></div>
    </>
  )
}
