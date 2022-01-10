import { useState } from "react"
import { BsFillGearFill } from "react-icons/bs"
import { FaUserCircle } from "react-icons/fa"
import Modal from "./Modal"
import ScorecardOptions from "./ScorecardOptions"

export default function Scorecard({ card }) {
  const startDate = new Date(card.start_time).toLocaleString()
  const endDate = new Date(card.end_time).toLocaleTimeString()

  const coursePar = card.holes.reduce((prev, cur) => prev + cur.par, 0)

  const [open, setOpen] = useState(false)

  const getColor = (hole, score) => {
    let difference = score.strokes - hole.par
    let color = "bg-white"

    if (difference === -1) {
      color = "bg-blue-200"
    }

    if (difference === -2) {
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
    <>
      <main className="w-full">
        {/* Title */}
        <div className="py-3 px-2 flex justify-between">
          <div>
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
          <button className="p-3" onClick={() => setOpen(true)}>
            <BsFillGearFill size="1.5em" />
          </button>
        </div>

        {/* Total Scores  */}
        <section className="my-3 px-2">
          <ul>
            {card.players.map((player) => {
              const playerStrokes = card.holes.reduce(
                (prev, cur) =>
                  prev +
                  cur.scores.filter((score) => score.id === player.id)[0]
                    .strokes,
                0
              )
              let playerPar = playerStrokes - coursePar

              if (playerPar === 0) {
                playerPar = "E"
              } else if (playerPar > 0) {
                playerPar = "+" + playerPar
              }

              return (
                <li
                  key={player.id}
                  className="flex justify-between items-center mt-2"
                >
                  <div className="flex items-center gap-4 font-bold">
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
          <ul className="grid grid-cols-12 text-center text-sm">
            {card.holes.map((hole) => (
              <>
                {/* Create new row */}
                {hole.number % 9 === 1 && (
                  <li key={hole.number * -1} className="col-span-3 mr-2">
                    <div className="text-gray-500 text-right mt-4">
                      <p>Hole</p>
                      <p>Dist</p>
                      <p>Par</p>
                    </div>

                    {/* Show usernames */}
                    <ul>
                      {hole.scores.map((score) => (
                        <li key={score.id} className="font-bold text-left">
                          {score.username}
                        </li>
                      ))}
                    </ul>
                  </li>
                )}

                {/* Show scores */}
                <li key={hole.number}>
                  {/* Hole data */}
                  <div className="text-gray-500 mt-4">
                    <p className="font-bold text-black">{hole.number}</p>
                    <p>{hole.distance}</p>
                    <p>{hole.par}</p>
                  </div>

                  {/* Hole scores */}
                  <ul>
                    {hole.scores.map((score) => (
                      <li key={score.id} className={getColor(hole, score)}>
                        {score.strokes}
                      </li>
                    ))}
                  </ul>
                </li>
              </>
            ))}
          </ul>
        </section>

        {/* Spacer */}
        <div className="h-60"></div>
      </main>
      <Modal open={open} close={() => setOpen(false)}>
        <ScorecardOptions card={card} />
      </Modal>
    </>
  )
}
