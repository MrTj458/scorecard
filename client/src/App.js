import { Route, Routes } from "react-router-dom"
import AppLayout from "./components/layout/app/AppLayout"
import SiteLayout from "./components/layout/site/SiteLayout"
import RequireAuthRoute from "./components/RequireAuthRoute"
import { UserProvider } from "./context/UserContext"
import BagPage from "./pages/app/BagPage"
import MorePage from "./pages/app/MorePage"
import ScorecardDetailPage from "./pages/app/scorecard/ScorecardDetailPage"
import ScorecardNewPage from "./pages/app/scorecard/ScorecardNewPage"
import ScorecardsListPage from "./pages/app/scorecard/ScorecardsListPage"
import HomePage from "./pages/HomePage"
import LoginPage from "./pages/LoginPage"
import SignOutPage from "./pages/SignOutPage"
import SignUpPage from "./pages/SignUpPage"

export default function App() {
  return (
    <UserProvider>
      <Routes>
        {/* Site routes */}
        <Route element={<SiteLayout />}>
          <Route path="/" element={<HomePage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/signup" element={<SignUpPage />} />
          <Route path="/signout" element={<SignOutPage />} />
        </Route>

        {/* App Routes */}
        <Route
          element={
            <RequireAuthRoute>
              <AppLayout />
            </RequireAuthRoute>
          }
        >
          <Route path="/app/scorecards" element={<ScorecardsListPage />} />
          <Route path="/app/scorecards/new" element={<ScorecardNewPage />} />
          <Route path="/app/scorecards/:id" element={<ScorecardDetailPage />} />
          <Route path="/app/bag" element={<BagPage />} />
          <Route path="/app/more" element={<MorePage />} />
        </Route>
      </Routes>
    </UserProvider>
  )
}
