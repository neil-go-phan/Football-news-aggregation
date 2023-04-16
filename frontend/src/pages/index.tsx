import { useRouter } from 'next/router';
import { useEffect } from 'react';
export default function Home() {
  const router = useRouter()
  
  useEffect(() => {
    router.push('/news/tin+tuc+bong+da?league=Tin+tức+bóng+đá')
  })
  
  return (
    <></>
  );
}