import axios from "axios"
import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import ScorecardDetail from "../../components/ScorecardDetail"
import ScorecardScoring from "../../components/ScorecardScoring"

export default function ScorecardPage() {
  const params = useParams()

  const [card, setCard] = useState(null)

  const fetchScorecard = async () => {
    try {
      const res = await axios.get(`/api/scorecards/${params.id}`)
      setCard(res.data)
    } catch (e) {
      console.error(e.response.data)
    }
  }

  useEffect(() => {
    fetchScorecard()
  }, [])

  if (!card) {
    return <p>Loading...</p>
  }

  if (!card.end_time) {
    return <ScorecardScoring card={card} />
  }

  return <ScorecardDetail card={card} />
}
