CREATE TABLE IF NOT EXISTS commands (
   id bigserial PRIMARY KEY,
   name text UNIQUE NOT NULL,
   text text NOT NULL,
   permission integer NOT NULL
);

INSERT INTO commands (name,"text","permission") VALUES
	 ('repeat','xset r rate 175 50',0),
	 ('eurkey','setxkbmap -layout eu',0);