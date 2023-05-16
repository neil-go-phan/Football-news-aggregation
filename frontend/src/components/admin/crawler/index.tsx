import React from 'react'
import EmbedWeb from './embedWeb'

function Crawler() {

  const handleClick = (event:Event):string => {
    const target = event.target as HTMLElement;
    return target.className
  };
  
  return (
    <div className='adminCrawler'>
      <EmbedWeb url='https://vnexpress.net/the-thao' handleClick={handleClick}/>
    </div>
  )
}

export default Crawler