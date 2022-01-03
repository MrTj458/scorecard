import { Link } from "react-router-dom"

export default function NavBar() {
  return (
    <div className="fixed bottom-0 left-0 bg-gray-200 p-4 w-full border-t-2 border-t-gray-400">
      <ul className="flex justify-around">
        <li>
          <Link to="/rounds">Rounds</Link>
        </li>
        <div className="text-gray-400">|</div>
        <li>
          <Link to="/profile">Profile</Link>
        </li>
        <div className="text-gray-400">|</div>
        <li>
          <Link to="/more">More</Link>
        </li>
      </ul>
    </div>
  )
}
