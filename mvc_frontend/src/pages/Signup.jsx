import { useState } from "react"
import { useNavigate } from "react-router-dom"
import Navbar2 from '../components/NavbarLogin'
import HeroImg from '../assets/cafe.jpg'
import axios from "axios"

export default function Signup() {
    const navigate = useNavigate()
    const url = import.meta.env.VITE_URL
    const [error, setError] = useState('')
    const [FormData, setFormData] = useState({
        first_name: '',
        last_name: '',
        contact: '',
        email: '',
        password: '',
    })

    function handleFormDataChange (e) {
        const fName = e.target.name
        setFormData({...FormData, [fName] : e.target.value})
    }

    const handleSubmit = async (e) => {
        try {
            e.preventDefault()
            const res = await axios.post(`/api/v1/client/signup`, FormData)
            const data = res.data
            setFormData({first_name: '',last_name: '',contact: '',email: '',password: '',})
            navigate('/login')
        }
        catch (err) {
            if (err.response && err.response.data && err.response.data.error) {
                setError(err.response.data.error)
            } else {
                setError('Signup failed. Please try again.')
            }
        }
    }

    return (

<>
<Navbar2/>

<div className="container-fluid  hero-bg" style={{ backgroundImage: `url(${HeroImg})`, height: "100vh"}}>
    <div className="container-fluid glass">

        <div className="container" >
            <div className="mx-auto w-50">
                <form onSubmit={handleSubmit} className="border mt-5 p-4 p-md-5">
    <h3 className="mb-4 text-center">Sign Up</h3>

    {error && <div className="alert alert-danger text-center">{error}</div>}

    <div className="row">
        <div className="col-md-6 mb-3">
            <label htmlFor="first_name" className="form-label">First Name</label>
            <input
                type="text"
                className="form-control"
                id="first_name"
                name="first_name"
                placeholder="Enter your first name" 
                value={FormData.first_name}       
                onChange={handleFormDataChange}
                required
            />
        </div>
        <div className="col-md-6 mb-3">
            <label htmlFor="last_name" className="form-label">Last Name</label>
            <input
                type="text"
                className="form-control"
                id="last_name"
                name="last_name"
                placeholder="Enter your last name"
                value={FormData.last_name}       
                onChange={handleFormDataChange}
                required
            />
        </div>
    </div>

    <div className="mb-3">
        <label htmlFor="contact" className="form-label">Contact</label>
        <input
            type="text"
            className="form-control"
            id="contact"
            name="contact"
            placeholder="Enter your Number"
            value={FormData.contact}
            onChange={handleFormDataChange}
            required
        />
    </div>

    <div className="mb-3">
        <label htmlFor="email" className="form-label">Email address</label>
        <input
            type="email"
            className="form-control"
            id="email"
            name="email"
            placeholder="Enter your email"
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
            placeholder="Enter your password"
            value={FormData.password}
            onChange={handleFormDataChange}
            required
        />
    </div>

    <button type="submit" className="btn btn-success w-100 mt-3">Sign Up!</button>
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