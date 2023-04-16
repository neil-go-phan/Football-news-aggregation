import useGetLeaguesName, { LeaguesNames } from '@/helpers/cacheQuery/getLeaguesNames';
import { _ROUTES } from '@/helpers/constants';
import { formatRoute } from '@/helpers/format';
import Link from 'next/link';
import { useRouter } from 'next/router';
import React, { useState } from 'react';
import { Nav } from 'react-bootstrap';
import { ThreeDots } from 'react-loader-spinner';

export default function SidebarNav() {
  const router = useRouter();
  const [expanded, setExpanded] = useState(false);

  const { data, isLoading } = useGetLeaguesName();

  if (!isLoading) {
    if (data) {
      let route = router.asPath.substring(0, router.asPath.indexOf('?'));
      const leaguesForDisplay: LeaguesNames = expanded ? data : data!.slice(0, 10);
      return (
        <ul className="list-unstyled">
          {leaguesForDisplay!.map((league) => (
            <Nav.Item
              key={`navbar_item_${league}`}
              className={`px-3 py-2 d-flex align-items-center nav__item ${
                route === `${_ROUTES.NEWS_PAGE}/${formatRoute(league)}`
                  ? 'active'
                  : ''
              }`}
            >
              <Link
                href={{
                  pathname: `${_ROUTES.NEWS_PAGE}/${formatRoute(league)}`,
                  query: {league: league},
                }}
                className="text-decoration-none text-dark link"
              >
                {league}
              </Link>
            </Nav.Item>
          ))}
          <p className="sidebar--showmore" onClick={() => setExpanded(!expanded)}>
            {expanded ? 'Show less' : 'Show more...'}
          </p>
        </ul>
      );
    }
 
  }
  return (
    <ul className="list-unstyled">
      <div className="sidebar--loading">
        <ThreeDots
          height="50"
          width="50"
          radius="9"
          color="#4fa94d"
          ariaLabel="three-dots-loading"
          visible={true}
        />
      </div>
    </ul>
  );
}
