
export default function Blog(props){
    return(
        <article>
            <div>
                <h1>{props.title}</h1>
            </div>
            <div>
                <p>{props.content}</p>
            </div>
        </article>
    )
}