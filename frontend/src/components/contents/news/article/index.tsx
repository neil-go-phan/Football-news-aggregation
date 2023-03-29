import React, { FunctionComponent } from 'react'

interface Props {
  article: ArticleType
}

export type ArticleType = {
  title: string
  discription: string
  thumbnail: string
  link: string
}


const Article: FunctionComponent<Props> = (props:Props) =>{
  const {article} = props
  return (
    <div className="articles__article">
        {article.title}
    </div>
  )
}

export default Article