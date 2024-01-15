CREATE TABLE services (
	id SERIAL PRIMARY KEY,
    service_name VARCHAR(255),
    status VARCHAR(255),
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    uptime VARCHAR(255),
    metadata JSONB
    );
CREATE OR REPLACE FUNCTION update_end_time_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.end_time = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_end_time_column
BEFORE UPDATE ON services
FOR EACH ROW
EXECUTE FUNCTION update_end_time_column();



CREATE TABLE services (
	service_id SERIAL PRIMARY KEY,
    service_name VARCHAR(255),
    status VARCHAR(255),
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    uptime interval,
    metadata JSONB,
    node_id INT references nodes(node_id)
    );
   
CREATE TABLE nodes (
	node_id SERIAL PRIMARY KEY,
    node_name VARCHAR(255),
    node_ip VARCHAR(255),
    metadata JSONB
);