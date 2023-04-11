import React, { PropsWithChildren } from 'react';
import { Button, Container, Navbar } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBars } from '@fortawesome/free-solid-svg-icons';
import Head from 'next/head';
import { SidebarOverlay, Sidebar } from '@/components/sideBar';
function MatchDetailLayout({ children }: PropsWithChildren) {
  return (
    <>
      <Head>
        <title>Chi tiết trận đấu</title>
        <meta name="description" content="Neil intern demo 1" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className="wrapper d-flex flex-column min-vh-100 bg-light">
        <Navbar bg="secondary" expand="lg">
          <Container fluid className="header-navbar d-flex align-items-center">
            <Navbar.Brand className="text-light" href="/">
              Tin tức bóng đá
            </Navbar.Brand>
          </Container>
        </Navbar>
        <div className="body flex-grow-1 p-2">{children}</div>
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
    </>
  );
}

export default MatchDetailLayout;