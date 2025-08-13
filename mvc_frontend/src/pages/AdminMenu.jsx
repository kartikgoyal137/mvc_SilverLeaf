import NavbarAdmin from "../components/NavbarAdmin"
import { useNavigate } from "react-router-dom"
import { useState } from "react"

export default function AdminMenu () {

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

    function handleSubmit (e) {
        
    }


    return (
        <>
        <NavbarAdmin/>
        <div className="container">
            <div className="row mt-5">
                <div className="col-6">
                    <form onSubmit={handleSubmit} className="container p-3 shadow rounded-5">
            <h3 className="mb-4 text-center">Add Item</h3>
{error && <div className="alert alert-danger text-center">{error}</div>}

            <div className="mb-3">
                <label htmlFor="email" className="form-label">ProductID</label>
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
            <div className="mb-3">
                <label htmlFor="email" className="form-label">CategoryID</label>
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
            <div className="mb-3">
                <label htmlFor="email" className="form-label">Product Name</label>
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

            <div className="mb-3">
                <label htmlFor="email" className="form-label">Ingredients</label>
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
                <label htmlFor="password" className="form-label">Image URL</label>
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


        <div className="col-6">
                    <form onSubmit={handleSubmit} className="container p-3 shadow rounded-5">
            <h3 className="mb-4 text-center">Remove Item</h3>
{error && <div className="alert alert-danger text-center">{error}</div>}

            <div className="mb-3">
                <label htmlFor="email" className="form-label">ProductID</label>
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
                <label htmlFor="password" className="form-label">Reason</label>
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
        </div>
        </>
    )
}