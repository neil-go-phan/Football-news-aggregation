import React, { PropsWithChildren, useState } from 'react';
import { Button, Container, Navbar } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBars } from '@fortawesome/free-solid-svg-icons';
import Head from 'next/head';
import { SidebarOverlay, Sidebar } from '@/components/sideBar';
import Link from 'next/link';
export default function NewsLayout({ children }: PropsWithChildren) {
  // Show status for xs screen
  const [isShowSmallScreenSidebar, setIsShowSmallScreenSidebar] =
    useState(false);
  const toggleIsShowSmallScreenSidebar = () => {
    setIsShowSmallScreenSidebar(!isShowSmallScreenSidebar);
  };
  return (
    <>
      <Head>
        <title>Tổng hợp tin tức bóng đá</title>
        <meta name="description" content="Neil intern demo 1" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className="wrapper d-flex flex-column min-vh-100 bg-light">
        <Navbar bg="secondary" expand="lg">
          <Container fluid className="header-navbar d-flex align-items-center">
          <Navbar.Brand className="text-light">
              <Link className="linkToHome" href={'/news/tin+tuc+bong+da?league=Tin+tức+bóng+đá'} >
                Football news
              </Link>
            </Navbar.Brand>
            <Button
              variant="link"
              className="sidebar-toggler d-md-none px-md-0 me-md-3 rounded-0 shadow-none"
              type="button"
              onClick={toggleIsShowSmallScreenSidebar}
            >
              <FontAwesomeIcon icon={faBars} />
            </Button>
          </Container>
        </Navbar>
        <div className="d-flex body flex-grow-1 p-2">
          <div className="col-xs-0 col-md-2">
            <Sidebar smallScreen={false} />
          </div>

          <div className="col-12 col-md-10">{children}</div>
        </div>
        {/* Footer */}
        <footer className="footer flex-column flex-md-row border-top d-flex align-items-center justify-content-between px-4 py-2 bg-secondary text-light">
          <div>
            <a
              className="text-decoration-none text-light"
              href="https://github.com/neil-go-phan/Football-news-aggregation"
            >
              Github{' '}
            </a>
          </div>
          <div className="ms-md-auto">
            Demo&nbsp;
            <span className="text-decoration-none">Neil Phan Golden Owl</span>
          </div>
        </footer>
      </div>

      <SidebarOverlay
        isShowSidebar={isShowSmallScreenSidebar}
        toggleSidebar={toggleIsShowSmallScreenSidebar}
      />
    </>
  );
}
