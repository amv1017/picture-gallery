const Polyvision = ({ paint, onClick }) => {
	return(
		<div className="max-w-sm rounded overflow-hidden shadow-lg bg-gray-200" onClick={onClick}>
			<div className="px-6 py-4">
				<img className="w-full" src={paint.url} />
				<p className="text-center">{paint.title}</p>
			</div>
		</div>
	)
}

export default Polyvision
