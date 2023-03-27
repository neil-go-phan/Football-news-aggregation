import '@fortawesome/fontawesome-svg-core/styles.css';
import '@/styles/globals.scss';
import ProgressBar from '@/components/processBar';
import type { AppProps } from 'next/app';

import { SSRProvider } from 'react-bootstrap';

export default function App({ Component, pageProps }: AppProps) {
  return (
    <SSRProvider>
      <ProgressBar />
      <Component {...pageProps} />
    </SSRProvider>
  );
}
