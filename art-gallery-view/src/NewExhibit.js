import React, { useState } from 'react'

const NewExhibit = () => {

    const [title,setTitle] = useState('');
    const [description,setDescription] = useState('');
    const [url,setUrl] = useState('');
    const [genre,setGenre] = useState('');
    const [author,setAuthor] = useState('');

    const handleFormSubmit = (e) => {
        e.preventDefault();

        send()

        setTitle('')
        setDescription('')
        setUrl('')
        setGenre('')
        setAuthor('')

    };

    const send = async () => {
        const data = JSON.stringify({
            'title': title,
            'description': description,
            'url': url,
            'genre': genre,
            'author': author
        })
        console.log(data)
        await fetch('http://localhost:8080/paintings', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: data
        })
    }

    return (
      
            <div className="flex flex-col">

                
                <form className="ml-40 mt-16 flex flex-col items-right" onSubmit={handleFormSubmit}>

                    <div>
                        <label htmlFor='title'>Title</label>
                        <input
                            type='text'
                            className={`w-full p-2 text-primary border rounded-md outline-none text-sm transition duration-150 ease-in-out mb-4`}
                            value={title}
                            placeholder='Enter Title'
                            onChange={(e) => setTitle(e.target.value)}
                        />
                    </div>

                    <div>
                        <label htmlFor='author'>Author</label>
                        <input
                            type='text'
                            className={`w-full p-2 text-primary border rounded-md outline-none text-sm transition duration-150 ease-in-out mb-4`}
                            value={author}
                            placeholder='Enter Author'
                            onChange={(e) => setAuthor(e.target.value)}
                        />
                    </div>

                    <div>
                        <label htmlFor='genre'>Genre</label>
                        <input
                            type='text'
                            className={`w-full p-2 text-primary border rounded-md outline-none text-sm transition duration-150 ease-in-out mb-4`}
                            value={genre}
                            placeholder='Enter Genre'
                            onChange={(e) => setGenre(e.target.value)}
                        />
                    </div>


                    <div>
                        <label htmlFor='description'>Description</label>
                        <input
                            type='text'
                            className={`w-full p-2 text-primary border rounded-md outline-none text-sm transition duration-150 ease-in-out mb-4`}
                            value={description}
                            placeholder='Enter Description'
                            onChange={(e) => setDescription(e.target.value)}
                        />
                    </div>

                    <div>
                        <label htmlFor='url'>URL</label>
                        <input
                            type='text'
                            className={`w-full p-2 text-primary border rounded-md outline-none text-sm transition duration-150 ease-in-out mb-4`}
                            value={url}
                            placeholder='Enter URL'
                            onChange={(e) => setUrl(e.target.value)}
                        />
                    </div>

                    <div className='flex justify-center items-center mt-6'>
                        <button
                            className={`bg-green-300 py-2 px-4 text-sm text-white rounded border border-green focus:outline-none focus:border-green-dark`}
                        >
                            Confirm
                        </button>

                    </div>
                </form>
            <div className="ml-40 mt-16 flex flex-col items-right">
                <ul>
                    Example
                    <li>Title: Black Square</li>
                    <li>Author: Kazimir Malevich</li>
                    <li>Genre: Still Life</li>
                    <li>Description: Masterpiece...</li>
                    <li>URL: https://img.gazeta.ru/files3/907/7888907/malevich21_-pic905-895x505-1560.jpg</li>
                </ul>

            </div>
            </div>
    )
}

export default NewExhibit
