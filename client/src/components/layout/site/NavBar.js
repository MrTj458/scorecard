import { Link } from "react-router-dom"

export default function NavBar() {
  return (
    <div className="bg-gray-800 text-white p-2 w-full sticky top-0 flex justify-center">
      <div className="flex justify-between items-center w-full max-w-5xl">
        <div>
          <Link to="/" className="text-3xl">
            Scorecard
          </Link>
        </div>

        <ul className="flex justify-center items-center gap-4">
          <li>
            <Link to="/login">Login</Link>
          </li>
          <li>
            <Link to="/signup">Sign Up</Link>
          </li>
        </ul>
      </div>
    </div>
  )
}
