import { useState, useEffect } from "react"
import Card from "../components/Card"
import postServices from "../services/services"

const ListPage = () => {
  const [posts, setPosts] = useState([])
  const [newPost, setNewPost] = useState('')
  const [newPostTitle, setNewPostTitle] = useState('')

  useEffect(() => {
    postServices.getAll().then((initialPosts) => {
      setPosts(initialPosts)
    })
  }, [])

  return (
    
    <div className="bg-gray-200">
      <Card />
    </div>
    
    
  )
}

export default ListPage