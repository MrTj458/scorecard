import { FaUserCircle } from "react-icons/fa"

export default function ScorecardDetail({ card }) {
  const startDate = new Date(card.start_time).toLocaleString()
  const endDate = new Date(card.end_time).toLocaleTimeString()

  const coursePar = card.holes.reduce((prev, cur) => prev + cur.par, 0)

  const handleDelete = async (e) => {
    e.preventDefault()
    // TODO: complete delete functionality
  }

  const getColor = (hole, score) => {
    let difference = score.strokes - hole.par
    let color = "bg-white"

    if (difference === -1) {
      color = "bg-blue-200"
    }

    if (difference == -2) {
      color = "bg-blue-400"
    }

    if (difference === 1) {
      color = "bg-orange-200"
    }

    if (difference === 2) {
      color = "bg-orange-400"
    }

    if (difference === 3) {
      color = "bg-orange-600"
    }

    return color
  }

  return (
    <main className="w-full">
      {/* Title */}
      <div className="py-3 px-2">
        <h1 className="text-xl font-bold">
          {card.course_name}
          <span className="font-normal text-gray-500">
            , {card.course_state} - {card.holes.length} holes
          </span>
        </h1>
        <p className="">
          {startDate} - {endDate}
        </p>
      </div>

      {/* Total Scores  */}
      <section className="my-3 px-2">
        <ul>
          {card.players.map((player) => {
            const playerStrokes = card.holes.reduce(
              (prev, cur) =>
                prev +
                cur.scores.filter((score) => score.id === player.id)[0].strokes,
              0
            )
            let playerPar = playerStrokes - coursePar

            if (playerPar === 0) {
              playerPar = "E"
            } else if (playerPar > 0) {
              playerPar = "+" + playerPar
            }

            return (
              <li className="flex justify-between items-center mt-2">
                <div className="flex items-center gap-4">
                  <FaUserCircle size="2em" />
                  {player.username}
                </div>
                <div className="flex items-center gap-4">
                  <p>{playerPar}</p>
                  <p>{playerStrokes}</p>
                </div>
              </li>
            )
          })}
        </ul>
      </section>

      {/* Score Breakdown */}
      <section className="w-full mt-3 px-2">
        <ul className="grid grid-cols-12 text-center">
          {card.holes.map((hole) => (
            <>
              {hole.number % 9 === 1 && (
                <li className="col-span-3 mr-2">
                  <div className="text-xs text-gray-500 text-right mt-4">
                    <p>Hole</p>
                    <p>Dist</p>
                    <p>Par</p>
                  </div>
                  <ul>
                    {hole.scores.map((score) => (
                      <li className="font-bold text-left">{score.username}</li>
                    ))}
                  </ul>
                </li>
              )}
              <li>
                <div className="text-xs text-gray-500 mt-4">
                  <p className="font-bold text-black">{hole.number}</p>
                  <p>{hole.distance}</p>
                  <p>{hole.par}</p>
                </div>
                <ul>
                  {hole.scores.map((score) => (
                    <li className={getColor(hole, score)}>{score.strokes}</li>
                  ))}
                </ul>
              </li>
            </>
          ))}
        </ul>
      </section>

      {/* Options */}
      <section className="mt-8">
        <form onSubmit={handleDelete}>
          <fieldset>
            <button className="w-full bg-red-500 text-white p-2">
              Delete Scorecard
            </button>
          </fieldset>
        </form>
      </section>

      {/* Spacer */}
      <div className="h-60"></div>
    </main>
  )
}
