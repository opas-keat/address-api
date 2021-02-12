-- public.geographies definition

-- Drop table

-- DROP TABLE public.geographies;

CREATE TABLE public.geographies (
	id int4 NOT NULL,
	"name" varchar(255) NOT NULL,
	CONSTRAINT geographies_pkey PRIMARY KEY (id)
);

-- public.provinces definition

-- Drop table

-- DROP TABLE public.provinces;

CREATE TABLE public.provinces (
	id int4 NOT NULL,
	code varchar(2) NOT NULL,
	name_th varchar(150) NOT NULL,
	name_en varchar(150) NOT NULL,
	geography_id int4 NOT NULL DEFAULT 0,
	CONSTRAINT provinces_pkey PRIMARY KEY (id)
);

-- public.amphures definition

-- Drop table

-- DROP TABLE public.amphures;

CREATE TABLE public.amphures (
	id int4 NOT NULL,
	code varchar(4) NOT NULL,
	name_th varchar(150) NOT NULL,
	name_en varchar(150) NOT NULL,
	province_id int4 NOT NULL DEFAULT 0,
	CONSTRAINT amphures_pkey PRIMARY KEY (id)
);

-- public.districts definition

-- Drop table

-- DROP TABLE public.districts;

CREATE TABLE public.districts (
	id varchar(6) NOT NULL,
	zip_code int4 NOT NULL,
	name_th varchar(150) NOT NULL,
	name_en varchar(150) NOT NULL,
	amphure_id int4 NOT NULL DEFAULT 0,
	CONSTRAINT districts_pkey PRIMARY KEY (id)
);

