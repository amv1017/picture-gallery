import PreView from './PreView'
import React, { useState, useEffect } from 'react'

const PaintingView = ({gallery}) => {


    const [showOne, setShowOne] = useState(false)
    const [ind, setInd] = useState(1)

 

    return(
        showOne ? 

        
        <div class="grid grid-cols-1 gap-4">
        <div onClick={() => setShowOne(!showOne)}>
        
            <PreView paint={gallery[ind]} />
            </div>
        </div> 
        
        : 
        
        <div class="grid grid-cols-3 gap-4">
        
        { gallery.map((i) => (
            <div onClick={() => { setShowOne(!showOne) }}>
                <PreView paint={i} onClick={() => { setInd(i) }} />
            </div>
        )) }
        </div>
        

    )
}

export default PaintingView
