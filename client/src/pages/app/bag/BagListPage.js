import axios from "axios"
import { useEffect, useState } from "react"
import toast from "react-hot-toast"
import DiscListItem from "../../../components/bag/DiscListItem"

export default function BagListPage() {
  const [discs, setDiscs] = useState([])

  const fetchDiscs = async () => {
    try {
      const res = await axios.get("/api/discs")
      setDiscs(res.data)
    } catch (e) {
      console.error(e)
      toast.error("Error fetching discs.")
    }
  }

  useEffect(() => {
    fetchDiscs()
  }, [])

  return (
    <>
      <ul className="w-full mt-3">
        {discs.map((disc) => (
          <DiscListItem key={disc.id} disc={disc} />
        ))}
      </ul>
    </>
  )
}
