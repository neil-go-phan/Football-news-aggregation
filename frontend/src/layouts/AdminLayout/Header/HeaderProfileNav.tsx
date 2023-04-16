import { Dropdown, Nav, NavItem } from 'react-bootstrap';
import Image from 'next/image';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser } from '@fortawesome/free-regular-svg-icons';
import { PropsWithChildren, useContext, useState } from 'react';
import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { faPowerOff } from '@fortawesome/free-solid-svg-icons';
import Popup from 'reactjs-popup';
import { useRouter } from 'next/router';
import { _ROUTES } from '@/helpers/constants';
import AdminProfile from '@/components/admin/adminProfile';
import { IsLoggedContext } from '@/common/contexts/isLoggedContext';
import { deleteCookie } from 'cookies-next';

type ItemWithIconProps = {
  icon: IconDefinition;
} & PropsWithChildren;

const ItemWithIcon = (props: ItemWithIconProps) => {
  const { icon, children } = props;

  return (
    <>
      <FontAwesomeIcon className="me-2" icon={icon} fixedWidth />
      {children}
    </>
  );
};

export default function HeaderProfileNav() {
  const router = useRouter();
  const logged = useContext(IsLoggedContext);

  const handleLogout = () => {
    logged?.setIsLogged(false);
    deleteCookie("token")
    router.push(_ROUTES.ADMIN_LOGIN);

  };

  const [isProfileOpen, setIsProfileOpen] = useState(false);
  const handleIsProfileClose = () => {
    setIsProfileOpen(false);
  };
  return (
    <Nav>
      <Dropdown as={NavItem}>
        <Dropdown.Toggle
          variant="link"
          bsPrefix="hide-caret"
          className="py-0 px-2 rounded-0"
          id="dropdown-profile"
        >
          <div className="avatar position-relative">
            <Image
              fill
              className="rounded-circle"
              src="/images/avatar.png"
              alt="user@email.com"
            />
          </div>
        </Dropdown.Toggle>
        <Dropdown.Menu className="pt-0">
          <Dropdown.Header className="bg-light fw-bold">
            Settings
          </Dropdown.Header>
          <Dropdown.Item onClick={() => setIsProfileOpen(!isProfileOpen)}>
            <ItemWithIcon icon={faUser}>Profile</ItemWithIcon>
          </Dropdown.Item>
          <Popup modal open={isProfileOpen} onClose={handleIsProfileClose}>
            <AdminProfile handleIsProfileClose={handleIsProfileClose} />
          </Popup>

          <Dropdown.Divider />

          <Dropdown.Item onClick={handleLogout}>
            <ItemWithIcon icon={faPowerOff}>Logout</ItemWithIcon>
          </Dropdown.Item>
        </Dropdown.Menu>
      </Dropdown>
    </Nav>
  );
}
