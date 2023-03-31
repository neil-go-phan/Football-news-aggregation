import React, { useState } from 'react';
import MainContents from './mainContents';
import RightSideBar from './rightSideBar';
import SearchTagContext from '../../common/contexts/searchTag';

function Contents() {
  const [searchTags, setSearchTags] = useState<Array<string>>([]);
  return (
    <SearchTagContext.Provider value={{ searchTags, setSearchTags }}>
      <div className="contents d-md-flex py-2">
        <div className="col-12 col-md-9">
          <MainContents />
        </div>
        <div className="col-md-3">
          <RightSideBar />
        </div>
      </div>
    </SearchTagContext.Provider>
  );
}

export default Contents;
