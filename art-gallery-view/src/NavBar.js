import React from 'react'
import { Link } from 'react-router-dom'

function NavBar() {
	return (
		<div className="flex items-center h-8 bg-gradient-to-r from-purple-400 via-cyan-500 to-blue-500">
			<div className="ml-16 text-black">ART GALLERY</div>
			<ul className="flex flex-auto justify-end pr-40">
				<li className="text-gray-100 m-9"><Link to="/">Home</Link></li>
				<li className="text-gray-100 m-9"><Link to="/new">New</Link></li>
			</ul>
		</div>
	)
}

export default NavBar
