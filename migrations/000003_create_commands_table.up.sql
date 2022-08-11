CREATE TABLE IF NOT EXISTS commands (
   id bigserial PRIMARY KEY,
   name text UNIQUE NOT NULL,
   text text NOT NULL,
   level integer NOT NULL
);

INSERT INTO commands (name,"text","level") VALUES
	 ('repeat','xset r rate 175 50',0),
	 ('eurkey','setxkbmap -layout eu',0),
	 ('clueless','ch02 ch21 ch31',0),
	 ('feelsdankman','⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⢀⣾⣿⣿⣿⣿⣷⣄⠹⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⣿⣿⣿⣿⣿⠟⢁⣴⣿⣿⣿⣿⣿⣿⣿⣿⣦⡈⢻⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⣿⣿⣿⡿⠁⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠙⢿⣿⣿⣿⣿⣿ ⣿⣿⣿⣿⣿⡿⠃⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⡈⢻⣿⣿⣿⣿ ⣿⣿⣿⣿⡿⢁⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠹⣿⣿⣿ ⣿⣿⣿⣿⠁⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠿⠿⠿⠿⠿⠆⠘⢿⣿ ⣿⣿⠟⠉⠄⠄⠄⠄⢤⣀⣦⣤⣤⣤⣤⣀⣀⡀⠄⠄⡀⠄⠄⠄⠄⠄⠄⠄⠙ ⣿⠃⠄⠄⠄⠄⠄⠄⠙⠿⣿⣿⠋⠩⠉⠉⢹⣿⣧⣤⣴⣶⣷⣿⠟⠛⠛⣿⣷ ⠇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠁⠒⠄⠄⠄⠄⠈⠉⠛⢻⣿⣿⢿⠁⠄⠄⠁⠘⢁ ⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣂⣀⣐⣂⣐⣒⣃⠠⠥⠤⠴⠶⠖⠦⠤⠖⢂⣽ ⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠛⠂⠐⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣠⣴⣶⣿⣿ ⠃⣠⣄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣠⣤⣄⠚⢿⣿⣿⣿⣿ ⣾⣿⣿⣿⣶⣦⣤⣤⣄⣀⣀⣀⣀⣀⣀⣠⣤⣤⣶⣿⣿⣿⣿⣷⡄⢻⣿⣿⣿ ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠈⣿⣿⣿',1000),
	 ('dankhug','⣼⣿⣿⣿⣷⡄⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣾⣿⣿⣿⣿⣿⣄⠀⠀⠀⠀⠀⠀ ⣿⣿⣿⣿⣿⣿⣆⠀⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀ ⣿⣿⣿⣿⣿⣿⣿⣧⡀⠀⠀⠀⣠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡀⠀⠀⠀ ⣿⣿⣿⣿⣿⣿⣿⣿⣷⡄⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⠀⠀ ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣄⠺⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠟⠛⠛⠀⠀ ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠆⠒⠀⠀⠶⢶⣶⣶⣭⣤⠹⠟⣛⢉⣉⣉⣀⣀ ⣿⣿⣾⣿⣶⣶⣶⣶⣶⣶⣿⣿⣶⠀⢬⣒⣂⡀⠀⠀⠀⠀⣈⣉⣉⣉⣉⡉⠅ ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⢭⣭⠭⠭⠉⠠⣷⣆⣂⣐⣐⣒⣒⡈ ⢿⣿⣿⣿⠋⢁⣄⡈⠉⠛⠛⠻⡿⠟⢠⡻⣿⣿⣛⣛⡋⠉⣀⠤⣚⠙⠛⠉⠁ ⠀⠙⠛⠛⠀⠘⠛⠛⠛⠛⠋⠀⠨⠀⠀⠀⠒⠒⠒⠒⠒⠒⠒⠊⡀⠀⠀⠀⠀ ⠀⠀⠀⠀⠀⠀⠀⠰⣾⣿⣿⣷⣦⠀⠀⠀⠀⠀⠀⠀⢀⣠⠴⢊⣭⣦⡀⠀⠀ ⠀⠀⠀⠀⠀⠀⠀⣠⣌⠻⣿⣿⣷⣞⣠⣖⣠⣶⣴⣶⣶⣶⣾⣿⣿⣿⣿⡀⠀ ⠀⢀⣀⣀⣠⣴⣾⣿⣿⣷⣌⠛⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⢻⣿⣷⡁ ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢻⣶⣬⣭⣉⣛⠛⠛⢛⣛⣉⣭⣴⣾⣿⣿⣿⡇',1000);
