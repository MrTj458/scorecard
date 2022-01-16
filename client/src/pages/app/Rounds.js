import axios from "axios"
import { useEffect, useState } from "react"
import { Link } from "react-router-dom"
import ScorecardListItem from "../../components/ScorecardListItem"

export default function Rounds() {
  const [cards, setCards] = useState([])

  const fetchCards = async () => {
    try {
      const res = await axios.get(`/api/scorecards`)
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
            <ScorecardListItem card={card} />
          </li>
        ))}
      </ul>

      <div className="h-60"></div>
    </>
  )
}
