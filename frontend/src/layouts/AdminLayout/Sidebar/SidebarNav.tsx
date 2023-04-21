import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faNewspaper,
  IconDefinition,
} from '@fortawesome/free-regular-svg-icons';
import { faGauge, faTag, faTrophy } from '@fortawesome/free-solid-svg-icons';
import React, { PropsWithChildren } from 'react';
import { Nav } from 'react-bootstrap';
import Link from 'next/link';
import { _ROUTES } from '@/helpers/constants';

type SidebarNavItemProps = {
  href: string;
  icon?: IconDefinition;
} & PropsWithChildren;

const SidebarNavItem = (props: SidebarNavItemProps) => {
  const { icon, children, href } = props;

  return (
    <Nav.Item>
      <Link href={href} passHref legacyBehavior>
        <Nav.Link className="px-3 py-2 d-flex align-items-center">
          {icon ? (
            <FontAwesomeIcon className="nav-icon ms-n3" icon={icon} />
          ) : (
            <span className="nav-icon ms-n3" />
          )}
          {children}
        </Nav.Link>
      </Link>
    </Nav.Item>
  );
};

export default function SidebarNav() {
  return (
    <ul className="list-unstyled">
      <SidebarNavItem icon={faGauge} href={_ROUTES.ADMIN_PAGE}>
        Dashboard
      </SidebarNavItem>
      <SidebarNavItem icon={faNewspaper} href={_ROUTES.ADMIN_ARTICLES}>
        Articles
      </SidebarNavItem>
      <SidebarNavItem icon={faTrophy} href={_ROUTES.ADMIN_LEAGUES}>
        Leagues
      </SidebarNavItem>
      <SidebarNavItem icon={faTag} href={_ROUTES.ADMIN_TAGS}>
        Tags
      </SidebarNavItem>
    </ul>
  );
}
