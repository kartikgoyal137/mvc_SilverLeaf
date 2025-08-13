import React from "react"
import { useNavigate } from "react-router-dom"
import { useAuth } from "../context/AuthContext"
import { Link } from "react-router-dom"

export default function Navbar() {
    const nav = useNavigate()
    const {logout} = useAuth()

    const Logout = async () => {
    localStorage.clear('token')
    localStorage.clear('order_id')
    logout()
    nav('/login')
    }


    return (    
        <>
        <nav className="navbar navbar-expand-lg ">
            <div className="container-fluid">
                <a className="navbar-brand" href="#">Silver Leaf</a>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                <ul className="navbar-nav me-auto mx-auto mb-2 mb-lg-0">
                    <li className="nav-item">
                    <Link to='/home' className="mx-4 nav-link active" aria-current="page" >Home</Link>
                    </li>
                    <li className="nav-item">
                    <Link to='/menu' className="mx-4 nav-link active">Menu</Link>
                    </li>
                    <li className="nav-item">
                    <Link to='/checkout' className=" mx-4 nav-link active">Cart</Link>
                    </li>
                    <li className="nav-item">
                    <Link to='/myorders' className="mx-4 nav-link active">Orders</Link>
                    </li>
                </ul>
                <button onClick={Logout} className="btn btn-outline-success mx-2" type="submit">Logout</button>
            </div>
        </div>
        </nav>
        </>
    )
}