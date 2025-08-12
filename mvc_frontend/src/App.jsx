import Home from './pages/Home.jsx'
import Menu from './pages/Menu.jsx'
import Login from './pages/Login.jsx'
import Signup from './pages/Signup.jsx'
import { Routes, Route, Link } from 'react-router-dom';
import './App.css'

function App() {
  return (
    <>
    <Routes>
      <Route path='/home' element={<Home/>} />
      <Route path='/menu' element={<Menu/>} />
      <Route path='/login' element={<Login/>} />
      <Route path='/signup' element={<Signup/>} />
    </Routes>
    </>
  )
}

export default App
