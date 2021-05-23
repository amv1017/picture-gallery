import React, {useState,useEffect} from 'react'
import PaintingView from './PaintingView';

function App() {

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
    const res = await fetch('http://localhost:8080/paintings',{method:"GET"});
    const data = await res.json();
    return data;
}

  return (
    <div className="container mx-auto">
      

      <h1 className="text-3xl">
        Paintings
      </h1>
      {paintings.length > 0 ?
        (<PaintingView gallery={paintings} />) :
        ('No exhibits found')

      }

    </div>
  );
}

export default App
