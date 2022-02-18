import axios from "axios"
import { useEffect, useState } from "react"
import toast from "react-hot-toast"
import { Link } from "react-router-dom"
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
      <Link
        to="/app/bag/new"
        className="w-full m-2 bg-orange-500 text-white text-center py-2"
      >
        Add New Disc
      </Link>

      {discs.length === 0 ? (
        <p>You don't have any stored discs.</p>
      ) : (
        <ul className="w-full mt-3">
          {discs.map((disc) => (
            <DiscListItem key={disc.id} disc={disc} />
          ))}
        </ul>
      )}
    </>
  )
}
