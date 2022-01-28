import { FaUserCircle } from "react-icons/fa"

export default function ScorecardStrokeCounter({ player, card, setStrokes }) {
  const strokes = player.strokes

  const coursePar = card.holes.reduce((prev, cur) => prev + cur.par, 0)
  const totalStrokes = card.holes.reduce(
    (prev, cur) =>
      prev + cur.scores.filter((score) => score.id === player.id)[0].strokes,
    0
  )
  let par = totalStrokes - coursePar

  if (par === 0) {
    par = "E"
  } else if (par > 0) {
    par = "+" + par
  }

  return (
    <div className="w-full flex justify-between items-center bg-gray-100 p-4 my-2">
      <div className="flex items-center gap-2 text-sm">
        <FaUserCircle size="2em" />
        <div>
          <p className="font-bold">{player.username}</p>
          <p className="text-gray-500">
            {par} ({totalStrokes})
          </p>
        </div>
      </div>
      <div className="flex items-center gap-3">
        <button
          type="button"
          onClick={() => strokes > 1 && setStrokes(player.id, strokes - 1)}
          className="bg-orange-400 w-10 h-10 rounded-full p-2"
        >
          -
        </button>
        <p>{player.strokes}</p>
        <button
          type="button"
          onClick={() => setStrokes(player.id, strokes + 1)}
          className="bg-orange-400 w-10 h-10 rounded-full p-2"
        >
          +
        </button>
      </div>
    </div>
  )
}
