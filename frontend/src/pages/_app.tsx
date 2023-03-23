import ProgressBar from '@/components/ProcessBar';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'nprogress/nprogress.css';
// import '@/styles/globals.css';
import type { AppProps } from 'next/app';
import ClientLayout from '@/layouts/clientLayout';

export default function App({ Component, pageProps }: AppProps) {
  return (
    <ClientLayout>
      <ProgressBar />
      <Component {...pageProps} />
    </ClientLayout>
  );
}
