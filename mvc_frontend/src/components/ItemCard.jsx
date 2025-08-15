import { useState } from "react";
import axios from "axios";
import img from '../assets/table.png'

export default function Card(props) {
    const orderID = JSON.parse(localStorage.getItem('order_id'))
    const url = import.meta.env.VITE_URL
    const myToken = JSON.parse(localStorage.getItem('token'))
    const [count, setCount] = useState(0);
    const handleIncrement = () => {
        setCount(prevCount => prevCount + 1);
    };
    const handleDecrement = () => {
        if (count > 1) {
            setCount(prevCount => prevCount - 1);
        }
    };

    const Add = async () => {
        try
        {const res = await axios.post(`/api/v1/cart/add`,{"order_id": orderID, "product_id": props.product_id, "quantity": count} ,{headers: {Authorization : `${myToken}` }})
        const data = res.data
        console.log(data)}
        catch (err) {
            console.log(err)
        }
    }

    const Remove = async () => {
        try
        {
        const res = await axios.post(`/api/v1/cart/remove`,{"order_id": orderID, "product_id": props.product_id, "quantity": count} ,{headers: {Authorization : `${myToken}` }})
        const data = res.data
        setCount(0)
        console.log(data)
    }
        catch (err) {
            console.log(err)
        }
    }

    return (
        <>
        <div className="card mx-5 my-4 rounded-4" style={{width: '20em'}}>
        <img src={props.img || img} className="rounded-4 mt-2 card-img-top" alt="..."/>
        <div className="card-body">
            <h5 className="card-title">{props.name}</h5>
            <p className="card-text">Some quick example text to build on the card title and make up the bulk of the card’s content.</p>
        </div>
        <ul className="list-group list-group-flush">
            <li className="list-group-item">Price: <button className="btn ms-2 btn-warning">${props.price}</button></li>
            <li className="list-group-item">{props.ing}</li>
            <li className="list-group-item d-flex justify-content-between align-items-center">
                        <button onClick={handleDecrement} className="btn btn-warning btn-sm">-</button>
                        <span className="fw-bold fs-5">{count}</span>
                        <button onClick={handleIncrement} className="btn btn-warning btn-sm">+</button>
                    </li>
        </ul>
        <div className="card-body">
            <button onClick={Add} className="btn btn-success mt-1">Add to Cart</button>
            <button onClick={Remove} className="btn btn-danger ms-3 mt-1">Remove</button>
        </div>
        </div>
        </>
    )
}