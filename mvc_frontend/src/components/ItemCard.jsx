import { useState } from "react";


export default function Card(props) {

    const [count, setCount] = useState(0);
    const handleIncrement = () => {
        setCount(prevCount => prevCount + 1);
    };
    const handleDecrement = () => {
        if (count > 0) {
            setCount(prevCount => prevCount - 1);
        }
    };

    return (
        <>
        <div className="card mx-5 my-4 rounded-4" style={{width: '20em'}}>
        <img src={props.img} className="rounded-4 mt-2 card-img-top" alt="..."/>
        <div className="card-body">
            <h5 className="card-title">{props.name}</h5>
            <p className="card-text">Some quick example text to build on the card title and make up the bulk of the card’s content.</p>
        </div>
        <ul className="list-group list-group-flush">
            <li className="list-group-item">Price: <button className="btn ms-2 btn-success">${props.price}</button></li>
            <li className="list-group-item">{props.ing}</li>
            <li className="list-group-item d-flex justify-content-between align-items-center">
                        <button onClick={handleDecrement} className="btn btn-danger btn-sm">-</button>
                        <span className="fw-bold fs-5">{count}</span>
                        <button onClick={handleIncrement} className="btn btn-success btn-sm">+</button>
                    </li>
        </ul>
        <div className="card-body">
            <button className="btn btn-success">Add to Cart</button>
        </div>
        </div>
        </>
    )
}