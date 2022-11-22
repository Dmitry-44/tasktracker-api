CREATE TABLE IF NOT EXISTS public.users (
	id bigserial NOT NULL,
	"name" varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.tasks (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	title varchar(255) NOT NULL,
	status int4 NOT NULL DEFAULT 0,
	created_by int4 NOT NULL DEFAULT 0,
	priority int4 NOT NULL DEFAULT 0,
	description varchar(1000) NOT NULL,
	group_id int4 NOT NULL,
	CONSTRAINT tasks_pk PRIMARY KEY (id),
	CONSTRAINT u_id FOREIGN KEY (created_by) REFERENCES public.users(id),
	CONSTRAINT g_id FOREIGN KEY (group_id) REFERENCES public.groups(id)
);

CREATE TABLE IF NOT EXISTS public."group" (
	id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	title text NOT NULL,
	description text NOT NULL,
	created_by int4 NOT NULL,
	CONSTRAINT group_pk PRIMARY KEY (id),
	CONSTRAINT u_id FOREIGN KEY (created_by) REFERENCES public.users(id)
);