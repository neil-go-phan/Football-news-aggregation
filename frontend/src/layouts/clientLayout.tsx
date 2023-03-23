import React, { PropsWithChildren } from 'react';
import { Container, Navbar } from 'react-bootstrap';
import Head from 'next/head';
export default function AdminLayout({ children }: PropsWithChildren) {
  return (
    <>
      <Head>
        <title>Football news aggregation</title>
        <meta name="description" content="Neil intern demo 1" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className="wrapper d-flex flex-column min-vh-100 bg-light">
        <Navbar bg="secondary" expand="lg">
          <Container>
            <Navbar.Brand className="text-light" href="/">
              Football news
            </Navbar.Brand>
          </Container>
        </Navbar>
        <div className="body flex-grow-1 px-3">
          <Container fluid="lg">{children}</Container>
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
    </>
  );
}
