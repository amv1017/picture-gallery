import React, { useState, useEffect } from 'react'
import Screen from './Screen'

const Paintings = (props) => {

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
        const res = await fetch(`http://localhost:8080/${props.match.params.section}/${props.match.params.id}`,{method:"GET"});
        const data = await res.json();
        return data;
    }
    



    return(
        <div class="w-full max-w-6xl mx-auto px-4 border-2 border-grey-600">
        
        <Screen gallery={paintings} />
        
        </div>
    )
}

export default Paintings
