import React, { FunctionComponent, useContext } from 'react'
import SearchTagContext from '../contexts/searchTag';

interface Props {
  tagName: string
}

const Tag: FunctionComponent<Props> = (props: Props) => {
  const { searchTags, setSearchTags } = useContext(SearchTagContext);
  const handleAddTag = () => {
    const tags = searchTags
    if (!tags.includes(props.tagName)) {
      tags.push(props.tagName)
      setSearchTags(tags)
    }
  }
  return (
    <span className="tag" onClick={handleAddTag}>
      {props.tagName}
    </span>
  )
}

export default Tag