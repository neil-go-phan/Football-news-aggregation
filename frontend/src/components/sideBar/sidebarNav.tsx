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
          router.asPath === '/' ? 'active' : ''
        }`}
      >
        <Link href={'/'} className="text-decoration-none text-dark link">
          Tin tức chung
        </Link>
      </Nav.Item>
      <Nav.Item
        className={`px-3 py-2 d-flex align-items-center nav__item ${
          router.asPath === '/ngoai-hang-anh' ? 'active' : ''
        }`}
      >
        <Link
          href={'/ngoai-hang-anh'}
          className="text-decoration-none text-dark link"
        >
          Ngoại hạng anh
        </Link>
      </Nav.Item>
      <Nav.Item
        className={`px-3 py-2 d-flex align-items-center nav__item ${
          router.asPath === '/cup-c1' ? 'active' : ''
        }`}
      >
        <Link href={'/cup-c1'} className="text-decoration-none text-dark link">
          Cup C1
        </Link>
      </Nav.Item>
    </ul>
  );
}
