CREATE TABLE articles (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  title text,
  description text,
  link text
);

CREATE TABLE admins (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  username text UNIQUE,
  password text
  email text UNIQUE
);

CREATE TABLE overview_items (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  time timestamp with time zone,
  match_id SERIAL,
  club_id SERIAL,
  info text,
  image_type text,
  time_string text
);

CREATE TABLE tags (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  tag_name text
);

CREATE TABLE players (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  match_line_up_id SERIAL,
  match_id SERIAL,
  row bigint,
  col bigint,
  player_name text,
  player_number text
);

CREATE TABLE schedules (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  date timestamp with time zone,
  league_name text
);

CREATE TABLE leagues (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  active boolean,
  league_name text
);

CREATE TABLE match_events (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  time timestamp with time zone,
  match_id SERIAL,
  time_string text,
  content text
);

CREATE TABLE statistics_items (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  match_id SERIAL,
  stat_content text,
  stat_club2 text,
  stat_club1 text
);

CREATE TABLE matches (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  time_start timestamp with time zone,
  club1_id SERIAL,
  club2_id SERIAL,
  lineup_club1_id SERIAL,
  lineup_club2_id SERIAL,
  schedule_id SERIAL,
  time text,
  round text,
  match_detail_link text,
  match_status text,
  scores text
);

CREATE TABLE clubs (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  name text,
  logo text
);

CREATE TABLE article_tag (
  article_id SERIAL,
  tag_id SERIAL,
  PRIMARY KEY (article_id, tag_id)
);

CREATE TABLE match_line_ups (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  club_name text,
  formation text,
  shirt_color text
);

CREATE TABLE config_crawlers (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  url text UNIQUE,
  article_div text,
  article_title text,
  article_description text,
  article_link text,
  next_page text,
  netx_page_type text
);

-- ADD FOREIGN KEY
ALTER TABLE
  article_tag
ADD
  CONSTRAINT fk_article_tag_tag FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE;

ALTER TABLE
  article_tag
ADD
  CONSTRAINT fk_article_tag_article FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE;

ALTER TABLE
  match_events
ADD
  CONSTRAINT fk_matches_events FOREIGN KEY (match_id) REFERENCES matches(id);

ALTER TABLE
  matches
ADD
  CONSTRAINT fk_matches_club2 FOREIGN KEY (club2_id) REFERENCES clubs(id);

ALTER TABLE
  matches
ADD
  CONSTRAINT fk_schedules_matches FOREIGN KEY (schedule_id) REFERENCES schedules(id);

ALTER TABLE
  matches
ADD
  CONSTRAINT fk_matches_club1 FOREIGN KEY (club1_id) REFERENCES clubs(id);

ALTER TABLE
  overview_items
ADD
  CONSTRAINT fk_matches_club2_overview FOREIGN KEY (match_id) REFERENCES matches(id);

ALTER TABLE
  overview_items
ADD
  CONSTRAINT fk_matches_club1_overview FOREIGN KEY (match_id) REFERENCES matches(id);

ALTER TABLE
  players
ADD
  CONSTRAINT fk_match_line_ups_players FOREIGN KEY (match_line_up_id) REFERENCES match_line_ups(id);

ALTER TABLE
  statistics_items
ADD
  CONSTRAINT fk_matches_statistics FOREIGN KEY (match_id) REFERENCES matches(id);

-- ADD INDEX
CREATE UNIQUE INDEX idx_admins_pkey ON public.admins USING btree (id);

CREATE INDEX idx_admins_deleted_at ON public.admins USING btree (deleted_at);

CREATE UNIQUE INDEX idx_admins_username ON public.admins USING btree (username);

CREATE UNIQUE INDEX idx_article_tag_pkey ON public.article_tag USING btree (article_id, tag_id);

CREATE UNIQUE INDEX idx_articles_pkey ON public.articles USING btree (id);

CREATE INDEX idx_articles_deleted_at ON public.articles USING btree (deleted_at);

CREATE UNIQUE INDEX idx_title_link ON public.articles USING btree (title, link);

CREATE UNIQUE INDEX idx_clubs_pkey ON public.clubs USING btree (id);

CREATE UNIQUE INDEX idx_club_name ON public.clubs USING btree (name);

CREATE INDEX idx_clubs_deleted_at ON public.clubs USING btree (deleted_at);

CREATE UNIQUE INDEX idx_clubs_name ON public.clubs USING btree (name);

CREATE UNIQUE INDEX idx_league_name ON public.leagues USING btree (league_name);

CREATE INDEX idx_leagues_deleted_at ON public.leagues USING btree (deleted_at);

CREATE UNIQUE INDEX idx_leagues_league_name ON public.leagues USING btree (league_name);

CREATE UNIQUE INDEX idx_leagues_pkey ON public.leagues USING btree (id);

CREATE INDEX idx_match_events_deleted_at ON public.match_events USING btree (deleted_at);

CREATE UNIQUE INDEX idx_match_events_pkey ON public.match_events USING btree (id);

CREATE INDEX idx_match_line_ups_deleted_at ON public.match_line_ups USING btree (deleted_at);

CREATE UNIQUE INDEX idx_match_line_ups_pkey ON public.match_line_ups USING btree (id);

CREATE INDEX idx_matches_deleted_at ON public.matches USING btree (deleted_at);

CREATE UNIQUE INDEX idx_matches_pkey ON public.matches USING btree (id);

CREATE INDEX idx_overview_items_deleted_at ON public.overview_items USING btree (deleted_at);

CREATE UNIQUE INDEX idx_overview_items_pkey ON public.overview_items USING btree (id);

CREATE INDEX idx_players_deleted_at ON public.players USING btree (deleted_at);

CREATE UNIQUE INDEX idx_players_pkey ON public.players USING btree (id);

CREATE UNIQUE INDEX idx_schedule_league_name ON public.schedules USING btree (date, league_name);

CREATE INDEX idx_schedules_deleted_at ON public.schedules USING btree (deleted_at);

CREATE UNIQUE INDEX idx_schedules_pkey ON public.schedules USING btree (id);

CREATE INDEX idx_statistics_items_deleted_at ON public.statistics_items USING btree (deleted_at);

CREATE UNIQUE INDEX idx_statistics_items_pkey ON public.statistics_items USING btree (id);

CREATE INDEX idx_tags_deleted_at ON public.tags USING btree (deleted_at);

CREATE UNIQUE INDEX idx_tags_tag_name ON public.tags USING btree (tag_name);

CREATE UNIQUE INDEX idx_tags_pkey ON public.tags USING btree (id);

INSERT INTO
  admins(id, created_at, username, password)
values
  (
    1,
    current_timestamp,
    'admin2023',
    'fa585d89c851dd338a70dcf535aa2a92fee7836dd6aff1226583e88e0996293f16bc009c652826e0fc5c706695a03cddce372f139eff4d13959da6f1f5d3eabe'
  );

INSERT INTO
  tags(id, created_at, tag_name)
values
  (1, current_timestamp, 'tin tuc bong da');

INSERT INTO
  leagues(id, created_at, league_name, active)
values
  (1, current_timestamp, 'Tin tức bóng đá', true);