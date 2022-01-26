import { Link } from "react-router-dom"

export default function More() {
  return (
    <>
      <Link
        to="/signout"
        className="w-full bg-gray-200 p-3 text-red-500 m-2 text-center"
      >
        Sign Out
      </Link>
    </>
  )
}
