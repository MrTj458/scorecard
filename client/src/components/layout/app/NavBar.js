import { BsThreeDots } from "react-icons/bs"
import { FaListAlt } from "react-icons/fa"
import { MdBackpack } from "react-icons/md"
import { Link } from "react-router-dom"

export default function NavBar() {
  return (
    <div className="fixed flex justify-center bottom-0 left-0 bg-gray-200 w-full border-t-2 border-t-gray-400">
      <ul className="flex justify-around items-center p-4 w-full max-w-2xl">
        <Link to="/app/scorecards">
          <li className="flex flex-col items-center">
            <FaListAlt />
            Scorecards
          </li>
        </Link>
        <div className="h-9 w-[1px] bg-gray-300"></div>
        <Link to="/app/bag">
          <li className="flex flex-col items-center">
            <MdBackpack />
            Bag
          </li>
        </Link>
        <div className="h-9 w-[1px] bg-gray-300"></div>
        <Link to="/app/more">
          <li className="flex flex-col items-center">
            <BsThreeDots />
            More
          </li>
        </Link>
      </ul>
    </div>
  )
}
