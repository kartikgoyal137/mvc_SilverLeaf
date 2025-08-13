import React from "react"
import { useAuth } from "../context/AuthContext"
import { useNavigate } from "react-router-dom"

export default function NavbarChef() {
    const {logout} = useAuth()
    const navigate = useNavigate()

    const Logout = async () => {
    localStorage.removeItem('token')
    localStorage.removeItem('order_id')
    logout()
    navigate('/login')
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
                <button onClick={Logout} className=" ms-auto btn btn-outline-success mx-2" type="submit">Logout</button>
            </div>
        </div>
        </nav>
        </>
    )
}