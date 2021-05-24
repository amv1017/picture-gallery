import React, { useState, useEffect } from 'react'
import Paintings from './Paintings'
import { BrowserRouter, Route } from 'react-router-dom'
import SideBar from './SideBar';
import NavBar from './NavBar';
import NewExhibit from './NewExhibit'


function App() {


  return (
    <div className="container mx-auto">
      
      <BrowserRouter>
      <NavBar />
      
      <div className="flex flex-col md:flex-row">
        
        <SideBar />
        
        <Route exact path='/new' component={NewExhibit} />
        <Route exact path='/:section/:id' component={Paintings} />
        
      </div>

      


      </BrowserRouter>
    </div>
    
  );
}

export default App
