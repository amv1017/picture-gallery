import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'

const SideBar = () => {

	const [authors, setAuthors]  = useState([])

	useEffect(() => {
		const getAuthors = async () => {
		  setAuthors(await fetchAuthors())
		}
		getAuthors()
	}, [])

	const fetchAuthors = async () => {
		const res = await fetch('http://localhost:8080/api/authors', { method: 'GET' })
		const data = await res.json()
		return data
	}

	return (
		<div>
			<div className="mb-3 text-center bg-blue-200 rounded-md p-3">AUTHORS
				<div className="flex flex-col items-center pb-6">
					{ authors.map((i) => (
						<div>
							<Link className="bg-blue-400 rounded-md px-5 mb-4" to={'/author/'+i.id}>{i.name}</Link>
						</div>
					)) }
				</div>
			</div>
			<div className="mb-3 text-center bg-green-200 rounded-md p-3">GENRES
				<div className="flex flex-col items-center">
					<Link className="bg-green-400 rounded-md py-2 px-5 mb-1" to={'/genre/1'}>Landscape</Link>
					<Link className="bg-green-400 rounded-md py-2 px-5 mb-1" to={'/genre/2'}>Portrait</Link>
					<Link className="bg-green-400 rounded-md py-2 px-5" to={'/genre/3'}>Still Life</Link>
				</div>
			</div>
		</div>
	)
}

export default SideBar
