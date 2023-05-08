-- public.assets definition

-- Drop table

-- DROP TABLE public.assets;

CREATE TABLE public.assets (
	address varchar NULL,
	tokenaddress varchar NULL,
	chainname varchar NULL,
	amount varchar NULL,
	createddate timestamp NULL DEFAULT now(),
	updateddate timestamp NULL DEFAULT now(),
	top varchar NULL
);


-- public.crypto definition

-- Drop table

-- DROP TABLE public.crypto;

CREATE TABLE public.crypto (
	id int4 NULL,
	"key" varchar NULL,
	"name" varchar NULL,
	symbol varchar NULL,
	chainname varchar NULL,
	address varchar NULL,
	"type" varchar NULL,
	totalsupply float8 NULL,
	image varchar NULL,
	marketcap float8 NULL,
	volume24h float8 NULL,
	priceusd float8 NULL,
	createddate varchar NULL,
	updateddate varchar NULL,
	des varchar NULL
);