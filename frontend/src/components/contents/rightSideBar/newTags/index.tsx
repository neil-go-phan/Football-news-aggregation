import React, { FunctionComponent, useState } from 'react';
import Tag from '../../../../common/tag';
import useGetTags, { Tags } from '@/helpers/cacheQuery/getTags';
import { ThreeDots } from 'react-loader-spinner';

const NewTags: FunctionComponent = () => {
  const [expanded, setExpanded] = useState(false);

  const { data, isLoading } = useGetTags();

  if (!isLoading) {
    const tagForDisplay: Tags = expanded ? data : data!.slice(0, 10);

    return (
      <div className="rightSideBar__tags">
        <p className="rightSideBar__tags--title">Tags</p>
        <div className="rightSideBar__tags--line mb-3"></div>
        <div className="rightSideBar__tags--list">
          {tagForDisplay!.map((tag) => {
            if (tag !== 'tin tuc bong da') {
              return <Tag key={`rightSideBar__tags_${tag}`} tagName={tag} />;
            }
          })}
          {data!.length > 10 ? (
            <p
              className="rightSideBar__tags--showmore"
              onClick={() => setExpanded(!expanded)}
            >
              {expanded ? 'Show less' : 'Show more...'}
            </p>
          ) : (
            <></>
          )}
        </div>
      </div>
    );
  }
  return (
    <div className="rightSideBar__tags">
      <p className="rightSideBar__tags--title">Tags</p>
      <div className="rightSideBar__tags--line mb-3"></div>
      <div className="rightSideBar__tags--list">
        <div className="rightSideBar__tags--loading">
          <ThreeDots
            height="50"
            width="50"
            radius="9"
            color="#4fa94d"
            ariaLabel="three-dots-loading"
            visible={true}
          />
        </div>
      </div>
    </div>
  );
};

export default NewTags;
