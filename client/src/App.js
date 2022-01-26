import { Route, Routes } from "react-router-dom"
import AppLayout from "./components/layout/app/AppLayout"
import SiteLayout from "./components/layout/site/SiteLayout"
import RequireAuth from "./components/RequireAuth"
import { UserProvider } from "./context/UserContext"
import Bag from "./pages/app/Bag"
import More from "./pages/app/More"
import NewScorecard from "./pages/app/NewScorecard"
import Rounds from "./pages/app/Rounds"
import ScorecardDetail from "./pages/app/ScorecardDetail"
import Home from "./pages/Home"
import Login from "./pages/Login"
import SignOut from "./pages/SignOut"
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
          <Route path="/signout" element={<SignOut />} />
        </Route>

        {/* App Routes */}
        <Route
          element={
            <RequireAuth>
              <AppLayout />
            </RequireAuth>
          }
        >
          <Route path="/rounds" element={<Rounds />} />
          <Route path="/bag" element={<Bag />} />
          <Route path="/more" element={<More />} />
          <Route path="/scorecards/new" element={<NewScorecard />} />
          <Route path="/scorecards/:id" element={<ScorecardDetail />} />
        </Route>
      </Routes>
    </UserProvider>
  )
}
