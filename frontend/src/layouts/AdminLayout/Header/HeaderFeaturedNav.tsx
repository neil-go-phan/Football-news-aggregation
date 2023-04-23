import { _ROUTES } from '@/helpers/constants'
import Link from 'next/link'
import { Nav } from 'react-bootstrap'

export default function HeaderFeaturedNav() {
  return (
    <Nav>
      <Nav.Item>
        <Link href={_ROUTES.ADMIN_PAGE} passHref legacyBehavior>
          <Nav.Link className="p-2">Admin</Nav.Link>
        </Link>
      </Nav.Item>
      <Nav.Item>
        <Link href={'/news/tin+tuc+bong+da?league=Tin+tức+bóng+đá'} passHref legacyBehavior>
          <Nav.Link className="p-2" target='_blank'>News Page</Nav.Link>
        </Link>
      </Nav.Item>
    </Nav>
  )
}
