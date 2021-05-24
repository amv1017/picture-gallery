import React, { useState, useEffect } from 'react'

const Exhibit = ({ museumid }) => {

    const [paintings, setPaintings]  = useState({})

    useEffect(() => {
        const getPaintings = async () => {
          setPaintings(await fetchPaintings())
        }
        getPaintings();
    }, [])

    const fetchPaintings = async () => {
        const res = await fetch(`http://localhost:8080/painting/${museumid}`,{method:'GET'});
        const data = await res.json();
        return data;
    }

    const remove = async () => {
        await fetch(`http://localhost:8080/painting/${museumid}`,{ method:'DELETE'});
        console.log('removing')
    }

    return(
        <div>
            
    <div className="flex flex-col justify-center" >


        <img className="object-contain h-screen w-full" src={paintings.url} />
        <p className="text-blue-900 text-4xl text-center mb-4">{paintings.title}</p>
        <p className="text-blue-900">{paintings.description}</p>

        <div className="flex justify-end">
                <button className="bg-red-400 rounded-md" onClick={() => remove()}>Delete This Exhibit</button>
            </div>

    </div>
    </div>
    )
}

export default Exhibit
