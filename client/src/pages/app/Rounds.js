import axios from "axios"
import { useContext, useEffect, useState } from "react"
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
      <p className="mt-2 bg-gray-100 w-full text-center">
        {cards.length} Rounds Total
      </p>
      {cards.map((card) => (
        <ScoreCard key={card.id} card={card} />
      ))}
    </>
  )
}
