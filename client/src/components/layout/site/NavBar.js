import { useContext } from "react"
import { Link } from "react-router-dom"
import UserContext from "../../../context/UserContext"

export default function NavBar() {
  const { user } = useContext(UserContext)

  return (
    <div className="bg-gray-800 text-white p-2 w-full sticky top-0 flex justify-center">
      <div className="flex justify-between items-center w-full max-w-5xl">
        <div>
          <Link to="/" className="text-3xl">
            Scorecard
          </Link>
        </div>

        <ul className="flex justify-center items-center gap-4">
          {user ? (
            <>
              <li>
                <Link to="/rounds">Go To App</Link>
              </li>
              <li>
                <Link to="/signout">Sign Out</Link>
              </li>
            </>
          ) : (
            <>
              <li>
                <Link to="/login">Login</Link>
              </li>
              <li>
                <Link to="/signup">Sign Up</Link>
              </li>
            </>
          )}
        </ul>
      </div>
    </div>
  )
}
