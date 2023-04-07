import axiosClient from '@/helpers/axiosClient';
import { _ROUTES } from '@/helpers/constants';
import Link from 'next/link';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { Nav } from 'react-bootstrap';
import { toast } from 'react-toastify';

export default function SidebarNav() {
  const router = useRouter();
  const [leagues, setLeagues] = useState<Array<any>>([]);
  const [expanded, setExpanded] = useState(false);

  const formatRoute = (league: string) => {
    const route = league
      .normalize('NFD')
      .replace(/[\u0300-\u036f]/g, '')
      .replace(/đ/g, 'd')
      .replace(/Đ/g, 'D');

    return route.replaceAll(' ', '+');
  };

  useEffect(() => {
    const getLeagues = async () => {
      try {
        const { data } = await axiosClient.get('leagues/list');
        setLeagues(data.leagues.leagues);
      } catch (error) {
        toast.error('Error occurred while geting leagues', {
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
    getLeagues();
  }, []);
  const leaguesForDisplay: Array<string> = expanded
    ? leagues
    : leagues.slice(0, 10);
  return (
    <ul className="list-unstyled">
      {leaguesForDisplay.map((league) => (
        <Nav.Item
          key={`navbar_item_${league}`}
          className={`px-3 py-2 d-flex align-items-center nav__item ${
            router.asPath === `${_ROUTES.NEWS_PAGE}/${formatRoute(league)}`
              ? 'active'
              : ''
          }`}
        >
          <Link
            href={`${_ROUTES.NEWS_PAGE}/${formatRoute(league)}`}
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
