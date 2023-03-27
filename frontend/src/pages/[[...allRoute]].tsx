// import Contents from '@/components/contents';
// import ClientLayout from '@/layouts/clientLayout';
// import { useRouter } from 'next/router';
// import { useEffect } from 'react';
// export default function Home() {
//   const router = useRouter()
  
//   useEffect(() => {
//     router.push("/football-news")
//   })
  
//   return (
//     <></>
//   );
// }

import Contents from '@/components/contents';
import ClientLayout from '@/layouts/clientLayout';

export default function FootballNews() {
  return (
    <ClientLayout>
      <Contents />
    </ClientLayout>
  );
}