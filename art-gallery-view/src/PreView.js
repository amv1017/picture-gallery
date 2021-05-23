const PreView = ({ paint }) => {
    return(
    <div className="grid grid-cols-2 gap-1">
       
            <img src={paint.url} />
            <p className="bg-gray-300 place-self-center">{paint.description}</p>
       
    </div>
    )
}

export default PreView
