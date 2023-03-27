import React, { useEffect } from 'react';
import classNames from 'classnames';
import SidebarNav from './sidebarNav';
import { useRouter } from 'next/router';

export function Sidebar(props: { smallScreen: boolean }) {
  const { smallScreen } = props;
  return (
    <div
      className={`sidebar d-flex flex-column h-100 bg-light ${
        !smallScreen ? 'lagreScreen' : ''
      }`}
      id="sidebar"
    >
      <div className="sidebar-nav flex-fill">
        <SidebarNav />
      </div>
    </div>
  );
}

export const SidebarOverlay = (props: {
  isShowSidebar: boolean;
  toggleSidebar: () => void;
}) => {
  const { isShowSidebar, toggleSidebar } = props;
  const router = useRouter()
  useEffect(() => {
    if (isShowSidebar) {
      toggleSidebar()
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [router.asPath])
  
  return (
    <div
      tabIndex={-1}
      aria-hidden
      className={classNames(
        'sidebar-overlay position-fixed top-0 w-100 h-100 d-flex',
        {
          'd-none': !isShowSidebar,
        }
      )}
    >
      <div className="">
        <Sidebar smallScreen={true} />
      </div>
      <div className="col-10 bg-dark opacity-50" onClick={toggleSidebar}></div>
    </div>
  );
};
