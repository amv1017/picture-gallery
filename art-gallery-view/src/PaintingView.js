import PreView from './PreView'

const PaintingView = ({gallery}) => {
    return(
        <div>
        { gallery.map((i) => (
            <PreView paint={i} />
        )) }
        </div>
    )
}

export default PaintingView
