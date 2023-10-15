CREATE TABLE public.master_user (
	id bigserial NOT NULL,
  username varchar not null,
  "password" varchar not null,
  fullname varchar null,
  is_active bool not null default true,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL DEFAULT now(),
	modified_date timestamp NOT NULL DEFAULT now(),
	modified_by int8 NOT NULL,
	CONSTRAINT master_user_pk PRIMARY KEY (id)
);