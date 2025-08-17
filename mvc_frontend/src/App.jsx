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
import AdminMenu from './pages/AdminMenu.jsx'
import MyOrders from './pages/MyOrders.jsx'
import ProtectedRoute from './components/ProtectedRoutes.jsx'
import './App.css'

function App() {
  return (
    <>
    <Routes>
      <Route path='/' element={<Login/>} />
      <Route path='/login' element={<Login/>} />
      <Route path='/signup' element={<Signup/>} />

      <Route path='/home' element={<ProtectedRoute allowedRoles={['customer', 'chef', 'administrator']}><Home/></ProtectedRoute>}/>
      <Route path='/menu' element={<ProtectedRoute allowedRoles={['customer', 'chef', 'administrator']}><Menu/></ProtectedRoute>}/>


      <Route path='/checkout' element={<ProtectedRoute allowedRoles={['customer', 'chef', 'administrator']}><Checkout/></ProtectedRoute>}/>
      <Route path='/myorders' element={<ProtectedRoute allowedRoles={['customer', 'chef', 'administrator']}><MyOrders/></ProtectedRoute>}/>

      <Route path='/chef' element={<ProtectedRoute allowedRoles={['chef']}><Chef/></ProtectedRoute>}/>

      <Route path='/admin/order' element={<ProtectedRoute allowedRoles={['administrator']}><AdminOrder/></ProtectedRoute>} />
      <Route path='/admin/user' element={<ProtectedRoute allowedRoles={['administrator']}><AdminUser/></ProtectedRoute>} />
      <Route path='/admin' element={<ProtectedRoute allowedRoles={['administrator']}><AdminUser/></ProtectedRoute>} />
      <Route path='/admin/pay' element={<ProtectedRoute allowedRoles={['administrator']}><AdminPay/></ProtectedRoute>} />
      <Route path='/admin/menu' element={<ProtectedRoute allowedRoles={['administrator']}><AdminMenu/></ProtectedRoute>}/>

    </Routes>
    </>
  )
}

export default App
