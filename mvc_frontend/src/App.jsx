import Home from './pages/Home.jsx'
import Menu from './pages/Menu.jsx'
import Login from './pages/Login.jsx'
import Signup from './pages/Signup.jsx'
import Checkout from './pages/Checkout.jsx'
import Chef from './pages/Chef.jsx'
import { Routes, Route, Link } from 'react-router-dom';
import AdminOrder from './pages/AdminOrder.jsx'
import AdminUser from './pages/AdminUser.jsx'
import AdminPay from './pages/AdminPay.jsx'
import MyOrders from './pages/MyOrders.jsx'
import './App.css'

function App() {
  return (
    <>
    <Routes>
      <Route path='/home' element={<Home/>} />
      <Route path='/menu' element={<Menu/>} />
      <Route path='/' element={<Login/>} />
      <Route path='/signup' element={<Signup/>} />
      <Route path='/checkout' element={<Checkout/>} />
      <Route path='/myorders' element={<MyOrders/>} />
      <Route path='/chef' element={<Chef/>} />
      <Route path='/admin/order' element={<AdminOrder/>} />
      <Route path='/admin/user' element={<AdminUser/>} />
      <Route path='/admin/pay' element={<AdminPay/>} />
    </Routes>
    </>
  )
}

export default App
