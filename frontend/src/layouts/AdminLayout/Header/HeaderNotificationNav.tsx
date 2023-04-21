import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faBell } from '@fortawesome/free-regular-svg-icons'
import {
  Badge, Dropdown, Nav, NavLink,
} from 'react-bootstrap'
import Link from 'next/link'
import React, { useEffect, useState } from 'react'
import Image from 'next/image'
import { fetchEventSource } from '@microsoft/fetch-event-source'

export default function HeaderNotificationNav() {
  // const [data, setData] = useState<nerver>([]);
  var sse = new EventSource(`${process.env.NEXT_PUBLIC_BACKEND_DOMAIN}notification/get`)
  sse.onmessage = console.log 
  // useEffect(() => {
  //   const fetchData = async () => {
  //     await fetchEventSource(`${process.env.NEXT_PUBLIC_BACKEND_DOMAIN}notification/get`, {
  //       onmessage(ev) {
  //           console.log(ev.data);
  //       }
  //   });
  //   };
  //   fetchData();
  // }, []);
  // console.log(data)
  return (
    <Nav>
      <Nav.Item>
        <Dropdown>
          <Dropdown.Toggle as={NavLink} bsPrefix="hide-caret" id="dropdown-mail">
            <FontAwesomeIcon icon={faBell} size="lg" />
            <Badge pill bg="primary" className="position-absolute top-0 right-0">
              7
            </Badge>
          </Dropdown.Toggle>
          <Dropdown.Menu className="pt-0" align="end">
            <Dropdown.Header className="bg-light fw-bold rounded-top">You have 4 messages</Dropdown.Header>
            <Link href="/" passHref legacyBehavior>
              <Dropdown.Item>
                <div className="message">
                  <div className="py-3 me-3 float-start">
                    <div className="avatar d-inline-flex position-relative">
                      <Image
                        fill
                        className="rounded-circle"
                        src="/assets/img/avatars/1.jpg"
                        alt="user@email.com"
                      />
                      <span
                        className="avatar-status position-absolute d-block bottom-0 end-0 bg-success rounded-circle border border-white"
                      />
                    </div>
                  </div>
                  <div>
                    <small className="text-muted">John Doe</small>
                    <small className="text-muted float-end mt-1">Just now</small>
                  </div>
                  <div className="text-truncate font-weight-bold">
                    <span className="text-danger">!</span>
                    {' '}
                    Pet Pikachu
                  </div>
                  <div className="small text-truncate text-muted">
                    Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                    sed do eiusmod tempor incididunt
                  </div>
                </div>
              </Dropdown.Item>
            </Link>
            <Link href="/" passHref legacyBehavior>
              <Dropdown.Item>
                <div className="message">
                  <div className="py-3 me-3 float-start">
                    <div className="avatar d-inline-flex position-relative">
                      <Image
                        fill
                        className="rounded-circle"
                        src="/assets/img/avatars/2.jpg"
                        alt="user@email.com"
                      />
                      <span
                        className="avatar-status position-absolute d-block bottom-0 end-0 bg-warning rounded-circle border border-white"
                      />
                    </div>
                  </div>
                  <div>
                    <small className="text-muted">John Doe</small>
                    <small className="text-muted float-end mt-1">5 mins ago</small>
                  </div>
                  <div className="text-truncate font-weight-bold">
                    Dress Eevee
                  </div>
                  <div className="small text-truncate text-muted">
                    Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                    sed do eiusmod tempor incididunt
                  </div>
                </div>
              </Dropdown.Item>
            </Link>
            <Link href="/" passHref legacyBehavior>
              <Dropdown.Item>
                <div className="message">
                  <div className="py-3 me-3 float-start">
                    <div className="avatar d-inline-flex position-relative">
                      <Image
                        fill
                        className="rounded-circle"
                        src="/assets/img/avatars/3.jpg"
                        alt="user@email.com"
                      />
                      <span
                        className="avatar-status position-absolute d-block bottom-0 end-0 bg-danger rounded-circle border border-white"
                      />
                    </div>
                  </div>
                  <div>
                    <small className="text-muted">John Doe</small>
                    <small className="text-muted float-end mt-1">1:52 PM</small>
                  </div>
                  <div className="text-truncate font-weight-bold">
                    Team up training
                  </div>
                  <div className="small text-truncate text-muted">
                    Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                    sed do eiusmod tempor incididunt
                  </div>
                </div>
              </Dropdown.Item>
            </Link>
            <Link href="/" passHref legacyBehavior>
              <Dropdown.Item>
                <div className="message">
                  <div className="py-3 me-3 float-start">
                    <div className="avatar d-inline-flex position-relative">
                      <Image
                        fill
                        className="rounded-circle"
                        src="/assets/img/avatars/4.jpg"
                        alt="user@email.com"
                      />
                      <span
                        className="avatar-status position-absolute d-block bottom-0 end-0 bg-primary rounded-circle border border-white"
                      />
                    </div>
                  </div>
                  <div>
                    <small className="text-muted">John Doe</small>
                    <small className="text-muted float-end mt-1">4:03 PM</small>
                  </div>
                  <div className="text-truncate font-weight-bold">
                    Go to Safari Zone
                  </div>
                  <div className="small text-truncate text-muted">
                    Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                    sed do eiusmod tempor incididunt
                  </div>
                </div>
              </Dropdown.Item>
            </Link>
          </Dropdown.Menu>
        </Dropdown>
      </Nav.Item>
    </Nav>
  )
}
