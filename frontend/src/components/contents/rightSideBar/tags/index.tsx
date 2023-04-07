import axiosClient from '@/helpers/axiosClient';
import React, { FunctionComponent, useEffect, useState } from 'react';
import { toast } from 'react-toastify';
import Tag from '../../../../common/tag';

const Tags: FunctionComponent = () => {
  const [tags, setTags] = useState<Array<string>>([]);
  const [expanded, setExpanded] = useState(false);
  useEffect(() => {
    const getTags = async () => {
      try {
        const { data } = await axiosClient.get('tags/list');
        setTags(data.tags.tags);
      } catch (error) {
        toast.error('Error occurred while geting tags', {
          position: 'top-right',
          autoClose: 3000,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: 'light',
        });
      }
    };
    getTags();
  }, []);

  const tagForDisplay = expanded ? tags : tags.slice(0, 10);
  return (
    <div className="rightSideBar__tags">
      <p className="rightSideBar__tags--title">Tags</p>
      <div className="rightSideBar__tags--line mb-3"></div>
      <div className="rightSideBar__tags--list">
        {tagForDisplay.map((tag) => (
          <Tag key={`rightSideBar__tags_${tag}`} tagName={tag} />
        ))}
        <p className="rightSideBar__tags--showmore" onClick={() => setExpanded(!expanded)}>
          {expanded ? 'Show less' : 'Show more...'}
        </p>
      </div>
    </div>
  );
};

export default Tags;
