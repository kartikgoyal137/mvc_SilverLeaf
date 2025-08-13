import React from "react"
import { useNavigate } from "react-router-dom"

export default function NavbarAdmin() {
    const nav = useNavigate()

    function Users() {
        nav('/admin/user')
    }
    function Orders() {
        nav('/admin/order')
    }
    function Payments() {
        nav('/admin/pay')
    }
    function Menu() {
        nav('/admin/menu')
    }

    const Logout = async () => {
    localStorage.removeItem('token')
    localStorage.removeItem('order_id')
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
                <ul className="navbar-nav me-auto mx-auto mb-2 mb-lg-0">
                    <li className="nav-item">
                    <a onClick={Users} className="mx-4 nav-link active" aria-current="page" href="#">Users</a>
                    </li>
                    <li className="nav-item">
                    <a onClick={Orders} className="mx-4 nav-link active" href="#">Orders</a>
                    </li>
                    <li className="nav-item">
                    <a onClick={Payments} className=" mx-4 nav-link active" href="#">Payments</a>
                    </li>
                    <li className="nav-item">
                    <a onClick={Menu} className="mx-4 nav-link active" href="#">Menu</a>
                    </li>
                </ul>
                <button onClick={Logout} className="btn btn-outline-success mx-2" type="submit">Logout</button>
            </div>
        </div>
        </nav>
        </>
    )
}