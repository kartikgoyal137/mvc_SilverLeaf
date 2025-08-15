import NavbarAdmin from "../components/NavbarAdmin"
import { useState } from "react"
import axios from "axios"

export default function AdminMenu () {
    const url = import.meta.env.VITE_URL
    const [addError, setAddError] = useState('')
    const [removeError, setRemoveError] = useState('')
    const [addSuccess, setAddSuccess] = useState('')
    const [removeSuccess, setRemoveSuccess] = useState('')
    const myToken = JSON.parse(localStorage.getItem('token'))

    const [addItemData, setAddItemData] = useState({
        product_id: '',
        category_id: '',
        product_name: '',
        price: '',
        image_url: '',
        ingredient_list: ''
    })

    const [removeItemData, setRemoveItemData] = useState({
        product_id: '',
        reason: ''
    })

    function handleAddItemChange (e) {
        setAddItemData({...addItemData, [e.target.name]: e.target.value})
    }

    function handleRemoveItemChange (e) {
        setRemoveItemData({...removeItemData, [e.target.name]: e.target.value})
    }

    async function handleAddItemSubmit (e) {
        e.preventDefault()
        setAddError('')
        setAddSuccess('')
        try {
            const payload = {
                ...addItemData,
                product_id: parseInt(addItemData.product_id, 10),
                category_id: parseInt(addItemData.category_id, 10),
                price: parseFloat(addItemData.price)
            };
            const res = await axios.patch(`/api/v1/menu/add`, payload, {
                headers: { Authorization: `${myToken}` }
            })
            setAddSuccess("Item added successfully!")
            setAddItemData({
                product_id: '',
                category_id: '',
                product_name: '',
                price: '',
                image_url: '',
                ingredient_list: ''
            })
        } catch (err) {
            if (err.response && err.response.data && err.response.data.error) {
                setAddError(err.response.data.error)
            } else {
                setAddError('Failed to add item. Please try again.')
            }
        }
    }

    async function handleRemoveItemSubmit (e) {
        e.preventDefault()
        setRemoveError('')
        setRemoveSuccess('')
        try {
            const res = await axios.delete(`/api/v1/menu/remove/${removeItemData.product_id}`, {
                headers: { Authorization: `${myToken}` }
            })
            setRemoveSuccess("Item removed successfully!")
            setRemoveItemData({ product_id: '', reason: '' })
        } catch (err) {
            if (err.response && err.response.data && err.response.data.error) {
                setRemoveError(err.response.data.error)
            } else {
                setRemoveError('Failed to remove item. Please try again.')
            }
        }
    }

    return (
        <>
        <NavbarAdmin/>
        <div className="container">
            <div className="row mt-5">
                <div className="col-md-6 mb-4">
                    <form onSubmit={handleAddItemSubmit} className="container p-4 shadow rounded-5">
                        <h3 className="mb-4 text-center">Add Item</h3>
                        {addError && <div className="alert alert-danger text-center">{addError}</div>}
                        {addSuccess && <div className="alert alert-success text-center">{addSuccess}</div>}

                        <div className="mb-3">
                            <label htmlFor="product_id_add" className="form-label">Product ID</label>
                            <input
                                type="number"
                                className="form-control"
                                id="product_id_add"
                                name="product_id"
                                placeholder="Enter Product ID"
                                value={addItemData.product_id}
                                onChange={handleAddItemChange}
                                required
                            />
                        </div>
                        <div className="mb-3">
                            <label htmlFor="category_id" className="form-label">Category ID</label>
                            <input
                                type="number"
                                className="form-control"
                                id="category_id"
                                name="category_id"
                                placeholder="Enter Category ID"
                                value={addItemData.category_id}
                                onChange={handleAddItemChange}
                                required
                            />
                        </div>
                        <div className="mb-3">
                            <label htmlFor="product_name" className="form-label">Product Name</label>
                            <input
                                type="text"
                                className="form-control"
                                id="product_name"
                                name="product_name"
                                placeholder="Enter Product Name"
                                value={addItemData.product_name}
                                onChange={handleAddItemChange}
                                required
                            />
                        </div>
                         <div className="mb-3">
                            <label htmlFor="price" className="form-label">Price</label>
                            <input
                                type="number"
                                step="0.01"
                                className="form-control"
                                id="price"
                                name="price"
                                placeholder="Enter Price"
                                value={addItemData.price}
                                onChange={handleAddItemChange}
                                required
                            />
                        </div>
                        <div className="mb-3">
                            <label htmlFor="ingredient_list" className="form-label">Ingredients</label>
                            <input
                                type="text"
                                className="form-control"
                                id="ingredient_list"
                                name="ingredient_list"
                                placeholder="Enter Ingredients (comma-separated)"
                                value={addItemData.ingredient_list}
                                onChange={handleAddItemChange}
                                required
                            />
                        </div>
                        <div className="mb-4">
                            <label htmlFor="image_url" className="form-label">Image URL</label>
                            <input
                                type="text"
                                className="form-control"
                                id="image_url"
                                name="image_url"
                                placeholder="Enter Image URL"
                                value={addItemData.image_url}
                                onChange={handleAddItemChange}
                                required
                            />
                        </div>
                        <button type="submit" className="btn btn-success w-100">Add Item</button>
                    </form>
                </div>

                <div className="col-md-6 mb-4">
                    <form onSubmit={handleRemoveItemSubmit} className="container p-4 shadow rounded-5">
                        <h3 className="mb-4 text-center">Remove Item</h3>
                        {removeError && <div className="alert alert-danger text-center">{removeError}</div>}
                        {removeSuccess && <div className="alert alert-success text-center">{removeSuccess}</div>}

                        <div className="mb-3">
                            <label htmlFor="product_id_remove" className="form-label">Product ID</label>
                            <input
                                type="number"
                                className="form-control"
                                id="product_id_remove"
                                name="product_id"
                                placeholder="Enter Product ID to remove"
                                value={removeItemData.product_id}
                                onChange={handleRemoveItemChange}
                                required
                            />
                        </div>
                        <div className="mb-4">
                            <label htmlFor="reason" className="form-label">Reason</label>
                            <input
                                type="text"
                                className="form-control"
                                id="reason"
                                name="reason"
                                placeholder="Reason for removal (optional)"
                                value={removeItemData.reason}
                                onChange={handleRemoveItemChange}
                            />
                        </div>
                        <button type="submit" className="btn btn-danger w-100">Remove Item</button>
                    </form>
                </div>
            </div>
        </div>
        </>
    )
}
