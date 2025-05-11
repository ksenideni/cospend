-- Пользователи
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email varchar(100) NOT NULL,
	phone_number varchar(20) NOT NULL,
	"name" varchar(250) NOT NULL,
	"password" varchar(150) NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	deleted_at timestamptz NULL,
	deleted_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT user_pkey PRIMARY KEY (id)
);

-- Группы (поездки, мероприятия и т.д.)
CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_by INTEGER REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Участники групп
CREATE TABLE IF NOT EXISTS group_members (
    id SERIAL PRIMARY KEY,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(group_id, user_id) -- нельзя дважды вступить в группу
);

-- Расходы
CREATE TABLE IF NOT EXISTS expenses (
    id SERIAL PRIMARY KEY,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    created_by INTEGER REFERENCES users(id) ON DELETE SET NULL,
    description TEXT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Доли расходов: кто сколько платил и кому сколько начислено
CREATE TABLE IF NOT EXISTS expense_shares (
    id SERIAL PRIMARY KEY,
    expense_id INTEGER NOT NULL REFERENCES expenses(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    paid NUMERIC(10, 2) NOT NULL DEFAULT 0,
    owed NUMERIC(10, 2) NOT NULL DEFAULT 0
);

-- Долги между пользователями внутри группы
CREATE TABLE IF NOT EXISTS debts (
    id SERIAL PRIMARY KEY,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    from_user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    to_user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount NUMERIC(10, 2) NOT NULL,
    UNIQUE(group_id, from_user_id, to_user_id)
);

-- Погашенные долги
CREATE TABLE IF NOT EXISTS settlements (
    id SERIAL PRIMARY KEY,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    from_user_id INTEGER NOT NULL REFERENCES users(id),
    to_user_id INTEGER NOT NULL REFERENCES users(id),
    amount NUMERIC(10, 2) NOT NULL,
    settled_at TIMESTAMP DEFAULT NOW()
);
