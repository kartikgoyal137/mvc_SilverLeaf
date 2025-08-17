import React, { useEffect, useState } from "react";
import Navbar from "../components/Navbar";
import axios from "axios";
import OrderCard from "../components/OrderCard";


export default function MyOrders() {
    const [orders, setOrders] = useState([]);
    const myToken = JSON.parse(localStorage.getItem('token'));

    useEffect(() => {
        const fetchOrdersAndProducts = async () => {
      
            if (!myToken) return;
            try {
           
                const res = await axios.get(`/api/v1/orders/user`, { headers: { Authorization: `${myToken}` } });
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

        fetchOrdersAndProducts();
    }, [myToken]); 

    return (
        <>
            <Navbar />
            <div className="container py-5">
                <div className="text-center mb-5">
                    <h1 className="display-4">My Orders</h1>
                    <p className="lead text-muted">Here's a history of all your past and current orders.</p>
                </div>

                {orders.length > 0 ? (
                    <div className="row">
                   
                        {orders.map(order => (
                            <OrderCard key={order.order_id} order={order} />
                        ))}
                    </div>
                ) : (
                   
                    <div className="text-center">
                        <p className="fs-5">You haven't placed any orders yet.</p>
                    </div>
                )}
            </div>
        </>
    );
}
