import React, { FunctionComponent } from 'react'
import NewTags from './newTags'


const RightSideBar: FunctionComponent= () => {
  return (
    <div className='rightSideBar px-2'>
      <NewTags />
    </div>
  )
}

export default RightSideBar