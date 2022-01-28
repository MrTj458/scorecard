import { Link } from "react-router-dom"
import ScorecardScoreIcon from "./ScorecardScoreIcon"

export default function ScorecardListItem({ card }) {
  const date = new Date(card.start_time).toLocaleString()
  const par = card.holes.reduce((prev, cur) => prev + cur.par, 0)

  return (
    <Link
      to={`/app/scorecards/${card.id}`}
      className={`block w-ful my-2 px-2 py-1 ${
        card.end_time ? "bg-gray-100" : "bg-blue-200"
      }`}
    >
      <h3 className="text-sm text-gray-500">
        <span className="text-black font-bold text-lg">
          {card.course_name}, {card.course_state}
        </span>{" "}
        - {card.holes.length} Holes - Par {par}
      </h3>

      <p className="text-xs text-gray-600">{date}</p>

      <div className="mt-1 flex flex-wrap justify-around">
        {card.players.map((player) => (
          <ScorecardScoreIcon
            key={player.id}
            player={player}
            holes={card.holes}
            coursePar={par}
          />
        ))}
      </div>
    </Link>
  )
}
