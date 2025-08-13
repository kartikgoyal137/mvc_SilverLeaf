import React, { useEffect } from "react";
import { useState } from "react";
import NavbarAdmin from "../components/NavbarAdmin";
import axios from "axios";

export default function AdminPay() {
    const url = import.meta.env.VITE_URL
    const [orders, setOrders] = useState([])
    const myToken = JSON.parse(localStorage.getItem('token'))


    useEffect(() => {
    const fetchOrdersAndProducts = async () => {
        try {
            const res = await axios.get(`${url}/api/v1/payments/admin/all`, { headers: { Authorization: `${myToken}` } });
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
                <div className="btn btn-warning mx-2 col-1">PaymentID</div>
                <div className="btn btn-warning mx-2 col-1">UserID</div>
                <div className="btn btn-warning mx-2 col-1">OrderID</div>
                <div className="btn btn-warning mx-2 col-1">Price</div>
                <div className="btn btn-warning mx-2 col-2">Timestamp</div>
                <div className="btn btn-warning mx-2 col-1">Tip</div>
                <div className="btn btn-warning mx-2 col-2">Status</div>
            </div>
            {orders.map(m => {
                return (
                <div className="mx-2 row my-3 ">
                    <div className="mx-2 col-1">{m.transaction_id}</div>
                    <div className="mx-2 col-1">{m.user_id}</div>
                    <div className="mx-2 col-1">{m.order_id}</div>
                    <div className="mx-2 col-1">{m.food_total}</div>
                    <div className="mx-2 col-2">{m.created_at}</div>
                    <div className="mx-2 col-1">{m.tip}</div>
                    <div className="btn btn-success mx-2 col-2">{m.status}</div>
                </div>
                )
            })}
        </div>
        </>
    )
}