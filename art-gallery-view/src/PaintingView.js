import PreView from './PreView'
import Exhibit from './Exhibit'
import React, { useState, useEffect } from 'react'

const PaintingView = ({gallery}) => {


    const [showOne, setShowOne] = useState(false)
    const [ind, setInd] = useState(1)

 

    return(
        showOne ? 

        
        <div class="grid grid-cols-1 gap-4">
        <div onClick={() => setShowOne(!showOne)}>
        
                <Exhibit url={gallery[ind].url} desc={gallery[ind].description} />
            </div>
        </div> 
        
        : 
        
        <div class="grid grid-cols-3 gap-4">
        
        { gallery.map((i) => (
            <div onClick={() => { setShowOne(!showOne) }}>
                <PreView paint={i} onClick={() => { setInd(gallery.indexOf(i)) }} />
                {console.log(gallery)}
            </div>
        )) 

        

        }
        </div>
        

    )
}

export default PaintingView
