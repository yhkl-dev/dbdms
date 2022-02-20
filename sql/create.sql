CREATE TABLE public.t_databases (
	id serial4 NOT NULL,
	name varchar NULL,
	host varchar NULL,
	port int4 NULL,
	username varchar NULL,
	password varchar NULL,
	comment varchar NULL,
    schema varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT movies_pkey PRIMARY KEY (id)
);

CREATE TABLE public.t_genre (
	id serial4 NOT NULL,
	genre_name int4 NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT genre_pkey PRIMARY KEY (id)
);

CREATE TABLE public.t_database_genre (
	id serial4 NOT NULL,
	database_id int NULL,
	genre_id int NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT database_genre_pkey PRIMARY KEY (id)
);

CREATE TABLE public.t_database_document (
	id serial4 NOT NULL,
	version_id int4 NULL,
	database_id int NULL,
	document text NULL,
	created_at timestamp NULL,
	CONSTRAINT document_pkey PRIMARY KEY (id)
);

CREATE TABLE public.t_user (
	id serial4 NOT NULL,
	username varchar NULL,
	email varchar  NULL,
	password varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT user_pkey PRIMARY KEY (id)
);