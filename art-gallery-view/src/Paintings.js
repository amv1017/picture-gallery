import React, { useState, useEffect } from 'react'
import Screen from './Screen'

const Paintings = (props) => {

	const [paintings, setPaintings] = useState([])

	useEffect(() => {
		const getPaintings = async () => {
			const res = await fetch(`http://localhost:8080/api/${props.match.params.section}/${props.match.params.id}`, { method: 'GET' })
			const data = await res.json()
			setPaintings(data)
		}
		getPaintings()
	}, [props.match.params.id])
 
	return(
		<div class="w-full max-w-6xl mx-auto px-4 border-2 border-grey-600">
			<Screen gallery={paintings} />
		</div>
	)
}

export default Paintings
