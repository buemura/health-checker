--- TABLES
CREATE UNLOGGED TABLE endpoints (
	id TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	url TEXT NOT NULL,
	status TEXT,
	check_frequency INTEGER NOT NULL,
	last_checked TIMESTAMP,
	notify_to TEXT NOT NULL
);

CREATE UNLOGGED TABLE notifications (
	id TEXT PRIMARY KEY,
	endpoint_id TEXT NOT NULL,
	destination TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	CONSTRAINT fk_endpoints_notifications_id
		FOREIGN KEY (endpoint_id) REFERENCES endpoints(id)
);
