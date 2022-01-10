import { useContext } from "react"
import UserContext from "../../context/UserContext"

export default function Bag() {
  const { user } = useContext(UserContext)

  return (
    <>
      <h1>{user.username}'s Bag</h1>
      <p>Feature coming soon.</p>
    </>
  )
}
