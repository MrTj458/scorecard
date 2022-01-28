import { GrClose } from "react-icons/gr"

export default function Modal({ open, close, children }) {
  if (open) {
    return (
      <div className="fixed flex justify-center items-center top-0 left-0 right-0 bottom-0 bg-gray-700 bg-opacity-90">
        <div className="bg-white rounded-md shadow-md p-6 relative">
          <button onClick={close} className="absolute top-3 right-3">
            <GrClose />
          </button>
          {children}
        </div>
      </div>
    )
  }
  return <></>
}
