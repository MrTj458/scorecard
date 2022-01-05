import { Link } from "react-router-dom"

export default function TitleBar() {
  return (
    <div className="bg-gray-800 text-white p-2 w-full flex justify-center sticky top-0">
      <Link to="/" className="text-2xl">
        DG Scorecard
      </Link>
    </div>
  )
}
