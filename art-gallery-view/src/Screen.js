import React, { useState } from 'react'
import Polyvision from './Polyvision'
import Exhibit from './Exhibit'

const Screen = ({ gallery }) => {

	const [showOne, setShowOne] = useState(false)
	const [ind, setInd] = useState(1)

	return(
		<>

			{ gallery === null ? ('No Exhibits') : (

				showOne ?

				(<div className="grid grid-cols-1 gap-4">
					<div onClick={() => setShowOne(!showOne)}>
						<Exhibit id={gallery[ind].id} />
					</div>
				</div>) :

				(<div className="grid grid-cols-3 gap-4">
					{ gallery.map((i) => (
						<div onClick={() => { setShowOne(!showOne) }}>
							<Polyvision paint={i} onClick={() => { setInd(gallery.indexOf(i)) }} />
						</div>
					))}
				</div>)

			)}

		</>
	)
}

export default Screen
