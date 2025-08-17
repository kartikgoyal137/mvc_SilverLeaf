import React, { useEffect, useState } from 'react';
import Navbar from '../components/Navbar'
import HeroImg from '../assets/cafe.jpg'
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

export default function Checkout() {
    
    const nav = useNavigate()
    const url = import.meta.env.VITE_URL
    const [cartItems, setCartItems] = useState([]);
    const [total, setTotal] = useState(0)
    const [showConfirmationModal, setShowConfirmationModal] = useState(false);

    const [formData, setFormData] = useState({
        table_no : "",
        description : "",
        tip : ""
    });

    const orderID = JSON.parse(localStorage.getItem('order_id'))
    const myToken = JSON.parse(localStorage.getItem('token'))
    const userID = JSON.parse(localStorage.getItem('user_id'))
    

    const calculateSubtotal = async () => {
        if (!orderID) {
            console.error("No Order ID found. Cannot fetch cart items.");
            return; 
        }
        const res = await axios.get(`/api/v1/payments/total/${orderID}`, {headers: {Authorization: `${myToken}`}})
        const data = res.data || 0
        console.log(data)
        const ans =  parseInt(data, 10)*1.08
        return ans.toFixed(2)
    };
    

    const handleInputChange = (e) => {
        const { id, value } = e.target;
        setFormData(prevData => ({ ...prevData, [id]: value }));
    };
    

    const handleSubmit = async (e) => {
        e.preventDefault();
        if(cartItems.length === 0) {
            alert("please order something")
            nav('/menu')
            return
        }
        const res1 =  await axios.post(`/api/v1/orders/place`, {"order_id" : orderID, "table_no" : parseInt(formData.table_no, 10), "tip" : parseInt(formData.tip,10), "instructions" : formData.description}, {headers: {Authorization: `${myToken}`}})
        const data1 = res1.data
        const res2 =  await axios.post(`/api/v1/payments/new`, {"order_id" : orderID, "user_id" : parseInt(userID,10), "food_total" : total, "tip" : parseInt(formData.tip,10)}, {headers: {Authorization: `${myToken}`}})
        const data2 = res2.data
        localStorage.removeItem('order_id')
        setShowConfirmationModal(true);
    };

    

    useEffect(() => {

        const fetch = async () => {
            if (!orderID) {
            console.error("No Order ID found. Cannot fetch cart items.");
            return;
        }
            const res = await axios.get(`/api/v1/cart/get/${orderID}`, {headers : {Authorization : `${myToken}`}})
            let data = res.data
            if(data===null) {
                data=[]
            }
            setCartItems(data)
            console.log(data)
        }
        fetch()
    }, [])

    useEffect( () => {
        const fetchSubtotal = async () => {
        const tp = await calculateSubtotal();
        const ans = parseInt(tp, 10)
        setTotal(ans);
    };

    fetchSubtotal();
    }, [cartItems])

    return (
        <>
        <Navbar/>
        <div className="container-fluid hero-bg" style={{ backgroundImage: `url(${HeroImg})`, height: "100vh" }}>
            <div className="hero-overlay glass" style={{height: '100vh'}}></div>
        <div className="container py-5">
            <main>
                <div className="py-5 text-center">
                    <h2>Checkout</h2>
                    <p className="lead">
                        Please fill out the form below to complete your purchase.
                    </p>
                </div>

                <div className="row g-5">
                    <div className="col-md-5 col-lg-4 order-md-last">
                        <h4 className="d-flex justify-content-between align-items-center mb-3">
                            <span className="text-success">Your cart (8% tax)</span>
                            <span className="badge bg-success rounded-pill">{cartItems.length}</span>
                        </h4>
                        <ul className="list-group mb-3 rounded-4">
                            {cartItems.map(item => (
                                <li key={item.id} className=" list-group-item d-flex justify-content-between lh-sm">
                                    <div>
                                        <h6 className="my-0 fs-4">{item.name}</h6>
                                    <small className="text-muted fs-6">Quantity: {item.quantity}</small>
                                    </div>
                                    <span className="text-muted fs-4">${(item.price * item.quantity).toFixed(2)}</span>
                                </li>
                            ))}
                            <li className="fs-4 py-4 list-group-item d-flex justify-content-between">
                                <span>Total (USD)</span>
                                <strong>${total}</strong>
                            </li>
                        </ul>
                    </div>

                    <div className="col-md-7 col-lg-8">
                        <h4 className="mb-3 text-success">Billing Information</h4>
                        <form className="needs-validation px-4 py-4" noValidate onSubmit={handleSubmit}>
                            <div className="row g-3">
                                <div className="col-12">
                                    <label htmlFor="firstName" className="form-label fs-5">OrderID</label>
                                    <input disabled type="text" className="form-control fs-4" id="firstName" value={orderID} onChange={handleInputChange} required />
                                </div>

                                <div className="col-6">
                                    <label htmlFor="tip" className="form-label fs-5">Tip</label>
                                    <input required type="number" className="form-control fs-4" id="tip" value={formData.tip} onChange={handleInputChange}  />
                                </div>
                                
                                <div className="col-6">
                                    <label htmlFor="table_no" className="form-label fs-5">Table No</label>
                                    <input required type="number" className="form-control fs-4" id="table_no" value={formData.table_no} onChange={handleInputChange}  />
                                </div>
                                
                                
                                <div className="col-12">
                                    <label htmlFor="description" className="form-label fs-5">Message for the chef<span className="text-muted fs-6"> (Optional)</span></label>
                                    <textarea className="form-control fs-4 " id="description" rows="3" placeholder="I don't like salt" value={formData.description} onChange={handleInputChange}></textarea>
                                    
                                </div>
                            </div>

                            <button onClick={handleSubmit} className="w-50 mt-4 btn btn-primary btn-lg" type="submit">Continue to checkout</button>
                        </form>
                    </div>
                </div>
            </main>

            {showConfirmationModal && (
                <div className="modal fade show" style={{ display: 'block' }} tabIndex="-1">
                    <div className="modal-dialog modal-dialog-centered">
                        <div className="modal-content">
                            <div className="modal-header">
                                <h5 className="modal-title">Thank You For Your Order!</h5>
                                <button type="button" className="btn-close" onClick={() => setShowConfirmationModal(false)}></button>
                            </div>
                            <div className="modal-body">
                                <p>Your order has been successfully placed.</p>
                                <hr />
                            </div>
                            <div className="modal-footer">
                                <button type="button" className="btn btn-secondary" onClick={() => {setShowConfirmationModal(false); nav('/home')}}>Close</button>
                            </div>
                        </div>
                    </div>
                </div>
            )}
            {showConfirmationModal && <div className="modal-backdrop fade show"></div>}
        </div>
        </div>
        </>
    );
}
