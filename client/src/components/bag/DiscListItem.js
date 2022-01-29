import { Link } from "react-router-dom"

export default function DiscListItem({ disc }) {
  const toggleInBag = async () => {
    // TODO
  }

  return (
    <li>
      <div className="mb-3 bg-gray-100 w-full flex justify-between items-center">
        <div className="block flex-grow">
          <Link to={`/app/bag/${disc.id}`}>
            <div className="p-4">
              <div>
                <h2 className="font-bold text-xl">{disc.name}</h2>
                <p className="font-bold text-gray-800">
                  {disc.manufacturer} - {disc.weight}g
                </p>
              </div>

              <div className="flex gap-4 text-gray-500 text-xs">
                <p>{disc.type}</p>
                <p>
                  {disc.speed} | {disc.glide} | {disc.turn} | {disc.fade}
                </p>
              </div>
            </div>
          </Link>
        </div>

        <div className="flex flex-col justify-center gap-2 p-4">
          <input
            className="block"
            type="checkbox"
            id="in-bag"
            checked={disc.in_bag}
            onChange={toggleInBag}
          />
          <label className="block text-gray-500 text-xs" htmlFor="in-bag">
            IN BAG
          </label>
        </div>
      </div>
    </li>
  )
}
