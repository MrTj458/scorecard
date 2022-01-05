export default function ScorePlayer({ player, holes, coursePar }) {
  const totalStrokes = holes.reduce(
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
    <div>
      <p className="font-bold">{player.username}</p>
      <p className="text-gray-500">
        {par} ({totalStrokes})
      </p>
    </div>
  )
}
