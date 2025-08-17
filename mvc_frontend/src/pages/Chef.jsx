import React, { useEffect } from "react";
import { useState } from "react";
import NavbarChef from "../components/NavbarChef";
import axios from "axios";

export default function Chef() {
    const url = import.meta.env.VITE_URL
    const [orders, setOrders] = useState([])
    const myToken = JSON.parse(localStorage.getItem('token'))

    async function ChangeStatus(m) {
        let newStatus;

        switch (m.status) {
        case 'Yet to start':
            newStatus = 'Cooking';
            break;
        case 'Cooking':
            newStatus = 'Completed';
            break;
        case 'Completed':
            newStatus = 'Yet to start';
            break;
        default:
            newStatus = 'Yet to start'; 
            break;
        }
        const res = await axios.post(`/api/v1/orders/chef/status`, {"order_id" : m.order_id, "status" : `${newStatus}`},{ headers: { Authorization: `${myToken}` } });
        const changed = res.data || []; 
        fetchOrdersAndProducts()
    }

    const fetchOrdersAndProducts = async () => {
        try {
            const res = await axios.get(`/api/v1/orders/chef/active`, { headers: { Authorization: `${myToken}` } });
            const initialOrders = res.data || []; 

            const populatedOrdersPromises = initialOrders.map(async (order) => {
                const productsRes = await axios.get(`/api/v1/cart/get/${order.order_id}`, { headers: { Authorization: `${myToken}` } });
                
                return {
                    ...order,
                    products: productsRes.data || []
                };
            });

            const populatedOrders = await Promise.all(populatedOrdersPromises);

            setOrders(populatedOrders);

        } catch (error) {
            console.error("Failed to fetch orders or products:", error);
        }
    };

    useEffect(() => {
        fetchOrdersAndProducts();
    }, []);

    return (
        <>
        <NavbarChef/>
        <div className="container">
            <div className="row mt-5">
                <div className="btn btn-warning mx-1 col-1">OrderID</div>
                <div className="btn btn-warning mx-1 col-2">Dish / Quantity</div>
                <div className="btn btn-warning mx-1 col-2">Timestamp</div>
                <div className="btn btn-warning mx-1 col-2">Instructions</div>
                <div className="btn btn-warning mx-1 col-1">Table no</div>
                <div className="btn btn-warning mx-1 col-2">Status</div>
            </div>
            {orders.map(m => {
                return (
                    <div className="mx-1 row my-3 border-2">
                <div className="mx-1 col-1">{m.order_id}</div>
                <div className="mx-1 col-2">{m.products.map(t => 
                        <>
                        <p className="m-0 p-0">{t.name} / {t.quantity}</p>
                        </>
                )
                }</div>
                <div className="mx-1 col-2">{m.created_at}</div>
                <div className="mx-1 col-2">{m.instructions}</div>
                <div className="mx-1 col-1">{m.table_no}</div>
                <div className="btn btn-primary col-2 text-center">{m.status}</div>
                <button onClick={() => {ChangeStatus(m)}} className="btn btn-success mx-1 col-1">UP</button>
                </div>
                )
            })}
        </div>
        </>
    )
}