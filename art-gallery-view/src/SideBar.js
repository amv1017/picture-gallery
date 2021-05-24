import React, { useState, useEffect } from "react";
import { Link } from 'react-router-dom'
const SideBar = () => {


    const [authors, setAuthors]  = useState([])
    

    
    useEffect(() => {
        const getPaintings = async () => {
          setAuthors(await fetchPaintings())
        }
        getPaintings();
    }, [])
    
    const fetchPaintings = async () => {
        const res = await fetch('http://localhost:8080/authors',{method:"GET"});
        const data = await res.json();
        return data;
    }
    
    

    return (
        <div>
        <div className="flex flex-col items-center">
                    { authors.map((i) => (
            <div>
                <a href={'/author/'+i.id}>{i.name}</a>
            </div>
            )) }
        </div>

        <div className="mb-3"></div>

        <div className="flex flex-col items-center">
            <a href={'/genre/1'}>Landscape</a>
            <a href={'/genre/2'}>Portrait</a>
            <a href={'/genre/3'}>Still Life</a>
        </div>
        </div>
    );
}
export default SideBar;
