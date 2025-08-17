import React, { useEffect } from "react";
import { useState } from "react";
import NavbarAdmin from "../components/NavbarAdmin";
import axios from "axios";
import FormattedDate from "../components/FormatDate";

export default function AdminPay() {
    const url = import.meta.env.VITE_URL
    const [payments, setPayments] = useState([])
    const myToken = JSON.parse(localStorage.getItem('token'))


    async function ChangeStatus(m) {
        const status = m.status === 'Pending' ? 'Completed' : 'Pending'
        const res = await axios.patch(`/api/v1/payments/admin/status`, {"order_id" : m.order_id, "status" : `${status}`},{ headers: { Authorization: `${myToken}` } });
        const changed = res.data || []; 
        fetchPayments()
    }

    const fetchPayments = async () => {
        try {
            const res = await axios.get(`/api/v1/payments/admin/all`, { headers: { Authorization: `${myToken}` } });
            const payments = res.data || []; 

            setPayments(payments);

        } catch (error) {
            console.error("Failed to fetch orders or products:", error);
        }
    };

    useEffect(() => {
    

        fetchPayments();
    }, []);

    return (
        <>
        <NavbarAdmin/>
        <div className="container">
            <div className="row my-5">
                <div className="btn btn-warning mx-1 col-1">PaymentID</div>
                <div className="btn btn-warning mx-1 col-1">UserID</div>
                <div className="btn btn-warning mx-1 col-1">OrderID</div>
                <div className="btn btn-warning mx-1 col-1">Price</div>
                <div className="btn btn-warning mx-1 col-2">Timestamp</div>
                <div className="btn btn-warning mx-1 col-1">Tip</div>
                <div className="btn btn-warning mx-1 col-3">Status</div>
            </div>
            {payments.map(m => {
                return (
                <div className="mx-1 row my-3 py-1 d-flex align-items-center border-bottom border-2">
                    <div className="mx-1 col-1">{m.transaction_id}</div>
                    <div className="mx-1 col-1">{m.user_id}</div>
                    <div className="mx-1 col-1">{m.order_id}</div>
                    <div className="mx-1 col-1">{m.food_total}</div>
                    <div className="mx-1 col-2"><FormattedDate isoDateString={m.created_at}/></div>
                    <div className="mx-1 col-1">{m.tip}</div>
                    <div className="btn btn-primary col-2">{m.status}
                    </div>
                    <button onClick={() => {ChangeStatus(m)}} className="btn btn-success mx-1 col-1">UP</button>
                </div>
                )
            })}
        </div>
        </>
    )
}