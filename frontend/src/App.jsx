import Blog from "./components/Blog"
import data from "./data.js"
export default function App() {
  // This for now reads static data
  const blogs = data.map((blog) => {
    return <Blog 
      key={blog.slug}
      {...blog}
    />
  })
  return(
  <>
    <h1>Welcome to my Lovely blog</h1>
    {blogs}
  </>

  )
}