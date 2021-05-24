import React, { useState, useEffect } from "react";

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
            <div className="mb-3 text-center bg-blue-200 rounded-md p-3">AUTHORS
        <div className="flex flex-col items-center pb-6">
                    { authors.map((i) => (
            <div>
                <a className="bg-blue-400 rounded-md px-5 mb-4" href={'/author/'+i.id}>{i.name}</a>
            </div>
            )) }
        </div>
        </div>

        <div className="mb-3 text-center bg-green-200 rounded-md p-3">GENRES

        <div className="flex flex-col items-center">
            <a className="bg-green-400 rounded-md py-2 px-5 mb-1" href={'/genre/1'}>Landscape</a>
            <a className="bg-green-400 rounded-md py-2 px-5 mb-1" href={'/genre/2'}>Portrait</a>
            <a className="bg-green-400 rounded-md py-2 px-5" href={'/genre/3'}>Still Life</a>
        </div>
        </div>

        </div>
    );
}
export default SideBar;
