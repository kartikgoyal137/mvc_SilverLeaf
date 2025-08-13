import React, { useEffect } from "react";
import { useState } from "react";
import NavbarChef from "../components/NavbarChef";
import axios from "axios";

export default function Chef() {
    const url = import.meta.env.VITE_URL
    const [orders, setOrders] = useState([])
    const myToken = JSON.parse(localStorage.getItem('token'))


    useEffect(() => {
    const fetchOrdersAndProducts = async () => {
        try {
            const res = await axios.get(`${url}/api/v1/orders/chef/active`, { headers: { Authorization: `${myToken}` } });
            const initialOrders = res.data || []; 

            const populatedOrdersPromises = initialOrders.map(async (order) => {
                const productsRes = await axios.get(`${url}/api/v1/cart/get/${order.order_id}`, { headers: { Authorization: `${myToken}` } });
                
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

        fetchOrdersAndProducts();
    }, []);

    return (
        <>
        <NavbarChef/>
        <div className="container">
            <div className="row mt-5">
                <div className="btn btn-warning mx-2 col-1">OrderID</div>
                <div className="btn btn-warning mx-2 col-2">Dish / Quantity</div>
                <div className="btn btn-warning mx-2 col-2">Timestamp</div>
                <div className="btn btn-warning mx-2 col-2">Instructions</div>
                <div className="btn btn-warning mx-2 col-1">Table no</div>
                <div className="btn btn-warning mx-2 col-2">Status</div>
            </div>
            {orders.map(m => {
                return (
                    <div className="mx-2 row my-3 ">
                <div className="mx-2 col-1">{m.order_id}</div>
                <div className="mx-2 col-2">{m.products.map(t => 
                        <>
                        <p className="m-0 p-0">{t.name} / {t.quantity}</p>
                        </>
                )
                }</div>
                <div className="mx-2 col-2">{m.created_at}</div>
                <div className="mx-2 col-2">{m.instructions}</div>
                <div className="mx-2 col-1">{m.table_no}</div>
                <div className="btn btn-success mx-2 col-2">{m.status}</div>
                </div>
                )
            })}
        </div>
        </>
    )
}