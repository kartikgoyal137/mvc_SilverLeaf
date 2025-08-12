import React, { useEffect, useState } from 'react';
import Navbar from '../components/Navbar'
import HeroImg from '../assets/cafe.jpg'
import './css/checkout.css'
import axios from 'axios';

export default function Checkout() {
   
    const url = import.meta.env.VITE_URL
    const [cartItems, setCartItems] = useState([]);
    const [total, setTotal] = useState(0)
    const [showConfirmationModal, setShowConfirmationModal] = useState(false);

    const [formData, setFormData] = useState({
        table_no : 0,
        description : "",
        tip : 0
    });
    

    const calculateSubtotal = () => {
        return cartItems.reduce((total, item) => total + item.price * item.quantity, 0);
    };
    

    const handleInputChange = (e) => {
        const { id, value } = e.target;
        setFormData(prevData => ({ ...prevData, [id]: value }));
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        setShowConfirmationModal(true);
    };

    const orderID = JSON.parse(localStorage.getItem('order_id'))
    const myToken = JSON.parse(localStorage.getItem('token'))

    useEffect(() => {
        const fetch = async () => {
            const res = await axios.get(`${url}/api/v1/cart/get/${orderID}`, {headers : {Authorization : `${myToken}`}})
            const data = res.data
            setCartItems(data)
            const total = calculateSubtotal()*1.08;
            setTotal(total)
            console.log(data)
        }
        fetch()
    }, [])

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
                            <span className="text-primary">Your cart</span>
                            <span className="badge bg-primary rounded-pill">{cartItems.length}</span>
                        </h4>
                        <ul className="list-group mb-3">
                            {cartItems.map(item => (
                                <li key={item.id} className="list-group-item d-flex justify-content-between lh-sm">
                                    <div>
                                        <h6 className="my-0">{item.name}</h6>
                                        <small className="text-muted">Quantity: {item.quantity}</small>
                                    </div>
                                    <span className="text-muted">${(item.price * item.quantity).toFixed(2)}</span>
                                </li>
                            ))}
                            <li className="list-group-item d-flex justify-content-between">
                                <span>Total (USD)</span>
                                <strong>${total.toFixed(2)}</strong>
                            </li>
                        </ul>
                    </div>

                    <div className="col-md-7 col-lg-8">
                        <h4 className="mb-3">Billing Information</h4>
                        <form className="needs-validation px-4 py-4" noValidate onSubmit={handleSubmit}>
                            <div className="row g-3">
                                <div className="col-12">
                                    <label htmlFor="firstName" className="form-label fs-5">OrderID</label>
                                    <input disabled type="text" className="form-control fs-4" id="firstName" value={orderID} onChange={handleInputChange} required />
                                </div>

                                <div className="col-6">
                                    <label htmlFor="Tip" className="form-label fs-5">Tip</label>
                                    <input required type="number" className="form-control fs-4" id="Tip" value={formData.tip} onChange={handleInputChange}  />
                                </div>
                                
                                <div className="col-6">
                                    <label htmlFor="table_no" className="form-label fs-5">Table No</label>
                                    <input required type="number" className="form-control fs-4" id="table_no" value={formData.table_no} onChange={handleInputChange}  />
                                </div>
                                
                                
                                <div className="col-12">
                                    <label htmlFor="description" className="form-label fs-5">Message for the chef<span className="text-muted fs-6"> (Optional)</span></label>
                                    <textarea className="form-control fs-4 " id="description" rows="3" placeholder="I don't like salt" value={formData.description} onChange={handleInputChange}></textarea>
                                    <button type="button" className="btn btn-outline-primary btn-sm mt-2">
                                    </button>
                                </div>
                            </div>

                            <button className="w-50 btn btn-primary btn-lg" type="submit">Continue to checkout</button>
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
                                <p>Your order has been successfully placed. A confirmation email has been sent to {formData.email}.</p>
                                <hr />
                            </div>
                            <div className="modal-footer">
                                <button type="button" className="btn btn-secondary" onClick={() => setShowConfirmationModal(false)}>Close</button>
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
