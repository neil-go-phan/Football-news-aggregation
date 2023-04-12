import React, { useState } from 'react'

function RelatedNews() {
  const [expanded, setExpanded] = useState(false);
  return (
    <div className='matchDetail__relatedNews p-3'>
      <div className="matchDetail__relatedNews--title">
        Tin liên quan
      </div>
      <div className="matchDetail__relatedNews--line my-3"></div>
      <div className="matchDetail__relatedNews--news">
      <p className="showmore" onClick={() => setExpanded(!expanded)}>
          {expanded ? 'Show less' : 'Show more...'}
        </p>
      </div>

    </div>
  )
}

export default RelatedNews