import axios from 'axios'
import { useEffect } from 'react'
import { useState } from 'react'
import Navbar from "../components/Navbar"
import HeroImg from '../assets/cafe.jpg'
import Card from '../components/ItemCard'

export default function  Menu() {
    const url = import.meta.env.VITE_URL
    const [itemV, setItemV] = useState(0)
    const [category, setCategory] = useState(1)
    const [menu, setMenu] = useState([])
    const [categories, setCategories] = useState([]);
    const myToken = JSON.parse(localStorage.getItem('token'))

    useEffect(() => {
        const fetch = async () => {
            const response = await axios.get(`/api/v1/menu/cat/${category}`)
            const data = response.data
            setMenu(data)
        }
        fetch()
    }, [category])

    function handleSetCat(e){
        setCategory(e)
    }

    async function fetchCategories() {
        const response = await axios.get(`/api/v1/menu/cat/all`)
        const data = await response.data
        return data
    }
    

    useEffect(() => {
        const fetch = async () => {
            const data = await fetchCategories()
            setCategories(data)
        }
        fetch()
        const orderID = JSON.parse(localStorage.getItem('order_id'))
        if(orderID !== null) setItemV(1)
        
    }, [])

    useEffect(() => {
        const btn = document.querySelector('.start')
        if(itemV===1) {
            btn.style.visibility = "hidden"
        }
    }, [itemV])

    

    async function StartOrder() {
        try {
            
            setItemV(1)
            const res = await axios.post(`/api/v1/orders/start`,{}, {headers: {Authorization : `${myToken}` }})
            const data = res.data
            localStorage.setItem('order_id', JSON.stringify(data.order_id))
            
        }
        catch (err) {
            console.log(err)
        }
        
    }

    

    return (
        <>
        <Navbar/>
        <div className="container-fluid hero-bg" style={{backgroundImage: `url(${HeroImg})`}}>
            
            <div className="hero-overlay h-100 mb-5"></div>    
            <div className="container text-center d-flex flex-column align-items-center justify-content-center h-75">
                <h1 className="pt-5 cormorant heading display-1 fw-bold w-75">Our Menu</h1>
            <p className="fs-3 mb-5 pt-2 w-75">Experience a symphony of flavours crafted with passion and the finest ingredients</p>
            <button onClick={StartOrder} className='btn btn-primary my-5 start fs-4'>Start Order</button>
            </div>
            
            </div>

        <div className="container d-flex justify-content-center my-5">
            {categories.map(cat => {
        return (
            <button className='btn btn-warning fs-5 fw-semibold mx-3' onClick={() => {handleSetCat(cat.category_id)}} key={cat.category_id}>{cat.category_name}</button>
        )
        })}
        </div>
        

        <div className="container">
            <div className="row menu-item">
        {menu.map((item)=> {
                    return (
                        <Card className="menu-item" product_id={item.product_id} key={item.product_id} ing={item.ingredient_list} price={item.price} img={item.image_url} name={item.product_name}/>
                    )
            })}
            </div>
        </div>
        
        
        </>
    )
}