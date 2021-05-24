import React, { useState, useEffect } from "react";
import { Link } from 'react-router-dom'
const SideBar = () => {


    const [authors, setPaintings]  = useState(
        []
    )
    
    useEffect(() => {
        const getPaintings = async () => {
          setPaintings(await fetchPaintings())
        }
        getPaintings();
    }, [])
    
    const fetchPaintings = async () => {
        const res = await fetch('http://localhost:8080/authors',{method:"GET"});
        const data = await res.json();
        return data;
    }
    



    return (
        <div className="flex flex-col items-center">
                    { authors.map((i) => (
            <div>
                <Link to={'/author/'+i.id}>{i.name}</Link>
            </div>
            )) }
        </div>
    );
}
export default SideBar;
