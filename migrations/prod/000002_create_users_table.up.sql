CREATE TABLE IF NOT EXISTS users (
	id bigserial PRIMARY KEY,
	added_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	login text UNIQUE NOT NULL,
	twitchid text NOT NULL,
	level integer,
	location text,
	lastfm_username text
);

INSERT INTO users (added_at,login,twitchid,"level",location,lastfm_username) VALUES
	 (NOW(),'nouryxd','31437432',1000,'vilnius','nouryqt'),
	 (NOW(),'nourybot','596581605',1000,'',''),
	 (NOW(),'uudelleenkytkeytynyt','465178364',1000,'',''),
	 (NOW(),'xnoury','197780373',500,'',''),
	 (NOW(),'noemience','135447564',500,'','');
