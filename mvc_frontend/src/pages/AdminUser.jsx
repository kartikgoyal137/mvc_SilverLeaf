import React, { useEffect } from "react";
import { useState } from "react";
import NavbarAdmin from "../components/NavbarAdmin";
import axios from "axios";

export default function AdminUser() {
    const url = import.meta.env.VITE_URL
    const [orders, setOrders] = useState([])
    const myToken = JSON.parse(localStorage.getItem('token'))


    useEffect(() => {
    const fetchOrdersAndProducts = async () => {
        try {
            const res = await axios.get(`${url}/api/v1/client/admin/all`, { headers: { Authorization: `${myToken}` } });
            const payments = res.data || []; 

            setOrders(payments);

        } catch (error) {
            console.error("Failed to fetch orders or products:", error);
        }
    };

        fetchOrdersAndProducts();
    }, []);

    return (
        <>
        <NavbarAdmin/>
        <div className="container">
            <div className="row mt-5">
                <div className="btn btn-warning mx-2 col-1">UserID</div>
                <div className="btn btn-warning mx-2 col-2">Name</div>
                <div className="btn btn-warning mx-2 col-2">Contact</div>
                <div className="btn btn-warning mx-2 col-3">Email</div>
                <div className="btn btn-warning mx-2 col-2">Role</div>
            </div>
            {orders.map(m => {
                return (
                <div className="mx-2 row my-3 ">
                    <div className="mx-2 col-1">{m.user_id}</div>
                    <div className="mx-2 col-2">{m.first_name} {m.last_name}</div>
                    <div className="mx-2 col-2">{m.contact}</div>
                    <div className="mx-2 col-3">{m.email}</div>
                    <div className="btn btn-success mx-2 col-2">{m.role}</div>
                </div>
                )
            })}
        </div>
        </>
    )
}