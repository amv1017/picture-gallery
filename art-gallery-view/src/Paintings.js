import React, { useState, useEffect } from 'react'
import PaintingView from './PaintingView'

const Paintings = ({path}) => {

    const [paintings, setPaintings]  = useState(
        []
    )
    
    useEffect(() => {
        const getPaintings = async () => {
          setPaintings(await fetchPaintings())
        }
        getPaintings();
    }, [])
    
    const fetchPaintings = async () => {
        console.log("Path: ",path)
        const res = await fetch(path,{method:"GET"});
        const data = await res.json();
        return data;
    }
    



    return(
        <div class="w-full max-w-6xl mx-auto px-4 border-2 border-grey-600">
        
        <PaintingView gallery={paintings} />
        
        </div>
    )
}

export default Paintings
