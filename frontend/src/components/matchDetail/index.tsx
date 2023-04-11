import React from 'react'
import RelatedNews from './relatedNews'
import Detail from './detail'

function MatchDetail() {
  return (
    <>
      <div className="matchDetail d-md-flex py-2">
        <div className="col-12 col-md-9">
          <Detail />
        </div>
        <div className="col-md-3">
          <RelatedNews />
        </div>
      </div>
    </>
  )
}

export default MatchDetail