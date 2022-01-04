import { Route, Routes } from "react-router-dom"
import AppLayout from "./components/layout/app/AppLayout"
import SiteLayout from "./components/layout/site/SiteLayout"
import { UserProvider } from "./context/UserContext"
import Bag from "./pages/app/Bag"
import More from "./pages/app/More"
import Profile from "./pages/app/Profile"
import Rounds from "./pages/app/Rounds"
import Home from "./pages/Home"
import Login from "./pages/Login"
import SignUp from "./pages/SignUp"

export default function App() {
  return (
    <UserProvider>
      <Routes>
        {/* Site routes */}
        <Route element={<SiteLayout />}>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<SignUp />} />
        </Route>

        {/* App Routes */}
        <Route element={<AppLayout />}>
          <Route path="/rounds" element={<Rounds />} />
          <Route path="/profile" element={<Profile />} />
          <Route path="/bag" element={<Bag />} />
          <Route path="/more" element={<More />} />
        </Route>
      </Routes>
    </UserProvider>
  )
}
