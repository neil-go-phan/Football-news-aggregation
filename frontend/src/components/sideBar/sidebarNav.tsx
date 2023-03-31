import { _ROUTES } from '@/helpers/constants';
import Link from 'next/link';
import { useRouter } from 'next/router';
import React from 'react';
import { Nav } from 'react-bootstrap';

export default function SidebarNav() {
  const router = useRouter();
  return (
    <ul className="list-unstyled">
      <Nav.Item
        className={`px-3 py-2 d-flex align-items-center nav__item ${
          router.asPath === _ROUTES.NEWS_PAGE ? 'active' : ''
        }`}
      >
        <Link href={_ROUTES.NEWS_PAGE} className="text-decoration-none text-dark link">
          Tin tức chung
        </Link>
      </Nav.Item>
      <Nav.Item
        className={`px-3 py-2 d-flex align-items-center nav__item ${
          router.asPath === `${_ROUTES.NEWS_PAGE}/ngoai-hang-anh` ? 'active' : ''
        }`}
      >
        <Link
          href={`${_ROUTES.NEWS_PAGE}/ngoai-hang-anh`}
          className="text-decoration-none text-dark link"
        >
          Ngoại hạng anh
        </Link>
      </Nav.Item>
      <Nav.Item
        className={`px-3 py-2 d-flex align-items-center nav__item ${
          router.asPath === `${_ROUTES.NEWS_PAGE}/cup-c1` ? 'active' : ''
        }`}
      >
        <Link href={`${_ROUTES.NEWS_PAGE}/cup-c1`} className="text-decoration-none text-dark link">
          Cup C1
        </Link>
      </Nav.Item>
    </ul>
  );
}
