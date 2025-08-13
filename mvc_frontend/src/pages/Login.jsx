import { useState } from "react"
import { useNavigate } from "react-router-dom"
import Navbar from '../components/Navbar'
import HeroImg from '../assets/cafe.jpg'
import './css/login.css'
import axios from "axios"
import { useAuth } from '../context/AuthContext';

export default function Signup() {
    const navigate = useNavigate()
    const url = import.meta.env.VITE_URL
    const [error, setError] = useState('')

    const [FormData, setFormData] = useState({
        email: '',
        password: ''
    })

    function handleFormDataChange (e) {
        setFormData({...FormData, [e.target.name]: e.target.value})
    }

    const {user, login} = useAuth()

    const handleSubmit = async (e) => {
        try {
            e.preventDefault()
            localStorage.clear()
            const res = await axios.post(`${url}/api/v1/client/login`, FormData)
            const data = res.data 
            localStorage.setItem('token', JSON.stringify(data.token)) 
            setFormData({email: '', password: ''})
            

            login(data.name, data.role, data.user_id)
                
            switch (data.role) {
                case 'customer':
                    navigate('/home')
                    break;
                case 'chef':
                    navigate('/chef')
                    break;
                case 'administrator':
                    navigate('/admin')
                    break;
                default:
                    break;
                }

        }
        catch (err) {
            if (err.response && err.response.data && err.response.data.error) {
                setError(err.response.data.error)
            } else {
                setError('Login failed. Please try again.')
            }
        }
    }


    return (

<>
<Navbar/>

<div className="container-fluid  hero-bg" style={{ backgroundImage: `url(${HeroImg})`, height: "100vh"}}>
    <div className="container-fluid glass">

        <div className="container w-25 mt-5" >
            <div className="mx-auto">
                    <form onSubmit={handleSubmit} className="container p-3 shadow rounded-5">
            <h3 className="mb-4 text-center">Login</h3>
{error && <div className="alert alert-danger text-center">{error}</div>}

            <div className="mb-3">
                <label htmlFor="email" className="form-label">Email address</label>
                <input
                type="email"
                className="form-control"
                id="email"
                name="email"
                placeholder="Enter email"
                value={FormData.email}
                onChange={handleFormDataChange}
                required
                />
            </div>

            <div className="mb-4">
                <label htmlFor="password" className="form-label">Password</label>
                <input
                type="password"
                className="form-control"
                id="password"
                name="password"
                placeholder="Enter password"
                value={FormData.password}
                onChange={handleFormDataChange}
                required
                />
            </div>

            <button type="submit" className="btn btn-primary w-100">Login</button>
                    </form>
            </div>
        </div>

        <div>
    </div>
    </div>

</div>
</>
    )
}