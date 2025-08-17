import React, { useEffect } from "react";
import { useState } from "react";
import NavbarAdmin from "../components/NavbarAdmin";
import axios from "axios";

export default function AdminUser() {
    const url = import.meta.env.VITE_URL
    const [users, setUsers] = useState([])
    const myToken = JSON.parse(localStorage.getItem('token'))

    async function ChangeRole(m) {
        let newRole;

        switch (m.role) {
        case 'customer':
            newRole = 'chef';
            break;
        case 'chef':
            newRole = 'administrator';
            break;
        case 'administrator':
            newRole = 'customer';
            break;
        default:
            newRole = 'customer';
            break;
        }
        
        const id = String(m.user_id)
        const res = await axios.patch(`/api/v1/client/admin/status/${newRole}/${id}`,{} ,{ headers: { Authorization: `${myToken}` } });
        const changed = res.data || []; 
        fetchUsers()
    }

    const fetchUsers = async () => {
        try {
            const res = await axios.get(`/api/v1/client/admin/all`, { headers: { Authorization: `${myToken}` } });
            const usersdata = res.data || []; 

            setUsers(usersdata);

        } catch (error) {
            console.error("Failed to fetch orders or products:", error);
        }
    };

    useEffect(() => {
    

        fetchUsers();
    }, []);

    return (
        <>
        <NavbarAdmin/>
        <div className="container">
            <div className="row my-5">
                <div className="btn btn-warning mx-1 col-1">UserID</div>
                <div className="btn btn-warning mx-1 col-2">Name</div>
                <div className="btn btn-warning mx-1 col-2">Contact</div>
                <div className="btn btn-warning mx-1 col-3">Email</div>
                <div className="btn btn-warning mx-1 col-2">Role</div>
            </div>
            {users.map(m => {
                return (
                <div className="mx-1 row my-3 border-2 ">
                    <div className="mx-1 col-1">{m.user_id}</div>
                    <div className="mx-1 col-2">{m.first_name} {m.last_name}</div>
                    <div className="mx-1 col-2">{m.contact}</div>
                    <div className="mx-1 col-3">{m.email}</div>
                    <div className="btn btn-primary col-2">{m.role}</div>
                    <button onClick={() => {ChangeRole(m)}} className="btn btn-success mx-1 col-1">UP</button>
                </div>
                )
            })}
        </div>
        </>
    )
}