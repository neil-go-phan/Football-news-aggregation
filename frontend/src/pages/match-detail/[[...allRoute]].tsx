import MatchDetail from '@/components/matchDetail';
import MatchDetailLayout from '@/layouts/matchDetailLayout';

export default function FootballNews() {
  // check routes
  return (
    <MatchDetailLayout>
      <MatchDetail />
    </MatchDetailLayout>
  );
}