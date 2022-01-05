import { Link } from "react-router-dom"
import ScorePlayer from "./ScorePlayer"

export default function ScoreCard({ card }) {
  const date = new Date(card.start_time).toLocaleString()
  const par = card.holes.reduce((prev, cur) => prev + cur.par, 0)

  return (
    <Link
      to={`/scorecards/${card.id}`}
      className="w-full bg-gray-100 m-2 px-2 py-1"
    >
      <h3 className="text-sm text-gray-500">
        <span className="text-black font-bold text-lg">
          {card.course_name}, {card.course_state}
        </span>{" "}
        - {card.holes.length} holes - Par {par}
      </h3>

      <p className="text-xs text-gray-600">{date}</p>

      <div className="mt-1 flex flex-wrap justify-around">
        {card.players.map((player) => (
          <ScorePlayer
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
