import axios from "axios"
import { createContext, useEffect, useState } from "react"

const UserContext = createContext()

export function UserProvider({ children }) {
  const [user, setUser] = useState(null)
  const [loading, setLoading] = useState(true)

  const fetchUser = async () => {
    try {
      const res = await axios.get("/api/users/me")
      setUser(res.data)
      setLoading(false)
    } catch (e) {
      setUser(null)
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchUser()
  }, [])

  return (
    <UserContext.Provider value={{ user, setUser }}>
      {!loading && children}
    </UserContext.Provider>
  )
}

export default UserContext
