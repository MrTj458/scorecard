import axios from "axios"
import { useContext, useEffect, useState } from "react"
import toast from "react-hot-toast"
import { useParams } from "react-router-dom"
import ScorecardDetail from "../../components/ScorecardDetail"
import ScorecardScoring from "../../components/ScorecardScoring"
import UserContext from "../../context/UserContext"

export default function ScorecardPage() {
  const params = useParams()
  const { user } = useContext(UserContext)

  const [card, setCard] = useState(null)

  const fetchScorecard = async () => {
    try {
      const res = await axios.get(`/api/scorecards/${params.id}`)
      setCard(res.data)
    } catch (e) {
      console.error(e.response.data)
    }
  }

  const complete = async () => {
    try {
      const res = await axios.post(`/api/scorecards/${card.id}/complete`)
      setCard(res.data)
      toast.success("Scorecard marked as finished")
    } catch (e) {
      console.error(e.response.data)
      toast.error("Unable to mark scorecard as finished, please try again")
    }
  }

  useEffect(() => {
    fetchScorecard()
  }, [])

  if (!card) {
    return <p>Loading...</p>
  }

  if (!card.end_time && card.created_by === user.id) {
    return <ScorecardScoring card={card} complete={complete} />
  }

  return <ScorecardDetail card={card} />
}
