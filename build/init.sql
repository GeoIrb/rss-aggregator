CREATE TABLE public.news (
	id serial NOT NULL,
	title varchar NOT NULL,
	pubDate timestamp NOT NULL,
	CONSTRAINT news_pkey PRIMARY KEY 
);