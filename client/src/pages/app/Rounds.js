import { useContext } from "react"
import UserContext from "../../context/UserContext"

export default function Rounds() {
  const { user } = useContext(UserContext)

  return (
    <>
      <h1>Rounds for {user.username}</h1>
    </>
  )
}
