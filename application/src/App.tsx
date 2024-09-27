import { HashRouter, Routes, Route } from 'react-router-dom'
import LoginPage from './pages/LoginPage'

function App() {
  return (
    <HashRouter>
      <Routes>
        <Route path='/' element={<LoginPage />} />
      </Routes>
    </HashRouter>
  )
}

export default App
