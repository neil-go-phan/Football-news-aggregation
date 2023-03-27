import React from 'react'
import Article from './article'
import SearchBar from './searchBar'

export default function News() {
  return (
    <div className='news'>
      <SearchBar />
      <div className="news__articles">
        <Article />
        <Article />
      </div>
    </div>
  )
}
