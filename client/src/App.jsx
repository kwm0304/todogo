import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import './App.css'
import Nav from './pages/Nav'
import Signup from './pages/Signup'
import Login from './pages/Login'
import ListPage from './pages/ListPage'

function App() {
  

  return (
    <>
    
     <Router>
      <Nav />
      
        <Routes>
          
            <Route path="/" element={<ListPage />} />
            <Route path="/signup" element={<Signup />} />
            <Route path="/login" element={<Login />} />
          
        </Routes>
        
     </Router>
    </>
  )
}

export default App
